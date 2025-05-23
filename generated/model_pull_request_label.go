/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

// Corresponds to a VCS Label associated with a Pull Request.
type PullRequestLabel struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the VCS Label.
	Id int64 `json:"id"`
	// The color of the VCS label.
	Color string `json:"color"`
	// The description of the VCS label.
	Description string `json:"description,omitempty"`
	// The name of the VCS label.
	Name string `json:"name"`
}
