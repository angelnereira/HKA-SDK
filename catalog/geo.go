// Package catalog bundles the reference catalogs and code-format helpers that the
// HKA / DGI electronic-invoice scheme depends on: the geographic location catalog
// (provincia / distrito / corregimiento), ITBMS tax rates, the CPBS product/service
// catalog, and parsers/validators for cédula, RUC, and CUFE.
//
// The geographic and CPBS catalogs ship with authoritative top-level data embedded
// via go:embed and can be regenerated in full from their official sources with the
// generator in tools/gencatalog. See docs/CATALOGS.md for sources and refresh steps.
package catalog

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed data/ubicaciones.json
var ubicacionesJSON []byte

// Provincia is a top-level administrative division. PrefijoCedula is the leading
// segment used by cédula numbers registered in that province.
type Provincia struct {
	Codigo        string     `json:"codigo"`
	Nombre        string     `json:"nombre"`
	PrefijoCedula string     `json:"prefijoCedula"`
	Distritos     []Distrito `json:"distritos"`
}

// Distrito is a district within a province.
type Distrito struct {
	Codigo         string          `json:"codigo"`
	Nombre         string          `json:"nombre"`
	Corregimientos []Corregimiento `json:"corregimientos"`
}

// Corregimiento is the smallest administrative unit, referenced by codigoUbicacion.
type Corregimiento struct {
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
}

type ubicacionesFile struct {
	Provincias []Provincia `json:"provincias"`
}

var (
	provincias    []Provincia
	provinciaByID = map[string]*Provincia{}
)

func init() {
	var f ubicacionesFile
	if err := json.Unmarshal(ubicacionesJSON, &f); err != nil {
		panic("catalog: invalid embedded ubicaciones.json: " + err.Error())
	}
	provincias = f.Provincias
	for i := range provincias {
		provinciaByID[provincias[i].Codigo] = &provincias[i]
	}
}

// Provincias returns all provinces (and comarcas) of Panama.
func Provincias() []Provincia { return provincias }

// ProvinciaByCodigo returns the province with the given numeric code (1..13).
func ProvinciaByCodigo(codigo string) (*Provincia, bool) {
	p, ok := provinciaByID[codigo]
	return p, ok
}

// Ubicacion is a parsed codigoUbicacion value.
type Ubicacion struct {
	Provincia     string // numeric provincia code
	Distrito      string // numeric distrito code
	Corregimiento string // numeric corregimiento code
}

func (u Ubicacion) String() string {
	return u.Provincia + "-" + u.Distrito + "-" + u.Corregimiento
}

var reUbicacion = regexp.MustCompile(`^(\d{1,2})-(\d{1,3})-(\d{1,3})$`)

// ParseUbicacion parses a codigoUbicacion of the form "provincia-distrito-corregimiento"
// (e.g. "8-8-7"). It validates the shape and that the province code is known; it does
// not require the distrito/corregimiento to be present in the embedded sample data.
func ParseUbicacion(code string) (Ubicacion, error) {
	m := reUbicacion.FindStringSubmatch(strings.TrimSpace(code))
	if m == nil {
		return Ubicacion{}, fmt.Errorf("catalog: codigoUbicacion %q must be provincia-distrito-corregimiento (e.g. 8-8-7)", code)
	}
	u := Ubicacion{Provincia: m[1], Distrito: m[2], Corregimiento: m[3]}
	if n, _ := strconv.Atoi(u.Provincia); n < 1 || n > 13 {
		return Ubicacion{}, fmt.Errorf("catalog: provincia %q out of range 1..13", u.Provincia)
	}
	return u, nil
}

// ValidateUbicacion reports whether code is a structurally valid codigoUbicacion.
func ValidateUbicacion(code string) bool {
	_, err := ParseUbicacion(code)
	return err == nil
}

// Resolve returns the names of the provincia, distrito and corregimiento for a code,
// to the extent they are present in the embedded catalog. ok is false when the
// provincia is unknown. Empty distrito/corregimiento names mean the embedded sample
// does not yet include that entry (regenerate with tools/gencatalog for full data).
func (u Ubicacion) Resolve() (provincia, distrito, corregimiento string, ok bool) {
	p, found := provinciaByID[u.Provincia]
	if !found {
		return "", "", "", false
	}
	provincia = p.Nombre
	for i := range p.Distritos {
		if p.Distritos[i].Codigo == u.Distrito {
			distrito = p.Distritos[i].Nombre
			for j := range p.Distritos[i].Corregimientos {
				if p.Distritos[i].Corregimientos[j].Codigo == u.Corregimiento {
					corregimiento = p.Distritos[i].Corregimientos[j].Nombre
				}
			}
		}
	}
	return provincia, distrito, corregimiento, true
}
