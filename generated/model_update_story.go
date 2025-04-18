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

type UpdateStory struct {
	// The description of the story.
	Description string `json:"description,omitempty"`
	// True if the story is archived, otherwise false.
	Archived bool `json:"archived,omitempty"`
	// An array of labels attached to the story.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// An array of IDs of Pull/Merge Requests attached to the story.
	PullRequestIds []int64 `json:"pull_request_ids,omitempty"`
	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type,omitempty"`
	// A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField.
	CustomFields []CustomFieldValueParams `json:"custom_fields,omitempty"`
	// One of \"first\" or \"last\". This can be used to move the given story to the first or last position in the workflow state.
	MoveTo string `json:"move_to,omitempty"`
	// An array of IDs of files attached to the story.
	FileIds []int64 `json:"file_ids,omitempty"`
	// A manual override for the time/date the Story was completed.
	CompletedAtOverride time.Time `json:"completed_at_override,omitempty"`
	// The title of the story.
	Name string `json:"name,omitempty"`
	// The ID of the epic the story belongs to.
	EpicId int64 `json:"epic_id,omitempty"`
	// An array of External Links associated with this story.
	ExternalLinks []string `json:"external_links,omitempty"`
	// An array of IDs of Branches attached to the story.
	BranchIds []int64 `json:"branch_ids,omitempty"`
	// An array of IDs of Commits attached to the story.
	CommitIds []int64 `json:"commit_ids,omitempty"`
	// The ID of the member that requested the story.
	RequestedById string `json:"requested_by_id,omitempty"`
	// The ID of the iteration the story belongs to.
	IterationId int64 `json:"iteration_id,omitempty"`
	// A manual override for the time/date the Story was started.
	StartedAtOverride time.Time `json:"started_at_override,omitempty"`
	// The ID of the group to associate with this story
	GroupId string `json:"group_id,omitempty"`
	// The ID of the workflow state to put the story in.
	WorkflowStateId int64 `json:"workflow_state_id,omitempty"`
	// An array of UUIDs of the followers of this story.
	FollowerIds []string `json:"follower_ids,omitempty"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids,omitempty"`
	// The ID of the story we want to move this story before.
	BeforeId int64 `json:"before_id,omitempty"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate int64 `json:"estimate,omitempty"`
	// The ID of the story we want to move this story after.
	AfterId int64 `json:"after_id,omitempty"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id,omitempty"`
	// An array of IDs of linked files attached to the story.
	LinkedFileIds []int64 `json:"linked_file_ids,omitempty"`
	// The due date of the story.
	Deadline time.Time `json:"deadline,omitempty"`
}
