package types

// SendResponse is the typed response from Send().
type SendResponse struct {
	Code                    ReturnCode
	Result                  string
	Message                 string
	CUFE                    string
	QR                      string
	FechaRecepcionDGI       string
	NroProtocoloAutorizacion string
}

// StatusResponse is the typed response from DocumentStatus().
type StatusResponse struct {
	Code                    ReturnCode
	Message                 string
	CUFE                    string
	FechaEmisionDocumento   string
	FechaRecepcionDocumento string
	EstatusDocumento        string
	MensajeDocumento        string
	Result                  string
}

// CancelResponse is the typed response from Cancel().
type CancelResponse struct {
	Code             ReturnCode
	Result           string
	Message          string
	EstatusDocumento string
}

// DownloadResponse carries the file content in Base64 encoding.
type DownloadResponse struct {
	Code     ReturnCode
	Result   string
	Message  string
	Document string // Base64-encoded XML or PDF
}

// FoliosResponse details folio availability for the authenticated tenant.
type FoliosResponse struct {
	Code                  ReturnCode
	Message               string
	Licencia              string
	FechaLicencia         string
	Ciclo                 string
	FechaCiclo            string
	FoliosTotalesCiclo    string
	FoliosUtilizadosCiclo string
	FoliosDisponibles     string
	FoliosTotales         string
	Result                string
}

// RUCResponse is the typed response from QueryRUC().
type RUCResponse struct {
	Code        ReturnCode
	TipoRuc     RUCType
	RUC         string
	DV          string
	RazonSocial string
	AfiliadoFE  string
	Message     string
	Result      string
}

// TrackEmailResponse is the typed response from TrackEmail().
type TrackEmailResponse struct {
	Code    ReturnCode
	Message string
	Items   []EmailTrack
	Result  string
}

// EmailTrack holds tracking information for a single email delivery event.
type EmailTrack struct {
	Correo    string
	CreadoEn  string
	Estado    string
	MessageID string
}

// BaseResponse covers simple responses such as SendEmail().
type BaseResponse struct {
	Code    ReturnCode
	Result  string
	Message string
}
