package docbuilder

import "github.com/angelnereira/hka-sdk/types"

// Cliente constructors build a well-formed buyer for each fiscal client category,
// filling the fields that the corresponding validation rules require and leaving
// the ones that must be absent unset.

// ClienteContribuyente builds a registered taxpayer client (TipoClienteFE 01).
// The taxpayer type defaults to juridical; use ClienteContribuyenteNatural for a
// natural person.
func ClienteContribuyente(ruc, dv, razonSocial, direccion string) types.Cliente {
	return types.Cliente{
		TipoClienteFE:        types.ClienteContribuyente,
		TipoContribuyente:    types.ContribuyenteJuridico,
		NumeroRUC:            ruc,
		DigitoVerificadorRUC: dv,
		RazonSocial:          razonSocial,
		Direccion:            direccion,
		Pais:                 types.CountryPA,
	}
}

// ClienteContribuyenteNatural builds a registered natural-person taxpayer (01).
func ClienteContribuyenteNatural(ruc, dv, razonSocial, direccion string) types.Cliente {
	c := ClienteContribuyente(ruc, dv, razonSocial, direccion)
	c.TipoContribuyente = types.ContribuyenteNatural
	return c
}

// ClienteConsumidorFinal builds a final-consumer client (TipoClienteFE 02).
// Final consumers carry no RUC; only a name and address are needed.
func ClienteConsumidorFinal(razonSocial, direccion string) types.Cliente {
	return types.Cliente{
		TipoClienteFE:     types.ClienteConsumidorFinal,
		TipoContribuyente: types.ContribuyenteNatural,
		RazonSocial:       razonSocial,
		Direccion:         direccion,
		Pais:              types.CountryPA,
	}
}

// ClienteGobierno builds a public-administration client (TipoClienteFE 03).
// Items sold to government clients must carry CPBS classification fields.
func ClienteGobierno(ruc, dv, razonSocial, direccion string) types.Cliente {
	return types.Cliente{
		TipoClienteFE:        types.ClienteGobierno,
		TipoContribuyente:    types.ContribuyenteJuridico,
		NumeroRUC:            ruc,
		DigitoVerificadorRUC: dv,
		RazonSocial:          razonSocial,
		Direccion:            direccion,
		Pais:                 types.CountryPA,
	}
}

// ClienteExtranjero builds a foreign client (TipoClienteFE 04), used for export
// invoices. Pass the foreign identification type and number; pais defaults to the
// "ZZ" sentinel when you also set PaisOtro, otherwise set Pais explicitly.
func ClienteExtranjero(nombre, direccion string, tipoID types.TipoIdentificacion, nroID string, pais types.CountryCode) types.Cliente {
	return types.Cliente{
		TipoClienteFE:               types.ClienteExtranjero,
		RazonSocial:                 nombre,
		Direccion:                   direccion,
		TipoIdentificacion:          tipoID,
		NroIdentificacionExtranjero: nroID,
		Pais:                        pais,
	}
}
