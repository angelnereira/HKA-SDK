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
			TipoDocumento:          types.TipoDocFacturaInterna,
			NumeroDocumentoFiscal:  hka.PadDocNumber(4),
			PuntoFacturacionFiscal: hka.PadPunto(1),
			FechaEmision:           hka.FormatDateTime(now),
			NaturalezaOperacion:    types.NatVenta,
			TipoOperacion:          types.OperacionSalida,
			DestinoOperacion:       types.DestinoPanama,
			FormatoCAFE:            types.CAFEPapelCarta,
			EntregaCAFE:            types.EntregaElectronica,
			EnvioContenedor:        types.ContenedorNormal,
			ProcesoGeneracion:      "1",
			Cliente: types.Cliente{
				TipoClienteFE:        types.ClienteGobierno,
				TipoContribuyente:    types.ContribuyenteJuridico,
				NumeroRUC:            "1-NT-1-35649",
				DigitoVerificadorRUC: "00",
				RazonSocial:          "Ministerio de Salud",
				Direccion:            "Edif. Gorgas, Ancón",
				Pais:                 types.CountryPA,
			},
		},
		ListaItems: []types.Item{
			{
				// Government clients require CodigoCPBSAbrev, CodigoCPBS, and UnidadMedidaCPBS
				Descripcion:             "Equipos de computo",
				CodigoCPBSAbrev:         "IT",
				CodigoCPBS:              "3000",
				UnidadMedidaCPBS:        "UNID",
				Cantidad:                "5.000000",
				PrecioUnitario:          "200.000000",
				PrecioUnitarioDescuento: "0.000000",
				PrecioItem:              "1000.000000",
				ValorTotal:              "1070.000000",
				TasaITBMS:               types.ITBMS7,
				ValorITBMS:              "70.000000",
			},
		},
		TotalesSubTotales: types.TotalesSubTotales{
			TotalPrecioNeto:    "1000.00",
			TotalITBMS:         "70.00",
			TotalISC:           "0.00",
			TotalMontoGravado:  "70.00",
			TotalFactura:       "1070.00",
			TotalValorRecibido: "1070.00",
			TiempoPago:         types.PagoInmediato,
			NroItems:           1,
			TotalTodosItems:    "1070.00",
			ListaFormaPago: []types.FormaPagoItem{
				{
					FormaPagoFact:    types.PagoTransferencia,
					ValorCuotaPagada: "1070.00",
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
