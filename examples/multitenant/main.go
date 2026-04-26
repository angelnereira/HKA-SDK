package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/types"
)

// tenant represents a single company using the SDK.
type tenant struct {
	name  string
	creds hka.Credentials
}

func main() {
	// A single client is shared across all tenants — it is goroutine-safe.
	client := hka.NewDemo()

	tenants := []tenant{
		{
			name: "Empresa A",
			creds: hka.Credentials{
				TokenEmpresa:  "TOKEN_EMPRESA_A",
				TokenPassword: "TOKEN_PASSWORD_A",
			},
		},
		{
			name: "Empresa B",
			creds: hka.Credentials{
				TokenEmpresa:  "TOKEN_EMPRESA_B",
				TokenPassword: "TOKEN_PASSWORD_B",
			},
		},
	}

	var wg sync.WaitGroup
	for i, t := range tenants {
		wg.Add(1)
		go func(idx int, ten tenant) {
			defer wg.Done()
			resp, err := sendInvoice(client, ten.creds, int64(idx+1))
			if err != nil {
				log.Printf("[%s] error: %v", ten.name, err)
				return
			}
			fmt.Printf("[%s] CUFE: %s\n", ten.name, resp.CUFE)
		}(i, t)
	}
	wg.Wait()
}

func sendInvoice(client *hka.Client, creds hka.Credentials, num int64) (*types.SendResponse, error) {
	now := time.Now()
	doc := &types.DocumentoElectronico{
		CodigoSucursalEmisor: "0000",
		DatosTransaccion: types.DatosTransaccion{
			TipoEmision:            types.EmisionAUPNormal,
			TipoDocumento:          types.TipoDocFacturaInterna,
			NumeroDocumentoFiscal:  hka.PadDocNumber(num),
			PuntoFacturacionFiscal: hka.PadPunto(1),
			FechaEmision:           hka.FormatDateTime(now),
			NaturalezaOperacion:    types.NatVenta,
			TipoOperacion:          types.OperacionSalida,
			DestinoOperacion:       types.DestinoPanama,
			FormatoCAFE:            types.CAFEPapelCarta,
			EntregaCAFE:            types.EntregaElectronica,
			EnvioContenedor:        types.ContenedorNormal,
			ProcesoGeneracion:      "1",
			TipoVenta:              types.VentaServicio,
			Cliente: types.Cliente{
				TipoClienteFE:        types.ClienteContribuyente,
				TipoContribuyente:    types.ContribuyenteJuridico,
				NumeroRUC:            "155596713-2-2015",
				DigitoVerificadorRUC: "59",
				RazonSocial:          "Cliente Demo S.A.",
				Direccion:            "Ave. Central, Ciudad de Panama",
				Pais:                 types.CountryPA,
			},
		},
		ListaItems: []types.Item{
			{
				Descripcion:             "Servicio mensual",
				Cantidad:                "1.000000",
				PrecioUnitario:          "100.000000",
				PrecioUnitarioDescuento: "0.000000",
				PrecioItem:              "100.000000",
				ValorTotal:              "107.000000",
				TasaITBMS:               types.ITBMS7,
				ValorITBMS:              "7.000000",
			},
		},
		TotalesSubTotales: types.TotalesSubTotales{
			TotalPrecioNeto:    "100.00",
			TotalITBMS:         "7.00",
			TotalISC:           "0.00",
			TotalMontoGravado:  "7.00",
			TotalFactura:       "107.00",
			TotalValorRecibido: "107.00",
			TiempoPago:         types.PagoInmediato,
			NroItems:           1,
			TotalTodosItems:    "107.00",
			ListaFormaPago: []types.FormaPagoItem{
				{
					FormaPagoFact:    types.PagoEfectivo,
					ValorCuotaPagada: "107.00",
				},
			},
		},
	}
	return client.Send(context.Background(), creds, doc)
}
