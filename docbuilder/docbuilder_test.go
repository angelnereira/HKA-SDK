package docbuilder_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/angelnereira/hka-sdk/docbuilder"
	"github.com/angelnereira/hka-sdk/types"
	"github.com/angelnereira/hka-sdk/validate"
)

// reproduces the official HKA factura interna example: one item at 5.55, 7% ITBMS.
func TestBuild_MatchesOfficialExample(t *testing.T) {
	doc, err := docbuilder.NewFacturaInterna().
		Sucursal("0000").
		Numero(1).
		Punto(1).
		Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "Mi Cliente S.A.", "Ave. La Paz, Edificio 100")).
		AddItem(docbuilder.Item{
			Descripcion:    "Muebles",
			UnidadMedida:   "um",
			Cantidad:       1,
			PrecioUnitario: 5.55,
			TasaITBMS:      types.ITBMS7,
		}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	item := doc.ListaItems[0]
	if item.PrecioItem != "5.550000" {
		t.Errorf("PrecioItem = %q, want 5.550000", item.PrecioItem)
	}
	if item.ValorITBMS != "0.388500" {
		t.Errorf("ValorITBMS = %q, want 0.388500", item.ValorITBMS)
	}
	if item.ValorTotal != "5.938500" {
		t.Errorf("ValorTotal = %q, want 5.938500", item.ValorTotal)
	}

	tot := doc.TotalesSubTotales
	checks := map[string]string{
		"TotalPrecioNeto":   "5.55",
		"TotalITBMS":        "0.39",
		"TotalMontoGravado": "0.39",
		"TotalFactura":      "5.94",
		"TotalTodosItems":   "5.94",
	}
	got := map[string]string{
		"TotalPrecioNeto":   tot.TotalPrecioNeto,
		"TotalITBMS":        tot.TotalITBMS,
		"TotalMontoGravado": tot.TotalMontoGravado,
		"TotalFactura":      tot.TotalFactura,
		"TotalTodosItems":   tot.TotalTodosItems,
	}
	for k, want := range checks {
		if got[k] != want {
			t.Errorf("%s = %q, want %q", k, got[k], want)
		}
	}
	if tot.NroItems != 1 {
		t.Errorf("NroItems = %d, want 1", tot.NroItems)
	}
	if len(tot.ListaFormaPago) != 1 || tot.ListaFormaPago[0].ValorCuotaPagada != "5.94" {
		t.Errorf("expected auto cash payment of 5.94, got %+v", tot.ListaFormaPago)
	}
}

// multi-item document with a discount and a 10% line.
func TestBuild_MultiItemTotals(t *testing.T) {
	doc, err := docbuilder.NewFacturaInterna().
		Cliente(docbuilder.ClienteConsumidorFinal("Juan Perez", "Calle 50")).
		AddItem(docbuilder.Item{Descripcion: "Producto A", Cantidad: 2, PrecioUnitario: 10, TasaITBMS: types.ITBMS7}).
		AddItem(docbuilder.Item{Descripcion: "Producto B", Cantidad: 1, PrecioUnitario: 50, Descuento: 5, TasaITBMS: types.ITBMS10}).
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// A: precioItem=20, itbms=1.40 ; B: precioItem=45, itbms=4.50
	// neto=65.00 itbms=5.90 factura=70.90
	tot := doc.TotalesSubTotales
	if tot.TotalPrecioNeto != "65.00" {
		t.Errorf("TotalPrecioNeto = %q, want 65.00", tot.TotalPrecioNeto)
	}
	if tot.TotalITBMS != "5.90" {
		t.Errorf("TotalITBMS = %q, want 5.90", tot.TotalITBMS)
	}
	if tot.TotalFactura != "70.90" {
		t.Errorf("TotalFactura = %q, want 70.90", tot.TotalFactura)
	}
}

