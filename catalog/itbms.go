package catalog

import (
	"strings"

	"github.com/angelnereira/hka-sdk/types"
)

// ITBMS (Impuesto a la Transferencia de Bienes Muebles y Servicios) rates.
//
// Panama applies three positive rates plus an exempt category. The applicable
// rate is determined by the *kind of good or service*, not by the customer:
//
//   - 7%  — tasa general: most goods and services.
//   - 10% — bebidas alcohólicas (importación y venta) y servicios de hospedaje/alojamiento.
//   - 15% — productos derivados del tabaco.
//   - 0%/Exento — bienes y servicios exentos por ley (p. ej. ciertos alimentos,
//     medicamentos, productos de canasta básica, educación, transporte).
//
// Sources: DGI (https://dgi.mef.gob.pa/itbms/Generalidades) and the ITBMS rate
// schedule. These categories are authoritative; the keyword classifier below is a
// best-effort *suggestion* and must not be treated as a legal determination.

// ITBMSCategoria is a semantic ITBMS rate category.
type ITBMSCategoria int

const (
	ITBMSGeneral         ITBMSCategoria = iota // 7%
	ITBMSAlcohol                               // 10% — bebidas alcohólicas
	ITBMSHospedaje                             // 10% — hospedaje/alojamiento
	ITBMSTabaco                                // 15% — derivados del tabaco
	ITBMSExentoCategoria                       // exento
)

// Tasa returns the SDK TasaITBMS code for the category.
func (c ITBMSCategoria) Tasa() types.TasaITBMS {
	switch c {
	case ITBMSAlcohol, ITBMSHospedaje:
		return types.ITBMS10
	case ITBMSTabaco:
		return types.ITBMS15
	case ITBMSExentoCategoria:
		return types.ITBMSExento
	default:
		return types.ITBMS7
	}
}

// Porcentaje returns the rate as a percentage (7, 10, 15 or 0).
func (c ITBMSCategoria) Porcentaje() int {
	switch c.Tasa() {
	case types.ITBMS7:
		return 7
	case types.ITBMS10:
		return 10
	case types.ITBMS15:
		return 15
	default:
		return 0
	}
}

// Descripcion returns a human-readable description of the category.
func (c ITBMSCategoria) Descripcion() string {
	switch c {
	case ITBMSAlcohol:
		return "Bebidas alcohólicas — 10%"
	case ITBMSHospedaje:
		return "Servicio de hospedaje/alojamiento — 10%"
	case ITBMSTabaco:
		return "Productos derivados del tabaco — 15%"
	case ITBMSExentoCategoria:
		return "Bien o servicio exento — 0%"
	default:
		return "Tasa general — 7%"
	}
}

// PorcentajeDeTasa returns the percentage for a TasaITBMS code.
func PorcentajeDeTasa(t types.TasaITBMS) int {
	switch t {
	case types.ITBMS7:
		return 7
	case types.ITBMS10:
		return 10
	case types.ITBMS15:
		return 15
	default:
		return 0
	}
}

var (
	kwTabaco    = []string{"cigarrillo", "cigarro", "tabaco", "puro", "cigarette", "vapeo", "vape"}
	kwAlcohol   = []string{"cerveza", "vino", "ron", "whisky", "whiskey", "vodka", "licor", "aguardiente", "ginebra", "tequila", "champ", "alcohol", "seco ", "bebida alcoh"}
	kwHospedaje = []string{"hospedaje", "alojamiento", "hotel", "hostal", "habitaci", "pernocta", "lodging"}
)

// SugerirCategoria suggests an ITBMS category from a free-text product description.
// It is a heuristic aid for data entry only — the emisor remains responsible for
// applying the correct legal rate. When no special keyword matches it returns the
// general 7% category. It never infers the exempt category (which is law-driven).
func SugerirCategoria(descripcion string) ITBMSCategoria {
	d := strings.ToLower(descripcion)
	if containsAny(d, kwTabaco) {
		return ITBMSTabaco
	}
	if containsAny(d, kwAlcohol) {
		return ITBMSAlcohol
	}
	if containsAny(d, kwHospedaje) {
		return ITBMSHospedaje
	}
	return ITBMSGeneral
}

// SugerirTasa is a convenience wrapper returning the suggested TasaITBMS directly.
func SugerirTasa(descripcion string) types.TasaITBMS {
	return SugerirCategoria(descripcion).Tasa()
}

func containsAny(s string, kws []string) bool {
	for _, k := range kws {
		if strings.Contains(s, k) {
			return true
		}
	}
	return false
}
