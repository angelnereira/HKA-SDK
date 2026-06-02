package catalog

import (
	"fmt"
	"regexp"

	"github.com/angelnereira/hka-sdk/types"
)

// CUFE / CAFE — the two pillars of identification and representation of Panama's
// electronic invoicing.
//
// CUFE — Código Único de Factura Electrónica
//
//	An alphanumeric identifier (~66 characters) that uniquely identifies each
//	invoice authorized by the DGI. It is produced by the emitting system following
//	the algorithm dictated by the DGI; in the HKA integration model the PAC returns
//	it once the document is validated and processed successfully (return code 200).
//	It is the logical/relational value to persist: it is mandatory when referencing
//	a document from a credit/debit note (types 04/05), it identifies a document in
//	RastreoCorreo (TrackEmail), and it lets anyone verify the invoice on the DGI
//	portal. The emitter never invents it — this package validates and inspects it.
//
//	Example: FE0120000155596713-2-2015-5900012019052800055000155650121566749040
//
// CAFE — Comprobante Auxiliar de Factura Electrónica
//
//	The graphical/printable representation of the authorized invoice (the document
//	delivered to the receiver). It is generated from the authorized XML and obtained
//	via DescargaPDF (DownloadPDF), which returns a Base64-encoded PDF. Its distinctive
//	elements are the QR code, the CUFE and the authorization data. It must remain
//	legible for at least six months for immediate verification.
//
// The full field-by-field composition of the CUFE and its check-digit algorithm are
// defined in the DGI "Ficha Técnica de Factura Electrónica"
// (https://dgi.mef.gob.pa/_7facturaelectronica/ftPAC). This package decodes only the
// fixed leading fields that are documented and verified against real CUFEs.

const (
	// CUFELength is the length of a CUFE in the HKA FE scheme.
	CUFELength = 66
	// CUFEPrefix is the leading marker of an HKA electronic-invoice CUFE.
	CUFEPrefix = "FE"
)

// AmbienteCUFE indicates the environment a CUFE was issued in.
type AmbienteCUFE string

const (
	AmbienteProduccion AmbienteCUFE = "1"
	AmbientePruebas    AmbienteCUFE = "2"
)

// Descripcion returns a human-readable description of the environment.
func (a AmbienteCUFE) Descripcion() string {
	switch a {
	case AmbienteProduccion:
		return "Producción"
	case AmbientePruebas:
		return "Pruebas / Demo"
	default:
		return "Ambiente desconocido"
	}
}

var reCUFE = regexp.MustCompile(`^FE[A-Z0-9-]{64}$`)

// ValidateCUFE reports whether cufe has the required 66-character HKA FE shape.
func ValidateCUFE(cufe string) bool {
	return len(cufe) == CUFELength && reCUFE.MatchString(cufe)
}

// CUFEInfo holds the leading fields decoded from a CUFE. Only the fields that are
// documented and verified are exposed; Raw carries the full value for persistence.
type CUFEInfo struct {
	Raw           string
	TipoDocumento types.TipoDocumento // characters 3-4
	Ambiente      AmbienteCUFE        // character 5: 1=producción, 2=pruebas
}

// ParseCUFE validates and decodes the leading fields of a CUFE. Deeper field
// extraction (RUC, fecha, número, dígito verificador) requires the DGI ficha
// técnica layout and is intentionally not attempted here to avoid mis-parsing.
func ParseCUFE(cufe string) (CUFEInfo, error) {
	if !ValidateCUFE(cufe) {
		return CUFEInfo{}, fmt.Errorf("catalog: CUFE %q must be %d characters starting with %q", cufe, CUFELength, CUFEPrefix)
	}
	return CUFEInfo{
		Raw:           cufe,
		TipoDocumento: types.TipoDocumento(cufe[2:4]),
		Ambiente:      AmbienteCUFE(cufe[4:5]),
	}, nil
}

// DescribeCUFE returns a short human-readable explanation of the CUFE.
func DescribeCUFE() string {
	return "CUFE — Código Único de Factura Electrónica: identificador único (~66 " +
		"caracteres) generado por el emisor según el algoritmo de la DGI y devuelto por " +
		"el PAC al autorizar (código 200). Es el dato lógico a persistir para referenciar " +
		"(notas 04/05), rastrear y verificar el documento."
}

// DescribeCAFE returns a short human-readable explanation of the CAFE.
func DescribeCAFE() string {
	return "CAFE — Comprobante Auxiliar de Factura Electrónica: representación PDF " +
		"legible del documento autorizado (con QR, CUFE y datos de autorización), " +
		"obtenida vía DescargaPDF. Es el entregable visual para el receptor."
}

// DescribeFormatoCAFE explains a FormatoCAFE code.
func DescribeFormatoCAFE(f types.FormatoCAFE) string {
	switch f {
	case types.CAFESinGeneracion:
		return "Sin generación de CAFE"
	case types.CAFECintaPapel:
		return "Cinta de papel — punto de venta (POS) / retail"
	case types.CAFEPapelCarta:
		return "Papel formato carta — facturación administrativa / B2B"
	default:
		return "Formato CAFE desconocido"
	}
}
