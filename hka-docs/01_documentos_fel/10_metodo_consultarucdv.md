# Método ConsultarRucDV()

> Fuente: https://felwiki.thefactoryhka.com.pa/consultarucdv

Permite al usuario **consultar el dígito verificador (DV) del RUC** y validar información básica del contribuyente.

## REQUEST: Parámetros a Enviar

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | TokenEmpresa / TokenPassword | Suministrado por el proveedor tecnológico |
| String | tipoRuc | Tipo de Contribuyente: **1: Natural. 2: Jurídico** |
| String | ruc | RUC del Contribuyente |

## Ejemplo XML — REQUEST

```xml
<tem:ConsultarRucDV>
  <tem:consultarRucDVRequest>
    <ser:tokenEmpresa><!--SOLICITAR--></ser:tokenEmpresa>
    <ser:tokenPassword><!--SOLICITAR--></ser:tokenPassword>
    <ser:tipoRuc>Tipo de contribuyente</ser:tipoRuc>
    <ser:ruc>Ruc del contribuyente</ser:ruc>
  </tem:consultarRucDVRequest>
</tem:ConsultarRucDV>
```

## RESPONSE

| Tipo | Identificador | Descripción |
|---|---|---|
| String | codigo | Código del resultado |
| String | tipoRuc | 1: Natural. 2: Jurídico |
| String | ruc | RUC del Contribuyente |
| String | dv | **DV calculado** del Contribuyente |
| String | razonSocial | Razón social del Contribuyente |
| String | afiliadoFE | **Estado de afiliación a la FE** |
| String | mensaje | Mensaje adicional |
| String | resultado | Resultado de la operación |

## Ejemplo XML — RESPONSE

```xml
<ConsultarRucDVResponse>
  <ConsultarRucDVResult>
    <a:codigo>200</a:codigo>
    <a:infoRuc>
      <a:tipoRuc>Tipo de contribuyente</a:tipoRuc>
      <a:ruc>Ruc del contribuyente</a:ruc>
      <a:dv>DV del contribuyente</a:dv>
      <a:razonSocial>Razón social del contribuyente</a:razonSocial>
      <a:afiliadoFE>Mensaje de estado de afiliación a la FE</a:afiliadoFE>
    </a:infoRuc>
    <a:mensaje>Mensaje adicional de la operación</a:mensaje>
    <a:resultado>Procesado</a:resultado>
  </ConsultarRucDVResult>
</ConsultarRucDVResponse>
```

## Casos de uso

- **Validación en onboarding de clientes**: Al registrar un nuevo cliente B2B, consultar el DV y la razón social para auto-completar/verificar datos
- **Pre-validación antes de emitir factura**: Evitar errores de DV inválido en `Enviar()` que causarían rechazo
- **Confirmación de afiliación a FE**: Antes de enviar facturas a un cliente contribuyente, validar si está afiliado al sistema de FE
