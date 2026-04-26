# Factura a Cliente Gobierno

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_a_cliente_gobierno
> **tipoDocumento = 01 con tipoClienteFE = 03**

Factura emitida a una entidad gubernamental (Ministerio, Contraloría, etc.). Requiere campos adicionales de CPBS obligatorios.

## Campos clave distintivos

| Campo | Valor | Observación |
|---|---|---|
| tipoClienteFE | `03` | **Gobierno** |
| numeroRUC | `8-NT-1-13656` | RUC gubernamental tipo "NT" (Número Tributario) |
| codigoCPBSAbrev | `72` | **OBLIGATORIO solo cuando tipoCliente=03** (Segmentos y Familias CPBS) |
| codigoCPBS | `7210` | **OBLIGATORIO cuando es venta a Administración Pública** |
| unidadMedidaCPBS | `und` | **OBLIGATORIO cuando tipoCliente=03** |

> ⚠️ **CRÍTICO**: Las facturas a Gobierno tienen reglas de CPBS (Codificación Panameña de Bienes y Servicios) OBLIGATORIAS que no aplican para otros tipos de cliente. Si falta cualquiera de `codigoCPBSAbrev`, `codigoCPBS` o `unidadMedidaCPBS` la factura será rechazada por la DGI.

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
        <ser:numeroDocumentoFiscal>0000000001</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>001</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2022-05-23T11:25:48-05:00</ser:fechaEmision>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>1</ser:formatoCAFE>
        <ser:entregaCAFE>3</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:tipoVenta>1</ser:tipoVenta>
        <ser:informacionInteres>Ejemplo de tipo gobierno</ser:informacionInteres>
        <ser:cliente>
          <ser:tipoClienteFE>03</ser:tipoClienteFE>
          <ser:tipoContribuyente>2</ser:tipoContribuyente>
          <ser:numeroRUC>8-NT-1-13656</ser:numeroRUC>
          <ser:digitoVerificadorRUC>43</ser:digitoVerificadorRUC>
          <ser:razonSocial>MINISTERIO DE EDUCACION</ser:razonSocial>
          <ser:direccion>Calle Ernesto Jaén Guardia, Edificio 6525</ser:direccion>
          <ser:codigoUbicacion>8-8-14</ser:codigoUbicacion>
          <ser:provincia>PANAMA</ser:provincia>
          <ser:distrito>PANAMÁ</ser:distrito>
          <ser:corregimiento>ANCON</ser:corregimiento>
          <ser:pais>PA</ser:pais>
        </ser:cliente>
      </ser:datosTransaccion>
      <ser:listaItems>
        <ser:item>
          <ser:descripcion>Materiales de escuela</ser:descripcion>
          <ser:codigo>01</ser:codigo>
          <ser:unidadMedida>und</ser:unidadMedida>
          <ser:cantidad>1.000</ser:cantidad>
          <ser:codigoCPBSAbrev>72</ser:codigoCPBSAbrev>
          <ser:codigoCPBS>7210</ser:codigoCPBS>
          <ser:unidadMedidaCPBS>und</ser:unidadMedidaCPBS>
          <ser:precioUnitario>93454.100000</ser:precioUnitario>
          <ser:precioUnitarioDescuento>0.000000</ser:precioUnitarioDescuento>
          <ser:precioItem>93454.100000</ser:precioItem>
          <ser:valorTotal>99995.887000</ser:valorTotal>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>6541.787000</ser:valorITBMS>
        </ser:item>
      </ser:listaItems>
      <ser:totalesSubTotales>
        <ser:totalPrecioNeto>93454.10</ser:totalPrecioNeto>
        <ser:totalITBMS>6541.78</ser:totalITBMS>
        <ser:totalMontoGravado>6541.78</ser:totalMontoGravado>
        <ser:totalDescuento>0.00</ser:totalDescuento>
        <ser:totalAcarreoCobrado>0.00</ser:totalAcarreoCobrado>
        <ser:valorSeguroCobrado>0.00</ser:valorSeguroCobrado>
        <ser:totalFactura>99995.88</ser:totalFactura>
        <ser:totalValorRecibido>99995.88</ser:totalValorRecibido>
        <ser:vuelto>0.00</ser:vuelto>
        <ser:tiempoPago>3</ser:tiempoPago>
        <ser:nroItems>1</ser:nroItems>
        <ser:totalTodosItems>99995.88</ser:totalTodosItems>
        <ser:listaFormaPago>
          <ser:formaPago>
            <ser:formaPagoFact>02</ser:formaPagoFact>
            <ser:valorCuotaPagada>99995.88</ser:valorCuotaPagada>
          </ser:formaPago>
        </ser:listaFormaPago>
        <ser:listaPagoPlazo>
          <ser:pagoPlazo>
            <ser:fechaVenceCuota>2022-05-23T11:25:48-05:00</ser:fechaVenceCuota>
            <ser:valorCuota>99995.88</ser:valorCuota>
            <ser:infoPagoCuota>Cuota automatica de sistema</ser:infoPagoCuota>
          </ser:pagoPlazo>
        </ser:listaPagoPlazo>
      </ser:totalesSubTotales>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```
