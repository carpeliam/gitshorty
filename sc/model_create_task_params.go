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

// Request parameters for creating a Task on a Story.
type CreateTaskParams struct {
	// The Task description.
	Description string `json:"description"`
	// True/false boolean indicating whether the Task is completed. Defaults to false.
	Complete bool `json:"complete,omitempty"`
	// An array of UUIDs for any members you want to add as Owners on this new Task.
	OwnerIds []string `json:"owner_ids,omitempty"`
	// This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id,omitempty"`
	// Defaults to the time/date the Task is created but can be set to reflect another creation time/date.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Defaults to the time/date the Task is created in Shortcut but can be set to reflect another time/date.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
