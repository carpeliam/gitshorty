/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

type CustomFieldEnumValue struct {
	// The unique public ID for the Custom Field.
	Id string `json:"id"`
	// A string value within the domain of this Custom Field.
	Value string `json:"value"`
	// An integer indicating the position of this Value with respect to the other CustomFieldEnumValues in the enumeration.
	Position int64 `json:"position"`
	// A color key associated with this CustomFieldEnumValue.
	ColorKey string `json:"color_key"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// When true, the CustomFieldEnumValue can be selected for the CustomField.
	Enabled bool `json:"enabled"`
}
