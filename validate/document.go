package validate

import "github.com/angelnereira/hka-sdk/types"

// DocumentFull runs all validation rules against a DocumentoElectronico.
// It is a convenience alias kept for clarity; the core logic lives in rules.go.
func DocumentFull(doc *types.DocumentoElectronico) *ValidationError {
	return Document(doc)
}
