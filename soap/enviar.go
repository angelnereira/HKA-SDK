package soap

import (
	"fmt"
	"strings"

	"github.com/angelnereira/hka-sdk/internal"
	"github.com/angelnereira/hka-sdk/types"
)

const (
	nsEnv = "soapenv"
	nsTem = "tem"
	nsSer = "ser"

	nsEnvURI    = "http://schemas.xmlsoap.org/soap/envelope/"
	nsTemURI    = "http://tempuri.org/"
	nsSerDocURI = "http://schemas.datacontract.org/2004/07/Services.ObjComprobante.v1_0"
)

// BuildEnviarEnvelope constructs the complete SOAP XML for the Enviar() method.
func BuildEnviarEnvelope(tokenEmpresa, tokenPassword string, doc *types.DocumentoElectronico) string {
	var sb strings.Builder

	sb.WriteString(`<?xml version="1.0" encoding="utf-8"?>`)
	sb.WriteString(`<soapenv:Envelope`)
	sb.WriteString(` xmlns:soapenv="` + nsEnvURI + `"`)
	sb.WriteString(` xmlns:tem="` + nsTemURI + `"`)
	sb.WriteString(` xmlns:ser="` + nsSerDocURI + `">`)
	sb.WriteString(`<soapenv:Header/>`)
	sb.WriteString(`<soapenv:Body>`)
	sb.WriteString(`<tem:Enviar>`)
	internal.Tag(&sb, nsTem, "tokenEmpresa", tokenEmpresa)
	internal.Tag(&sb, nsTem, "tokenPassword", tokenPassword)
	internal.Open(&sb, nsTem, "documento")
	buildDocumento(&sb, doc)
	internal.Close(&sb, nsTem, "documento")
	sb.WriteString(`</tem:Enviar>`)
	sb.WriteString(`</soapenv:Body>`)
	sb.WriteString(`</soapenv:Envelope>`)

	return sb.String()
}

func buildDocumento(sb *strings.Builder, doc *types.DocumentoElectronico) {
	internal.Tag(sb, nsSer, "codigoSucursalEmisor", doc.CodigoSucursalEmisor)
	internal.Tag(sb, nsSer, "tipoSucursal", string(doc.TipoSucursal))
	buildDatosTransaccion(sb, &doc.DatosTransaccion)
	buildListaItems(sb, doc.ListaItems)
	buildTotales(sb, &doc.TotalesSubTotales)

	if len(doc.AutorizadosDescarga) > 0 {
		internal.Open(sb, nsSer, "autorizadosDescarga")
		for _, a := range doc.AutorizadosDescarga {
			internal.Open(sb, nsSer, "autorizadoDescarga")
			internal.Tag(sb, nsSer, "nombreAutorizado", a.NombreAutorizado)
			internal.Close(sb, nsSer, "autorizadoDescarga")
		}
		internal.Close(sb, nsSer, "autorizadosDescarga")
	}

	if doc.PedidoGlobal != nil {
		internal.Open(sb, nsSer, "pedidoGlobal")
		internal.Tag(sb, nsSer, "numeroPedido", doc.PedidoGlobal.NumeroPedido)
		internal.Tag(sb, nsSer, "fechaPedido", doc.PedidoGlobal.FechaPedido)
		internal.Tag(sb, nsSer, "nroLinea", doc.PedidoGlobal.NroLinea)
		internal.Tag(sb, nsSer, "montoTotal", doc.PedidoGlobal.MontoTotal)
		internal.Tag(sb, nsSer, "monedaPedido", doc.PedidoGlobal.MonedaPedido)
		internal.Close(sb, nsSer, "pedidoGlobal")
	}

	if doc.InfoLogistica != nil {
		buildInfoLogistica(sb, doc.InfoLogistica)
	}

	if doc.InfoEntrega != nil {
		buildInfoEntrega(sb, doc.InfoEntrega)
	}

	if doc.UsoPosterior != nil {
		internal.Open(sb, nsSer, "usoPosterior")
		internal.TagForce(sb, nsSer, "cufe", doc.UsoPosterior.CUFE)
		internal.Close(sb, nsSer, "usoPosterior")
	}
}

