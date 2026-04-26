# Instructivo de Consumo de Servicios Web PAC

> Fuente: https://felwiki.thefactoryhka.com.pa/instructivo_de_consumo_de_servicios_web_pac

## Proceso para envío de documentos al PAC

El presente documento detalla el proceso paso a paso para enviar documentos electrónicos **directamente a los servicios PAC** (Proveedor Autorizado de Certificación) — es decir, el servicio oficial de la **DGI** que recibe el XML ya firmado. Esto es diferente del servicio de integración de HKA (`/ws/obj/v1.0/Service.svc`), que firma el XML por ti.

## Autenticación

> **Nota**: El proceso de autenticación y uso de los métodos del servicio PAC se realiza basado en el **certificado digital** al momento de establecer la comunicación. Debe ser provisto el **Certificado Digital de Transmisión** (parte pública).

## Métodos del servicio PAC

El servicio PAC consta de 4 métodos:

1. **RecepciónFE** — Envía FE firmada para certificación
2. **Recepción de Eventos** — Anulación y manifestación de documentos
3. **ConsultaFE** — Consulta FE por CUFE
4. **Descarga por Criterios** — Búsqueda por filtros (fecha, RUC, tipo)

---

## 1. Método de RecepciónFE

Recibe una factura electrónica ya firmada para su **certificación** por la DGI.

### REQUEST

| Identificador | Descripción |
|---|---|
| dVerForm | Versión del Formato |
| dId | Identificador de control de envío |
| iAmb | Ambiente de destino (1=Producción, 2=Pruebas) |
| xFe | XML de la FE transmitida (firmado) |

### Ejemplo Request (resumido)

```xml
<soapenv:Body>
  <fer:feDatosMsg>
    <fer:rEnviFe>
      <fer:dVerForm>1.00</fer:dVerForm>
      <fer:dId>01</fer:dId>
      <fer:iAmb>2</fer:iAmb>
      <fer:xFe><![CDATA[
        <rFE xmlns="http://dgi-fep.mef.gob.pa">
          <dVerForm>1.00</dVerForm>
          <dId>FE0120000155596713-2-2015-5900002020021000000050185050221010989249</dId>
          <gDGen>...datos generales...</gDGen>
          <gItem>...items...</gItem>
          <gTot>...totales...</gTot>
          <Signature xmlns="http://www.w3.org/2000/09/xmldsig#">...firma XMLDSig...</Signature>
        </rFE>
      ]]></fer:xFe>
    </fer:rEnviFe>
  </fer:feDatosMsg>
</soapenv:Body>
```

### RESPONSE

| Identificador | Descripción |
|---|---|
| dVerForm | Versión del Formato |
| iAmb | Ambiente de destino |
| dVerApl | Versión del aplicativo de recepción |
| dId | Identificador de control de envío |
| dCUFE | **CUFE de la FE Procesada** |
| dFecProc | Fecha y hora del procesamiento |
| dProAut | **Número de la autorización de uso** |
| dCodRes | Código del resultado |
| dMsgRes | Mensaje del resultado |

### Ejemplo Response

```xml
<rProtFe>
  <dVerForm>1.00</dVerForm>
  <gInfProt>
    <dId>ID20200000000000000322</dId>
    <iAmb>2</iAmb>
    <dVerApl>0.0.1</dVerApl>
    <dCUFE>FE0120000155596713-2-2015-5900002020071300000185800520124121596372</dCUFE>
    <dFecProc>2020-07-13T21:09:04-05:00</dFecProc>
    <dProAut>20200000000000000322</dProAut>
    <dDigVal>J5ZhLRLorQ9iw0WGs9CXt7dsGGvMrc4X1JCZFHu2/LM=</dDigVal>
    <gResProc>
      <dCodRes>0260</dCodRes>
      <dMsgRes>Autorizado el uso de la FE</dMsgRes>
    </gResProc>
  </gInfProt>
</rProtFe>
```

---

## 2. Método de Recepción de Eventos

Recibe y almacena los **eventos de anulación y manifestación** de los documentos electrónicos autorizados.

### REQUEST

| Identificador | Descripción |
|---|---|
| dVerForm | Versión del Formato |
| dId | Identificador de control de envío |
| iAmb | Ambiente de destino |
| dEvReg | **XML del evento para registrar** (firmado) |

