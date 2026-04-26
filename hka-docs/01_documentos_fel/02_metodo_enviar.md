# Método Enviar()

> Fuente: https://felwiki.thefactoryhka.com.pa/enviar

La función encargada de **construir y enviar el documento electrónico** al sistema de facturación electrónica de The Factory HKA.

## REQUEST: Parámetros a Enviar

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | TokenEmpresa / TokenPassword | Suministrado por el proveedor tecnológico |
| Objeto | DocumentoElectronico | Ver detalle abajo |

## Documento Electrónico

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|4 | codigoSucursalEmisor | 0000: Casa matriz. 0001 en adelante: Otras sucursales. Puede utilizar números y letras. |

## DatosFactura

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|2 | tipoEmision | 01: Autorización de Uso Previa, operación normal. 02: Autorización de Uso Previa, contingencia. 03: Autorización Posterior, normal. 04: Autorización Posterior, contingencia. |
| String | C/C | AN\|25 | fechaInicioContingencia | AAAA-MM-DD hh:mm:ssTZH. Obligatorio si tipoEmision = 02 / 04 |
| String | C/C | AN\|15..250 | motivoContingencia | Obligatorio si tipoEmision = 02 / 04. Si la contingencia dura más de 72 horas también debe explicar las razones para no haber regresado a la operación normal. |
| String | SI | N\|2 | tipoDocumento | 01: Factura operación interna. 02: Factura importación. 03: Factura exportación. 04: Nota de Crédito referente a FE. 05: Nota de Débito referente a FE. 06: Nota de Crédito genérica. 07: Nota de Débito genérica. 08: Factura Zona Franca. 09: Reembolso. 10: Factura operación extranjera |
| String | SI | N\|10 | numeroDocumentoFiscal | Número del documento fiscal en la serie correspondiente, de 0000000001 a 9999999999, no se permite reinicio de numeración. Llenar con ceros a la izquierda. |
| String | SI | N\|3 | puntoFacturacionFiscal | Punto de Facturación. Permite secuencias independientes de numeración. No se admite el valor "cero". Llenar con ceros a la izquierda. |
| String | SI | AN\|25 | fechaEmision | AAAA-MM-DDThh:mm:ssTZH |
| String | NO | N\|25 | fechaSalida | AAAA-MM-DDThh:mm:ssTZH. Informar cuando sea conocida. |
| String | SI | N\|2 | naturalezaOperacion | 01: Venta. 02: Exportación. 03: Re-exportación. 04: Venta de fuente extranjera. 10: Transferencia/Traspaso. 11: Devolución. 12: Consignación. 13: Remesa. 14: Entrega gratuita. 20: Compra. 21: Importación. |
| String | SI | N\|1 | tipoOperacion | 1: Salida o venta. 2: Entrada o compra (factura de compra — para comercio informal, ej: taxista, trabajadores manuales). |
| String | SI | N\|1 | destinoOperacion | 1: Panamá. 2: Extranjero. |
| String | SI | N\|1 | formatoCAFE | 1: Sin generación de CAFE. 2: Cinta de papel. 3: Papel formato carta. |
| String | SI | N\|1 | entregaCAFE | 1: Sin generación de CAFE. 2: CAFE entregado al receptor en papel. 3: CAFE enviado al receptor en formato electrónico. |
| String | SI | N\|1 | envioContenedor | 1: Normal. 2: El receptor exceptúa al emisor de la obligatoriedad de envío del contenedor. |
| String | SI | N\|1 | procesoGeneracion | 1: Generación por el sistema de facturación del contribuyente. |
| String | SI | N\|1 | tipoVenta | Tipo de Venta: 1: Venta de giro del negocio. 2: Venta Activo Fijo. 3: Venta de Bienes Raíces. 4: Prestación de Servicio. Si no es venta, no informar. |
| String | NO | N\|1 | tipoSucursal | 1: Mayor cantidad de operaciones al detal (retail). 2: Mayor cantidad al por mayor. Si la transacción es B2B y tipoSucursal=01 y dvtot. |
| String | NO | AN\|...5000 | informacionInteres | Información del interés del emisor con respecto a la FE. |

