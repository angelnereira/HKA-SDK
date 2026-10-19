package docbuilder

import (
	"time"

	"github.com/angelnereira/hka-sdk/types"
)

// The ten document-type constructors. Each preselects TipoDocumento plus the
// transaction defaults that the HKA rules require for that type, so a document
// built from it passes pre-flight validation as long as the caller supplies a
// client and at least one item.

// NewFacturaInterna creates a builder for a standard domestic invoice (type 01).
func NewFacturaInterna() *Builder {
	return newBuilder(types.TipoDocFacturaInterna)
}

// NewFacturaImportacion creates a builder for an import invoice (type 02).
func NewFacturaImportacion() *Builder {
	b := newBuilder(types.TipoDocFacturaImportacion)
	b.doc.DatosTransaccion.NaturalezaOperacion = types.NatImportacion
	return b
}

// NewFacturaExportacion creates a builder for an export invoice (type 03).
// It pins DestinoOperacion to foreign and NaturalezaOperacion to export; you must
// supply a foreign client (ClienteExtranjero) and call Exportacion().
func NewFacturaExportacion() *Builder {
	b := newBuilder(types.TipoDocFacturaExportacion)
	b.doc.DatosTransaccion.NaturalezaOperacion = types.NatExportacion
	b.doc.DatosTransaccion.DestinoOperacion = types.DestinoExtranjero
	return b
}

// NewNotaCreditoReferenciada creates a builder for a credit note that references
// an existing electronic invoice (type 04). Call Referencia() with the CUFE.
func NewNotaCreditoReferenciada() *Builder {
	b := newBuilder(types.TipoDocNotaCreditoRef)
	b.doc.DatosTransaccion.NaturalezaOperacion = types.NatDevolucion
	return b
}

// NewNotaDebitoReferenciada creates a builder for a debit note that references an
// existing electronic invoice (type 05). Call Referencia() with the CUFE.
func NewNotaDebitoReferenciada() *Builder {
	b := newBuilder(types.TipoDocNotaDebitoRef)
	return b
}

// NewNotaCreditoGenerica creates a builder for a generic credit note (type 06),
// which must not reference any prior document.
func NewNotaCreditoGenerica() *Builder {
	b := newBuilder(types.TipoDocNotaCreditoGen)
	b.doc.DatosTransaccion.NaturalezaOperacion = types.NatDevolucion
	return b
}

// NewNotaDebitoGenerica creates a builder for a generic debit note (type 07),
// which must not reference any prior document.
func NewNotaDebitoGenerica() *Builder {
	return newBuilder(types.TipoDocNotaDebitoGen)
}

// NewFacturaZonaFranca creates a builder for a free-zone invoice (type 08).
func NewFacturaZonaFranca() *Builder {
	return newBuilder(types.TipoDocFacturaZonaFranca)
}

// NewReembolso creates a builder for a reimbursement document (type 09).
func NewReembolso() *Builder {
	return newBuilder(types.TipoDocReembolso)
}

// NewFacturaExtranjera creates a builder for a foreign-operation invoice (type 10).
func NewFacturaExtranjera() *Builder {
	return newBuilder(types.TipoDocFacturaExtranjera)
}

// NewFacturaGobierno creates a builder for an invoice to a Panamanian government entity (type 01, client 03).
// Every item added must carry valid CPBS codes (CodigoCPBSAbrev, CodigoCPBS, UnidadMedidaCPBS).
func NewFacturaGobierno() *Builder {
	b := newBuilder(types.TipoDocFacturaInterna)
	b.doc.DatosTransaccion.Cliente.TipoClienteFE = types.ClienteGobierno
	b.doc.DatosTransaccion.Cliente.TipoContribuyente = types.ContribuyenteJuridico
	b.doc.DatosTransaccion.Cliente.CodigoUbicacion = "8-8-8"
	b.doc.DatosTransaccion.Cliente.Pais = types.CountryPA
	return b
}

// NewFacturaConsumidorFinal creates a builder for a final-consumer domestic invoice (type 01, client 02).
func NewFacturaConsumidorFinal() *Builder {
	b := newBuilder(types.TipoDocFacturaInterna)
	b.doc.DatosTransaccion.Cliente.TipoClienteFE = types.ClienteConsumidorFinal
	b.doc.DatosTransaccion.Cliente.TipoContribuyente = types.ContribuyenteNatural
	b.doc.DatosTransaccion.Cliente.CodigoUbicacion = "8-8-8"
	b.doc.DatosTransaccion.Cliente.Pais = types.CountryPA
	return b
}

// Exportacion attaches the export data block required for export invoices (type 03)
// and any document whose DestinoOperacion is foreign.
func (b *Builder) Exportacion(d types.DatosExportacion) *Builder {
	b.doc.DatosTransaccion.DatosFacturaExportacion = &d
	return b
}

// Referencia adds a referenced electronic invoice by CUFE, required for credit and
// debit notes of type 04/05. The emission date of the referenced document is
// recorded so the SDK can verify the 180-day window.
func (b *Builder) Referencia(cufe string, fechaEmision time.Time) *Builder {
	b.doc.DatosTransaccion.ListaDocsFiscalReferenciados = append(
		b.doc.DatosTransaccion.ListaDocsFiscalReferenciados,
		types.DocFiscalReferenciado{
			CufeFEReferenciada:                cufe,
			FechaEmisionDocFiscalReferenciado: formatDateTime(fechaEmision),
		},
	)
	return b
}

// Retencion attaches a withholding block to the document.
func (b *Builder) Retencion(r types.Retencion) *Builder {
	b.doc.DatosTransaccion.Retencion = &r
	b.doc.TotalesSubTotales.Retencion = &r
	return b
}

// DescuentoGlobal attaches a document-level discount to the totals block.
func (b *Builder) DescuentoGlobal(monto float64, descripcion string) *Builder {
	b.doc.TotalesSubTotales.DescuentoBonificacion = &types.DescuentoBonificacion{
		MontoDescuento:       fmtAmount(monto, totalDecimals),
		DescripcionDescuento: descripcion,
	}
	return b
}

// Ubicacion sets or overrides the province-district-corregimiento code for the client.
func (b *Builder) Ubicacion(codigo string) *Builder {
	b.doc.DatosTransaccion.Cliente.CodigoUbicacion = codigo
	return b
}

// AutorizadoDescarga adds an entity authorized to download this electronic document and its events.
func (b *Builder) AutorizadoDescarga(tipo, ruc, dv, nombre string) *Builder {
	b.doc.AutorizadosDescarga = append(b.doc.AutorizadosDescarga, types.AutorizadoDescarga{
		TipoContribuyente: tipo,
		RUCReceptor:       ruc,
		DigitoVerifRUC:    dv,
		NombreAutorizado:  nombre,
	})
	return b
}

// PedidoGlobal sets the global purchase order details.
func (b *Builder) PedidoGlobal(numero string, fecha time.Time, monto float64, moneda string) *Builder {
	b.doc.PedidoGlobal = &types.PedidoComercialGlobal{
		NumeroPedido: numero,
		FechaPedido:  formatDateTime(fecha),
		MontoTotal:   fmtAmount(monto, totalDecimals),
		MonedaPedido: moneda,
	}
	return b
}
