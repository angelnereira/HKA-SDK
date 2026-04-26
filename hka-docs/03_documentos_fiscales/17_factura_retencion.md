# Factura con Retención

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_con_retencion

Factura donde se aplica **retención** al ITBMS. Común en operaciones con Gobierno o en pagos a no domiciliados.

## Códigos de Retención (codigoRetencion)

| Código | Descripción |
|---|---|
| 1 | Pago por servicio profesional al estado — **100%** |
| 2 | Pago por venta de bienes/servicios al estado — **50%** |
| 3 | Pago o acreditación a no domiciliado o empresa constituida en el exterior — **100%** |
| 4 | Pago o acreditación por compra de bienes/servicios — **50%** |
| 7 | Pago a comercio afiliado a sistema de TC/TD — **50%** |
| 8 | Otros (disminución de la retención) |

## Fórmula

```
montoRetencion = tasaITBMS × totalITBMS
```

## Nodo `retencion` requerido

```xml
<ser:retencion>
  <ser:codigoRetencion>3</ser:codigoRetencion>
  <ser:montoRetencion>7.00</ser:montoRetencion>
</ser:retencion>
```

## Ejemplo de XML

Este ejemplo usa cliente Gobierno (tipoClienteFE=03) + retención código 3 (no domiciliado 100%).

```xml
<soapenv:Body>
  <tem:Enviar>
    <tem:tokenEmpresa></tem:tokenEmpresa>
    <tem:tokenPassword></tem:tokenPassword>
    <tem:documento>
      <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
      <ser:tipoSucursal>1</ser:tipoSucursal>
      <ser:datosTransaccion>
        <ser:tipoEmision>01</ser:tipoEmision>
        <ser:tipoDocumento>01</ser:tipoDocumento>
        <ser:numeroDocumentoFiscal>31040001</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>001</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2022-10-18T08:28:28-05:00</ser:fechaEmision>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>3</ser:formatoCAFE>
        <ser:entregaCAFE>3</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:informacionInteres>Factura interna</ser:informacionInteres>
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
          <ser:descripcion>SISTEMAS DE CABLES</ser:descripcion>
          <ser:unidadMedida>um</ser:unidadMedida>
          <ser:cantidad>1.000000</ser:cantidad>
          <ser:codigoCPBSAbrev>72</ser:codigoCPBSAbrev>
          <ser:codigoCPBS>7210</ser:codigoCPBS>
          <ser:precioUnitario>100.000000</ser:precioUnitario>
          <ser:precioItem>100.000000</ser:precioItem>
          <ser:valorTotal>107.000000</ser:valorTotal>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>7.00</ser:valorITBMS>
        </ser:item>
      </ser:listaItems>
      <ser:totalesSubTotales>
        <ser:totalPrecioNeto>100.00</ser:totalPrecioNeto>
        <ser:totalITBMS>7.00</ser:totalITBMS>
        <ser:totalMontoGravado>7.00</ser:totalMontoGravado>
        <ser:totalFactura>107.00</ser:totalFactura>
        <ser:totalValorRecibido>107.00</ser:totalValorRecibido>
        <ser:vuelto>0.00</ser:vuelto>
        <ser:tiempoPago>1</ser:tiempoPago>
        <ser:nroItems>1</ser:nroItems>
        <ser:totalTodosItems>107.00</ser:totalTodosItems>
        <ser:listaFormaPago>
          <ser:formaPago>
            <ser:formaPagoFact>02</ser:formaPagoFact>
            <ser:valorCuotaPagada>107.00</ser:valorCuotaPagada>
          </ser:formaPago>
        </ser:listaFormaPago>
        <ser:retencion>
          <ser:codigoRetencion>3</ser:codigoRetencion>
          <ser:montoRetencion>7.00</ser:montoRetencion>
        </ser:retencion>
      </ser:totalesSubTotales>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```
