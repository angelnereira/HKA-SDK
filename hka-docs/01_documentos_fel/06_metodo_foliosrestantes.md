# Método FoliosRestantes()

> Fuente: https://felwiki.thefactoryhka.com.pa/foliosrestantes

Permite al usuario consultar los **folios disponibles** de la licencia para su uso.

## REQUEST: Parámetros a Enviar

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | TokenEmpresa / TokenPassword | Proporcionado por The Factory HKA |

## Ejemplo XML — REQUEST

```xml
<FoliosRestantes>
  <tokenEmpresa>SOLICITAR</tokenEmpresa>
  <tokenPassword>SOLICITAR</tokenPassword>
</FoliosRestantes>
```

## RESPONSE

| Tipo | Identificador | Descripción |
|---|---|---|
| String | codigo | Código del resultado |
| String | mensaje | Mensaje adicional |
| String | licencia | Código de la licencia asignada |
| String | fechaLicencia | Fecha de vigencia (formato `YYYY-MM-DD / YYYY-MM-DD`) |
| String | ciclo | Número del ciclo de la licencia |
| String | fechaCiclo | Fecha del ciclo |
| String | foliosTotalesCiclo | Cantidad total de folios del ciclo |
| String | foliosUtilizadosCiclo | Cantidad utilizada de folios del ciclo |
| String | foliosDisponiblesCiclo | Cantidad disponible de folios del ciclo |
| String | foliosTotales | Cantidad total de folios de la licencia |
| String | resultado | Resultado de la operación |

## Ejemplo XML — RESPONSE

```xml
<FoliosRestantesResponse>
  <FoliosRestantesResult>
    <a:codigo>200</a:codigo>
    <a:mensaje>Se retorna Licencia y Disponibilidad</a:mensaje>
    <a:licencia>URFT-F1E3-VSGY</a:licencia>
    <a:fechaLicencia>2020-08-12 / 2021-08-12</a:fechaLicencia>
    <a:ciclo>2</a:ciclo>
    <a:fechaCiclo>2020-09-12 / 2020-10-12</a:fechaCiclo>
    <a:foliosTotalesCiclo>5000</a:foliosTotalesCiclo>
    <a:foliosUtilizadosCiclo>0</a:foliosUtilizadosCiclo>
    <a:foliosDisponibleCiclo>5000</a:foliosDisponibleCiclo>
    <a:foliosTotales>60000</a:foliosTotales>
    <a:resultado>procesado</a:resultado>
  </FoliosRestantesResult>
</FoliosRestantesResponse>
```

## Notas de implementación

- Las licencias tienen **vigencia anual**
- Los ciclos son de **consumo mensual**
- `fechaLicencia` y `fechaCiclo` vienen con formato concatenado: `"YYYY-MM-DD / YYYY-MM-DD"`
- Útil para monitorear consumo y alertar antes de agotar folios
