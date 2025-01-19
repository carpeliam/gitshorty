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

// A map of story attributes this template populates.
type CreateStoryContents struct {
	// The description of the story.
	Description string `json:"description,omitempty"`
	// An array of labels to be populated by the template.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type,omitempty"`
	// An array of maps specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField.
	CustomFields []CustomFieldValueParams `json:"custom_fields,omitempty"`
	// An array of the attached file IDs to be populated.
	FileIds []int64 `json:"file_ids,omitempty"`
	// The name of the story.
	Name string `json:"name,omitempty"`
	// The ID of the epic the to be populated.
	EpicId int64 `json:"epic_id,omitempty"`
	// An array of external links to be populated.
	ExternalLinks []string `json:"external_links,omitempty"`
	// The ID of the iteration the to be populated.
	IterationId int64 `json:"iteration_id,omitempty"`
	// An array of tasks to be populated by the template.
	Tasks []BaseTaskParams `json:"tasks,omitempty"`
	// The ID of the group to be populated.
	GroupId string `json:"group_id,omitempty"`
	// The ID of the workflow state to be populated.
	WorkflowStateId int64 `json:"workflow_state_id,omitempty"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids,omitempty"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids,omitempty"`
	// The numeric point estimate to be populated.
	Estimate int64 `json:"estimate,omitempty"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id,omitempty"`
	// An array of the linked file IDs to be populated.
	LinkedFileIds []int64 `json:"linked_file_ids,omitempty"`
	// The due date of the story.
	Deadline time.Time `json:"deadline,omitempty"`
}
