package docbuilder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/angelnereira/hka-sdk/types"
	"github.com/angelnereira/hka-sdk/validate"
)

// DynamicDocument represents a flexible, dynamic payload structure that can be unmarshaled
// from simplified JSON, webhooks, ERP integrations, or REST requests.
type DynamicDocument struct {
	Tipo               string              `json:"tipo"`                    // 01..10, or "interna", "exportacion", "nc_ref", "gobierno", etc.
	TipoDocumento      string              `json:"tipoDocumento,omitempty"` // alias for tipo
	Sucursal           any                 `json:"sucursal,omitempty"`      // int or string (e.g. 0 or "0000")
	Numero             any                 `json:"numero,omitempty"`        // int, float, or string (e.g. 100 or "0000000100")
	Punto              any                 `json:"punto,omitempty"`         // int or string (e.g. 1 or "001")
	FechaEmision       any                 `json:"fechaEmision,omitempty"`  // string (RFC3339/ISO) or time.Time
	TipoVenta          string              `json:"tipoVenta,omitempty"`
	InformacionInteres string              `json:"informacionInteres,omitempty"`
	AutoTasaITBMS      bool                `json:"autoTasaITBMS,omitempty"`
	Cliente            DynamicCliente      `json:"cliente"`
	Items              []DynamicItem       `json:"items"`
	FormasPago         []DynamicFormaPago  `json:"formasPago,omitempty"`
	PagosPlazo         []DynamicPagoPlazo  `json:"pagosPlazo,omitempty"`
	PagoContado        string              `json:"pagoContado,omitempty"`
	Exportacion        *DynamicExportacion `json:"exportacion,omitempty"`
	Referencias        []DynamicReferencia `json:"referencias,omitempty"`
	Retencion          *DynamicRetencion   `json:"retencion,omitempty"`
	DescuentoGlobal    *DynamicDescuento   `json:"descuentoGlobal,omitempty"`
	Autorizados        []DynamicAutorizado `json:"autorizados,omitempty"`
	PedidoGlobal       *DynamicPedido      `json:"pedidoGlobal,omitempty"`
}

type DynamicCliente struct {
	Tipo               string `json:"tipo"`                        // "01", "contribuyente", "02", "consumidorFinal", "03", "gobierno", "04", "extranjero"
	TipoClienteFE      string `json:"tipoClienteFE,omitempty"`     // alias
	TipoContribuyente  string `json:"tipoContribuyente,omitempty"` // "1": Natural, "2": Jurídico
	RUC                string `json:"ruc,omitempty"`
	NumeroRUC          string `json:"numeroRUC,omitempty"` // alias
	DV                 any    `json:"dv,omitempty"`        // int or string
	DigitoVerifRUC     any    `json:"digitoVerificadorRUC,omitempty"`
	RazonSocial        string `json:"razonSocial,omitempty"`
	Nombre             string `json:"nombre,omitempty"` // alias
	Direccion          string `json:"direccion,omitempty"`
	CodigoUbicacion    string `json:"codigoUbicacion,omitempty"` // e.g. "8-8-8"
	Provincia          string `json:"provincia,omitempty"`
	Distrito           string `json:"distrito,omitempty"`
	Corregimiento      string `json:"corregimiento,omitempty"`
	TipoIdentificacion string `json:"tipoIdentificacion,omitempty"` // "01", "02", "99"
	NroIdentificacion  string `json:"nroIdentificacion,omitempty"`
	PaisExtranjero     string `json:"paisExtranjero,omitempty"`
	Telefono           string `json:"telefono,omitempty"`
	Telefono1          string `json:"telefono1,omitempty"`
	Correo             string `json:"correo,omitempty"`
	CorreoElectronico1 string `json:"correoElectronico1,omitempty"`
	Pais               string `json:"pais,omitempty"`     // default "PA"
	PaisOtro           string `json:"paisOtro,omitempty"` // when Pais = "ZZ"
}

