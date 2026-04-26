# Factura a Cliente Extranjero

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_a_cliente_extranjero
> **tipoDocumento = 01 con tipoClienteFE = 04**

Factura emitida a un cliente extranjero **pero la operación se realiza en Panamá** (destinoOperacion=1). Diferente de la Factura de Exportación (donde destinoOperacion=2).

## Campos clave distintivos

| Campo | Valor | Observación |
|---|---|---|
| tipoDocumento | `01` | Factura operación interna (normal) |
| destinoOperacion | `1` | **Panamá** (la operación es local) |
| tipoClienteFE | `04` | **Extranjero** |
| tipoIdentificacion | `01` | Pasaporte (otros: 02=Núm. Tributario, 99=Otro) |
| nroIdentificacionExtranjero | `123456789` | Obligatorio |
| paisExtranjero | `Colombia` | País de origen del cliente (nombre completo si pasaporte) |
| pais | `PA` | **PA porque destinoOperacion=1** |

**No enviar cuando tipoClienteFE=04:** `tipoContribuyente`, `numeroRUC`, `digitoVerificadorRUC`, `codigoUbicacion`, `provincia`, `distrito`, `corregimiento`

## Ejemplo de XML

```xml
<soapenv:Body>
  <tem:Enviar>
    <tem:tokenEmpresa>Credenciales</tem:tokenEmpresa>
    <tem:tokenPassword>Credenciales</tem:tokenPassword>
    <tem:documento>
      <ser:codigoSucursalEmisor>0001</ser:codigoSucursalEmisor>
      <ser:tipoSucursal>1</ser:tipoSucursal>
      <ser:datosTransaccion>
        <ser:tipoEmision>01</ser:tipoEmision>
        <ser:tipoDocumento>01</ser:tipoDocumento>
        <ser:numeroDocumentoFiscal>0000056</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>001</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2022-05-24T08:28:28-05:00</ser:fechaEmision>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>3</ser:formatoCAFE>
        <ser:entregaCAFE>3</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:informacionInteres>Ejemplo Factura Cliente Extranjero</ser:informacionInteres>
        <ser:cliente>
          <ser:tipoClienteFE>04</ser:tipoClienteFE>
          <ser:razonSocial>Cliente Extranjero</ser:razonSocial>
          <ser:tipoIdentificacion>01</ser:tipoIdentificacion>
          <ser:nroIdentificacionExtranjero>123456789</ser:nroIdentificacionExtranjero>
          <ser:paisExtranjero>Colombia</ser:paisExtranjero>
          <ser:pais>PA</ser:pais>
        </ser:cliente>
      </ser:datosTransaccion>
      <ser:listaItems>
        <ser:item>
          <ser:descripcion>Escritorio</ser:descripcion>
          <ser:unidadMedida>um</ser:unidadMedida>
          <ser:cantidad>1.00</ser:cantidad>
          <ser:codigoCPBS>1310</ser:codigoCPBS>
          <ser:unidadMedidaCPBS>cm</ser:unidadMedidaCPBS>
          <ser:precioUnitario>5.55</ser:precioUnitario>
          <ser:precioItem>5.55</ser:precioItem>
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
