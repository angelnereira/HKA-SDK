// Package hka provides a Go SDK for The Factory HKA electronic invoicing API (Panama — DGI).
//
// The client is stateless with respect to credentials: pass a Credentials value on
// every call to support multi-tenant use without any shared mutable state.
//
// Quick start:
//
//	client := hka.NewDemo()
//	resp, err := client.Send(ctx, hka.Credentials{
//	    TokenEmpresa:  os.Getenv("HKA_TOKEN_EMPRESA"),
//	    TokenPassword: os.Getenv("HKA_TOKEN_PASSWORD"),
//	}, doc)
package hka

import (
	"context"
	"fmt"
	"time"

	"github.com/angelnereira/hka-sdk/internal"
	"github.com/angelnereira/hka-sdk/soap"
	"github.com/angelnereira/hka-sdk/types"
	"github.com/angelnereira/hka-sdk/validate"
)

const (
	// DemoEndpoint is the HKA integration/testing environment URL.
	DemoEndpoint = "https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc"

	defaultTimeout = 30 * time.Second
)

// SOAPActions maps each method to its required SOAPAction header value.
var soapActions = map[string]string{
	"Enviar":             "http://tempuri.org/IService/Enviar",
	"EstadoDocumento":    "http://tempuri.org/IService/EstadoDocumento",
	"AnulacionDocumento": "http://tempuri.org/IService/AnulacionDocumento",
	"DescargaXML":        "http://tempuri.org/IService/DescargaXML",
	"DescargaPDF":        "http://tempuri.org/IService/DescargaPDF",
	"FoliosRestantes":    "http://tempuri.org/IService/FoliosRestantes",
	"EnvioCorreo":        "http://tempuri.org/IService/EnvioCorreo",
	"RastreoCorreo":      "http://tempuri.org/IService/RastreoCorreo",
	"ConsultarRucDV":     "http://tempuri.org/IService/ConsultarRucDV",
}

// Config holds optional configuration for Client construction.
type Config struct {
	// Endpoint overrides the default demo URL. Use this to point to production.
	Endpoint string
	// Timeout overrides the default 30-second HTTP timeout.
	Timeout time.Duration
}

// Client is the HKA SDK entry point. It is goroutine-safe and intended to be
// shared across goroutines and tenants. Credentials are per-call.
type Client struct {
	http *internal.HTTPClient
}

// New creates a Client with the provided configuration.
// Pass nil to use defaults (demo endpoint, 30s timeout).
func New(cfg *Config) *Client {
	endpoint := DemoEndpoint
	timeout := defaultTimeout
	if cfg != nil {
		if cfg.Endpoint != "" {
			endpoint = cfg.Endpoint
		}
		if cfg.Timeout > 0 {
			timeout = cfg.Timeout
		}
	}
	return &Client{
		http: internal.NewHTTPClient(endpoint, timeout),
	}
}

// NewDemo creates a Client pointed at the HKA demo/integration environment.
func NewDemo() *Client {
	return New(nil)
}

// Credentials holds per-tenant authentication tokens issued by HKA.
// Credentials are scoped to a single RUC and must not be shared across companies.
type Credentials struct {
	TokenEmpresa  string
	TokenPassword string
}

// GetTokens implements the interface consumed by validate.Credentials.
func (c Credentials) GetTokens() (string, string) {
	return c.TokenEmpresa, c.TokenPassword
}

// DocumentKey uniquely identifies a fiscal document using its five-field composite key.
type DocumentKey struct {
	CodigoSucursalEmisor   string
	NumeroDocumentoFiscal  string
	PuntoFacturacionFiscal string
	TipoDocumento          types.TipoDocumento
	TipoEmision            types.TipoEmision
}

// HKAError is returned when the HKA service responds with a non-success code.
type HKAError struct {
	Code    types.ReturnCode
	Message string
	Raw     string
}

func (e *HKAError) Error() string {
	return fmt.Sprintf("hka error %s: %s", e.Code, e.Message)
}

// NetworkError wraps connection or HTTP-level failures.
type NetworkError struct {
	Cause error
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("hka network error: %v", e.Cause)
}

func (e *NetworkError) Unwrap() error {
	return e.Cause
}

// ValidationError re-exports the validate package type so callers only need
// to import "github.com/angelnereira/hka-sdk" to handle all error types.
type ValidationError = validate.ValidationError