func buildDatosTransaccion(sb *strings.Builder, dt *types.DatosTransaccion) {
	internal.Open(sb, nsSer, "datosTransaccion")
	internal.Tag(sb, nsSer, "tipoEmision", string(dt.TipoEmision))
	internal.Tag(sb, nsSer, "fechaInicioContingencia", dt.FechaInicioContingencia)
	internal.Tag(sb, nsSer, "motivoContingencia", dt.MotivoContingencia)
	internal.Tag(sb, nsSer, "tipoDocumento", string(dt.TipoDocumento))
	internal.Tag(sb, nsSer, "numeroDocumentoFiscal", dt.NumeroDocumentoFiscal)
	internal.Tag(sb, nsSer, "puntoFacturacionFiscal", dt.PuntoFacturacionFiscal)
	internal.Tag(sb, nsSer, "fechaEmision", dt.FechaEmision)
	internal.Tag(sb, nsSer, "fechaSalida", dt.FechaSalida)
	internal.Tag(sb, nsSer, "naturalezaOperacion", string(dt.NaturalezaOperacion))
	internal.Tag(sb, nsSer, "tipoOperacion", string(dt.TipoOperacion))
	internal.Tag(sb, nsSer, "destinoOperacion", string(dt.DestinoOperacion))
	internal.Tag(sb, nsSer, "formatoCAFE", string(dt.FormatoCAFE))
	internal.Tag(sb, nsSer, "entregaCAFE", string(dt.EntregaCAFE))
	internal.Tag(sb, nsSer, "envioContenedor", string(dt.EnvioContenedor))
	internal.Tag(sb, nsSer, "procesoGeneracion", dt.ProcesoGeneracion)
	internal.Tag(sb, nsSer, "tipoVenta", string(dt.TipoVenta))
	internal.Tag(sb, nsSer, "informacionInteres", dt.InformacionInteres)
	buildCliente(sb, &dt.Cliente)

	if len(dt.ListaDocsFiscalReferenciados) > 0 {
		internal.Open(sb, nsSer, "listaDocsFiscalReferenciados")
		for _, ref := range dt.ListaDocsFiscalReferenciados {
			internal.Open(sb, nsSer, "docFiscalReferenciado")
			internal.Tag(sb, nsSer, "fechaEmisionDocFiscalReferenciado", ref.FechaEmisionDocFiscalReferenciado)
			internal.Tag(sb, nsSer, "cufeFEReferenciada", ref.CufeFEReferenciada)
			internal.Tag(sb, nsSer, "nroFacturaPapel", ref.NroFacturaPapel)
			internal.Tag(sb, nsSer, "nroFacturaImpFiscal", ref.NroFacturaImpFiscal)
			internal.Close(sb, nsSer, "docFiscalReferenciado")
		}
		internal.Close(sb, nsSer, "listaDocsFiscalReferenciados")
	}

	if dt.DatosFacturaExportacion != nil {
		buildExportacion(sb, dt.DatosFacturaExportacion)
	}

	if dt.Retencion != nil {
		internal.Open(sb, nsSer, "retencion")
		internal.Tag(sb, nsSer, "codigoRetencion", string(dt.Retencion.CodigoRetencion))
		internal.Tag(sb, nsSer, "fechaRetencion", dt.Retencion.FechaRetencion)
		internal.Tag(sb, nsSer, "montoRetencion", dt.Retencion.MontoRetencion)
		internal.Close(sb, nsSer, "retencion")
	}

	internal.Close(sb, nsSer, "datosTransaccion")
}

