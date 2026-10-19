package types

// DatosExportacion is required when DestinoOperacion = "2".
type DatosExportacion struct {
	CondicionesEntrega          Incoterm     `json:"condicionesEntrega"`
	MonedaOperExportacion       CurrencyCode `json:"monedaOperExportacion"`
	MonedaOperExportacionNonDef string       `json:"monedaOperExportacionNonDef,omitempty"` // required when MonedaOperExportacion = "ZZZ"
	TipoDeCambio                string       `json:"tipoDeCambio,omitempty"`                // required for non-USD currencies
	MontoMonedaExtranjera       string       `json:"montoMonedaExtranjera,omitempty"`       // tipoDeCambio * totalFactura
	PuertoEmbarque              string       `json:"puertoEmbarque,omitempty"`
}
