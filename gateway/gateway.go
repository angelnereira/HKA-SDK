// Package gateway exposes the HKA SDK as a language-neutral JSON/HTTP service.
//
// It wraps the safe-by-construction docbuilder, the validation layer, the catalog
// helpers and the SOAP client behind clean REST endpoints so that any language
// (TypeScript, JavaScript, Python, Java, …) can emit compliant documents without
// touching SOAP or recomputing totals. Credentials are passed per request via
// headers, preserving the SDK's stateless multi-tenant model.
//
// See docs/GATEWAY.md and openapi.yaml for the full contract.
package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/catalog"
	"github.com/angelnereira/hka-sdk/docbuilder"
	"github.com/angelnereira/hka-sdk/types"
	"github.com/angelnereira/hka-sdk/validate"
)

// Handler holds the dependencies for the gateway HTTP handlers.
type Handler struct {
	client *hka.Client
}

// NewHandler builds a gateway backed by the given HKA client.
func NewHandler(client *hka.Client) *Handler {
	return &Handler{client: client}
}

// Routes returns an http.Handler with every endpoint registered.
func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Document build / send.
	mux.HandleFunc("POST /v1/documents/build", h.buildOnly)
	mux.HandleFunc("POST /v1/documents/send", h.send)

	// Key-based operations.
	mux.HandleFunc("POST /v1/documents/status", h.status)
	mux.HandleFunc("POST /v1/documents/cancel", h.cancel)
	mux.HandleFunc("POST /v1/documents/xml", h.downloadXML)
	mux.HandleFunc("POST /v1/documents/pdf", h.downloadPDF)
	mux.HandleFunc("POST /v1/documents/email", h.sendEmail)
	mux.HandleFunc("POST /v1/email/track", h.trackEmail)
	mux.HandleFunc("GET /v1/folios", h.folios)
	mux.HandleFunc("POST /v1/ruc/query", h.queryRUC)

	// Catalog (no credentials required).
	mux.HandleFunc("GET /v1/catalog/provincias", h.provincias)
	mux.HandleFunc("GET /v1/catalog/ubicacion/{code}", h.ubicacion)
	mux.HandleFunc("GET /v1/catalog/cufe/{cufe}", h.cufe)
	mux.HandleFunc("GET /v1/catalog/cpbs/{codigo}", h.cpbs)
	mux.HandleFunc("POST /v1/catalog/itbms/suggest", h.itbmsSuggest)

	return mux
}

// --- document build / send ---

func (h *Handler) buildDocument(req *DocumentRequest) (*types.DocumentoElectronico, error) {
	ctor, ok := constructors[req.Tipo]
	if !ok {
		return nil, &badRequest{msg: "tipo must be one of 01..10"}
	}
	b := ctor()
	if req.Sucursal != "" {
		b.Sucursal(req.Sucursal)
	}
	if req.Numero > 0 {
		b.Numero(req.Numero)
	}
	if req.Punto > 0 {
		b.Punto(req.Punto)
	}
	if req.FechaEmision != nil {
		b.FechaEmision(*req.FechaEmision)
	}
	if req.TipoVenta != "" {
		b.TipoVenta(types.TipoVenta(req.TipoVenta))
	}
	if req.InformacionInteres != "" {
		b.InformacionInteres(req.InformacionInteres)
	}
	if req.AutoTasaITBMS {
		b.AutoTasaITBMS()
	}

	cli, err := req.Cliente.toCliente()
	if err != nil {
		return nil, &badRequest{msg: err.Error()}
	}
	b.Cliente(cli)

	for _, it := range req.Items {
		b.AddItem(it.toItem())
	}
	for _, fp := range req.FormasPago {
		b.AddFormaPago(types.FormaPago(fp.Forma), fp.Valor)
	}
	for _, pp := range req.PagosPlazo {
		b.AddPagoPlazo(pp.FechaVence, pp.Valor, pp.Descripcion)
	}
	if req.PagoContado != "" {
		b.PagoContado(types.FormaPago(req.PagoContado))
	}
	if req.Exportacion != nil {
		b.Exportacion(types.DatosExportacion{
			CondicionesEntrega:          types.Incoterm(req.Exportacion.CondicionesEntrega),
			MonedaOperExportacion:       types.CurrencyCode(req.Exportacion.Moneda),
			MonedaOperExportacionNonDef: req.Exportacion.MonedaNoDef,
			TipoDeCambio:                req.Exportacion.TipoDeCambio,
			MontoMonedaExtranjera:       req.Exportacion.MontoMonedaExtranjera,
			PuertoEmbarque:              req.Exportacion.PuertoEmbarque,
		})
	}
	for _, ref := range req.Referencias {
		b.Referencia(ref.CUFE, ref.FechaEmision)
	}
	if req.Retencion != nil {
		b.Retencion(types.Retencion{
			CodigoRetencion: types.CodigoRetencion(req.Retencion.Codigo),
			FechaRetencion:  req.Retencion.Fecha,
			MontoRetencion:  req.Retencion.Monto,
		})
	}
	return b.Build()
}

