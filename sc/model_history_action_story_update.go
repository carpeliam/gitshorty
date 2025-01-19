/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc

// An action representing a Story being updated.
type HistoryActionStoryUpdate struct {
	// The ID of the entity referenced.
	Id int64 `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The action of the entity referenced.
	Action string `json:"action"`
	// The application URL of the Story.
	AppUrl string `json:"app_url"`
	Changes *HistoryChangesStory `json:"changes,omitempty"`
	// The name of the Story.
	Name string `json:"name"`
	// The type of Story; either feature, bug, or chore.
	StoryType string `json:"story_type"`
}
