# hka-sdk

SDK para la API de facturación electrónica de **The Factory HKA** (Panamá — DGI).

Abstrae el protocolo SOAP, ofrece tipos de dominio fuertemente tipados, **construye
los 10 tipos de documento de forma segura** (calculando ITBMS y todos los totales),
valida antes de cada llamada HTTP y devuelve errores tipados. Es **stateless**
respecto a las credenciales — un cliente atiende varios contribuyentes en paralelo —
y, mediante el gateway REST/JSON, es **políglota**: se consume desde TypeScript,
JavaScript, Python, Java o cualquier lenguaje.

## Características

- ✅ **Los 10 tipos de documento fiscal** (factura interna, importación, exportación,
  notas de crédito/débito referenciadas y genéricas, zona franca, reembolso, extranjera).
- 🛡️ **Builder seguro por construcción**: rellena los campos obligatorios por tipo y
  calcula `valorITBMS`, `precioItem`, `valorTotal` y todos los totales — es muy difícil
  emitir un documento que incumpla.
- 🌍 **Gateway REST/JSON + OpenAPI** para integración en cualquier lenguaje.
- 📚 **Catálogos y formatos** como tipos del SDK: ubicación (provincia/distrito/
  corregimiento), cédula, RUC, CUFE/CAFE, ITBMS y CPBS.
- 🔍 **Validación pre-flight** (~48 reglas) que reporta todos los problemas a la vez.
- 🧵 **Multi-tenant** y seguro para concurrencia; sin dependencias externas (solo stdlib).

> 📖 **¿Primera vez?** La [**Guía de uso paso a paso**](./docs/GUIA_USO.md) explica
> todo el flujo en detalle (Go y otros lenguajes), con casos especiales y FAQ.

## Mapa de paquetes

| Paquete | Propósito |
|---|---|
| `hka` (raíz) | Cliente, operaciones (`Send`, `Cancel`, descargas, …), errores tipados, helpers |
| `docbuilder` | Construcción segura de los 10 documentos con auto-cálculo de montos |
| `catalog` | Catálogos (ubicación, CPBS) y formatos (cédula, RUC, CUFE, CAFE, ITBMS) |
| `types` | Tipos de dominio, constantes y enumeraciones |
| `validate` | Validación pre-flight independiente |
| `gateway` + `cmd/hka-gateway` | Servicio REST/JSON políglota (oculta el SOAP) |
| `tools/gencatalog` | Regenera el catálogo geográfico desde una fuente oficial |

## Requisitos e instalación

- Go 1.23 o superior — sin dependencias externas.

```
go get github.com/angelnereira/hka-sdk
```

## Inicio rápido (recomendado)

Usa el `docbuilder`: tú aportas los datos naturales y el SDK calcula los montos,
fija los campos obligatorios del tipo y valida antes de devolver el documento.

```go
package main

import (
    "context"
    "errors"
    "fmt"
    "log"
    "os"

    hka "github.com/angelnereira/hka-sdk"
    "github.com/angelnereira/hka-sdk/docbuilder"
    "github.com/angelnereira/hka-sdk/types"
)

func main() {
    client := hka.NewDemo()
    creds := hka.Credentials{
        TokenEmpresa:  os.Getenv("HKA_TOKEN_EMPRESA"),
        TokenPassword: os.Getenv("HKA_TOKEN_PASSWORD"),
    }

    doc, err := docbuilder.NewFacturaInterna().
        Sucursal("0000").
        Numero(1).
        Punto(1).
        Cliente(docbuilder.ClienteContribuyente(
            "155596713-2-2015", "59", "Mi Cliente S.A.", "Ave. La Paz, Edificio 100",
        )).
        AddItem(docbuilder.Item{
            Descripcion:    "Servicio de consultoría",
            Cantidad:       1,
            PrecioUnitario: 100,
            TasaITBMS:      types.ITBMS7,
        }).
        Build()
    if err != nil {
        var ve *hka.ValidationError
        if errors.As(err, &ve) {
            for _, fe := range ve.Fields {
                fmt.Printf("validación: %s — %s\n", fe.Field, fe.Message)
            }
        }
        log.Fatal(err)
    }

    resp, err := client.Send(context.Background(), creds, doc)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("CUFE: %s\n", resp.CUFE)
}
```

Las condiciones de pago se infieren: sin llamadas adicionales se agrega un pago de
contado por el total; `AddPagoPlazo(...)` cambia a plazo y combinarlo con
`AddFormaPago(...)` resulta en mixto. Ejemplo completo en
[`examples/builder_quickstart`](./examples/builder_quickstart).

