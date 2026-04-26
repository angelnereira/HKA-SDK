# Factura con Otro Tipo de Impuesto (OTI)

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_con_otro_tipo_de_impuesto

Factura que incluye **Otros Tipos de Impuestos (OTI)** además del ITBMS: tasas específicas como SUME 911, Portabilidad Numérica, impuestos sobre seguros, cargos aeroportuarios, FECI, etc.

## Códigos OTI (tasaOTI / codigoTotalOTI)

| Código | Descripción |
|---|---|
| 01 | SUME 911 |
| 02 | Tasa Portabilidad Numérica |
| 03 | Impuesto sobre seguro (5%) |
| 04 | ATTT — Impuestos sobre seguro autos (1%) |
| 05 | Tasa Salida Aeropuerto (FZ) |
| 06 | Cargo Incentivo Desarrollo Aeropuerto (F3) |
| 07 | Cargo Seguridad Aeropuerto (AH) |
| 08 | Otros Cargos (XT) |
| 09 | Combustible (YQ) |
| 10 | FECI |
| 11 | Intereses |

## Estructura OTI en dos niveles

### A nivel de ítem: `listaItemOTI`
Cada ítem lleva los OTI que le aplican.

```xml
<ser:listaItemOTI>
  <ser:oti>
    <ser:tasaOTI>01</ser:tasaOTI>
    <ser:valorTasa>0.01</ser:valorTasa>
  </ser:oti>
</ser:listaItemOTI>
```

### A nivel de factura: `listaTotalOTI`
Sumatoria por código OTI a través de todos los ítems.

```xml
<ser:listaTotalOTI>
  <ser:totalOti>
    <ser:codigoTotalOTI>01</ser:codigoTotalOTI>
    <ser:valorTotalOTI>0.01</ser:valorTotalOTI>
  </ser:totalOti>
</ser:listaTotalOTI>
```

## Fórmula

```
totalMontoGravado = totalITBMS + totalISC + Σ valorTotalOTI
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
        <ser:numeroDocumentoFiscal>0000000901</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>226</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2023-02-10T02:00:00-05:00</ser:fechaEmision>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>1</ser:formatoCAFE>
        <ser:entregaCAFE>1</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:tipoVenta>1</ser:tipoVenta>
        <ser:informacionInteres>Linea1&lt;br/> Linea2</ser:informacionInteres>
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
          <ser:descripcion>Lapiz</ser:descripcion>
          <ser:codigo>TMP-MGT-085/2022</ser:codigo>
          <ser:unidadMedida>m</ser:unidadMedida>
          <ser:cantidad>1.00</ser:cantidad>
          <ser:codigoCPBSAbrev>10</ser:codigoCPBSAbrev>
          <ser:codigoCPBS>1010</ser:codigoCPBS>
          <ser:infoItem>Lapiz con Goma</ser:infoItem>
          <ser:precioUnitario>1.00</ser:precioUnitario>
          <ser:precioItem>1.00</ser:precioItem>
          <ser:valorTotal>1.08</ser:valorTotal>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>0.07</ser:valorITBMS>
          <ser:listaItemOTI>
            <ser:oti>
              <ser:tasaOTI>01</ser:tasaOTI>
              <ser:valorTasa>0.01</ser:valorTasa>
            </ser:oti>
          </ser:listaItemOTI>
        </ser:item>
      </ser:listaItems>
      <ser:totalesSubTotales>
        <ser:totalPrecioNeto>1.00</ser:totalPrecioNeto>
        <ser:totalSC>0.00</ser:totalSC>
        <ser:totalITBMS>0.07</ser:totalITBMS>
        <ser:totalMontoGravado>0.08</ser:totalMontoGravado>
        <ser:totalDescuento>0.00</ser:totalDescuento>
        <ser:totalFactura>1.08</ser:totalFactura>
        <ser:totalValorRecibido>1.08</ser:totalValorRecibido>
        <ser:vuelto>0.00</ser:vuelto>
        <ser:tiempoPago>1</ser:tiempoPago>
        <ser:nroItems>1</ser:nroItems>
        <ser:totalTodosItems>1.08</ser:totalTodosItems>
        <ser:listaFormaPago>
          <ser:formaPago>
            <ser:formaPagoFact>02</ser:formaPagoFact>
            <ser:valorCuotaPagada>1.08</ser:valorCuotaPagada>
          </ser:formaPago>
        </ser:listaFormaPago>
        <ser:listaTotalOTI>
          <ser:totalOti>
            <ser:codigoTotalOTI>01</ser:codigoTotalOTI>
            <ser:valorTotalOTI>0.01</ser:valorTotalOTI>
          </ser:totalOti>
        </ser:listaTotalOTI>
      </ser:totalesSubTotales>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```
