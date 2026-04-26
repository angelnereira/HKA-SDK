# Preguntas Frecuentes

> Fuente: https://felwiki.thefactoryhka.com.pa/preguntas_frecuentes

## Conceptos fundamentales

### ¿Qué es el CAFE?

Por sus siglas: **Comprobante Auxiliar de Factura Electrónica**. Es el **PDF** que se entrega cuando el documento es autorizado.

### ¿Qué es el CUFE?

Por sus siglas: **Código Único de Factura Electrónica**. Es el código que identifica a cada documento autorizado.

### ¿Cuál es el protocolo que manejan, SOAP o REST?

Se maneja protocolo **SOAP**.

---

## Integración

### ¿Cuánto tiempo toma la integración?

Varía dependiendo de quién hace la integración (interno de la empresa o un tercero) y el tiempo dedicado al proyecto. Los ejemplos de código permiten emitir documentos **inmediatamente**.

### ¿Los datos creados en el ambiente demo son válidos para producción?

**No**. Los datos en ambientes de pruebas (sucursales, clientes, puntos de facturación) no son válidos para producción. Una vez se cambie de ambiente se debe **volver a cargar la data**.

### Si tengo varias empresas, ¿puedo usar las mismas credenciales?

**No**, las credenciales son **únicas por cada RUC**.

---

## Portal Dfactura

### ¿Existe límite de usuarios para la plataforma Dfactura?

**No**, no hay límite de usuarios para crear en Dfactura.

### ¿Admite carga masiva de productos y clientes?

**Sí**, consultar el Manual de Usuario Portal Dfactura.

### ¿Se pueden generar reportes desde el portal Dfactura?

**Sí**, en la sección de Documentos Emitidos se permite filtrar por fecha, RUC, razón social, tipo de documento, etc.

### ¿Se puede agregar el logo al CAFE?

**Sí**, se puede agregar desde el portal Dfactura.

### ¿Se puede acceder al portal Dfactura de manera simultánea?

**No**, cada sesión es única. No se puede tener un usuario conectado desde distintos equipos.

---

## Reglas fiscales

### ¿Puedo declarar un RUC jurídico como consumidor final?

**No**. Para cada tipo de cliente hay campos requeridos y opcionales. Si es **RUC jurídico, debe declararse obligatoriamente como contribuyente**. Declararlo como consumidor final generará **rechazo del servicio**.

### ¿Cuál es la diferencia entre factura de operación interna y de exportación?

- **Factura de operación interna (01)**: Documento regular dentro de Panamá ("factura" estándar).
- **Factura de exportación (03)**: Cuando el producto o servicio va hacia un país destino distinto de Panamá.

### ¿Puedo crear clientes exentos de ITBMS en Dfactura?

**No se crean clientes exentos** de impuestos, **sino productos o servicios**.

---

## Plazos y anulaciones

### ¿Cuántos días tengo para realizar una nota de crédito a una factura?

El periodo máximo es de **180 días** (6 meses aproximadamente).

### ¿En qué casos aplica Anulación de documentos?

Se puede anular cuando **no hubo transacción fiscal** en base al producto o servicio — es decir, en casos de **errores** en el documento.

- **Período para anular**: **máximo 7 días continuos**
- Luego de ese período: generará rechazo y se debe emitir **Nota de Crédito**

---

## Contacto de soporte

Para dudas adicionales, enviar correo a:

**📧 soporte_fel_pa@thefactoryhka.com**

---

## Resumen de plazos críticos

| Operación | Plazo máximo |
|---|---|
| Anular factura | 7 días continuos |
| Emitir Nota de Crédito a factura | 180 días (6 meses aprox.) |
| Regresar a operación normal post-contingencia | 72 horas (si dura más, explicar razones) |
