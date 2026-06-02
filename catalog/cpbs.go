package catalog

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

//go:embed data/cpbs.json
var cpbsJSON []byte

// CPBS (Catálogo de Productos, Bienes y Servicios).
//
// Two fields reference it on each item and are mandatory when selling to a
// government client (tipoClienteFE = 03):
//
//   - codigoCPBSAbrev: 2-digit category/division.
//   - codigoCPBS: 4-digit code whose first two digits equal codigoCPBSAbrev.
//
// The catalog derives from the DGI Anexos Técnicos (based on the CPC). The embedded
// data is a verified sample; regenerate the full catalog with tools/gencatalog.

// CPBSCategoria is a 2-digit CPBS category and its products.
type CPBSCategoria struct {
	Abrev     string        `json:"abrev"`
	Nombre    string        `json:"nombre"`
	Productos []CPBSProducto `json:"productos"`
}

// CPBSProducto is a 4-digit CPBS code and its name.
type CPBSProducto struct {
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
}

type cpbsFile struct {
	Categorias []CPBSCategoria `json:"categorias"`
}

var (
	cpbsCategorias []CPBSCategoria
	cpbsByAbrev    = map[string]*CPBSCategoria{}
	cpbsByCodigo   = map[string]CPBSProducto{}
)

func init() {
	var f cpbsFile
	if err := json.Unmarshal(cpbsJSON, &f); err != nil {
		panic("catalog: invalid embedded cpbs.json: " + err.Error())
	}
	cpbsCategorias = f.Categorias
	for i := range cpbsCategorias {
		cpbsByAbrev[cpbsCategorias[i].Abrev] = &cpbsCategorias[i]
		for _, p := range cpbsCategorias[i].Productos {
			cpbsByCodigo[p.Codigo] = p
		}
	}
}

var (
	reCPBSAbrev = regexp.MustCompile(`^\d{2}$`)
	reCPBS      = regexp.MustCompile(`^\d{4}$`)
)

// CPBSCategorias returns the catalog categories.
func CPBSCategorias() []CPBSCategoria { return cpbsCategorias }

// CPBSByCodigo returns the product entry for a 4-digit CPBS code, if known.
func CPBSByCodigo(codigo string) (CPBSProducto, bool) {
	p, ok := cpbsByCodigo[codigo]
	return p, ok
}

// AbrevForCPBS returns the 2-digit codigoCPBSAbrev implied by a 4-digit codigoCPBS.
func AbrevForCPBS(codigoCPBS string) (string, error) {
	if !reCPBS.MatchString(codigoCPBS) {
		return "", fmt.Errorf("catalog: codigoCPBS %q must be 4 digits", codigoCPBS)
	}
	return codigoCPBS[:2], nil
}

// ValidateCPBS checks that codigoCPBSAbrev and codigoCPBS are well-formed and
// consistent (the abbreviated code must be the 2-digit prefix of the full code).
func ValidateCPBS(abrev, codigoCPBS string) error {
	abrev = strings.TrimSpace(abrev)
	codigoCPBS = strings.TrimSpace(codigoCPBS)
	if !reCPBSAbrev.MatchString(abrev) {
		return fmt.Errorf("catalog: codigoCPBSAbrev %q must be 2 digits", abrev)
	}
	if !reCPBS.MatchString(codigoCPBS) {
		return fmt.Errorf("catalog: codigoCPBS %q must be 4 digits", codigoCPBS)
	}
	if codigoCPBS[:2] != abrev {
		return fmt.Errorf("catalog: codigoCPBSAbrev %q must be the first two digits of codigoCPBS %q", abrev, codigoCPBS)
	}
	return nil
}
