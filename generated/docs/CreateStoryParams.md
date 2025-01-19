# CreateStoryParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the story. | [optional] [default to null]
**Archived** | **bool** | Controls the story&#x27;s archived state. | [optional] [default to null]
**StoryLinks** | [**[]CreateStoryLinkParams**](CreateStoryLinkParams.md) | An array of story links attached to the story. | [optional] [default to null]
**Labels** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels attached to the story. | [optional] [default to null]
**StoryType** | **string** | The type of story (feature, bug, chore). | [optional] [default to null]
**CustomFields** | [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] [default to null]
**MoveTo** | **string** | One of \&quot;first\&quot; or \&quot;last\&quot;. This can be used to move the given story to the first or last position in the workflow state. | [optional] [default to null]
**FileIds** | **[]int64** | An array of IDs of files attached to the story. | [optional] [default to null]
**SourceTaskId** | **int64** | Given this story was converted from a task in another story, this is the original task ID that was converted to this story. | [optional] [default to null]
**CompletedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Story was completed. | [optional] [default to null]
**Name** | **string** | The name of the story. | [default to null]
**Comments** | [**[]CreateStoryCommentParams**](CreateStoryCommentParams.md) | An array of comments to add to the story. | [optional] [default to null]
**EpicId** | **int64** | The ID of the epic the story belongs to. | [optional] [default to null]
**StoryTemplateId** | **string** | The id of the story template used to create this story, if applicable. This is just an association; no content from the story template is inherited by the story simply by setting this field. | [optional] [default to null]
**ExternalLinks** | **[]string** | An array of External Links associated with this story. | [optional] [default to null]
**SubTasks** | [**[]CreateSubTaskParams**](CreateSubTaskParams.md) | An array of sub tasks to create. | [optional] [default to null]
**RequestedById** | **string** | The ID of the member that requested the story. | [optional] [default to null]
**IterationId** | **int64** | The ID of the iteration the story belongs to. | [optional] [default to null]
**Tasks** | [**[]CreateTaskParams**](CreateTaskParams.md) | An array of tasks connected to the story. | [optional] [default to null]
**StartedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Story was started. | [optional] [default to null]
**GroupId** | **string** | The id of the group to associate with this story. | [optional] [default to null]
**WorkflowStateId** | **int64** | The ID of the workflow state the story will be in. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Story was updated. | [optional] [default to null]
**FollowerIds** | **[]string** | An array of UUIDs of the followers of this story. | [optional] [default to null]
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | [optional] [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here. | [optional] [default to null]
**Estimate** | **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] [default to null]
**ProjectId** | **int64** | The ID of the project the story belongs to. | [optional] [default to null]
**LinkedFileIds** | **[]int64** | An array of IDs of linked files attached to the story. | [optional] [default to null]
**Deadline** | [**time.Time**](time.Time.md) | The due date of the story. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Story was created. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

