package catalog

import (
	"fmt"
	"regexp"
	"strings"
)

// RUC (Registro Único de Contribuyente) handling.
//
// Panama distinguishes two RUC shapes:
//
//   - Natural person: the RUC is the person's cédula, e.g. "8-123-456".
//   - Legal entity (jurídica): a Registro Público-derived number, commonly
//     written as "FOLIO-ROLLO-AÑO" / matrícula, e.g. "155596713-2-2015".
//
// Every RUC has an associated dígito verificador (DV) issued by the DGI and sent
// separately in the digitoVerificadorRUC field. The DV cannot be derived locally
// in a guaranteed-correct way; the authoritative source is the QueryRUC operation
// on the HKA client. This package therefore validates RUC *structure* only and
// leaves DV verification to QueryRUC.

// RUCKind classifies a RUC as natural-person or legal-entity.
type RUCKind int

const (
	RUCDesconocido RUCKind = iota
	RUCNatural             // cédula-based
	RUCJuridico            // Registro Público-based
)

func (k RUCKind) String() string {
	switch k {
	case RUCNatural:
		return "natural"
	case RUCJuridico:
		return "jurídico"
	default:
		return "desconocido"
	}
}

// RUC is a parsed RUC value.
type RUC struct {
	Raw    string
	Kind   RUCKind
	Cedula *Cedula // populated when Kind == RUCNatural
	// Segments holds the dash-separated parts (useful for juridical RUCs).
	Segments []string
}

var reJuridico = regexp.MustCompile(`^\d{2,12}-\d{1,4}-\d{4}$`)

// ParseRUC parses and structurally validates a RUC string. It infers whether the
// RUC belongs to a natural person (cédula form) or a legal entity.
func ParseRUC(s string) (RUC, error) {
	s = strings.ToUpper(strings.TrimSpace(s))
	r := RUC{Raw: s, Segments: strings.Split(s, "-")}

	if c, err := ParseCedula(s); err == nil {
		r.Kind = RUCNatural
		r.Cedula = &c
		return r, nil
	}
	if reJuridico.MatchString(s) {
		r.Kind = RUCJuridico
		return r, nil
	}
	return RUC{}, fmt.Errorf("catalog: RUC %q is neither a valid cédula nor a juridical RUC (FOLIO-ROLLO-AÑO)", s)
}

// ValidateRUC reports whether s is a structurally valid RUC (natural or juridical).
func ValidateRUC(s string) bool {
	_, err := ParseRUC(s)
	return err == nil
}
