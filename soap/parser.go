// Package soap handles SOAP envelope serialization and response parsing for each HKA method.
package soap

import (
	"regexp"
	"strings"

	"github.com/angelnereira/hka-sdk/types"
)

// extract returns the inner text of the first occurrence of <tag>...</tag>.
func extract(xml, tag string) string {
	re := regexp.MustCompile(`(?i)<` + regexp.QuoteMeta(tag) + `(?:\s[^>]*)?>([^<]*)<\/` + regexp.QuoteMeta(tag) + `>`)
	m := re.FindStringSubmatch(xml)
	if len(m) > 1 {
		return strings.TrimSpace(unescapeXML(m[1]))
	}
	return ""
}

// extractAll returns every inner text occurrence of <tag>...</tag>.
func extractAll(xml, tag string) []string {
	re := regexp.MustCompile(`(?i)<` + regexp.QuoteMeta(tag) + `(?:\s[^>]*)?>([^<]*)<\/` + regexp.QuoteMeta(tag) + `>`)
	matches := re.FindAllStringSubmatch(xml, -1)
	out := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) > 1 {
			out = append(out, strings.TrimSpace(unescapeXML(m[1])))
		}
	}
	return out
}

// extractBlock returns the inner content of the first occurrence of <tag>...</tag> including sub-elements.
func extractBlock(xml, tag string) string {
	open := "<" + tag
	close := "</" + tag + ">"
	start := strings.Index(strings.ToLower(xml), strings.ToLower(open))
	if start == -1 {
		return ""
	}
	end := strings.Index(xml[start:], close)
	if end == -1 {
		return ""
	}
	block := xml[start : start+end+len(close)]
	inner := strings.Index(block, ">")
	if inner == -1 {
		return ""
	}
	return block[inner+1 : len(block)-len(close)]
}

func unescapeXML(s string) string {
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&quot;", "\"")
	s = strings.ReplaceAll(s, "&apos;", "'")
	return s
}

// ParseSendResponse parses the SOAP response from Enviar().
func ParseSendResponse(raw []byte) *types.SendResponse {
	xml := string(raw)
	return &types.SendResponse{
		Code:                     types.ReturnCode(extract(xml, "codigo")),
		Result:                   extract(xml, "resultado"),
		Message:                  extract(xml, "mensaje"),
		CUFE:                     extract(xml, "cufe"),
		QR:                       extract(xml, "qr"),
		FechaRecepcionDGI:        extract(xml, "fechaRecepcionDGI"),
		NroProtocoloAutorizacion: extract(xml, "nroProtocoloAutorizacion"),
	}
}

// ParseStatusResponse parses the SOAP response from EstadoDocumento().
func ParseStatusResponse(raw []byte) *types.StatusResponse {
	xml := string(raw)
	return &types.StatusResponse{
		Code:                    types.ReturnCode(extract(xml, "codigo")),
		Message:                 extract(xml, "mensaje"),
		CUFE:                    extract(xml, "cufe"),
		FechaEmisionDocumento:   extract(xml, "fechaEmisionDocumento"),
		FechaRecepcionDocumento: extract(xml, "fechaRecepcionDocumento"),
		EstatusDocumento:        extract(xml, "estatusDocumento"),
		MensajeDocumento:        extract(xml, "mensajeDocumento"),
		Result:                  extract(xml, "resultado"),
	}
}

// ParseCancelResponse parses the SOAP response from AnulacionDocumento().
func ParseCancelResponse(raw []byte) *types.CancelResponse {
	xml := string(raw)
	return &types.CancelResponse{
		Code:             types.ReturnCode(extract(xml, "codigo")),
		Result:           extract(xml, "resultado"),
		Message:          extract(xml, "mensaje"),
		EstatusDocumento: extract(xml, "estatusDocumento"),
	}
}

// ParseDownloadResponse parses the SOAP response from DescargaXML() or DescargaPDF().
func ParseDownloadResponse(raw []byte) *types.DownloadResponse {
	xml := string(raw)
	return &types.DownloadResponse{
		Code:     types.ReturnCode(extract(xml, "codigo")),
		Result:   extract(xml, "resultado"),
		Message:  extract(xml, "mensaje"),
		Document: extract(xml, "documento"),
	}
}

// ParseFoliosResponse parses the SOAP response from FoliosRestantes().
func ParseFoliosResponse(raw []byte) *types.FoliosResponse {
	xml := string(raw)
	return &types.FoliosResponse{
		Code:                  types.ReturnCode(extract(xml, "codigo")),
		Message:               extract(xml, "mensaje"),
		Licencia:              extract(xml, "licencia"),
		FechaLicencia:         extract(xml, "fechaLicencia"),
		Ciclo:                 extract(xml, "ciclo"),
		FechaCiclo:            extract(xml, "fechaCiclo"),
		FoliosTotalesCiclo:    extract(xml, "foliosTotalesCiclo"),
		FoliosUtilizadosCiclo: extract(xml, "foliosUtilizadosCiclo"),
		FoliosDisponibles:     extract(xml, "foliosDisponibles"),
		FoliosTotales:         extract(xml, "foliosTotales"),
		Result:                extract(xml, "resultado"),
	}
}

// ParseBaseResponse parses a minimal SOAP response (e.g. EnvioCorreo).
func ParseBaseResponse(raw []byte) *types.BaseResponse {
	xml := string(raw)
	return &types.BaseResponse{
		Code:    types.ReturnCode(extract(xml, "codigo")),
		Result:  extract(xml, "resultado"),
		Message: extract(xml, "mensaje"),
	}
}

// ParseRUCResponse parses the SOAP response from ConsultarRucDV().
func ParseRUCResponse(raw []byte) *types.RUCResponse {
	xml := string(raw)
	return &types.RUCResponse{
		Code:        types.ReturnCode(extract(xml, "codigo")),
		TipoRuc:     types.RUCType(extract(xml, "tipoRuc")),
		RUC:         extract(xml, "ruc"),
		DV:          extract(xml, "dv"),
		RazonSocial: extract(xml, "razonSocial"),
		AfiliadoFE:  extract(xml, "afiliadoFE"),
		Message:     extract(xml, "mensaje"),
		Result:      extract(xml, "resultado"),
	}
}

// ParseTrackEmailResponse parses the SOAP response from RastreoCorreo().
func ParseTrackEmailResponse(raw []byte) *types.TrackEmailResponse {
	xml := string(raw)
	resp := &types.TrackEmailResponse{
		Code:    types.ReturnCode(extract(xml, "codigo")),
		Message: extract(xml, "mensaje"),
		Result:  extract(xml, "resultado"),
	}
	// parse each email track item block
	correos := extractAll(xml, "correo")
	estados := extractAll(xml, "estado")
	creados := extractAll(xml, "creadoEn")
	msgIDs := extractAll(xml, "messageId")
	for i, c := range correos {
		t := types.EmailTrack{Correo: c}
		if i < len(estados) {
			t.Estado = estados[i]
		}
		if i < len(creados) {
			t.CreadoEn = creados[i]
		}
		if i < len(msgIDs) {
			t.MessageID = msgIDs[i]
		}
		resp.Items = append(resp.Items, t)
	}
	return resp
}
