# Ejemplo en lenguaje Python

> Fuente: https://felwiki.thefactoryhka.com.pa/lenguaje_python

## Introducción

Este documento detalla el proceso para consumir los métodos del web service FEL de Panamá utilizando el **lenguaje Python 3**.

## Requisitos

### Python 3

**Descargar Python 3**: https://www.python.org/downloads

### Cliente SOAP: `zeep`

Una vez instalado Python 3, instalar el cliente SOAP `zeep` utilizando pip:

**Documentación oficial de zeep**: https://docs.python-zeep.org/en/master/

## Pasos para instalar zeep

1. Abrir **CMD** (o terminal) en la carpeta que contiene los archivos de Python
2. Ejecutar el siguiente comando:

```bash
pip install zeep
```

## Descargar archivos de ejemplo

- [Ejemplos Python (ZIP 6 KB)](https://felwiki.thefactoryhka.com.pa/_media/metodospython.zip)

---

## Ejemplo base de integración con zeep

> **Nota**: El wiki oficial solo lista requisitos e instalación. El código completo está en el ZIP. A continuación un ejemplo base reconstruido del patrón típico de integración SOAP con `zeep`:

```python
from zeep import Client
from datetime import datetime
import pytz

# Configuración
WSDL_URL = "https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc?wsdl"
TOKEN_EMPRESA = "SOLICITAR"
TOKEN_PASSWORD = "SOLICITAR"

# Crear cliente SOAP
client = Client(WSDL_URL)

# Fecha emisión formato requerido
tz_panama = pytz.timezone('America/Panama')
fecha_emision = datetime.now(tz_panama).strftime('%Y-%m-%dT%H:%M:%S-05:00')

# Construir documento electrónico
documento = {
    'codigoSucursalEmisor': '0000',
    'tipoSucursal': '1',
    'datosTransaccion': {
        'tipoEmision': '01',
        'tipoDocumento': '01',
        'numeroDocumentoFiscal': '0000000001',
        'puntoFacturacionFiscal': '001',
        'fechaEmision': fecha_emision,
        'naturalezaOperacion': '01',
        'tipoOperacion': '1',
        'destinoOperacion': '1',
        'formatoCAFE': '1',
        'entregaCAFE': '1',
        'envioContenedor': '1',
        'procesoGeneracion': '1',
        'tipoVenta': '1',
        'informacionInteres': 'Factura de prueba',
        'cliente': {
            'tipoClienteFE': '01',
            'tipoContribuyente': '2',
            'numeroRUC': '155596713-2-2015',
            'digitoVerificadorRUC': '59',
            'razonSocial': 'Ambiente de pruebas',
            'direccion': 'Ave. La Paz',
            'codigoUbicacion': '1-1-1',
            'provincia': 'Bocas del Toro',
            'distrito': 'Bocas del Toro',
            'corregimiento': 'Bocas del Toro',
            'telefono1': '9999-9999',
            'correoElectronico1': 'usuario@pruebas.com',
            'pais': 'PA',
        }
    },
    'listaItems': {
        'item': [{
            'descripcion': 'Lapiz',
            'codigo': 'CA-001',
            'unidadMedida': 'm',
            'cantidad': '1.00',
            'precioUnitario': '10.00',
            'precioItem': '10.00',
            'valorTotal': '10.70',
            'tasaITBMS': '01',
            'valorITBMS': '0.70',
        }]
    },
    'totalesSubTotales': {
        'totalPrecioNeto': '10.00',
        'totalITBMS': '0.70',
        'totalMontoGravado': '0.70',
        'totalFactura': '10.70',
        'totalValorRecibido': '10.70',
        'tiempoPago': '1',
        'nroItems': '1',
        'totalTodosItems': '10.70',
        'listaFormaPago': {
            'formaPago': [{
                'formaPagoFact': '02',
                'valorCuotaPagada': '10.70',
            }]
        }
    }
}

# Llamar al método Enviar
try:
    response = client.service.Enviar(
        tokenEmpresa=TOKEN_EMPRESA,
        tokenPassword=TOKEN_PASSWORD,
        documento=documento
    )

    print(f"Código: {response.codigo}")
    print(f"Resultado: {response.resultado}")
    print(f"Mensaje: {response.mensaje}")
    print(f"CUFE: {response.cufe}")
    print(f"QR: {response.qr}")
except Exception as e:
    print(f"Error: {e}")
```

## Otros métodos

Los demás métodos (`EstadoDocumento`, `AnulacionDocumento`, `DescargaXML`, `FoliosRestantes`, `EnvioCorreo`, `DescargaPDF`, `RastreoCorreo`, `ConsultarRucDV`) siguen el mismo patrón: `client.service.NombreMetodo(parametros)`.
