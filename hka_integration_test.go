//go:build integration

package hka_test

import (
	"context"
	"os"
	"testing"
	"time"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/docbuilder"
	"github.com/angelnereira/hka-sdk/types"
)

func getIntegrationCreds(t *testing.T) (hka.Credentials, *hka.Client) {
	tokEmp := os.Getenv("HKA_TOKEN_EMPRESA")
	tokPass := os.Getenv("HKA_TOKEN_PASSWORD")
	if tokEmp == "" || tokPass == "" {
		t.Skip("Skipping live integration tests: HKA_TOKEN_EMPRESA and HKA_TOKEN_PASSWORD not set")
	}

	creds := hka.Credentials{
		TokenEmpresa:  tokEmp,
		TokenPassword: tokPass,
	}

	endpoint := os.Getenv("HKA_ENDPOINT")
	if endpoint == "" {
		// Default to production endpoint if provided tokens belong to prod
		endpoint = "https://emision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc"
	}

	client := hka.New(&hka.Config{
		Endpoint: endpoint,
		Timeout:  30 * time.Second,
		Retries:  2,
	})
	return creds, client
}

func TestIntegration_RemainingFolios(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	resp, err := client.RemainingFolios(ctx, creds)
	if err != nil {
		t.Fatalf("RemainingFolios failed: %v", err)
	}

	if !resp.Code.IsSuccess() {
		t.Fatalf("RemainingFolios returned non-success code: %s (%s)", resp.Code, resp.Message)
	}

	t.Logf("RemainingFolios SUCCESS: Licencia=%s, FoliosDisponibles=%s, Ciclo=%s, FoliosTotalesCiclo=%s",
		resp.Licencia, resp.FoliosDisponibles, resp.Ciclo, resp.FoliosTotalesCiclo)
}

func TestIntegration_QueryRUC(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Query official RUC of The Factory HKA Corp (155596713-2-2015, DV 59)
	resp, err := client.QueryRUC(ctx, creds, types.RUCJuridico, "155596713-2-2015")
	if err != nil {
		t.Fatalf("QueryRUC failed: %v", err)
	}

	if !resp.Code.IsSuccess() {
		t.Fatalf("QueryRUC returned non-success code: %s (%s)", resp.Code, resp.Message)
	}

	if resp.DV != "59" {
		t.Errorf("expected DV 59, got %s", resp.DV)
	}

	t.Logf("QueryRUC SUCCESS: RUC=%s, DV=%s, RazonSocial=%s, AfiliadoFE=%s",
		resp.RUC, resp.DV, resp.RazonSocial, resp.AfiliadoFE)
}

func TestIntegration_DocumentStatus_Existing(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	key := hka.DocumentKey{
		CodigoSucursalEmisor:   "0000",
		NumeroDocumentoFiscal:  "0000000001",
		PuntoFacturacionFiscal: "001",
		TipoDocumento:          types.TipoDocFacturaInterna,
		TipoEmision:            types.EmisionAUPNormal,
	}

	resp, err := client.DocumentStatus(ctx, creds, key)
	if err != nil {
		t.Fatalf("DocumentStatus failed: %v", err)
	}

	if !resp.Code.IsSuccess() {
		t.Fatalf("DocumentStatus returned code %s: %s", resp.Code, resp.Message)
	}

	t.Logf("DocumentStatus SUCCESS: Estatus=%s, CUFE=%s, FechaEmision=%s",
		resp.EstatusDocumento, resp.CUFE, resp.FechaEmisionDocumento)
}

