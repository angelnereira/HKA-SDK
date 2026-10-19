package types

// DocumentoElectronico is the root payload for Send().
type DocumentoElectronico struct {
	CodigoSucursalEmisor string                 `json:"codigoSucursalEmisor,omitempty"`
	TipoSucursal         TipoSucursal           `json:"tipoSucursal,omitempty"`
	DatosTransaccion     DatosTransaccion       `json:"datosTransaccion"`
	ListaItems           []Item                 `json:"listaItems"`
	TotalesSubTotales    TotalesSubTotales      `json:"totalesSubTotales"`
	AutorizadosDescarga  []AutorizadoDescarga   `json:"autorizadosDescarga,omitempty"`
	PedidoGlobal         *PedidoComercialGlobal `json:"pedidoGlobal,omitempty"`
	InfoLogistica        *InfoLogistica         `json:"infoLogistica,omitempty"`
	InfoEntrega          *InfoEntrega           `json:"infoEntrega,omitempty"`
	UsoPosterior         *UsoPosterior          `json:"usoPosterior,omitempty"`
}

// DatosTransaccion holds all transaction-level fields for a fiscal document.
type DatosTransaccion struct {
	TipoEmision                  TipoEmision             `json:"tipoEmision"`
	FechaInicioContingencia      string                  `json:"fechaInicioContingencia,omitempty"` // required when TipoEmision = 02/04; format YYYY-MM-DDThh:mm:ss-05:00
	MotivoContingencia           string                  `json:"motivoContingencia,omitempty"`      // required when TipoEmision = 02/04; 15..250 chars
	TipoDocumento                TipoDocumento           `json:"tipoDocumento"`
	NumeroDocumentoFiscal        string                  `json:"numeroDocumentoFiscal"`  // N|10 zero-padded
	PuntoFacturacionFiscal       string                  `json:"puntoFacturacionFiscal"` // N|3, not "000"
	FechaEmision                 string                  `json:"fechaEmision"`           // YYYY-MM-DDThh:mm:ss-05:00
	FechaSalida                  string                  `json:"fechaSalida,omitempty"`  // optional, same format
	NaturalezaOperacion          NaturalezaOperacion     `json:"naturalezaOperacion"`
	TipoOperacion                TipoOperacion           `json:"tipoOperacion"`
	DestinoOperacion             DestinoOperacion        `json:"destinoOperacion"`
	FormatoCAFE                  FormatoCAFE             `json:"formatoCAFE"`
	EntregaCAFE                  EntregaCAFE             `json:"entregaCAFE"`
	EnvioContenedor              EnvioContenedor         `json:"envioContenedor"`
	ProcesoGeneracion            string                  `json:"procesoGeneracion"`            // always "1"
	TipoVenta                    TipoVenta               `json:"tipoVenta,omitempty"`          // omit when TipoOperacion = "2"
	InformacionInteres           string                  `json:"informacionInteres,omitempty"` // optional, up to 5000 chars
	Cliente                      Cliente                 `json:"cliente"`
	ListaDocsFiscalReferenciados []DocFiscalReferenciado `json:"listaDocsFiscalReferenciados,omitempty"` // required for types 04/05; absent for 06/07
	DatosFacturaExportacion      *DatosExportacion       `json:"datosFacturaExportacion,omitempty"`      // required when DestinoOperacion = "2"
	Retencion                    *Retencion              `json:"retencion,omitempty"`
}

// Cliente holds the buyer/recipient information for the fiscal document.
type Cliente struct {
	TipoClienteFE               TipoClienteFE      `json:"tipoClienteFE"`
	TipoContribuyente           TipoContribuyente  `json:"tipoContribuyente,omitempty"`           // omit when TipoClienteFE = "04"
	NumeroRUC                   string             `json:"numeroRUC,omitempty"`                   // omit when TipoClienteFE = "04"
	DigitoVerificadorRUC        string             `json:"digitoVerificadorRUC,omitempty"`        // omit when TipoClienteFE = "04"
	RazonSocial                 string             `json:"razonSocial,omitempty"`                 // required when TipoClienteFE = "01"/"03"
	Direccion                   string             `json:"direccion,omitempty"`                   // required when TipoClienteFE = "01"/"03"
	CodigoUbicacion             string             `json:"codigoUbicacion,omitempty"`             // omit when TipoClienteFE = "04"
	Provincia                   string             `json:"provincia,omitempty"`                   // omit when TipoClienteFE = "04"
	Distrito                    string             `json:"distrito,omitempty"`                    // omit when TipoClienteFE = "04"
	Corregimiento               string             `json:"corregimiento,omitempty"`               // omit when TipoClienteFE = "04"
	TipoIdentificacion          TipoIdentificacion `json:"tipoIdentificacion,omitempty"`          // required when TipoClienteFE = "04"
	NroIdentificacionExtranjero string             `json:"nroIdentificacionExtranjero,omitempty"` // required when TipoClienteFE = "04"
	PaisExtranjero              string             `json:"paisExtranjero,omitempty"`              // only when TipoIdentificacion = "01" (passport)
	Telefono1                   string             `json:"telefono1,omitempty"`
	CorreoElectronico1          string             `json:"correoElectronico1,omitempty"`
	CorreoElectronico2          string             `json:"correoElectronico2,omitempty"`
	CorreoElectronico3          string             `json:"correoElectronico3,omitempty"`
	Pais                        CountryCode        `json:"pais,omitempty"`     // "PA" when DestinoOperacion = "1"; "ZZ" if not in catalog
	PaisOtro                    string             `json:"paisOtro,omitempty"` // required when Pais = "ZZ"
}

// AutorizadoDescarga represents an entity authorized to download the document and its events.
type AutorizadoDescarga struct {
	TipoContribuyente string `json:"tipoContribuyente,omitempty"` // "1": Natural, "2": Jurídico
	RUCReceptor       string `json:"rucReceptor,omitempty"`       // RUC del receptor autorizado
	DigitoVerifRUC    string `json:"digitoVerifRUC,omitempty"`    // Dígito verificador
	NombreAutorizado  string `json:"nombreAutorizado,omitempty"`  // fallback / description
}

// PedidoComercialGlobal contains global purchase order data.
type PedidoComercialGlobal struct {
	NumeroPedido string `json:"numeroPedido,omitempty"`
	FechaPedido  string `json:"fechaPedido,omitempty"`
	NroLinea     string `json:"nroLinea,omitempty"`
	MontoTotal   string `json:"montoTotal,omitempty"`
	MonedaPedido string `json:"monedaPedido,omitempty"`
}

// UsoPosterior is required for documents with TipoEmision = "03" or "04".
type UsoPosterior struct {
	CUFE string `json:"cufe,omitempty"`
}
