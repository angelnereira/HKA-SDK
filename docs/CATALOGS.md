# Catálogos y formatos (`catalog`)

El paquete [`catalog`](../catalog) reúne los catálogos de referencia y los
ayudantes de formato que exige el esquema de factura electrónica de HKA / DGI.
Todo se incluye como **tipos dentro del SDK** y los catálogos grandes pueden
regenerarse desde su fuente oficial.

## Qué incluye

| Área | API | Fuente / estado |
|------|-----|-----------------|
| Ubicación geográfica (provincia-distrito-corregimiento) | `Provincias`, `ProvinciaByCodigo`, `ParseUbicacion`, `ValidateUbicacion`, `Ubicacion.Resolve` | INEC — 13 provincias autoritativas embebidas; distritos/corregimientos regenerables |
| Cédula | `ParseCedula`, `ValidateCedula`, `DescribePrefijo`, `SpecialCedulaPrefijos` | Formato `PREFIJO-TOMO-PARTIDA` |
| RUC | `ParseRUC`, `ValidateRUC` (natural = cédula, jurídico = `FOLIO-ROLLO-AÑO`) | Estructura; el DV autoritativo se obtiene con `client.QueryRUC` |
| CUFE / CAFE | `ValidateCUFE`, `CUFELength`, `DescribeCUFE`, `DescribeCAFE` | Validación de forma (66 chars, prefijo `FE`) |
| ITBMS | `SugerirCategoria`, `SugerirTasa`, `ITBMSCategoria`, `PorcentajeDeTasa` | DGI — 7% / 10% / 15% / exento |
| CPBS | `ValidateCPBS`, `AbrevForCPBS`, `CPBSByCodigo`, `CPBSCategorias` | Estructura 2/4 dígitos; catálogo regenerable |

## Formatos

### codigoUbicacion
`provincia-distrito-corregimiento`, p. ej. `8-8-7` (Panamá → Panamá → Bella Vista).
Provincia es un código numérico 1..13 (10 provincias + 3 comarcas; 13 = Panamá Oeste).

### Cédula
`PREFIJO-TOMO-PARTIDA`, p. ej. `8-123-456`. El prefijo es el código de provincia de
inscripción (1..13) o un prefijo especial: `PE` (nacido en el extranjero/antigua Zona
del Canal), `E` (extranjero residente), `N` (naturalizado), `AV` (casos especiales).

### RUC
- **Natural**: el RUC es la cédula.
- **Jurídico**: número del Registro Público, p. ej. `155596713-2-2015`.
- El **dígito verificador (DV)** lo emite la DGI y viaja aparte en
  `digitoVerificadorRUC`. No se calcula localmente; usa `QueryRUC` para obtenerlo.

### CUFE / CAFE
- **CUFE** (Código Único de Factura Electrónica): identificador único de 66
  caracteres que comienza con `FE`, generado y firmado por HKA/DGI. El cliente
  nunca lo construye; el SDK valida su forma.
- **CAFE** (Comprobante Auxiliar de Factura Electrónica): el PDF legible del
  documento autorizado.
- Composición campo a campo y algoritmo del DV: ver la *Ficha Técnica de Factura
  Electrónica* de la DGI (https://dgi.mef.gob.pa/_7facturaelectronica/ftPAC).

### ITBMS
- **7%** tasa general.
- **10%** bebidas alcohólicas y servicio de hospedaje/alojamiento.
- **15%** productos derivados del tabaco.
- **0% / exento** bienes y servicios exentos por ley.

`SugerirTasa(descripcion)` es una ayuda heurística para captura de datos; la
responsabilidad legal de aplicar la tasa correcta es del emisor.

## Regenerar el catálogo geográfico completo

Los catálogos grandes (distritos/corregimientos, CPBS) se incluyen como muestra
verificada y se regeneran a partir de una exportación oficial normalizada:

```bash
go run ./tools/gencatalog -csv ubicaciones.csv -out catalog/data/ubicaciones.json
go test ./catalog/   # verificar
```

El CSV debe tener la cabecera:

```
provincia_codigo,provincia_nombre,prefijo_cedula,distrito_codigo,distrito_nombre,corregimiento_codigo,corregimiento_nombre
```

## Fuentes oficiales

- INEC — División Político-Administrativa: https://www.inec.gob.pa/
- Catálogo unificado (Alanube): https://developer.alanube.co/v1.0-PAN/reference/catalogo-unificado-de-provincias-distritos-y-corregimientos
- DGI — Generalidades del ITBMS: https://dgi.mef.gob.pa/itbms/Generalidades
- DGI — Ficha Técnica de Factura Electrónica (CUFE, DV): https://dgi.mef.gob.pa/_7facturaelectronica/ftPAC
- HKA — Anexos Técnicos: https://felwiki.thefactoryhka.com.pa/_media/anexos_tecnico-hka.pdf
