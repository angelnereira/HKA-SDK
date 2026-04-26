# Ejemplo en lenguaje JAVA

> Fuente: https://felwiki.thefactoryhka.com.pa/lenguaje_java

Este código explica de manera sencilla cómo integrarse al servicio web de facturación electrónica.

[Código Ejemplo + Librerías (ZIP 16.5 MB)](https://felwiki.thefactoryhka.com.pa/_media/tf_pa_ejemplodir_enviar_wiki.zip)

## Pasos para realizar la integración

1. **PASO 1**: Configurar la ruta del WS y tokens de prueba
2. **PASO 2**: Construir el objeto `documentoElectronico`
3. **PASO 3**: Enviar mediante el método `enviar()`
4. **PASO 4**: Obtener la respuesta

## Estructura inicial

```java
package tf_pa_ejemplo_directo_java;

public class Tf_Pa_ejemplo_directo_java {
    public static void main(String[] args) {
        EnviarDocumentoElectronico enviar;
        enviar = new EnviarDocumentoElectronico();
    }
}
```

## Imports requeridos

```java
package tf_pa_ejemplo_directo_java;

import java.net.MalformedURLException;
import java.net.URISyntaxException;
import java.net.URL;
import java.text.SimpleDateFormat;
import java.util.Date;

import javax.xml.ws.BindingProvider;
import org.datacontract.schemas._2004._07.services.EnviarResponse;
import org.datacontract.schemas._2004._07.services_objcomprobante.ObjectFactory;
import org.datacontract.schemas._2004._07.services_objcomprobante.DocumentoElectronico;
import org.datacontract.schemas._2004._07.services_objcomprobante.Cliente;
import org.datacontract.schemas._2004._07.services_objcomprobante.DatosTransaccion;
import org.datacontract.schemas._2004._07.services_objcomprobante.DocFiscalReferenciado;
import org.datacontract.schemas._2004._07.services_objcomprobante.ListaDocsFiscalReferenciados;
import org.datacontract.schemas._2004._07.services_objcomprobante.ListaItems;
import org.datacontract.schemas._2004._07.services_objcomprobante.Item;
import org.datacontract.schemas._2004._07.services_objcomprobante.TotalesSubTotales;
import org.datacontract.schemas._2004._07.services_objcomprobante.ListaFormaPago;
import org.datacontract.schemas._2004._07.services_objcomprobante.FormaPago;
import org.datacontract.schemas._2004._07.services_objcomprobante.Retencion;
```

## PASO 1: Configurar WSDL y tokens

Cuenta demo para pruebas:
- **LINK**: `https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc?wsdl`
- **Usuario**: `"SOLICITAR"`
- **Password**: `"SOLICITAR"`

```java
public class EnviarDocumentoElectronico {

    String wsdl = "https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc?wsdl";
    String tokenEmpresa = "SOLICITAR";
    String tokenPassword = "SOLICITAR";

    org.tempuri.Service service = new org.tempuri.Service();
    org.tempuri.IService port = service.getBasicHttpBindingIService();
    BindingProvider bindingProvider = (BindingProvider) port;

    DocumentoElectronico documentoElectronico = new DocumentoElectronico();
    DatosTransaccion datosTransaccion = new DatosTransaccion();
    Item item = new Item();
    ObjectFactory factory = new ObjectFactory();
    Date fechaHoy = new Date();
    SimpleDateFormat fechaHoyFormat = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss'-05:00'");
    String fechaE = fechaHoyFormat.format(fechaHoy);
```

## PASO 2: Construir el Documento Electrónico

```java
private void buildDocumentoElectronico() {
    try {
        // ============== CLIENTE ==============
        Cliente cliente = new Cliente();
        cliente.setTipoClienteFE(factory.createClienteTipoClienteFE("01"));        // Contribuyente
        cliente.setTipoContribuyente(factory.createClienteTipoContribuyente("2")); // Juridico
        cliente.setNumeroRUC(factory.createClienteNumeroRUC("155596713-2-2015"));
        cliente.setDigitoVerificadorRUC(factory.createClienteDigitoVerificadorRUC("59"));
        cliente.setRazonSocial(factory.createClienteRazonSocial("Ambiente de pruebas"));
        cliente.setCodigoUbicacion(factory.createClienteCodigoUbicacion("1-2-3"));
        cliente.setDireccion(factory.createClienteDireccion("Ave. La Paz"));
        cliente.setCorregimiento(factory.createClienteCorregimiento("GUABITO"));
        cliente.setDistrito(factory.createClienteDistrito("CHANGUINOLA"));
        cliente.setProvincia(factory.createClienteProvincia("BOCAS DEL TORO"));
        cliente.setTelefono1(factory.createClienteTelefono1("9999-9999"));
        cliente.setCorreoElectronico1(factory.createClienteCorreoElectronico1("sucursal000@empresa.com"));
        cliente.setPais(factory.createClientePais("PA"));
        cliente.setNroIdentificacionExtranjero(factory.createClienteNroIdentificacionExtranjero(""));
        cliente.setPaisExtranjero(factory.createClientePaisExtranjero(""));
        datosTransaccion.setCliente(factory.createCliente(cliente));

        // ============== DATOS TRANSACCION ==============
        datosTransaccion.setNumeroDocumentoFiscal(factory.createDatosTransaccionNumeroDocumentoFiscal("7011"));
        datosTransaccion.setTipoEmision(factory.createDatosTransaccionTipoEmision("01"));
        datosTransaccion.setFechaInicioContingencia(factory.createDatosTransaccionFechaInicioContingencia(""));
        datosTransaccion.setMotivoContingencia(factory.createDatosTransaccionMotivoContingencia(""));
        datosTransaccion.setTipoDocumento(factory.createDatosTransaccionTipoDocumento("01"));
        datosTransaccion.setPuntoFacturacionFiscal(factory.createDatosTransaccionPuntoFacturacionFiscal("002"));
        datosTransaccion.setFechaEmision(factory.createDatosTransaccionFechaEmision(this.fechaE));
        datosTransaccion.setFechaSalida(factory.createDatosTransaccionFechaSalida(""));
        datosTransaccion.setNaturalezaOperacion(factory.createDatosTransaccionNaturalezaOperacion("01"));
        datosTransaccion.setTipoOperacion(factory.createDatosTransaccionTipoOperacion("1"));
        datosTransaccion.setDestinoOperacion(factory.createDatosTransaccionDestinoOperacion("1"));
        datosTransaccion.setFormatoCAFE(factory.createDatosTransaccionFormatoCAFE("1"));
        datosTransaccion.setEntregaCAFE(factory.createDatosTransaccionEntregaCAFE("1"));
        datosTransaccion.setEnvioContenedor(factory.createDatosTransaccionEnvioContenedor("1"));
        datosTransaccion.setProcesoGeneracion(factory.createDatosTransaccionProcesoGeneracion("1"));
        datosTransaccion.setTipoVenta(factory.createDatosTransaccionTipoVenta("1"));
        datosTransaccion.setInformacionInteres(factory.createDatosTransaccionInformacionInteres("Demo para Facturacion Electronica"));

        // Documento Fiscal Referenciado
        ListaDocsFiscalReferenciados listaDocsFiscalReferenciados = new ListaDocsFiscalReferenciados();
        DocFiscalReferenciado docFiscalReferenciado = new DocFiscalReferenciado();
        docFiscalReferenciado.setCufeFEReferenciada(factory.createDocFiscalReferenciadoCufeFEReferenciada(""));
        docFiscalReferenciado.setFechaEmisionDocFiscalReferenciado(factory.createDocFiscalReferenciadoFechaEmisionDocFiscalReferenciado(""));
        listaDocsFiscalReferenciados.getDocFiscalReferenciado().add(null);
        datosTransaccion.setListaDocsFiscalReferenciados(factory.createListaDocsFiscalReferenciados(null));

        documentoElectronico.setCodigoSucursalEmisor(factory.createDocumentoElectronicoCodigoSucursalEmisor("0000"));
        documentoElectronico.setTipoSucursal(factory.createDocumentoElectronicoTipoSucursal(""));
        documentoElectronico.setDatosTransaccion(factory.createDatosTransaccion(datosTransaccion));

        // ============== ITEM ==============
        ListaItems items = new ListaItems();
        item.setDescripcion(factory.createItemDescripcion("Lapiz"));
        item.setCodigo(factory.createItemCodigo("CA-001"));
        item.setUnidadMedida(factory.createItemUnidadMedida("cm"));
        item.setCantidad(factory.createItemCantidad("2.00"));
        item.setFechaFabricacion(factory.createItemFechaFabricacion(""));
        item.setFechaCaducidad(factory.createItemFechaCaducidad(""));
        item.setInfoItem(factory.createItemInfoItem("Lapiz con Goma"));
        item.setPrecioUnitario(factory.createItemPrecioUnitario("69.000000"));
        item.setPrecioUnitarioDescuento(factory.createItemPrecioUnitarioDescuento("0.00"));
        item.setPrecioItem(factory.createItemPrecioItem("138.00"));
        item.setPrecioAcarreo(factory.createItemPrecioAcarreo("1.01"));
        item.setPrecioSeguro(factory.createItemPrecioSeguro("12.01"));
        item.setValorTotal(factory.createItemValorTotal("171.72"));
        item.setCodigoGTIN(factory.createItemCodigoGTIN(""));
        item.setCantGTINCom(factory.createItemCantGTINCom(""));
        item.setCodigoGTINInv(factory.createItemCodigoGTINInv(""));
        item.setCantGTINComInv(factory.createItemCantGTINComInv(""));
        item.setTasaITBMS(factory.createItemTasaITBMS("03"));
        item.setValorITBMS(factory.createItemValorITBMS("20.70"));
        item.setTasaISC(factory.createItemTasaISC(""));
        item.setValorISC(factory.createItemValorISC(""));
        item.setCodigoCPBS(factory.createItemCodigoCPBS("1410"));
        item.setUnidadMedidaCPBS(factory.createItemUnidadMedidaCPBS("cm"));
        items.getItem().add(0, item);
        documentoElectronico.setListaItems(factory.createListaItems(items));

        // ============== TOTALES ==============
        TotalesSubTotales totalesSubTotales = new TotalesSubTotales();
        totalesSubTotales.setNroItems(factory.createTotalesSubTotalesNroItems("1"));
        totalesSubTotales.setTiempoPago(factory.createTotalesSubTotalesTiempoPago("1"));
        totalesSubTotales.setTotalAcarreoCobrado(factory.createTotalesSubTotalesTotalAcarreoCobrado(""));
        totalesSubTotales.setTotalDescuento(factory.createTotalesSubTotalesTotalDescuento("0.00"));
        totalesSubTotales.setTotalFactura(factory.createTotalesSubTotalesTotalFactura("171.72"));
        totalesSubTotales.setTotalISC(factory.createTotalesSubTotalesTotalISC(""));
        totalesSubTotales.setTotalITBMS(factory.createTotalesSubTotalesTotalITBMS("20.70"));
        totalesSubTotales.setTotalMontoGravado(factory.createTotalesSubTotalesTotalMontoGravado("20.70"));
        totalesSubTotales.setTotalPrecioNeto(factory.createTotalesSubTotalesTotalPrecioNeto("138.00"));
        totalesSubTotales.setTotalValorRecibido(factory.createTotalesSubTotalesTotalValorRecibido("171.72"));
        totalesSubTotales.setVuelto(factory.createTotalesSubTotalesVuelto("0.00"));
        totalesSubTotales.setValorSeguroCobrado(factory.createTotalesSubTotalesValorSeguroCobrado(""));
        totalesSubTotales.setTotalTodosItems(factory.createTotalesSubTotalesTotalTodosItems("171.72"));

        // Forma de Pago
        FormaPago formaPago = new FormaPago();
        ListaFormaPago listaFormaPago = new ListaFormaPago();
        formaPago.setFormaPagoFact(factory.createFormaPagoFormaPagoFact("01"));  // Credito
        formaPago.setValorCuotaPagada(factory.createFormaPagoValorCuotaPagada("171.72"));
        formaPago.setDescFormaPago(factory.createFormaPagoDescFormaPago(""));
        listaFormaPago.getFormaPago().add(0, formaPago);
        totalesSubTotales.setListaFormaPago(factory.createListaFormaPago(listaFormaPago));

        // Retención
        Retencion retencion = new Retencion();
        retencion.setCodigoRetencion(factory.createRetencionCodigoRetencion("8"));
        retencion.setMontoRetencion(factory.createRetencionMontoRetencion("1.00"));
        totalesSubTotales.setRetencion(factory.createRetencion(retencion));
        documentoElectronico.setTotalesSubTotales(factory.createTotalesSubTotales(totalesSubTotales));

    } catch (Exception e) {
        System.out.println("Exception :" + e.toString());
    }
}
```

## PASO 3: Enviar al WS

```java
public EnviarDocumentoElectronico() {
    EnviarResponse enviarResponse;
    try {
        bindingProvider.getRequestContext().put(BindingProvider.ENDPOINT_ADDRESS_PROPERTY, wsdl);

        // Construir el objeto
        buildDocumentoElectronico();

        // Enviar
        System.out.println("Enviando Documento Fiscal: " +
                          datosTransaccion.getNumeroDocumentoFiscal().getValue());
        enviarResponse = port.enviar(tokenEmpresa, tokenPassword, documentoElectronico);
```

## PASO 4: Obtener la respuesta

```java
        // Respuesta
        System.out.println("Respuesta: " + enviarResponse.getCodigo().getValue() +
                          " - " + enviarResponse.getMensaje().getValue());
        System.out.println("Cufe: " + enviarResponse.getCufe().getValue());

        // QR link
        URL url = new URL(enviarResponse.getQr().getValue());
        System.out.println(url.toURI());

    } catch (MalformedURLException | URISyntaxException e) {
        System.out.println("Exception :" + e.toString());
    }
}
```
