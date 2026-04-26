package soap

import (
	"strings"

	"github.com/angelnereira/hka-sdk/internal"
)

// BuildRastreoCorreoEnvelope constructs the SOAP XML for RastreoCorreo().
// Identifies the document by CUFE, not by DocumentKey.
// Only uses the tem namespace.
func BuildRastreoCorreoEnvelope(tokenEmpresa, tokenPassword, cufe string) string {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	sb.WriteString(`<soapenv:Envelope xmlns:soapenv="` + nsEnvURI + `" xmlns:tem="` + nsTemURI + `">`)
	sb.WriteString(`<soapenv:Header/>`)
	sb.WriteString(`<soapenv:Body>`)
	sb.WriteString(`<tem:RastreoCorreo>`)
	internal.Tag(&sb, nsTem, "tokenEmpresa", tokenEmpresa)
	internal.Tag(&sb, nsTem, "tokenPassword", tokenPassword)
	internal.Tag(&sb, nsTem, "cufe", cufe)
	sb.WriteString(`</tem:RastreoCorreo>`)
	sb.WriteString(`</soapenv:Body>`)
	sb.WriteString(`</soapenv:Envelope>`)
	return sb.String()
}
