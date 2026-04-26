# Manual de Integración Directa al Web Service

> Fuente: https://felwiki.thefactoryhka.com.pa/manual_de_integracion_directa_al_ws

## Métodos del Web Service

El servicio web esta implementado para realizar la comunicación entre los sistemas de los contribuyentes y el sistema de **facturación electrónica de The Factory HKA**.

Para el consumo de este servicio es **necesario e imprescindible** el uso de las credenciales de acceso conformadas por el **Token de Empresa** y **Token Password**. Por supuesto también es necesario conocer el URL del Web Service del ambiente de integración.

El servicio web consta de los siguientes métodos:

1. **Enviar()** — Envío del documento electrónico
2. **EstadoDocumento()** — Consulta estatus
3. **Anulacion()** — Anulación de documentos (método real: `AnulacionDocumento`)
4. **DescargaXML()** — Descarga del XML
5. **FoliosRestantes()** — Consulta folios disponibles
6. **EnvioCorreo()** — Envío por correo
7. **DescargaPDF()** — Descarga del CAFE en PDF
8. **RastreoCorreo()** — Traza de correos
9. **ConsultarRucDV()** — Consulta dígito verificador del RUC

## Formatos de Parámetros

Los parámetros a incorporar a los métodos deben cumplir con el formato y directivas que correspondan, tomando en cuenta la siguiente norma:

| Formato | Descripción |
| --- | --- |
| A | Carácter alfabético |
| N | Carácter numérico |
| AN | Carácter alfanumérico |
| A3 | 3 caracteres alfabéticos (longitud fija) |
| N3 | 3 caracteres numéricos (longitud fija) |
| AN3 | 3 caracteres alfanuméricos (longitud fija) |
| A..3 | Hasta 3 caracteres alfabéticos (longitud variable) |
| N..3 | Hasta 3 caracteres numéricos (longitud variable) |
| AN..3 | Hasta 3 caracteres alfanuméricos (longitud variable) |
| N..8/5.2 | Hasta 8 caracteres alfanuméricos, conformados por (hasta) 5 enteros y dos decimales separados por punto |

## Requerido

| Requerido | Descripción |
| --- | --- |
| SI | Si es requerido |
| NO | No es requerido |
| C/C | Requerido con condición |

## ⚠️ NOTA IMPORTANTE

> Al momento de crear la estructura del XML en el método de **Enviar**, los campos que no sean requeridos en caso de que correspondan, **no informar o no dejar sin valor esos campos**, para que se procese el XML en el servicio sin inconveniente alguno.

## Documentación adicional

- [ANEXOS TÉCNICOS (PDF 2.2 MB)](https://felwiki.thefactoryhka.com.pa/_media/anexos_tecnico-hka.pdf)
- [Catálogo de códigos de retorno del servicio (PDF 72.7 KB)](https://felwiki.thefactoryhka.com.pa/_media/catalogo_de_codigos_de_retorno_del_servicio-08-2023.pdf)