## Cliente

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|2 | tipoClienteFE | 01: Contribuyente. 02: Consumidor final. 03: Gobierno. 04: Extranjero. |
| String | SI | N\|1 | tipoContribuyente | 1: Natural. 2: Jurídico. **No informar si tipoClienteFE = 04** |
| String | C/C | AN\|20 | numeroRUC | RUC del Receptor. Si es 02:Consumidor Final puede llenarse con la cédula si se solicita. **No enviar cuando tipoClienteFE = 04** |
| String | C/C | N\|2 | digitoVerificadorRUC | Calculado por el algoritmo del RUC publicado por la DGI. Obligatorio y válido si tipoClienteFE = 01/03. **No enviar cuando tipoClienteFE = 04** |
| String | C/C | AN\|2..200 | razonSocial | Obligatorio si tipoClienteFE = 01/03 |
| String | C/C | AN\|5..100 | direccion | Urbanización, Calle, Casa/Edificio, Número de Local. Obligatorio si tipoClienteFE = 01/03 |
| String | C/C | AN\|..8 | codigoUbicacion | Código compuesto de la ubicación. **No enviar cuando tipoClienteFE = 04** |
| String | C/C | AN\|..50 | provincia | Según Catálogo unificado. **No enviar cuando tipoClienteFE = 04** |
| String | C/C | AN\|..50 | distrito | Según Catálogo unificado. **No enviar cuando tipoClienteFE = 04** |
| String | C/C | AN\|..50 | corregimiento | Según Catálogo unificado. **No enviar cuando tipoClienteFE = 04** |
| String | C/C | N\|2 | tipoIdentificacion | Obligatorio si tipoClienteFE = 04. 01: Pasaporte. 02: Número Tributario. 99: Otro. |
| String | C/C | AN\|..50 | nroIdentificacionExtranjero | Obligatorio si tipoClienteFE = 04 |
| String | C/C | AN\|..50 | paisExtranjero | Solo si nroIdentificacionExtranjero es pasaporte. Nombre completo del país. |
| String | NO | AN\|7..16 | telefono1 | Formatos: 999-9999, 9999-9999, 999999-999-9999 o 999999-9999-9999 |
| String | NO | AN\|..50 | correoElectronico1 | Correo válido. |
| String | NO | AN\|..50 | correoElectronico2 | Correo válido. |
| String | NO | AN\|..50 | correoElectronico3 | Correo válido. |
| String | SI | AN\|2 | pais | Catálogo unificado. **Debe ser PA si destinoOperacion=1**, no puede ser PA si destinoOperacion=2. **Debe ser ZZ si el país no existe en el catálogo**. |
| String | C/C | AN\|5..50 | paisOtro | Obligatorio si pais = ZZ. Nombre completo del país. |

## datosFacturaExportacion (si destinoOperacion ≠ Panamá)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|3 | condicionesEntrega | Condiciones de entrega según Tabla 1. INCOTERMS |
| String | SI | AN\|3 | monedaOperExportacion | Código de moneda (Tabla 2. ISO 4217) |
| String | C/C | AN\|5..50 | monedaOperExportacionNonDef | Obligatorio si monedaOperExportacion = ZZZ |
| String | C/C | N\|1..11\|2.4 | tipoDeCambio | Para monedas diferentes al USD |
| String | C/C | N\|1..20\|0.4 | montoMonedaExtranjera | Producto de tipoDeCambio por totalFactura |
| String | NO | AN\|5..50 | puertoEmbarque | Puerto de Embarque de la mercancía |

## List<docFiscalReferenciado> (Opcional)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|25 | fechaEmisionDocFiscalReferenciado | Fecha de emisión del documento referenciado |
| String | SI | AN\|66 | cufeFEReferenciada | CUFE de FE referenciada (66 posiciones) |
| String | NO | AN\|22 | nroFacturaPapel | No informar si existe cufeFEReferenciada |
| String | NO | AN\|22 | nroFacturaIF | Número factura emitida por impresora fiscal. No informar si existe cufeFEReferenciada |

## List<autorizadoDescargaFEyEventos>

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|1 | tipoContribuyente | 1: Natural. 2: Jurídico |
| String | SI | AN\|20 | rucReceptor | RUC autorizado a descargar la FE. Ej: empresa que da acceso a su firma de contadores o a un tercero. |
| String | SI | N\|1 | digitoVerifRucReceptor | Según algoritmo del RUC publicado por la DGI |

