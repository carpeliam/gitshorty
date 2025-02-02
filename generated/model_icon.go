/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated
import (
	"time"
)

// Icons are used to attach images to Groups, Workspaces, Members, and Loading screens in the Shortcut web application.
type Icon struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The unique ID of the Icon.
	Id string `json:"id"`
	// The time/date that the Icon was created.
	CreatedAt time.Time `json:"created_at"`
	// The time/date that the Icon was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The URL of the Icon.
	Url string `json:"url"`
}
