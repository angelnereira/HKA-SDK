package catalog

import "regexp"

// CUFE / CAFE definitions.
//
// CUFE — Código Único de Factura Electrónica: the unique code the DGI/PAC assigns
// to each authorized document. In the HKA "FE" scheme it is a 66-character string
// that begins with "FE" and embeds, among other fields, the emitter RUC, the
// document type, the billing point, the document number, the emission date, a
// security code and a check digit.
//
// CAFE — Comprobante Auxiliar de Factura Electrónica: the human-readable PDF
// representation of an authorized document (what the buyer receives). It is a
// rendering of the document, not a code; its layout is defined by the DGI.
//
// The authoritative, field-by-field composition of the CUFE and its check-digit
// algorithm are defined in the DGI "Ficha Técnica de Factura Electrónica"
// (https://dgi.mef.gob.pa/_7facturaelectronica/ftPAC). To avoid encoding an
// unverified layout, this package validates the CUFE's shape and length only; the
// CUFE itself is produced and signed by HKA/DGI, never constructed by the client.

const (
	// CUFELength is the fixed length of a CUFE in the HKA FE scheme.
	CUFELength = 66
	// CUFEPrefix is the leading marker of an HKA electronic-invoice CUFE.
	CUFEPrefix = "FE"
)

var reCUFE = regexp.MustCompile(`^FE[A-Z0-9-]{64}$`)

// ValidateCUFE reports whether cufe has the required 66-character HKA FE shape.
func ValidateCUFE(cufe string) bool {
	return len(cufe) == CUFELength && reCUFE.MatchString(cufe)
}

// DescribeCUFE returns a short human-readable explanation of the CUFE.
func DescribeCUFE() string {
	return "CUFE — Código Único de Factura Electrónica: identificador único de 66 " +
		"caracteres asignado por la DGI/PAC a cada documento autorizado. Es generado y " +
		"firmado por HKA/DGI; el cliente nunca lo construye."
}

// DescribeCAFE returns a short human-readable explanation of the CAFE.
func DescribeCAFE() string {
	return "CAFE — Comprobante Auxiliar de Factura Electrónica: representación PDF " +
		"legible del documento autorizado que se entrega al receptor."
}