### Ejemplo Request

```xml
<soapenv:Body>
  <fer:feDatosMsg>
    <fer:rEnviEventoFe>
      <fer:dVerForm>1.00</fer:dVerForm>
      <fer:dId>1</fer:dId>
      <fer:iAmb>2</fer:iAmb>
      <fer:dEvReg><![CDATA[
        <rEvAnulaFe xmlns="http://dgi-fep.mef.gob.pa">
          <dVerForm>1.00</dVerForm>
          <gInfProt>
            <dIdFirma>ID20230000000000000576</dIdFirma>
            <iAmb>2</iAmb>
            <dCufe>FE0120000155596713-2-2015-5900002021031500000991815050128384073517</dCufe>
            <gRucEm>
              <dTipContEm>2</dTipContEm>
              <dRucEm>155596713-2-2015</dRucEm>
              <dDvEm>59</dDvEm>
            </gRucEm>
            <dMotivoAn>La transaccion no se finiquito por lo cual se requiere anular.</dMotivoAn>
          </gInfProt>
          <Signature xmlns="http://www.w3.org/2000/09/xmldsig#">...firma...</Signature>
        </rEvAnulaFe>
      ]]></fer:dEvReg>
    </fer:rEnviEventoFe>
  </fer:feDatosMsg>
</soapenv:Body>
```

### RESPONSE

```xml
<s:Body>
  <rRetEnviEventoFe xmlns="http://dgi-fep.mef.gob.pa/wsdl/FeRecepFE">
    <dVerForm>1.00</dVerForm>
    <iAmb>2</iAmb>
    <dVerApl>0.1.0</dVerApl>
    <dFecProc>2021-05-24T14:45:42-05:00</dFecProc>
    <gResProc>
      <dCodRes>0600</dCodRes>
      <dMsgRes>Evento registrado con éxito</dMsgRes>
    </gResProc>
  </rRetEnviEventoFe>
</s:Body>
```

---

## 3. Método de ConsultaFE

Devuelve el archivo de una FE y los **eventos registrados** hasta la fecha, identificado por su CUFE.

### REQUEST

| Identificador | Descripción |
|---|---|
| dVerForm | Versión del Formato |
| dId | Identificador de control de envío |
| iAmb | Ambiente de destino |
| dCufe | **Código Único de Factura Electrónica** |

### Ejemplo Request

```xml
<soapenv:Body>
  <fer:feDatosMsg>
    <fer:rEnviConsFe>
      <fer:dVerForm>1.00</fer:dVerForm>
      <fer:dId>01</fer:dId>
      <fer:iAmb>2</fer:iAmb>
      <fer:dCufe>FE0120000155596713-2-2015-5900002021051200000992155050121410079104</fer:dCufe>
    </fer:rEnviConsFe>
  </fer:feDatosMsg>
</soapenv:Body>
```

### RESPONSE

| Identificador | Descripción |
|---|---|
| dVerForm | Versión del Formato |
| iAmb | Ambiente de destino |
| dId | Identificador de control |
| dVerApl | Versión del aplicativo |
| dFecProc | Fecha y hora del procesamiento |
| dCodRes | Código del resultado |
| dMsgRes | Mensaje del resultado |
| xContenFE | **Contenedor completo de la FE** (XML + firma) |

### Ejemplo Response (resumido)

```xml
<rRetEnviConsFe xmlns="http://dgi-fep.mef.gob.pa/wsdl/FeRecepFE">
  <dVerForm>1.00</dVerForm>
  <iAmb>2</iAmb>
  <dVerApl>0.1.0</dVerApl>
  <dFecProc>2021-05-24T14:50:20-05:00</dFecProc>
  <dCodRes>0422</dCodRes>
  <dMsgRes>Exito en la consulta</dMsgRes>
  <xContenFE>
    <rContFe>
      <dVerForm>1.00</dVerForm>
      <xFe>
        <rFE xmlns="http://dgi-fep.mef.gob.pa">
          ... FE completa con firma ...
        </rFE>
      </xFe>
    </rContFe>
  </xContenFE>
</rRetEnviConsFe>
```

---

## 4. Método de Descarga por Criterios

