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

// A Comment is any note added within the Comment field of a Story.
type StoryComment struct {
	// The Shortcut application url for the Comment.
	AppUrl string `json:"app_url"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// True/false boolean indicating whether the Comment has been deleted.
	Deleted bool `json:"deleted"`
	// The ID of the Story on which the Comment appears.
	StoryId int64 `json:"story_id"`
	// `Deprecated:` use `member_mention_ids`.
	MentionIds []string `json:"mention_ids"`
	// The unique ID of the Member who is the Comment's author.
	AuthorId string `json:"author_id"`
	// The unique IDs of the Member who are mentioned in the Comment.
	MemberMentionIds []string `json:"member_mention_ids"`
	// Marks the comment as a blocker that can be surfaced to permissions or teams mentioned in the comment. Can only be used on a top-level comment.
	Blocker bool `json:"blocker,omitempty"`
	// Whether the Comment is currently the root of a thread that is linked to Slack.
	LinkedToSlack bool `json:"linked_to_slack"`
	// The time/date when the Comment was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The unique IDs of the Group who are mentioned in the Comment.
	GroupMentionIds []string `json:"group_mention_ids"`
	// This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The ID of the parent Comment this Comment is threaded under.
	ParentId int64 `json:"parent_id,omitempty"`
	// The unique ID of the Comment.
	Id int64 `json:"id"`
	// The Comments numerical position in the list from oldest to newest.
	Position int64 `json:"position"`
	// Marks the comment as an unblocker to its  blocker parent. Can only be set on a threaded comment who has a parent with `blocker` set.
	UnblocksParent bool `json:"unblocks_parent,omitempty"`
	// A set of Reactions to this Comment.
	Reactions []StoryReaction `json:"reactions"`
	// The time/date when the Comment was created.
	CreatedAt time.Time `json:"created_at"`
	// The text of the Comment. In the case that the Comment has been deleted, this field can be set to nil.
	Text string `json:"text"`
}
