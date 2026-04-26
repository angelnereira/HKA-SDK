package soap

import (
	"strings"

	"github.com/angelnereira/hka-sdk/internal"
)

const nsSerApiRestURI = "http://schemas.datacontract.org/2004/07/Services.ApiRest"

// BuildConsultarRucDVEnvelope constructs the SOAP XML for ConsultarRucDV().
// Tokens go inside the consultarRucDVRequest sub-object with the ApiRest namespace.
func BuildConsultarRucDVEnvelope(tokenEmpresa, tokenPassword, tipoRuc, ruc string) string {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	sb.WriteString(`<soapenv:Envelope xmlns:soapenv="` + nsEnvURI + `" xmlns:tem="` + nsTemURI + `" xmlns:ser="` + nsSerApiRestURI + `">`)
	sb.WriteString(`<soapenv:Header/>`)
	sb.WriteString(`<soapenv:Body>`)
	sb.WriteString(`<tem:ConsultarRucDV>`)
	internal.Open(&sb, nsTem, "consultarRucDVRequest")
	internal.Tag(&sb, nsSer, "tokenEmpresa", tokenEmpresa)
	internal.Tag(&sb, nsSer, "tokenPassword", tokenPassword)
	internal.Tag(&sb, nsSer, "tipoRuc", tipoRuc)
	internal.Tag(&sb, nsSer, "ruc", ruc)
	internal.Close(&sb, nsTem, "consultarRucDVRequest")
	sb.WriteString(`</tem:ConsultarRucDV>`)
	sb.WriteString(`</soapenv:Body>`)
	sb.WriteString(`</soapenv:Envelope>`)
	return sb.String()
}