// every constructor must yield a document that passes validation when given a
// suitable client and item.
func TestAllConstructors_Validate(t *testing.T) {
	item := docbuilder.Item{Descripcion: "Servicio", Cantidad: 1, PrecioUnitario: 100, TasaITBMS: types.ITBMS7}
	gobItem := docbuilder.Item{
		Descripcion: "Servicio", Cantidad: 1, PrecioUnitario: 100, TasaITBMS: types.ITBMS7,
		CodigoCPBSAbrev: "10", CodigoCPBS: "1310", UnidadMedidaCPBS: "cm",
	}
	contrib := docbuilder.ClienteContribuyente("155596713-2-2015", "59", "Cliente S.A.", "Ave. La Paz")
	extranjero := docbuilder.ClienteExtranjero("Foreign Corp", "123 Main St", types.IdTributario, "TAX-99", types.CountryUS)
	expData := types.DatosExportacion{CondicionesEntrega: types.IncoFOB, MonedaOperExportacion: types.CurrencyUSD, PuertoEmbarque: "Balboa"}

	cufe := "FE01" + strings.Repeat("0", 62) // exactly 66 chars

	cases := []struct {
		name string
		b    *docbuilder.Builder
	}{
		{"01 interna", docbuilder.NewFacturaInterna().Cliente(contrib).AddItem(item)},
		{"02 importacion", docbuilder.NewFacturaImportacion().Cliente(contrib).AddItem(item)},
		{"03 exportacion", docbuilder.NewFacturaExportacion().Cliente(extranjero).Exportacion(expData).AddItem(item)},
		{"04 nc ref", docbuilder.NewNotaCreditoReferenciada().Cliente(contrib).Referencia(cufe, time.Now()).AddItem(item)},
		{"05 nd ref", docbuilder.NewNotaDebitoReferenciada().Cliente(contrib).Referencia(cufe, time.Now()).AddItem(item)},
		{"06 nc gen", docbuilder.NewNotaCreditoGenerica().Cliente(contrib).AddItem(item)},
		{"07 nd gen", docbuilder.NewNotaDebitoGenerica().Cliente(contrib).AddItem(item)},
		{"08 zona franca", docbuilder.NewFacturaZonaFranca().Cliente(contrib).AddItem(item)},
		{"09 reembolso", docbuilder.NewReembolso().Cliente(contrib).AddItem(item)},
		{"10 extranjera", docbuilder.NewFacturaExtranjera().Cliente(contrib).AddItem(item)},
		{"03 gobierno", docbuilder.NewFacturaInterna().Cliente(docbuilder.ClienteGobierno("155596713-2-2015", "59", "Entidad", "Ave. Balboa")).AddItem(gobItem)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := tc.b.Build()
			if err != nil {
				t.Fatalf("Build() error: %v", err)
			}
			if ve := validate.Document(doc); ve != nil {
				t.Fatalf("built document failed validation: %v", ve)
			}
		})
	}
}

func TestBuild_NoItems_Errors(t *testing.T) {
	_, err := docbuilder.NewFacturaInterna().
		Cliente(docbuilder.ClienteConsumidorFinal("Juan", "Calle 50")).
		Build()
	if err == nil {
		t.Fatal("expected error when no items are added")
	}
}

func TestBuild_DeferredPayment(t *testing.T) {
	doc, err := docbuilder.NewFacturaInterna().
		Cliente(docbuilder.ClienteConsumidorFinal("Juan", "Calle 50")).
		AddItem(docbuilder.Item{Descripcion: "Servicio", Cantidad: 1, PrecioUnitario: 100, TasaITBMS: types.ITBMS7}).
		AddPagoPlazo(time.Now().Add(30*24*time.Hour), 107, "Cuota 1").
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if doc.TotalesSubTotales.TiempoPago != types.PagoPlazo {
		t.Errorf("TiempoPago = %q, want %q", doc.TotalesSubTotales.TiempoPago, types.PagoPlazo)
	}
	if len(doc.TotalesSubTotales.ListaPagoPlazo) != 1 {
		t.Errorf("expected 1 installment, got %d", len(doc.TotalesSubTotales.ListaPagoPlazo))
	}
}

func TestBuild_InvalidPropagatesValidationError(t *testing.T) {
	// export invoice without export data must surface a *ValidationError
	_, err := docbuilder.NewFacturaExportacion().
		Cliente(docbuilder.ClienteExtranjero("Foreign", "Addr", types.IdTributario, "X", types.CountryUS)).
		AddItem(docbuilder.Item{Descripcion: "Servicio", Cantidad: 1, PrecioUnitario: 100, TasaITBMS: types.ITBMSExento}).
		Build()
	if err == nil {
		t.Fatal("expected validation error for export invoice without DatosExportacion")
	}
	var ve *validate.ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("expected *validate.ValidationError, got %T", err)
	}
}
