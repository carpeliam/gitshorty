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

// Repository refers to a VCS repository.
type Repository struct {
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The shorthand name of the VCS repository.
	Name string `json:"name"`
	// The VCS provider for the Repository.
	Type_ string `json:"type"`
	// The time/date the Repository was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The VCS unique identifier for the Repository.
	ExternalId string `json:"external_id"`
	// The ID associated to the VCS repository in Shortcut.
	Id int64 `json:"id"`
	// The URL of the Repository.
	Url string `json:"url"`
	// The full name of the VCS repository.
	FullName string `json:"full_name"`
	// The time/date the Repository was created.
	CreatedAt time.Time `json:"created_at"`
}
