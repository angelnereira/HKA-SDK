package hka_test

import (
	"strings"
	"testing"
	"time"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/docbuilder"
	"github.com/angelnereira/hka-sdk/soap"
	"github.com/angelnereira/hka-sdk/types"
	"github.com/angelnereira/hka-sdk/validate"
)

func TestAllDocumentTypes_EmissionContracts(t *testing.T) {
	cufeRef := "FE01" + strings.Repeat("9", 62) // 66 chars
	now := time.Now()

	cases := []struct {
		name        string
		docType     types.TipoDocumento
		buildFn     func() (*types.DocumentoElectronico, error)
		expectedTag string
	}{
		{
			name:    "01 Factura Interna",
			docType: types.TipoDocFacturaInterna,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaInterna().
					Numero(101).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE S.A.", "Calle 50")).
					AddItem(docbuilder.Item{Descripcion: "Item Normal", Cantidad: 1, PrecioUnitario: 100, TasaITBMS: types.ITBMS7}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>01</ser:tipoDocumento>",
		},
		{
			name:    "02 Factura Importacion",
			docType: types.TipoDocFacturaImportacion,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaImportacion().
					Numero(102).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "IMPORTADORA S.A.", "Via Tocumen")).
					AddItem(docbuilder.Item{Descripcion: "Mercancia Importada", Cantidad: 10, PrecioUnitario: 50, TasaITBMS: types.ITBMS7}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>02</ser:tipoDocumento>",
		},
		{
			name:    "03 Factura Exportacion",
			docType: types.TipoDocFacturaExportacion,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaExportacion().
					Numero(103).
					Punto(1).
					Cliente(docbuilder.ClienteExtranjero("US Imports LLC", "100 Wall St, NY", types.IdTributario, "TAX-12345", types.CountryUS)).
					Exportacion(types.DatosExportacion{
						CondicionesEntrega:    types.IncoFOB,
						MonedaOperExportacion: types.CurrencyUSD,
						PuertoEmbarque:        "Balboa",
					}).
					AddItem(docbuilder.Item{Descripcion: "Bienes Exportados", Cantidad: 5, PrecioUnitario: 200, TasaITBMS: types.ITBMSExento}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>03</ser:tipoDocumento>",
		},
		{
			name:    "04 Nota Credito Referenciada",
			docType: types.TipoDocNotaCreditoRef,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewNotaCreditoReferenciada().
					Numero(104).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE S.A.", "Calle 50")).
					Referencia(cufeRef, now.Add(-24*time.Hour)).
					AddItem(docbuilder.Item{Descripcion: "Ajuste de precio", Cantidad: 1, PrecioUnitario: 50, TasaITBMS: types.ITBMS7}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>04</ser:tipoDocumento>",
		},
		{
			name:    "05 Nota Debito Referenciada",
			docType: types.TipoDocNotaDebitoRef,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewNotaDebitoReferenciada().
					Numero(105).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE S.A.", "Calle 50")).
					Referencia(cufeRef, now.Add(-24*time.Hour)).
					AddItem(docbuilder.Item{Descripcion: "Intereses por mora", Cantidad: 1, PrecioUnitario: 20, TasaITBMS: types.ITBMS7}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>05</ser:tipoDocumento>",
		},
		{
			name:    "06 Nota Credito Generica",
			docType: types.TipoDocNotaCreditoGen,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewNotaCreditoGenerica().
					Numero(106).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE S.A.", "Calle 50")).
					AddItem(docbuilder.Item{Descripcion: "Descuento comercial generico", Cantidad: 1, PrecioUnitario: 30, TasaITBMS: types.ITBMS7}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>06</ser:tipoDocumento>",
		},
		{
			name:    "07 Nota Debito Generica",
			docType: types.TipoDocNotaDebitoGen,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewNotaDebitoGenerica().
					Numero(107).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE S.A.", "Calle 50")).
					AddItem(docbuilder.Item{Descripcion: "Gasto bancario generico", Cantidad: 1, PrecioUnitario: 15, TasaITBMS: types.ITBMS7}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>07</ser:tipoDocumento>",
		},
		{
			name:    "08 Factura Zona Franca",
			docType: types.TipoDocFacturaZonaFranca,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaZonaFranca().
					Numero(108).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "EMPRESA ZONA LIBRE", "Zona Libre Colon")).
					AddItem(docbuilder.Item{Descripcion: "Reexportacion", Cantidad: 1, PrecioUnitario: 300, TasaITBMS: types.ITBMSExento}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>08</ser:tipoDocumento>",
		},
		{
			name:    "09 Reembolso",
			docType: types.TipoDocReembolso,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewReembolso().
					Numero(109).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE S.A.", "Calle 50")).
					AddItem(docbuilder.Item{Descripcion: "Gasto reembolsable", Cantidad: 1, PrecioUnitario: 80, TasaITBMS: types.ITBMSExento}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>09</ser:tipoDocumento>",
		},
		{
			name:    "10 Factura Extranjera",
			docType: types.TipoDocFacturaExtranjera,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaExtranjera().
					Numero(110).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE S.A.", "Calle 50")).
					AddItem(docbuilder.Item{Descripcion: "Operacion extraterritorial", Cantidad: 1, PrecioUnitario: 150, TasaITBMS: types.ITBMSExento}).
					Build()
			},
			expectedTag: "<ser:tipoDocumento>10</ser:tipoDocumento>",
		},
		{
			name:    "Factura Gobierno con CPBS",
			docType: types.TipoDocFacturaInterna,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaGobierno().
					Numero(111).
					Punto(1).
					Cliente(docbuilder.ClienteGobierno("8-NT-1-13656", "43", "MINISTERIO DE SALUD", "Ancon")).
					AddItem(docbuilder.Item{
						Descripcion:      "Medicamentos e insumos",
						Cantidad:         10,
						PrecioUnitario:   25,
						TasaITBMS:        types.ITBMS7,
						CodigoCPBSAbrev:  "13",
						CodigoCPBS:       "1310",
						UnidadMedidaCPBS: "und",
					}).
					Build()
			},
			expectedTag: "<ser:codigoCPBS>1310</ser:codigoCPBS>",
		},
		{
			name:    "Factura Consumidor Final",
			docType: types.TipoDocFacturaInterna,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaConsumidorFinal().
					Numero(112).
					Punto(1).
					Cliente(docbuilder.ClienteConsumidorFinal("Maria Rodriguez", "San Francisco")).
					AddItem(docbuilder.Item{Descripcion: "Consumo restaurante", Cantidad: 1, PrecioUnitario: 45, TasaITBMS: types.ITBMS7}).
					Build()
			},
			expectedTag: "<ser:tipoClienteFE>02</ser:tipoClienteFE>",
		},
		{
			name:    "Factura Pago a Plazo",
			docType: types.TipoDocFacturaInterna,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaInterna().
					Numero(113).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE CREDITO S.A.", "Calle 50")).
					AddItem(docbuilder.Item{Descripcion: "Equipo a credito", Cantidad: 1, PrecioUnitario: 1000, TasaITBMS: types.ITBMS7}).
					AddPagoPlazo(now.Add(30*24*time.Hour), 535, "Cuota 1 de 2").
					AddPagoPlazo(now.Add(60*24*time.Hour), 535, "Cuota 2 de 2").
					Build()
			},
			expectedTag: "<ser:tiempoPago>2</ser:tiempoPago>",
		},
		{
			name:    "Factura con Descuento Global y Retención",
			docType: types.TipoDocFacturaInterna,
			buildFn: func() (*types.DocumentoElectronico, error) {
				return docbuilder.NewFacturaInterna().
					Numero(114).
					Punto(1).
					Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "CLIENTE S.A.", "Calle 50")).
					AddItem(docbuilder.Item{Descripcion: "Servicio profesional", Cantidad: 1, PrecioUnitario: 500, TasaITBMS: types.ITBMS7}).
					DescuentoGlobal(25.0, "Descuento por volumen").
					Retencion(types.Retencion{
						CodigoRetencion: types.RetencionServicioEstado100,
						FechaRetencion:  "2026-10-18",
						MontoRetencion:  "35.00",
					}).
					Build()
			},
			expectedTag: "<ser:descDescuento>Descuento por volumen</ser:descDescuento>",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := tc.buildFn()
			if err != nil {
				t.Fatalf("Build failed: %v", err)
			}

			// Validate with pre-flight validator
			if ve := validate.Document(doc); ve != nil {
				t.Fatalf("Pre-flight validation failed: %v", ve)
			}

			// Serialize to SOAP envelope
			xml := soap.BuildEnviarEnvelope("tokEmp", "tokPass", doc)
			if !strings.Contains(xml, tc.expectedTag) {
				t.Errorf("SOAP envelope does not contain expected tag %s", tc.expectedTag)
			}
		})
	}
}