func buildCliente(sb *strings.Builder, c *types.Cliente) {
	internal.Open(sb, nsSer, "cliente")
	internal.Tag(sb, nsSer, "tipoClienteFE", string(c.TipoClienteFE))
	internal.Tag(sb, nsSer, "tipoContribuyente", string(c.TipoContribuyente))
	internal.Tag(sb, nsSer, "numeroRUC", c.NumeroRUC)
	internal.Tag(sb, nsSer, "digitoVerificadorRUC", c.DigitoVerificadorRUC)
	internal.Tag(sb, nsSer, "razonSocial", c.RazonSocial)
	internal.Tag(sb, nsSer, "direccion", c.Direccion)
	internal.Tag(sb, nsSer, "codigoUbicacion", c.CodigoUbicacion)
	internal.Tag(sb, nsSer, "provincia", c.Provincia)
	internal.Tag(sb, nsSer, "distrito", c.Distrito)
	internal.Tag(sb, nsSer, "corregimiento", c.Corregimiento)
	internal.Tag(sb, nsSer, "tipoIdentificacion", string(c.TipoIdentificacion))
	internal.Tag(sb, nsSer, "nroIdentificacionExtranjero", c.NroIdentificacionExtranjero)
	internal.Tag(sb, nsSer, "paisExtranjero", c.PaisExtranjero)
	internal.Tag(sb, nsSer, "telefono1", c.Telefono1)
	internal.Tag(sb, nsSer, "correoElectronico1", c.CorreoElectronico1)
	internal.Tag(sb, nsSer, "correoElectronico2", c.CorreoElectronico2)
	internal.Tag(sb, nsSer, "correoElectronico3", c.CorreoElectronico3)
	internal.Tag(sb, nsSer, "pais", string(c.Pais))
	internal.Tag(sb, nsSer, "paisOtro", c.PaisOtro)
	internal.Close(sb, nsSer, "cliente")
}

func buildExportacion(sb *strings.Builder, exp *types.DatosExportacion) {
	internal.Open(sb, nsSer, "datosFacturaExportacion")
	internal.Tag(sb, nsSer, "condicionesEntrega", string(exp.CondicionesEntrega))
	internal.Tag(sb, nsSer, "monedaOperExportacion", string(exp.MonedaOperExportacion))
	internal.Tag(sb, nsSer, "monedaOperExportacionNonDef", exp.MonedaOperExportacionNonDef)
	internal.Tag(sb, nsSer, "tipoDeCambio", exp.TipoDeCambio)
	internal.Tag(sb, nsSer, "montoMonedaExtranjera", exp.MontoMonedaExtranjera)
	internal.Tag(sb, nsSer, "puertoEmbarque", exp.PuertoEmbarque)
	internal.Close(sb, nsSer, "datosFacturaExportacion")
}

func buildListaItems(sb *strings.Builder, items []types.Item) {
	internal.Open(sb, nsSer, "listaItems")
	for _, item := range items {
		buildItem(sb, &item)
	}
	internal.Close(sb, nsSer, "listaItems")
}

