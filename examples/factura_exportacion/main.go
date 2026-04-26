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

	doc := &types.DocumentoElectronico{
		CodigoSucursalEmisor: "0000",
		DatosTransaccion: types.DatosTransaccion{
			TipoEmision:            types.EmisionAUPNormal,
			TipoDocumento:          types.TipoDocFacturaExportacion,
			NumeroDocumentoFiscal:  hka.PadDocNumber(2),
			PuntoFacturacionFiscal: hka.PadPunto(1),
			FechaEmision:           hka.FormatDateTime(now),
			NaturalezaOperacion:    types.NatExportacion,
			TipoOperacion:          types.OperacionSalida,
			DestinoOperacion:       types.DestinoExtranjero,
			FormatoCAFE:            types.CAFEPapelCarta,
			EntregaCAFE:            types.EntregaElectronica,
			EnvioContenedor:        types.ContenedorNormal,
			ProcesoGeneracion:      "1",
			Cliente: types.Cliente{
				TipoClienteFE:               types.ClienteExtranjero,
				TipoIdentificacion:          types.IdPasaporte,
				NroIdentificacionExtranjero: "P12345678",
				RazonSocial:                 "Foreign Corp LLC",
				Pais:                        types.CountryUS,
			},
			DatosFacturaExportacion: &types.DatosExportacion{
				CondicionesEntrega:    types.IncoFOB,
				MonedaOperExportacion: types.CurrencyUSD,
			},
		},
		ListaItems: []types.Item{
			{
				Descripcion:             "Producto de exportacion",
				Cantidad:                "10.000000",
				PrecioUnitario:          "50.000000",
				PrecioUnitarioDescuento: "0.000000",
				PrecioItem:              "500.000000",
				ValorTotal:              "500.000000",
				TasaITBMS:               types.ITBMSExento,
				ValorITBMS:              "0.000000",
			},
		},
		TotalesSubTotales: types.TotalesSubTotales{
			TotalPrecioNeto:    "500.00",
			TotalITBMS:         "0.00",
			TotalISC:           "0.00",
			TotalMontoGravado:  "0.00",
			TotalFactura:       "500.00",
			TotalValorRecibido: "500.00",
			TiempoPago:         types.PagoInmediato,
			NroItems:           1,
			TotalTodosItems:    "500.00",
			ListaFormaPago: []types.FormaPagoItem{
				{
					FormaPagoFact:    types.PagoTransferencia,
					ValorCuotaPagada: "500.00",
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
	fmt.Printf("QR: %s\n", resp.QR)
}
