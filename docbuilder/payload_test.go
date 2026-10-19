package docbuilder_test

import (
	"strings"
	"testing"

	"github.com/angelnereira/hka-sdk/docbuilder"
	"github.com/angelnereira/hka-sdk/types"
	"github.com/angelnereira/hka-sdk/validate"
)

func TestFromJSON_SimplifiedInterna(t *testing.T) {
	jsonPayload := `{
		"tipo": "01",
		"sucursal": "0000",
		"numero": 12,
		"punto": 1,
		"cliente": {
			"tipo": "contribuyente",
			"ruc": "155596713-2-2015",
			"dv": "59",
			"razonSocial": "THE FACTORY HKA CORP",
			"direccion": "Calle 50, Plaza Morica"
		},
		"items": [
			{
				"descripcion": "Servicio de desarrollo",
				"cantidad": 2,
				"precioUnitario": 100.0,
				"tasaITBMS": "01"
			}
		]
	}`

	doc, err := docbuilder.FromJSON([]byte(jsonPayload))
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	if doc.DatosTransaccion.NumeroDocumentoFiscal != "0000000012" {
		t.Errorf("expected 0000000012, got %s", doc.DatosTransaccion.NumeroDocumentoFiscal)
	}
	if doc.DatosTransaccion.PuntoFacturacionFiscal != "001" {
		t.Errorf("expected 001, got %s", doc.DatosTransaccion.PuntoFacturacionFiscal)
	}
	if doc.TotalesSubTotales.TotalPrecioNeto != "200.00" {
		t.Errorf("expected 200.00, got %s", doc.TotalesSubTotales.TotalPrecioNeto)
	}
	if doc.TotalesSubTotales.TotalITBMS != "14.00" {
		t.Errorf("expected 14.00, got %s", doc.TotalesSubTotales.TotalITBMS)
	}
	if doc.TotalesSubTotales.TotalFactura != "214.00" {
		t.Errorf("expected 214.00, got %s", doc.TotalesSubTotales.TotalFactura)
	}
	if doc.DatosTransaccion.Cliente.CodigoUbicacion != "8-8-8" {
		t.Errorf("expected default 8-8-8, got %s", doc.DatosTransaccion.Cliente.CodigoUbicacion)
	}
}

func TestFromJSON_FullHKASchema(t *testing.T) {
	jsonPayload := `{
		"codigoSucursalEmisor": "0000",
		"datosTransaccion": {
			"tipoEmision": "01",
			"tipoDocumento": "01",
			"numeroDocumentoFiscal": "0000000055",
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
				"razonSocial": "THE FACTORY HKA CORP",
				"direccion": "Calle 50, Plaza Morica",
				"codigoUbicacion": "8-8-8",
				"pais": "PA"
			}
		},
		"listaItems": [
			{
				"descripcion": "Licencia de software",
				"cantidad": "1.000000",
				"precioUnitario": "50.000000",
				"precioUnitarioDescuento": "0.000000",
				"tasaITBMS": "01"
			}
		]
	}`

	doc, err := docbuilder.FromJSON([]byte(jsonPayload))
	if err != nil {
		t.Fatalf("FromJSON full schema failed: %v", err)
	}

	if doc.TotalesSubTotales.TotalFactura != "53.50" {
		t.Errorf("expected computed TotalFactura 53.50, got %s", doc.TotalesSubTotales.TotalFactura)
	}
}

func TestFromJSON_GobiernoWithCPBS(t *testing.T) {
	jsonPayload := `{
		"tipo": "gobierno",
		"numero": 10,
		"punto": 1,
		"cliente": {
			"tipo": "gobierno",
			"ruc": "8-NT-1-13656",
			"dv": "43",
			"razonSocial": "MINISTERIO DE EDUCACION",
			"direccion": "Calle Ernesto Jaen Guardia"
		},
		"items": [
			{
				"descripcion": "Materiales escolares",
				"cantidad": 10,
				"precioUnitario": 5.0,
				"tasaITBMS": "01",
				"codigoCPBSAbrev": "13",
				"codigoCPBS": "1310",
				"unidadMedidaCPBS": "und"
			}
		]
	}`

	doc, err := docbuilder.FromJSON([]byte(jsonPayload))
	if err != nil {
		t.Fatalf("FromJSON gobierno failed: %v", err)
	}

	if doc.DatosTransaccion.Cliente.TipoClienteFE != types.ClienteGobierno {
		t.Errorf("expected ClienteGobierno (03), got %s", doc.DatosTransaccion.Cliente.TipoClienteFE)
	}
	if doc.ListaItems[0].CodigoCPBS != "1310" {
		t.Errorf("expected CPBS 1310, got %s", doc.ListaItems[0].CodigoCPBS)
	}
	if err := validate.Document(doc); err != nil {
		t.Fatalf("built gobierno document failed validation: %v", err)
	}
}

