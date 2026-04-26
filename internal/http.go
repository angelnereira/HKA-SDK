package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient is the internal SOAP transport. Not exposed to SDK consumers.
type HTTPClient struct {
	client   *http.Client
	endpoint string
}

// NewHTTPClient constructs an HTTPClient with the given endpoint and timeout.
func NewHTTPClient(endpoint string, timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		endpoint: endpoint,
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxIdleConns:        10,
				IdleConnTimeout:     90 * time.Second,
				TLSHandshakeTimeout: 10 * time.Second,
			},
		},
	}
}

// Post sends a SOAP request and returns the raw response body.
func (h *HTTPClient) Post(ctx context.Context, soapAction, body string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.endpoint, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", fmt.Sprintf("%q", soapAction))

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
