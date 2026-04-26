# Factura de Exportación

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_de_exportacion
> **tipoDocumento = 03**

Factura emitida para operaciones de exportación a clientes en el extranjero.

## Campos clave distintivos

| Campo | Valor | Observación |
|---|---|---|
| tipoDocumento | `03` | Factura de exportación |
| destinoOperacion | `2` | **Extranjero** |
| tipoClienteFE | `04` | **Extranjero** — no RUC/DV |
| tipoIdentificacion | `99` | Otro (o 01:Pasaporte, 02:Número Tributario) |
| nroIdentificacionExtranjero | `123456789` | Obligatorio |
| pais | `VE` | Código ISO del país destino (**no PA**) |
| datosFacturaExportacion | OBLIGATORIO | Incluye condicionesEntrega, monedaOperExportacion, puertoEmbarque |

## Nodo `datosFacturaExportacion` requerido

```xml
<ser:datosFacturaExportacion>
  <ser:condicionesEntrega>EXW</ser:condicionesEntrega>
  <ser:monedaOperExportacion>USD</ser:monedaOperExportacion>
  <ser:puertoEmbarque>Prueba</ser:puertoEmbarque>
</ser:datosFacturaExportacion>
```

## Ejemplo de XML

```xml
<soapenv:Body>
  <tem:Enviar>
    <tem:tokenEmpresa>USUARIO</tem:tokenEmpresa>
    <tem:tokenPassword>PASSWORD</tem:tokenPassword>
    <tem:documento>
      <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
      <ser:tipoSucursal>1</ser:tipoSucursal>
      <ser:datosTransaccion>
        <ser:tipoEmision>01</ser:tipoEmision>
        <ser:tipoDocumento>03</ser:tipoDocumento>
        <ser:numeroDocumentoFiscal>5000006</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>052</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2020-12-29T08:28:28-05:00</ser:fechaEmision>
        <ser:fechaSalida>2020-12-29T08:28:28-05:00</ser:fechaSalida>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>2</ser:destinoOperacion>
        <ser:formatoCAFE>3</ser:formatoCAFE>
        <ser:entregaCAFE>3</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:informacionInteres>Factura de Exportacion</ser:informacionInteres>
        <ser:cliente>
          <ser:tipoClienteFE>04</ser:tipoClienteFE>
          <ser:razonSocial>TFHKA</ser:razonSocial>
          <ser:direccion>Ave. La Paz</ser:direccion>
          <ser:tipoIdentificacion>99</ser:tipoIdentificacion>
          <ser:nroIdentificacionExtranjero>123456789</ser:nroIdentificacionExtranjero>
          <ser:pais>VE</ser:pais>
        </ser:cliente>
        <ser:datosFacturaExportacion>
          <ser:condicionesEntrega>EXW</ser:condicionesEntrega>
          <ser:monedaOperExportacion>USD</ser:monedaOperExportacion>
          <ser:puertoEmbarque>Prueba</ser:puertoEmbarque>
        </ser:datosFacturaExportacion>
      </ser:datosTransaccion>
      <ser:listaItems>
        <ser:item>
          <ser:descripcion>Escritorio</ser:descripcion>
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
