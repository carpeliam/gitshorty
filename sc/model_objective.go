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

// An Objective is a collection of Epics that represent a release or some other large initiative that you are working on.
type Objective struct {
	// The Shortcut application url for the Objective.
	AppUrl string `json:"app_url"`
	// The Objective's description.
	Description string `json:"description"`
	// A boolean indicating whether the Objective has been archived or not.
	Archived bool `json:"archived"`
	// A true/false boolean indicating if the Objective has been started.
	Started bool `json:"started"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// A manual override for the time/date the Objective was completed.
	CompletedAtOverride time.Time `json:"completed_at_override"`
	// The time/date the Objective was started.
	StartedAt time.Time `json:"started_at"`
	// The time/date the Objective was completed.
	CompletedAt time.Time `json:"completed_at"`
	// The name of the Objective.
	Name string `json:"name"`
	GlobalId string `json:"global_id"`
	// A true/false boolean indicating if the Objectivehas been completed.
	Completed bool `json:"completed"`
	// The workflow state that the Objective is in.
	State string `json:"state"`
	// A manual override for the time/date the Objective was started.
	StartedAtOverride time.Time `json:"started_at_override"`
	// The time/date the Objective was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// An array of Categories attached to the Objective.
	Categories []Category `json:"categories"`
	// The unique ID of the Objective.
	Id int64 `json:"id"`
	// The IDs of the Key Results associated with the Objective.
	KeyResultIds []string `json:"key_result_ids"`
	// A number representing the position of the Objective in relation to every other Objective within the Workspace.
	Position int64 `json:"position"`
	Stats *ObjectiveStats `json:"stats"`
	// The time/date the Objective was created.
	CreatedAt time.Time `json:"created_at"`
}
