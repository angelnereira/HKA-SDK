// Package types contains all domain types, constants, and enumerations for the HKA SDK.
package types

import "fmt"

// TipoDocumento identifies the fiscal document type.
type TipoDocumento string

const (
	TipoDocFacturaInterna     TipoDocumento = "01"
	TipoDocFacturaImportacion TipoDocumento = "02"
	TipoDocFacturaExportacion TipoDocumento = "03"
	TipoDocNotaCreditoRef     TipoDocumento = "04"
	TipoDocNotaDebitoRef      TipoDocumento = "05"
	TipoDocNotaCreditoGen     TipoDocumento = "06"
	TipoDocNotaDebitoGen      TipoDocumento = "07"
	TipoDocFacturaZonaFranca  TipoDocumento = "08"
	TipoDocReembolso          TipoDocumento = "09"
	TipoDocFacturaExtranjera  TipoDocumento = "10"
)

// TipoEmision identifies the emission authorization type.
type TipoEmision string

const (
	EmisionAUPNormal       TipoEmision = "01"
	EmisionAUPContingencia TipoEmision = "02"
	EmisionAUPosteriorNorm TipoEmision = "03"
	EmisionAUPosteriorCont TipoEmision = "04"
)

// TipoClienteFE identifies the fiscal client category.
type TipoClienteFE string

const (
	ClienteContribuyente   TipoClienteFE = "01"
	ClienteConsumidorFinal TipoClienteFE = "02"
	ClienteGobierno        TipoClienteFE = "03"
	ClienteExtranjero      TipoClienteFE = "04"
)

// TipoContribuyente distinguishes natural persons from legal entities.
type TipoContribuyente string

const (
	ContribuyenteNatural  TipoContribuyente = "1"
	ContribuyenteJuridico TipoContribuyente = "2"
)

// NaturalezaOperacion describes the commercial nature of the transaction.
type NaturalezaOperacion string

const (
	NatVenta            NaturalezaOperacion = "01"
	NatExportacion      NaturalezaOperacion = "02"
	NatReexportacion    NaturalezaOperacion = "03"
	NatFuenteExtranjera NaturalezaOperacion = "04"
	NatTransferencia    NaturalezaOperacion = "10"
	NatDevolucion       NaturalezaOperacion = "11"
	NatConsignacion     NaturalezaOperacion = "12"
	NatRemesa           NaturalezaOperacion = "13"
	NatEntregaGratuita  NaturalezaOperacion = "14"
	NatCompra           NaturalezaOperacion = "20"
	NatImportacion      NaturalezaOperacion = "21"
)

// TipoOperacion indicates whether the operation is an outgoing sale or incoming purchase.
type TipoOperacion string

const (
	OperacionSalida  TipoOperacion = "1"
	OperacionEntrada TipoOperacion = "2"
)

// DestinoOperacion specifies whether the transaction destination is domestic or foreign.
type DestinoOperacion string

const (
	DestinoPanama     DestinoOperacion = "1"
	DestinoExtranjero DestinoOperacion = "2"
)

// FormatoCAFE specifies the CAFE (fiscal document) generation format.
type FormatoCAFE string

const (
	CAFESinGeneracion FormatoCAFE = "1"
	CAFECintaPapel    FormatoCAFE = "2"
	CAFEPapelCarta    FormatoCAFE = "3"
)

// EntregaCAFE specifies how the CAFE is delivered to the client.
type EntregaCAFE string

const (
	EntregaSinCAFE     EntregaCAFE = "1"
	EntregaEnPapel     EntregaCAFE = "2"
	EntregaElectronica EntregaCAFE = "3"
)

// EnvioContenedor indicates whether the container shipment is standard or exempt.
type EnvioContenedor string

const (
	ContenedorNormal     EnvioContenedor = "1"
	ContenedorExceptuado EnvioContenedor = "2"
)

// TipoVenta classifies what is being sold.
type TipoVenta string

const (
	VentaGiroNegocio  TipoVenta = "1"
	VentaActivoFijo   TipoVenta = "2"
	VentaBienesRaices TipoVenta = "3"
	VentaServicio     TipoVenta = "4"
)

// TipoSucursal classifies the branch type.
type TipoSucursal string

const (
	SucursalDetal    TipoSucursal = "1"
	SucursalMayorista TipoSucursal = "2"
)

// TasaITBMS represents the ITBMS (VAT) rate applied to an item.
type TasaITBMS string

const (
	ITBMSExento TasaITBMS = "00"
	ITBMS7      TasaITBMS = "01"
	ITBMS10     TasaITBMS = "02"
	ITBMS15     TasaITBMS = "03"
)