func TestFromJSON_ZonaFranca(t *testing.T) {
	jsonPayload := `{
		"tipo": "08",
		"numero": 20,
		"punto": 1,
		"cliente": {
			"tipo": "contribuyente",
			"ruc": "155596713-2-2015",
			"dv": "59",
			"razonSocial": "ZONA FRANCA USER S.A.",
			"direccion": "Colon Free Zone"
		},
		"items": [
			{
				"descripcion": "Carga y reexportación",
				"cantidad": 1,
				"precioUnitario": 500.0,
				"tasaITBMS": "00"
			}
		]
	}`

	doc, err := docbuilder.FromJSON([]byte(jsonPayload))
	if err != nil {
		t.Fatalf("FromJSON zona franca failed: %v", err)
	}

	if doc.DatosTransaccion.TipoDocumento != types.TipoDocFacturaZonaFranca {
		t.Errorf("expected TipoDocFacturaZonaFranca (08), got %s", doc.DatosTransaccion.TipoDocumento)
	}
	if doc.TotalesSubTotales.TotalFactura != "500.00" {
		t.Errorf("expected 500.00, got %s", doc.TotalesSubTotales.TotalFactura)
	}
}

func TestFromJSON_Exportacion(t *testing.T) {
	jsonPayload := `{
		"tipo": "03",
		"numero": 30,
		"punto": 1,
		"cliente": {
			"tipo": "extranjero",
			"nombre": "Global Logistics Corp",
			"direccion": "100 Broadway, NY",
			"tipoIdentificacion": "02",
			"nroIdentificacion": "TAX-998877",
			"pais": "US"
		},
		"exportacion": {
			"condicionesEntrega": "FOB",
			"moneda": "USD",
			"puertoEmbarque": "Balboa"
		},
		"items": [
			{
				"descripcion": "Exportacion de café panameño",
				"cantidad": 100,
				"precioUnitario": 15.0,
				"tasaITBMS": "00"
			}
		]
	}`

	doc, err := docbuilder.FromJSON([]byte(jsonPayload))
	if err != nil {
		t.Fatalf("FromJSON exportacion failed: %v", err)
	}

	if doc.DatosTransaccion.TipoDocumento != types.TipoDocFacturaExportacion {
		t.Errorf("expected TipoDocFacturaExportacion (03), got %s", doc.DatosTransaccion.TipoDocumento)
	}
	if doc.DatosTransaccion.DestinoOperacion != types.DestinoExtranjero {
		t.Errorf("expected DestinoExtranjero (2), got %s", doc.DatosTransaccion.DestinoOperacion)
	}
	if doc.DatosTransaccion.DatosFacturaExportacion == nil || doc.DatosTransaccion.DatosFacturaExportacion.CondicionesEntrega != types.IncoFOB {
		t.Errorf("expected FOB incoterm, got %v", doc.DatosTransaccion.DatosFacturaExportacion)
	}
}

func TestFromJSON_NotaCreditoReferenciada(t *testing.T) {
	cufe := "FE01" + strings.Repeat("1", 62)
	jsonPayload := `{
		"tipo": "04",
		"numero": 40,
		"punto": 1,
		"cliente": {
			"tipo": "contribuyente",
			"ruc": "155596713-2-2015",
			"dv": "59",
			"razonSocial": "CLIENTE S.A.",
			"direccion": "Calle 50"
		},
		"referencias": [
			{
				"cufe": "` + cufe + `",
				"fechaEmision": "2026-10-15T09:00:00-05:00"
			}
		],
		"items": [
			{
				"descripcion": "Devolucion de mercancia defectuosa",
				"cantidad": 1,
				"precioUnitario": 100.0,
				"tasaITBMS": "01"
			}
		]
	}`

	doc, err := docbuilder.FromJSON([]byte(jsonPayload))
	if err != nil {
		t.Fatalf("FromJSON NC Referenciada failed: %v", err)
	}

	if doc.DatosTransaccion.TipoDocumento != types.TipoDocNotaCreditoRef {
		t.Errorf("expected TipoDocNotaCreditoRef (04), got %s", doc.DatosTransaccion.TipoDocumento)
	}
	if len(doc.DatosTransaccion.ListaDocsFiscalReferenciados) == 0 {
		t.Fatalf("expected referenced document list")
	}
	if doc.DatosTransaccion.ListaDocsFiscalReferenciados[0].CufeFEReferenciada != cufe {
		t.Errorf("expected cufe %s, got %s", cufe, doc.DatosTransaccion.ListaDocsFiscalReferenciados[0].CufeFEReferenciada)
	}
}

