/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

// The Identity of the VCS user that authored the Commit.
type Identity struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// This is your login in VCS.
	Name string `json:"name"`
	// The service this Identity is for.
	Type_ string `json:"type"`
}
