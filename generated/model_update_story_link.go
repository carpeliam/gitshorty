/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

type UpdateStoryLink struct {
	// The type of link.
	Verb string `json:"verb,omitempty"`
	// The ID of the subject Story.
	SubjectId int64 `json:"subject_id,omitempty"`
	// The ID of the object Story.
	ObjectId int64 `json:"object_id,omitempty"`
}
