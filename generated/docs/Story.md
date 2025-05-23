# Story

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Story. | [default to null]
**Description** | **string** | The description of the story. | [default to null]
**Archived** | **bool** | True if the story has been archived or not. | [default to null]
**Started** | **bool** | A true/false boolean indicating if the Story has been started. | [default to null]
**StoryLinks** | [**[]TypedStoryLink**](TypedStoryLink.md) | An array of story links attached to the Story. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**Labels** | [**[]LabelSlim**](LabelSlim.md) | An array of labels attached to the story. | [default to null]
**MentionIds** | **[]string** | &#x60;Deprecated:&#x60; use &#x60;member_mention_ids&#x60;. | [default to null]
**SyncedItem** | [***SyncedItem**](SyncedItem.md) |  | [optional] [default to null]
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Story description. | [default to null]
**StoryType** | **string** | The type of story (feature, bug, chore). | [default to null]
**CustomFields** | [**[]StoryCustomField**](StoryCustomField.md) | An array of CustomField value assertions for the story. | [optional] [default to null]
**LinkedFiles** | [**[]LinkedFile**](LinkedFile.md) | An array of linked files attached to the story. | [default to null]
**WorkflowId** | **int64** | The ID of the workflow the story belongs to. | [default to null]
**CompletedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Story was completed. | [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | The time/date the Story was started. | [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) | The time/date the Story was completed. | [default to null]
**Name** | **string** | The name of the story. | [default to null]
**GlobalId** | **string** |  | [default to null]
**Completed** | **bool** | A true/false boolean indicating if the Story has been completed. | [default to null]
**Comments** | [**[]StoryComment**](StoryComment.md) | An array of comments attached to the story. | [default to null]
**Blocker** | **bool** | A true/false boolean indicating if the Story is currently a blocker of another story. | [default to null]
**Branches** | [**[]Branch**](Branch.md) | An array of Git branches attached to the story. | [default to null]
**EpicId** | **int64** | The ID of the epic the story belongs to. | [default to null]
**UnresolvedBlockerComments** | **[]int64** | The IDs of any unresolved blocker comments on the Story. | [optional] [default to null]
**StoryTemplateId** | **string** | The ID of the story template used to create this story, or null if not created using a template. | [default to null]
**ExternalLinks** | **[]string** | An array of external links (strings) associated with a Story | [default to null]
**PreviousIterationIds** | **[]int64** | The IDs of the iteration the story belongs to. | [default to null]
**RequestedById** | **string** | The ID of the Member that requested the story. | [default to null]
**IterationId** | **int64** | The ID of the iteration the story belongs to. | [default to null]
**SubTaskStoryIds** | **[]int64** |  | [optional] [default to null]
**Tasks** | [**[]Task**](Task.md) | An array of tasks connected to the story. | [default to null]
**LabelIds** | **[]int64** | An array of label ids attached to the story. | [default to null]
**StartedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Story was started. | [default to null]
**GroupId** | **string** | The ID of the group associated with the story. | [default to null]
**WorkflowStateId** | **int64** | The ID of the workflow state the story is currently in. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Story was updated. | [default to null]
**PullRequests** | [**[]PullRequest**](PullRequest.md) | An array of Pull/Merge Requests attached to the story. | [default to null]
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Story description. | [default to null]
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | [default to null]
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here. | [default to null]
**Id** | **int64** | The unique ID of the Story. | [default to null]
**LeadTime** | **int64** | The lead time (in seconds) of this story when complete. | [optional] [default to null]
**ParentStoryId** | **int64** |  | [optional] [default to null]
**Estimate** | **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [default to null]
**Commits** | [**[]Commit**](Commit.md) | An array of commits attached to the story. | [default to null]
**Files** | [**[]UploadedFile**](UploadedFile.md) | An array of files attached to the story. | [default to null]
**Position** | **int64** | A number representing the position of the story in relation to every other story in the current project. | [default to null]
**Blocked** | **bool** | A true/false boolean indicating if the Story is currently blocked. | [default to null]
**ProjectId** | **int64** | The ID of the project the story belongs to. | [default to null]
**Deadline** | [**time.Time**](time.Time.md) | The due date of the story. | [default to null]
**Stats** | [***StoryStats**](StoryStats.md) |  | [default to null]
**CycleTime** | **int64** | The cycle time (in seconds) of this story when complete. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Story was created. | [default to null]
**MovedAt** | [**time.Time**](time.Time.md) | The time/date the Story was last changed workflow-state. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

