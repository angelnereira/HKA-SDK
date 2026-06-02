package catalog_test

import (
	"testing"

	"github.com/angelnereira/hka-sdk/catalog"
	"github.com/angelnereira/hka-sdk/types"
)

func TestProvincias(t *testing.T) {
	ps := catalog.Provincias()
	if len(ps) != 13 {
		t.Fatalf("expected 13 provincias (incl. comarcas), got %d", len(ps))
	}
	p, ok := catalog.ProvinciaByCodigo("8")
	if !ok || p.Nombre != "Panamá" {
		t.Fatalf("provincia 8 = %+v, ok=%v", p, ok)
	}
	if p.PrefijoCedula != "8" {
		t.Errorf("prefijoCedula = %q, want 8", p.PrefijoCedula)
	}
}

func TestParseUbicacion(t *testing.T) {
	u, err := catalog.ParseUbicacion("8-8-7")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Provincia != "8" || u.Distrito != "8" || u.Corregimiento != "7" {
		t.Errorf("parsed = %+v", u)
	}
	prov, _, corr, ok := u.Resolve()
	if !ok || prov != "Panamá" || corr != "Bella Vista" {
		t.Errorf("resolve = %q/%q ok=%v", prov, corr, ok)
	}

	for _, bad := range []string{"", "8-8", "99-1-1", "a-b-c", "8/8/7"} {
		if catalog.ValidateUbicacion(bad) {
			t.Errorf("expected %q to be invalid", bad)
		}
	}
}

func TestParseCedula(t *testing.T) {
	c, err := catalog.ParseCedula("8-123-456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !c.EsProvincial() || c.Prefijo != "8" {
		t.Errorf("parsed = %+v", c)
	}
	if _, err := catalog.ParseCedula("E-50-123"); err != nil {
		t.Errorf("E-prefix cédula should be valid: %v", err)
	}
	for _, bad := range []string{"", "14-1-1", "8-1", "ZZ-1-1"} {
		if catalog.ValidateCedula(bad) {
			t.Errorf("expected %q to be invalid", bad)
		}
	}
}

func TestParseRUC(t *testing.T) {
	r, err := catalog.ParseRUC("155596713-2-2015")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Kind != catalog.RUCJuridico {
		t.Errorf("kind = %v, want jurídico", r.Kind)
	}
	r2, err := catalog.ParseRUC("8-123-456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r2.Kind != catalog.RUCNatural || r2.Cedula == nil {
		t.Errorf("natural RUC = %+v", r2)
	}
	if catalog.ValidateRUC("not-a-ruc!") {
		t.Error("expected invalid RUC to fail")
	}
}

func TestCUFE(t *testing.T) {
	valid := "FE" + repeat("0", 64)
	if !catalog.ValidateCUFE(valid) {
		t.Errorf("expected 66-char FE string to be valid")
	}
	if catalog.ValidateCUFE("FE123") {
		t.Error("short CUFE should be invalid")
	}
	if len(valid) != catalog.CUFELength {
		t.Errorf("len = %d, want %d", len(valid), catalog.CUFELength)
	}
}

func TestParseCUFE(t *testing.T) {
	// real-world example: factura interna (01) emitted in the test environment (2)
	cufe := "FE0120000155596713-2-2015-5900012019052800055000155650121566749040"
	if len(cufe) != catalog.CUFELength {
		t.Fatalf("example CUFE len = %d, want 66", len(cufe))
	}
	info, err := catalog.ParseCUFE(cufe)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if info.TipoDocumento != types.TipoDocFacturaInterna {
		t.Errorf("TipoDocumento = %q, want 01", info.TipoDocumento)
	}
	if info.Ambiente != catalog.AmbientePruebas {
		t.Errorf("Ambiente = %q, want pruebas (2)", info.Ambiente)
	}
	if _, err := catalog.ParseCUFE("FE123"); err == nil {
		t.Error("short CUFE should fail to parse")
	}
}

func TestDescribeFormatoCAFE(t *testing.T) {
	if got := catalog.DescribeFormatoCAFE(types.CAFECintaPapel); got == "" || got == "Formato CAFE desconocido" {
		t.Errorf("unexpected description for cinta papel: %q", got)
	}
	if got := catalog.DescribeFormatoCAFE(types.CAFEPapelCarta); got == "" || got == "Formato CAFE desconocido" {
		t.Errorf("unexpected description for papel carta: %q", got)
	}
}

func TestITBMSClassifier(t *testing.T) {
	cases := []struct {
		desc string
		want types.TasaITBMS
	}{
		{"Cerveza nacional", types.ITBMS10},
		{"Botella de whisky", types.ITBMS10},
		{"Noche de hospedaje en hotel", types.ITBMS10},
		{"Cartón de cigarrillos", types.ITBMS15},
		{"Servicio de consultoría", types.ITBMS7},
		{"Tornillos de acero", types.ITBMS7},
	}
	for _, tc := range cases {
		if got := catalog.SugerirTasa(tc.desc); got != tc.want {
			t.Errorf("SugerirTasa(%q) = %q, want %q", tc.desc, got, tc.want)
		}
	}
	if catalog.PorcentajeDeTasa(types.ITBMS15) != 15 {
		t.Error("ITBMS15 should be 15%")
	}
}

func TestCPBS(t *testing.T) {
	if err := catalog.ValidateCPBS("13", "1310"); err != nil {
		t.Errorf("13/1310 should be valid: %v", err)
	}
	if err := catalog.ValidateCPBS("72", "1310"); err == nil {
		t.Error("mismatched abrev/codigo should fail")
	}
	if err := catalog.ValidateCPBS("1", "1310"); err == nil {
		t.Error("1-digit abrev should fail")
	}
	abrev, err := catalog.AbrevForCPBS("7210")
	if err != nil || abrev != "72" {
		t.Errorf("AbrevForCPBS(7210) = %q, %v", abrev, err)
	}
	if p, ok := catalog.CPBSByCodigo("1310"); !ok || p.Nombre != "Muebles" {
		t.Errorf("CPBSByCodigo(1310) = %+v ok=%v", p, ok)
	}
}

func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
