package types

// Item represents a single line item in the fiscal document.
type Item struct {
	Descripcion             string // AN|2..500, required
	Codigo                  string // AN|1..20, optional
	UnidadMedida            string
	Cantidad                string // N|1..9 with 2..6 decimals
	FechaFabricacion        string // YYYY-MM-DD; required for pharma/food items
	FechaCaducidad          string // YYYY-MM-DD, conditional
	CodigoCPBSAbrev         string // AN|2; required when TipoClienteFE = "03"
	CodigoCPBS              string // AN|4; required when selling to public administration
	UnidadMedidaCPBS        string // required when TipoClienteFE = "03"
	InfoItem                string // optional, up to 5000 chars
	PrecioUnitario          string // N with 2..6 decimals
	PrecioUnitarioDescuento string // discount in Balboas; "0.000000" when none
	PrecioItem              string // = Cantidad * (PrecioUnitario - PrecioUnitarioDescuento)
	PrecioAcarreo           string
	PrecioSeguro            string
	ValorTotal              string // = PrecioItem + taxes
	CodigoGTIN              string
	CantGTINCom             string
	CodigoGTINInv           string
	CantGTINComInv          string
	TasaITBMS               TasaITBMS // required
	ValorITBMS              string    // = TasaITBMS.Rate() * PrecioItem
	TasaISC                 string
	ValorISC                string
	ListaOTI                []OTI
	Medicina                *Medicina
	Vehiculo                *Vehiculo
	PedidoItem              *PedidoComercialItem
}

// OTI represents one entry of Otros Tipos de Impuesto applied to an item.
type OTI struct {
	Tasa      TasaOTI
	ValorOTI  string
	FlagMonto string // "1" = amount, "2" = percentage
}

// Medicina holds pharmaceutical batch information for controlled medicines.
type Medicina struct {
	NroLote           string // required
	CantProductosLote string // required
}

// Vehiculo holds information required when the item is a new vehicle.
type Vehiculo struct {
	Marca           string
	Modelo          string
	Anio            string
	Color           string
	NumeroVIN       string
	NumeroMotor     string
	NumeroPlaca     string
	TipoCombustible string
	PrecioBase      string
	NumeroDocAduana string
}

// PedidoComercialItem links a line item to a purchase order line.
type PedidoComercialItem struct {
	NumeroPedido string
	NroLinea     string
}
