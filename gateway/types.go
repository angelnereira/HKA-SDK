package gateway

import (
	"fmt"
	"time"

	"github.com/angelnereira/hka-sdk/docbuilder"
	"github.com/angelnereira/hka-sdk/types"
)

// This file defines the clean JSON DTOs the gateway accepts and the mapping from
// those DTOs onto the safe-by-construction docbuilder. Any language can produce
// this JSON; the gateway hides the SOAP protocol and the monetary math entirely.

// ClienteRequest is a language-neutral description of the buyer. Tipo selects the
// fiscal client category and which fields are required.
type ClienteRequest struct {
	Tipo               string `json:"tipo"` // contribuyente | contribuyenteNatural | consumidorFinal | gobierno | extranjero
	RUC                string `json:"ruc,omitempty"`
	DV                 string `json:"dv,omitempty"`
	RazonSocial        string `json:"razonSocial,omitempty"`
	Direccion          string `json:"direccion,omitempty"`
	TipoIdentificacion string `json:"tipoIdentificacion,omitempty"` // extranjero: 01 pasaporte, 02 tributario, 99 otro
	NroIdentificacion  string `json:"nroIdentificacion,omitempty"`  // extranjero
	Pais               string `json:"pais,omitempty"`               // extranjero (ISO alpha-2)
	CodigoUbicacion    string `json:"codigoUbicacion,omitempty"`    // provincia-distrito-corregimiento
	Telefono           string `json:"telefono,omitempty"`
	Correo             string `json:"correo,omitempty"`
}

func (c ClienteRequest) toCliente() (types.Cliente, error) {
	var cli types.Cliente
	switch c.Tipo {
	case "contribuyente":
		cli = docbuilder.ClienteContribuyente(c.RUC, c.DV, c.RazonSocial, c.Direccion)
	case "contribuyenteNatural":
		cli = docbuilder.ClienteContribuyenteNatural(c.RUC, c.DV, c.RazonSocial, c.Direccion)
	case "consumidorFinal":
		cli = docbuilder.ClienteConsumidorFinal(c.RazonSocial, c.Direccion)
	case "gobierno":
		cli = docbuilder.ClienteGobierno(c.RUC, c.DV, c.RazonSocial, c.Direccion)
	case "extranjero":
		pais := types.CountryCode(c.Pais)
		if pais == "" {
			pais = types.CountryUS
		}
		cli = docbuilder.ClienteExtranjero(c.RazonSocial, c.Direccion, types.TipoIdentificacion(c.TipoIdentificacion), c.NroIdentificacion, pais)
	default:
		return types.Cliente{}, fmt.Errorf("cliente.tipo %q is not one of contribuyente|contribuyenteNatural|consumidorFinal|gobierno|extranjero", c.Tipo)
	}
	if c.CodigoUbicacion != "" {
		cli.CodigoUbicacion = c.CodigoUbicacion
	}
	if c.Telefono != "" {
		cli.Telefono1 = c.Telefono
	}
	if c.Correo != "" {
		cli.CorreoElectronico1 = c.Correo
	}
	return cli, nil
}

// ItemRequest mirrors docbuilder.Item with a string ITBMS rate (empty to let the
// gateway infer it when autoTasaITBMS is set).
type ItemRequest struct {
	Descripcion      string  `json:"descripcion"`
	Codigo           string  `json:"codigo,omitempty"`
	UnidadMedida     string  `json:"unidadMedida,omitempty"`
	Cantidad         float64 `json:"cantidad"`
	PrecioUnitario   float64 `json:"precioUnitario"`
	Descuento        float64 `json:"descuento,omitempty"`
	TasaITBMS        string  `json:"tasaITBMS,omitempty"` // 00 exento, 01 7%, 02 10%, 03 15%
	CodigoCPBSAbrev  string  `json:"codigoCPBSAbrev,omitempty"`
	CodigoCPBS       string  `json:"codigoCPBS,omitempty"`
	UnidadMedidaCPBS string  `json:"unidadMedidaCPBS,omitempty"`
}

