# Ejemplo en lenguaje PHP

> Fuente: https://felwiki.thefactoryhka.com.pa/lenguaje_php

Ejemplifica el proceso de conexión, acceso y uso de las propiedades y métodos ofrecidos por el Servicio Web de Integración de TFHKA a través de **PHP (v7 o superior)**.

## Criterios de integración

- Establecer ruta y Tokens para conectarse y usar los métodos del WS
- Construir el objeto **DocumentoElectronico** (Factura) a enviar
- Enviar el documento al WS para su procesamiento
- Recibir y procesar la respuesta del WS

## Pasos para realizar la integración

1. **PASO 1**: Definir variables con la ruta y los tokens requeridos por el WS
2. **PASO 2**: Construir el Documento Electrónico a enviar al WS
3. **PASO 3**: Enviar el Documento Electrónico al WS
4. **PASO 4**: Recibir y procesar la Respuesta del WS

> **IMPORTANTE**: El Documento procesado será almacenado en la **NUBE** para que pueda ser consultado en cualquier momento.

---

## PASO 1: Definir ruta y tokens

```php
$ruta = "https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc?singleWsdl";
$tokenEmpresa = "SOLICITAR";
$tokenPassword = "SOLICITAR";
```

## PASO 2: Construir el documento electrónico

```php
$numeroDocFiscal = "0000000000";
Mensajero::InfoNumero($numeroDocFiscal);

$factura = new DocumentoElectronico();
$cliente = new Cliente();
$datos = new DatosTransaccion();
$datos->cliente = $cliente;

// Debe establecerse el formato solicitado para la fecha de emisión "yyyy-MM-ddTHH:mm:ss-05:00"
$datos->fechaEmision = (new \DateTime('America/Panama'))->format('Y-m-d\TH:i:s-05:00');
$datos->numeroDocumentoFiscal = $numeroDocFiscal;
$factura->datosTransaccion = $datos;

$item = new Item();
$factura->listaItems = array($item);

$totales = new Totales();
$formaPago = new FormaPago();
$totales->listaFormaPago = array($formaPago);
$factura->totalesSubTotales = $totales;
```

## PASO 3: Enviar al Web Service

```php
try {
  $wsPa = new SoapClient($ruta);
  $parametros = array(
    'tokenEmpresa' => $tokenEmpresa,
    'tokenPassword' => $tokenPassword,
    'documento' => $factura,
  );

  // Enviamos el documento al método "Enviar" del WS
  $respWsPa = json_decode(json_encode($wsPa->__soapCall('Enviar', array($parametros))), true);
} catch (\Exception $e) {
  die($e);
}
```

## PASO 4: Procesar la respuesta

La respuesta incluye: **código, resultado, mensaje, CUFE, QR, entre otros**.

Si el código es **200** o **102**, se presentará un enlace al QR para ver los detalles:

```php
Mensajero::MostrarRespuesta($respWsPa);
// Acceso: $respWsPa["EnviarResult"]['codigo'], ['resultado'], ['mensaje'], ['cufe'], ['qr']
```

---

## Definición de clases

```php
class DocumentoElectronico
{
  public $codigoSucursalEmisor = "0000";
  public $tipoSucursal = "1";
  public $datosTransaccion;
  public $listaItems;
  public $totalesSubTotales;
}

class Cliente
{
  public $tipoClienteFE = "01";
  public $tipoContribuyente = "2";
  public $numeroRUC = "155596713-2-2015";
  public $digitoVerificadorRUC = "59";
  public $razonSocial = "FE general";
  public $direccion = "Av. Balboa";
  public $codigoUbicacion = "1-2-3";
  public $corregimiento = "Guabito";
  public $distrito = "Changuinola";
  public $provincia = "Bocas del Toro";
  public $telefono1 = "997-8243";
  public $telefono2 = "";
  public $telefono3 = "";
  public $correoElectronico1 = "fep@gmail.com";
  public $pais = "PA";
  public $paisOtro = "";
}

class DatosTransaccion
{
  public $tipoEmision = "01";
  public $tipoDocumento = "01";
  public $numeroDocumentoFiscal;
  public $puntoFacturacionFiscal = "001";
  public $fechaEmision;
  public $naturalezaOperacion = "01";
  public $tipoOperacion = "1";
  public $destinoOperacion = "1";
  public $formatoCAFE = "1";
  public $entregaCAFE = "1";
  public $envioContenedor = "1";
  public $procesoGeneracion = "1";
  public $tipoVenta = "1";
  public $informacionInteres = "Prueba de Información de interés";
  public $cliente;
}

class Item
{
  public $descripcion = "Cuadernos";
  public $codigo = "T";
  public $unidadMedida = "und";
  public $cantidad = "2.00";
  public $fechaFabricacion = "2020-12-25";
  public $unidadMedidaCPBS = "cm";
  public $precioUnitario = "69.00";
  public $precioUnitarioDescuento = "0.00";
  public $precioAcarreo = "1.01";
  public $precioSeguro = "12.01";
  public $precioItem = "138.00";
  public $valorTotal = "171.72";
  public $codigoGTIN = "0";
  public $cantGTINCom = "0.99";
  public $codigoGTINInv = "0";
  public $cantGTINComInv = "1.00";
  public $tasaITBMS = "03";
  public $valorITBMS = "20.70";
  public $tasaISC = "0.00";
  public $valorISC = "0.00";
  public $codigoCPBS = "1410";
}

class Totales
{
  public $totalPrecioNeto = "138.00";
  public $totalITBMS = "20.70";
  public $totalISC = "0.00";
  public $totalMontoGravado = "20.70";
  public $totalDescuento = "";
  public $totalAcarreoCobrado = "";
  public $valorSeguroCobrado = "";
  public $totalFactura = "171.72";
  public $totalValorRecibido = "171.72";
  public $vuelto = "0.00";
  public $tiempoPago = "1";
  public $nroItems = "1";
  public $totalTodosItems = "171.72";
  public $listaFormaPago;
}

class FormaPago
{
  public $formaPagoFact = "02";
  public $valorCuotaPagada = "171.72";
  public $descFormaPago = "";
}
```