func buildItem(sb *strings.Builder, item *types.Item) {
	internal.Open(sb, nsSer, "item")
	internal.Tag(sb, nsSer, "descripcion", item.Descripcion)
	internal.Tag(sb, nsSer, "codigo", item.Codigo)
	internal.Tag(sb, nsSer, "unidadMedida", item.UnidadMedida)
	internal.Tag(sb, nsSer, "cantidad", item.Cantidad)
	internal.Tag(sb, nsSer, "fechaFabricacion", item.FechaFabricacion)
	internal.Tag(sb, nsSer, "fechaCaducidad", item.FechaCaducidad)
	internal.Tag(sb, nsSer, "codigoCPBSAbrev", item.CodigoCPBSAbrev)
	internal.Tag(sb, nsSer, "codigoCPBS", item.CodigoCPBS)
	internal.Tag(sb, nsSer, "unidadMedidaCPBS", item.UnidadMedidaCPBS)
	internal.Tag(sb, nsSer, "infoItem", item.InfoItem)
	internal.Tag(sb, nsSer, "precioUnitario", item.PrecioUnitario)
	internal.Tag(sb, nsSer, "precioUnitarioDescuento", item.PrecioUnitarioDescuento)
	internal.Tag(sb, nsSer, "precioItem", item.PrecioItem)
	internal.Tag(sb, nsSer, "precioAcarreo", item.PrecioAcarreo)
	internal.Tag(sb, nsSer, "precioSeguro", item.PrecioSeguro)
	internal.Tag(sb, nsSer, "valorTotal", item.ValorTotal)
	internal.Tag(sb, nsSer, "codigoGTIN", item.CodigoGTIN)
	internal.Tag(sb, nsSer, "cantGTINCom", item.CantGTINCom)
	internal.Tag(sb, nsSer, "codigoGTINInv", item.CodigoGTINInv)
	internal.Tag(sb, nsSer, "cantGTINComInv", item.CantGTINComInv)
	internal.Tag(sb, nsSer, "tasaITBMS", string(item.TasaITBMS))
	internal.Tag(sb, nsSer, "valorITBMS", item.ValorITBMS)
	internal.Tag(sb, nsSer, "tasaISC", item.TasaISC)
	internal.Tag(sb, nsSer, "valorISC", item.ValorISC)

	if len(item.ListaOTI) > 0 {
		internal.Open(sb, nsSer, "listaOTI")
		for _, oti := range item.ListaOTI {
			internal.Open(sb, nsSer, "oti")
			internal.Tag(sb, nsSer, "tasa", string(oti.Tasa))
			internal.Tag(sb, nsSer, "valorOTI", oti.ValorOTI)
			internal.Tag(sb, nsSer, "flagMonto", oti.FlagMonto)
			internal.Close(sb, nsSer, "oti")
		}
		internal.Close(sb, nsSer, "listaOTI")
	}

	if item.Medicina != nil {
		internal.Open(sb, nsSer, "medicina")
		internal.Tag(sb, nsSer, "nroLote", item.Medicina.NroLote)
		internal.Tag(sb, nsSer, "cantProductosLote", item.Medicina.CantProductosLote)
		internal.Close(sb, nsSer, "medicina")
	}

	if item.Vehiculo != nil {
		internal.Open(sb, nsSer, "vehiculo")
		internal.Tag(sb, nsSer, "marca", item.Vehiculo.Marca)
		internal.Tag(sb, nsSer, "modelo", item.Vehiculo.Modelo)
		internal.Tag(sb, nsSer, "anio", item.Vehiculo.Anio)
		internal.Tag(sb, nsSer, "color", item.Vehiculo.Color)
		internal.Tag(sb, nsSer, "numeroVIN", item.Vehiculo.NumeroVIN)
		internal.Tag(sb, nsSer, "numeroMotor", item.Vehiculo.NumeroMotor)
		internal.Tag(sb, nsSer, "numeroPlaca", item.Vehiculo.NumeroPlaca)
		internal.Tag(sb, nsSer, "tipoCombustible", item.Vehiculo.TipoCombustible)
		internal.Tag(sb, nsSer, "precioBase", item.Vehiculo.PrecioBase)
		internal.Tag(sb, nsSer, "numeroDocAduana", item.Vehiculo.NumeroDocAduana)
		internal.Close(sb, nsSer, "vehiculo")
	}

	if item.PedidoItem != nil {
		internal.Open(sb, nsSer, "pedidoItem")
		internal.Tag(sb, nsSer, "numeroPedido", item.PedidoItem.NumeroPedido)
		internal.Tag(sb, nsSer, "nroLinea", item.PedidoItem.NroLinea)
		internal.Close(sb, nsSer, "pedidoItem")
	}

	internal.Close(sb, nsSer, "item")
}