func (h *Handler) buildOnly(w http.ResponseWriter, r *http.Request) {
	var req DocumentRequest
	if !decode(w, r, &req) {
		return
	}
	doc, err := h.buildDocument(&req)
	if err != nil {
		writeBuildError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, doc)
}

func (h *Handler) send(w http.ResponseWriter, r *http.Request) {
	creds, ok := credsFrom(w, r)
	if !ok {
		return
	}
	var req DocumentRequest
	if !decode(w, r, &req) {
		return
	}
	doc, err := h.buildDocument(&req)
	if err != nil {
		writeBuildError(w, err)
		return
	}
	resp, err := h.client.Send(r.Context(), creds, doc)
	if err != nil {
		writeClientError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

// --- key-based operations ---

func (k KeyRequest) toKey() hka.DocumentKey {
	return hka.DocumentKey{
		CodigoSucursalEmisor:   k.Sucursal,
		NumeroDocumentoFiscal:  k.Numero,
		PuntoFacturacionFiscal: k.Punto,
		TipoDocumento:          types.TipoDocumento(k.TipoDocumento),
		TipoEmision:            types.TipoEmision(k.TipoEmision),
	}
}

func (h *Handler) status(w http.ResponseWriter, r *http.Request) {
	h.withKey(w, r, func(ctx context.Context, creds hka.Credentials, key hka.DocumentKey, _ map[string]any) (any, error) {
		return h.client.DocumentStatus(ctx, creds, key)
	})
}

func (h *Handler) cancel(w http.ResponseWriter, r *http.Request) {
	h.withKey(w, r, func(ctx context.Context, creds hka.Credentials, key hka.DocumentKey, extra map[string]any) (any, error) {
		reason, _ := extra["motivo"].(string)
		return h.client.Cancel(ctx, creds, key, reason)
	})
}

func (h *Handler) downloadXML(w http.ResponseWriter, r *http.Request) {
	h.withKey(w, r, func(ctx context.Context, creds hka.Credentials, key hka.DocumentKey, _ map[string]any) (any, error) {
		return h.client.DownloadXML(ctx, creds, key)
	})
}

func (h *Handler) downloadPDF(w http.ResponseWriter, r *http.Request) {
	h.withKey(w, r, func(ctx context.Context, creds hka.Credentials, key hka.DocumentKey, extra map[string]any) (any, error) {
		serial, _ := extra["serial"].(string)
		return h.client.DownloadPDF(ctx, creds, key, serial)
	})
}

func (h *Handler) sendEmail(w http.ResponseWriter, r *http.Request) {
	h.withKey(w, r, func(ctx context.Context, creds hka.Credentials, key hka.DocumentKey, extra map[string]any) (any, error) {
		email, _ := extra["email"].(string)
		return h.client.SendEmail(ctx, creds, key, email)
	})
}

func (h *Handler) trackEmail(w http.ResponseWriter, r *http.Request) {
	creds, ok := credsFrom(w, r)
	if !ok {
		return
	}
	var body struct {
		CUFE string `json:"cufe"`
	}
	if !decode(w, r, &body) {
		return
	}
	resp, err := h.client.TrackEmail(r.Context(), creds, body.CUFE)
	if err != nil {
		writeClientError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) folios(w http.ResponseWriter, r *http.Request) {
	creds, ok := credsFrom(w, r)
	if !ok {
		return
	}
	resp, err := h.client.RemainingFolios(r.Context(), creds)
	if err != nil {
		writeClientError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) queryRUC(w http.ResponseWriter, r *http.Request) {
	creds, ok := credsFrom(w, r)
	if !ok {
		return
	}
	var body struct {
		Tipo string `json:"tipo"`
		RUC  string `json:"ruc"`
	}
	if !decode(w, r, &body) {
		return
	}
	resp, err := h.client.QueryRUC(r.Context(), creds, types.RUCType(body.Tipo), body.RUC)
	if err != nil {
		writeClientError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

// keyOp runs an operation that needs credentials plus a document key. The request
// body carries the key plus any extra fields (motivo, serial, email) flattened in.
type keyOp func(ctx context.Context, creds hka.Credentials, key hka.DocumentKey, extra map[string]any) (any, error)

func (h *Handler) withKey(w http.ResponseWriter, r *http.Request, op keyOp) {
	creds, ok := credsFrom(w, r)
	if !ok {
		return
	}
	var body struct {
		Key   KeyRequest     `json:"key"`
		Extra map[string]any `json:"-"`
	}
	// Decode into a generic map to capture both the key and the extra fields.
	var raw map[string]json.RawMessage
	if !decode(w, r, &raw) {
		return
	}
	if keyRaw, ok := raw["key"]; ok {
		if err := json.Unmarshal(keyRaw, &body.Key); err != nil {
			writeJSON(w, http.StatusBadRequest, errResp("invalid key: "+err.Error()))
			return
		}
	}
	extra := map[string]any{}
	for k, v := range raw {
		if k == "key" {
			continue
		}
		var val any
		_ = json.Unmarshal(v, &val)
		extra[k] = val
	}
	resp, err := op(r.Context(), creds, body.Key.toKey(), extra)
	if err != nil {
		writeClientError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

// --- catalog ---

func (h *Handler) provincias(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, catalog.Provincias())
}

func (h *Handler) ubicacion(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	u, err := catalog.ParseUbicacion(code)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, errResp(err.Error()))
		return
	}
	prov, dist, corr, ok := u.Resolve()
	writeJSON(w, http.StatusOK, map[string]any{
		"codigo":        u.String(),
		"provincia":     prov,
		"distrito":      dist,
		"corregimiento": corr,
		"conocido":      ok,
	})
}

func (h *Handler) cufe(w http.ResponseWriter, r *http.Request) {
	info, err := catalog.ParseCUFE(r.PathValue("cufe"))
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, errResp(err.Error()))
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"raw":           info.Raw,
		"tipoDocumento": string(info.TipoDocumento),
		"ambiente":      string(info.Ambiente),
		"ambienteDesc":  info.Ambiente.Descripcion(),
	})
}

func (h *Handler) cpbs(w http.ResponseWriter, r *http.Request) {
	codigo := r.PathValue("codigo")
	p, ok := catalog.CPBSByCodigo(codigo)
	abrev, _ := catalog.AbrevForCPBS(codigo)
	writeJSON(w, http.StatusOK, map[string]any{
		"codigo":   codigo,
		"abrev":    abrev,
		"nombre":   p.Nombre,
		"conocido": ok,
	})
}

func (h *Handler) itbmsSuggest(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Descripcion string `json:"descripcion"`
	}
	if !decode(w, r, &body) {
		return
	}
	cat := catalog.SugerirCategoria(body.Descripcion)
	writeJSON(w, http.StatusOK, map[string]any{
		"descripcion": body.Descripcion,
		"tasa":        string(cat.Tasa()),
		"porcentaje":  cat.Porcentaje(),
		"categoria":   cat.Descripcion(),
	})
}

// --- helpers ---

var constructors = map[string]func() *docbuilder.Builder{
	"01": docbuilder.NewFacturaInterna,
	"02": docbuilder.NewFacturaImportacion,
	"03": docbuilder.NewFacturaExportacion,
	"04": docbuilder.NewNotaCreditoReferenciada,
	"05": docbuilder.NewNotaDebitoReferenciada,
	"06": docbuilder.NewNotaCreditoGenerica,
	"07": docbuilder.NewNotaDebitoGenerica,
	"08": docbuilder.NewFacturaZonaFranca,
	"09": docbuilder.NewReembolso,
	"10": docbuilder.NewFacturaExtranjera,
}

type badRequest struct{ msg string }

func (e *badRequest) Error() string { return e.msg }

func credsFrom(w http.ResponseWriter, r *http.Request) (hka.Credentials, bool) {
	c := hka.Credentials{
		TokenEmpresa:  r.Header.Get("X-HKA-Token-Empresa"),
		TokenPassword: r.Header.Get("X-HKA-Token-Password"),
	}
	if c.TokenEmpresa == "" || c.TokenPassword == "" {
		writeJSON(w, http.StatusUnauthorized, errResp("missing X-HKA-Token-Empresa / X-HKA-Token-Password headers"))
		return hka.Credentials{}, false
	}
	return c, true
}

func decode(w http.ResponseWriter, r *http.Request, dst any) bool {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(dst); err != nil {
		writeJSON(w, http.StatusBadRequest, errResp("invalid JSON body: "+err.Error()))
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(v)
}

func errResp(msg string) map[string]any {
	return map[string]any{"error": msg}
}

func writeBuildError(w http.ResponseWriter, err error) {
	var ve *validate.ValidationError
	if errors.As(err, &ve) {
		fields := make([]map[string]string, 0, len(ve.Fields))
		for _, fe := range ve.Fields {
			fields = append(fields, map[string]string{"field": fe.Field, "message": fe.Message})
		}
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{
			"error":  "validation failed",
			"fields": fields,
		})
		return
	}
	writeJSON(w, http.StatusBadRequest, errResp(err.Error()))
}

func writeClientError(w http.ResponseWriter, err error) {
	var ve *validate.ValidationError
	if errors.As(err, &ve) {
		writeBuildError(w, err)
		return
	}
	var hkaErr *hka.HKAError
	if errors.As(err, &hkaErr) {
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"error":   "hka service error",
			"code":    string(hkaErr.Code),
			"message": hkaErr.Message,
		})
		return
	}
	var netErr *hka.NetworkError
	if errors.As(err, &netErr) {
		writeJSON(w, http.StatusBadGateway, errResp("network error: "+netErr.Error()))
		return
	}
	writeJSON(w, http.StatusInternalServerError, errResp(err.Error()))
}
