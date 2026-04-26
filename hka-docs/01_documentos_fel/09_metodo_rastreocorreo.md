# Método RastreoCorreo()

> Fuente: https://felwiki.thefactoryhka.com.pa/rastreocorreo

Permite al usuario **consultar la traza de correos enviados** de los documentos generados.

> ⚠️ **Nota técnica**: Este método identifica el documento por **CUFE**, no por los parámetros estándar de documento (codigoSucursalEmisor / numeroDocumentoFiscal / etc). Esto es **distinto** al resto de los métodos auxiliares.

## REQUEST: Parámetros a Enviar

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | TokenEmpresa / TokenPassword | Suministrado por el proveedor tecnológico |
| String | cufe | **Código Único de Factura Electrónica** |

## Ejemplo XML — REQUEST

```xml
<RastreoCorreo>
  <tokenEmpresa><!--SOLICITAR--></tokenEmpresa>
  <tokenPassword><!--SOLICITAR--></tokenPassword>
  <tem:cufe><!--CUFE--></tem:cufe>
</RastreoCorreo>
```

## RESPONSE

| Tipo | Identificador | Descripción |
|---|---|---|
| String | codigo | Código del resultado |
| String | mensaje | Mensaje adicional |
| String | listaRastreo | Lista con: correo, fecha de creación, estado del correo e Id del mensaje |
| String | resultado | Resultado del envío |

### Campos dentro de `listaRastreo > listTracking`

| Campo | Descripción |
|---|---|
| correo | Dirección de email destino |
| creado_en | Fecha de creación del correo |
| estado | Estado del correo (delivered, bounced, etc.) |
| messageId | ID único del mensaje (para tracking con proveedor SMTP) |

## Ejemplo XML — RESPONSE

```xml
<RastreoCorreoResponse>
  <RastreoCorreoResult>
    <a:codigo>200</a:codigo>
    <a:mensaje>Se retorna mensaje de la consulta</a:mensaje>
    <a:listaRastreo>
      <b:listTracking>
        <b:correo>Correo</b:correo>
        <b:creado_en>Fecha de creación del correo</b:creado_en>
        <b:estado>Estado del correo</b:estado>
        <b:messageId>Id del mensaje del correo</b:messageId>
      </b:listTracking>
    </a:listaRastreo>
    <a:resultado>procesado</a:resultado>
  </RastreoCorreoResult>
</RastreoCorreoResponse>
```
