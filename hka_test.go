package hka_test

import (
	"testing"
	"time"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/types"
	"github.com/angelnereira/hka-sdk/validate"
)

// --- helpers ---

func validInternalInvoice(docNum int64) *types.DocumentoElectronico {
	now := time.Now()
	return &types.DocumentoElectronico{
		CodigoSucursalEmisor: "0000",
		DatosTransaccion: types.DatosTransaccion{
			TipoEmision:            types.EmisionAUPNormal,
			TipoDocumento:          types.TipoDocFacturaInterna,
			NumeroDocumentoFiscal:  hka.PadDocNumber(docNum),
			PuntoFacturacionFiscal: hka.PadPunto(1),
			FechaEmision:           hka.FormatDateTime(now),
			NaturalezaOperacion:    types.NatVenta,
			TipoOperacion:          types.OperacionSalida,
			DestinoOperacion:       types.DestinoPanama,
			FormatoCAFE:            types.CAFEPapelCarta,
			EntregaCAFE:            types.EntregaElectronica,
			EnvioContenedor:        types.ContenedorNormal,
			ProcesoGeneracion:      "1",
			TipoVenta:              types.VentaServicio,
			Cliente: types.Cliente{
				TipoClienteFE:        types.ClienteContribuyente,
				TipoContribuyente:    types.ContribuyenteJuridico,
				NumeroRUC:            "155596713-2-2015",
				DigitoVerificadorRUC: "59",
				RazonSocial:          "Mi Cliente S.A.",
				Direccion:            "Ave. La Paz, Edificio 100",
				Pais:                 types.CountryPA,
			},
		},
		ListaItems: []types.Item{
			{
				Descripcion:             "Servicio de prueba",
				Cantidad:                "1.000000",
				PrecioUnitario:          "100.000000",
				PrecioUnitarioDescuento: "0.000000",
				PrecioItem:              "100.000000",
				ValorTotal:              "107.000000",
				TasaITBMS:               types.ITBMS7,
				ValorITBMS:              "7.000000",
			},
		},
		TotalesSubTotales: types.TotalesSubTotales{
			TotalPrecioNeto:    "100.00",
			TotalITBMS:         "7.00",
			TotalISC:           "0.00",
			TotalMontoGravado:  "7.00",
			TotalFactura:       "107.00",
			TotalValorRecibido: "107.00",
			TiempoPago:         types.PagoInmediato,
			NroItems:           1,
			TotalTodosItems:    "107.00",
			ListaFormaPago: []types.FormaPagoItem{
				{FormaPagoFact: types.PagoEfectivo, ValorCuotaPagada: "107.00"},
			},
		},
	}
}

// --- validation unit tests ---

func TestValidation_ValidDocument_NoErrors(t *testing.T) {
	doc := validInternalInvoice(1)
	err := validate.Document(doc)
	if err != nil {
		t.Fatalf("expected no validation errors, got: %v", err)
	}
}

func TestValidation_EmptyTipoEmision(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.TipoEmision = ""
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for empty TipoEmision")
	}
	assertFieldError(t, err, "TipoEmision")
}

func TestValidation_InvalidDocNumber(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.NumeroDocumentoFiscal = "123"
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for short NumeroDocumentoFiscal")
	}
	assertFieldError(t, err, "NumeroDocumentoFiscal")
}

func TestValidation_PuntoFacturacion_Zero(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.PuntoFacturacionFiscal = "000"
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for PuntoFacturacionFiscal '000'")
	}
	assertFieldError(t, err, "PuntoFacturacionFiscal")
}

func TestValidation_InvalidFechaEmision(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.FechaEmision = "2024/01/01"
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for FechaEmision format")
	}
	assertFieldError(t, err, "FechaEmision")
}

func TestValidation_ContingencyRequiresMotivo(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.TipoEmision = types.EmisionAUPContingencia
	// FechaInicioContingencia not set
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for missing contingency fields")
	}
	assertFieldError(t, err, "FechaInicioContingencia")
}

func TestValidation_CreditNoteRefRequiresCUFE(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.TipoDocumento = types.TipoDocNotaCreditoRef
	doc.DatosTransaccion.ListaDocsFiscalReferenciados = nil
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for missing ListaDocsFiscalReferenciados on type 04")
	}
	assertFieldError(t, err, "ListaDocsFiscalReferenciados")
}

func TestValidation_GenericNoteProhibitsRefs(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.TipoDocumento = types.TipoDocNotaCreditoGen
	doc.DatosTransaccion.ListaDocsFiscalReferenciados = []types.DocFiscalReferenciado{
		{CufeFEReferenciada: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"},
	}
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for refs present on type 06")
	}
	assertFieldError(t, err, "ListaDocsFiscalReferenciados")
}

func TestValidation_ExportRequiresDatosExportacion(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.TipoDocumento = types.TipoDocFacturaExportacion
	doc.DatosTransaccion.DestinoOperacion = types.DestinoExtranjero
	doc.DatosTransaccion.Cliente.TipoClienteFE = types.ClienteExtranjero
	doc.DatosTransaccion.Cliente.NumeroRUC = ""
	doc.DatosTransaccion.Cliente.DigitoVerificadorRUC = ""
	doc.DatosTransaccion.Cliente.TipoContribuyente = ""
	doc.DatosTransaccion.Cliente.TipoIdentificacion = types.IdPasaporte
	doc.DatosTransaccion.Cliente.NroIdentificacionExtranjero = "P1234567"
	doc.DatosTransaccion.Cliente.Pais = types.CountryUS
	// DatosFacturaExportacion intentionally left nil
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for missing DatosFacturaExportacion on type 03")
	}
	assertFieldError(t, err, "DatosFacturaExportacion")
}

