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

// Projects typically map to teams (such as Frontend, Backend, Mobile, Devops, etc) but can represent any open-ended product, component, or initiative.
type Project struct {
	// The Shortcut application url for the Project.
	AppUrl string `json:"app_url"`
	// The description of the Project.
	Description string `json:"description"`
	// True/false boolean indicating whether the Project is in an Archived state.
	Archived bool `json:"archived"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The number of days before the thermometer appears in the Story summary.
	DaysToThermometer int64 `json:"days_to_thermometer"`
	// The color associated with the Project in the Shortcut member interface.
	Color string `json:"color"`
	// The ID of the workflow the project belongs to.
	WorkflowId int64 `json:"workflow_id"`
	// The name of the Project
	Name string `json:"name"`
	// The Global ID of the Project.
	GlobalId string `json:"global_id"`
	// The date at which the Project was started.
	StartTime time.Time `json:"start_time"`
	// The time/date that the Project was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
	// This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The unique ID of the Project.
	Id int64 `json:"id"`
	// Configuration to enable or disable thermometers in the Story summary.
	ShowThermometer bool `json:"show_thermometer"`
	// The ID of the team the project belongs to.
	TeamId int64 `json:"team_id"`
	// The number of weeks per iteration in this Project.
	IterationLength int64 `json:"iteration_length"`
	// The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation string `json:"abbreviation"`
	Stats *ProjectStats `json:"stats"`
	// The time/date that the Project was created.
	CreatedAt time.Time `json:"created_at"`
}
