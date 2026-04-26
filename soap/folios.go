package soap

import (
	"strings"

	"github.com/angelnereira/hka-sdk/internal"
)

// BuildFoliosEnvelope constructs the SOAP XML for FoliosRestantes().
// This method only uses the tem namespace — no ser namespace.
func BuildFoliosEnvelope(tokenEmpresa, tokenPassword string) string {
	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	sb.WriteString(`<soapenv:Envelope xmlns:soapenv="` + nsEnvURI + `" xmlns:tem="` + nsTemURI + `">`)
	sb.WriteString(`<soapenv:Header/>`)
	sb.WriteString(`<soapenv:Body>`)
	sb.WriteString(`<tem:FoliosRestantes>`)
	internal.Tag(&sb, nsTem, "tokenEmpresa", tokenEmpresa)
	internal.Tag(&sb, nsTem, "tokenPassword", tokenPassword)
	sb.WriteString(`</tem:FoliosRestantes>`)
	sb.WriteString(`</soapenv:Body>`)
	sb.WriteString(`</soapenv:Envelope>`)
	return sb.String()
}