type DynamicItem struct {
	Descripcion             string  `json:"descripcion"`
	Codigo                  string  `json:"codigo,omitempty"`
	UnidadMedida            string  `json:"unidadMedida,omitempty"`
	Cantidad                any     `json:"cantidad"`                          // float, int, or string
	PrecioUnitario          any     `json:"precioUnitario"`                    // float, int, or string
	Descuento               any     `json:"descuento,omitempty"`               // float, int, or string
	PrecioUnitarioDescuento any     `json:"precioUnitarioDescuento,omitempty"` // alias
	TasaITBMS               string  `json:"tasaITBMS,omitempty"`               // "00", "01", "02", "03"
	TasaISC                 string  `json:"tasaISC,omitempty"`
	ValorISC                any     `json:"valorISC,omitempty"`
	CodigoCPBSAbrev         string  `json:"codigoCPBSAbrev,omitempty"`
	CodigoCPBS              string  `json:"codigoCPBS,omitempty"`
	UnidadMedidaCPBS        string  `json:"unidadMedidaCPBS,omitempty"`
	FechaFabricacion        string  `json:"fechaFabricacion,omitempty"`
	FechaCaducidad          string  `json:"fechaCaducidad,omitempty"`
	InfoItem                string  `json:"infoItem,omitempty"`
	Medicina                *Med    `json:"medicina,omitempty"`
	Vehiculo                *Veh    `json:"vehiculo,omitempty"`
	PedidoItem              *PedItm `json:"pedidoItem,omitempty"`
}

type Med struct {
	NroLote           string `json:"nroLote"`
	CantProductosLote string `json:"cantProductosLote"`
}

type Veh struct {
	Marca           string `json:"marca,omitempty"`
	Modelo          string `json:"modelo,omitempty"`
	Anio            string `json:"anio,omitempty"`
	Color           string `json:"color,omitempty"`
	NumeroVIN       string `json:"numeroVIN,omitempty"`
	NumeroMotor     string `json:"numeroMotor,omitempty"`
	NumeroPlaca     string `json:"numeroPlaca,omitempty"`
	TipoCombustible string `json:"tipoCombustible,omitempty"`
	PrecioBase      string `json:"precioBase,omitempty"`
	NumeroDocAduana string `json:"numeroDocAduana,omitempty"`
}

type PedItm struct {
	NumeroPedido string `json:"numeroPedido"`
	NroLinea     string `json:"nroLinea"`
}

type DynamicFormaPago struct {
	Forma            string `json:"forma"` // "01".."09", "99"
	FormaPagoFact    string `json:"formaPagoFact,omitempty"`
	DescFormaPago    string `json:"descFormaPago,omitempty"`
	Valor            any    `json:"valor"` // float or string
	ValorCuotaPagada any    `json:"valorCuotaPagada,omitempty"`
}

type DynamicPagoPlazo struct {
	FechaVence       any    `json:"fechaVence"` // string or time.Time
	FechaPago        any    `json:"fechaPago,omitempty"`
	Valor            any    `json:"valor"` // float or string
	ValorCuotaPlazo  any    `json:"valorCuotaPlazo,omitempty"`
	Descripcion      string `json:"descripcion,omitempty"`
	DescripcionCuota string `json:"descripcionCuota,omitempty"`
}

type DynamicExportacion struct {
	CondicionesEntrega    string `json:"condicionesEntrega"` // Incoterm
	Moneda                string `json:"moneda"`             // ISO currency
	MonedaOperExportacion string `json:"monedaOperExportacion,omitempty"`
	MonedaNoDef           string `json:"monedaNoDef,omitempty"`
	TipoDeCambio          any    `json:"tipoDeCambio,omitempty"`
	MontoMonedaExtranjera any    `json:"montoMonedaExtranjera,omitempty"`
	PuertoEmbarque        string `json:"puertoEmbarque,omitempty"`
}

