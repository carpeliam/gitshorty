/*
 * Shortcut API
 *
 * Shortcut API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package generated

// The changes that have occurred as a result of the action.
type HistoryChangesStory struct {
	Description *StoryHistoryChangeOldNewStr `json:"description,omitempty"`
	Archived *StoryHistoryChangeOldNewBool `json:"archived,omitempty"`
	Started *StoryHistoryChangeOldNewBool `json:"started,omitempty"`
	TaskIds *StoryHistoryChangeAddsRemovesInt `json:"task_ids,omitempty"`
	MentionIds *StoryHistoryChangeAddsRemovesUuid `json:"mention_ids,omitempty"`
	StoryType *StoryHistoryChangeOldNewStr `json:"story_type,omitempty"`
	Name *StoryHistoryChangeOldNewStr `json:"name,omitempty"`
	Completed *StoryHistoryChangeOldNewBool `json:"completed,omitempty"`
	Blocker *StoryHistoryChangeOldNewBool `json:"blocker,omitempty"`
	EpicId *StoryHistoryChangeOldNewInt `json:"epic_id,omitempty"`
	BranchIds *StoryHistoryChangeAddsRemovesInt `json:"branch_ids,omitempty"`
	CommitIds *StoryHistoryChangeAddsRemovesInt `json:"commit_ids,omitempty"`
	RequestedById *StoryHistoryChangeOldNewUuid `json:"requested_by_id,omitempty"`
	IterationId *StoryHistoryChangeOldNewInt `json:"iteration_id,omitempty"`
	LabelIds *StoryHistoryChangeAddsRemovesInt `json:"label_ids,omitempty"`
	GroupId *StoryHistoryChangeOldNewUuid `json:"group_id,omitempty"`
	WorkflowStateId *StoryHistoryChangeOldNewInt `json:"workflow_state_id,omitempty"`
	ObjectStoryLinkIds *StoryHistoryChangeAddsRemovesInt `json:"object_story_link_ids,omitempty"`
	FollowerIds *StoryHistoryChangeAddsRemovesUuid `json:"follower_ids,omitempty"`
	OwnerIds *StoryHistoryChangeAddsRemovesUuid `json:"owner_ids,omitempty"`
	CustomFieldValueIds *StoryHistoryChangeAddsRemovesUuid `json:"custom_field_value_ids,omitempty"`
	Estimate *StoryHistoryChangeOldNewInt `json:"estimate,omitempty"`
	SubjectStoryLinkIds *StoryHistoryChangeAddsRemovesInt `json:"subject_story_link_ids,omitempty"`
	Blocked *StoryHistoryChangeOldNewBool `json:"blocked,omitempty"`
	ProjectId *StoryHistoryChangeOldNewInt `json:"project_id,omitempty"`
	Deadline *StoryHistoryChangeOldNewStr `json:"deadline,omitempty"`
}
