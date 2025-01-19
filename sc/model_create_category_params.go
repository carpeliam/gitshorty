/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc

// Request parameters for creating a Category with a Objective.
type CreateCategoryParams struct {
	// The name of the new Category.
	Name string `json:"name"`
	// The hex color to be displayed with the Category (for example, \"#ff0000\").
	Color string `json:"color,omitempty"`
	// This field can be set to another unique ID. In the case that the Category has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id,omitempty"`
}