type DynamicReferencia struct {
	CUFE                              string `json:"cufe"`
	CufeFEReferenciada                string `json:"cufeFEReferenciada,omitempty"`
	FechaEmision                      any    `json:"fechaEmision"`
	FechaEmisionDocFiscalReferenciado any    `json:"fechaEmisionDocFiscalReferenciado,omitempty"`
	NroFacturaPapel                   string `json:"nroFacturaPapel,omitempty"`
	NroFacturaImpFiscal               string `json:"nroFacturaImpFiscal,omitempty"`
}

type DynamicRetencion struct {
	Codigo          string `json:"codigo"`
	CodigoRetencion string `json:"codigoRetencion,omitempty"`
	Fecha           string `json:"fecha,omitempty"`
	FechaRetencion  string `json:"fechaRetencion,omitempty"`
	Monto           any    `json:"monto"`
	MontoRetencion  any    `json:"montoRetencion,omitempty"`
}

type DynamicDescuento struct {
	Monto                any    `json:"monto"`
	MontoDescuento       any    `json:"montoDescuento,omitempty"`
	Descripcion          string `json:"descripcion,omitempty"`
	DescripcionDescuento string `json:"descripcionDescuento,omitempty"`
}

type DynamicAutorizado struct {
	TipoContribuyente string `json:"tipoContribuyente"`
	Tipo              string `json:"tipo,omitempty"`
	RUC               string `json:"ruc"`
	RUCReceptor       string `json:"rucReceptor,omitempty"`
	DV                any    `json:"dv"`
	DigitoVerifRUC    any    `json:"digitoVerifRUC,omitempty"`
	Nombre            string `json:"nombre,omitempty"`
	NombreAutorizado  string `json:"nombreAutorizado,omitempty"`
}

type DynamicPedido struct {
	Numero       string `json:"numero"`
	NumeroPedido string `json:"numeroPedido,omitempty"`
	Fecha        any    `json:"fecha"`
	FechaPedido  any    `json:"fechaPedido,omitempty"`
	Monto        any    `json:"monto"`
	MontoTotal   any    `json:"montoTotal,omitempty"`
	Moneda       string `json:"moneda,omitempty"`
	MonedaPedido string `json:"monedaPedido,omitempty"`
}

// FromPayload accepts any supported payload type (JSON bytes, JSON string, io.Reader,
// map[string]any, or typed document) and converts it into a validated *types.DocumentoElectronico.
func FromPayload(v any) (*types.DocumentoElectronico, error) {
	switch val := v.(type) {
	case *types.DocumentoElectronico:
		if ve := validate.Document(val); ve != nil {
			return nil, ve
		}
		return val, nil
	case types.DocumentoElectronico:
		if ve := validate.Document(&val); ve != nil {
			return nil, ve
		}
		return &val, nil
	case []byte:
		return FromJSON(val)
	case string:
		return FromJSON([]byte(val))
	case io.Reader:
		data, err := io.ReadAll(val)
		if err != nil {
			return nil, fmt.Errorf("docbuilder: error reading payload stream: %w", err)
		}
		return FromJSON(data)
	case map[string]any:
		return FromMap(val)
	default:
		// Attempt JSON round-trip
		data, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("docbuilder: unable to marshal payload: %w", err)
		}
		return FromJSON(data)
	}
}

// FromMap converts a map[string]any into a validated *types.DocumentoElectronico.
func FromMap(m map[string]any) (*types.DocumentoElectronico, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("docbuilder: error marshaling map: %w", err)
	}
	return FromJSON(data)
}

// FromJSON parses arbitrary JSON into a validated *types.DocumentoElectronico.
// It automatically recognizes whether the input is a full HKA WCF/SOAP JSON document
// or a simplified dynamic payload.
func FromJSON(data []byte) (*types.DocumentoElectronico, error) {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return nil, fmt.Errorf("docbuilder: payload is empty")
	}

	// 1. Detect if this is already a full HKA DocumentoElectronico structure
	var peek map[string]json.RawMessage
	if err := json.Unmarshal(trimmed, &peek); err != nil {
		return nil, fmt.Errorf("docbuilder: invalid JSON payload: %w", err)
	}

	if _, hasDT := peek["datosTransaccion"]; hasDT {
		return parseFullDocumentJSON(trimmed)
	}
	if _, hasDT := peek["DatosTransaccion"]; hasDT {
		return parseFullDocumentJSON(trimmed)
	}

	// 2. Otherwise parse as simplified DynamicDocument
	var dyn DynamicDocument
	if err := json.Unmarshal(trimmed, &dyn); err != nil {
		return nil, fmt.Errorf("docbuilder: error unmarshaling dynamic document: %w", err)
	}

	return dyn.Build()
}

