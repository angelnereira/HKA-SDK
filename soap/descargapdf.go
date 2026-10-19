package soap

import (
	"strings"

	"github.com/angelnereira/hka-sdk/internal"
)

// BuildDescargaPDFEnvelope constructs the SOAP XML for DescargaPDF().
func BuildDescargaPDFEnvelope(tokenEmpresa, tokenPassword, sucursal, numDoc, punto, tipoDoc, tipoEmision, deviceSerial string) string {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	sb.WriteString(`<soapenv:Envelope xmlns:soapenv="` + nsEnvURI + `" xmlns:tem="` + nsTemURI + `" xmlns:ser="` + nsSerModelURI + `">`)
	sb.WriteString(`<soapenv:Header/>`)
	sb.WriteString(`<soapenv:Body>`)
	sb.WriteString(`<tem:DescargaPDF>`)
	internal.Tag(&sb, nsTem, "tokenEmpresa", tokenEmpresa)
	internal.Tag(&sb, nsTem, "tokenPassword", tokenPassword)
	internal.Open(&sb, nsTem, "datosDocumento")
	internal.Tag(&sb, nsSer, "codigoSucursalEmisor", sucursal)
	internal.Tag(&sb, nsSer, "numeroDocumentoFiscal", numDoc)
	internal.Tag(&sb, nsSer, "puntoFacturacionFiscal", punto)
	if deviceSerial != "" {
		internal.Tag(&sb, nsSer, "serialDispositivo", deviceSerial)
	}
	internal.Tag(&sb, nsSer, "tipoDocumento", tipoDoc)
	internal.Tag(&sb, nsSer, "tipoEmision", tipoEmision)
	internal.Close(&sb, nsTem, "datosDocumento")
	sb.WriteString(`</tem:DescargaPDF>`)
	sb.WriteString(`</soapenv:Body>`)
	sb.WriteString(`</soapenv:Envelope>`)
	return sb.String()
}
