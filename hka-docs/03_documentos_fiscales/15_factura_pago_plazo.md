# Factura con Pago a Plazo

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_con_pago_plazo

Factura pagadera en **cuotas con fechas de vencimiento diferidas**. Usa `tiempoPago=2` (Plazo) y `listaPagoPlazo` con múltiples entradas.

## Campos clave distintivos

| Campo | Valor | Observación |
|---|---|---|
| tiempoPago | `2` | **Plazo** (otros: 1=Inmediato, 3=Mixto) |
| listaPagoPlazo | Múltiples | Una entrada `pagoPlazo` por cada cuota |
| fechaVenceCuota | ISO 8601 | Fecha de vencimiento de cada cuota |
| valorCuota | Numérico | Valor de cada cuota |

## Suma de cuotas

```
Σ valorCuota  = totalFactura
100.00 + 7.00 = 107.00 ✓
```

## Ejemplo de XML

```xml
<soapenv:Body>
  <tem:Enviar>
    <tem:tokenEmpresa>credenciales</tem:tokenEmpresa>
    <tem:tokenPassword>credenciales</tem:tokenPassword>
    <tem:documento>
      <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
      <ser:tipoSucursal>1</ser:tipoSucursal>
      <ser:datosTransaccion>
        <ser:tipoEmision>01</ser:tipoEmision>
        <ser:tipoDocumento>01</ser:tipoDocumento>
        <ser:numeroDocumentoFiscal>51340011</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>097</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2022-07-01T11:01:28-05:00</ser:fechaEmision>
        <ser:naturalezaOperacion>20</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>2</ser:formatoCAFE>
        <ser:entregaCAFE>3</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:informacionInteres><![CDATA[Prueba de facturacion electronica <br/> prueba2]]></ser:informacionInteres>
        <ser:cliente>
          <ser:tipoClienteFE>02</ser:tipoClienteFE>
          <ser:tipoContribuyente>1</ser:tipoContribuyente>
          <ser:numeroRUC>000000</ser:numeroRUC>
          <ser:razonSocial>Consumidor final Prueba</ser:razonSocial>
          <ser:direccion>Ciudad de Panama</ser:direccion>
          <ser:codigoUbicacion>7-3-1</ser:codigoUbicacion>
          <ser:provincia>Bastimentos</ser:provincia>
          <ser:distrito>Bastimentos</ser:distrito>
          <ser:corregimiento>Bocas del Toro</ser:corregimiento>
          <ser:telefono1>235-2352</ser:telefono1>
          <ser:pais>PA</ser:pais>
        </ser:cliente>
      </ser:datosTransaccion>
      <ser:listaItems>
        <ser:item>
          <ser:descripcion>ITEM FACT ELECTRONICA</ser:descripcion>
          <ser:codigo>FEL002</ser:codigo>
          <ser:unidadMedida>und</ser:unidadMedida>
          <ser:cantidad>1.00</ser:cantidad>
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
            <ser:fechaVenceCuota>2022-07-21T10:37:38-05:00</ser:fechaVenceCuota>
            <ser:valorCuota>100.00</ser:valorCuota>
          </ser:pagoPlazo>
          <ser:pagoPlazo>
            <ser:fechaVenceCuota>2022-07-29T10:37:38-05:00</ser:fechaVenceCuota>
            <ser:valorCuota>7.00</ser:valorCuota>
          </ser:pagoPlazo>
        </ser:listaPagoPlazo>
      </ser:totalesSubTotales>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```
