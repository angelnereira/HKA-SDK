# Análisis del proyecto y hoja de ruta

> Estado al 2026-06-02. Objetivo: un SDK basado en Go para The Factory HKA
> (facturación electrónica, DGI Panamá) que sea **fácil de usar**, **políglota**
> (importable desde TypeScript, JavaScript, Python, Go, etc.) y en el que sea
> **casi imposible** generar un documento que incumpla las normas de HKA.

## 1. Estado actual

El núcleo en Go está sólido y compila/pasa pruebas:

- **Cliente** stateless y seguro para concurrencia (`hka.go`) con las 9 operaciones
  del servicio de integración: `Send`, `DocumentStatus`, `Cancel`, `DownloadXML`,
  `DownloadPDF`, `SendEmail`, `TrackEmail`, `RemainingFolios`, `QueryRUC`.
- **Los 10 tipos de documento** como constantes tipadas (`types/codes.go`).
- **Validación pre-flight** (~48 reglas) antes de cualquier llamada HTTP (`validate/`).
- **Constructores SOAP** por método, parser de respuestas, errores tipados.
- Catálogos completos (Incoterms, monedas, países, ITBMS, formas de pago, retención).
- Documentación de la wiki HKA en `hka-docs/` y ejemplos.

## 2. Brechas detectadas

| # | Brecha | Impacto | Estado |
|---|--------|---------|--------|
| 1 | No es importable fuera de Go (no políglota) | Bloquea el objetivo principal | **Resuelto** — gateway REST/JSON + OpenAPI (`gateway/`, `cmd/hka-gateway`, `openapi.yaml`) |
| 7 | Catálogos (ubicación, CPBS) y formatos (cédula, RUC, CUFE, CAFE, ITBMS) no incluidos | Cumplimiento y ergonomía | **Resuelto** — paquete `catalog` + `tools/gencatalog` |
| 2 | Construir un documento es manual y propenso a error (~60 líneas, montos calculados a mano) | Riesgo de incumplimiento | **Resuelto** — paquete `docbuilder` |
| 3 | Sin motor de auto-cálculo de ITBMS/totales | Descuadres y redondeos | **Resuelto** — `docbuilder.Build()` |
| 4 | `listaPagoPlazo` emitía `fechaPago`/`valorCuotaPlazo` en lugar de `fechaVenceCuota`/`valorCuota` | Documentos a plazo rechazados | **Resuelto** — corregido en `soap/enviar.go` |
| 5 | Falta `LICENSE`, CI/CD, idempotencia/reintentos en `Send` | Madurez de producción | Pendiente |
| 6 | Camino PAC-directo (firma con certificado) documentado pero no implementado | Funcionalidad avanzada | Pendiente (opcional) |

## 3. Lo implementado en esta iteración

### Capa segura de documentos: `docbuilder`

Un constructor por tipo de documento que:

1. Rellena automáticamente los campos obligatorios de la transacción con el valor
   correcto para ese tipo (`ProcesoGeneracion`, formatos CAFE, naturaleza, destino…).
2. Recibe ítems con valores naturales (`Cantidad`, `PrecioUnitario`, `Descuento`,
   `TasaITBMS`) y **calcula** `precioItem`, `valorITBMS` y `valorTotal`.
3. **Calcula todos los totales** (`TotalPrecioNeto`, `TotalITBMS`,
   `TotalMontoGravado`, `TotalFactura`, `TotalTodosItems`, `NroItems`) de forma
   coherente con los ejemplos oficiales de HKA (redondeo a 6 decimales por ítem,
   2 decimales en totales).
4. Infiere las condiciones de pago (contado / plazo / mixto).
5. Ejecuta la validación completa como última red de seguridad y devuelve un
   `*ValidationError` con todos los problemas si algo no cuadra.

Constructores de cliente (`ClienteContribuyente`, `ClienteConsumidorFinal`,
`ClienteGobierno`, `ClienteExtranjero`, …) que rellenan exactamente los campos
que cada categoría exige.

Resultado: una factura válida se genera en pocas líneas y es **muy difícil**
producir un documento que no cumpla.

## 4. Próximos pasos (acordados)

### Fase siguiente — Gateway REST/JSON + OpenAPI (políglota)

Para que cualquier lenguaje pueda integrarse "casi inevitablemente":

1. **Servicio gateway en Go** que envuelve `docbuilder` + `validate` + cliente HKA
   y expone endpoints JSON limpios (uno por operación y por tipo de documento),
   ocultando por completo el SOAP.
2. **Especificación OpenAPI 3.1** del gateway, desde la cual se autogeneran
   clientes tipados para TypeScript, JavaScript, Python, Java, etc.
3. Imagen de contenedor y ejemplos de consumo por lenguaje.

Con esto el SDK Go sigue siendo la fuente de verdad y todos los demás lenguajes
consumen la misma lógica de validación y cálculo a través de HTTP/JSON.

### Endurecimiento de producción

- `LICENSE`, pipeline de CI (build + vet + test), versionado semántico.
- Reintentos con backoff e idempotencia en `Send`.
- Pruebas de integración contra el entorno demo (tras `-tags=integration`).
