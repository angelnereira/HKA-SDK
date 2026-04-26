# Ejemplo en lenguaje C# (CSharp)

> Fuente: https://felwiki.thefactoryhka.com.pa/lenguaje_c

Ejemplifica el proceso de conexión al **Servicio Web de Integración de TFHKA** a través de **C# (plataforma .Net)**. Independiente del tipo de proyecto: Consola, Winforms o WPF.

## Criterios de integración

- Establecer una **Referencia al Servicio Web** de TFHKA
- Instanciar el objeto `ServiceClient` para acceder a los métodos
- Construir el objeto **DocumentoElectronico**
- Enviar el documento al WS
- Recibir y procesar la respuesta

## Pasos para realizar la integración

1. **PASO 1**: Agregar una nueva Referencia de Servicio con la URL del WS
2. **PASO 2**: Crear una instancia del objeto `ServiceClient`
3. **PASO 3**: Construir el Documento Electrónico
4. **PASO 4**: Enviar al WS con Tokens
5. **PASO 5**: Recibir y procesar la respuesta

> **IMPORTANTE**: El Documento procesado será almacenado en la **NUBE** para ser consultado en cualquier momento.

---

## PASO 1: Agregar Referencia de Servicio

En **Visual Studio**:
- En el Explorador de Soluciones, click derecho sobre el nombre del Proyecto
- Seleccionar **"Agregar → Referencia de servicio..."**
- En el campo `Dirección:` ingresar: `https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc`
- Click en **"Ir"** y esperar a que descargue la información
- En **"Espacio de nombres:"** colocar: `TfhkaPaWebService` (referencial, no obligatorio)

> **IMPORTANTE**: Para que el ejemplo funcione, debe agregar la Referencia Web a través de **"Agregar Referencia de servicio..."**, lo cual generará el archivo `Reference.cs` donde se crearán las clases y objetos requeridos por el `Data Contract` del SOAP Service.

```csharp
using System;
using EjemploIntegraciondirectaCSharp.TfhkaPaWebService;

namespace EjemploIntegraciondirectaCSharp
{
    class DemoIntCSharpTfhka
    {
        public const string tokenEmpresa = ""; // Suministrado por soporte TFHKA
        public const string tokenPassword = ""; // Suministrado por soporte TFHKA
```

## PASO 2: Crear instancia del ServiceClient

```csharp
static void Main(string[] args)
{
    try
    {
        ServiceClient webService = new ServiceClient();

        Consola.mensaje("\r\n - EJEMPLO DE INTEGRACIÓN DIRECTA C# TFHKA - \r\n", 1, 0, 0);
        String numero = "";
        while (numero == "")
        {
            Consola.mensaje("  Introduzca el Número de Documento Electrónico a enviar:\r\n", 0, 2, 0);
            numero = Console.ReadLine();
        }
```

## PASO 3: Construir el Documento Electrónico

```csharp
DocumentoElectronico documento = Factura.armarFactura(numero);
```

## PASO 4: Enviar al WS

```csharp
EnviarResponse respuestaWebService = new EnviarResponse();
Consola.mensaje("\r\n  Enviando el Documento...", 2, 0, 0);
respuestaWebService = webService.Enviar(tokenEmpresa, tokenPassword, documento);
```

## PASO 5: Procesar la respuesta

```csharp
Consola.mensaje("\r\n  RESPUESTA: " + respuestaWebService.codigo + ". " +
                respuestaWebService.resultado + ". " +
                respuestaWebService.mensaje, 0, 0, 0);

if (respuestaWebService.codigo == "200")
{
    // Abrir página de Consulta de Facturas por QR en el navegador
    System.Diagnostics.Process.Start(respuestaWebService.qr);
}
```

---

## Clases para armar el documento

### Factura

```csharp
public class Factura
{
    public static DocumentoElectronico factura = new DocumentoElectronico();
    public static DocumentoElectronico armarFactura(string numero)
    {
        factura.codigoSucursalEmisor = "0000";
        factura.tipoSucursal = "1";

        // Datos de la Transacción
        DatosFactura datos = new DatosFactura();
        datosTransaccion datosFactura = datos.crearDatosFactura(numero);
        factura.datosTransaccion = datosFactura;

        // Datos de los Productos / Items
        factura.listaItems = new listaItems();
        factura.listaItems.Add(Producto.crearItem());

        // Datos de los Totales y Sub-Totales
        factura.totalesSubTotales = Totales.generarTotales();

        return factura;
    }
}
```

### Cliente

