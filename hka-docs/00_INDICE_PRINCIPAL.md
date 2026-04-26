# WIKI THE FACTORY HKA PANAMÁ — Documentación Completa

> **Fuente**: https://felwiki.thefactoryhka.com.pa/
> **Extraído**: 20 de abril de 2026
> **Propósito**: Documentación de referencia para integración con el PAC The Factory HKA (Panamá — DGI / Facturación Electrónica)

---

## ¡Bienvenido!

Esta sección es para clientes de integración y contiene toda la información referente a la facturación electrónica y manuales útiles en el proceso de integración.

---

## Estructura de la documentación

### 📁 01_documentos_fel/
Manuales principales de integración con el servicio de facturación electrónica.

- `01_manual_integracion_ws.md` — Visión general del Web Service y formatos de parámetros
- `02_metodo_enviar.md` — **Método principal** para envío de documentos electrónicos
- `03_metodo_estadodocumento.md` — Consulta del estatus de un documento
- `04_metodo_anulacion.md` — Anulación de documentos autorizados
- `05_metodo_descargaxml.md` — Descarga del XML firmado
- `06_metodo_foliosrestantes.md` — Consulta de folios disponibles en la licencia
- `07_metodo_enviocorreo.md` — Envío del documento por correo
- `08_metodo_descargapdf.md` — Descarga del CAFE en PDF
- `09_metodo_rastreocorreo.md` — Consulta de traza de envíos por correo
- `10_metodo_consultarucdv.md` — Consulta de dígito verificador de RUC
- `11_manual_portal_dfactura.md` — Manual del Portal Dfactura (usuario final)
- `12_listado_documentos_fel.md` — Listado mínimo de documentos exigidos
- `13_ws_integration_manual_english.md` — Versión en inglés

### 📁 02_ejemplos_codigo/
Ejemplos de integración con el Web Service SOAP.

- `01_lenguaje_php.md`
- `02_lenguaje_csharp.md`
- `03_lenguaje_java.md`
- `04_lenguaje_python.md`

### 📁 03_documentos_fiscales/
Ejemplos XML completos para cada tipo de documento fiscal.

1. `01_factura_operacion_interna.md`
2. `02_factura_importacion.md`
3. `03_factura_exportacion.md`
4. `04_nota_credito_referente_FE.md`
5. `05_nota_debito_referente_FE.md`
6. `06_nota_credito_generica.md`
7. `07_nota_debito_generica.md`
8. `08_factura_zona_franca.md`
9. `09_reembolso.md`
10. `10_factura_cliente_extranjero.md`
11. `11_factura_cliente_gobierno.md`
12. `12_factura_consumidor_final.md`
13. `13_factura_descuento_global.md`
14. `14_factura_descuento_por_item.md`
15. `15_factura_pago_plazo.md`
16. `16_factura_otro_tipo_impuesto.md`
17. `17_factura_retencion.md`

### 📁 04_documentos_pac/
- `01_instructivo_servicios_web_pac.md`

### 📁 05_ayuda/
- `01_preguntas_frecuentes.md`

---

## URLs clave

| Recurso | URL |
|---|---|
| Wiki oficial | https://felwiki.thefactoryhka.com.pa/ |
| WS Demo (integración) | https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc?singleWsdl |
| WS Demo (WSDL) | https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc?wsdl |
| Anexos Técnicos | https://felwiki.thefactoryhka.com.pa/_media/anexos_tecnico-hka.pdf |
| Catálogo códigos de retorno | https://felwiki.thefactoryhka.com.pa/_media/catalogo_de_codigos_de_retorno_del_servicio-08-2023.pdf |
| Postman Collection | https://felwiki.thefactoryhka.com.pa/_media/demo.postman_collection.zip |

## Credenciales

Para consumir el servicio web se necesitan:
- `tokenEmpresa` — Suministrado por The Factory HKA
- `tokenPassword` — Suministrado por The Factory HKA
- URL del Web Service del ambiente de integración

---

## Tipos de Documento (tipoDocumento)

| Código | Descripción |
|---|---|
| 01 | Factura de operación interna |
| 02 | Factura de importación |
| 03 | Factura de exportación |
| 04 | Nota de Crédito referente a una FE |
| 05 | Nota de Débito referente a una FE |
| 06 | Nota de Crédito genérica |
| 07 | Nota de Débito genérica |
| 08 | Factura de Zona Franca |
| 09 | Reembolso |
| 10 | Factura de operación extranjera |

## Tipos de Emisión (tipoEmision)

| Código | Descripción |
|---|---|
| 01 | Autorización de Uso Previa, operación normal |
| 02 | Autorización de Uso Previa, operación en contingencia |
| 03 | Autorización de Uso Posterior, operación normal |
| 04 | Autorización de Uso Posterior, operación en contingencia |

## Tipos de Cliente (tipoClienteFE)

| Código | Descripción |
|---|---|
| 01 | Contribuyente |
| 02 | Consumidor final |
| 03 | Gobierno |
| 04 | Extranjero |
