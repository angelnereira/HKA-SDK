package validate

import "github.com/angelnereira/hka-sdk/types"

// Items validates a slice of items in isolation, using the provided client type
// to check government-specific rules.
func Items(items []types.Item, tipoCliente types.TipoClienteFE) *ValidationError {
	ve := &ValidationError{}
	validateItems(ve, items, tipoCliente)
	if ve.HasErrors() {
		return ve
	}
	return nil
}
