package types

// DocFiscalReferenciado references a prior FE document for credit/debit notes (types 04/05).
// The referenced document must have been issued within the last 180 days.
type DocFiscalReferenciado struct {
	FechaEmisionDocFiscalReferenciado string `json:"fechaEmisionDocFiscalReferenciado,omitempty"` // YYYY-MM-DDThh:mm:ss-05:00
	CufeFEReferenciada                string `json:"cufeFEReferenciada,omitempty"`                // exactly 66 characters
	NroFacturaPapel                   string `json:"nroFacturaPapel,omitempty"`                   // only when no CUFE is available
	NroFacturaImpFiscal               string `json:"nroFacturaImpFiscal,omitempty"`               // only when no CUFE is available
}
