package validate

import "github.com/angelnereira/hka-sdk/types"

// Totals validates a TotalesSubTotales structure against the provided items slice.
func Totals(t *types.TotalesSubTotales, items []types.Item) *ValidationError {
	ve := &ValidationError{}
	validateTotales(ve, t, items)
	if ve.HasErrors() {
		return ve
	}
	return nil
}
