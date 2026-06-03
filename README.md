# hka-sdk

Go SDK for The Factory HKA electronic invoicing API (Panama — DGI).

Abstracts the SOAP protocol, provides strongly typed domain types, runs pre-flight validation before every HTTP call, and returns typed errors. Stateless with respect to credentials — one client handles multiple tenants concurrently.

## Requirements

- Go 1.23 or later
- No external dependencies (stdlib only)

## Installation

```
go get github.com/angelnereira/hka-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    hka "github.com/angelnereira/hka-sdk"
    "github.com/angelnereira/hka-sdk/types"
)

func main() {
    client := hka.NewDemo()

    creds := hka.Credentials{
        TokenEmpresa:  os.Getenv("HKA_TOKEN_EMPRESA"),
        TokenPassword: os.Getenv("HKA_TOKEN_PASSWORD"),
    }

    doc := &types.DocumentoElectronico{
        CodigoSucursalEmisor: "0000",
        DatosTransaccion: types.DatosTransaccion{
            TipoEmision:            types.EmisionAUPNormal,
            TipoDocumento:          types.TipoDocFacturaInterna,
            NumeroDocumentoFiscal:  hka.PadDocNumber(1),
            PuntoFacturacionFiscal: hka.PadPunto(1),
            FechaEmision:           hka.FormatDateTime(time.Now()),
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
                Descripcion:             "Servicio de consultoria",
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

    resp, err := client.Send(context.Background(), creds, doc)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("CUFE: %s\n", resp.CUFE)
}
```

## Document Builders (safe-by-construction)

Building a `DocumentoElectronico` by hand means filling dozens of fields and
computing every monetary total yourself. The `docbuilder` package removes that
burden: pick a constructor for one of the ten document types, add items with
their natural values, and the builder computes item ITBMS, item totals, and all
document totals for you, then validates before returning.

```go
import "github.com/angelnereira/hka-sdk/docbuilder"

doc, err := docbuilder.NewFacturaInterna().
    Sucursal("0000").
    Numero(1).
    Punto(1).
    Cliente(docbuilder.ClienteContribuyente(
        "155596713-2-2015", "59", "Mi Cliente S.A.", "Ave. La Paz, Edificio 100",
    )).
    AddItem(docbuilder.Item{
        Descripcion:    "Servicio de consultoria",
        Cantidad:       1,
        PrecioUnitario: 100,
        TasaITBMS:      types.ITBMS7,
    }).
    Build()
if err != nil { /* *ValidationError with every problem found */ }

resp, err := client.Send(ctx, creds, doc)
```

One constructor per document type, each preselecting the transaction defaults the
HKA rules require:

| Constructor | Type |
|---|---|
| `NewFacturaInterna()` | 01 — domestic invoice |
| `NewFacturaImportacion()` | 02 — import invoice |
| `NewFacturaExportacion()` | 03 — export invoice (`.Exportacion(...)`, foreign client) |
| `NewNotaCreditoReferenciada()` | 04 — credit note (`.Referencia(cufe, fecha)`) |
| `NewNotaDebitoReferenciada()` | 05 — debit note (`.Referencia(cufe, fecha)`) |
| `NewNotaCreditoGenerica()` | 06 — generic credit note |
| `NewNotaDebitoGenerica()` | 07 — generic debit note |
| `NewFacturaZonaFranca()` | 08 — free-zone invoice |
| `NewReembolso()` | 09 — reimbursement |
| `NewFacturaExtranjera()` | 10 — foreign-operation invoice |

Client constructors fill the fields each category requires:
`ClienteContribuyente`, `ClienteContribuyenteNatural`, `ClienteConsumidorFinal`,
`ClienteGobierno`, `ClienteExtranjero`.

Payment terms are inferred: with no calls the builder adds a single cash payment
for the full total (immediate); `AddPagoPlazo(...)` switches to deferred, and
combining both yields mixed. See [`examples/builder_quickstart`](./examples/builder_quickstart).

## API Reference

| Method | Description |
|---|---|
| `Send(ctx, creds, doc)` | Submit a fiscal document |
| `DocumentStatus(ctx, creds, key)` | Query document status |
| `Cancel(ctx, creds, key, reason)` | Cancel a document (7-day window) |
| `DownloadXML(ctx, creds, key)` | Download signed XML (Base64) |
| `DownloadPDF(ctx, creds, key, serial)` | Download PDF (Base64) |
| `SendEmail(ctx, creds, key, email)` | Email the document to a recipient |
| `TrackEmail(ctx, creds, cufe)` | Track email delivery by CUFE |
| `RemainingFolios(ctx, creds)` | Query available folio count |
| `QueryRUC(ctx, creds, type, ruc)` | Look up a RUC and its check digit |

## Document Types

| Code | Constant | Key Constraints |
|---|---|---|
| 01 | `TipoDocFacturaInterna` | Standard domestic invoice |
| 02 | `TipoDocFacturaImportacion` | Import invoice |
| 03 | `TipoDocFacturaExportacion` | Requires `DestinoOperacion=2`, `TipoClienteFE=04`, `DatosFacturaExportacion` |
| 04 | `TipoDocNotaCreditoRef` | Requires `ListaDocsFiscalReferenciados` with CUFE, 180-day window |
| 05 | `TipoDocNotaDebitoRef` | Same as 04 |
| 06 | `TipoDocNotaCreditoGen` | `ListaDocsFiscalReferenciados` must be absent |
| 07 | `TipoDocNotaDebitoGen` | Same as 06 |
| 08 | `TipoDocFacturaZonaFranca` | Free zone invoice |
| 09 | `TipoDocReembolso` | Reimbursement |
| 10 | `TipoDocFacturaExtranjera` | Foreign operation invoice |

