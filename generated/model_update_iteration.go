/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

type UpdateIteration struct {
	// An array of UUIDs for any Members you want to add as Followers.
	FollowerIds []string `json:"follower_ids,omitempty"`
	// An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI.
	GroupIds []string `json:"group_ids,omitempty"`
	// An array of Labels attached to the Iteration.
	Labels []CreateLabelParams `json:"labels,omitempty"`
	// The description of the Iteration.
	Description string `json:"description,omitempty"`
	// The name of this Iteration
	Name string `json:"name,omitempty"`
	// The date this Iteration begins, e.g. 2019-07-01
	StartDate string `json:"start_date,omitempty"`
	// The date this Iteration ends, e.g. 2019-07-05.
	EndDate string `json:"end_date,omitempty"`
}
