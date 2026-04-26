# Factura con Descuento por Ítem

> Fuente: https://felwiki.thefactoryhka.com.pa/factura_con_descuento_por_item

Factura con **descuentos aplicados individualmente a cada ítem** mediante `precioUnitarioDescuento`. Diferente del descuento global (que se resta del total).

## Campos clave distintivos

| Campo | Valor | Observación |
|---|---|---|
| precioUnitarioDescuento | `100.00` | **Monto en Balboas, NO porcentaje** |
| precioItem | `400.00` | Resultado tras aplicar descuento: `cantidad × (precioUnitario - precioUnitarioDescuento)` |
| totalDescuento | `0.00` | **NO se suman** los descuentos por ítem aquí (solo descuentos globales) |

## Fórmula del ítem con descuento

```
precioItem = cantidad × (precioUnitario - precioUnitarioDescuento)
          = 1.00 × (500.00 - 100.00) = 400.00

valorTotal = precioItem + valorITBMS + valorISC + precioAcarreo + precioSeguro
          = 400.00 + 28.00 = 428.00
```

## Diferencia crítica

- **Descuento por ítem**: `precioUnitarioDescuento` reduce el `precioItem` directamente. El ITBMS se calcula sobre el precio ya descontado.
- **Descuento global**: `totalDescuento` se resta al final sin afectar los precios individuales ni el cálculo de impuestos.

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
        <ser:numeroDocumentoFiscal>5000015</ser:numeroDocumentoFiscal>
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
          <ser:razonSocial>Consumidor final</ser:razonSocial>
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
          <ser:precioUnitario>500.00</ser:precioUnitario>
          <ser:precioUnitarioDescuento>100.00</ser:precioUnitarioDescuento>
          <ser:precioItem>400.00</ser:precioItem>
          <ser:valorTotal>428.00</ser:valorTotal>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>28.00</ser:valorITBMS>
        </ser:item>
        <ser:item>
          <ser:descripcion>Pencil</ser:descripcion>
          <ser:codigo>PENC01</ser:codigo>
          <ser:unidadMedida>m</ser:unidadMedida>
          <ser:cantidad>1.00</ser:cantidad>
          <ser:precioUnitario>400.00</ser:precioUnitario>
          <ser:precioItem>400.00</ser:precioItem>
          <ser:valorTotal>428.00</ser:valorTotal>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>28.00</ser:valorITBMS>
        </ser:item>
      </ser:listaItems>
      <ser:totalesSubTotales>
        <ser:totalPrecioNeto>800.00</ser:totalPrecioNeto>
        <ser:totalITBMS>56.00</ser:totalITBMS>
        <ser:totalMontoGravado>56.00</ser:totalMontoGravado>
        <ser:totalDescuento>0.00</ser:totalDescuento>
        <ser:totalFactura>856.00</ser:totalFactura>
        <ser:totalValorRecibido>856.00</ser:totalValorRecibido>
        <ser:tiempoPago>1</ser:tiempoPago>
        <ser:nroItems>2</ser:nroItems>
        <ser:totalTodosItems>856.00</ser:totalTodosItems>
        <ser:listaFormaPago>
          <ser:formaPago>
            <ser:formaPagoFact>02</ser:formaPagoFact>
            <ser:valorCuotaPagada>856.00</ser:valorCuotaPagada>
          </ser:formaPago>
        </ser:listaFormaPago>
      </ser:totalesSubTotales>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```
