// Example: generating and sending a fiscal document with the safe-by-construction
// docbuilder. Compare this with examples/factura_interna, which builds the same
// document by hand — here every monetary field and total is computed for you.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	hka "github.com/angelnereira/hka-sdk"
	"github.com/angelnereira/hka-sdk/docbuilder"
	"github.com/angelnereira/hka-sdk/types"
)

func main() {
	client := hka.NewDemo()
	creds := hka.Credentials{
		TokenEmpresa:  os.Getenv("HKA_TOKEN_EMPRESA"),
		TokenPassword: os.Getenv("HKA_TOKEN_PASSWORD"),
	}

	// Build a domestic invoice. The builder fills the mandatory transaction
	// fields, computes item ITBMS and totals, and validates before returning.
	doc, err := docbuilder.NewFacturaInterna().
		Sucursal("0000").
		Numero(1).
		Punto(1).
		Cliente(docbuilder.ClienteContribuyente(
			"155596713-2-2015", "59",
			"Mi Cliente S.A.", "Ave. La Paz, Edificio 100",
		)).
		AddItem(docbuilder.Item{
			Descripcion:    "Servicio de consultoria",
			Cantidad:       1,
			PrecioUnitario: 100,
			TasaITBMS:      types.ITBMS7,
		}).
		AddItem(docbuilder.Item{
			Descripcion:    "Licencia de software",
			Cantidad:       2,
			PrecioUnitario: 50,
			Descuento:      5, // per-unit discount
			TasaITBMS:      types.ITBMS7,
		}).
		Build()
	if err != nil {
		var ve *hka.ValidationError
		if errors.As(err, &ve) {
			for _, fe := range ve.Fields {
				fmt.Printf("validation: %s — %s\n", fe.Field, fe.Message)
			}
		}
		log.Fatal(err)
	}

	fmt.Printf("Total factura: %s\n", doc.TotalesSubTotales.TotalFactura)

	resp, err := client.Send(context.Background(), creds, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CUFE: %s\n", resp.CUFE)
}
