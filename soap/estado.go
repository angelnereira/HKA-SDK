package soap

import (
	"strings"

	"github.com/angelnereira/hka-sdk/internal"
)

const nsSerModelURI = "http://schemas.datacontract.org/2004/07/Services.Model"

// BuildEstadoEnvelope constructs the SOAP XML for EstadoDocumento().
func BuildEstadoEnvelope(tokenEmpresa, tokenPassword, sucursal, numDoc, punto, tipoDoc, tipoEmision string) string {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	sb.WriteString(`<soapenv:Envelope xmlns:soapenv="` + nsEnvURI + `" xmlns:tem="` + nsTemURI + `" xmlns:ser="` + nsSerModelURI + `">`)
	sb.WriteString(`<soapenv:Header/>`)
	sb.WriteString(`<soapenv:Body>`)
	sb.WriteString(`<tem:EstadoDocumento>`)
	internal.Tag(&sb, nsTem, "tokenEmpresa", tokenEmpresa)
	internal.Tag(&sb, nsTem, "tokenPassword", tokenPassword)
	internal.Open(&sb, nsTem, "estadoDocumentoRequest")
	internal.Tag(&sb, nsSer, "codigoSucursalEmisor", sucursal)
	internal.Tag(&sb, nsSer, "numeroDocumentoFiscal", numDoc)
	internal.Tag(&sb, nsSer, "puntoFacturacionFiscal", punto)
	internal.Tag(&sb, nsSer, "tipoDocumento", tipoDoc)
	internal.Tag(&sb, nsSer, "tipoEmision", tipoEmision)
	internal.Close(&sb, nsTem, "estadoDocumentoRequest")
	sb.WriteString(`</tem:EstadoDocumento>`)
	sb.WriteString(`</soapenv:Body>`)
	sb.WriteString(`</soapenv:Envelope>`)
	return sb.String()
}