## Los 10 tipos de documento

| Código | Constructor (`docbuilder`) | Constante (`types`) | Notas / requisitos |
|---|---|---|---|
| 01 | `NewFacturaInterna()` | `TipoDocFacturaInterna` | Factura de operación interna |
| 02 | `NewFacturaImportacion()` | `TipoDocFacturaImportacion` | Factura de importación |
| 03 | `NewFacturaExportacion()` | `TipoDocFacturaExportacion` | `.Exportacion(...)` + cliente extranjero; destino exterior |
| 04 | `NewNotaCreditoReferenciada()` | `TipoDocNotaCreditoRef` | `.Referencia(cufe, fecha)`; ventana de 180 días |
| 05 | `NewNotaDebitoReferenciada()` | `TipoDocNotaDebitoRef` | Igual que 04 |
| 06 | `NewNotaCreditoGenerica()` | `TipoDocNotaCreditoGen` | Sin documento referenciado |
| 07 | `NewNotaDebitoGenerica()` | `TipoDocNotaDebitoGen` | Igual que 06 |
| 08 | `NewFacturaZonaFranca()` | `TipoDocFacturaZonaFranca` | Factura de zona franca |
| 09 | `NewReembolso()` | `TipoDocReembolso` | Reembolso |
| 10 | `NewFacturaExtranjera()` | `TipoDocFacturaExtranjera` | Operación extranjera |

Constructores de cliente que rellenan lo que cada categoría exige:
`ClienteContribuyente`, `ClienteContribuyenteNatural`, `ClienteConsumidorFinal`,
`ClienteGobierno`, `ClienteExtranjero`.

## Operaciones

| Método | Descripción |
|---|---|
| `Send(ctx, creds, doc)` | Envía un documento fiscal |
| `DocumentStatus(ctx, creds, key)` | Consulta el estado de un documento |
| `Cancel(ctx, creds, key, reason)` | Anula un documento (ventana de 7 días) |
| `DownloadXML(ctx, creds, key)` | Descarga el XML firmado (Base64) |
| `DownloadPDF(ctx, creds, key, serial)` | Descarga el CAFE PDF (Base64) |
| `SendEmail(ctx, creds, key, email)` | Envía el documento por correo |
| `TrackEmail(ctx, creds, cufe)` | Rastrea la entrega por CUFE |
| `RemainingFolios(ctx, creds)` | Consulta los folios disponibles |
| `QueryRUC(ctx, creds, tipo, ruc)` | Consulta un RUC y su dígito verificador |

El resultado de `DownloadXML`/`DownloadPDF` trae el contenido en Base64;
`resp.Bytes()` lo decodifica a bytes listos para escribir a disco.

## Uso políglota (gateway REST/JSON)

Go es la fuente de verdad, pero cualquier lenguaje puede usar el SDK a través del
gateway JSON, que oculta el SOAP y calcula todos los totales:

```bash
go run ./cmd/hka-gateway        # escucha en :8080 (endpoint demo por defecto)
# o: docker build -t hka-gateway . && docker run -p 8080:8080 hka-gateway
```

```bash
curl -s localhost:8080/v1/documents/build -d '{
  "tipo":"01","numero":1,"punto":1,
  "cliente":{"tipo":"contribuyente","ruc":"155596713-2-2015","dv":"59",
             "razonSocial":"Mi Cliente S.A.","direccion":"Ave. La Paz"},
  "items":[{"descripcion":"Servicio","cantidad":1,"precioUnitario":100,"tasaITBMS":"01"}]
}'   # -> documento con TotalFactura "107.00", calculado por el gateway
```

Las credenciales se pasan por petición con los headers `X-HKA-Token-Empresa` /
`X-HKA-Token-Password`. Desde [`openapi.yaml`](./openapi.yaml) se autogeneran
clientes tipados (TypeScript, Python, Java, …). Contrato completo y ejemplos por
lenguaje en [`docs/GATEWAY.md`](./docs/GATEWAY.md).

## Catálogos y formatos

El paquete [`catalog`](./catalog) incluye, como tipos del SDK, los catálogos de
referencia y los ayudantes de formato del esquema DGI/HKA:

```go
import "github.com/angelnereira/hka-sdk/catalog"

catalog.Provincias()                    // 13 provincias/comarcas
catalog.ParseUbicacion("8-8-7")         // provincia-distrito-corregimiento
catalog.ParseCedula("8-123-456")        // cédula
catalog.ParseRUC("155596713-2-2015")    // natural (cédula) o jurídico
catalog.ParseCUFE(cufe)                 // decodifica tipo de documento y ambiente
catalog.ValidateCPBS("13", "1310")      // códigos de producto para Gobierno
catalog.SugerirTasa("Cerveza nacional") // -> ITBMS10 (10%)
```

