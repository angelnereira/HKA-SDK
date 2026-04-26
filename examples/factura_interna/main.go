package main

import (
	"context"
	"errors"
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
		TipoSucursal:         types.SucursalDetal,
		DatosTransaccion: types.DatosTransaccion{
			TipoEmision:            types.EmisionAUPNormal,
			TipoDocumento:          types.TipoDocFacturaInterna,
			NumeroDocumentoFiscal:  hka.PadDocNumber(1),
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
				RazonSocial:          "Mi Cliente S.A.",
				Direccion:            "Ave. La Paz, Edificio 100",
				CodigoUbicacion:      "8-8-14",
				Provincia:            "PANAMA",
				Distrito:             "PANAMA",
				Corregimiento:        "ANCON",
				Pais:                 types.CountryPA,
			},
		},
		ListaItems: []types.Item{
			{
				Descripcion:             "Servicio de consultoria",
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

	ctx := context.Background()
	resp, err := client.Send(ctx, creds, doc)
	if err != nil {
		var valErr *hka.ValidationError
		if errors.As(err, &valErr) {
			for _, fe := range valErr.Fields {
				fmt.Printf("validation error in %s: %s\n", fe.Field, fe.Message)
			}
		}
		log.Fatal(err)
	}

	fmt.Printf("CUFE: %s\n", resp.CUFE)
	fmt.Printf("QR: %s\n", resp.QR)
	fmt.Printf("Protocolo: %s\n", resp.NroProtocoloAutorizacion)
}