func buildTotales(sb *strings.Builder, t *types.TotalesSubTotales) {
	internal.Open(sb, nsSer, "totalesSubTotales")
	internal.Tag(sb, nsSer, "totalPrecioNeto", t.TotalPrecioNeto)
	internal.Tag(sb, nsSer, "totalITBMS", t.TotalITBMS)
	internal.Tag(sb, nsSer, "totalISC", t.TotalISC)
	internal.Tag(sb, nsSer, "totalMontoGravado", t.TotalMontoGravado)
	internal.Tag(sb, nsSer, "totalDescuento", t.TotalDescuento)
	internal.Tag(sb, nsSer, "totalAcarreoCobrado", t.TotalAcarreoCobrado)
	internal.Tag(sb, nsSer, "valorSeguroCobrado", t.ValorSeguroCobrado)
	internal.Tag(sb, nsSer, "totalOtrosGastos", t.TotalOtrosGastos)
	internal.Tag(sb, nsSer, "totalFactura", t.TotalFactura)
	internal.Tag(sb, nsSer, "totalValorRecibido", t.TotalValorRecibido)
	internal.Tag(sb, nsSer, "vuelto", t.Vuelto)
	internal.Tag(sb, nsSer, "tiempoPago", string(t.TiempoPago))
	if t.NroItems > 0 {
		internal.Tag(sb, nsSer, "nroItems", fmt.Sprintf("%d", t.NroItems))
	}
	internal.Tag(sb, nsSer, "totalTodosItems", t.TotalTodosItems)

	if len(t.ListaFormaPago) > 0 {
		internal.Open(sb, nsSer, "listaFormaPago")
		for _, fp := range t.ListaFormaPago {
			internal.Open(sb, nsSer, "formaPago")
			internal.Tag(sb, nsSer, "formaPagoFact", string(fp.FormaPagoFact))
			internal.Tag(sb, nsSer, "descFormaPago", fp.DescFormaPago)
			internal.Tag(sb, nsSer, "valorCuotaPagada", fp.ValorCuotaPagada)
			internal.Close(sb, nsSer, "formaPago")
		}
		internal.Close(sb, nsSer, "listaFormaPago")
	}

	if len(t.ListaPagoPlazo) > 0 {
		internal.Open(sb, nsSer, "listaPagoPlazo")
		for _, pp := range t.ListaPagoPlazo {
			internal.Open(sb, nsSer, "pagoPlazo")
			internal.Tag(sb, nsSer, "fechaPago", pp.FechaPago)
			internal.Tag(sb, nsSer, "valorCuotaPlazo", pp.ValorCuotaPlazo)
			internal.Tag(sb, nsSer, "descripcionCuota", pp.DescripcionCuota)
			internal.Close(sb, nsSer, "pagoPlazo")
		}
		internal.Close(sb, nsSer, "listaPagoPlazo")
	}

	if len(t.ListaTotalOTI) > 0 {
		internal.Open(sb, nsSer, "listaTotalOTI")
		for _, oti := range t.ListaTotalOTI {
			internal.Open(sb, nsSer, "totalOTI")
			internal.Tag(sb, nsSer, "tasa", string(oti.Tasa))
			internal.Tag(sb, nsSer, "valorTotalOTI", oti.ValorTotalOTI)
			internal.Close(sb, nsSer, "totalOTI")
		}
		internal.Close(sb, nsSer, "listaTotalOTI")
	}

	if t.DescuentoBonificacion != nil {
		internal.Open(sb, nsSer, "descuentoBonificacion")
		internal.Tag(sb, nsSer, "montoDescuento", t.DescuentoBonificacion.MontoDescuento)
		internal.Tag(sb, nsSer, "descripcionDescuento", t.DescuentoBonificacion.DescripcionDescuento)
		internal.Close(sb, nsSer, "descuentoBonificacion")
	}

	internal.Close(sb, nsSer, "totalesSubTotales")
}

func buildInfoLogistica(sb *strings.Builder, il *types.InfoLogistica) {
	internal.Open(sb, nsSer, "infoLogistica")
	internal.Tag(sb, nsSer, "entregaFisica", il.EntregaFisica)
	internal.Tag(sb, nsSer, "modalidadFlete", il.ModalidadFlete)
	internal.Tag(sb, nsSer, "clausulaSeguro", il.ClausulaSeguro)
	internal.Tag(sb, nsSer, "nroContenedor", il.NroContenedor)
	internal.Tag(sb, nsSer, "nroManifiesto", il.NroManifiesto)
	internal.Tag(sb, nsSer, "nroOrdenEmbarque", il.NroOrdenEmbarque)
	internal.Tag(sb, nsSer, "puertoOrigen", il.PuertoOrigen)
	internal.Tag(sb, nsSer, "puertoDestino", il.PuertoDestino)
	internal.Tag(sb, nsSer, "paisDestino", string(il.PaisDestino))
	internal.Tag(sb, nsSer, "fechaEstimadaEntrega", il.FechaEstimadaEntrega)
	internal.Close(sb, nsSer, "infoLogistica")
}

func buildInfoEntrega(sb *strings.Builder, ie *types.InfoEntrega) {
	internal.Open(sb, nsSer, "infoEntrega")
	internal.Tag(sb, nsSer, "entregaA", ie.EntregaA)
	internal.Tag(sb, nsSer, "direccionEntrega", ie.DireccionEntrega)
	internal.Tag(sb, nsSer, "provincia", ie.Provincia)
	internal.Tag(sb, nsSer, "distrito", ie.Distrito)
	internal.Tag(sb, nsSer, "corregimiento", ie.Corregimiento)
	internal.Tag(sb, nsSer, "paisEntrega", string(ie.PaisEntrega))
	internal.Close(sb, nsSer, "infoEntrega")
}
