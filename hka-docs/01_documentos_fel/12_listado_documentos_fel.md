# Listado de Documentos FEL

> Fuente: https://felwiki.thefactoryhka.com.pa/listado_de_documentos_fel

Listado de los **documentos mínimos requeridos** para toda empresa emisora de facturas electrónicas. Esta lista representa el **set de pruebas funcionales** que una empresa debe poder emitir de forma exitosa antes de pasar a producción.

## Lista obligatoria

1. **Factura de operación interna**
2. **Nota de débito genérica y referente a una FE**
3. **Nota de crédito genérica y referente a una FE**
4. **Documentos electrónicos con productos de distintas tasas ITBMS**
5. **Documentos electrónicos a precio 0.00**
6. **Documentos electrónicos con descuento unitario por item**
7. **Documentos electrónicos con descuentos por factura**
8. **Documentos electrónicos a cliente Contribuyente, Consumidor final, Gobierno, Extranjero**
9. **Documentos electrónicos con productos de precios usando 2 a 6 decimales** (validar redondeos)

> ⚠️ **NOTA**: Sujeto a cambios.

---

## Implicaciones para integración

Estos 9 casos definen un **conjunto mínimo de pruebas de integración** que se debe validar:

### Test plan sugerido para Sago One

| # | Caso | Endpoint Sago One | Validación |
|---|---|---|---|
| 1 | Factura operación interna básica | POST /invoice | CUFE válido, código 200 |
| 2 | Nota de débito referenciada | POST /credit-debit-note | Debe incluir `cufeFEReferenciada` |
| 3 | Nota de crédito referenciada | POST /credit-debit-note | Debe incluir `cufeFEReferenciada` |
| 4 | Nota de débito genérica | POST /credit-debit-note | Sin CUFE referenciado |
| 5 | Nota de crédito genérica | POST /credit-debit-note | Sin CUFE referenciado |
| 6 | Factura con múltiples tasas ITBMS (00/01/02/03) | POST /invoice | Desglose correcto en totales |
| 7 | Factura con ítems precio 0.00 | POST /invoice | `valorTotal` e `ITBMS` igual a 0.00 |
| 8 | Factura con descuento por ítem | POST /invoice | `precioUnitarioDescuento` en Balboas |
| 9 | Factura con descuento por factura | POST /invoice | `totalDescuento` > 0 |
| 10 | Factura a Contribuyente (tipoClienteFE=01) | POST /invoice | RUC + DV validados |
| 11 | Factura a Consumidor final (tipoClienteFE=02) | POST /invoice | Solo cédula opcional |
| 12 | Factura a Gobierno (tipoClienteFE=03) | POST /invoice | **codigoCPBS obligatorio** |
| 13 | Factura a Extranjero (tipoClienteFE=04) | POST /invoice | Sin RUC/DV/ubicación |
| 14 | Factura con 6 decimales en precio | POST /invoice | Redondeo correcto en totales |