```csharp
public class Cliente
{
    public static cliente cliente = new cliente();
    public static cliente crearCliente()
    {
        cliente.tipoClienteFE = "01";
        cliente.tipoContribuyente = "2";
        cliente.numeroRUC = "155596713-2-2015";
        cliente.digitoVerificadorRUC = "59";
        cliente.razonSocial = "FE general";
        cliente.direccion = "Av. Balboa";
        cliente.codigoUbicacion = "1-2-3";
        cliente.corregimiento = "Guabito";
        cliente.distrito = "Changuinola";
        cliente.provincia = "Bocas del Toro";
        cliente.telefono1 = "997-8243";
        cliente.telefono2 = "";
        cliente.telefono3 = "";
        cliente.correoElectronico1 = "fep@gmail.com";
        cliente.pais = "PA";
        cliente.paisOtro = "";

        return cliente;
    }
}
```

### DatosFactura

```csharp
public class DatosFactura
{
    public static datosTransaccion datosFactura = new datosTransaccion();
    public datosTransaccion crearDatosFactura(string num)
    {
        datosFactura.tipoEmision = "01";
        datosFactura.tipoDocumento = "01";
        datosFactura.numeroDocumentoFiscal = num;
        datosFactura.puntoFacturacionFiscal = "001";

        // Formato requerido: "yyyy-MM-ddTHH:mm:ss-05:00"
        var fechaEmision = Convert.ToDateTime(DateTime.Today.ToString("dd-MM-yyyy"))
            .ToString("yyyy-MM-ddTHH:mm:ss-05:00");
        datosFactura.fechaEmision = fechaEmision;

        datosFactura.naturalezaOperacion = "01";
        datosFactura.tipoOperacion = "1";
        datosFactura.destinoOperacion = "1";
        datosFactura.formatoCAFE = "1";
        datosFactura.entregaCAFE = "1";
        datosFactura.envioContenedor = "1";
        datosFactura.procesoGeneracion = "1";
        datosFactura.tipoVenta = "1";
        datosFactura.informacionInteres = "Prueba de Información de interés";
        datosFactura.cliente = Cliente.crearCliente();

        return datosFactura;
    }
}
```

### Producto (Item)

```csharp
public class Producto
{
    public static Item item = new Item();
    public static Item crearItem()
    {
        item.descripcion = "Cuadernos";
        item.codigo = "T";
        item.unidadMedida = "und";
        item.cantidad = "2.00";
        item.fechaFabricacion = "2020-12-25";
        item.unidadMedidaCPBS = "cm";
        item.precioUnitario = "69.00";
        item.precioUnitarioDescuento = "0.00";
        item.precioAcarreo = "1.01";
        item.precioSeguro = "12.01";
        item.precioItem = "138.00";
        item.valorTotal = "171.72";
        item.codigoGTIN = "0";
        item.cantGTINCom = "0.99";
        item.codigoGTINInv = "0";
        item.cantGTINComInv = "1.00";
        item.tasaITBMS = "03";
        item.valorITBMS = "20.70";
        item.tasaISC = "0.00";
        item.valorISC = "0.00";
        item.codigoCPBS = "1410";

        return item;
    }
}
```

### Totales

```csharp
public class Totales
{
    public static totalesSubTotales TotalesSubTotales = new totalesSubTotales();
    public static totalesSubTotales generarTotales()
    {
        TotalesSubTotales.totalPrecioNeto = "138.00";
        TotalesSubTotales.totalITBMS = "20.70";
        TotalesSubTotales.totalISC = "0.00";
        TotalesSubTotales.totalMontoGravado = "20.70";
        TotalesSubTotales.totalDescuento = "";
        TotalesSubTotales.totalAcarreoCobrado = "";
        TotalesSubTotales.valorSeguroCobrado = "";
        TotalesSubTotales.totalFactura = "171.72";
        TotalesSubTotales.totalValorRecibido = "171.72";
        TotalesSubTotales.vuelto = "0.00";
        TotalesSubTotales.tiempoPago = "1";
        TotalesSubTotales.nroItems = "1";
        TotalesSubTotales.totalTodosItems = "171.72";

        TotalesSubTotales.listaFormaPago = new listaFormaPago();
        FormaPago formaPago1 = new FormaPago();
        formaPago1.formaPagoFact = "02";
        formaPago1.valorCuotaPagada = "171.72";
        formaPago1.descFormaPago = "";
        TotalesSubTotales.listaFormaPago.Add(formaPago1);

        return TotalesSubTotales;
    }
}
```