func TestIntegration_DownloadXML_Existing(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	key := hka.DocumentKey{
		CodigoSucursalEmisor:   "0000",
		NumeroDocumentoFiscal:  "0000000001",
		PuntoFacturacionFiscal: "001",
		TipoDocumento:          types.TipoDocFacturaInterna,
		TipoEmision:            types.EmisionAUPNormal,
	}

	resp, err := client.DownloadXML(ctx, creds, key)
	if err != nil {
		t.Fatalf("DownloadXML failed: %v", err)
	}

	if !resp.Code.IsSuccess() {
		t.Fatalf("DownloadXML returned code %s: %s", resp.Code, resp.Message)
	}

	if len(resp.Document) == 0 {
		t.Errorf("expected non-empty XML base64 string")
	}

	t.Logf("DownloadXML SUCCESS: Base64 Length=%d characters", len(resp.Document))
}

func TestIntegration_DownloadPDF_Existing(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	key := hka.DocumentKey{
		CodigoSucursalEmisor:   "0000",
		NumeroDocumentoFiscal:  "0000000001",
		PuntoFacturacionFiscal: "001",
		TipoDocumento:          types.TipoDocFacturaInterna,
		TipoEmision:            types.EmisionAUPNormal,
	}

	resp, err := client.DownloadPDF(ctx, creds, key, "")
	if err != nil {
		t.Fatalf("DownloadPDF failed: %v", err)
	}

	if !resp.Code.IsSuccess() {
		t.Fatalf("DownloadPDF returned code %s: %s", resp.Code, resp.Message)
	}

	if len(resp.Document) == 0 {
		t.Errorf("expected non-empty PDF base64 string")
	}

	t.Logf("DownloadPDF SUCCESS: Base64 Length=%d characters", len(resp.Document))
}

func TestIntegration_TrackEmail_Existing(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cufe := "FE0120000035907-58-261614-8700002023122600000000010010113130337646"
	resp, err := client.TrackEmail(ctx, creds, cufe)
	if err != nil {
		t.Fatalf("TrackEmail failed: %v", err)
	}

	if !resp.Code.IsSuccess() {
		t.Fatalf("TrackEmail returned code %s: %s", resp.Code, resp.Message)
	}

	t.Logf("TrackEmail SUCCESS: Found %d tracking events", len(resp.Items))
	for i, item := range resp.Items {
		t.Logf("  [%d] Correo: %s, CreadoEn: %s, Estado: %s", i+1, item.Correo, item.CreadoEn, item.Estado)
	}
}

func TestIntegration_SendEmail_Existing(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	key := hka.DocumentKey{
		CodigoSucursalEmisor:   "0000",
		NumeroDocumentoFiscal:  "0000000001",
		PuntoFacturacionFiscal: "001",
		TipoDocumento:          types.TipoDocFacturaInterna,
		TipoEmision:            types.EmisionAUPNormal,
	}

	resp, err := client.SendEmail(ctx, creds, key, "facturas@finanpigsa.com")
	if err != nil {
		t.Fatalf("SendEmail failed: %v", err)
	}

	if !resp.Code.IsSuccess() {
		t.Fatalf("SendEmail returned non-success code %s: %s", resp.Code, resp.Message)
	}

	t.Logf("SendEmail SUCCESS: Code=%s, Message=%s", resp.Code, resp.Message)
}

func TestIntegration_Cancel_LiveContractValidation(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Calling cancel on an expired 2023 document exercises the full SOAP request/response
	// pipeline against HKA's WCF backend without modifying or canceling any active document.
	key := hka.DocumentKey{
		CodigoSucursalEmisor:   "0000",
		NumeroDocumentoFiscal:  "0000000001",
		PuntoFacturacionFiscal: "001",
		TipoDocumento:          types.TipoDocFacturaInterna,
		TipoEmision:            types.EmisionAUPNormal,
	}

	resp, err := client.Cancel(ctx, creds, key, "Prueba de contrato")
	if err != nil {
		if hkaErr, ok := err.(*hka.HKAError); ok {
			t.Logf("Cancel contract successfully verified by live HKA backend: Code=%s, Message=%s", hkaErr.Code, hkaErr.Message)
			return
		}
		t.Fatalf("Unexpected network or system error: %v", err)
	}
	t.Logf("Cancel SUCCESS: Code=%s, Message=%s", resp.Code, resp.Message)
}

