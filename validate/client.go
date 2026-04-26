package validate

import "github.com/angelnereira/hka-sdk/types"

// Client validates only the Cliente sub-structure in isolation.
// Useful for pre-validating client data before building the full document.
func Client(c *types.Cliente) *ValidationError {
	ve := &ValidationError{}
	validateCliente(ve, c)
	if ve.HasErrors() {
		return ve
	}
	return nil
}
