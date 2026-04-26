# Método EstadoDocumento()

> Fuente: https://felwiki.thefactoryhka.com.pa/estadodocumento

Permite al usuario consultar el **estatus de un documento** indicando su número.

## REQUEST: Parámetros a Enviar

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | TokenEmpresa / TokenPassword | Proporcionado por The Factory HKA |
| Objeto | DocumentoElectronico | Ver DatosDocumento |

### DatosDocumento

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|2 | codigoSucursalEmisor | 0000: casa matriz. 0001 en delante: otras sucursales |
| String | SI | N\|10 | numeroDocumentoFiscal | Número del documento fiscal (llenar con ceros a la izquierda) |
| String | SI | N\|3 | puntoFacturacionFiscal | Punto de Facturación (no se admite "cero") |
| String | SI | N\|2 | tipoDocumento | 01-09 (ver catálogo) |
| String | SI | N\|2 | tipoEmision | 01-04 (ver catálogo) |

## Ejemplo XML — REQUEST

```xml
<tem:EstadoDocumento>
  <tem:tokenEmpresa>SOLICITAR</tem:tokenEmpresa>
  <tem:tokenPassword>SOLICITAR</tem:tokenPassword>
  <tem:datosDocumento>
    <ser:codigoSucursalEmisor>0001</ser:codigoSucursalEmisor>
    <ser:numeroDocumentoFiscal>99176</ser:numeroDocumentoFiscal>
    <ser:puntoFacturacionFiscal>565</ser:puntoFacturacionFiscal>
    <ser:tipoDocumento>01</ser:tipoDocumento>
    <ser:tipoEmision>01</ser:tipoEmision>
  </tem:datosDocumento>
</tem:EstadoDocumento>
```

## RESPONSE: Parámetros a Recibir

| Tipo | Identificador | Descripción |
|---|---|---|
| String | codigo | Código del resultado |
| String | mensaje | Mensaje de retorno |
| String | fechaEmisionDocumento | Fecha de emisión |
| String | fechaRecepcionDocumento | Fecha de Recepción |
| String | estatusDocumento | Estado del documento |
| String | mensajeDocumento | Mensaje del estado |
| String | resultado | Mensaje de retorno |

## Ejemplo XML — RESPONSE

```xml
<EstadoDocumentoResponse xmlns="http://tempuri.org/">
  <EstadoDocumentoResult>
    <a:codigo>200</a:codigo>
    <a:mensaje>Se retornan datos del documento.</a:mensaje>
    <a:cufe>FE01100000000000-000-0000-4000002023032000000010992260127062273602</a:cufe>
    <a:fechaEmisionDocumento>2023-04-10 10:29:09</a:fechaEmisionDocumento>
    <a:fechaRecepcionDocumento>2023-04-10 10:29:07</a:fechaRecepcionDocumento>
    <a:estatusDocumento>Autorizada</a:estatusDocumento>
    <a:mensajeDocumento>Autorizado el uso de la FE</a:mensajeDocumento>
    <a:resultado>procesado</a:resultado>
  </EstadoDocumentoResult>
</EstadoDocumentoResponse>
```
