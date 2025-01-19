/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc

// An action representing a Story Comment being created.
type HistoryActionStoryCommentCreate struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The application URL of the Story Comment.
	AppUrl string `json:"app_url"`
	// The text of the Story Comment.
	Text string `json:"text"`
	// The Member ID of who created the Story Comment.
	AuthorId string `json:"author_id"`
}
