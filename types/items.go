package types

// Item represents a single line item in the fiscal document.
type Item struct {
	Descripcion             string               `json:"descripcion"`      // AN|2..500, required
	Codigo                  string               `json:"codigo,omitempty"` // AN|1..20, optional
	UnidadMedida            string               `json:"unidadMedida,omitempty"`
	Cantidad                string               `json:"cantidad"`                          // N|1..9 with 2..6 decimals
	FechaFabricacion        string               `json:"fechaFabricacion,omitempty"`        // YYYY-MM-DD; required for pharma/food items
	FechaCaducidad          string               `json:"fechaCaducidad,omitempty"`          // YYYY-MM-DD, conditional
	CodigoCPBSAbrev         string               `json:"codigoCPBSAbrev,omitempty"`         // AN|2; required when TipoClienteFE = "03"
	CodigoCPBS              string               `json:"codigoCPBS,omitempty"`              // AN|4; required when selling to public administration
	UnidadMedidaCPBS        string               `json:"unidadMedidaCPBS,omitempty"`        // required when TipoClienteFE = "03"
	InfoItem                string               `json:"infoItem,omitempty"`                // optional, up to 5000 chars
	PrecioUnitario          string               `json:"precioUnitario"`                    // N with 2..6 decimals
	PrecioUnitarioDescuento string               `json:"precioUnitarioDescuento,omitempty"` // discount in Balboas; "0.000000" when none
	PrecioItem              string               `json:"precioItem"`                        // = Cantidad * (PrecioUnitario - PrecioUnitarioDescuento)
	PrecioAcarreo           string               `json:"precioAcarreo,omitempty"`
	PrecioSeguro            string               `json:"precioSeguro,omitempty"`
	ValorTotal              string               `json:"valorTotal"` // = PrecioItem + taxes
	CodigoGTIN              string               `json:"codigoGTIN,omitempty"`
	CantGTINCom             string               `json:"cantGTINCom,omitempty"`
	CodigoGTINInv           string               `json:"codigoGTINInv,omitempty"`
	CantGTINComInv          string               `json:"cantGTINComInv,omitempty"`
	TasaITBMS               TasaITBMS            `json:"tasaITBMS"`  // required
	ValorITBMS              string               `json:"valorITBMS"` // = TasaITBMS.Rate() * PrecioItem
	TasaISC                 string               `json:"tasaISC,omitempty"`
	ValorISC                string               `json:"valorISC,omitempty"`
	ListaOTI                []OTI                `json:"listaOTI,omitempty"`
	Medicina                *Medicina            `json:"medicina,omitempty"`
	Vehiculo                *Vehiculo            `json:"vehiculo,omitempty"`
	PedidoItem              *PedidoComercialItem `json:"pedidoItem,omitempty"`
}

// OTI represents one entry of Otros Tipos de Impuesto applied to an item.
type OTI struct {
	Tasa      TasaOTI `json:"tasa"`
	ValorOTI  string  `json:"valorOTI"`
	FlagMonto string  `json:"flagMonto"` // "1" = amount, "2" = percentage
}

// Medicina holds pharmaceutical batch information for controlled medicines.
type Medicina struct {
	NroLote           string `json:"nroLote"`           // required
	CantProductosLote string `json:"cantProductosLote"` // required
}

// Vehiculo holds information required when the item is a new vehicle.
type Vehiculo struct {
	Marca           string `json:"marca,omitempty"`
	Modelo          string `json:"modelo,omitempty"`
	Anio            string `json:"anio,omitempty"`
	Color           string `json:"color,omitempty"`
	NumeroVIN       string `json:"numeroVIN,omitempty"`
	NumeroMotor     string `json:"numeroMotor,omitempty"`
	NumeroPlaca     string `json:"numeroPlaca,omitempty"`
	TipoCombustible string `json:"tipoCombustible,omitempty"`
	PrecioBase      string `json:"precioBase,omitempty"`
	NumeroDocAduana string `json:"numeroDocAduana,omitempty"`
}

// PedidoComercialItem links a line item to a purchase order line.
type PedidoComercialItem struct {
	NumeroPedido string `json:"numeroPedido"`
	NroLinea     string `json:"nroLinea"`
}
