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

// EpicSlim represents the same resource as an Epic but is more light-weight, including all Epic fields except the comments array. The description string can be optionally included. Use the [Get Epic](#Get-Epic) endpoint to fetch the unabridged payload for an Epic.
type EpicSlim struct {
	// The Shortcut application url for the Epic.
	AppUrl string `json:"app_url"`
	// The Epic's description.
	Description string `json:"description,omitempty"`
	// True/false boolean that indicates whether the Epic is archived or not.
	Archived bool `json:"archived"`
	// A true/false boolean indicating if the Epic has been started.
	Started bool `json:"started"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// An array of Labels attached to the Epic.
	Labels []LabelSlim `json:"labels"`
	// `Deprecated:` use `member_mention_ids`.
	MentionIds []string `json:"mention_ids"`
	// An array of Member IDs that have been mentioned in the Epic description.
	MemberMentionIds []string `json:"member_mention_ids"`
	// An array containing Group IDs and Group-owned story counts for the Epic's associated groups.
	AssociatedGroups []EpicAssociatedGroup `json:"associated_groups"`
	// The IDs of Projects related to this Epic.
	ProjectIds []int64 `json:"project_ids"`
	// The number of stories in this epic which are not associated with a project.
	StoriesWithoutProjects int64 `json:"stories_without_projects"`
	// A manual override for the time/date the Epic was completed.
	CompletedAtOverride time.Time `json:"completed_at_override"`
	// The ID of the associated productboard integration.
	ProductboardPluginId string `json:"productboard_plugin_id"`
	// The time/date the Epic was started.
	StartedAt time.Time `json:"started_at"`
	// The time/date the Epic was completed.
	CompletedAt time.Time `json:"completed_at"`
	// An array of IDs for Objectives to which this epic is related.
	ObjectiveIds []int64 `json:"objective_ids"`
	// The name of the Epic.
	Name string `json:"name"`
	GlobalId string `json:"global_id"`
	// A true/false boolean indicating if the Epic has been completed.
	Completed bool `json:"completed"`
	// The URL of the associated productboard feature.
	ProductboardUrl string `json:"productboard_url"`
	// The Epic's planned start date.
	PlannedStartDate time.Time `json:"planned_start_date"`
	// `Deprecated` The workflow state that the Epic is in.
	State string `json:"state"`
	// `Deprecated` The ID of the Objective this Epic is related to. Use `objective_ids`.
	MilestoneId int64 `json:"milestone_id"`
	// The ID of the Member that requested the epic.
	RequestedById string `json:"requested_by_id"`
	// The ID of the Epic State.
	EpicStateId int64 `json:"epic_state_id"`
	// An array of Label ids attached to the Epic.
	LabelIds []int64 `json:"label_ids"`
	// A manual override for the time/date the Epic was started.
	StartedAtOverride time.Time `json:"started_at_override"`
	// `Deprecated` The ID of the group to associate with the epic. Use `group_ids`.
	GroupId string `json:"group_id"`
	// The time/date the Epic was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// An array of Group IDs that have been mentioned in the Epic description.
	GroupMentionIds []string `json:"group_mention_ids"`
	// The ID of the associated productboard feature.
	ProductboardId string `json:"productboard_id"`
	// An array of UUIDs for any Members you want to add as Followers on this Epic.
	FollowerIds []string `json:"follower_ids"`
	// An array of UUIDS for Groups to which this Epic is related.
	GroupIds []string `json:"group_ids"`
	// An array of UUIDs for any members you want to add as Owners on this new Epic.
	OwnerIds []string `json:"owner_ids"`
	// This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The unique ID of the Epic.
	Id int64 `json:"id"`
	// The Epic's relative position in the Epic workflow state.
	Position int64 `json:"position"`
	// The name of the associated productboard feature.
	ProductboardName string `json:"productboard_name"`
	// The Epic's deadline.
	Deadline time.Time `json:"deadline"`
	Stats *EpicStats `json:"stats"`
	// The time/date the Epic was created.
	CreatedAt time.Time `json:"created_at"`
}