// Rate returns the decimal multiplier for computing valorITBMS.
func (t TasaITBMS) Rate() float64 {
	switch t {
	case ITBMS7:
		return 0.07
	case ITBMS10:
		return 0.10
	case ITBMS15:
		return 0.15
	default:
		return 0.00
	}
}

// TasaOTI represents other tax types (OTI — Otros Tipos de Impuesto).
type TasaOTI string

const (
	OTISume911          TasaOTI = "01"
	OTIPortabilidad     TasaOTI = "02"
	OTISeguro5          TasaOTI = "03"
	OTISeguroAutos1     TasaOTI = "04"
	OTISalidaAeropuerto TasaOTI = "05"
	OTICargoDesarrollo  TasaOTI = "06"
	OTISeguridad        TasaOTI = "07"
	OTIOtrosCargos      TasaOTI = "08"
	OTICombustible      TasaOTI = "09"
	OTIFECI             TasaOTI = "10"
	OTIIntereses        TasaOTI = "11"
)

// FormaPago enumerates accepted payment methods.
type FormaPago string

const (
	PagoCredito        FormaPago = "01"
	PagoEfectivo       FormaPago = "02"
	PagoTarjetaCredito FormaPago = "03"
	PagoTarjetaDebito  FormaPago = "04"
	PagoFidelizacion   FormaPago = "05"
	PagoVale           FormaPago = "06"
	PagoTarjetaRegalo  FormaPago = "07"
	PagoTransferencia  FormaPago = "08"
	PagoCheque         FormaPago = "09"
	PagoOtro           FormaPago = "99"
)

// TiempoPago indicates the payment timing.
type TiempoPago string

const (
	PagoInmediato TiempoPago = "1"
	PagoPlazo     TiempoPago = "2"
	PagoMixto     TiempoPago = "3"
)

// CodigoRetencion identifies the withholding tax type.
type CodigoRetencion string

const (
	RetencionServicioEstado100 CodigoRetencion = "1"
	RetencionVentaEstado50     CodigoRetencion = "2"
	RetencionNoDomiciliado100  CodigoRetencion = "3"
	RetencionCompraBienes50    CodigoRetencion = "4"
	RetencionTC_TD50           CodigoRetencion = "7"
	RetencionOtros             CodigoRetencion = "8"
)

// TipoIdentificacion applies to foreign clients (TipoClienteFE = "04").
type TipoIdentificacion string

const (
	IdPasaporte  TipoIdentificacion = "01"
	IdTributario TipoIdentificacion = "02"
	IdOtro       TipoIdentificacion = "99"
)

// RUCType distinguishes between natural-person and legal-entity RUCs.
type RUCType string

const (
	RUCNatural  RUCType = "1"
	RUCJuridico RUCType = "2"
)

// Incoterm enumerates valid delivery condition codes (Tabla 1 — condicionesEntrega).
type Incoterm string

const (
	IncoEXW Incoterm = "EXW"
	IncoFCA Incoterm = "FCA"
	IncoCPT Incoterm = "CPT"
	IncoCIP Incoterm = "CIP"
	IncoDPU Incoterm = "DPU"
	IncoDAP Incoterm = "DAP"
	IncoDDP Incoterm = "DDP"
	IncoFAS Incoterm = "FAS"
	IncoFOB Incoterm = "FOB"
	IncoCFR Incoterm = "CFR"
	IncoCIF Incoterm = "CIF"
)

// CurrencyCode represents an ISO 4217 currency code.
type CurrencyCode string

const (
	CurrencyUSD CurrencyCode = "USD"
	CurrencyEUR CurrencyCode = "EUR"
	CurrencyGBP CurrencyCode = "GBP"
	CurrencyCAD CurrencyCode = "CAD"
	CurrencyMXN CurrencyCode = "MXN"
	CurrencyCOP CurrencyCode = "COP"
	CurrencyVES CurrencyCode = "VES"
	CurrencyCRC CurrencyCode = "CRC"
	CurrencyBRL CurrencyCode = "BRL"
	CurrencyARS CurrencyCode = "ARS"
	CurrencyCLP CurrencyCode = "CLP"
	CurrencyPEN CurrencyCode = "PEN"
	CurrencyJPY CurrencyCode = "JPY"
	CurrencyCNY CurrencyCode = "CNY"
	CurrencyZZZ CurrencyCode = "ZZZ" // unlisted currency — requires MonedaOperExportacionNonDef
)

// CountryCode represents an ISO 3166-1 alpha-2 country code.
type CountryCode string