func (i ItemRequest) toItem() docbuilder.Item {
	return docbuilder.Item{
		Descripcion:      i.Descripcion,
		Codigo:           i.Codigo,
		UnidadMedida:     i.UnidadMedida,
		Cantidad:         i.Cantidad,
		PrecioUnitario:   i.PrecioUnitario,
		Descuento:        i.Descuento,
		TasaITBMS:        types.TasaITBMS(i.TasaITBMS),
		CodigoCPBSAbrev:  i.CodigoCPBSAbrev,
		CodigoCPBS:       i.CodigoCPBS,
		UnidadMedidaCPBS: i.UnidadMedidaCPBS,
	}
}

// FormaPagoRequest is one explicit payment entry.
type FormaPagoRequest struct {
	Forma string  `json:"forma"` // 01..09, 99
	Valor float64 `json:"valor"`
}

// PagoPlazoRequest is one deferred-payment installment.
type PagoPlazoRequest struct {
	FechaVence  time.Time `json:"fechaVence"`
	Valor       float64   `json:"valor"`
	Descripcion string    `json:"descripcion,omitempty"`
}

// ExportacionRequest carries the export data block (types 03 / foreign destination).
type ExportacionRequest struct {
	CondicionesEntrega    string `json:"condicionesEntrega"` // Incoterm
	Moneda                string `json:"moneda"`             // ISO 4217
	MonedaNoDef           string `json:"monedaNoDef,omitempty"`
	TipoDeCambio          string `json:"tipoDeCambio,omitempty"`
	MontoMonedaExtranjera string `json:"montoMonedaExtranjera,omitempty"`
	PuertoEmbarque        string `json:"puertoEmbarque,omitempty"`
}

// ReferenciaRequest references a prior FE by CUFE (types 04 / 05).
type ReferenciaRequest struct {
	CUFE         string    `json:"cufe"`
	FechaEmision time.Time `json:"fechaEmision"`
}

// RetencionRequest carries withholding data.
type RetencionRequest struct {
	Codigo string `json:"codigo"`
	Fecha  string `json:"fecha,omitempty"`
	Monto  string `json:"monto,omitempty"`
}

// DocumentRequest is the top-level body for building/sending a document.
type DocumentRequest struct {
	Tipo               string              `json:"tipo"` // 01..10
	Sucursal           string              `json:"sucursal,omitempty"`
	Numero             int64               `json:"numero,omitempty"`
	Punto              int                 `json:"punto,omitempty"`
	FechaEmision       *time.Time          `json:"fechaEmision,omitempty"`
	TipoVenta          string              `json:"tipoVenta,omitempty"`
	InformacionInteres string              `json:"informacionInteres,omitempty"`
	AutoTasaITBMS      bool                `json:"autoTasaITBMS,omitempty"`
	Cliente            ClienteRequest      `json:"cliente"`
	Items              []ItemRequest       `json:"items"`
	FormasPago         []FormaPagoRequest  `json:"formasPago,omitempty"`
	PagosPlazo         []PagoPlazoRequest  `json:"pagosPlazo,omitempty"`
	PagoContado        string              `json:"pagoContado,omitempty"` // forma de pago para el auto-contado
	Exportacion        *ExportacionRequest `json:"exportacion,omitempty"`
	Referencias        []ReferenciaRequest `json:"referencias,omitempty"`
	Retencion          *RetencionRequest   `json:"retencion,omitempty"`
}

// KeyRequest is the five-field composite key identifying a document.
type KeyRequest struct {
	Sucursal      string `json:"sucursal"`
	Numero        string `json:"numero"`
	Punto         string `json:"punto"`
	TipoDocumento string `json:"tipoDocumento"`
	TipoEmision   string `json:"tipoEmision"`
}
