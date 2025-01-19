/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc

// An action representing a Task being updated.
type HistoryActionTaskUpdate struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	Changes *HistoryChangesTask `json:"changes"`
	// Whether or not the Task is complete.
	Complete bool `json:"complete,omitempty"`
	// The description of the Task.
	Description string `json:"description"`
	// The Story ID that contains the Task.
	StoryId int64 `json:"story_id"`
}