const (
	CountryPA CountryCode = "PA"
	CountryUS CountryCode = "US"
	CountryMX CountryCode = "MX"
	CountryCO CountryCode = "CO"
	CountryVE CountryCode = "VE"
	CountryCR CountryCode = "CR"
	CountryHN CountryCode = "HN"
	CountryGT CountryCode = "GT"
	CountrySV CountryCode = "SV"
	CountryNI CountryCode = "NI"
	CountryBR CountryCode = "BR"
	CountryAR CountryCode = "AR"
	CountryCL CountryCode = "CL"
	CountryPE CountryCode = "PE"
	CountryES CountryCode = "ES"
	CountryDE CountryCode = "DE"
	CountryFR CountryCode = "FR"
	CountryGB CountryCode = "GB"
	CountryCN CountryCode = "CN"
	CountryJP CountryCode = "JP"
	CountryZZ CountryCode = "ZZ" // unlisted country — requires PaisOtro
)

// ReturnCode is the service-level response code returned in the `codigo` field.
type ReturnCode string

const (
	CodeSuccess       ReturnCode = "200"
	CodeResend        ReturnCode = "300"
	CodeDuplicate     ReturnCode = "102"
	CodeNoFolios      ReturnCode = "119"
	CodeInvalidData   ReturnCode = "109"
	CodeUnauthorized  ReturnCode = "401"
	CodeNotFound      ReturnCode = "404"
	CodeCancelExpired ReturnCode = "112"
	CodeRefExpired    ReturnCode = "180"
)

// IsSuccess returns true when the code indicates a successful operation.
func (r ReturnCode) IsSuccess() bool {
	return r == CodeSuccess
}

// RequiresResend returns true when the document must be resent.
func (r ReturnCode) RequiresResend() bool {
	return r == CodeResend
}

func (r ReturnCode) String() string {
	return string(r)
}

// validIncoterms is the set of accepted Incoterm codes for fast lookup.
var validIncoterms = map[Incoterm]struct{}{
	IncoEXW: {}, IncoFCA: {}, IncoCPT: {}, IncoCIP: {}, IncoDPU: {},
	IncoDAP: {}, IncoDDP: {}, IncoFAS: {}, IncoFOB: {}, IncoCFR: {}, IncoCIF: {},
}

// IsValidIncoterm reports whether i is a recognized Incoterm.
func IsValidIncoterm(i Incoterm) bool {
	_, ok := validIncoterms[i]
	return ok
}

// validCurrencies is the set of accepted ISO 4217 codes plus "ZZZ".
var validCurrencies = map[CurrencyCode]struct{}{
	CurrencyUSD: {}, CurrencyEUR: {}, CurrencyGBP: {}, CurrencyCAD: {},
	CurrencyMXN: {}, CurrencyCOP: {}, CurrencyVES: {}, CurrencyCRC: {},
	CurrencyBRL: {}, CurrencyARS: {}, CurrencyCLP: {}, CurrencyPEN: {},
	CurrencyJPY: {}, CurrencyCNY: {}, CurrencyZZZ: {},
}

// IsValidCurrency reports whether c is a recognized currency code.
func IsValidCurrency(c CurrencyCode) bool {
	_, ok := validCurrencies[c]
	return ok
}

// validNaturalezas is the set of accepted NaturalezaOperacion codes.
var validNaturalezas = map[NaturalezaOperacion]struct{}{
	NatVenta: {}, NatExportacion: {}, NatReexportacion: {}, NatFuenteExtranjera: {},
	NatTransferencia: {}, NatDevolucion: {}, NatConsignacion: {}, NatRemesa: {},
	NatEntregaGratuita: {}, NatCompra: {}, NatImportacion: {},
}

// IsValidNaturaleza reports whether n is a recognized operation nature.
func IsValidNaturaleza(n NaturalezaOperacion) bool {
	_, ok := validNaturalezas[n]
	return ok
}

// String implements fmt.Stringer for debugging.
func (t TipoDocumento) String() string    { return string(t) }
func (t TipoEmision) String() string      { return string(t) }
func (t TipoClienteFE) String() string    { return string(t) }
func (t TipoContribuyente) String() string { return string(t) }
func (t TipoOperacion) String() string    { return string(t) }
func (t DestinoOperacion) String() string { return string(t) }
func (t TasaITBMS) String() string        { return string(t) }
func (t FormaPago) String() string        { return string(t) }
func (t TiempoPago) String() string       { return string(t) }
func (t CountryCode) String() string      { return string(t) }
func (t CurrencyCode) String() string     { return string(t) }

// formatNotEmpty is a compile-time guard; used in soap layer.
var _ = fmt.Sprintf
