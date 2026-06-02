# HKA Gateway — SDK políglota vía REST/JSON

El gateway expone el SDK de Go como un servicio **JSON/HTTP neutral al lenguaje**.
Oculta por completo el protocolo SOAP, **calcula todos los totales** con el builder
seguro y aplica la validación de cumplimiento. Cualquier lenguaje (TypeScript,
JavaScript, Python, Java, PHP, C#, …) lo consume con HTTP, y desde
[`openapi.yaml`](../openapi.yaml) se **autogeneran clientes tipados**.

## Ejecutar

```bash
# Demo (por defecto)
go run ./cmd/hka-gateway              # escucha en :8080

# Producción
HKA_ENDPOINT=https://emision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc \
ADDR=:8080 go run ./cmd/hka-gateway

# Docker
docker build -t hka-gateway .
docker run -p 8080:8080 hka-gateway
```

Las **credenciales se pasan por petición** (modelo multi-tenant) mediante headers:

```
X-HKA-Token-Empresa: <token>
X-HKA-Token-Password: <token>
```

## Endpoints

| Método y ruta | Descripción | Credenciales |
|---|---|---|
| `POST /v1/documents/build` | Construye y valida (dry-run, sin enviar) | no |
| `POST /v1/documents/send` | Construye, valida y envía | sí |
| `POST /v1/documents/status` | Estado de un documento | sí |
| `POST /v1/documents/cancel` | Anula (ventana de 7 días) | sí |
| `POST /v1/documents/xml` | Descarga XML firmado (Base64) | sí |
| `POST /v1/documents/pdf` | Descarga CAFE PDF (Base64) | sí |
| `POST /v1/documents/email` | Envía el documento por correo | sí |
| `POST /v1/email/track` | Rastrea entregas por CUFE | sí |
| `GET  /v1/folios` | Folios disponibles | sí |
| `POST /v1/ruc/query` | Consulta RUC y dígito verificador | sí |
| `GET  /v1/catalog/provincias` | Provincias/comarcas | no |
| `GET  /v1/catalog/ubicacion/{code}` | Resuelve `codigoUbicacion` | no |
| `GET  /v1/catalog/cufe/{cufe}` | Decodifica campos del CUFE | no |
| `GET  /v1/catalog/cpbs/{codigo}` | Consulta CPBS | no |
| `POST /v1/catalog/itbms/suggest` | Sugiere tasa ITBMS por descripción | no |

## Ejemplos de consumo

### curl

```bash
# Dry-run: construye y valida una factura interna; el gateway calcula los totales.
curl -s localhost:8080/v1/documents/build -d '{
  "tipo": "01", "numero": 1, "punto": 1,
  "cliente": {"tipo":"contribuyente","ruc":"155596713-2-2015","dv":"59",
              "razonSocial":"Mi Cliente S.A.","direccion":"Ave. La Paz"},
  "items": [{"descripcion":"Servicio de consultoría","cantidad":1,"precioUnitario":100,"tasaITBMS":"01"}]
}' | jq .TotalesSubTotales.TotalFactura     # "107.00"

# Enviar (requiere credenciales)
curl -s localhost:8080/v1/documents/send \
  -H 'X-HKA-Token-Empresa: '"$HKA_TOKEN_EMPRESA" \
  -H 'X-HKA-Token-Password: '"$HKA_TOKEN_PASSWORD" \
  -d @factura.json | jq .CUFE
```

### TypeScript / JavaScript (fetch)

```ts
const res = await fetch("http://localhost:8080/v1/documents/send", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "X-HKA-Token-Empresa": process.env.HKA_TOKEN_EMPRESA!,
    "X-HKA-Token-Password": process.env.HKA_TOKEN_PASSWORD!,
  },
  body: JSON.stringify({
    tipo: "01",
    numero: 1, punto: 1,
    cliente: { tipo: "contribuyente", ruc: "155596713-2-2015", dv: "59",
               razonSocial: "Mi Cliente S.A.", direccion: "Ave. La Paz" },
    items: [{ descripcion: "Servicio", cantidad: 1, precioUnitario: 100, tasaITBMS: "01" }],
  }),
});
if (res.status === 422) throw new Error(JSON.stringify(await res.json())); // errores de validación
const { CUFE } = await res.json();
```

### Python (requests)

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
r.raise_for_status()
print(r.json()["CUFE"])
```

## Generar clientes tipados desde OpenAPI

```bash
# TypeScript
openapi-generator-cli generate -i openapi.yaml -g typescript-fetch -o ./client-ts
# Python
openapi-generator-cli generate -i openapi.yaml -g python -o ./client-py
# Java, C#, PHP, Go, etc. — ver https://openapi-generator.tech/docs/generators
```

## Manejo de errores

| Código HTTP | Significado |
|---|---|
| `200` | Operación exitosa |
| `400` | JSON inválido o `tipo`/cliente no reconocido |
| `401` | Faltan los headers de credenciales |
| `422` | Validación fallida — el cuerpo incluye `fields[]` con cada problema |
| `502` | Error del servicio HKA (con `code`/`message`) o de red |

Como el gateway construye el documento con el builder seguro y valida antes de
enviar, los errores de cumplimiento se devuelven **antes** de tocar a HKA, con el
detalle por campo para corregirlos todos de una vez.
