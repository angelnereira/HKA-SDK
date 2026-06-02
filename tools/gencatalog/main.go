// Command gencatalog regenerates the embedded geographic catalog
// (catalog/data/ubicaciones.json) from an official, normalized CSV export.
//
// Why a CSV? The DGI "Anexos Técnicos" PDF is a scanned document and the INEC
// portal publishes the División Político-Administrativa in several formats. To
// keep regeneration reproducible and source-agnostic, this tool consumes a CSV
// with one row per corregimiento and the following columns (header required):
//
//	provincia_codigo,provincia_nombre,prefijo_cedula,distrito_codigo,distrito_nombre,corregimiento_codigo,corregimiento_nombre
//
// Reliable sources to build that CSV (see docs/CATALOGS.md):
//   - INEC — https://www.inec.gob.pa/ (División Político-Administrativa)
//   - Catálogo unificado (Alanube) — https://developer.alanube.co/v1.0-PAN/reference/catalogo-unificado-de-provincias-distritos-y-corregimientos
//   - Datasets comunitarios en GitHub (verificar contra INEC antes de usar)
//
// Usage:
//
//	go run ./tools/gencatalog -csv path/to/ubicaciones.csv -out catalog/data/ubicaciones.json
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type corregimiento struct {
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
}

type distrito struct {
	Codigo         string          `json:"codigo"`
	Nombre         string          `json:"nombre"`
	Corregimientos []corregimiento `json:"corregimientos"`
}

type provincia struct {
	Codigo        string     `json:"codigo"`
	Nombre        string     `json:"nombre"`
	PrefijoCedula string     `json:"prefijoCedula"`
	Distritos     []distrito `json:"distritos"`
}

type output struct {
	Meta       map[string]any `json:"_meta"`
	Provincias []provincia    `json:"provincias"`
}

func main() {
	csvPath := flag.String("csv", "", "path to normalized ubicaciones CSV (required)")
	outPath := flag.String("out", "catalog/data/ubicaciones.json", "output JSON path")
	flag.Parse()

	if *csvPath == "" {
		fmt.Fprintln(os.Stderr, "error: -csv is required")
		flag.Usage()
		os.Exit(2)
	}

	rows, err := readCSV(*csvPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading CSV:", err)
		os.Exit(1)
	}

	provs := build(rows)
	out := output{
		Meta: map[string]any{
			"description":    "Catálogo de ubicaciones de Panamá (provincia-distrito-corregimiento) para el campo codigoUbicacion.",
			"formato_codigo": "provincia-distrito-corregimiento (ej. 8-8-7)",
			"fuente_oficial": "INEC / Contraloría — División Político-Administrativa",
			"generado_por":   "tools/gencatalog",
		},
		Provincias: provs,
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error encoding JSON:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(*outPath, append(data, '\n'), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "error writing output:", err)
		os.Exit(1)
	}

	var nDist, nCorr int
	for _, p := range provs {
		nDist += len(p.Distritos)
		for _, d := range p.Distritos {
			nCorr += len(d.Corregimientos)
		}
	}
	fmt.Printf("wrote %s: %d provincias, %d distritos, %d corregimientos\n", *outPath, len(provs), nDist, nCorr)
}

type row struct {
	provCod, provNom, prefijo, distCod, distNom, corrCod, corrNom string
}

func readCSV(path string) ([]row, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) < 2 {
		return nil, fmt.Errorf("CSV has no data rows")
	}

	idx := map[string]int{}
	for i, h := range records[0] {
		idx[h] = i
	}
	required := []string{"provincia_codigo", "provincia_nombre", "prefijo_cedula", "distrito_codigo", "distrito_nombre", "corregimiento_codigo", "corregimiento_nombre"}
	for _, col := range required {
		if _, ok := idx[col]; !ok {
			return nil, fmt.Errorf("missing required column %q", col)
		}
	}

	out := make([]row, 0, len(records)-1)
	for _, rec := range records[1:] {
		out = append(out, row{
			provCod: rec[idx["provincia_codigo"]],
			provNom: rec[idx["provincia_nombre"]],
			prefijo: rec[idx["prefijo_cedula"]],
			distCod: rec[idx["distrito_codigo"]],
			distNom: rec[idx["distrito_nombre"]],
			corrCod: rec[idx["corregimiento_codigo"]],
			corrNom: rec[idx["corregimiento_nombre"]],
		})
	}
	return out, nil
}

func build(rows []row) []provincia {
	provOrder := []string{}
	provMap := map[string]*provincia{}
	distMap := map[string]*distrito{} // key: provCod|distCod

	for _, r := range rows {
		p, ok := provMap[r.provCod]
		if !ok {
			p = &provincia{Codigo: r.provCod, Nombre: r.provNom, PrefijoCedula: r.prefijo}
			provMap[r.provCod] = p
			provOrder = append(provOrder, r.provCod)
		}
		dk := r.provCod + "|" + r.distCod
		d, ok := distMap[dk]
		if !ok {
			d = &distrito{Codigo: r.distCod, Nombre: r.distNom}
			distMap[dk] = d
		}
		d.Corregimientos = append(d.Corregimientos, corregimiento{Codigo: r.corrCod, Nombre: r.corrNom})
	}

	// Assemble provincias with their distritos in numeric order.
	out := make([]provincia, 0, len(provOrder))
	for _, pc := range provOrder {
		p := provMap[pc]
		dl := []distrito{}
		seen := map[string]bool{}
		for _, r := range rows {
			if r.provCod != pc {
				continue
			}
			if seen[r.distCod] {
				continue
			}
			seen[r.distCod] = true
			dl = append(dl, *distMap[pc+"|"+r.distCod])
		}
		sort.SliceStable(dl, func(i, j int) bool { return numLess(dl[i].Codigo, dl[j].Codigo) })
		p.Distritos = dl
		out = append(out, *p)
	}
	sort.SliceStable(out, func(i, j int) bool { return numLess(out[i].Codigo, out[j].Codigo) })
	return out
}

func numLess(a, b string) bool {
	ai, errA := strconv.Atoi(a)
	bi, errB := strconv.Atoi(b)
	if errA == nil && errB == nil {
		return ai < bi
	}
	return a < b
}
