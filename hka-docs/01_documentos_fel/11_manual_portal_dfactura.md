# Manual de Usuario — Portal Dfactura

> Fuente: https://felwiki.thefactoryhka.com.pa/manual_de_usuario_portal_dfactura

## Descripción General Del Sistema

**Dfactura** es una plataforma con capacidad de emitir documentos bajo el esquema de facturación electrónica desde un **portal web**.

### Funcionalidades

- Acceso seguro con email, contraseña y validación **reCAPTCHA**
- Manejo de catálogos de clientes, productos, sucursales y puntos de facturación con capacidad de visualizar, añadir, editar y eliminar registros configurables
- Emisión de documentos electrónicos como Factura de Operación Interna, Nota de Débito Genérica, Nota de Crédito Genérica, Factura de Zona Franca y Reembolso, con capacidad de editar cliente, conceptos e información adicional

La emisión directa desde el portal está diseñada para **pequeñas y medianas empresas o profesionales independientes con un bajo volumen de operaciones**. El portal es la herramienta de configuración de todos los medios que interactúan con la plataforma (Aplicación móvil o ERP).

---

## Ingreso

Para ingresar al Portal Dfactura, colocar el **correo electrónico y contraseña** suministrados por Soporte para iniciar sesión.

## Configuración

Ir a **Configuraciones → Mi Cuenta**.

### Mi cuenta
Editar los datos de la cuenta.

### Ingreso de Datos Fiscales
Presionando el botón **Editar** se encuentran los campos de la información fiscal. Una vez completada la información presionar **Actualizar información**.

### Cambio de Contraseña
Ir a **Seguridad** → solicitará contraseña actual y la nueva con su confirmación → **Cambiar Contraseña**.

### Olvido de Contraseña
En el inicio de sesión pulsar **"¿Olvidó su contraseña?"** → ingresar correo electrónico → **Restablecer contraseña** y validar el reCAPTCHA.

---

## Estructura de Navegación

### Información de la Cuenta
Muestra fecha actual, nombre del usuario, licencia activa, información de folios (totales y utilizados).

### Estadísticas y reportes de comprobantes
Reportes por año o por mes, cantidad de documentos, facturas, documentos anulados, notas de crédito y notas de débito emitidos.

### Información de monto de comprobantes
Gráfica de montos emitidos por comprobante sobre base de tiempo mensual.

### Estadística de Catálogos
Total de cada uno de los catálogos.

---

## Catálogos

Clientes, productos, sucursales, puntos de facturación, dispositivos, usuarios, transportistas y vehículos son manejados mediante catálogos.

### Clientes

**Agregar**: `Catálogos → Clientes → Agregar Cliente` → seleccionar tipo de cliente → completar campos (documento de identidad cumpliendo formato válido) → **Guardar**.

**Consumidor final genérico**: En tipo de cliente seleccionar **"Consumidor final"** → tipo de contribuyente **"Persona Natural"** → datos básicos.

> NOTA: Este método permite emitir un documento sin dar información de empresa particular, solo datos básicos de un consumidor.

**Editar**: Icono **Editar** en columna Acciones.
**Eliminar**: Icono **Eliminar** o casilla de selección para varios/todos.
**Carga Masiva**: Botón **Carga Masiva** → archivo CSV o TXT. Descargar ejemplos disponibles.

### Sucursales

**Agregar**: `Catálogos → Sucursales → Agregar Sucursal` → código, nombre, ubicación, dirección, correo y teléfonos → **Guardar**.
**Editar**: Icono **Editar** en columna Acciones.

### Punto de Facturación

**Agregar**: `Catálogos → Punto de Facturación → Agregar Punto de Facturación` → código, nombre, estatus, sucursal asociada → **Guardar**.
**Editar**: Icono **Editar** en columna Acciones.

### Dispositivos

Visualizan todos los dispositivos asignados al usuario (id, nombre, marca, modelo, serial, tipo). Los dispositivos pueden asociarse y desasociarse a las licencias.

- **Asociar**: seleccionar licencia → botón **Asignar**
- **Desasociar**: botón **Desasociar**
- Los dispositivos son asignados por el personal de soporte

### Usuarios

**Agregar Usuario**: Seleccionar tipo de rol → completar datos personales (razón social, tipo de documento, número de identificación, correo, estatus) → **Siguiente** → sección de **Permisos** → **Guardar**.

**Editar Usuario**: Icono **Editar** en columna Acciones para información.
**Editar Permisos**: Otro icono **Editar** para permisos específicos.

### Transportistas

**Agregar Transportista**: Tipo contribuyente, Razón Social, Número de identificación y Dígito verificador → **Guardar**.
**Editar Transportista**: Icono **Editar**.

### Vehículos

**Agregar Vehículo**: Tipo de vehículo, Chasis, Potencia del motor, Capacidad del motor, Peso neto, Peso bruto, Capacidad de tracción, Distancia entre ejes, Tipo de pintura, Tipo de combustible → **Guardar**.