## List<Item>

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|2..500 | descripcion | Descripción del producto o servicio |
| String | NO | AN\|1..20 | codigo | Código del producto en el catálogo del emisor |
| String | NO | AN\|1..20 | unidadMedida | Según catálogo de unidades (Tabla 3) |
| String | SI | N\|1..9\|2.6 | cantidad | Cantidad del producto en la unidad del código interno |
| String | C/C | AN\|10 | fechaFabricacion | AAAA-MM-DD. Obligatorio si es materia prima farmacéutica, medicina, alimento |
| String | C/C | AN\|10 | fechaCaducidad | AAAA-MM-DD |
| String | NO | AN\|2 | codigoCPBSAbrev | Segmentos y Familias CPBS (Tabla 4). Obligatorio solo cuando tipoCliente = 03 |
| String | C/C | AN\|4 | codigoCPBS | Codificación Panameña de Bienes y Servicios. Obligatorio si es venta a Administración Pública |
| String | NO | AN\|1..30 | unidadMedidaCPBS | Catálogo de unidades (Tabla 3). Obligatorio solo cuando tipoCliente = 03 |
| String | NO | AB\|0..5000 | infoItem | Información del interés del emisor sobre el ítem |
| String | SI | N\|0..11\|2.6 | precioUnitario | Si el ítem no es valorable: 0.00 |
| String | SI | N\|1..9\|2.6 | precioUnitarioDescuento | Descuento en Balboas, no en porcentaje. Menor que precioUnitario |
| String | SI | N\|1..9\|2..6 | precioItem | Si el ítem no es valorable: 0.00. `precioItem = cantidad * (precioUnitario - precioUnitarioDescuento)` |
| String | NO | N\|1..9\|2 | precioAcarreo | Precio del acarreo para este ítem. Si se informa, no se puede informar en el total |
| String | NO | N\|1..9\|2 | precioSeguro | Precio del seguro para este ítem. Si se informa, no se puede informar en el total |
| String | SI | N\|1..9\|2.6 | valorTotal | Suma del precio del ítem con los montos de los impuestos. `precioItem + precioAcarreo + precioSeguro + valorITBMS + valorISC` |
| String | NO | N\|..14 | codigoGTIN | Código GTIN para unidad de comercialización. Vacío si no tiene código |
| String | NO | N\|1..9\|0.6 | cantGTINCom | Cantidad en código GTIN |
| String | NO | N\|..14 | codigoGTINInv | Código GTIN para unidad de inventario |
| String | NO | N\|1..9\|0.6 | cantGTINComInv | Cantidad en código GTIN de inventario |
| String | SI | N\|2 | tasaITBMS | **00: 0% (exento). 01: 7%. 02: 10%. 03: 15%** |
| String | SI | N\|1..9\|2.6 | valorITBMS | `tasaITBMS * precioItem` |
| String | C/C | N\|1..11\|2.4 | tasaISC | Tasa ISC porcentual (según legislación vigente) |
| String | C/C | N\|1..9\|2.6 | valorISC | Monto ISC. Valor fijo o `tasaISC * precioItem` |

### listaItemOTI — List<oti>

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | C/C | N\|2 | tasaOTI | Códigos OTI: 01: SUME 911. 02: Tasa Portabilidad Numérica. 03: Impuesto sobre seguro (5%). 04: ATTT-Impuestos sobre seguro autos (1%). 05: Tasa Salida Aeropuerto (FZ). 06: Cargo Incentivo Desarrollo Aeropuerto (F3). 07: Cargo Seguridad Aeropuerto (AH). 08: Otros Cargos (XT). 09: Combustible (YQ). 10: FECI. 11: Intereses |
| String | C/C | N\|1..9\|2.6 | valorTasa | Valor de la tasa OTI |

