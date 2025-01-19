/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc

// The starting value of the Key Result.
type KeyResultValue struct {
	// The numeric value, as a decimal string. No more than two decimal places are allowed.
	NumericValue string `json:"numeric_value,omitempty"`
	// The boolean value.
	BooleanValue bool `json:"boolean_value,omitempty"`
}
