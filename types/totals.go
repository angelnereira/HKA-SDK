package types

// TotalesSubTotales aggregates all monetary totals for the document.
type TotalesSubTotales struct {
	TotalPrecioNeto       string                 `json:"totalPrecioNeto"`
	TotalITBMS            string                 `json:"totalITBMS"`
	TotalISC              string                 `json:"totalISC,omitempty"`
	TotalMontoGravado     string                 `json:"totalMontoGravado"`
	TotalDescuento        string                 `json:"totalDescuento,omitempty"`
	TotalAcarreoCobrado   string                 `json:"totalAcarreoCobrado,omitempty"`
	ValorSeguroCobrado    string                 `json:"valorSeguroCobrado,omitempty"`
	TotalOtrosGastos      string                 `json:"totalOtrosGastos,omitempty"`
	TotalFactura          string                 `json:"totalFactura"`
	TotalValorRecibido    string                 `json:"totalValorRecibido,omitempty"`
	Vuelto                string                 `json:"vuelto,omitempty"`
	TiempoPago            TiempoPago             `json:"tiempoPago"`
	NroItems              int                    `json:"nroItems"`
	TotalTodosItems       string                 `json:"totalTodosItems"`
	ListaFormaPago        []FormaPagoItem        `json:"listaFormaPago"`
	ListaPagoPlazo        []CuotaPlazo           `json:"listaPagoPlazo,omitempty"`
	ListaTotalOTI         []TotalOTI             `json:"listaTotalOTI,omitempty"`
	DescuentoBonificacion *DescuentoBonificacion `json:"descuentoBonificacion,omitempty"`
	Retencion             *Retencion             `json:"retencion,omitempty"`
}

// FormaPagoItem represents one payment method entry.
type FormaPagoItem struct {
	FormaPagoFact    FormaPago `json:"formaPagoFact"`
	DescFormaPago    string    `json:"descFormaPago,omitempty"` // required when FormaPagoFact = "99"
	ValorCuotaPagada string    `json:"valorCuotaPagada"`
}

// CuotaPlazo represents one installment in a deferred payment schedule.
// It maps to the <pagoPlazo> element: FechaPago -> fechaVenceCuota and
// ValorCuotaPlazo -> valorCuota.
type CuotaPlazo struct {
	FechaPago        string `json:"fechaPago"`                  // fechaVenceCuota, YYYY-MM-DDThh:mm:ss-05:00
	ValorCuotaPlazo  string `json:"valorCuotaPlazo"`            // valorCuota
	DescripcionCuota string `json:"descripcionCuota,omitempty"` // not part of the HKA schema; retained for caller notes
}

// TotalOTI aggregates the total for one OTI type across all items.
type TotalOTI struct {
	Tasa          TasaOTI `json:"tasa"`
	ValorTotalOTI string  `json:"valorTotalOTI"`
}

// DescuentoBonificacion holds a document-level discount or bonus.
type DescuentoBonificacion struct {
	MontoDescuento       string `json:"montoDescuento"`
	DescripcionDescuento string `json:"descripcionDescuento"`
}
