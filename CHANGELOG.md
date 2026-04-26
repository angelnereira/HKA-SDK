# Changelog

All notable changes to this project will be documented in this file.
Format: [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).
Versioning: [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - In Progress

### Added

- `Client` with `New()` and `NewDemo()` constructors — stateless, goroutine-safe
- `Send()` — submit any of the 10 fiscal document types with pre-flight validation
- `DocumentStatus()` — query the DGI authorization status of a document
- `Cancel()` — request cancellation within the 7-day window
- `DownloadXML()` — retrieve the signed XML as Base64
- `DownloadPDF()` — retrieve the PDF representation as Base64
- `SendEmail()` — send the document to a recipient email address
- `TrackEmail()` — track email delivery events by CUFE
- `RemainingFolios()` — query available folio count for the license
- `QueryRUC()` — look up a Panamanian RUC number and check digit
- Full type catalog in `types/` covering all HKA domain concepts
- Pre-flight validation in `validate/` covering all 48 rules from the HKA spec
- SOAP envelope builders in `soap/` — one per method, correct namespace per method
- Internal HTTP client with configurable timeout and connection pooling
- Helper functions: `FormatDateTime`, `FormatDate`, `PadDocNumber`, `PadPunto`, `PadSucursal`, `FormatAmount`, `CalculateITBMS`, `CalculatePrecioItem`, `CalculateValorTotal`, `ValidateCUFE`, `ValidateRUCFormat`, `IsWithin180Days`, `IsWithin7Days`
- Typed errors: `ValidationError`, `HKAError`, `NetworkError`
- Examples for: factura interna, factura exportacion, nota credito referenciada, factura gobierno, multi-tenant
- HKA wiki documentation in `hka-docs/`
- Unit test suite covering all validation rules
