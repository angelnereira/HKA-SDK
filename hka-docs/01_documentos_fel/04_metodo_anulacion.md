# Método Anulación()

> Fuente: https://felwiki.thefactoryhka.com.pa/anulacion

Permite al usuario **anular documentos** generados en los cuales no se haya realizado la operación registrada en una FE que haya obtenido autorización de uso.

> ⚠️ **Nota técnica importante**: El nombre real del método SOAP es `AnulacionDocumento`, aunque en la documentación aparezca referenciado como "Anulacion". Esta discrepancia es común en la wiki — verificar el método exacto usando la colección Postman oficial.

## REQUEST: Parámetros a Enviar

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | TokenEmpresa / TokenPassword | Proporcionado por The Factory HKA |
| String | motivoAnulacion | Motivo de la Anulación de FE |
| Objeto | DatosDocumento | Ver detalle abajo |

### DatosDocumento

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|4 | codigoSucursalEmisor | 0000: casa matriz. 0001+: otras sucursales |
| String | SI | N\|10 | numeroDocumentoFiscal | Número del documento fiscal |
| String | SI | N\|3 | puntoFacturacionFiscal | Punto de Facturación |
| String | SI | N\|2 | tipoDocumento | 01-09 (ver catálogo) |
| String | SI | N\|2 | tipoEmision | 01-04 (ver catálogo) |

## Ejemplo XML — REQUEST

```xml
<soapenv:Body>
  <tem:AnulacionDocumento>
    <tem:tokenEmpresa>-SOLICITAR-</tem:tokenEmpresa>
    <tem:tokenPassword>-SOLICITAR-</tem:tokenPassword>
    <tem:motivoAnulacion>Motivo de la Anulacion</tem:motivoAnulacion>
    <tem:datosDocumento>
      <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
      <ser:numeroDocumentoFiscal>1212</ser:numeroDocumentoFiscal>
      <ser:puntoFacturacionFiscal>999</ser:puntoFacturacionFiscal>
      <ser:tipoDocumento>01</ser:tipoDocumento>
      <ser:tipoEmision>01</ser:tipoEmision>
    </tem:datosDocumento>
  </tem:AnulacionDocumento>
</soapenv:Body>
```

## RESPONSE: Parámetros a Recibir

| Tipo | Identificador | Descripción |
|---|---|---|
| String | codigo | Código del resultado |
| String | estatusDocumento | Estatus del documento |
| String | resultado | Resultado de la anulación |
| String | mensaje | Mensaje adicional |

## Ejemplo XML — RESPONSE

```xml
<s:Body>
  <AnulacionDocumentoResponse xmlns="http://tempuri.org/">
    <AnulacionDocumentoResult>
      <a:codigo>200</a:codigo>
      <a:resultado>procesado</a:resultado>
      <a:mensaje>Proceso de Anulación ejecutado con éxito.</a:mensaje>
    </AnulacionDocumentoResult>
  </AnulacionDocumentoResponse>
</s:Body>
```