func TestFromMap_And_FromPayload(t *testing.T) {
	m := map[string]any{
		"tipo":            "01",
		"numero":          50,
		"punto":           1,
		"cliente":         map[string]any{"tipo": "consumidorFinal", "razonSocial": "Ana Perez", "direccion": "San Francisco"},
		"items":           []any{map[string]any{"descripcion": "Consulta", "cantidad": 1, "precioUnitario": 80.0, "tasaITBMS": "01"}},
		"retencion":       map[string]any{"codigo": "01", "fecha": "2026-10-18", "monto": 8.0},
		"descuentoGlobal": map[string]any{"monto": 5.0, "descripcion": "Descuento cortesía"},
	}

	doc, err := docbuilder.FromMap(m)
	if err != nil {
		t.Fatalf("FromMap failed: %v", err)
	}

	if doc.TotalesSubTotales.TotalPrecioNeto != "80.00" {
		t.Errorf("expected 80.00, got %s", doc.TotalesSubTotales.TotalPrecioNeto)
	}
	if doc.TotalesSubTotales.DescuentoBonificacion == nil || doc.TotalesSubTotales.DescuentoBonificacion.MontoDescuento != "5.00" {
		t.Errorf("expected global discount 5.00, got %v", doc.TotalesSubTotales.DescuentoBonificacion)
	}
	if doc.TotalesSubTotales.Retencion == nil || doc.TotalesSubTotales.Retencion.MontoRetencion != "8" {
		t.Errorf("expected retention 8, got %v", doc.TotalesSubTotales.Retencion)
	}

	// Test FromPayload with string
	jsonStr := `{"tipo":"01","numero":51,"punto":1,"cliente":{"tipo":"consumidorFinal","razonSocial":"Pedro","direccion":"Costa del Este"},"items":[{"descripcion":"Item","cantidad":1,"precioUnitario":20,"tasaITBMS":"01"}]}`
	doc2, err := docbuilder.FromPayload(jsonStr)
	if err != nil {
		t.Fatalf("FromPayload(string) failed: %v", err)
	}
	if doc2.DatosTransaccion.NumeroDocumentoFiscal != "0000000051" {
		t.Errorf("expected 0000000051, got %s", doc2.DatosTransaccion.NumeroDocumentoFiscal)
	}
}

func TestFromJSON_AutoTasaITBMS(t *testing.T) {
	jsonPayload := `{
		"tipo": "01",
		"numero": 60,
		"punto": 1,
		"autoTasaITBMS": true,
		"cliente": {
			"tipo": "consumidorFinal",
			"razonSocial": "Carlos",
			"direccion": "El Cangrejo"
		},
		"items": [
			{
				"descripcion": "Cerveza nacional artesanal",
				"cantidad": 2,
				"precioUnitario": 3.0
			}
		]
	}`

	doc, err := docbuilder.FromJSON([]byte(jsonPayload))
	if err != nil {
		t.Fatalf("FromJSON auto tasa failed: %v", err)
	}

	if doc.ListaItems[0].TasaITBMS != types.ITBMS10 {
		t.Errorf("expected auto inferred ITBMS 10%% (02), got %s", doc.ListaItems[0].TasaITBMS)
	}
	if doc.TotalesSubTotales.TotalFactura != "6.60" { // 6.00 + 0.60
		t.Errorf("expected 6.60, got %s", doc.TotalesSubTotales.TotalFactura)
	}
}

func TestNewConstructors_Gobierno_ConsumidorFinal(t *testing.T) {
	// Factura Gobierno
	gobDoc, err := docbuilder.NewFacturaGobierno().
		Numero(70).
		Punto(1).
		Cliente(docbuilder.ClienteGobierno("8-NT-1-13656", "43", "MINISTERIO", "Ancon")).
		AddItem(docbuilder.Item{
			Descripcion:      "Papelería",
			Cantidad:         1,
			PrecioUnitario:   10,
			TasaITBMS:        types.ITBMS7,
			CodigoCPBSAbrev:  "13",
			CodigoCPBS:       "1310",
			UnidadMedidaCPBS: "und",
		}).
		Build()
	if err != nil {
		t.Fatalf("NewFacturaGobierno failed: %v", err)
	}
	if gobDoc.DatosTransaccion.Cliente.TipoClienteFE != types.ClienteGobierno {
		t.Errorf("expected ClienteGobierno, got %s", gobDoc.DatosTransaccion.Cliente.TipoClienteFE)
	}

	// Factura Consumidor Final
	cfDoc, err := docbuilder.NewFacturaConsumidorFinal().
		Numero(71).
		Punto(1).
		Cliente(docbuilder.ClienteConsumidorFinal("Cliente Final", "Via España")).
		AddItem(docbuilder.Item{
			Descripcion:    "Panaderia",
			Cantidad:       1,
			PrecioUnitario: 5,
			TasaITBMS:      types.ITBMSExento,
		}).
		Build()
	if err != nil {
		t.Fatalf("NewFacturaConsumidorFinal failed: %v", err)
	}
	if cfDoc.DatosTransaccion.Cliente.TipoClienteFE != types.ClienteConsumidorFinal {
		t.Errorf("expected ClienteConsumidorFinal, got %s", cfDoc.DatosTransaccion.Cliente.TipoClienteFE)
	}
}
