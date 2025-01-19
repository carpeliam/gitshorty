/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sc
import (
	"time"
)

// Stories are the standard unit of work in Shortcut and represent individual features, bugs, and chores.
type Story struct {
	// The Shortcut application url for the Story.
	AppUrl string `json:"app_url"`
	// The description of the story.
	Description string `json:"description"`
	// True if the story has been archived or not.
	Archived bool `json:"archived"`
	// A true/false boolean indicating if the Story has been started.
	Started bool `json:"started"`
	// An array of story links attached to the Story.
	StoryLinks []TypedStoryLink `json:"story_links"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// An array of labels attached to the story.
	Labels []LabelSlim `json:"labels"`
	// `Deprecated:` use `member_mention_ids`.
	MentionIds []string `json:"mention_ids"`
	SyncedItem *SyncedItem `json:"synced_item,omitempty"`
	// An array of Member IDs that have been mentioned in the Story description.
	MemberMentionIds []string `json:"member_mention_ids"`
	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// An array of CustomField value assertions for the story.
	CustomFields []StoryCustomField `json:"custom_fields,omitempty"`
	// An array of linked files attached to the story.
	LinkedFiles []LinkedFile `json:"linked_files"`
	// The ID of the workflow the story belongs to.
	WorkflowId int64 `json:"workflow_id"`
	// A manual override for the time/date the Story was completed.
	CompletedAtOverride time.Time `json:"completed_at_override"`
	// The time/date the Story was started.
	StartedAt time.Time `json:"started_at"`
	// The time/date the Story was completed.
	CompletedAt time.Time `json:"completed_at"`
	// The name of the story.
	Name string `json:"name"`
	GlobalId string `json:"global_id"`
	// A true/false boolean indicating if the Story has been completed.
	Completed bool `json:"completed"`
	// An array of comments attached to the story.
	Comments []StoryComment `json:"comments"`
	// A true/false boolean indicating if the Story is currently a blocker of another story.
	Blocker bool `json:"blocker"`
	// An array of Git branches attached to the story.
	Branches []Branch `json:"branches"`
	// The ID of the epic the story belongs to.
	EpicId int64 `json:"epic_id"`
	// The IDs of any unresolved blocker comments on the Story.
	UnresolvedBlockerComments []int64 `json:"unresolved_blocker_comments,omitempty"`
	// The ID of the story template used to create this story, or null if not created using a template.
	StoryTemplateId string `json:"story_template_id"`
	// An array of external links (strings) associated with a Story
	ExternalLinks []string `json:"external_links"`
	// The IDs of the iteration the story belongs to.
	PreviousIterationIds []int64 `json:"previous_iteration_ids"`
	// The ID of the Member that requested the story.
	RequestedById string `json:"requested_by_id"`
	// The ID of the iteration the story belongs to.
	IterationId int64 `json:"iteration_id"`
	SubTaskStoryIds []int64 `json:"sub_task_story_ids,omitempty"`
	// An array of tasks connected to the story.
	Tasks []Task `json:"tasks"`
	// An array of label ids attached to the story.
	LabelIds []int64 `json:"label_ids"`
	// A manual override for the time/date the Story was started.
	StartedAtOverride time.Time `json:"started_at_override"`
	// The ID of the group associated with the story.
	GroupId string `json:"group_id"`
	// The ID of the workflow state the story is currently in.
	WorkflowStateId int64 `json:"workflow_state_id"`
	// The time/date the Story was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// An array of Pull/Merge Requests attached to the story.
	PullRequests []PullRequest `json:"pull_requests"`
	// An array of Group IDs that have been mentioned in the Story description.
	GroupMentionIds []string `json:"group_mention_ids"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids"`
	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId string `json:"external_id"`
	// The unique ID of the Story.
	Id int64 `json:"id"`
	// The lead time (in seconds) of this story when complete.
	LeadTime int64 `json:"lead_time,omitempty"`
	ParentStoryId int64 `json:"parent_story_id,omitempty"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate int64 `json:"estimate"`
	// An array of commits attached to the story.
	Commits []Commit `json:"commits"`
	// An array of files attached to the story.
	Files []UploadedFile `json:"files"`
	// A number representing the position of the story in relation to every other story in the current project.
	Position int64 `json:"position"`
	// A true/false boolean indicating if the Story is currently blocked.
	Blocked bool `json:"blocked"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id"`
	// The due date of the story.
	Deadline time.Time `json:"deadline"`
	Stats *StoryStats `json:"stats"`
	// The cycle time (in seconds) of this story when complete.
	CycleTime int64 `json:"cycle_time,omitempty"`
	// The time/date the Story was created.
	CreatedAt time.Time `json:"created_at"`
	// The time/date the Story was last changed workflow-state.
	MovedAt time.Time `json:"moved_at"`
}