// Send submits a fiscal document to the HKA service.
// Pre-flight validation runs before any HTTP call; a *ValidationError is returned
// if the document is structurally invalid.
func (c *Client) Send(ctx context.Context, creds Credentials, doc *types.DocumentoElectronico) (*types.SendResponse, error) {
	if err := validate.Credentials(creds); err != nil {
		return nil, err
	}
	if err := validate.Document(doc); err != nil {
		return nil, err
	}

	envelope := soap.BuildEnviarEnvelope(creds.TokenEmpresa, creds.TokenPassword, doc)
	raw, err := c.http.Post(ctx, soapActions["Enviar"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}

	resp := soap.ParseSendResponse(raw)
	if !resp.Code.IsSuccess() && !resp.Code.RequiresResend() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}

// DocumentStatus queries the current status of a previously submitted document.
func (c *Client) DocumentStatus(ctx context.Context, creds Credentials, key DocumentKey) (*types.StatusResponse, error) {
	envelope := soap.BuildEstadoEnvelope(
		creds.TokenEmpresa, creds.TokenPassword,
		key.CodigoSucursalEmisor, key.NumeroDocumentoFiscal, key.PuntoFacturacionFiscal,
		string(key.TipoDocumento), string(key.TipoEmision),
	)
	raw, err := c.http.Post(ctx, soapActions["EstadoDocumento"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}
	resp := soap.ParseStatusResponse(raw)
	if !resp.Code.IsSuccess() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}

// Cancel requests cancellation of a fiscal document.
// The HKA service enforces a 7-day window; use IsWithin7Days to check beforehand.
func (c *Client) Cancel(ctx context.Context, creds Credentials, key DocumentKey, reason string) (*types.CancelResponse, error) {
	envelope := soap.BuildAnulacionEnvelope(
		creds.TokenEmpresa, creds.TokenPassword,
		key.CodigoSucursalEmisor, key.NumeroDocumentoFiscal, key.PuntoFacturacionFiscal,
		string(key.TipoDocumento), string(key.TipoEmision), reason,
	)
	raw, err := c.http.Post(ctx, soapActions["AnulacionDocumento"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}
	resp := soap.ParseCancelResponse(raw)
	if !resp.Code.IsSuccess() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}

// DownloadXML retrieves the signed XML file for a document as Base64.
func (c *Client) DownloadXML(ctx context.Context, creds Credentials, key DocumentKey) (*types.DownloadResponse, error) {
	envelope := soap.BuildDescargaXMLEnvelope(
		creds.TokenEmpresa, creds.TokenPassword,
		key.CodigoSucursalEmisor, key.NumeroDocumentoFiscal, key.PuntoFacturacionFiscal,
		string(key.TipoDocumento), string(key.TipoEmision),
	)
	raw, err := c.http.Post(ctx, soapActions["DescargaXML"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}
	resp := soap.ParseDownloadResponse(raw)
	if !resp.Code.IsSuccess() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}

// DownloadPDF retrieves the PDF representation of a document as Base64.
func (c *Client) DownloadPDF(ctx context.Context, creds Credentials, key DocumentKey, deviceSerial string) (*types.DownloadResponse, error) {
	envelope := soap.BuildDescargaPDFEnvelope(
		creds.TokenEmpresa, creds.TokenPassword,
		key.CodigoSucursalEmisor, key.NumeroDocumentoFiscal, key.PuntoFacturacionFiscal,
		string(key.TipoDocumento), string(key.TipoEmision), deviceSerial,
	)
	raw, err := c.http.Post(ctx, soapActions["DescargaPDF"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}
	resp := soap.ParseDownloadResponse(raw)
	if !resp.Code.IsSuccess() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}

// SendEmail requests the HKA service to email the document to the provided address.
func (c *Client) SendEmail(ctx context.Context, creds Credentials, key DocumentKey, email string) (*types.BaseResponse, error) {
	envelope := soap.BuildEnvioCorreoEnvelope(
		creds.TokenEmpresa, creds.TokenPassword,
		key.CodigoSucursalEmisor, key.NumeroDocumentoFiscal, key.PuntoFacturacionFiscal,
		string(key.TipoDocumento), string(key.TipoEmision), email,
	)
	raw, err := c.http.Post(ctx, soapActions["EnvioCorreo"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}
	resp := soap.ParseBaseResponse(raw)
	if !resp.Code.IsSuccess() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}

// TrackEmail retrieves email delivery tracking information for a document identified by CUFE.
func (c *Client) TrackEmail(ctx context.Context, creds Credentials, cufe string) (*types.TrackEmailResponse, error) {
	envelope := soap.BuildRastreoCorreoEnvelope(creds.TokenEmpresa, creds.TokenPassword, cufe)
	raw, err := c.http.Post(ctx, soapActions["RastreoCorreo"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}
	resp := soap.ParseTrackEmailResponse(raw)
	if !resp.Code.IsSuccess() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}

// RemainingFolios queries the number of available folios (document slots) for the tenant.
func (c *Client) RemainingFolios(ctx context.Context, creds Credentials) (*types.FoliosResponse, error) {
	envelope := soap.BuildFoliosEnvelope(creds.TokenEmpresa, creds.TokenPassword)
	raw, err := c.http.Post(ctx, soapActions["FoliosRestantes"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}
	resp := soap.ParseFoliosResponse(raw)
	if !resp.Code.IsSuccess() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}

// QueryRUC looks up a Panamanian RUC number and returns its check digit and registration status.
func (c *Client) QueryRUC(ctx context.Context, creds Credentials, rucType types.RUCType, ruc string) (*types.RUCResponse, error) {
	envelope := soap.BuildConsultarRucDVEnvelope(creds.TokenEmpresa, creds.TokenPassword, string(rucType), ruc)
	raw, err := c.http.Post(ctx, soapActions["ConsultarRucDV"], envelope)
	if err != nil {
		return nil, &NetworkError{Cause: err}
	}
	resp := soap.ParseRUCResponse(raw)
	if !resp.Code.IsSuccess() {
		return nil, &HKAError{Code: resp.Code, Message: resp.Message, Raw: string(raw)}
	}
	return resp, nil
}
