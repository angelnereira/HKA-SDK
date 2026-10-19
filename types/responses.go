package types

import "encoding/base64"

// SendResponse is the typed response from Send().
type SendResponse struct {
	Code                     ReturnCode `json:"code"`
	Result                   string     `json:"result,omitempty"`
	Message                  string     `json:"message"`
	CUFE                     string     `json:"cufe,omitempty"`
	QR                       string     `json:"qr,omitempty"`
	FechaRecepcionDGI        string     `json:"fechaRecepcionDGI,omitempty"`
	NroProtocoloAutorizacion string     `json:"nroProtocoloAutorizacion,omitempty"`
}

// StatusResponse is the typed response from DocumentStatus().
type StatusResponse struct {
	Code                    ReturnCode `json:"code"`
	Message                 string     `json:"message"`
	CUFE                    string     `json:"cufe,omitempty"`
	FechaEmisionDocumento   string     `json:"fechaEmisionDocumento,omitempty"`
	FechaRecepcionDocumento string     `json:"fechaRecepcionDocumento,omitempty"`
	EstatusDocumento        string     `json:"estatusDocumento,omitempty"`
	MensajeDocumento        string     `json:"mensajeDocumento,omitempty"`
	Result                  string     `json:"result,omitempty"`
}

// CancelResponse is the typed response from Cancel().
type CancelResponse struct {
	Code             ReturnCode `json:"code"`
	Result           string     `json:"result,omitempty"`
	Message          string     `json:"message"`
	EstatusDocumento string     `json:"estatusDocumento,omitempty"`
}

// DownloadResponse carries the file content in Base64 encoding.
type DownloadResponse struct {
	Code     ReturnCode `json:"code"`
	Result   string     `json:"result,omitempty"`
	Message  string     `json:"message"`
	Document string     `json:"document"` // Base64-encoded XML or PDF
}

// Bytes decodes the Base64 Document into its raw bytes. For DownloadPDF this is the
// CAFE PDF; for DownloadXML it is the signed XML. The result can be written directly
// to a file (e.g. os.WriteFile("cafe.pdf", b, 0o644)).
func (r *DownloadResponse) Bytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(r.Document)
}

// FoliosResponse details folio availability for the authenticated tenant.
type FoliosResponse struct {
	Code                  ReturnCode `json:"code"`
	Message               string     `json:"message"`
	Licencia              string     `json:"licencia,omitempty"`
	FechaLicencia         string     `json:"fechaLicencia,omitempty"`
	Ciclo                 string     `json:"ciclo,omitempty"`
	FechaCiclo            string     `json:"fechaCiclo,omitempty"`
	FoliosTotalesCiclo    string     `json:"foliosTotalesCiclo,omitempty"`
	FoliosUtilizadosCiclo string     `json:"foliosUtilizadosCiclo,omitempty"`
	FoliosDisponibles     string     `json:"foliosDisponibles,omitempty"`
	FoliosTotales         string     `json:"foliosTotales,omitempty"`
	Result                string     `json:"result,omitempty"`
}

// RUCResponse is the typed response from QueryRUC().
type RUCResponse struct {
	Code        ReturnCode `json:"code"`
	TipoRuc     RUCType    `json:"tipoRuc,omitempty"`
	RUC         string     `json:"ruc,omitempty"`
	DV          string     `json:"dv,omitempty"`
	RazonSocial string     `json:"razonSocial,omitempty"`
	AfiliadoFE  string     `json:"afiliadoFE,omitempty"`
	Message     string     `json:"message"`
	Result      string     `json:"result,omitempty"`
}

// TrackEmailResponse is the typed response from TrackEmail().
type TrackEmailResponse struct {
	Code    ReturnCode   `json:"code"`
	Message string       `json:"message"`
	Items   []EmailTrack `json:"items,omitempty"`
	Result  string       `json:"result,omitempty"`
}

// EmailTrack holds tracking information for a single email delivery event.
type EmailTrack struct {
	Correo    string `json:"correo"`
	CreadoEn  string `json:"creadoEn"`
	Estado    string `json:"estado"`
	MessageID string `json:"messageID,omitempty"`
}

// BaseResponse covers simple responses such as SendEmail().
type BaseResponse struct {
	Code    ReturnCode `json:"code"`
	Result  string     `json:"result,omitempty"`
	Message string     `json:"message"`
}
