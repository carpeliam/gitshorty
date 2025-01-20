/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

// Request parameters for creating an entirely new entity template.
type CreateEntityTemplate struct {
	// The name of the new entity template
	Name string `json:"name"`
	// The id of the user creating this template.
	AuthorId string `json:"author_id,omitempty"`
	StoryContents *CreateStoryContents `json:"story_contents"`
}
