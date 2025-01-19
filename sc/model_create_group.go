/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc

type CreateGroup struct {
	// The description of the Group.
	Description string `json:"description,omitempty"`
	// The Member ids to add to this Group.
	MemberIds []string `json:"member_ids,omitempty"`
	// The Workflow ids to add to the Group.
	WorkflowIds []int64 `json:"workflow_ids,omitempty"`
	// The name of this Group.
	Name string `json:"name"`
	// The mention name of this Group.
	MentionName string `json:"mention_name"`
	// The color you wish to use for the Group in the system.
	Color string `json:"color,omitempty"`
	// The color key you wish to use for the Group in the system.
	ColorKey string `json:"color_key,omitempty"`
	// The Icon id for the avatar of this Group.
	DisplayIconId string `json:"display_icon_id,omitempty"`
}