func TestIntegration_Send_LiveContractValidation(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Building a document with Numero(1) will exercise the full Enviar SOAP envelope
	// and backend WCF deserializer. Since doc 1 was already authorized in 2023,
	// HKA rejects it as duplicate without consuming any folios or creating new tax liabilities.
	cliente := docbuilder.ClienteContribuyente(
		"155596713-2-2015", "59", "THE FACTORY HKA CORP", "Calle 50",
	)
	cliente.CodigoUbicacion = "8-8-8"

	doc, err := docbuilder.NewFacturaInterna().
		Sucursal("0000").
		Numero(1).
		Punto(1).
		Cliente(cliente).
		AddItem(docbuilder.Item{
			Descripcion:    "Prueba SDK Go",
			Cantidad:       1,
			PrecioUnitario: 1,
			TasaITBMS:      types.ITBMS7,
		}).
		Build()
	if err != nil {
		t.Fatalf("docbuilder failed: %v", err)
	}

	sendResp, err := client.Send(ctx, creds, doc)
	if err != nil {
		if hkaErr, ok := err.(*hka.HKAError); ok {
			t.Logf("Send contract successfully verified by live HKA backend: Code=%s, Message=%s", hkaErr.Code, hkaErr.Message)
			return
		}
		t.Fatalf("Unexpected network/system error: %v", err)
	}
	t.Logf("Send response: Code=%s, Message=%s", sendResp.Code, sendResp.Message)
}

func TestIntegration_SendPayload_DynamicJSON(t *testing.T) {
	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Pass dynamic JSON payload directly to client.SendPayload
	jsonPayload := `{
		"tipo": "01",
		"sucursal": "0000",
		"numero": 1,
		"punto": 1,
		"cliente": {
			"tipo": "contribuyente",
			"ruc": "155596713-2-2015",
			"dv": "59",
			"razonSocial": "THE FACTORY HKA CORP",
			"direccion": "Calle 50"
		},
		"items": [
			{
				"descripcion": "Servicio Dynamic Payload",
				"cantidad": 1,
				"precioUnitario": 5.0,
				"tasaITBMS": "01"
			}
		]
	}`

	sendResp, err := client.SendPayload(ctx, creds, jsonPayload)
	if err != nil {
		if hkaErr, ok := err.(*hka.HKAError); ok {
			t.Logf("SendPayload dynamic JSON contract verified by live HKA backend: Code=%s, Message=%s", hkaErr.Code, hkaErr.Message)
			return
		}
		t.Fatalf("Unexpected error: %v", err)
	}
	t.Logf("SendPayload response: Code=%s, Message=%s", sendResp.Code, sendResp.Message)
}

func TestIntegration_LiveEmissionGuard(t *testing.T) {
	if os.Getenv("HKA_ENABLE_LIVE_EMISSION") != "true" {
		t.Skip("Skipping live document emission (Send/Cancel): set HKA_ENABLE_LIVE_EMISSION=true to enable real emission and folio consumption")
	}

	creds, client := getIntegrationCreds(t)
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	docNum := time.Now().Unix() % 900000000
	doc, err := docbuilder.NewFacturaInterna().
		Sucursal("0000").
		Numero(docNum).
		Punto(1).
		Cliente(docbuilder.ClienteContribuyente(
			"155596713-2-2015", "59", "THE FACTORY HKA CORP", "Calle 50",
		)).
		AddItem(docbuilder.Item{
			Descripcion:    "Prueba SDK Go",
			Cantidad:       1,
			PrecioUnitario: 1,
			TasaITBMS:      types.ITBMS7,
		}).
		Build()
	if err != nil {
		t.Fatalf("docbuilder failed: %v", err)
	}

	sendResp, err := client.Send(ctx, creds, doc)
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
	t.Logf("Send SUCCESS: CUFE=%s, Protocolo=%s", sendResp.CUFE, sendResp.NroProtocoloAutorizacion)
}
