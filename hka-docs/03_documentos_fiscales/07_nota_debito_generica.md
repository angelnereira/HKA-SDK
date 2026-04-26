# Nota de Débito Genérica

> Fuente: https://felwiki.thefactoryhka.com.pa/nota_de_debito_generica
> **tipoDocumento = 07**

Nota de débito **sin referencia** a una FE anterior. Estructura idéntica a la Nota de Crédito Genérica, solo cambia `tipoDocumento = 07`.

## Campos clave distintivos

| Campo | Valor | Observación |
|---|---|---|
| tipoDocumento | `07` | **Nota de Débito genérica** (sin CUFE referenciado) |
| listaDocsFiscalReferenciados | AUSENTE | No se incluye este nodo |

## Ejemplo de XML

```xml
<soapenv:Body>
  <tem:Enviar>
    <tem:tokenEmpresa>SOLICITAR</tem:tokenEmpresa>
    <tem:tokenPassword>SOLICITAR</tem:tokenPassword>
    <tem:documento>
      <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
      <ser:tipoSucursal>1</ser:tipoSucursal>
      <ser:datosTransaccion>
        <ser:tipoEmision>01</ser:tipoEmision>
        <ser:tipoDocumento>07</ser:tipoDocumento>
        <ser:numeroDocumentoFiscal>5000010</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>001</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2020-12-29T08:28:28-05:00</ser:fechaEmision>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>3</ser:formatoCAFE>
        <ser:entregaCAFE>3</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:informacionInteres>Factura de nota de debito generica</ser:informacionInteres>
        <ser:cliente>
          <ser:tipoClienteFE>01</ser:tipoClienteFE>
          <ser:tipoContribuyente>2</ser:tipoContribuyente>
          <ser:numeroRUC>155596713-2-2015</ser:numeroRUC>
          <ser:digitoVerificadorRUC>59</ser:digitoVerificadorRUC>
          <ser:razonSocial>TFHKA</ser:razonSocial>
          <ser:direccion>Ave. La Paz</ser:direccion>
          <ser:codigoUbicacion>1-1-2</ser:codigoUbicacion>
          <ser:provincia>BOCAS DEL TORO</ser:provincia>
          <ser:distrito>BOCAS DEL TORO</ser:distrito>
          <ser:corregimiento>BASTIMENTOS</ser:corregimiento>
          <ser:telefono1>222-3456</ser:telefono1>
          <ser:pais>PA</ser:pais>
        </ser:cliente>
      </ser:datosTransaccion>
      <ser:listaItems>
        <ser:item>
          <ser:descripcion>Muebles</ser:descripcion>
          <ser:unidadMedida>um</ser:unidadMedida>
          <ser:cantidad>1.000000</ser:cantidad>
          <ser:codigoCPBS>1310</ser:codigoCPBS>
          <ser:unidadMedidaCPBS>cm</ser:unidadMedidaCPBS>
          <ser:precioUnitario>5.550000</ser:precioUnitario>
          <ser:precioItem>5.550000</ser:precioItem>
          <ser:valorTotal>5.938500</ser:valorTotal>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>0.3885</ser:valorITBMS>
        </ser:item>
      </ser:listaItems>
      <ser:totalesSubTotales>
        <ser:totalPrecioNeto>5.55</ser:totalPrecioNeto>
        <ser:totalITBMS>0.39</ser:totalITBMS>
        <ser:totalMontoGravado>0.39</ser:totalMontoGravado>
        <ser:totalFactura>5.94</ser:totalFactura>
        <ser:totalValorRecibido>5.94</ser:totalValorRecibido>
        <ser:tiempoPago>3</ser:tiempoPago>
        <ser:nroItems>1</ser:nroItems>
        <ser:totalTodosItems>5.94</ser:totalTodosItems>
        <ser:listaFormaPago>
          <ser:formaPago>
            <ser:formaPagoFact>02</ser:formaPagoFact>
            <ser:valorCuotaPagada>5.94</ser:valorCuotaPagada>
          </ser:formaPago>
        </ser:listaFormaPago>
      </ser:totalesSubTotales>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```
