# Factura con Descuento Global

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_con_descuento_global

Factura con **descuento aplicado a nivel de factura completa** (no a ítems individuales). El descuento no modifica los valores de los ítems sino que se resta del total.

## Campos clave distintivos

| Campo | Valor | Observación |
|---|---|---|
| totalDescuento | `100.00` | Suma de todos los descuentos globales |
| listaDescBonificacion | OBLIGATORIO | Lista con descripción y monto del descuento |
| totalTodosItems | `1070.00` | Total items con impuestos ANTES del descuento |
| totalFactura | `970.00` | totalTodosItems - totalDescuento |

## Fórmula clave

```
totalFactura = totalTodosItems + totalAcarreoCobrado + valorSeguroCobrado - totalDescuento
```

## Nodo `listaDescBonificacion` requerido

```xml
<ser:listaDescBonificacion>
  <ser:descuentoBonificacion>
    <ser:descDescuento>DESCUENTO DE 100 BALBOAS</ser:descDescuento>
    <ser:montoDescuento>100.00</ser:montoDescuento>
  </ser:descuentoBonificacion>
</ser:listaDescBonificacion>
```

## Ejemplo de XML

```xml
<soapenv:Body>
  <tem:Enviar>
    <tem:tokenEmpresa>Credenciales</tem:tokenEmpresa>
    <tem:tokenPassword>Credenciales</tem:tokenPassword>
    <tem:documento>
      <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
      <ser:tipoSucursal>1</ser:tipoSucursal>
      <ser:datosTransaccion>
        <ser:tipoEmision>01</ser:tipoEmision>
        <ser:tipoDocumento>01</ser:tipoDocumento>
        <ser:numeroDocumentoFiscal>500001579</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>007</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2022-06-30T01:00:00-05:00</ser:fechaEmision>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>1</ser:formatoCAFE>
        <ser:entregaCAFE>1</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:tipoVenta>1</ser:tipoVenta>
        <ser:informacionInteres>Informacion Interes</ser:informacionInteres>
        <ser:cliente>
          <ser:tipoClienteFE>02</ser:tipoClienteFE>
          <ser:razonSocial>Consumidor final para pruebas</ser:razonSocial>
          <ser:telefono1>214-4490</ser:telefono1>
          <ser:pais>PA</ser:pais>
        </ser:cliente>
      </ser:datosTransaccion>
      <ser:listaItems>
        <ser:item>
          <ser:descripcion>Pencil</ser:descripcion>
          <ser:codigo>PENC01</ser:codigo>
          <ser:unidadMedida>m</ser:unidadMedida>
          <ser:cantidad>1.00</ser:cantidad>
          <ser:infoItem>Pencil</ser:infoItem>
          <ser:precioUnitario>500.00</ser:precioUnitario>
          <ser:precioItem>500.00</ser:precioItem>
          <ser:valorTotal>535.00</ser:valorTotal>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>35.00</ser:valorITBMS>
        </ser:item>
        <ser:item>
          <ser:descripcion>Pencil</ser:descripcion>
          <ser:codigo>PENC01</ser:codigo>
          <ser:unidadMedida>m</ser:unidadMedida>
          <ser:cantidad>1.00</ser:cantidad>
          <ser:precioUnitario>500.00</ser:precioUnitario>
          <ser:precioItem>500.00</ser:precioItem>
          <ser:valorTotal>535.00</ser:valorTotal>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>35.00</ser:valorITBMS>
        </ser:item>
      </ser:listaItems>
      <ser:totalesSubTotales>
        <ser:totalPrecioNeto>1000.00</ser:totalPrecioNeto>
        <ser:totalITBMS>70.00</ser:totalITBMS>
        <ser:totalMontoGravado>70.00</ser:totalMontoGravado>
        <ser:totalDescuento>100.00</ser:totalDescuento>
        <ser:totalFactura>970.00</ser:totalFactura>
        <ser:totalValorRecibido>970.00</ser:totalValorRecibido>
        <ser:tiempoPago>1</ser:tiempoPago>
        <ser:nroItems>2</ser:nroItems>
        <ser:totalTodosItems>1070.00</ser:totalTodosItems>
        <ser:listaDescBonificacion>
          <ser:descuentoBonificacion>
            <ser:descDescuento>DESCUENTO DE 100 BALBOAS</ser:descDescuento>
            <ser:montoDescuento>100.00</ser:montoDescuento>
          </ser:descuentoBonificacion>
        </ser:listaDescBonificacion>
        <ser:listaFormaPago>
          <ser:formaPago>
            <ser:formaPagoFact>02</ser:formaPagoFact>
            <ser:valorCuotaPagada>970.00</ser:valorCuotaPagada>
          </ser:formaPago>
        </ser:listaFormaPago>
      </ser:totalesSubTotales>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```
