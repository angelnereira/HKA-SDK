# Método EnvioCorreo()

> Fuente: https://felwiki.thefactoryhka.com.pa/enviocorreo

Permite al usuario **enviar el documento electrónico (XML + PDF) por correo** al destinatario indicado.

## REQUEST: Parámetros a Enviar

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | TokenEmpresa / TokenPassword | Proporcionado por The Factory HKA |
| Objeto | DatosDocumento | Ver detalles abajo |
| String | correo | Correo electrónico destino. Ej: `nombre@dominio.com` |

### DatosDocumento

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|2 | codigoSucursalEmisor | 0000: casa matriz |
| String | SI | N\|10 | numeroDocumentoFiscal | Número del documento fiscal |
| String | SI | N\|3 | puntoFacturacionFiscal | Punto de Facturación |
| String | SI | N\|2 | tipoDocumento | 01-09 |
| String | SI | N\|2 | tipoEmision | 01-04 |

## Ejemplo XML — REQUEST

```xml
<soapenv:Body>
  <tem:EnvioCorreo>
    <tem:tokenEmpresa>SOLICITAR</tem:tokenEmpresa>
    <tem:tokenPassword>SOLICITAR</tem:tokenPassword>
    <tem:datosDocumento>
      <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
      <ser:numeroDocumentoFiscal>0001</ser:numeroDocumentoFiscal>
      <ser:puntoFacturacionFiscal>001</ser:puntoFacturacionFiscal>
      <ser:tipoDocumento>01</ser:tipoDocumento>
      <ser:tipoEmision>01</ser:tipoEmision>
    </tem:datosDocumento>
    <tem:correo>nombre@dominio.com</tem:correo>
  </tem:EnvioCorreo>
</soapenv:Body>
```

## RESPONSE

| Tipo | Identificador | Descripción |
|---|---|---|
| String | codigo | Código del resultado |
| String | resultado | Resultado de operación |
| String | mensaje | Mensaje del resultado |

## Ejemplo XML — RESPONSE

```xml
<s:Body>
  <EnvioCorreoResponse xmlns="http://tempuri.org/">
    <EnvioCorreoResult>
      <a:codigo>200</a:codigo>
      <a:resultado>procesado</a:resultado>
      <a:mensaje>Se ha enviado satisfactoriamente el email.</a:mensaje>
    </EnvioCorreoResult>
  </EnvioCorreoResponse>
</s:Body>
```