## Vehículo (caso vehículo nuevo)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|2 | modalidadOperacionVenta | 01: venta a representante. 02: venta al consumidor final. 03: venta a gobierno. 04: venta a flota. 99: otros |
| String | C/C | AN\|10..50 | modalidadOperacionVentaNoDef | Obligatorio si modalidadOperacionVenta = 99 |
| String | SI | AN\|17 | chasis | Código VIN |
| String | C/C | AN\|1..4 | codigoColor | Color del vehículo (código interno del fabricante) |
| String | NO | AN\|1..40 | colorNombre | Descripción del color |
| String | SI | N\|1..4 | potenciaMotor | Potencia del motor (CV) |
| String | C/C | N\|1..4 | capacidadMotor | Expresada en litros |
| String | NO | N\|1..11\|0.4 | pesoNeto | Peso Neto en Toneladas |
| String | NO | N\|1..11\|0.4 | pesoBruto | Peso Bruto en Toneladas |
| String | SI | N\|2 | tipoCombustible | 01: gasolina. 02: diésel. 03: etanol. 08: eléctrico. 09: gasolina/eléctrico. 99: otro |
| String | C/C | AN\|10..50 | tipoCombustibleNoDef | Obligatorio si tipoCombustible = 99 |
| String | SI | AN\|1..21 | numeroMotor | Número del motor |
| String | NO | N\|1..11\|0.4 | capacidadTraccion | Capacidad máxima de tracción en Toneladas |
| String | NO | N\|1..11\|0.4 | distanciaEjes | Distancia entre ejes en Metros |
| String | C/C | N\|4 | anoModelo | Año del modelo |
| String | NO | N\|4 | anoFabricacion | Año de fabricación |
| String | NO | N\|1 | tipoPintura | 1: sólida. 2: metálica. 3: perla. 4: fosca. 9: otra |
| String | C/C | AN\|10..50 | tipoPinturaNodef | Obligatorio si tipoPintura = 9 |
| String | SI | N\|1 | tipoVehiculo | Tipos del Registro Vehicular: 1:motocicleta 2:bus 3:camión 4:sedán 5:camioneta 6:pickup 7:Mini Pick-Up 8:Pick Up Grua 9:Pick Cabina Sencilla 10:Pick Up Cabina y Media 11:Pick-Up Doble Cabina 12:Sedan Híbrido 13:Microbus 14:Mini Bus 15:Mini Microbus 16:Busito 17:Hatcback 18:Mini Panel 19:Panel 20:Convertible 21:Coupe 22:Coupe Convertible 23:Mini Car 24:Motorscooter 25:Mototriciclo 26:Scooter 27:Triciclo 28:Limosina 29:Mini Vans 30:Van 31:Trailer 32:Semi Trailer 33:Suburvan 34:Vagon 35:Vagoneta 36:Remolque 37:Four Wheel 38:SUV |
| String | SI | N\|1 | usoVehiculo | 1:comercial 2:particular 3:Diplomatica 4:Oficial 5:Especial |
| String | C/C | N\|1 | condicionVehiculo | 1: acabado. 2: semi-acabado. 3: inacabado |
| String | NO | N\|1..3 | capacidadPasajeros | Capacidad máxima de pasajeros sentados |

## Medicina (obligatorio si ítem es medicina o materia prima)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|5..35 | nroLote | Número de lote de medicinas |
| String | SI | N\|1..9 | cantProductosLote | Cantidad de productos en el lote |

## TotalesSubTotales

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|1..11\|1.2 | totalPrecioNeto | Suma de precios de ítems antes de impuestos |
| String | SI | N\|1..11\|2 | totalITBMS | Suma de ocurrencias de valorITBMS |
| String | SI | N\|1..11\|2 | totalISC | Suma de ocurrencias de valorISC |
| String | SI | N\|1..11\|1.2 | totalMontoGravado | totalITBMS + totalISC + valorTotalOTI |
| String | NO | N\|1..11\|1.2 | totalDescuento | Suma de ocurrencias de montoDescuento |
| String | NO | N\|1..11\|1.2 | totalAcarreoCobrado | Si se informa, no informar en precioAcarreo |
| String | NO | N\|1..11\|1.2 | valorSeguroCobrado | Si se informa, no informar en precioSeguro |
| String | SI | N\|1..11\|1.2 | totalFactura | totalTodosItems + totalAcarreoCobrado + valorSeguroCobrado - totalDescuento |
| String | SI | N\|1..11\|1.2 | totalValorRecibido | Suma de ocurrencias de valorCuotaPagada |
| String | C/C | N\|1..11\|1.2 | vuelto | Diferencia entre totalValorRecibido y totalFactura. No informar si no existe vuelto |
| String | SI | N\|1 | tiempoPago | **1: Inmediato. 2: Plazo. 3: Mixto** |
| String | SI | N\|1..11 | nroItems | Número total de ítems |
| String | SI | N\|1..11\|1.2 | totalTodosItems | Total de todos los ítems (suma de valorTotal) |
| String | NO | N\|1..11\|1.2 | totalOtrosGastos | Otros Gastos cobrados en el precio Total |

