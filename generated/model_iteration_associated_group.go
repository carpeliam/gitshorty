/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

type IterationAssociatedGroup struct {
	// The Group ID of the associated group.
	GroupId string `json:"group_id"`
	// The number of stories this Group owns in the Iteration.
	AssociatedStoriesCount int64 `json:"associated_stories_count,omitempty"`
}
