/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc

// The synced item for the story.
type SyncedItem struct {
	// The id used to reference an external entity.
	ExternalId string `json:"external_id"`
	// The url to the external entity.
	Url string `json:"url"`
}