## DescuentoBonificacion (Opcional)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|..500 | descDescuento | Descripción de descuentos/bonificaciones aplicados a la factura |
| String | SI | N\|1..11\|2 | montoDescuento | Monto del descuento/bonificación |

## List<formaPago>

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|2 | formaPagoFact | **01: Crédito. 02: Efectivo. 03: Tarjeta Crédito. 04: Tarjeta Débito. 05: Tarjeta Fidelización. 06: Vale. 07: Tarjeta de Regalo. 08: Transf/Deposito cta. Bancaria. 09: Cheque. 99: Otro** |
| String | C/C | AN\|10..100 | descFormaPago | Obligatorio si formaPagoFact = 99 |
| String | SI | N\|1..11\|1..2 | valorCuotaPagada | Valor de la cuota pagada con esta forma de pago |

## Retencion (cuando aplica)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|2 | codigoRetencion | **1: Pago por servicio profesional al estado 100%. 2: Pago por venta de bienes/servicios al estado 50%. 3: Pago o acreditación a no domiciliado 100%. 4: Pago o acreditación por compra bienes/servicios 50%. 7: Pago a comercio afiliado a sistema TC/TD 50%. 8: Otros (disminución de retención)** |
| String | SI | N\|1..11\|1..2 | montoRetencion | Monto de la retención. Cálculo: `tasaITBMS * totalITBMS` |

## List<pagoPlazo> (Opcional)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|25 | fechaVenceCuota | Fecha de vencimiento de la cuota |
| String | SI | N\|1..11\|1..2 | valorCuota | Valor de la cuota |
| String | NO | AN\|15..1000 | infoPagoCuota | Información sobre esta cuota |

## listaTotalOTI — List<totalOTI>

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|2 | codigoTotalOTI | Mismos códigos OTI que en tasaOTI del ítem |
| String | SI | N\|1..11\|1..2 | valorTotalOTI | Sumatoria de codigoTotalOTI para ítems con ese código |

## pedidoComercialItem (Opcional)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|1..9 | nroPedidoCompraItem | Número del pedido de compra |
| String | SI | AN\|1..3 | nroItem | Número secuencial del ítem en el pedido |
| String | NO | AB\|0..5000 | infoItem | Información del pedido comercial relacionado al ítem |

## pedidoComercialGlobal (Opcional)

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | AN\|1..12 | nroPedidoCompraGlobal | Número del pedido de compra |
| String | C/C | AN\|1..12 | codigoReceptor | Código con que el facturador identifica al cliente |
| String | C/C | AN\|1..12 | nroAceptacion | Número de aceptación del pedido de compra |
| String | C/C | A\|1..50 | codigoSistemaEmisor | Código del sistema que emite la factura |
| String | C/C | AN\|0..5000 | InfoPedido | Información del pedido comercial global |

## InfoLogistica

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | NO | N\|1..3 | nroVolumenes | Número de volúmenes |
| String | NO | N\|1..4\|2..6 | pesoCarga | Peso total de la carga |
| String | NO | N\|1 | unidadPesoTotal | 1: g. 2: kg. 3: toneladas. 4: libras |
| String | NO | AN\|6 | licVehiculoCarga | Licencia del vehículo de carga |
| String | SI | AN\|15..100 | razonSocialTransportista | Nombre o Razón social del transportista |
| String | SI | N\|1 | tipoRucTransportista | 1: Natural. 2: Jurídico |
| String | SI | AN\|5..20 | rucTransportista | RUC del Transportista |
| String | SI | AN\|2 | digitoVerifRucTransportista | DV calculado por algoritmo DGI |
| String | NO | AN\|0..5000 | infoLogisticaEmisor | Información sobre logística |

## InfoEntrega

| Tipo | Req | Formato | Identificador | Descripción |
|---|---|---|---|---|
| String | SI | N\|1 | tipoRucEntrega | 1: Natural. 2: Jurídico |
| String | SI | AN\|5..20 | numeroRucEntrega | RUC del receptor |
| String | SI | AN\|2 | digitoVerifRucEntrega | DV calculado por algoritmo DGI |
| String | SI | An\|2..100 | razonSocialEntrega | Razón Social (Jurídica) o Nombre y Apellido (Natural) del local |
| String | SI | AN\|100 | direccionEntrega | Dirección del local de la entrega |
| String | SI | AN\|8 | codigoUbicacionEntrega | Código compuesto de la ubicación |
| String | SI | AN\|50 | corregimientoEntrega | Según Catálogo |
| String | SI | AN\|50 | distritoEntrega | Según Catálogo |
| String | SI | AN\|50 | provinciaEntrega | Según Catálogo |
| String | SI | N\|7..12 | telefonoEntrega | Formatos: 999-9999 ó 9999-9999 |
| String | SI | N\|7..12 | telefonoEntregaAlt | Teléfono alterno del local |

