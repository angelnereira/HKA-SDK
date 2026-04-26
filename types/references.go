package types

// DocFiscalReferenciado references a prior FE document for credit/debit notes (types 04/05).
// The referenced document must have been issued within the last 180 days.
type DocFiscalReferenciado struct {
	FechaEmisionDocFiscalReferenciado string // YYYY-MM-DDThh:mm:ss-05:00
	CufeFEReferenciada                string // exactly 66 characters
	NroFacturaPapel                   string // only when no CUFE is available
	NroFacturaImpFiscal               string // only when no CUFE is available
}
