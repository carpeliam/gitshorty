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

type CreateCommentComment struct {
	// The comment text.
	Text string `json:"text"`
	// The Member ID of the Comment's author. Defaults to the user identified by the API token.
	AuthorId string `json:"author_id,omitempty"`
	// Defaults to the time/date the comment is created, but can be set to reflect another date.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Defaults to the time/date the comment is last updated, but can be set to reflect another date.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// This field can be set to another unique ID. In the case that the comment has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id,omitempty"`
}