## Ejemplo de XML — REQUEST

```xml
<soapenv:Body>
  <tem:Enviar>
    <tem:tokenEmpresa>SOLICITAR</tem:tokenEmpresa>
    <tem:tokenPassword>SOLICITAR</tem:tokenPassword>
    <tem:documento>
      <ser:codigoSucursalEmisor>0000</ser:codigoSucursalEmisor>
      <ser:tipoSucursal>1</ser:tipoSucursal>
      <ser:datosTransaccion>
        <ser:tipoEmision>01</ser:tipoEmision>
        <ser:tipoDocumento>01</ser:tipoDocumento>
        <ser:numeroDocumentoFiscal>176</ser:numeroDocumentoFiscal>
        <ser:puntoFacturacionFiscal>505</ser:puntoFacturacionFiscal>
        <ser:fechaEmision>2022-07-15T07:49:27-05:00</ser:fechaEmision>
        <ser:fechaSalida>2022-07-15T07:49:27-05:00</ser:fechaSalida>
        <ser:naturalezaOperacion>01</ser:naturalezaOperacion>
        <ser:tipoOperacion>1</ser:tipoOperacion>
        <ser:destinoOperacion>1</ser:destinoOperacion>
        <ser:formatoCAFE>1</ser:formatoCAFE>
        <ser:entregaCAFE>1</ser:entregaCAFE>
        <ser:envioContenedor>1</ser:envioContenedor>
        <ser:procesoGeneracion>1</ser:procesoGeneracion>
        <ser:tipoVenta>1</ser:tipoVenta>
        <ser:informacionInteres>Informacion Interes</ser:informacionInteres>
        <ser:cliente>
          <ser:tipoClienteFE>01</ser:tipoClienteFE>
          <ser:tipoContribuyente>2</ser:tipoContribuyente>
          <ser:numeroRUC>155596713-2-2015</ser:numeroRUC>
          <ser:digitoVerificadorRUC>59</ser:digitoVerificadorRUC>
          <ser:razonSocial>Ambiente de pruebas</ser:razonSocial>
          <ser:direccion>Ave. La Paz</ser:direccion>
          <ser:codigoUbicacion>1-1-1</ser:codigoUbicacion>
          <ser:provincia>Bocas del Toro</ser:provincia>
          <ser:distrito>Bocas del Toro</ser:distrito>
          <ser:corregimiento>Bocas del Toro</ser:corregimiento>
          <ser:tipoIdentificacion></ser:tipoIdentificacion>
          <ser:nroIdentificacionExtranjero></ser:nroIdentificacionExtranjero>
          <ser:paisExtranjero></ser:paisExtranjero>
          <ser:telefono1>9999-9999</ser:telefono1>
          <ser:correoElectronico1>usuario@pruebas.com</ser:correoElectronico1>
          <ser:pais>PA</ser:pais>
          <ser:paisOtro/>
        </ser:cliente>
      </ser:datosTransaccion>
      <ser:listaItems>
        <ser:item>
          <ser:descripcion>Lapiz</ser:descripcion>
          <ser:codigo>CA-001</ser:codigo>
          <ser:unidadMedida>m</ser:unidadMedida>
          <ser:cantidad>1.00</ser:cantidad>
          <ser:fechaFabricacion>1999-05-10</ser:fechaFabricacion>
          <ser:fechaCaducidad>2019-05-10</ser:fechaCaducidad>
          <ser:unidadMedidaCPBS></ser:unidadMedidaCPBS>
          <ser:infoItem>Lapiz con Goma</ser:infoItem>
          <ser:precioUnitario>1234.12</ser:precioUnitario>
          <ser:precioUnitarioDescuento></ser:precioUnitarioDescuento>
          <ser:precioItem>1234.12</ser:precioItem>
          <ser:valorTotal>1335.3384</ser:valorTotal>
          <ser:codigoGTIN>0</ser:codigoGTIN>
          <ser:cantGTINCom>0.99</ser:cantGTINCom>
          <ser:codigoGTINInv>0</ser:codigoGTINInv>
          <ser:cantGTINComInv>1.00</ser:cantGTINComInv>
          <ser:tasaITBMS>01</ser:tasaITBMS>
          <ser:valorITBMS>86.3884</ser:valorITBMS>
          <ser:listaItemOTI>
            <ser:oti>
              <ser:tasaOTI>01</ser:tasaOTI>
              <ser:valorTasa>14.81</ser:valorTasa>
            </ser:oti>
            <ser:oti>
              <ser:tasaOTI>02</ser:tasaOTI>
              <ser:valorTasa>0.02</ser:valorTasa>
            </ser:oti>
          </ser:listaItemOTI>
        </ser:item>
      </ser:listaItems>
      <ser:totalesSubTotales>
        <ser:totalPrecioNeto>1234.12</ser:totalPrecioNeto>
        <ser:totalITBMS>86.39</ser:totalITBMS>
        <ser:totalMontoGravado>101.22</ser:totalMontoGravado>
        <ser:totalDescuento>0.00</ser:totalDescuento>
        <ser:totalAcarreoCobrado>0.00</ser:totalAcarreoCobrado>
        <ser:valorSeguroCobrado>0.00</ser:valorSeguroCobrado>
        <ser:totalOtrosGastos>0.00</ser:totalOtrosGastos>
        <ser:totalFactura>1335.34</ser:totalFactura>
        <ser:totalValorRecibido>1335.34</ser:totalValorRecibido>
        <ser:tiempoPago>1</ser:tiempoPago>
        <ser:nroItems>1</ser:nroItems>
        <ser:totalTodosItems>1335.34</ser:totalTodosItems>
        <ser:listaFormaPago>
          <ser:formaPago>
            <ser:formaPagoFact>02</ser:formaPagoFact>
            <ser:descFormaPago/>
            <ser:valorCuotaPagada>1335.34</ser:valorCuotaPagada>
          </ser:formaPago>
        </ser:listaFormaPago>
        <ser:listaTotalOTI>
          <ser:totalOti>
            <ser:codigoTotalOTI>01</ser:codigoTotalOTI>
            <ser:valorTotalOTI>14.81</ser:valorTotalOTI>
          </ser:totalOti>
          <ser:totalOti>
            <ser:codigoTotalOTI>02</ser:codigoTotalOTI>
            <ser:valorTotalOTI>0.02</ser:valorTotalOTI>
          </ser:totalOti>
        </ser:listaTotalOTI>
      </ser:totalesSubTotales>
      <ser:usoPosterior>
        <ser:cufe></ser:cufe>
      </ser:usoPosterior>
    </tem:documento>
  </tem:Enviar>
</soapenv:Body>
```

