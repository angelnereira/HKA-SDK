// Package docbuilder provides safe-by-construction builders for the ten HKA
// fiscal document types.
//
// The goal is to make it nearly impossible to emit a non-compliant document:
// the builder fills in every mandatory transaction field with the correct
// default for the chosen document type, and computes every monetary field
// (item ITBMS, item totals, and all document totals) automatically so they can
// never disagree with each other. Build() runs the full pre-flight validation
// as a final safety net and returns a ready-to-send *types.DocumentoElectronico.
//
// Typical use:
//
//	doc, err := docbuilder.NewFacturaInterna().
//	    Sucursal("0000").
//	    Numero(1).
//	    Punto(1).
//	    Cliente(docbuilder.ClienteContribuyente("155596713-2-2015", "59", "Mi Cliente S.A.", "Ave. La Paz")).
//	    AddItem(docbuilder.Item{
//	        Descripcion:    "Servicio de consultoria",
//	        Cantidad:       1,
//	        PrecioUnitario: 100,
//	        TasaITBMS:      types.ITBMS7,
//	    }).
//	    Build()
//
// All money math is rounded half-away-from-zero: item-level fields to six
// decimals and document totals to two, matching HKA's published examples.
package docbuilder

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/angelnereira/hka-sdk/catalog"
	"github.com/angelnereira/hka-sdk/types"
	"github.com/angelnereira/hka-sdk/validate"
)

const (
	itemDecimals  = 6
	totalDecimals = 2
)

// Item is the minimal, human-friendly description of a line item. The builder
// derives precioItem, valorITBMS and valorTotal from these fields so the caller
// never has to compute (or mis-compute) them.
type Item struct {
	Descripcion    string          // required, 2..500 chars
	Codigo         string          // optional product code
	UnidadMedida   string          // optional unit of measure
	Cantidad       float64         // required, > 0
	PrecioUnitario float64         // required, >= 0
	Descuento      float64         // per-unit discount in Balboas (optional)
	TasaITBMS      types.TasaITBMS // required ITBMS rate

	// Optional ISC (selective consumption tax). When ValorISC > 0 it is added
	// to the item total and to the document's TotalISC.
	TasaISC  string
	ValorISC float64

	// CPBS classification — required when selling to a government client (03).
	CodigoCPBSAbrev  string
	CodigoCPBS       string
	UnidadMedidaCPBS string

	// Optional extras passed through verbatim.
	FechaFabricacion string
	FechaCaducidad   string
	InfoItem         string
	Medicina         *types.Medicina
	Vehiculo         *types.Vehiculo
}

// Builder accumulates the inputs for one fiscal document. Construct it with one
// of the New* functions, chain the setters, then call Build().
type Builder struct {
	doc        *types.DocumentoElectronico
	items      []Item
	formasPago []types.FormaPagoItem
	pagosPlazo []types.CuotaPlazo
	autoPago   types.FormaPago // when set, a single full-total payment is added at Build
	autoTasa   bool            // when true, empty item ITBMS rates are inferred from the description
	errs       []string
}

// newBuilder seeds a Builder with the defaults common to every document type.
func newBuilder(tipo types.TipoDocumento) *Builder {
	return &Builder{
		doc: &types.DocumentoElectronico{
			CodigoSucursalEmisor: "0000",
			DatosTransaccion: types.DatosTransaccion{
				TipoEmision:            types.EmisionAUPNormal,
				TipoDocumento:          tipo,
				NumeroDocumentoFiscal:  "0000000001",
				PuntoFacturacionFiscal: "001",
				FechaEmision:           formatDateTime(time.Now()),
				NaturalezaOperacion:    types.NatVenta,
				TipoOperacion:          types.OperacionSalida,
				DestinoOperacion:       types.DestinoPanama,
				FormatoCAFE:            types.CAFEPapelCarta,
				EntregaCAFE:            types.EntregaElectronica,
				EnvioContenedor:        types.ContenedorNormal,
				ProcesoGeneracion:      "1",
				TipoVenta:              types.VentaGiroNegocio,
			},
		},
		autoPago: types.PagoEfectivo,
	}
}

// Sucursal sets the four-character emitter branch code (default "0000").
func (b *Builder) Sucursal(codigo string) *Builder {
	b.doc.CodigoSucursalEmisor = padLeft(codigo, 4)
	return b
}

// Numero sets the fiscal document number; it is zero-padded to 10 digits.
func (b *Builder) Numero(n int64) *Builder {
	b.doc.DatosTransaccion.NumeroDocumentoFiscal = fmt.Sprintf("%010d", n)
	return b
}

