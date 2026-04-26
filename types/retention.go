package types

// Retencion holds withholding tax information applied to the document.
type Retencion struct {
	CodigoRetencion CodigoRetencion
	FechaRetencion  string // YYYY-MM-DD
	MontoRetencion  string
}
