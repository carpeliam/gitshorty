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

// Comments associated with Epic Discussions.
type ThreadedComment struct {
	// The Shortcut application url for the Comment.
	AppUrl string `json:"app_url"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// True/false boolean indicating whether the Comment is deleted.
	Deleted bool `json:"deleted"`
	// `Deprecated:` use `member_mention_ids`.
	MentionIds []string `json:"mention_ids"`
	// The unique ID of the Member that authored the Comment.
	AuthorId string `json:"author_id"`
	// An array of Member IDs that have been mentioned in this Comment.
	MemberMentionIds []string `json:"member_mention_ids"`
	// A nested array of threaded comments.
	Comments []ThreadedComment `json:"comments"`
	// The time/date the Comment was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// An array of Group IDs that have been mentioned in this Comment.
	GroupMentionIds []string `json:"group_mention_ids"`
	// This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The unique ID of the Comment.
	Id int64 `json:"id"`
	// The time/date the Comment was created.
	CreatedAt time.Time `json:"created_at"`
	// The text of the Comment.
	Text string `json:"text"`
}
