package types

// TotalesSubTotales aggregates all monetary totals for the document.
type TotalesSubTotales struct {
	TotalPrecioNeto       string
	TotalITBMS            string
	TotalISC              string
	TotalMontoGravado     string
	TotalDescuento        string
	TotalAcarreoCobrado   string
	ValorSeguroCobrado    string
	TotalOtrosGastos      string
	TotalFactura          string
	TotalValorRecibido    string
	Vuelto                string
	TiempoPago            TiempoPago
	NroItems              int
	TotalTodosItems       string
	ListaFormaPago        []FormaPagoItem
	ListaPagoPlazo        []CuotaPlazo
	ListaTotalOTI         []TotalOTI
	DescuentoBonificacion *DescuentoBonificacion
}

// FormaPagoItem represents one payment method entry.
type FormaPagoItem struct {
	FormaPagoFact    FormaPago
	DescFormaPago    string // required when FormaPagoFact = "99"
	ValorCuotaPagada string
}

// CuotaPlazo represents one installment in a deferred payment schedule.
// It maps to the <pagoPlazo> element: FechaPago -> fechaVenceCuota and
// ValorCuotaPlazo -> valorCuota.
type CuotaPlazo struct {
	FechaPago        string // fechaVenceCuota, YYYY-MM-DDThh:mm:ss-05:00
	ValorCuotaPlazo  string // valorCuota
	DescripcionCuota string // not part of the HKA schema; retained for caller notes
}

// TotalOTI aggregates the total for one OTI type across all items.
type TotalOTI struct {
	Tasa          TasaOTI
	ValorTotalOTI string
}

// DescuentoBonificacion holds a document-level discount or bonus.
type DescuentoBonificacion struct {
	MontoDescuento       string
	DescripcionDescuento string
}
