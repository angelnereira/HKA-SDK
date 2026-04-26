# Factura a Consumidor Final

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_a_consumidor_final
> **tipoDocumento = 01 con tipoClienteFE = 02**

Factura emitida a un consumidor final (persona natural sin relación B2B). El nodo `cliente` es **mínimo** — solo requiere `tipoClienteFE` y `pais`. Es el caso más simple y común en POS / retail.

## Campos clave distintivos

| Campo | Valor | Observación |
|---|---|---|
| tipoClienteFE | `02` | **Consumidor final** |
| cliente | Mínimo | Solo `tipoClienteFE` y `pais` son obligatorios |
| numeroRUC | (opcional) | Si el cliente pide factura con su cédula, se puede llenar con la cédula |
| formatoCAFE | `2` | Típicamente cinta de papel (ticket) para consumidor final |

**Simplificación clave**: A diferencia de las facturas a Contribuyente (01), Gobierno (03) o Extranjero (04), al Consumidor Final (02) **no se le requiere razón social, dirección, ni codigoUbicacion**.

## Nota sobre múltiples formas de pago

Este ejemplo muestra un caso con **dos formas de pago simultáneas** (Crédito + Efectivo) y **pagos en cuotas** (`listaPagoPlazo` con 2 cuotas) — útil para ventas mixtas en POS.

## Ejemplo de XML

```xml
<tem:documento>
  <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
  <ser:tipoSucursal>1</ser:tipoSucursal>
  <ser:datosTransaccion>
    <ser:tipoEmision>01</ser:tipoEmision>
    <ser:tipoDocumento>01</ser:tipoDocumento>
    <ser:numeroDocumentoFiscal>51340013</ser:numeroDocumentoFiscal>
    <ser:puntoFacturacionFiscal>001</ser:puntoFacturacionFiscal>
    <ser:fechaEmision>2022-08-03T17:01:28-05:00</ser:fechaEmision>
    <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
    <ser:tipoOperacion>1</ser:tipoOperacion>
    <ser:destinoOperacion>1</ser:destinoOperacion>
    <ser:formatoCAFE>2</ser:formatoCAFE>
    <ser:entregaCAFE>3</ser:entregaCAFE>
    <ser:envioContenedor>1</ser:envioContenedor>
    <ser:procesoGeneracion>1</ser:procesoGeneracion>
    <ser:informacionInteres><![CDATA[Prueba de facturacion electronica <br/> prueba2]]></ser:informacionInteres>
    <ser:cliente>
      <ser:tipoClienteFE>02</ser:tipoClienteFE>
      <ser:pais>PA</ser:pais>
    </ser:cliente>
  </ser:datosTransaccion>
  <ser:listaItems>
    <ser:item>
      <ser:descripcion>ITEM FACT ELECTRONICA</ser:descripcion>
      <ser:codigo>FEL002</ser:codigo>
      <ser:unidadMedida>und</ser:unidadMedida>
      <ser:cantidad>1.00</ser:cantidad>
      <ser:infoItem>-prueba-</ser:infoItem>
      <ser:precioUnitario>100.00</ser:precioUnitario>
      <ser:precioUnitarioDescuento>0.00</ser:precioUnitarioDescuento>
      <ser:precioItem>100.00</ser:precioItem>
      <ser:valorTotal>107.00</ser:valorTotal>
      <ser:tasaITBMS>01</ser:tasaITBMS>
      <ser:valorITBMS>7.00</ser:valorITBMS>
    </ser:item>
  </ser:listaItems>
  <ser:totalesSubTotales>
    <ser:totalPrecioNeto>100.00</ser:totalPrecioNeto>
    <ser:totalITBMS>7.00</ser:totalITBMS>
    <ser:totalMontoGravado>7.00</ser:totalMontoGravado>
    <ser:totalDescuento>0.00</ser:totalDescuento>
    <ser:totalFactura>107.00</ser:totalFactura>
    <ser:totalValorRecibido>107.00</ser:totalValorRecibido>
    <ser:vuelto>0.00</ser:vuelto>
    <ser:tiempoPago>2</ser:tiempoPago>
    <ser:nroItems>1</ser:nroItems>
    <ser:totalTodosItems>107.00</ser:totalTodosItems>
    <ser:listaFormaPago>
      <ser:formaPago>
        <ser:formaPagoFact>01</ser:formaPagoFact>
        <ser:valorCuotaPagada>100.00</ser:valorCuotaPagada>
      </ser:formaPago>
      <ser:formaPago>
        <ser:formaPagoFact>02</ser:formaPagoFact>
        <ser:valorCuotaPagada>7.00</ser:valorCuotaPagada>
      </ser:formaPago>
    </ser:listaFormaPago>
    <ser:listaPagoPlazo>
      <ser:pagoPlazo>
        <ser:fechaVenceCuota>2022-08-21T10:37:38-05:00</ser:fechaVenceCuota>
        <ser:valorCuota>100.00</ser:valorCuota>
      </ser:pagoPlazo>
      <ser:pagoPlazo>
        <ser:fechaVenceCuota>2022-08-29T10:37:38-05:00</ser:fechaVenceCuota>
        <ser:valorCuota>7.00</ser:valorCuota>
      </ser:pagoPlazo>
    </ser:listaPagoPlazo>
  </ser:totalesSubTotales>
</tem:documento>
```
