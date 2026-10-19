package types

// InfoLogistica holds logistics and shipping information for the document.
type InfoLogistica struct {
	EntregaFisica        string      `json:"entregaFisica,omitempty"`
	ModalidadFlete       string      `json:"modalidadFlete,omitempty"`
	ClausulaSeguro       string      `json:"clausulaSeguro,omitempty"`
	NroContenedor        string      `json:"nroContenedor,omitempty"`
	NroManifiesto        string      `json:"nroManifiesto,omitempty"`
	NroOrdenEmbarque     string      `json:"nroOrdenEmbarque,omitempty"`
	PuertoOrigen         string      `json:"puertoOrigen,omitempty"`
	PuertoDestino        string      `json:"puertoDestino,omitempty"`
	PaisDestino          CountryCode `json:"paisDestino,omitempty"`
	FechaEstimadaEntrega string      `json:"fechaEstimadaEntrega,omitempty"`
}

// InfoEntrega holds delivery address information.
type InfoEntrega struct {
	EntregaA         string      `json:"entregaA,omitempty"`
	DireccionEntrega string      `json:"direccionEntrega,omitempty"`
	Provincia        string      `json:"provincia,omitempty"`
	Distrito         string      `json:"distrito,omitempty"`
	Corregimiento    string      `json:"corregimiento,omitempty"`
	PaisEntrega      CountryCode `json:"paisEntrega,omitempty"`
}
