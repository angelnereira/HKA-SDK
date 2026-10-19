package soap

import (
	"strings"
	"testing"

	"github.com/angelnereira/hka-sdk/types"
)

func TestBuildEstadoEnvelope(t *testing.T) {
	xml := BuildEstadoEnvelope("tokEmp", "tokPass", "0000", "0000000001", "001", "01", "01")
	if !strings.Contains(xml, "<tem:EstadoDocumento>") {
		t.Errorf("missing <tem:EstadoDocumento>")
	}
	if !strings.Contains(xml, "<tem:datosDocumento>") {
		t.Errorf("expected <tem:datosDocumento>, got: %s", xml)
	}
	if strings.Contains(xml, "estadoDocumentoRequest") {
		t.Errorf("found deprecated estadoDocumentoRequest tag")
	}
	if !strings.Contains(xml, "<ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>") {
		t.Errorf("missing sucursal tag")
	}
}

func TestBuildAnulacionEnvelope(t *testing.T) {
	xml := BuildAnulacionEnvelope("tokEmp", "tokPass", "0000", "0000000001", "001", "01", "01", "Factura duplicada")
	if !strings.Contains(xml, "<tem:AnulacionDocumento>") {
		t.Errorf("missing <tem:AnulacionDocumento>")
	}
	if !strings.Contains(xml, "<tem:motivoAnulacion>Factura duplicada</tem:motivoAnulacion>") {
		t.Errorf("expected <tem:motivoAnulacion> as direct child of AnulacionDocumento")
	}
	if !strings.Contains(xml, "<tem:datosDocumento>") {
		t.Errorf("expected <tem:datosDocumento>")
	}
	if strings.Contains(xml, "anulacionDocumentoRequest") {
		t.Errorf("found deprecated anulacionDocumentoRequest tag")
	}
}

func TestBuildDescargaXMLEnvelope(t *testing.T) {
	xml := BuildDescargaXMLEnvelope("tokEmp", "tokPass", "0000", "0000000001", "001", "01", "01")
	if !strings.Contains(xml, "<tem:DescargaXML>") {
		t.Errorf("missing <tem:DescargaXML>")
	}
	if !strings.Contains(xml, "<tem:datosDocumento>") {
		t.Errorf("expected <tem:datosDocumento>")
	}
	if strings.Contains(xml, "descargaXMLRequest") {
		t.Errorf("found deprecated descargaXMLRequest tag")
	}
}

func TestBuildDescargaPDFEnvelope(t *testing.T) {
	xml := BuildDescargaPDFEnvelope("tokEmp", "tokPass", "0000", "0000000001", "001", "01", "01", "DEV123")
	if !strings.Contains(xml, "<tem:DescargaPDF>") {
		t.Errorf("missing <tem:DescargaPDF>")
	}
	if !strings.Contains(xml, "<tem:datosDocumento>") {
		t.Errorf("expected <tem:datosDocumento>")
	}
	if !strings.Contains(xml, "<ser:serialDispositivo>DEV123</ser:serialDispositivo>") {
		t.Errorf("expected <ser:serialDispositivo>")
	}
	if strings.Contains(xml, "serieDispositivo") {
		t.Errorf("found invalid serieDispositivo tag (must be serialDispositivo)")
	}
	if strings.Contains(xml, "descargaPDFRequest") {
		t.Errorf("found deprecated descargaPDFRequest tag")
	}
}

func TestBuildEnvioCorreoEnvelope(t *testing.T) {
	xml := BuildEnvioCorreoEnvelope("tokEmp", "tokPass", "0000", "0000000001", "001", "01", "01", "cliente@empresa.com")
	if !strings.Contains(xml, "<tem:EnvioCorreo>") {
		t.Errorf("missing <tem:EnvioCorreo>")
	}
	if !strings.Contains(xml, "<tem:datosDocumento>") {
		t.Errorf("expected <tem:datosDocumento>")
	}
	if !strings.Contains(xml, "<tem:correo>cliente@empresa.com</tem:correo>") {
		t.Errorf("expected <tem:correo> as child of EnvioCorreo")
	}
	if strings.Contains(xml, "envioCorreoRequest") {
		t.Errorf("found deprecated envioCorreoRequest tag")
	}
}

func TestBuildFoliosEnvelope(t *testing.T) {
	xml := BuildFoliosEnvelope("tokEmp", "tokPass")
	if !strings.Contains(xml, "<tem:FoliosRestantes>") {
		t.Errorf("missing <tem:FoliosRestantes>")
	}
	if !strings.Contains(xml, "<tem:tokenEmpresa>tokEmp</tem:tokenEmpresa>") {
		t.Errorf("missing tokenEmpresa")
	}
}

