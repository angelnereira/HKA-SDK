# Factura de Operación Interna

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_de_operacion_interna
> **tipoDocumento = 01**

Documento fiscal más común: venta de bienes o servicios dentro de Panamá.

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
        <ser:tipoDocumento>01</ser:tipoDocumento>
        <ser:numeroDocumentoFiscal>5000009</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>001</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2020-12-29T08:28:28-05:00</ser:fechaEmision>
        <ser:fechaSalida>2020-12-29T08:28:28-05:00</ser:fechaSalida>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>3</ser:formatoCAFE>
        <ser:entregaCAFE>3</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:tipoVenta></ser:tipoVenta>
        <ser:informacionInteres>Factura interna</ser:informacionInteres>
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
          <ser:paisExtranjero></ser:paisExtranjero>
          <ser:telefono1>222-3456</ser:telefono1>
          <ser:pais>PA</ser:pais>
          <ser:paisOtro></ser:paisOtro>
        </ser:cliente>
      </ser:datosTransaccion>
      <ser:listaItems>
        <ser:item>
          <ser:descripcion>Muebles</ser:descripcion>
          <ser:codigo></ser:codigo>
          <ser:unidadMedida>um</ser:unidadMedida>
          <ser:cantidad>1.000000</ser:cantidad>
          <ser:codigoCPBS>1310</ser:codigoCPBS>
          <ser:unidadMedidaCPBS>cm</ser:unidadMedidaCPBS>
          <ser:precioUnitario>5.550000</ser:precioUnitario>
          <ser:precioItem>5.550000</ser:precioItem>
          <ser:valorTotal>5.938500</ser:valorTotal>
          <ser:codigoGTIN>0</ser:codigoGTIN>
          <ser:cantGTINCom>0.00</ser:cantGTINCom>
          <ser:codigoGTINInv>0</ser:codigoGTINInv>
          <ser:cantGTINComInv>0.00</ser:cantGTINComInv>
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
        <ser:listaPagoPlazo>
          <ser:pagoPlazo>
            <ser:fechaVenceCuota>2020-12-29T08:28:28-05:00</ser:fechaVenceCuota>
            <ser:valorCuota>5.94</ser:valorCuota>
          </ser:pagoPlazo>
        </ser:listaPagoPlazo>
      </ser:totalesSubTotales>
      <ser:usoPosterior>
        <ser:cufe></ser:cufe>
      </ser:usoPosterior>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```

## Campos clave

| Campo | Valor | Observación |
|---|---|---|
| tipoDocumento | `01` | Factura de operación interna |
| tipoEmision | `01` | Autorización Uso Previa, operación normal |
| naturalezaOperacion | `01` | Venta |
| tipoOperacion | `1` | Salida o venta |
| destinoOperacion | `1` | Panamá |
| tipoClienteFE | `01` | Contribuyente |
| pais | `PA` | Panamá |
| tasaITBMS | `01` | 7% (otros: 00=exento, 02=10%, 03=15%) |