// Punto sets the billing point; it is zero-padded to 3 digits.
func (b *Builder) Punto(n int) *Builder {
	if n <= 0 {
		b.errs = append(b.errs, "Punto must be greater than 0 (000 is not allowed)")
		return b
	}
	b.doc.DatosTransaccion.PuntoFacturacionFiscal = fmt.Sprintf("%03d", n)
	return b
}

// FechaEmision overrides the emission timestamp (defaults to now, Panama time).
func (b *Builder) FechaEmision(t time.Time) *Builder {
	b.doc.DatosTransaccion.FechaEmision = formatDateTime(t)
	return b
}

// Cliente sets the buyer. Use the Cliente* constructors in this package to
// build a well-formed value for each client category.
func (b *Builder) Cliente(c types.Cliente) *Builder {
	b.doc.DatosTransaccion.Cliente = c
	return b
}

// TipoVenta overrides the sale classification (default VentaGiroNegocio).
func (b *Builder) TipoVenta(t types.TipoVenta) *Builder {
	b.doc.DatosTransaccion.TipoVenta = t
	return b
}

// InformacionInteres attaches free-form information (up to 5000 chars).
func (b *Builder) InformacionInteres(s string) *Builder {
	b.doc.DatosTransaccion.InformacionInteres = s
	return b
}

// Sucursal/branch type, CAFE format and other rarely-changed knobs are exposed
// through the underlying document if needed.
func (b *Builder) Document() *types.DocumentoElectronico { return b.doc }

// AutoTasaITBMS enables inferring the ITBMS rate from each item's description when
// the item leaves TasaITBMS empty. The inference is a best-effort suggestion
// (alcohol/lodging -> 10%, tobacco -> 15%, otherwise 7%); the emisor remains
// responsible for the correct legal rate. Items that set TasaITBMS keep their value.
func (b *Builder) AutoTasaITBMS() *Builder {
	b.autoTasa = true
	return b
}

// AddItem appends a line item. Monetary fields are computed at Build time.
func (b *Builder) AddItem(item Item) *Builder {
	b.items = append(b.items, item)
	return b
}

// AddFormaPago records an explicit payment entry. When at least one is added the
// builder will not auto-generate a cash payment.
func (b *Builder) AddFormaPago(forma types.FormaPago, valor float64) *Builder {
	b.formasPago = append(b.formasPago, types.FormaPagoItem{
		FormaPagoFact:    forma,
		ValorCuotaPagada: fmtAmount(valor, totalDecimals),
	})
	return b
}

// PagoContado selects the payment method auto-applied for the full invoice total
// when no explicit AddFormaPago calls are made (default: cash).
func (b *Builder) PagoContado(forma types.FormaPago) *Builder {
	b.autoPago = forma
	return b
}

// AddPagoPlazo adds one installment and switches the document to deferred
// payment (TiempoPago = plazo). Mixing with AddFormaPago yields TiempoPago = mixto.
func (b *Builder) AddPagoPlazo(fechaVence time.Time, valor float64, descripcion string) *Builder {
	b.pagosPlazo = append(b.pagosPlazo, types.CuotaPlazo{
		FechaPago:        formatDateTime(fechaVence),
		ValorCuotaPlazo:  fmtAmount(valor, totalDecimals),
		DescripcionCuota: descripcion,
	})
	return b
}

