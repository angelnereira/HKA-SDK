package soap

import (
	"strings"

	"github.com/angelnereira/hka-sdk/internal"
)

// BuildDescargaXMLEnvelope constructs the SOAP XML for DescargaXML().
func BuildDescargaXMLEnvelope(tokenEmpresa, tokenPassword, sucursal, numDoc, punto, tipoDoc, tipoEmision string) string {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	sb.WriteString(`<soapenv:Envelope xmlns:soapenv="` + nsEnvURI + `" xmlns:tem="` + nsTemURI + `" xmlns:ser="` + nsSerModelURI + `">`)
	sb.WriteString(`<soapenv:Header/>`)
	sb.WriteString(`<soapenv:Body>`)
	sb.WriteString(`<tem:DescargaXML>`)
	internal.Tag(&sb, nsTem, "tokenEmpresa", tokenEmpresa)
	internal.Tag(&sb, nsTem, "tokenPassword", tokenPassword)
	internal.Open(&sb, nsTem, "descargaXMLRequest")
	internal.Tag(&sb, nsSer, "codigoSucursalEmisor", sucursal)
	internal.Tag(&sb, nsSer, "numeroDocumentoFiscal", numDoc)
	internal.Tag(&sb, nsSer, "puntoFacturacionFiscal", punto)
	internal.Tag(&sb, nsSer, "tipoDocumento", tipoDoc)
	internal.Tag(&sb, nsSer, "tipoEmision", tipoEmision)
	internal.Close(&sb, nsTem, "descargaXMLRequest")
	sb.WriteString(`</tem:DescargaXML>`)
	sb.WriteString(`</soapenv:Body>`)
	sb.WriteString(`</soapenv:Envelope>`)
	return sb.String()
}