## EnviarResponse: Parámetros a Recibir

| Tipo | Identificador | Descripción |
| --- | --- | --- |
| String | codigo | Código del resultado de la operación |
| String | resultado | Indica el resultado del envío |
| String | mensaje | Mensaje adicional de la operación |
| String | CUFE | Código Único de Factura Electrónica |
| String | QR | Código QR de la Factura Electrónica |
| String | fechaRecepcionDGI | Fecha de autorización ante la DGI |
| String | nroProtocoloAutorizacion | Número del protocolo de autorización ante la DGI |

## Ejemplo de XML — RESPONSE

```xml
<EnviarResponse xmlns="http://tempuri.org/">
  <EnviarResult xmlns:a="http://schemas.datacontract.org/2004/07/Services.Response" xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
    <a:codigo>200</a:codigo>
    <a:resultado>procesado</a:resultado>
    <a:mensaje>El documento se envió correctamente.</a:mensaje>
    <a:cufe>FE0120000155596713-2-2015-5900012019052800055000155650121566749040</a:cufe>
    <a:qr>https://dgi-fep-test.mef.gob.pa:40001/Consultas/FacturasPorQR?chFE=...</a:qr>
    <a:fechaRecepcionDGI>2019-05-29T13:32:21+00:00</a:fechaRecepcionDGI>
    <a:nroProtocoloAutorizacion>20190000000000179413</a:nroProtocoloAutorizacion>
    <a:fechaLimite i:nil="true"/>
  </EnviarResult>
</EnviarResponse>
```
