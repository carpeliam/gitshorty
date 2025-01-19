/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc

// A reference to a VCS Commit.
type HistoryReferenceCommit struct {
	// The ID of the entity referenced.
	Id *interface{} `json:"id"`
	// The type of entity referenced.
	EntityType string `json:"entity_type"`
	// The message from the Commit.
	Message string `json:"message"`
	// The external URL for the Branch.
	Url string `json:"url"`
}