func parseFullDocumentJSON(data []byte) (*types.DocumentoElectronico, error) {
	var doc types.DocumentoElectronico
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("docbuilder: error unmarshaling full document JSON: %w", err)
	}

	// Apply sensible defaults if omitted
	if doc.CodigoSucursalEmisor == "" {
		doc.CodigoSucursalEmisor = "0000"
	}
	if doc.DatosTransaccion.ProcesoGeneracion == "" {
		doc.DatosTransaccion.ProcesoGeneracion = "1"
	}
	if doc.DatosTransaccion.Cliente.Pais == "" {
		doc.DatosTransaccion.Cliente.Pais = types.CountryPA
	}
	if doc.DatosTransaccion.Cliente.Pais == types.CountryPA &&
		doc.DatosTransaccion.Cliente.TipoClienteFE != types.ClienteExtranjero &&
		doc.DatosTransaccion.Cliente.CodigoUbicacion == "" {
		doc.DatosTransaccion.Cliente.CodigoUbicacion = "8-8-8"
	}

	// If totals are missing or empty, calculate them from items
	if doc.TotalesSubTotales.TotalFactura == "" && len(doc.ListaItems) > 0 {
		var sumNeto, sumITBMS, sumISC, sumTodos float64
		for i := range doc.ListaItems {
			it := &doc.ListaItems[i]
			cant, _ := strconv.ParseFloat(it.Cantidad, 64)
			pu, _ := strconv.ParseFloat(it.PrecioUnitario, 64)
			desc, _ := strconv.ParseFloat(it.PrecioUnitarioDescuento, 64)
			if it.PrecioItem == "" {
				it.PrecioItem = fmtAmount(cant*(pu-desc), itemDecimals)
			}
			pi, _ := strconv.ParseFloat(it.PrecioItem, 64)
			if it.ValorITBMS == "" {
				it.ValorITBMS = fmtAmount(pi*it.TasaITBMS.Rate(), itemDecimals)
			}
			vi, _ := strconv.ParseFloat(it.ValorITBMS, 64)
			isc, _ := strconv.ParseFloat(it.ValorISC, 64)
			if it.ValorTotal == "" {
				it.ValorTotal = fmtAmount(pi+vi+isc, itemDecimals)
			}
			vt, _ := strconv.ParseFloat(it.ValorTotal, 64)

			sumNeto += pi
			sumITBMS += vi
			sumISC += isc
			sumTodos += vt
		}

		tot := &doc.TotalesSubTotales
		tot.TotalPrecioNeto = fmtAmount(sumNeto, totalDecimals)
		tot.TotalITBMS = fmtAmount(sumITBMS, totalDecimals)
		tot.TotalMontoGravado = fmtAmount(sumITBMS, totalDecimals)
		if sumISC > 0 {
			tot.TotalISC = fmtAmount(sumISC, totalDecimals)
		}
		totalFact := sumNeto + sumITBMS + sumISC
		tot.TotalFactura = fmtAmount(totalFact, totalDecimals)
		tot.NroItems = len(doc.ListaItems)
		tot.TotalTodosItems = fmtAmount(sumTodos, totalDecimals)

		if tot.TiempoPago == "" {
			if len(tot.ListaPagoPlazo) > 0 {
				tot.TiempoPago = types.PagoPlazo
			} else {
				tot.TiempoPago = types.PagoInmediato
				if len(tot.ListaFormaPago) == 0 {
					tot.ListaFormaPago = []types.FormaPagoItem{{
						FormaPagoFact:    types.PagoEfectivo,
						ValorCuotaPagada: fmtAmount(totalFact, totalDecimals),
					}}
				}
				tot.TotalValorRecibido = fmtAmount(totalFact, totalDecimals)
			}
		}
	}

	if ve := validate.Document(&doc); ve != nil {
		return nil, ve
	}
	return &doc, nil
}

