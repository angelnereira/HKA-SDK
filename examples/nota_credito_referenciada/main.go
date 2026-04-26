package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/types"
)

func main() {
	client := hka.NewDemo()

	creds := hka.Credentials{
		TokenEmpresa:  os.Getenv("HKA_TOKEN_EMPRESA"),
		TokenPassword: os.Getenv("HKA_TOKEN_PASSWORD"),
	}

	now := time.Now()

	// The CUFE of the original invoice being referenced.
	// Must be exactly 66 characters and issued within the last 180 days.
	originalCUFE := "REPLACE_WITH_REAL_66_CHAR_CUFE_FROM_ORIGINAL_INVOICE_HERE0000000"
	originalDate := now.Add(-24 * time.Hour) // issued yesterday

	if !hka.IsWithin180Days(originalDate) {
		log.Fatal("referenced document is outside the 180-day window")
	}

	doc := &types.DocumentoElectronico{
		CodigoSucursalEmisor: "0000",
		DatosTransaccion: types.DatosTransaccion{
			TipoEmision:            types.EmisionAUPNormal,
			TipoDocumento:          types.TipoDocNotaCreditoRef,
			NumeroDocumentoFiscal:  hka.PadDocNumber(3),
			PuntoFacturacionFiscal: hka.PadPunto(1),
			FechaEmision:           hka.FormatDateTime(now),
			NaturalezaOperacion:    types.NatDevolucion,
			TipoOperacion:          types.OperacionSalida,
			DestinoOperacion:       types.DestinoPanama,
			FormatoCAFE:            types.CAFEPapelCarta,
			EntregaCAFE:            types.EntregaElectronica,
			EnvioContenedor:        types.ContenedorNormal,
			ProcesoGeneracion:      "1",
			Cliente: types.Cliente{
				TipoClienteFE:        types.ClienteContribuyente,
				TipoContribuyente:    types.ContribuyenteJuridico,
				NumeroRUC:            "155596713-2-2015",
				DigitoVerificadorRUC: "59",
				RazonSocial:          "Mi Cliente S.A.",
				Direccion:            "Ave. La Paz, Edificio 100",
				Pais:                 types.CountryPA,
			},
			ListaDocsFiscalReferenciados: []types.DocFiscalReferenciado{
				{
					FechaEmisionDocFiscalReferenciado: hka.FormatDateTime(originalDate),
					CufeFEReferenciada:                originalCUFE,
				},
			},
		},
		ListaItems: []types.Item{
			{
				Descripcion:             "Devolucion parcial servicio consultoria",
				Cantidad:                "1.000000",
				PrecioUnitario:          "50.000000",
				PrecioUnitarioDescuento: "0.000000",
				PrecioItem:              "50.000000",
				ValorTotal:              "53.500000",
				TasaITBMS:               types.ITBMS7,
				ValorITBMS:              "3.500000",
			},
		},
		TotalesSubTotales: types.TotalesSubTotales{
			TotalPrecioNeto:    "50.00",
			TotalITBMS:         "3.50",
			TotalISC:           "0.00",
			TotalMontoGravado:  "3.50",
			TotalFactura:       "53.50",
			TotalValorRecibido: "53.50",
			TiempoPago:         types.PagoInmediato,
			NroItems:           1,
			TotalTodosItems:    "53.50",
			ListaFormaPago: []types.FormaPagoItem{
				{
					FormaPagoFact:    types.PagoEfectivo,
					ValorCuotaPagada: "53.50",
				},
			},
		},
	}

	ctx := context.Background()
	resp, err := client.Send(ctx, creds, doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("CUFE: %s\n", resp.CUFE)
}
