package types

// DocumentoElectronico is the root payload for Send().
type DocumentoElectronico struct {
	CodigoSucursalEmisor string
	TipoSucursal         TipoSucursal
	DatosTransaccion     DatosTransaccion
	ListaItems           []Item
	TotalesSubTotales    TotalesSubTotales
	AutorizadosDescarga  []AutorizadoDescarga
	PedidoGlobal         *PedidoComercialGlobal
	InfoLogistica        *InfoLogistica
	InfoEntrega          *InfoEntrega
	UsoPosterior         *UsoPosterior
}

// DatosTransaccion holds all transaction-level fields for a fiscal document.
type DatosTransaccion struct {
	TipoEmision                   TipoEmision
	FechaInicioContingencia       string // required when TipoEmision = 02/04; format YYYY-MM-DDThh:mm:ss-05:00
	MotivoContingencia            string // required when TipoEmision = 02/04; 15..250 chars
	TipoDocumento                 TipoDocumento
	NumeroDocumentoFiscal         string // N|10 zero-padded
	PuntoFacturacionFiscal        string // N|3, not "000"
	FechaEmision                  string // YYYY-MM-DDThh:mm:ss-05:00
	FechaSalida                   string // optional, same format
	NaturalezaOperacion           NaturalezaOperacion
	TipoOperacion                 TipoOperacion
	DestinoOperacion              DestinoOperacion
	FormatoCAFE                   FormatoCAFE
	EntregaCAFE                   EntregaCAFE
	EnvioContenedor               EnvioContenedor
	ProcesoGeneracion             string // always "1"
	TipoVenta                     TipoVenta // omit when TipoOperacion = "2"
	InformacionInteres            string // optional, up to 5000 chars
	Cliente                       Cliente
	ListaDocsFiscalReferenciados  []DocFiscalReferenciado // required for types 04/05; absent for 06/07
	DatosFacturaExportacion       *DatosExportacion       // required when DestinoOperacion = "2"
	Retencion                     *Retencion
}

// Cliente holds the buyer/recipient information for the fiscal document.
type Cliente struct {
	TipoClienteFE              TipoClienteFE
	TipoContribuyente          TipoContribuyente  // omit when TipoClienteFE = "04"
	NumeroRUC                  string             // omit when TipoClienteFE = "04"
	DigitoVerificadorRUC       string             // omit when TipoClienteFE = "04"
	RazonSocial                string             // required when TipoClienteFE = "01"/"03"
	Direccion                  string             // required when TipoClienteFE = "01"/"03"
	CodigoUbicacion            string             // omit when TipoClienteFE = "04"
	Provincia                  string             // omit when TipoClienteFE = "04"
	Distrito                   string             // omit when TipoClienteFE = "04"
	Corregimiento              string             // omit when TipoClienteFE = "04"
	TipoIdentificacion         TipoIdentificacion // required when TipoClienteFE = "04"
	NroIdentificacionExtranjero string            // required when TipoClienteFE = "04"
	PaisExtranjero             string             // only when TipoIdentificacion = "01" (passport)
	Telefono1                  string
	CorreoElectronico1         string
	CorreoElectronico2         string
	CorreoElectronico3         string
	Pais                       CountryCode // "PA" when DestinoOperacion = "1"; "ZZ" if not in catalog
	PaisOtro                   string      // required when Pais = "ZZ"
}

// AutorizadoDescarga represents a person authorized to download the document.
type AutorizadoDescarga struct {
	NombreAutorizado string
}

// PedidoComercialGlobal contains global purchase order data.
type PedidoComercialGlobal struct {
	NumeroPedido  string
	FechaPedido   string
	NroLinea      string
	MontoTotal    string
	MonedaPedido  string
}

// UsoPosterior is required for documents with TipoEmision = "03" or "04".
type UsoPosterior struct {
	CUFE string
}