El ITBMS se modela por tipo de bien/servicio: **7%** general, **10%** bebidas
alcohólicas y hospedaje, **15%** tabaco, **0%** exento. La validación también
verifica el formato de `codigoUbicacion` y la consistencia de los códigos CPBS.

Los catálogos grandes (geográfico y CPBS) incluyen datos de nivel superior
autoritativos y se regeneran completos con `tools/gencatalog`. Fuentes y
procedimiento en [`docs/CATALOGS.md`](./docs/CATALOGS.md).

## Validación

La validación pre-flight corre antes de cualquier llamada HTTP. Si el documento es
inválido, `Send()` devuelve un `*hka.ValidationError` sin hacer red, listando todos
los problemas para corregirlos de una vez:

```go
resp, err := client.Send(ctx, creds, doc)
if err != nil {
    var ve *hka.ValidationError
    if errors.As(err, &ve) {
        for _, fe := range ve.Fields {
            fmt.Printf("%s: %s\n", fe.Field, fe.Message)
        }
    }
}
```

Los validadores del paquete `validate` también se usan de forma independiente:

```go
if err := validate.Document(doc); err != nil { /* ... */ }
```

## Manejo de errores

```go
resp, err := client.Send(ctx, creds, doc)
if err != nil {
    switch e := err.(type) {
    case *hka.ValidationError:
        // fallo pre-flight — no se hizo ninguna llamada HTTP
        for _, fe := range e.Fields {
            fmt.Println(fe.Field, fe.Message)
        }
    case *hka.HKAError:
        // el servicio devolvió un código de error
        fmt.Println(e.Code, e.Message)
    case *hka.NetworkError:
        // fallo de conexión o HTTP
        fmt.Println(e.Cause)
    }
}
```

## Ambientes

```go
// Demo / integración
client := hka.NewDemo()

// Producción (endpoint provisto por HKA)
client := hka.New(&hka.Config{
    Endpoint: "https://emision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc",
    Timeout:  60 * time.Second,
})
```

## Uso multi-tenant

Un mismo `Client` se comparte entre goroutines; pasa `Credentials` distintas por llamada:

```go
client := hka.NewDemo()

go func() { client.Send(ctx, credsEmpresaA, docA) }()
go func() { client.Send(ctx, credsEmpresaB, docB) }()
```

## Construcción manual (avanzado)

Cuando necesites control total, puedes construir el `types.DocumentoElectronico` a
mano y enviarlo con `client.Send`. En ese caso eres responsable de calcular los
montos (`valorITBMS`, `precioItem`, totales); la validación los verifica pero no los
calcula. Para casi todos los casos, prefiere el `docbuilder`. Helpers útiles:

```go
hka.FormatDateTime(t)           // time.Time → "2006-01-02T15:04:05-05:00"
hka.FormatDate(t)               // time.Time → "2006-01-02"
hka.PadDocNumber(1)             // → "0000000001"
hka.PadPunto(1)                 // → "001"
hka.PadSucursal("1")            // → "0001"
hka.FormatAmount(107.0, 2)      // → "107.00"
hka.CalculateITBMS(100, tasa)   // → monto ITBMS
hka.IsWithin180Days(refDate)    // elegibilidad de nota referenciada
hka.IsWithin7Days(emitDate)     // elegibilidad de anulación
```

## Pruebas

```bash
# Pruebas unitarias (sin credenciales)
go test ./...

# Pruebas de integración contra el ambiente demo
HKA_TOKEN_EMPRESA=xxx HKA_TOKEN_PASSWORD=yyy go test -tags=integration ./... -v
```

## Documentación

- [`docs/GUIA_USO.md`](./docs/GUIA_USO.md) — **guía de uso detallada paso a paso** (empieza aquí).
- [`docs/ANALYSIS.md`](./docs/ANALYSIS.md) — análisis del proyecto y hoja de ruta.
- [`docs/GATEWAY.md`](./docs/GATEWAY.md) — gateway REST/JSON y consumo por lenguaje.
- [`docs/CATALOGS.md`](./docs/CATALOGS.md) — catálogos, formatos y fuentes oficiales.
- [`hka-docs/`](./hka-docs/) — documentación de la wiki de HKA (manuales, ejemplos
  XML por tipo de documento, ejemplos de código en PHP/C#/Java/Python, FAQ).

## Licencia

MIT — ver [LICENSE](./LICENSE).
