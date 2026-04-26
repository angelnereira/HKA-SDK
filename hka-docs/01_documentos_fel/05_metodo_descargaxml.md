# Método DescargaXML()

> Fuente: https://felwiki.thefactoryhka.com.pa/descargaxml

Permite al usuario **descargar un documento electrónico generado, en formato XML**.

## REQUEST: Parámetros a Enviar

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | TokenEmpresa / TokenPassword | Proporcionado por The Factory HKA |
| Objeto | DatosDocumento | Ver detalle abajo |

### DatosDocumento

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|4 | codigoSucursalEmisor | 0000: casa matriz |
| String | SI | N\|10 | numeroDocumentoFiscal | Número del documento fiscal |
| String | SI | N\|3 | puntoFacturacionFiscal | Punto de Facturación |
| String | SI | N\|2 | tipoDocumento | 01-09 |
| String | SI | N\|2 | tipoEmision | 01-04 |

## Ejemplo XML — REQUEST

```xml
<DescargaXML>
  <tokenEmpresa><!--SOLICITAR--></tokenEmpresa>
  <tokenPassword><!--SOLICITAR--></tokenPassword>
  <datosDocumento>
    <codigoSucursalEmisor>0001</codigoSucursalEmisor>
    <numeroDocumentoFiscal>891</numeroDocumentoFiscal>
    <puntoFacturacionFiscal>001</puntoFacturacionFiscal>
    <tipoDocumento>01</tipoDocumento>
    <ser:tipoEmision>01</ser:tipoEmision>
  </datosDocumento>
</DescargaXML>
```

## RESPONSE

| Tipo | Identificador | Descripción |
|---|---|---|
| String | codigo | Código del resultado |
| String | resultado | Resultado de operación |
| String | mensaje | Mensaje adicional |
| String | documento | **Documento descargado en formato Base64** |

## Ejemplo XML — RESPONSE

```xml
<DescargaXMLResponse>
  <DescargaXMLResult>
    <codigo>200</codigo>
    <resultado>procesado</resultado>
    <a:mensaje>Se retorna el archivo solicitado.</a:mensaje>
    <documento> [XML EN BASE 64] </documento>
  </DescargaXMLResult>
</DescargaXMLResponse>
```
