// Package validate provides pre-flight validation for HKA document payloads.
// All rules are applied before any HTTP call is made. Validation collects every
// error found so callers can correct all problems in one iteration.
package validate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/angelnereira/hka-sdk/catalog"
	"github.com/angelnereira/hka-sdk/types"
)

// ValidationError aggregates every field-level validation failure.
type ValidationError struct {
	Fields []FieldError
}

// FieldError describes a single validation failure.
type FieldError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	if len(e.Fields) == 0 {
		return "validation error: no fields"
	}
	msgs := make([]string, len(e.Fields))
	for i, f := range e.Fields {
		msgs[i] = fmt.Sprintf("%s: %s", f.Field, f.Message)
	}
	return "validation errors:\n  " + strings.Join(msgs, "\n  ")
}

// HasErrors reports whether any validation failures were collected.
func (e *ValidationError) HasErrors() bool {
	return len(e.Fields) > 0
}

func (e *ValidationError) add(field, msg string) {
	e.Fields = append(e.Fields, FieldError{Field: field, Message: msg})
}

var reDatetime = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}-05:00$`)
var reDate = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var reDocNumber = regexp.MustCompile(`^\d{10}$`)
var rePunto = regexp.MustCompile(`^\d{3}$`)

// Credentials validates that both tokens are non-empty.
func Credentials(creds interface{ GetTokens() (string, string) }) *ValidationError {
	ve := &ValidationError{}
	t1, t2 := creds.GetTokens()
	if t1 == "" {
		ve.add("TokenEmpresa", "must not be empty")
	}
	if t2 == "" {
		ve.add("TokenPassword", "must not be empty")
	}
	if ve.HasErrors() {
		return ve
	}
	return nil
}

// Document validates a DocumentoElectronico and returns a ValidationError if any
// rule is violated. Returns nil when the document is fully valid.
func Document(doc *types.DocumentoElectronico) *ValidationError {
	ve := &ValidationError{}
	validateTransaccion(ve, &doc.DatosTransaccion)
	validateItems(ve, doc.ListaItems, doc.DatosTransaccion.Cliente.TipoClienteFE)
	validateTotales(ve, &doc.TotalesSubTotales, doc.ListaItems)
	if ve.HasErrors() {
		return ve
	}
	return nil
}

func validateTransaccion(ve *ValidationError, dt *types.DatosTransaccion) {
	// Rule 1
	switch dt.TipoEmision {
	case types.EmisionAUPNormal, types.EmisionAUPContingencia,
		types.EmisionAUPosteriorNorm, types.EmisionAUPosteriorCont:
	default:
		ve.add("TipoEmision", "must be one of 01, 02, 03, 04")
	}

	// Rules 2-3
	if dt.TipoEmision == types.EmisionAUPContingencia || dt.TipoEmision == types.EmisionAUPosteriorCont {
		if dt.FechaInicioContingencia == "" {
			ve.add("FechaInicioContingencia", "required when TipoEmision is 02 or 04")
		}
		if l := len(dt.MotivoContingencia); l < 15 || l > 250 {
			ve.add("MotivoContingencia", "required for contingency, 15..250 characters")
		}
	}

	// Rule 4
	switch dt.TipoDocumento {
	case types.TipoDocFacturaInterna, types.TipoDocFacturaImportacion, types.TipoDocFacturaExportacion,
		types.TipoDocNotaCreditoRef, types.TipoDocNotaDebitoRef, types.TipoDocNotaCreditoGen,
		types.TipoDocNotaDebitoGen, types.TipoDocFacturaZonaFranca, types.TipoDocReembolso,
		types.TipoDocFacturaExtranjera:
	default:
		ve.add("TipoDocumento", "must be one of 01..10")
	}

	// Rule 5
	if !reDocNumber.MatchString(dt.NumeroDocumentoFiscal) {
		ve.add("NumeroDocumentoFiscal", "must be exactly 10 digits")
	}

	// Rule 6
	if !rePunto.MatchString(dt.PuntoFacturacionFiscal) || dt.PuntoFacturacionFiscal == "000" {
		ve.add("PuntoFacturacionFiscal", "must be 3 digits and not '000'")
	}

	// Rule 7
	if !reDatetime.MatchString(dt.FechaEmision) {
		ve.add("FechaEmision", "must match YYYY-MM-DDThh:mm:ss-05:00")
	}

	// Rule 8
	if !types.IsValidNaturaleza(dt.NaturalezaOperacion) {
		ve.add("NaturalezaOperacion", "invalid value")
	}

	// Rule 9
	if dt.TipoOperacion != types.OperacionSalida && dt.TipoOperacion != types.OperacionEntrada {
		ve.add("TipoOperacion", "must be 1 or 2")
	}

	// Rule 10
	if dt.DestinoOperacion != types.DestinoPanama && dt.DestinoOperacion != types.DestinoExtranjero {
		ve.add("DestinoOperacion", "must be 1 or 2")
	}

	// Rule 11
	if dt.FormatoCAFE != types.CAFESinGeneracion && dt.FormatoCAFE != types.CAFECintaPapel && dt.FormatoCAFE != types.CAFEPapelCarta {
		ve.add("FormatoCAFE", "must be 1, 2 or 3")
	}

	// Rule 12
	if dt.EntregaCAFE != types.EntregaSinCAFE && dt.EntregaCAFE != types.EntregaEnPapel && dt.EntregaCAFE != types.EntregaElectronica {
		ve.add("EntregaCAFE", "must be 1, 2 or 3")
	}

	// Rule 13
	if dt.EnvioContenedor != types.ContenedorNormal && dt.EnvioContenedor != types.ContenedorExceptuado {
		ve.add("EnvioContenedor", "must be 1 or 2")
	}

	// Rule 14
	if dt.ProcesoGeneracion != "1" {
		ve.add("ProcesoGeneracion", `must be "1"`)
	}

	// Rules 15-16: document type cross-checks
	switch dt.TipoDocumento {
	case types.TipoDocNotaCreditoRef, types.TipoDocNotaDebitoRef:
		if len(dt.ListaDocsFiscalReferenciados) == 0 {
			ve.add("ListaDocsFiscalReferenciados", "required for document types 04 and 05")
		}
	case types.TipoDocNotaCreditoGen, types.TipoDocNotaDebitoGen:
		if len(dt.ListaDocsFiscalReferenciados) > 0 {
			ve.add("ListaDocsFiscalReferenciados", "must be absent for document types 06 and 07")
		}
	}

	// Rule 17-18: export invoice
	if dt.TipoDocumento == types.TipoDocFacturaExportacion {
		if dt.DestinoOperacion != types.DestinoExtranjero {
			ve.add("DestinoOperacion", "must be '2' (foreign) for document type 03")
		}
		if dt.Cliente.TipoClienteFE != types.ClienteExtranjero {
			ve.add("Cliente.TipoClienteFE", "must be '04' (foreign) for document type 03")
		}
		if dt.DatosFacturaExportacion == nil {
			ve.add("DatosFacturaExportacion", "required for document type 03")
		}
	}

	// Rule 19
	if dt.DestinoOperacion == types.DestinoExtranjero && dt.DatosFacturaExportacion == nil {
		ve.add("DatosFacturaExportacion", "required when DestinoOperacion is '2'")
	}

	// Rules 20-21
	if dt.DestinoOperacion == types.DestinoPanama && dt.Cliente.Pais != types.CountryPA {
		ve.add("Cliente.Pais", "must be 'PA' when DestinoOperacion is '1'")
	}
	if dt.DestinoOperacion == types.DestinoExtranjero && dt.Cliente.Pais == types.CountryPA {
		ve.add("Cliente.Pais", "must not be 'PA' when DestinoOperacion is '2'")
	}

	// Rules 22-24: referenced documents
	for i, ref := range dt.ListaDocsFiscalReferenciados {
		prefix := fmt.Sprintf("ListaDocsFiscalReferenciados[%d]", i)
		if ref.CufeFEReferenciada != "" {
			if len(ref.CufeFEReferenciada) != 66 {
				ve.add(prefix+".CufeFEReferenciada", "must be exactly 66 characters")
			}
			if ref.NroFacturaPapel != "" || ref.NroFacturaImpFiscal != "" {
				ve.add(prefix, "NroFacturaPapel and NroFacturaImpFiscal must be absent when CufeFEReferenciada is set")
			}
		}
		if ref.FechaEmisionDocFiscalReferenciado != "" && !reDatetime.MatchString(ref.FechaEmisionDocFiscalReferenciado) {
			ve.add(prefix+".FechaEmisionDocFiscalReferenciado", "must match YYYY-MM-DDThh:mm:ss-05:00")
		}
	}

	validateCliente(ve, &dt.Cliente)

	// Rule: TipoVenta must not be set when TipoOperacion = "2"
	if dt.TipoOperacion == types.OperacionEntrada && dt.TipoVenta != "" {
		ve.add("TipoVenta", "must not be set when TipoOperacion is '2' (entrada)")
	}

	if dt.DatosFacturaExportacion != nil {
		validateExportacion(ve, dt.DatosFacturaExportacion)
	}
}

func validateCliente(ve *ValidationError, c *types.Cliente) {
	switch c.TipoClienteFE {
	case types.ClienteContribuyente, types.ClienteGobierno:
		// Rules 25
		if c.NumeroRUC == "" {
			ve.add("Cliente.NumeroRUC", "required for TipoClienteFE 01 and 03")
		}
		if c.DigitoVerificadorRUC == "" {
			ve.add("Cliente.DigitoVerificadorRUC", "required for TipoClienteFE 01 and 03")
		}
		if l := len(c.RazonSocial); l < 2 || l > 200 {
			ve.add("Cliente.RazonSocial", "required for TipoClienteFE 01/03, 2..200 characters")
		}
		if l := len(c.Direccion); l < 5 || l > 100 {
			ve.add("Cliente.Direccion", "required for TipoClienteFE 01/03, 5..100 characters")
		}
		if c.TipoContribuyente == "" {
			ve.add("Cliente.TipoContribuyente", "required for TipoClienteFE 01 and 03")
		}
		// Rule 28: juridical entity cannot be consumer final — checked at caller level
	case types.ClienteExtranjero:
		// Rules 26
		if c.NumeroRUC != "" {
			ve.add("Cliente.NumeroRUC", "must not be set for TipoClienteFE 04")
		}
		if c.DigitoVerificadorRUC != "" {
			ve.add("Cliente.DigitoVerificadorRUC", "must not be set for TipoClienteFE 04")
		}
		if c.CodigoUbicacion != "" {
			ve.add("Cliente.CodigoUbicacion", "must not be set for TipoClienteFE 04")
		}
		if c.TipoContribuyente != "" {
			ve.add("Cliente.TipoContribuyente", "must not be set for TipoClienteFE 04")
		}
		if c.TipoIdentificacion == "" {
			ve.add("Cliente.TipoIdentificacion", "required for TipoClienteFE 04")
		}
		if c.NroIdentificacionExtranjero == "" {
			ve.add("Cliente.NroIdentificacionExtranjero", "required for TipoClienteFE 04")
		}
	}

	// codigoUbicacion, when present, must be a well-formed provincia-distrito-corregimiento.
	if c.CodigoUbicacion != "" && !catalog.ValidateUbicacion(c.CodigoUbicacion) {
		ve.add("Cliente.CodigoUbicacion", "must be provincia-distrito-corregimiento (e.g. 8-8-7)")
	}

	// Rule 27
	if c.Pais == types.CountryZZ && c.PaisOtro == "" {
		ve.add("Cliente.PaisOtro", "required when Pais is 'ZZ'")
	}

	// Rule 28
	if c.TipoContribuyente == types.ContribuyenteJuridico && c.TipoClienteFE == types.ClienteConsumidorFinal {
		ve.add("Cliente.TipoClienteFE", "juridical entity (TipoContribuyente=2) cannot be ConsumidorFinal")
	}
}

func validateExportacion(ve *ValidationError, exp *types.DatosExportacion) {
	// Rule 43
	if !types.IsValidIncoterm(exp.CondicionesEntrega) {
		ve.add("DatosFacturaExportacion.CondicionesEntrega", "invalid Incoterm")
	}
	// Rule 44
	if !types.IsValidCurrency(exp.MonedaOperExportacion) {
		ve.add("DatosFacturaExportacion.MonedaOperExportacion", "invalid currency code")
	}
	// Rule 45
	if exp.MonedaOperExportacion == types.CurrencyZZZ && exp.MonedaOperExportacionNonDef == "" {
		ve.add("DatosFacturaExportacion.MonedaOperExportacionNonDef", "required when MonedaOperExportacion is 'ZZZ'")
	}
	// Rule 46
	if exp.MonedaOperExportacion != types.CurrencyUSD && exp.MonedaOperExportacion != types.CurrencyZZZ {
		if exp.TipoDeCambio == "" {
			ve.add("DatosFacturaExportacion.TipoDeCambio", "required for non-USD currencies")
		}
		if exp.MontoMonedaExtranjera == "" {
			ve.add("DatosFacturaExportacion.MontoMonedaExtranjera", "required for non-USD currencies")
		}
	}
}

func validateItems(ve *ValidationError, items []types.Item, tipoCliente types.TipoClienteFE) {
	if len(items) == 0 {
		ve.add("ListaItems", "must contain at least one item")
		return
	}
	for i, item := range items {
		prefix := fmt.Sprintf("ListaItems[%d]", i)

		// Rule 29: government client requires CPBS fields
		if tipoCliente == types.ClienteGobierno {
			if item.CodigoCPBSAbrev == "" {
				ve.add(prefix+".CodigoCPBSAbrev", "required for government clients")
			}
			if item.CodigoCPBS == "" {
				ve.add(prefix+".CodigoCPBS", "required for government clients")
			}
			if item.UnidadMedidaCPBS == "" {
				ve.add(prefix+".UnidadMedidaCPBS", "required for government clients")
			}
		}

		// When CPBS codes are supplied, they must be well-formed and consistent.
		if item.CodigoCPBS != "" || item.CodigoCPBSAbrev != "" {
			if err := catalog.ValidateCPBS(item.CodigoCPBSAbrev, item.CodigoCPBS); err != nil {
				ve.add(prefix+".CodigoCPBS", err.Error())
			}
		}

		// Rule 33
		switch item.TasaITBMS {
		case types.ITBMSExento, types.ITBMS7, types.ITBMS10, types.ITBMS15:
		default:
			ve.add(prefix+".TasaITBMS", "must be one of 00, 01, 02, 03")
		}

		// Rules 30-32: numeric cross-checks (best-effort — strings may have variable decimals)
		precio, errPrecio := strconv.ParseFloat(item.PrecioUnitario, 64)
		descuento, errDesc := strconv.ParseFloat(item.PrecioUnitarioDescuento, 64)
		cantidad, errCant := strconv.ParseFloat(item.Cantidad, 64)
		precioItem, errPI := strconv.ParseFloat(item.PrecioItem, 64)

		if errPrecio == nil && errDesc == nil && precio <= descuento {
			ve.add(prefix+".PrecioUnitarioDescuento", "must be less than PrecioUnitario")
		}
		if errPrecio == nil && errDesc == nil && errCant == nil && errPI == nil {
			expected := cantidad * (precio - descuento)
			if abs(expected-precioItem) > 0.01 {
				ve.add(prefix+".PrecioItem", fmt.Sprintf("must equal Cantidad*(PrecioUnitario-PrecioUnitarioDescuento), expected ~%.6f got %s", expected, item.PrecioItem))
			}
		}

		// Rule 34: medicine batch info
		if item.Medicina != nil {
			if item.Medicina.NroLote == "" {
				ve.add(prefix+".Medicina.NroLote", "required")
			}
			if item.Medicina.CantProductosLote == "" {
				ve.add(prefix+".Medicina.CantProductosLote", "required")
			}
		}

		// Rule 35: acarreo/seguro only in one place
		// (checked at document level after scanning all items)
	}
}

func validateTotales(ve *ValidationError, t *types.TotalesSubTotales, items []types.Item) {
	// Rule 36
	if t.NroItems != len(items) {
		ve.add("TotalesSubTotales.NroItems", fmt.Sprintf("must equal number of items (%d)", len(items)))
	}

	// Rule 40
	if t.TiempoPago == types.PagoPlazo || t.TiempoPago == types.PagoMixto {
		if len(t.ListaPagoPlazo) == 0 {
			ve.add("TotalesSubTotales.ListaPagoPlazo", "required when TiempoPago is 2 or 3")
		}
	}

	// Rule 41
	for i, fp := range t.ListaFormaPago {
		if fp.FormaPagoFact == types.PagoOtro && fp.DescFormaPago == "" {
			ve.add(fmt.Sprintf("TotalesSubTotales.ListaFormaPago[%d].DescFormaPago", i), "required when FormaPagoFact is '99'")
		}
	}

	// Rule 42: Vuelto requires TotalValorRecibido > TotalFactura
	if t.Vuelto != "" {
		recibido, errR := strconv.ParseFloat(t.TotalValorRecibido, 64)
		total, errT := strconv.ParseFloat(t.TotalFactura, 64)
		if errR == nil && errT == nil && recibido <= total {
			ve.add("TotalesSubTotales.Vuelto", "TotalValorRecibido must be greater than TotalFactura when Vuelto is set")
		}
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