// Build converts a DynamicDocument into a validated *types.DocumentoElectronico using Builder.
func (d *DynamicDocument) Build() (*types.DocumentoElectronico, error) {
	tipo := d.resolveTipo()
	ctor, ok := constructorMap[tipo]
	if !ok {
		return nil, fmt.Errorf("docbuilder: unsupported or invalid tipo %q (must be 01..10)", d.Tipo)
	}

	b := ctor()

	// Sucursal
	if d.Sucursal != nil {
		b.Sucursal(toString(d.Sucursal))
	}

	// Numero
	if d.Numero != nil {
		if n := toInt64(d.Numero); n > 0 {
			b.Numero(n)
		}
	}

	// Punto
	if d.Punto != nil {
		if p := int(toInt64(d.Punto)); p > 0 {
			b.Punto(p)
		}
	}

	// FechaEmision
	if d.FechaEmision != nil {
		if t, err := parseTime(d.FechaEmision); err == nil {
			b.FechaEmision(t)
		}
	}

	// TipoVenta
	if d.TipoVenta != "" {
		b.TipoVenta(types.TipoVenta(d.TipoVenta))
	}

	// InformacionInteres
	if d.InformacionInteres != "" {
		b.InformacionInteres(d.InformacionInteres)
	}

	// AutoTasaITBMS
	if d.AutoTasaITBMS {
		b.AutoTasaITBMS()
	}

	// Cliente
	cli, err := d.Cliente.toCliente()
	if err != nil {
		return nil, err
	}
	b.Cliente(cli)

	// Items
	if len(d.Items) == 0 {
		return nil, fmt.Errorf("docbuilder: items list must contain at least one item")
	}
	for i, it := range d.Items {
		item, err := it.toItem(i)
		if err != nil {
			return nil, err
		}
		b.AddItem(item)
	}

	// FormasPago
	for _, fp := range d.FormasPago {
		forma := fp.Forma
		if forma == "" {
			forma = fp.FormaPagoFact
		}
		if forma == "" {
			forma = string(types.PagoEfectivo)
		}
		val := toFloat64(fp.Valor)
		if val == 0 && fp.ValorCuotaPagada != nil {
			val = toFloat64(fp.ValorCuotaPagada)
		}
		b.AddFormaPago(types.FormaPago(forma), val)
	}

	// PagosPlazo
	for _, pp := range d.PagosPlazo {
		fv := pp.FechaVence
		if fv == nil {
			fv = pp.FechaPago
		}
		t, err := parseTime(fv)
		if err != nil {
			t = time.Now().Add(30 * 24 * time.Hour)
		}
		val := toFloat64(pp.Valor)
		if val == 0 && pp.ValorCuotaPlazo != nil {
			val = toFloat64(pp.ValorCuotaPlazo)
		}
		desc := pp.Descripcion
		if desc == "" {
			desc = pp.DescripcionCuota
		}
		b.AddPagoPlazo(t, val, desc)
	}

	// PagoContado
	if d.PagoContado != "" {
		b.PagoContado(types.FormaPago(d.PagoContado))
	}

	// Exportacion
	if d.Exportacion != nil {
		b.Exportacion(types.DatosExportacion{
			CondicionesEntrega:          types.Incoterm(d.Exportacion.CondicionesEntrega),
			MonedaOperExportacion:       types.CurrencyCode(firstNonEmpty(d.Exportacion.Moneda, d.Exportacion.MonedaOperExportacion)),
			MonedaOperExportacionNonDef: d.Exportacion.MonedaNoDef,
			TipoDeCambio:                toString(d.Exportacion.TipoDeCambio),
			MontoMonedaExtranjera:       toString(d.Exportacion.MontoMonedaExtranjera),
			PuertoEmbarque:              d.Exportacion.PuertoEmbarque,
		})
	}

	// Referencias
	for _, ref := range d.Referencias {
		cufe := firstNonEmpty(ref.CUFE, ref.CufeFEReferenciada)
		fe := ref.FechaEmision
		if fe == nil {
			fe = ref.FechaEmisionDocFiscalReferenciado
		}
		t, _ := parseTime(fe)
		b.Referencia(cufe, t)
	}

	// Retencion
	if d.Retencion != nil {
		b.Retencion(types.Retencion{
			CodigoRetencion: types.CodigoRetencion(firstNonEmpty(d.Retencion.Codigo, d.Retencion.CodigoRetencion)),
			FechaRetencion:  firstNonEmpty(d.Retencion.Fecha, d.Retencion.FechaRetencion),
			MontoRetencion:  toString(firstNonEmptyAny(d.Retencion.Monto, d.Retencion.MontoRetencion)),
		})
	}

	// DescuentoGlobal
	if d.DescuentoGlobal != nil {
		monto := toFloat64(firstNonEmptyAny(d.DescuentoGlobal.Monto, d.DescuentoGlobal.MontoDescuento))
		desc := firstNonEmpty(d.DescuentoGlobal.Descripcion, d.DescuentoGlobal.DescripcionDescuento)
		b.DescuentoGlobal(monto, desc)
	}

	// Autorizados
	for _, aut := range d.Autorizados {
		tipo := firstNonEmpty(aut.Tipo, aut.TipoContribuyente)
		ruc := firstNonEmpty(aut.RUC, aut.RUCReceptor)
		dv := toString(firstNonEmptyAny(aut.DV, aut.DigitoVerifRUC))
		nombre := firstNonEmpty(aut.Nombre, aut.NombreAutorizado)
		b.AutorizadoDescarga(tipo, ruc, dv, nombre)
	}

	// PedidoGlobal
	if d.PedidoGlobal != nil {
		num := firstNonEmpty(d.PedidoGlobal.Numero, d.PedidoGlobal.NumeroPedido)
		fecha, _ := parseTime(firstNonEmptyAny(d.PedidoGlobal.Fecha, d.PedidoGlobal.FechaPedido))
		monto := toFloat64(firstNonEmptyAny(d.PedidoGlobal.Monto, d.PedidoGlobal.MontoTotal))
		moneda := firstNonEmpty(d.PedidoGlobal.Moneda, d.PedidoGlobal.MonedaPedido)
		b.PedidoGlobal(num, fecha, monto, moneda)
	}

	return b.Build()
}

