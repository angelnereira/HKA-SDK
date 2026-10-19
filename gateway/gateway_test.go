package gateway_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/gateway"
	"github.com/angelnereira/hka-sdk/types"
)

func newServer() http.Handler {
	return gateway.NewHandler(hka.NewDemo()).Routes()
}

func do(t *testing.T, srv http.Handler, method, path, body string, headers map[string]string) (*httptest.ResponseRecorder, map[string]any) {
	t.Helper()
	var r *http.Request
	if body == "" {
		r = httptest.NewRequest(method, path, nil)
	} else {
		r = httptest.NewRequest(method, path, strings.NewReader(body))
	}
	for k, v := range headers {
		r.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	var out map[string]any
	if w.Body.Len() > 0 {
		_ = json.Unmarshal(w.Body.Bytes(), &out)
	}
	return w, out
}

func TestBuildOnly_ComputesTotals(t *testing.T) {
	srv := newServer()
	body := `{
      "tipo": "01",
      "numero": 1, "punto": 1,
      "cliente": {"tipo":"contribuyente","ruc":"155596713-2-2015","dv":"59","razonSocial":"Cliente S.A.","direccion":"Ave. La Paz"},
      "items": [{"descripcion":"Servicio","cantidad":1,"precioUnitario":100,"tasaITBMS":"01"}]
    }`
	w, _ := do(t, srv, "POST", "/v1/documents/build", body, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var doc types.DocumentoElectronico
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("failed to unmarshal built document: %v", err)
	}
	if doc.TotalesSubTotales.TotalFactura != "107.00" {
		t.Errorf("expected TotalFactura 107.00, got %v", doc.TotalesSubTotales.TotalFactura)
	}
}

func TestBuildOnly_AutoTasaITBMS(t *testing.T) {
	srv := newServer()
	body := `{
      "tipo": "01", "autoTasaITBMS": true,
      "cliente": {"tipo":"consumidorFinal","razonSocial":"Juan","direccion":"Calle 50"},
      "items": [{"descripcion":"Cerveza nacional","cantidad":1,"precioUnitario":2}]
    }`
	w, _ := do(t, srv, "POST", "/v1/documents/build", body, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var doc types.DocumentoElectronico
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("failed to unmarshal built document: %v", err)
	}
	if len(doc.ListaItems) == 0 || doc.ListaItems[0].TasaITBMS != "02" { // 10% for alcohol
		t.Errorf("expected inferred TasaITBMS 02, got %v", doc.ListaItems)
	}
}

func TestBuildOnly_FullHKASchema(t *testing.T) {
	srv := newServer()
	body := `{
		"codigoSucursalEmisor": "0000",
		"datosTransaccion": {
			"tipoEmision": "01",
			"tipoDocumento": "01",
			"numeroDocumentoFiscal": "0000000088",
			"puntoFacturacionFiscal": "001",
			"fechaEmision": "2026-10-18T10:00:00-05:00",
			"naturalezaOperacion": "01",
			"tipoOperacion": "1",
			"destinoOperacion": "1",
			"formatoCAFE": "3",
			"entregaCAFE": "3",
			"envioContenedor": "1",
			"procesoGeneracion": "1",
			"tipoVenta": "1",
			"cliente": {
				"tipoClienteFE": "01",
				"tipoContribuyente": "2",
				"numeroRUC": "155596713-2-2015",
				"digitoVerificadorRUC": "59",
				"razonSocial": "CLIENTE S.A.",
				"direccion": "Calle 50",
				"codigoUbicacion": "8-8-8",
				"pais": "PA"
			}
		},
		"listaItems": [
			{
				"descripcion": "Item directo",
				"cantidad": "2.000000",
				"precioUnitario": "25.000000",
				"tasaITBMS": "01"
			}
		]
	}`
	w, _ := do(t, srv, "POST", "/v1/documents/build", body, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var doc types.DocumentoElectronico
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if doc.TotalesSubTotales.TotalFactura != "53.50" {
		t.Errorf("expected 53.50, got %s", doc.TotalesSubTotales.TotalFactura)
	}
}

func TestBuildOnly_GobiernoDocument(t *testing.T) {
	srv := newServer()
	body := `{
		"tipo": "gobierno",
		"numero": 99,
		"punto": 1,
		"cliente": {
			"tipo": "gobierno",
			"ruc": "8-NT-1-13656",
			"dv": "43",
			"razonSocial": "MINISTERIO",
			"direccion": "Ancon"
		},
		"items": [
			{
				"descripcion": "Papeleria",
				"cantidad": 1,
				"precioUnitario": 100,
				"tasaITBMS": "01",
				"codigoCPBSAbrev": "13",
				"codigoCPBS": "1310",
				"unidadMedidaCPBS": "und"
			}
		]
	}`
	w, _ := do(t, srv, "POST", "/v1/documents/build", body, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	var doc types.DocumentoElectronico
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if doc.DatosTransaccion.Cliente.TipoClienteFE != types.ClienteGobierno {
		t.Errorf("expected gobierno client, got %s", doc.DatosTransaccion.Cliente.TipoClienteFE)
	}
}

func TestBuildOnly_ValidationError(t *testing.T) {
	srv := newServer()
	body := `{"tipo":"01","cliente":{"tipo":"consumidorFinal","razonSocial":"X","direccion":"Calle 50"},"items":[]}`
	w, out := do(t, srv, "POST", "/v1/documents/build", body, nil)
	if w.Code != http.StatusUnprocessableEntity && w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, body = %s", w.Code, w.Body.String())
	}
	if out["error"] == nil {
		t.Error("expected an error field")
	}
}

func TestBuildOnly_UnknownTipo(t *testing.T) {
	srv := newServer()
	body := `{"tipo":"99","cliente":{"tipo":"consumidorFinal","razonSocial":"X","direccion":"Calle 50"},"items":[]}`
	w, _ := do(t, srv, "POST", "/v1/documents/build", body, nil)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", w.Code)
	}
}

func TestSend_MissingCreds(t *testing.T) {
	srv := newServer()
	w, _ := do(t, srv, "POST", "/v1/documents/send", `{"tipo":"01"}`, nil)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", w.Code)
	}
}

