/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc
import (
	"time"
)

// A Label can be used to associate and filter Stories and Epics, and also create new Workspaces.
type Label struct {
	// The Shortcut application url for the Label.
	AppUrl string `json:"app_url"`
	// The description of the Label.
	Description string `json:"description"`
	// A true/false boolean indicating if the Label has been archived.
	Archived bool `json:"archived"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The hex color to be displayed with the Label (for example, \"#ff0000\").
	Color string `json:"color"`
	// The name of the Label.
	Name string `json:"name"`
	GlobalId string `json:"global_id"`
	// The time/date that the Label was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// This field can be set to another unique ID. In the case that the Label has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The unique ID of the Label.
	Id int64 `json:"id"`
	Stats *LabelStats `json:"stats,omitempty"`
	// The time/date that the Label was created.
	CreatedAt time.Time `json:"created_at"`
}
