package catalog

import (
	"fmt"
	"regexp"
	"strings"
)

// Cédula (national ID) format in Panama.
//
// A cédula is written as "PREFIJO-TOMO-PARTIDA":
//   - PREFIJO: the province of registration (numeric 1..13) or a special letter
//     prefix for non-standard cases.
//   - TOMO (a.k.a. "libro"): a sequential book number.
//   - PARTIDA (a.k.a. "folio"/"asiento"): the entry within the book.
//
// Example: "8-123-456" — born/registered in Panamá province.
//
// For natural persons the RUC is generally the cédula itself, so this parser is
// also used as the basis for natural-person RUC handling.

// CedulaPrefijo describes a recognized cédula prefix.
type CedulaPrefijo struct {
	Prefijo     string
	Descripcion string
}

// cedulaPrefijos lists the special (non-province) prefixes. Numeric province
// prefixes 1..13 are validated against the geographic catalog.
var cedulaPrefijos = map[string]string{
	"PE": "Panameño nacido en el extranjero / antigua Zona del Canal",
	"E":  "Extranjero residente",
	"N":  "Naturalizado",
	"AV": "Persona pendiente de inscripción / casos especiales",
}

// SpecialCedulaPrefijos returns the recognized non-province cédula prefixes.
func SpecialCedulaPrefijos() []CedulaPrefijo {
	out := make([]CedulaPrefijo, 0, len(cedulaPrefijos))
	for p, d := range cedulaPrefijos {
		out = append(out, CedulaPrefijo{Prefijo: p, Descripcion: d})
	}
	return out
}

// Cedula is a parsed cédula number.
type Cedula struct {
	Prefijo string // province code (1..13) or special letter prefix
	Tomo    string
	Partida string
}

func (c Cedula) String() string { return c.Prefijo + "-" + c.Tomo + "-" + c.Partida }

// EsProvincial reports whether the prefix is a numeric province code (vs. a special
// letter prefix such as E, N, PE).
func (c Cedula) EsProvincial() bool {
	return reNumeric.MatchString(c.Prefijo)
}

var (
	reNumeric = regexp.MustCompile(`^\d{1,2}$`)
	reCedula  = regexp.MustCompile(`^([0-9]{1,2}|PE|E|N|AV)-(\d{1,4})-(\d{1,6})$`)
)

// ParseCedula parses and validates a cédula string of the form "PREFIJO-TOMO-PARTIDA".
func ParseCedula(s string) (Cedula, error) {
	m := reCedula.FindStringSubmatch(strings.ToUpper(strings.TrimSpace(s)))
	if m == nil {
		return Cedula{}, fmt.Errorf("catalog: cédula %q must be PREFIJO-TOMO-PARTIDA (e.g. 8-123-456)", s)
	}
	c := Cedula{Prefijo: m[1], Tomo: m[2], Partida: m[3]}
	if c.EsProvincial() {
		if _, ok := provinciaByID[c.Prefijo]; !ok {
			return Cedula{}, fmt.Errorf("catalog: cédula province prefix %q out of range 1..13", c.Prefijo)
		}
	}
	return c, nil
}

// ValidateCedula reports whether s is a structurally valid cédula.
func ValidateCedula(s string) bool {
	_, err := ParseCedula(s)
	return err == nil
}

// DescribePrefijo returns a human-readable description of a cédula prefix.
func DescribePrefijo(prefijo string) string {
	prefijo = strings.ToUpper(prefijo)
	if d, ok := cedulaPrefijos[prefijo]; ok {
		return d
	}
	if p, ok := provinciaByID[prefijo]; ok {
		return "Inscrito en " + p.Nombre
	}
	return "Prefijo desconocido"
}
