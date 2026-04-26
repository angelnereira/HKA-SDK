package types

// InfoLogistica holds logistics and shipping information for the document.
type InfoLogistica struct {
	EntregaFisica    string
	ModalidadFlete   string
	ClausulaSeguro   string
	NroContenedor    string
	NroManifiesto    string
	NroOrdenEmbarque string
	PuertoOrigen     string
	PuertoDestino    string
	PaisDestino      CountryCode
	FechaEstimadaEntrega string
}

// InfoEntrega holds delivery address information.
type InfoEntrega struct {
	EntregaA    string
	DireccionEntrega string
	Provincia   string
	Distrito    string
	Corregimiento string
	PaisEntrega CountryCode
}
