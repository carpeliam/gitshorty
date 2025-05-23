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

type UpdateFile struct {
	// The description of the file.
	Description string `json:"description,omitempty"`
	// The time/date that the file was uploaded.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The time/date that the file was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the file.
	Name string `json:"name,omitempty"`
	// The unique ID assigned to the Member who uploaded the file to Shortcut.
	UploaderId string `json:"uploader_id,omitempty"`
	// An additional ID that you may wish to assign to the file.
	ExternalId string `json:"external_id,omitempty"`
}