func TestBuildRastreoCorreoEnvelope(t *testing.T) {
	cufe := "FE0120000155596713-2-2015-5900002021051200000992155050121410079104"
	xml := BuildRastreoCorreoEnvelope("tokEmp", "tokPass", cufe)
	if !strings.Contains(xml, "<tem:RastreoCorreo>") {
		t.Errorf("missing <tem:RastreoCorreo>")
	}
	if !strings.Contains(xml, "<tem:cufe>"+cufe+"</tem:cufe>") {
		t.Errorf("missing cufe tag")
	}
}

func TestBuildConsultarRucDVEnvelope(t *testing.T) {
	xml := BuildConsultarRucDVEnvelope("tokEmp", "tokPass", "2", "155596713-2-2015")
	if !strings.Contains(xml, "<tem:ConsultarRucDV>") {
		t.Errorf("missing <tem:ConsultarRucDV>")
	}
	if !strings.Contains(xml, "<tem:consultarRucDVRequest>") {
		t.Errorf("missing <tem:consultarRucDVRequest>")
	}
	if !strings.Contains(xml, "<ser:tipoRuc>2</ser:tipoRuc>") {
		t.Errorf("missing tipoRuc tag")
	}
	if !strings.Contains(xml, "<ser:ruc>155596713-2-2015</ser:ruc>") {
		t.Errorf("missing ruc tag")
	}
}

func TestBuildEnviarEnvelope_FullSchema(t *testing.T) {
	doc := &types.DocumentoElectronico{
		CodigoSucursalEmisor: "0000",
		TipoSucursal:         types.SucursalDetal,
		DatosTransaccion: types.DatosTransaccion{
			TipoEmision:            types.EmisionAUPNormal,
			TipoDocumento:          types.TipoDocFacturaInterna,
			NumeroDocumentoFiscal:  "0000000001",
			PuntoFacturacionFiscal: "001",
			FechaEmision:           "2026-10-18T10:00:00-05:00",
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
				RazonSocial:          "Empresa Demo S.A.",
				Direccion:            "Calle 50",
				Pais:                 types.CountryPA,
			},
		},
		ListaItems: []types.Item{
			{
				Descripcion:    "Consultoria",
				Cantidad:       "1.000000",
				PrecioUnitario: "100.000000",
				PrecioItem:     "100.000000",
				ValorTotal:     "107.000000",
				TasaITBMS:      types.ITBMS7,
				ValorITBMS:     "7.000000",
				ListaOTI: []types.OTI{
					{Tasa: types.OTISume911, ValorOTI: "1.00"},
				},
				PedidoItem: &types.PedidoComercialItem{
					NumeroPedido: "PO-001",
					NroLinea:     "1",
				},
			},
		},
		TotalesSubTotales: types.TotalesSubTotales{
			TotalPrecioNeto:   "100.00",
			TotalITBMS:        "7.00",
			TotalMontoGravado: "100.00",
			TotalFactura:      "108.00",
			TiempoPago:        types.PagoInmediato,
			NroItems:          1,
			TotalTodosItems:   "108.00",
			DescuentoBonificacion: &types.DescuentoBonificacion{
				DescripcionDescuento: "Descuento Pronto Pago",
				MontoDescuento:       "5.00",
			},
			ListaFormaPago: []types.FormaPagoItem{
				{FormaPagoFact: types.PagoEfectivo, ValorCuotaPagada: "108.00"},
			},
			Retencion: &types.Retencion{
				CodigoRetencion: types.RetencionServicioEstado100,
				MontoRetencion:  "7.00",
			},
			ListaTotalOTI: []types.TotalOTI{
				{Tasa: types.OTISume911, ValorTotalOTI: "1.00"},
			},
		},
		AutorizadosDescarga: []types.AutorizadoDescarga{
			{TipoContribuyente: "2", RUCReceptor: "155596713-2-2015", DigitoVerifRUC: "59"},
		},
		PedidoGlobal: &types.PedidoComercialGlobal{
			NumeroPedido: "PO-GLOBAL-99",
		},
	}

	xml := BuildEnviarEnvelope("tokEmp", "tokPass", doc)

	// Check listaItemOTI structure
	if !strings.Contains(xml, "<ser:listaItemOTI>") || !strings.Contains(xml, "<ser:tasaOTI>01</ser:tasaOTI>") {
		t.Errorf("missing or invalid listaItemOTI structure")
	}

	// Check pedidoComercialItem
	if !strings.Contains(xml, "<ser:pedidoComercialItem>") || !strings.Contains(xml, "<ser:nroPedidoCompraItem>PO-001</ser:nroPedidoCompraItem>") {
		t.Errorf("missing or invalid pedidoComercialItem structure")
	}

	// Check listaDescBonificacion
	if !strings.Contains(xml, "<ser:listaDescBonificacion>") || !strings.Contains(xml, "<ser:descDescuento>Descuento Pronto Pago</ser:descDescuento>") {
		t.Errorf("missing or invalid listaDescBonificacion")
	}

	// Check retencion in totalesSubTotales
	if !strings.Contains(xml, "<ser:retencion>") || !strings.Contains(xml, "<ser:codigoRetencion>1</ser:codigoRetencion>") {
		t.Errorf("missing or invalid retencion in totalesSubTotales")
	}

	// Check listaTotalOTI with totalOti and codigoTotalOTI
	if !strings.Contains(xml, "<ser:listaTotalOTI>") || !strings.Contains(xml, "<ser:codigoTotalOTI>01</ser:codigoTotalOTI>") {
		t.Errorf("missing or invalid totalOti with codigoTotalOTI")
	}

	// Check listaAutorizadosDescargaFEyEventos in datosTransaccion
	if !strings.Contains(xml, "<ser:listaAutorizadosDescargaFEyEventos>") {
		t.Errorf("missing listaAutorizadosDescargaFEyEventos in datosTransaccion")
	}

	// Check pedidoComercialGlobal
	if !strings.Contains(xml, "<ser:pedidoComercialGlobal>") || !strings.Contains(xml, "<ser:nroPedidoCompraGlobal>PO-GLOBAL-99</ser:nroPedidoCompraGlobal>") {
		t.Errorf("missing pedidoComercialGlobal")
	}
}

