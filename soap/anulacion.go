package soap

import (
	"strings"

	"github.com/angelnereira/hka-sdk/internal"
)

// BuildAnulacionEnvelope constructs the SOAP XML for AnulacionDocumento().
func BuildAnulacionEnvelope(tokenEmpresa, tokenPassword, sucursal, numDoc, punto, tipoDoc, tipoEmision, motivo string) string {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	sb.WriteString(`<soapenv:Envelope xmlns:soapenv="` + nsEnvURI + `" xmlns:tem="` + nsTemURI + `" xmlns:ser="` + nsSerModelURI + `">`)
	sb.WriteString(`<soapenv:Header/>`)
	sb.WriteString(`<soapenv:Body>`)
	sb.WriteString(`<tem:AnulacionDocumento>`)
	internal.Tag(&sb, nsTem, "tokenEmpresa", tokenEmpresa)
	internal.Tag(&sb, nsTem, "tokenPassword", tokenPassword)
	internal.Open(&sb, nsTem, "anulacionDocumentoRequest")
	internal.Tag(&sb, nsSer, "codigoSucursalEmisor", sucursal)
	internal.Tag(&sb, nsSer, "numeroDocumentoFiscal", numDoc)
	internal.Tag(&sb, nsSer, "puntoFacturacionFiscal", punto)
	internal.Tag(&sb, nsSer, "tipoDocumento", tipoDoc)
	internal.Tag(&sb, nsSer, "tipoEmision", tipoEmision)
	internal.Tag(&sb, nsSer, "motivoAnulacion", motivo)
	internal.Close(&sb, nsTem, "anulacionDocumentoRequest")
	sb.WriteString(`</tem:AnulacionDocumento>`)
	sb.WriteString(`</soapenv:Body>`)
	sb.WriteString(`</soapenv:Envelope>`)
	return sb.String()
}