Devuelve archivos de documentos electrónicos y eventos registrados basado en **filtros múltiples**.

### REQUEST

| Identificador | Descripción |
|---|---|
| dVerForm | Versión del Formato |
| dId | Identificador de control |
| iAmb | Ambiente de destino |
| gRucEmi | RUC del emisor |
| gRucRec | RUC del receptor |
| iDoc | Tipo de documento |
| dSucEm | Sucursal del emisor |
| iTipEv | Tipo de evento |
| iTipoRec | Tipo de receptor |
| dFechaEmDesde | Fecha desde de la emisión |
| dFechaEmHasta | Fecha hasta de la emisión |
| iTipoCon | Tipo de consulta |
| dPag | Página de la consulta |

### Ejemplo Request

```xml
<fer:feDatosMsg>
  <fer:rEnviDescFe>
    <fer:dVerForm>1.00</fer:dVerForm>
    <fer:dId>1</fer:dId>
    <fer:iAmb>2</fer:iAmb>
    <fer:gRucEmi>
      <fer:dRuc>155596713-2-2015</fer:dRuc>
    </fer:gRucEmi>
    <fer:gRucRec>
      <fer:dRuc>155596713-2-2015</fer:dRuc>
    </fer:gRucRec>
    <fer:iDoc>01</fer:iDoc>
    <fer:dSucEm>0000</fer:dSucEm>
    <fer:iTipEv>0001</fer:iTipEv>
    <fer:iTipoRec>01</fer:iTipoRec>
    <fer:dFechaEmDesde>2021-08-16</fer:dFechaEmDesde>
    <fer:dFechaEmHasta>2021-08-17</fer:dFechaEmHasta>
    <fer:iTipoCon>1</fer:iTipoCon>
    <fer:dPag>01</fer:dPag>
  </fer:rEnviDescFe>
</fer:feDatosMsg>
```

### RESPONSE

| Identificador | Descripción |
|---|---|
| dVerForm | Versión del Formato |
| iAmb | Ambiente de destino |
| dVerApl | Versión del aplicativo |
| dFecProc | Fecha y hora del procesamiento |
| dCodRes | Código de resultado |
| dMsgRes | Mensaje del resultado |
| dPag | Página de consulta |
| dTotPag | Total de páginas |
| dTotdoc | **Total de documentos electrónicos encontrados** |
| xContenFE | Contenedor de las FE |

### Ejemplo Response (resumido)

```xml
<rRetConsDescFe xmlns="">
  <dVerForm>1.00</dVerForm>
  <iAmb>2</iAmb>
  <dVerApl>0.0.1</dVerApl>
  <dFecProc>2021-10-18T21:56:49-05:00</dFecProc>
  <gResProc>
    <dCodRes>0761</dCodRes>
    <dMsgRes>Éxito en la consulta</dMsgRes>
  </gResProc>
  <dPag>1</dPag>
  <dTotPag>1</dTotPag>
  <dTotdoc>4</dTotdoc>
  <xContenFE>
    <rContFe xmlns="">
      ... 4 documentos FE completos ...
    </rContFe>
  </xContenFE>
</rRetConsDescFe>
```

---

## Diferencias entre la integración HKA vs. PAC directo

| Aspecto | HKA (Service.svc) | PAC directo (DGI) |
|---|---|---|
| URL | `https://demoemision.thefactoryhka.com.pa/ws/obj/v1.0/Service.svc` | DGI `dgi-fep.mef.gob.pa` |
| Autenticación | `tokenEmpresa` + `tokenPassword` | Certificado Digital de Transmisión |
| Firma XML | **HKA firma por ti** | **Tú debes firmar el XML** (XMLDSig RSA-SHA256) |
| Estructura | Objeto `documento` con múltiples listas | `rFE` según esquema oficial DGI |
| Método envío | `Enviar()` | `RecepciónFE` |
| Método anular | `AnulacionDocumento()` | `Recepción de Eventos` con `rEvAnulaFe` |
| Método consulta | `EstadoDocumento()` | `ConsultaFE` (por CUFE) |

**Recomendación para Sago One**: Usar la integración HKA (ya hecha en Sago One) para facturar con menos complejidad de firma. El PAC directo sería un fallback/complemento para operaciones avanzadas como búsqueda por criterios múltiples o contingencia.