func TestCatalogProvincias(t *testing.T) {
	srv := newServer()
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, httptest.NewRequest("GET", "/v1/catalog/provincias", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	var ps []any
	_ = json.Unmarshal(w.Body.Bytes(), &ps)
	if len(ps) != 13 {
		t.Errorf("expected 13 provincias, got %d", len(ps))
	}
}

func TestCatalogUbicacion(t *testing.T) {
	srv := newServer()
	w, out := do(t, srv, "GET", "/v1/catalog/ubicacion/8-8-7", "", nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if out["provincia"] != "Panamá" || out["corregimiento"] != "Bella Vista" {
		t.Errorf("unexpected resolve: %v", out)
	}
}

func TestCatalogCUFE(t *testing.T) {
	srv := newServer()
	cufe := "FE0120000155596713-2-2015-5900012019052800055000155650121566749040"
	w, out := do(t, srv, "GET", "/v1/catalog/cufe/"+cufe, "", nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", w.Code, w.Body.String())
	}
	if out["tipoDocumento"] != "01" || out["ambiente"] != "2" {
		t.Errorf("unexpected cufe parse: %v", out)
	}
}

func TestITBMSSuggest(t *testing.T) {
	srv := newServer()
	w, out := do(t, srv, "POST", "/v1/catalog/itbms/suggest", `{"descripcion":"Cartón de cigarrillos"}`, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if out["tasa"] != "03" { // 15% tobacco
		t.Errorf("expected tasa 03, got %v", out["tasa"])
	}
}

func TestHealthz(t *testing.T) {
	srv := newServer()
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, httptest.NewRequest("GET", "/healthz", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
}

// guard: ensure the build endpoint rejects malformed JSON.
func TestBuildOnly_BadJSON(t *testing.T) {
	srv := newServer()
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, httptest.NewRequest("POST", "/v1/documents/build", bytes.NewReader([]byte("{not json"))))
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", w.Code)
	}
}
