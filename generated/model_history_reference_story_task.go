/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

// A reference to a Story Task.
type HistoryReferenceStoryTask struct {
	// The ID of the entity referenced.
	Id *interface{} `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The description of the Story Task.
	Description string `json:"description"`
}