// Build computes every monetary field, runs full validation, and returns the
// finished document. If the inputs are inconsistent it returns a *ValidationError
// (or a plain error for builder-level misuse) and a nil document.
func (b *Builder) Build() (*types.DocumentoElectronico, error) {
	if len(b.errs) > 0 {
		return nil, fmt.Errorf("docbuilder: %s", strings.Join(b.errs, "; "))
	}
	if len(b.items) == 0 {
		return nil, fmt.Errorf("docbuilder: at least one item is required")
	}

	var (
		sumPrecioNeto float64
		sumITBMS      float64
		sumISC        float64
		sumTodosItems float64
	)

	built := make([]types.Item, 0, len(b.items))
	for i, in := range b.items {
		if in.Cantidad <= 0 {
			return nil, fmt.Errorf("docbuilder: item[%d] Cantidad must be greater than 0", i)
		}
		if in.TasaITBMS == "" && b.autoTasa {
			in.TasaITBMS = catalog.SugerirTasa(in.Descripcion)
		}
		precioItem := round(in.Cantidad*(in.PrecioUnitario-in.Descuento), itemDecimals)
		valorITBMS := round(precioItem*in.TasaITBMS.Rate(), itemDecimals)
		valorISC := round(in.ValorISC, itemDecimals)
		valorTotal := round(precioItem+valorITBMS+valorISC, itemDecimals)

		sumPrecioNeto += precioItem
		sumITBMS += valorITBMS
		sumISC += valorISC
		sumTodosItems += valorTotal

		it := types.Item{
			Descripcion:             in.Descripcion,
			Codigo:                  in.Codigo,
			UnidadMedida:            in.UnidadMedida,
			Cantidad:                fmtAmount(in.Cantidad, itemDecimals),
			PrecioUnitario:          fmtAmount(in.PrecioUnitario, itemDecimals),
			PrecioUnitarioDescuento: fmtAmount(in.Descuento, itemDecimals),
			PrecioItem:              fmtAmount(precioItem, itemDecimals),
			ValorTotal:              fmtAmount(valorTotal, itemDecimals),
			TasaITBMS:               in.TasaITBMS,
			ValorITBMS:              fmtAmount(valorITBMS, itemDecimals),
			TasaISC:                 in.TasaISC,
			CodigoCPBSAbrev:         in.CodigoCPBSAbrev,
			CodigoCPBS:              in.CodigoCPBS,
			UnidadMedidaCPBS:        in.UnidadMedidaCPBS,
			FechaFabricacion:        in.FechaFabricacion,
			FechaCaducidad:          in.FechaCaducidad,
			InfoItem:                in.InfoItem,
			Medicina:                in.Medicina,
			Vehiculo:                in.Vehiculo,
		}
		if valorISC > 0 {
			it.ValorISC = fmtAmount(valorISC, itemDecimals)
		}
		built = append(built, it)
	}

	totalFactura := round(sumPrecioNeto+sumITBMS+sumISC, totalDecimals)

	totals := types.TotalesSubTotales{
		TotalPrecioNeto:   fmtAmount(round(sumPrecioNeto, totalDecimals), totalDecimals),
		TotalITBMS:        fmtAmount(round(sumITBMS, totalDecimals), totalDecimals),
		TotalMontoGravado: fmtAmount(round(sumITBMS, totalDecimals), totalDecimals),
		TotalFactura:      fmtAmount(totalFactura, totalDecimals),
		NroItems:          len(built),
		TotalTodosItems:   fmtAmount(round(sumTodosItems, totalDecimals), totalDecimals),
	}
	if sumISC > 0 {
		totals.TotalISC = fmtAmount(round(sumISC, totalDecimals), totalDecimals)
	}

	// Payment terms.
	hasPlazo := len(b.pagosPlazo) > 0
	hasContado := len(b.formasPago) > 0
	switch {
	case hasPlazo && hasContado:
		totals.TiempoPago = types.PagoMixto
	case hasPlazo:
		totals.TiempoPago = types.PagoPlazo
	default:
		totals.TiempoPago = types.PagoInmediato
		if !hasContado {
			b.formasPago = []types.FormaPagoItem{{
				FormaPagoFact:    b.autoPago,
				ValorCuotaPagada: fmtAmount(totalFactura, totalDecimals),
			}}
		}
	}
	totals.ListaFormaPago = b.formasPago
	totals.ListaPagoPlazo = b.pagosPlazo

	// TotalValorRecibido is meaningful for immediate cash payment.
	if totals.TiempoPago == types.PagoInmediato {
		totals.TotalValorRecibido = fmtAmount(totalFactura, totalDecimals)
	}

	b.doc.ListaItems = built
	b.doc.TotalesSubTotales = totals

	if ve := validate.Document(b.doc); ve != nil {
		return nil, ve
	}
	return b.doc, nil
}

// --- money + formatting helpers ---

func round(v float64, decimals int) float64 {
	p := math.Pow10(decimals)
	if v < 0 {
		return math.Ceil(v*p-0.5) / p
	}
	return math.Floor(v*p+0.5) / p
}

func fmtAmount(v float64, decimals int) string {
	return strconv.FormatFloat(round(v, decimals), 'f', decimals, 64)
}

func formatDateTime(t time.Time) string {
	loc := time.FixedZone("America/Panama", -5*60*60)
	return t.In(loc).Format("2006-01-02T15:04:05-07:00")
}

func padLeft(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return strings.Repeat("0", width-len(s)) + s
}