**Uso del Vehículo**: En "Generar Documentos" → seleccionar ítem → activar **"Activar campos para vehículos"** → seleccionar el vehículo del catálogo → algunos campos se autocompletarán.

---

## Documentos Emitidos

Ir a `Emisión → Documentos Emitidos` para visualizar todas las transacciones.

### Visualización de documentos

**Filtros**: Tipo de documento, Número del documento, Sucursal, RUC del cliente, Estado de documento, Razón Social del cliente.

**Información mostrada**: Tipo de documento, Número, Sucursal, RUC receptor, Razón social receptor, Monto total, Estatus, Descarga XML, Descarga PDF, Envío de correo, Historial de correo, Anulación.

### Acciones disponibles

- **Envío de correo**: Reenviar XML+PDF al cliente
- **Historial de correo**: Consultar CUFE, email receptor, estatus del correo y fecha de envío
- **Anulación**: Colocar motivo → Anular documento autorizado

### Notas de Crédito/Débito Referenciadas

Al seleccionar el número de documento de cualquier documento autorizado, se puede generar una **nota de crédito o débito referenciada** con fecha de emisión y CUFE del documento original.

### Generar Reporte

Botón **Generar Reporte** → descarga archivo **.xlsx** y **PDF** de todos los documentos emitidos. Opción de **Ocultar Columnas**.

### Generar Documento

Ir a `Emisión → Generar Documento`.

> **NOTA**: Debe tener clientes, productos y puntos de facturación registrados en los catálogos para poder generar documentos.

**Pasos**:
1. Seleccionar cliente del catálogo
2. Seleccionar sucursal y punto de facturación asociado
3. Ingresar datos del comprobante
4. Seleccionar almacén asociado a la sucursal
5. Agregar productos
6. Verificar cantidad del producto (despliega cargos de impuestos y CPBS)
7. Ingresar valor del descuento (si aplica)
8. Mencionar impuesto OTI (si aplica)
9. Especificar tipo de retención (si aplica)
10. Mencionar acarreo en el ítem o en el total
11. Activar campos de Medicina (si aplica)
12. Activar campos de Vehículo (si aplica)
13. Activar información del pedido comercial
14. Activar información de logística
15. Activar información de autorizados a descargar la FE y eventos
16. Seleccionar plazo de pagos y agregar más cuotas (si aplica)
17. Seleccionar forma de pago y agregar más formas para pago por cuotas
18. Seleccionar información de exportación (si aplica)
19. **Generar documento**

Una vez generado, el cliente recibe en su correo el XML y el PDF.

### Anular Documento

Ir a `Emisión → Anular Documento`:
1. Sucursal de la FE
2. Tipo de documento a anular
3. Número del documento
4. Tipo de emisión
5. Punto de facturación
6. Motivo de anulación
7. **Enviar anulación**

### Plantillas Documentos

`Emisión → Plantillas Documentos`:

**Agregar**: Nombre, Tipo de Documento, Sucursal, Punto de Facturación, Naturaleza de Operación, Tipo de Venta, Tipo de Sucursal, Entrega CAFE, Formato CAFÉ, Información de Interés, Envío FE, Tipo de Operación, Forma de Pago → **Guardar Plantilla**.

**Uso**: Al generar documentos, seleccionar la opción de **Plantilla** → los campos Tipo de Documento, Sucursal, Punto de Facturación y Numero del Documento se autocompletarán.

---

## Inventario

### Productos

**Agregar Producto**: Tipo de producto, código, descripción, unidad de medida, precio unitario, tipo de impuesto.

> Los campos deben cumplir con las especificaciones de formato del Manual de Integración del WS.

**Editar**: Icono **Editar** en Acciones.
**Eliminar**: Icono **Eliminar** o casilla de selección.
**Carga Masiva**: Botón **Carga Masiva** → archivo CSV o TXT.

### Almacén

**Agregar Almacén**: Código, nombre, teléfono, ubicación, descripción → Asignar sucursal → Seleccionar productos.

**Editar Almacén**: Icono **Editar** en Acciones.

**Detalles del Almacén**: Visualiza productos agregados, busca producto específico, exporta en **.xlsx**, permite editar información.

**Generar reporte de movimientos**: Búsqueda con filtros → descarga **.xlsx** de todos los movimientos de los almacenes.

**Movimientos de Almacén**: Seleccionar tipo (**entrada o salida**) → sucursal y almacén → productos y cantidad.

---

## Licencias

Sección donde se visualizan:
- Licencias asignadas al usuario
- Cantidad de meses de vigencia y cadencia
- Cantidad de folios
- **Ciclos**: al seleccionar icono de **Ciclos** se muestran los ciclos de cada etapa de la licencia con los folios consumidos en cada uno

**Reglas**:
- Las licencias tienen **vigencia anual**
- Los ciclos son de **consumo mensual**