func TestValidation_JuridicaCannotBeConsumidorFinal(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.Cliente.TipoClienteFE = types.ClienteConsumidorFinal
	doc.DatosTransaccion.Cliente.TipoContribuyente = types.ContribuyenteJuridico
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for juridical entity as ConsumidorFinal")
	}
	assertFieldError(t, err, "Cliente.TipoClienteFE")
}

func TestValidation_GovernmentRequiresCPBS(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.Cliente.TipoClienteFE = types.ClienteGobierno
	doc.DatosTransaccion.Cliente.TipoContribuyente = types.ContribuyenteJuridico
	// items intentionally missing CPBS fields
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for missing CPBS on government client")
	}
	assertFieldError(t, err, "ListaItems[0].CodigoCPBSAbrev")
}

func TestValidation_NroItemsMismatch(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.TotalesSubTotales.NroItems = 99
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for NroItems mismatch")
	}
	assertFieldError(t, err, "TotalesSubTotales.NroItems")
}

func TestValidation_PagoPlazoRequired(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.TotalesSubTotales.TiempoPago = types.PagoPlazo
	doc.TotalesSubTotales.ListaPagoPlazo = nil
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for missing ListaPagoPlazo")
	}
	assertFieldError(t, err, "TotalesSubTotales.ListaPagoPlazo")
}

func TestValidation_PaisOtroRequired(t *testing.T) {
	doc := validInternalInvoice(1)
	doc.DatosTransaccion.DestinoOperacion = types.DestinoExtranjero
	doc.DatosTransaccion.Cliente.TipoClienteFE = types.ClienteExtranjero
	doc.DatosTransaccion.Cliente.TipoContribuyente = ""
	doc.DatosTransaccion.Cliente.NumeroRUC = ""
	doc.DatosTransaccion.Cliente.DigitoVerificadorRUC = ""
	doc.DatosTransaccion.Cliente.TipoIdentificacion = types.IdPasaporte
	doc.DatosTransaccion.Cliente.NroIdentificacionExtranjero = "P99999"
	doc.DatosTransaccion.Cliente.Pais = types.CountryZZ
	doc.DatosTransaccion.Cliente.PaisOtro = "" // intentionally empty
	doc.DatosTransaccion.DatosFacturaExportacion = &types.DatosExportacion{
		CondicionesEntrega:    types.IncoFOB,
		MonedaOperExportacion: types.CurrencyUSD,
	}
	err := validate.Document(doc)
	if err == nil {
		t.Fatal("expected validation error for missing PaisOtro when Pais = ZZ")
	}
	assertFieldError(t, err, "Cliente.PaisOtro")
}

// --- helpers ---

func assertFieldError(t *testing.T, err *validate.ValidationError, field string) {
	t.Helper()
	for _, fe := range err.Fields {
		if fe.Field == field {
			return
		}
	}
	t.Errorf("expected FieldError for field %q, got errors: %v", field, err)
}

// --- helpers tests ---

func TestPadDocNumber(t *testing.T) {
	if got := hka.PadDocNumber(1); got != "0000000001" {
		t.Errorf("PadDocNumber(1) = %q, want %q", got, "0000000001")
	}
}

func TestPadPunto(t *testing.T) {
	if got := hka.PadPunto(1); got != "001" {
		t.Errorf("PadPunto(1) = %q, want %q", got, "001")
	}
}

func TestValidateCUFE(t *testing.T) {
	valid := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	if !hka.ValidateCUFE(valid) {
		t.Error("expected 66-char CUFE to be valid")
	}
	if hka.ValidateCUFE("short") {
		t.Error("expected short CUFE to be invalid")
	}
}

func TestFormatAmount(t *testing.T) {
	if got := hka.FormatAmount(107.0, 2); got != "107.00" {
		t.Errorf("FormatAmount(107.0, 2) = %q, want %q", got, "107.00")
	}
}

func TestCalculateITBMS(t *testing.T) {
	result := hka.CalculateITBMS(100.0, types.ITBMS7)
	// floating-point multiplication may not be exactly 7.0; allow a small tolerance
	if result < 6.9999 || result > 7.0001 {
		t.Errorf("CalculateITBMS(100, ITBMS7) = %f, want ~7.0", result)
	}
}

func TestIsWithin180Days(t *testing.T) {
	recent := time.Now().Add(-10 * 24 * time.Hour)
	if !hka.IsWithin180Days(recent) {
		t.Error("expected date 10 days ago to be within 180 days")
	}
	old := time.Now().Add(-200 * 24 * time.Hour)
	if hka.IsWithin180Days(old) {
		t.Error("expected date 200 days ago to be outside 180 days")
	}
}

func TestIsWithin7Days(t *testing.T) {
	recent := time.Now().Add(-3 * 24 * time.Hour)
	if !hka.IsWithin7Days(recent) {
		t.Error("expected date 3 days ago to be within 7 days")
	}
	old := time.Now().Add(-10 * 24 * time.Hour)
	if hka.IsWithin7Days(old) {
		t.Error("expected date 10 days ago to be outside 7 days")
	}
}
