# Guía de uso — hka-sdk

Guía práctica y detallada para integrar la facturación electrónica de **The Factory
HKA (Panamá — DGI)** con este SDK. Está pensada para que cualquier desarrollador
—use Go u otro lenguaje— entienda el flujo completo y los casos especiales.

## Contenido

1. [Conceptos en 2 minutos](#1-conceptos-en-2-minutos)
2. [Requisitos previos](#2-requisitos-previos)
3. [El flujo completo (de la idea al CUFE)](#3-el-flujo-completo-de-la-idea-al-cufe)
4. [Paso a paso en Go](#4-paso-a-paso-en-go)
5. [Los 10 tipos de documento, con ejemplos](#5-los-10-tipos-de-documento-con-ejemplos)
6. [Casos especiales](#6-casos-especiales)
7. [Operaciones posteriores (estado, PDF, correo, anulación)](#7-operaciones-posteriores)
8. [Manejo de errores](#8-manejo-de-errores)
9. [Desde otro lenguaje (gateway REST/JSON)](#9-desde-otro-lenguaje-gateway-restjson)
10. [Catálogos y formatos](#10-catálogos-y-formatos)
11. [Errores comunes (FAQ)](#11-errores-comunes-faq)
12. [Checklist para producción](#12-checklist-para-producción)

---

## 1. Conceptos en 2 minutos

| Término | Qué es |
|---|---|
| **PAC** | Proveedor Autorizado de Calificación. The Factory HKA firma tu documento y lo envía a la DGI. |
| **Credenciales** | `tokenEmpresa` y `tokenPassword`, que te entrega HKA. Son **únicas por RUC**. |
| **Ambiente** | `demo` (pruebas) y `producción`. La data de demo **no** sirve en producción. |
| **CUFE** | Código Único de Factura Electrónica (~66 caracteres). Lo devuelve HKA al autorizar. **Es el dato que debes guardar.** |
| **CAFE** | Comprobante Auxiliar (el PDF) que se entrega al cliente. |
| **Folios** | Cupos de documentos de tu licencia. |
| **ITBMS** | El IVA de Panamá: 7% general, 10% alcohol/hospedaje, 15% tabaco, 0% exento. |

**Idea central del SDK:** tú describes la venta con datos naturales (cantidad,
precio, tasa) y el SDK **calcula los montos**, **fija los campos obligatorios** según
el tipo de documento y **valida** antes de enviar. Así es muy difícil incumplir.

---

## 2. Requisitos previos

- **Go 1.23+** si integras en Go. Sin dependencias externas.
- **Credenciales HKA** (`HKA_TOKEN_EMPRESA`, `HKA_TOKEN_PASSWORD`).
- Para otros lenguajes: solo necesitas correr el **gateway** (ver §9) y hacer HTTP.

Instalación en Go:

```bash
go get github.com/angelnereira/hka-sdk
```

Guarda tus credenciales como variables de entorno (nunca en el código):

```bash
export HKA_TOKEN_EMPRESA="..."
export HKA_TOKEN_PASSWORD="..."
```

---

## 3. El flujo completo (de la idea al CUFE)

```
1. Crear el cliente del SDK   ──>  hka.NewDemo() / hka.New(...)
2. Construir el documento     ──>  docbuilder.NewFacturaInterna()...Build()
        (calcula montos + valida automáticamente)
3. Enviar                     ──>  client.Send(ctx, creds, doc)
4. Leer la respuesta          ──>  resp.CUFE, resp.QR  (GUÁRDALOS)
5. (opcional) Descargar PDF   ──>  client.DownloadPDF(...)
6. (opcional) Enviar correo   ──>  client.SendEmail(...)
7. (opcional) Consultar/anular──>  client.DocumentStatus(...) / client.Cancel(...)
```

Si el documento está mal formado, el paso 2 o 3 devuelve un **`*ValidationError`
sin tocar la red**, con la lista de todos los problemas.

---

## 4. Paso a paso en Go

### 4.1 Crear el cliente

```go
import hka "github.com/angelnereira/hka-sdk"

// Demo / integración (por defecto)
client := hka.NewDemo()

// Producción (endpoint que te da HKA)
client := hka.New(&hka.Config{
    Endpoint: "https://emision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc",
    Timeout:  60 * time.Second,
})
```

El `Client` es **seguro para concurrencia**: créalo una vez y compártelo.

### 4.2 Las credenciales van por llamada (multi-tenant)

```go
creds := hka.Credentials{
    TokenEmpresa:  os.Getenv("HKA_TOKEN_EMPRESA"),
    TokenPassword: os.Getenv("HKA_TOKEN_PASSWORD"),
}
```

Si manejas varias empresas, pasa unas `Credentials` distintas en cada llamada.

### 4.3 Construir el documento con el builder

```go
import (
    "github.com/angelnereira/hka-sdk/docbuilder"
    "github.com/angelnereira/hka-sdk/types"
)

doc, err := docbuilder.NewFacturaInterna().
    Sucursal("0000").                 // código de sucursal (4 dígitos)
    Numero(1).                        // número de documento (se rellena a 10 dígitos)
    Punto(1).                         // punto de facturación (3 dígitos, ≠ 000)
    Cliente(docbuilder.ClienteContribuyente(
        "155596713-2-2015", "59",     // RUC y dígito verificador
        "Mi Cliente S.A.",            // razón social
        "Ave. La Paz, Edificio 100",  // dirección
    )).
    AddItem(docbuilder.Item{
        Descripcion:    "Servicio de consultoría",
        Cantidad:       1,
        PrecioUnitario: 100,
        TasaITBMS:      types.ITBMS7, // 7%
    }).
    Build()
```

Qué hizo el builder por ti:

- Fijó `tipoDocumento=01`, `tipoEmision=01`, `procesoGeneracion=1`, formatos CAFE, etc.
- Calculó del ítem: `precioItem=100.00`, `valorITBMS=7.00`, `valorTotal=107.00`.
- Calculó los totales: `totalPrecioNeto=100.00`, `totalITBMS=7.00`, `totalFactura=107.00`, `nroItems=1`.
- Agregó un pago de contado en efectivo por `107.00`.
- Validó todo el documento. Si algo falla, `doc` es `nil` y `err` lo explica.

### 4.4 Enviar y leer el CUFE

```go
resp, err := client.Send(context.Background(), creds, doc)
if err != nil {
    log.Fatal(err) // ver §8 para distinguir el tipo de error
}
fmt.Println("CUFE:", resp.CUFE)              // guárdalo en tu base de datos
fmt.Println("Protocolo:", resp.NroProtocoloAutorizacion)
fmt.Println("QR:", resp.QR)
```

> 💾 **Guarda el CUFE.** Lo necesitarás para descargar el PDF, rastrear el correo,
> anular o referenciar este documento desde una nota de crédito/débito.

### 4.5 La "llave" del documento

Varias operaciones (estado, anulación, descargas) identifican el documento con su
llave de 5 campos. La construyes con los mismos datos que usaste para emitir:

```go
key := hka.DocumentKey{
    CodigoSucursalEmisor:   doc.CodigoSucursalEmisor,
    NumeroDocumentoFiscal:  doc.DatosTransaccion.NumeroDocumentoFiscal,
    PuntoFacturacionFiscal: doc.DatosTransaccion.PuntoFacturacionFiscal,
    TipoDocumento:          doc.DatosTransaccion.TipoDocumento,
    TipoEmision:            doc.DatosTransaccion.TipoEmision,
}
```

---

## 5. Los 10 tipos de documento, con ejemplos

Cada constructor ya fija lo que el tipo exige. Solo cambias el constructor y, en
algunos casos, agregas un dato específico.

```go
// 01 — Factura de operación interna
docbuilder.NewFacturaInterna().Cliente(cli).AddItem(item)

// 02 — Factura de importación
docbuilder.NewFacturaImportacion().Cliente(cli).AddItem(item)

// 03 — Factura de exportación (cliente extranjero + datos de exportación)
docbuilder.NewFacturaExportacion().
    Cliente(docbuilder.ClienteExtranjero("Foreign Corp", "123 Main St",
        types.IdTributario, "TAX-99", types.CountryUS)).
    Exportacion(types.DatosExportacion{
        CondicionesEntrega:    types.IncoFOB,
        MonedaOperExportacion: types.CurrencyUSD,
        PuertoEmbarque:        "Balboa",
    }).
    AddItem(item)

// 04 — Nota de crédito referenciada (CUFE del documento original)
docbuilder.NewNotaCreditoReferenciada().
    Cliente(cli).
    Referencia("FE01...<66 chars>", fechaEmisionOriginal).
    AddItem(item)

// 05 — Nota de débito referenciada
docbuilder.NewNotaDebitoReferenciada().Cliente(cli).Referencia(cufe, fecha).AddItem(item)

// 06 — Nota de crédito genérica (sin referencia)
docbuilder.NewNotaCreditoGenerica().Cliente(cli).AddItem(item)

// 07 — Nota de débito genérica
docbuilder.NewNotaDebitoGenerica().Cliente(cli).AddItem(item)

// 08 — Factura de zona franca
docbuilder.NewFacturaZonaFranca().Cliente(cli).AddItem(item)

// 09 — Reembolso
docbuilder.NewReembolso().Cliente(cli).AddItem(item)

// 10 — Factura de operación extranjera
docbuilder.NewFacturaExtranjera().Cliente(cli).AddItem(item)
```

Todos terminan con `.Build()`.

---

## 6. Casos especiales

### 6.1 Tipos de cliente

```go
// Contribuyente (RUC jurídico)
docbuilder.ClienteContribuyente("155596713-2-2015", "59", "Empresa S.A.", "Ave. Balboa")

// Contribuyente persona natural
docbuilder.ClienteContribuyenteNatural("8-123-456", "10", "Juan Pérez", "Calle 50")

// Consumidor final (sin RUC)
docbuilder.ClienteConsumidorFinal("Juan Pérez", "Calle 50")

// Gobierno (los ítems requieren códigos CPBS — ver 6.4)
docbuilder.ClienteGobierno("155596713-2-2015", "59", "Ministerio X", "Ave. Balboa")

// Extranjero (para exportación)
docbuilder.ClienteExtranjero("Foreign Corp", "123 Main St", types.IdTributario, "TAX-99", types.CountryUS)
```

### 6.2 Varios ítems, descuento e ISC

```go
docbuilder.NewFacturaInterna().
    Cliente(cli).
    AddItem(docbuilder.Item{Descripcion: "Producto A", Cantidad: 2, PrecioUnitario: 10, TasaITBMS: types.ITBMS7}).
    AddItem(docbuilder.Item{
        Descripcion: "Producto B", Cantidad: 1, PrecioUnitario: 50,
        Descuento: 5,                 // descuento POR UNIDAD, en Balboas
        TasaITBMS: types.ITBMS10,     // 10%
    }).
    Build()
```

El builder suma todo y recalcula los totales; no tienes que cuadrar nada a mano.

### 6.3 Formas de pago y pago a plazo

```go
// Contado explícito por método (si no llamas nada, usa efectivo por el total)
b.PagoContado(types.PagoTransferencia)

// Pago a plazo (cambia TiempoPago a "plazo"; mezclar con AddFormaPago => "mixto")
b.AddPagoPlazo(time.Now().AddDate(0, 1, 0), 107.00, "Cuota 1")
```

### 6.4 Cliente de Gobierno (CPBS obligatorio)

```go
item := docbuilder.Item{
    Descripcion: "Servicio", Cantidad: 1, PrecioUnitario: 100, TasaITBMS: types.ITBMS7,
    CodigoCPBSAbrev:  "72",   // categoría (2 dígitos)
    CodigoCPBS:       "7210", // código (4 dígitos; sus 2 primeros = CPBSAbrev)
    UnidadMedidaCPBS: "cm",
}
docbuilder.NewFacturaInterna().Cliente(clienteGobierno).AddItem(item).Build()
```

### 6.5 Inferir la tasa de ITBMS por descripción (opcional)

```go
doc, _ := docbuilder.NewFacturaInterna().
    AutoTasaITBMS().          // activa la inferencia
    Cliente(cli).
    AddItem(docbuilder.Item{Descripcion: "Cerveza nacional", Cantidad: 1, PrecioUnitario: 2}).
    Build()
// El ítem queda con TasaITBMS=02 (10%) porque "cerveza" es bebida alcohólica.
```

> ⚠️ Es una **ayuda**, no una determinación legal. La responsabilidad de la tasa es
> del emisor. Si pones `TasaITBMS` explícita, esa manda.

### 6.6 Retención

```go
b.Retencion(types.Retencion{
    CodigoRetencion: types.RetencionServicioEstado100,
    FechaRetencion:  "2026-06-03",
    MontoRetencion:  "7.00",
})
```

---

## 7. Operaciones posteriores

```go
// Estado del documento
st, _ := client.DocumentStatus(ctx, creds, key)
fmt.Println(st.EstatusDocumento, st.MensajeDocumento)

// Descargar el CAFE (PDF) y guardarlo
pdfResp, _ := client.DownloadPDF(ctx, creds, key, "" /* serial opcional */)
pdfBytes, _ := pdfResp.Bytes()           // decodifica el Base64
os.WriteFile("cafe.pdf", pdfBytes, 0o644)

// Descargar el XML firmado
xmlResp, _ := client.DownloadXML(ctx, creds, key)
xmlBytes, _ := xmlResp.Bytes()

// Enviar por correo y rastrear
client.SendEmail(ctx, creds, key, "cliente@correo.com")
track, _ := client.TrackEmail(ctx, creds, resp.CUFE)

// Folios disponibles
folios, _ := client.RemainingFolios(ctx, creds)
fmt.Println("Disponibles:", folios.FoliosDisponibles)

// Anular (dentro de 7 días)
if hka.IsWithin7Days(emitidoEn) {
    client.Cancel(ctx, creds, key, "Error en el monto")
}

// Consultar el dígito verificador de un RUC
ruc, _ := client.QueryRUC(ctx, creds, types.RUCJuridico, "155596713-2-2015")
fmt.Println("DV:", ruc.DV)
```

---

## 8. Manejo de errores

Hay **tres** tipos de error, y conviene distinguirlos:

```go
resp, err := client.Send(ctx, creds, doc)
if err != nil {
    switch e := err.(type) {

    case *hka.ValidationError:
        // El documento no cumple. NO se hizo ninguna llamada a HKA.
        // e.Fields lista cada problema (campo + mensaje) para corregirlos todos.
        for _, fe := range e.Fields {
            fmt.Printf("  %s: %s\n", fe.Field, fe.Message)
        }

    case *hka.HKAError:
        // HKA respondió con un código de error (p. ej. 119 sin folios, 102 duplicado).
        fmt.Println("código:", e.Code, "mensaje:", e.Message)

    case *hka.NetworkError:
        // Falló la conexión / HTTP. OJO: el documento PUDO haberse recibido;
        // antes de reenviar, consulta el estado para no duplicar.
        fmt.Println("red:", e.Cause)
    }
    return
}
```

> Por seguridad de cumplimiento, el SDK **no reintenta `Send` automáticamente**: un
> reenvío tras un error de red podría crear una factura duplicada. Ante un
> `NetworkError`, consulta `DocumentStatus` y reenvía solo si no fue recibido.

---

## 9. Desde otro lenguaje (gateway REST/JSON)

Si no usas Go, levanta el **gateway**: oculta el SOAP, calcula los totales y expone
JSON limpio. Cualquier lenguaje lo consume con HTTP.

```bash
# Arrancar (endpoint demo por defecto)
go run ./cmd/hka-gateway          # :8080
# o con Docker
docker build -t hka-gateway . && docker run -p 8080:8080 hka-gateway
```

Las credenciales van en headers por petición:

```
X-HKA-Token-Empresa: <token>
X-HKA-Token-Password: <token>
```

### 9.1 Probar sin enviar (dry-run): el gateway calcula los totales

```bash
curl -s localhost:8080/v1/documents/build -d '{
  "tipo":"01","numero":1,"punto":1,
  "cliente":{"tipo":"contribuyente","ruc":"155596713-2-2015","dv":"59",
             "razonSocial":"Mi Cliente S.A.","direccion":"Ave. La Paz"},
  "items":[{"descripcion":"Servicio","cantidad":1,"precioUnitario":100,"tasaITBMS":"01"}]
}'
```

### 9.2 Enviar de verdad

```bash
curl -s localhost:8080/v1/documents/send \
  -H "X-HKA-Token-Empresa: $HKA_TOKEN_EMPRESA" \
  -H "X-HKA-Token-Password: $HKA_TOKEN_PASSWORD" \
  -d @factura.json
```

### 9.3 TypeScript / JavaScript

```ts
const res = await fetch("http://localhost:8080/v1/documents/send", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "X-HKA-Token-Empresa": process.env.HKA_TOKEN_EMPRESA!,
    "X-HKA-Token-Password": process.env.HKA_TOKEN_PASSWORD!,
  },
  body: JSON.stringify({
    tipo: "01", numero: 1, punto: 1,
    cliente: { tipo: "contribuyente", ruc: "155596713-2-2015", dv: "59",
               razonSocial: "Mi Cliente S.A.", direccion: "Ave. La Paz" },
    items: [{ descripcion: "Servicio", cantidad: 1, precioUnitario: 100, tasaITBMS: "01" }],
  }),
});
if (res.status === 422) {
  const { fields } = await res.json(); // errores de validación por campo
  throw new Error(JSON.stringify(fields));
}
const { CUFE } = await res.json();
```

### 9.4 Python

```python
import os, requests

r = requests.post("http://localhost:8080/v1/documents/send",
    headers={
        "X-HKA-Token-Empresa": os.environ["HKA_TOKEN_EMPRESA"],
        "X-HKA-Token-Password": os.environ["HKA_TOKEN_PASSWORD"],
    },
    json={
        "tipo": "01", "numero": 1, "punto": 1,
        "cliente": {"tipo": "contribuyente", "ruc": "155596713-2-2015", "dv": "59",
                    "razonSocial": "Mi Cliente S.A.", "direccion": "Ave. La Paz"},
        "items": [{"descripcion": "Servicio", "cantidad": 1, "precioUnitario": 100, "tasaITBMS": "01"}],
    })
if r.status_code == 422:
    raise SystemExit(r.json())   # errores de validación
r.raise_for_status()
print(r.json()["CUFE"])
```

### 9.5 Generar un cliente tipado automáticamente

```bash
openapi-generator-cli generate -i openapi.yaml -g typescript-fetch -o ./client-ts
openapi-generator-cli generate -i openapi.yaml -g python          -o ./client-py
```

El contrato completo (todos los endpoints, cuerpos y respuestas) está en
[`../openapi.yaml`](../openapi.yaml) y [`GATEWAY.md`](./GATEWAY.md).

---

## 10. Catálogos y formatos

El paquete `catalog` (o los endpoints `/v1/catalog/...` del gateway) te ayuda a
validar y entender los datos:

```go
import "github.com/angelnereira/hka-sdk/catalog"

catalog.ValidateUbicacion("8-8-7")        // codigoUbicacion válido?
catalog.ParseCedula("8-123-456")          // estructura de la cédula
catalog.ParseRUC("155596713-2-2015")      // natural vs jurídico
catalog.ParseCUFE(cufe)                    // tipo de documento + ambiente
catalog.ValidateCPBS("72", "7210")        // consistencia de códigos CPBS
catalog.SugerirTasa("Botella de ron")     // -> "02" (10%)
```

Detalle de formatos (cédula, RUC, CUFE, CAFE) y fuentes oficiales en
[`CATALOGS.md`](./CATALOGS.md).

---

## 11. Errores comunes (FAQ)

| Síntoma | Causa probable | Solución |
|---|---|---|
| `ValidationError` en `Cliente.Direccion` | Dirección con menos de 5 caracteres | Usa una dirección de 5..100 caracteres |
| `ValidationError` en `ListaItems` | No agregaste ítems | Agrega al menos un `AddItem(...)` |
| `ValidationError` en `CodigoCPBS` | Cliente de Gobierno sin CPBS, o `CPBSAbrev` no es prefijo del código | Llena CPBS; `Abrev` = 2 primeros dígitos del código |
| `HKAError` código `119` | Sin folios disponibles | Recarga folios / revisa licencia |
| `HKAError` código `102` | Documento duplicado (mismo número) | Usa un `Numero` no usado |
| Documento rechazado a plazo | Faltaba la cuota | Usa `AddPagoPlazo(...)` |
| Credenciales de demo no funcionan en producción | Data de demo no migra | Recarga sucursales/clientes en producción |
| `401` en el gateway | Faltan los headers de token | Envía `X-HKA-Token-Empresa` y `-Password` |

---

## 12. Checklist para producción

- [ ] Apuntar `Endpoint` al de producción (`hka.New(&hka.Config{Endpoint: ...})`).
- [ ] Cargar de nuevo la data (sucursales, clientes, puntos) en producción.
- [ ] Numeración de documentos secuencial y sin huecos por punto de facturación.
- [ ] **Persistir el CUFE** (y de ser posible el XML/PDF) de cada documento autorizado.
- [ ] Manejar los 3 tipos de error; ante `NetworkError`, consultar estado antes de reenviar.
- [ ] Conservar el CAFE legible al menos 6 meses.
- [ ] Credenciales en variables de entorno / gestor de secretos, nunca en el código.
- [ ] Monitorear folios disponibles (`RemainingFolios`).
```
