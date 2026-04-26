package types

// DatosExportacion is required when DestinoOperacion = "2".
type DatosExportacion struct {
	CondicionesEntrega          Incoterm
	MonedaOperExportacion       CurrencyCode
	MonedaOperExportacionNonDef string // required when MonedaOperExportacion = "ZZZ"
	TipoDeCambio                string // required for non-USD currencies
	MontoMonedaExtranjera       string // tipoDeCambio * totalFactura
	PuertoEmbarque              string
}