func (d *DynamicDocument) resolveTipo() string {
	raw := firstNonEmpty(d.Tipo, d.TipoDocumento)
	raw = strings.TrimSpace(strings.ToLower(raw))
	switch raw {
	case "01", "1", "interna", "factura_interna", "factura":
		if d.Cliente.isGobierno() {
			return "gobierno"
		}
		return "01"
	case "02", "2", "importacion", "factura_importacion":
		return "02"
	case "03", "3", "exportacion", "factura_exportacion":
		return "03"
	case "04", "4", "nc_ref", "nota_credito_ref", "nota_credito_referenciada":
		return "04"
	case "05", "5", "nd_ref", "nota_debito_ref", "nota_debito_referenciada":
		return "05"
	case "06", "6", "nc_gen", "nota_credito_gen", "nota_credito_generica":
		return "06"
	case "07", "7", "nd_gen", "nota_debito_gen", "nota_debito_generica":
		return "07"
	case "08", "8", "zona_franca", "factura_zona_franca":
		return "08"
	case "09", "9", "reembolso":
		return "09"
	case "10", "extranjera", "factura_extranjera":
		return "10"
	case "gobierno", "factura_gobierno":
		return "gobierno"
	case "consumidor_final":
		return "consumidor_final"
	default:
		return raw
	}
}