func TestDynamicPayload_UniversalIngestion(t *testing.T) {
	// 1. Ingestion via ParseDocument with simplified JSON
	jsonDoc := `{
		"tipo": "gobierno",
		"numero": 200,
		"punto": 1,
		"cliente": {
			"tipo": "gobierno",
			"ruc": "8-NT-1-13656",
			"dv": 43,
			"razonSocial": "ENTIDAD PUBLICA",
			"direccion": "Calle 50"
		},
		"items": [
			{
				"descripcion": "Suministros varios",
				"cantidad": 5,
				"precioUnitario": 12.5,
				"tasaITBMS": "01",
				"codigoCPBSAbrev": "13",
				"codigoCPBS": "1310",
				"unidadMedidaCPBS": "und"
			}
		]
	}`

	doc, err := hka.ParseDocument(jsonDoc)
	if err != nil {
		t.Fatalf("ParseDocument failed: %v", err)
	}

	if doc.DatosTransaccion.NumeroDocumentoFiscal != "0000000200" {
		t.Errorf("expected 0000000200, got %s", doc.DatosTransaccion.NumeroDocumentoFiscal)
	}
	if doc.DatosTransaccion.Cliente.TipoClienteFE != types.ClienteGobierno {
		t.Errorf("expected gobierno, got %s", doc.DatosTransaccion.Cliente.TipoClienteFE)
	}
	if doc.TotalesSubTotales.TotalPrecioNeto != "62.50" {
		t.Errorf("expected 62.50, got %s", doc.TotalesSubTotales.TotalPrecioNeto)
	}
}
