package types

// Retencion holds withholding tax information applied to the document.
type Retencion struct {
	CodigoRetencion CodigoRetencion `json:"codigoRetencion"`
	FechaRetencion  string          `json:"fechaRetencion"` // YYYY-MM-DD
	MontoRetencion  string          `json:"montoRetencion"`
}