func TestParseFoliosResponse(t *testing.T) {
	xml := `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
	  <s:Body>
	    <FoliosRestantesResponse xmlns="http://tempuri.org/">
	      <FoliosRestantesResult xmlns:a="http://schemas.datacontract.org/2004/07/Services.Response">
	        <a:codigo>200</a:codigo>
	        <a:mensaje>Consulta exitosa</a:mensaje>
	        <a:licencia>LIC-123</a:licencia>
	        <a:fechaLicencia>2026-12-31</a:fechaLicencia>
	        <a:ciclo>1</a:ciclo>
	        <a:fechaCiclo>2026-06-01</a:fechaCiclo>
	        <a:foliosTotalesCiclo>1000</a:foliosTotalesCiclo>
	        <a:foliosUtilizadosCiclo>250</a:foliosUtilizadosCiclo>
	        <a:foliosDisponibleCiclo>750</a:foliosDisponibleCiclo>
	        <a:foliosTotales>5000</a:foliosTotales>
	        <a:resultado>procesado</a:resultado>
	      </FoliosRestantesResult>
	    </FoliosRestantesResponse>
	  </s:Body>
	</s:Envelope>`

	resp := ParseFoliosResponse([]byte(xml))
	if resp.Code != "200" {
		t.Errorf("expected code 200, got %s", resp.Code)
	}
	if resp.FoliosDisponibles != "750" {
		t.Errorf("expected foliosDisponibles 750, got %s", resp.FoliosDisponibles)
	}
	if resp.FoliosTotalesCiclo != "1000" {
		t.Errorf("expected foliosTotalesCiclo 1000, got %s", resp.FoliosTotalesCiclo)
	}
}

func TestParseTrackEmailResponse(t *testing.T) {
	xml := `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
	  <s:Body>
	    <RastreoCorreoResponse xmlns="http://tempuri.org/">
	      <RastreoCorreoResult xmlns:a="http://schemas.datacontract.org/2004/07/Services.Response">
	        <a:codigo>200</a:codigo>
	        <a:mensaje>Consulta exitosa</a:mensaje>
	        <a:listaRastreo xmlns:b="http://schemas.datacontract.org/2004/07/Services.Logic.Caching.Correo">
	          <b:listTracking>
	            <b:correo>destinatario@empresa.com</b:correo>
	            <b:creado_en>2026-05-10T14:30:00-05:00</b:creado_en>
	            <b:estado>DELIVERED</b:estado>
	            <b:messageId>msg-xyz-123</b:messageId>
	          </b:listTracking>
	        </a:listaRastreo>
	        <a:resultado>procesado</a:resultado>
	      </RastreoCorreoResult>
	    </RastreoCorreoResponse>
	  </s:Body>
	</s:Envelope>`

	resp := ParseTrackEmailResponse([]byte(xml))
	if resp.Code != "200" {
		t.Errorf("expected code 200, got %s", resp.Code)
	}
	if len(resp.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Items))
	}
	item := resp.Items[0]
	if item.Correo != "destinatario@empresa.com" {
		t.Errorf("expected correo destinatario@empresa.com, got %s", item.Correo)
	}
	if item.CreadoEn != "2026-05-10T14:30:00-05:00" {
		t.Errorf("expected creado_en parsed into CreadoEn, got %s", item.CreadoEn)
	}
	if item.Estado != "DELIVERED" {
		t.Errorf("expected estado DELIVERED, got %s", item.Estado)
	}
}
