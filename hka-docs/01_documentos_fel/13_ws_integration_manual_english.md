# WS Integration Manual — English Version

> Source: https://felwiki.thefactoryhka.com.pa/ws_integration_manual_-_english_version

## Web Services Methods

The web service is implemented to perform the communication between the taxpayers' systems and the E-billing system of **The Factory HKA**.

For the use of this service it is necessary to use the access tokens created by The Factory HKA as well as the web service URL.

## The web service consists of the following methods

1. **Enviar()** — Send electronic document
2. **EstadoDocumento()** — Query document status
3. **Anulacion()** — Annul/cancel document
4. **DescargaXML()** — Download XML
5. **FoliosRestantes()** — Query remaining folios
6. **EnvioCorreo()** — Send document by email
7. **DescargaPDF()** — Download PDF
8. **RastreoCorreo()** — Email tracking
9. **ConsultaRucDV()** — Query RUC check digit

## Parameter Format Standard

| Format | Description |
| --- | --- |
| A | Alphabetic field |
| N | Numeric field |
| AN | Alphanumeric field |
| A3 | 3 alphabetical field (fixed length) |
| N3 | 3 number field (fixed length) |
| AN3 | 3 alphanumeric value field (fixed length) |
| A...3 | Up to 3 alphabetic characters (variable length) |
| N...3 | Up to 3 numeric characters (variable length) |
| AN...3 | Up to 3 alphanumeric characters (variable length) |
| N..8/5.2 | Up to 8 alphanumeric characters, consisting of (up to) 5 integers and two decimals separated by a period |

## Required

| Required | Description |
| --- | --- |
| YES | Required |
| NO | Not required |
| C/C | Condition based field |

---

## Notes

The English version of this wiki only includes the overview page above. Each individual method (Enviar, EstadoDocumento, etc.) has an English equivalent at `https://felwiki.thefactoryhka.com.pa/{metodo}_english` — the technical content is **identical** to the Spanish version, just translated. Refer to the Spanish method documentation in `01_documentos_fel/02_metodo_enviar.md` through `10_metodo_consultarucdv.md` for the complete specifications with the same parameter tables, XML examples, and response structures.