## Polyglot usage (REST/JSON gateway)

Go is the source of truth, but any language can drive the SDK through the JSON
gateway, which hides SOAP and computes every total for you:

```bash
go run ./cmd/hka-gateway        # listens on :8080 (demo endpoint by default)
# or: docker build -t hka-gateway . && docker run -p 8080:8080 hka-gateway
```

```bash
curl -s localhost:8080/v1/documents/build -d '{
  "tipo":"01","numero":1,"punto":1,
  "cliente":{"tipo":"contribuyente","ruc":"155596713-2-2015","dv":"59",
             "razonSocial":"Mi Cliente S.A.","direccion":"Ave. La Paz"},
  "items":[{"descripcion":"Servicio","cantidad":1,"precioUnitario":100,"tasaITBMS":"01"}]
}'   # -> document with TotalFactura "107.00", computed by the gateway
```

Credentials are passed per request via `X-HKA-Token-Empresa` /
`X-HKA-Token-Password` headers. Generate typed clients for TypeScript, Python,
Java, etc. from [`openapi.yaml`](./openapi.yaml). Full contract and per-language
examples in [`docs/GATEWAY.md`](./docs/GATEWAY.md).

## Catalogs & code formats

The [`catalog`](./catalog) package bundles the reference catalogs and format
helpers the DGI/HKA scheme depends on, as types inside the SDK:

```go
import "github.com/angelnereira/hka-sdk/catalog"

catalog.Provincias()                       // 13 provinces/comarcas
catalog.ParseUbicacion("8-8-7")            // provincia-distrito-corregimiento
catalog.ParseCedula("8-123-456")           // national ID
catalog.ParseRUC("155596713-2-2015")       // natural (cédula) or juridical
catalog.ValidateCUFE(cufe)                 // 66-char FE code shape
catalog.ValidateCPBS("13", "1310")         // government product codes
catalog.SugerirTasa("Cerveza nacional")    // -> ITBMS10 (10%)
```

ITBMS rates are modeled by kind of good/service: 7% general, 10% alcoholic
beverages and lodging, 15% tobacco, 0% exempt. Validation now also checks
`codigoUbicacion` shape and CPBS code consistency.

The large geographic and CPBS catalogs ship with authoritative top-level data and
can be regenerated in full from official sources with `tools/gencatalog`. See
[`docs/CATALOGS.md`](./docs/CATALOGS.md) for sources and the refresh procedure.

## Validation

Pre-flight validation runs before any HTTP call. If the document is invalid, `Send()` returns a `*hka.ValidationError` without making any network request. The error lists every problem found so all fields can be corrected at once:

```go
resp, err := client.Send(ctx, creds, doc)
if err != nil {
    var valErr *hka.ValidationError
    if errors.As(err, &valErr) {
        for _, fe := range valErr.Fields {
            fmt.Printf("%s: %s\n", fe.Field, fe.Message)
        }
    }
}
```

The standalone validators in the `validate` package can also be used independently:

```go
if err := validate.Document(doc); err != nil { ... }
if err := validate.Client(&cliente); err != nil { ... }
if err := validate.Items(items, tipoCliente); err != nil { ... }
```

## Error Handling

```go
resp, err := client.Send(ctx, creds, doc)
if err != nil {
    switch e := err.(type) {
    case *hka.ValidationError:
        // pre-flight failure — no HTTP call was made
        for _, fe := range e.Fields {
            fmt.Println(fe.Field, fe.Message)
        }
    case *hka.HKAError:
        // service returned an error code
        fmt.Println(e.Code, e.Message)
    case *hka.NetworkError:
        // connection or HTTP failure
        fmt.Println(e.Cause)
    }
}
```

## Environments

```go
// Demo / integration
client := hka.NewDemo()

// Production (endpoint provided by HKA)
client := hka.New(&hka.Config{
    Endpoint: "https://emision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc",
    Timeout:  60 * time.Second,
})
```

## Multi-tenant Usage

One `Client` can be shared across goroutines. Pass different `Credentials` per call:

```go
client := hka.NewDemo()

go func() { client.Send(ctx, credsEmpresaA, docA) }()
go func() { client.Send(ctx, credsEmpresaB, docB) }()
```

## Helpers

```go
hka.FormatDateTime(t)           // time.Time → "2006-01-02T15:04:05-05:00"
hka.FormatDate(t)               // time.Time → "2006-01-02"
hka.PadDocNumber(1)             // → "0000000001"
hka.PadPunto(1)                 // → "001"
hka.PadSucursal("1")            // → "0001"
hka.FormatAmount(107.0, 2)      // → "107.00"
hka.CalculateITBMS(100, tasa)   // → net ITBMS amount
hka.IsWithin180Days(refDate)    // check reference note eligibility
hka.IsWithin7Days(emitDate)     // check cancellation eligibility
hka.ValidateCUFE(cufe)          // check CUFE length (66 chars)
```

## Testing

```bash
# Unit tests (no credentials required)
go test ./...

# Integration tests against demo environment
HKA_TOKEN_EMPRESA=xxx HKA_TOKEN_PASSWORD=yyy go test -tags=integration ./... -v
```

## Documentation

Detailed HKA wiki documentation is available in [`hka-docs/`](./hka-docs/):

- `01_documentos_fel/` — Web service integration manuals and method references
- `02_ejemplos_codigo/` — Code examples in PHP, C#, Java, Python
- `03_documentos_fiscales/` — Fiscal document type examples
- `04_documentos_pac/` — PAC services guide
- `05_ayuda/` — FAQ

## License

MIT — see [LICENSE](./LICENSE).
