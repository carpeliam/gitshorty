/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

type CustomFieldValueParams struct {
	// The unique public ID for the CustomField.
	FieldId string `json:"field_id"`
	// The unique public ID for the CustomFieldEnumValue.
	ValueId string `json:"value_id"`
	// A literal value for the CustomField. Currently ignored.
	Value string `json:"value,omitempty"`
}