var constructorMap = map[string]func() *Builder{
	"01":               NewFacturaInterna,
	"02":               NewFacturaImportacion,
	"03":               NewFacturaExportacion,
	"04":               NewNotaCreditoReferenciada,
	"05":               NewNotaDebitoReferenciada,
	"06":               NewNotaCreditoGenerica,
	"07":               NewNotaDebitoGenerica,
	"08":               NewFacturaZonaFranca,
	"09":               NewReembolso,
	"10":               NewFacturaExtranjera,
	"gobierno":         NewFacturaGobierno,
	"consumidor_final": NewFacturaConsumidorFinal,
}

func (c *DynamicCliente) isGobierno() bool {
	t := strings.TrimSpace(strings.ToLower(firstNonEmpty(c.Tipo, c.TipoClienteFE)))
	return t == "03" || t == "gobierno"
}

func (c *DynamicCliente) toCliente() (types.Cliente, error) {
	tipo := strings.TrimSpace(strings.ToLower(firstNonEmpty(c.Tipo, c.TipoClienteFE)))
	ruc := firstNonEmpty(c.RUC, c.NumeroRUC)
	dv := toString(firstNonEmptyAny(c.DV, c.DigitoVerifRUC))
	nombre := firstNonEmpty(c.RazonSocial, c.Nombre)
	direccion := c.Direccion
	pais := firstNonEmpty(c.Pais, "PA")
	ubicacion := firstNonEmpty(c.CodigoUbicacion, "8-8-8")

	var cli types.Cliente
	switch tipo {
	case "01", "contribuyente", "":
		cli = ClienteContribuyente(ruc, dv, nombre, direccion)
	case "contribuyente_natural", "contribuyentenatural", "natural":
		cli = ClienteContribuyenteNatural(ruc, dv, nombre, direccion)
	case "02", "consumidor_final", "consumidorfinal", "final":
		cli = ClienteConsumidorFinal(nombre, direccion)
	case "03", "gobierno":
		cli = ClienteGobierno(ruc, dv, nombre, direccion)
	case "04", "extranjero":
		tipoID := types.TipoIdentificacion(c.TipoIdentificacion)
		if tipoID == "" {
			tipoID = types.IdPasaporte
		}
		cli = ClienteExtranjero(nombre, direccion, tipoID, c.NroIdentificacion, types.CountryCode(pais))
		if pais == "ZZ" && c.PaisOtro != "" {
			cli.PaisOtro = c.PaisOtro
		}
		return cli, nil
	default:
		// Attempt numeric code fallback
		switch types.TipoClienteFE(tipo) {
		case types.ClienteContribuyente:
			cli = ClienteContribuyente(ruc, dv, nombre, direccion)
		case types.ClienteConsumidorFinal:
			cli = ClienteConsumidorFinal(nombre, direccion)
		case types.ClienteGobierno:
			cli = ClienteGobierno(ruc, dv, nombre, direccion)
		case types.ClienteExtranjero:
			cli = ClienteExtranjero(nombre, direccion, types.IdPasaporte, c.NroIdentificacion, types.CountryCode(pais))
			return cli, nil
		default:
			return cli, fmt.Errorf("docbuilder: unknown cliente tipo %q (must be contribuyente, consumidorFinal, gobierno, or extranjero)", tipo)
		}
	}

	cli.CodigoUbicacion = ubicacion
	if c.Provincia != "" {
		cli.Provincia = c.Provincia
	}
	if c.Distrito != "" {
		cli.Distrito = c.Distrito
	}
	if c.Corregimiento != "" {
		cli.Corregimiento = c.Corregimiento
	}
	if tel := firstNonEmpty(c.Telefono, c.Telefono1); tel != "" {
		cli.Telefono1 = tel
	}
	if email := firstNonEmpty(c.Correo, c.CorreoElectronico1); email != "" {
		cli.CorreoElectronico1 = email
	}
	return cli, nil
}

