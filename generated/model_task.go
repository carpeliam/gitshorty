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

// A Task on a Story.
type Task struct {
	// Full text of the Task.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique identifier of the parent Story.
	StoryId int64 `json:"story_id"`
	// `Deprecated:` use `member_mention_ids`.
	MentionIds []string `json:"mention_ids"`
	// An array of UUIDs of Members mentioned in this Task.
	MemberMentionIds []string `json:"member_mention_ids"`
	// The time/date the Task was completed.
	CompletedAt time.Time `json:"completed_at"`
	// The time/date the Task was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// An array of UUIDs of Groups mentioned in this Task.
	GroupMentionIds []string `json:"group_mention_ids"`
	// An array of UUIDs of the Owners of this Task.
	OwnerIds []string `json:"owner_ids"`
	// This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The unique ID of the Task.
	Id int64 `json:"id"`
	// The number corresponding to the Task's position within a list of Tasks on a Story.
	Position int64 `json:"position"`
	// True/false boolean indicating whether the Task has been completed.
	Complete bool `json:"complete"`
	// The time/date the Task was created.
	CreatedAt time.Time `json:"created_at"`
}
