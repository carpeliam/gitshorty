/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated
import (
	"time"
)

type CustomField struct {
	// A string description of the CustomField
	Description string `json:"description,omitempty"`
	// A string that represents the icon that corresponds to this custom field.
	IconSetIdentifier string `json:"icon_set_identifier,omitempty"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The types of stories this CustomField is scoped to.
	StoryTypes []string `json:"story_types,omitempty"`
	// The name of the Custom Field.
	Name string `json:"name"`
	// When true, the CustomFieldEnumValues may not be reordered.
	FixedPosition bool `json:"fixed_position,omitempty"`
	// The instant when this CustomField was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The unique public ID for the CustomField.
	Id string `json:"id"`
	// A collection of legal values for a CustomField.
	Values []CustomFieldEnumValue `json:"values,omitempty"`
	// The type of Custom Field, eg. 'enum'.
	FieldType string `json:"field_type"`
	// An integer indicating the position of this Custom Field with respect to the other CustomField
	Position int64 `json:"position"`
	// The canonical name for a Shortcut-defined field.
	CanonicalName string `json:"canonical_name,omitempty"`
	// When true, the CustomField can be applied to entities in the Workspace.
	Enabled bool `json:"enabled"`
	// The instant when this CustomField was created.
	CreatedAt time.Time `json:"created_at"`
}