func (it *DynamicItem) toItem(idx int) (Item, error) {
	cant := toFloat64(it.Cantidad)
	if cant <= 0 {
		return Item{}, fmt.Errorf("docbuilder: items[%d].cantidad must be > 0", idx)
	}
	pu := toFloat64(it.PrecioUnitario)
	desc := toFloat64(firstNonEmptyAny(it.Descuento, it.PrecioUnitarioDescuento))

	tasa := it.TasaITBMS
	switch tasa {
	case "0", "00", "0%", "exento":
		tasa = "00"
	case "7", "01", "7%":
		tasa = "01"
	case "10", "02", "10%":
		tasa = "02"
	case "15", "03", "15%":
		tasa = "03"
	}

	item := Item{
		Descripcion:      it.Descripcion,
		Codigo:           it.Codigo,
		UnidadMedida:     it.UnidadMedida,
		Cantidad:         cant,
		PrecioUnitario:   pu,
		Descuento:        desc,
		TasaITBMS:        types.TasaITBMS(tasa),
		TasaISC:          it.TasaISC,
		ValorISC:         toFloat64(it.ValorISC),
		CodigoCPBSAbrev:  it.CodigoCPBSAbrev,
		CodigoCPBS:       it.CodigoCPBS,
		UnidadMedidaCPBS: it.UnidadMedidaCPBS,
		FechaFabricacion: it.FechaFabricacion,
		FechaCaducidad:   it.FechaCaducidad,
		InfoItem:         it.InfoItem,
	}

	if it.Medicina != nil {
		item.Medicina = &types.Medicina{
			NroLote:           it.Medicina.NroLote,
			CantProductosLote: it.Medicina.CantProductosLote,
		}
	}
	if it.Vehiculo != nil {
		item.Vehiculo = &types.Vehiculo{
			Marca:           it.Vehiculo.Marca,
			Modelo:          it.Vehiculo.Modelo,
			Anio:            it.Vehiculo.Anio,
			Color:           it.Vehiculo.Color,
			NumeroVIN:       it.Vehiculo.NumeroVIN,
			NumeroMotor:     it.Vehiculo.NumeroMotor,
			NumeroPlaca:     it.Vehiculo.NumeroPlaca,
			TipoCombustible: it.Vehiculo.TipoCombustible,
			PrecioBase:      it.Vehiculo.PrecioBase,
			NumeroDocAduana: it.Vehiculo.NumeroDocAduana,
		}
	}
	return item, nil
}

// --- Dynamic Conversion Helpers ---

func toFloat64(v any) float64 {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		f, _ := strconv.ParseFloat(strings.TrimSpace(val), 64)
		return f
	default:
		return 0
	}
}

func toInt64(v any) int64 {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case int64:
		return val
	case int:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		i, _ := strconv.ParseInt(strings.TrimSpace(val), 10, 64)
		return i
	default:
		return 0
	}
}

func toString(v any) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return strings.TrimSpace(val)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int:
		return strconv.Itoa(val)
	case int64:
		return strconv.FormatInt(val, 10)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func parseTime(v any) (time.Time, error) {
	if v == nil {
		return time.Now(), nil
	}
	switch val := v.(type) {
	case time.Time:
		return val, nil
	case string:
		s := strings.TrimSpace(val)
		if s == "" {
			return time.Now(), nil
		}
		// Try RFC3339
		if t, err := time.Parse(time.RFC3339, s); err == nil {
			return t, nil
		}
		// Try standard Panama format "2006-01-02T15:04:05-07:00"
		if t, err := time.Parse("2006-01-02T15:04:05-07:00", s); err == nil {
			return t, nil
		}
		// Try simple date "2006-01-02"
		if t, err := time.Parse("2006-01-02", s); err == nil {
			return t, nil
		}
		return time.Now(), fmt.Errorf("unable to parse date %q", s)
	default:
		return time.Now(), nil
	}
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func firstNonEmptyAny(vals ...any) any {
	for _, v := range vals {
		if v != nil {
			if s, ok := v.(string); ok && strings.TrimSpace(s) == "" {
				continue
			}
			return v
		}
	}
	return nil
}
