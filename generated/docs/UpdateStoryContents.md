# UpdateStoryContents

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the story. | [optional] [default to null]
**Labels** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be populated by the template. | [optional] [default to null]
**StoryType** | **string** | The type of story (feature, bug, chore). | [optional] [default to null]
**CustomFields** | [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | An array of maps specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] [default to null]
**FileIds** | **[]int64** | An array of the attached file IDs to be populated. | [optional] [default to null]
**Name** | **string** | The name of the story. | [optional] [default to null]
**EpicId** | **int64** | The ID of the epic the to be populated. | [optional] [default to null]
**ExternalLinks** | **[]string** | An array of external links to be populated. | [optional] [default to null]
**IterationId** | **int64** | The ID of the iteration the to be populated. | [optional] [default to null]
**Tasks** | [**[]BaseTaskParams**](BaseTaskParams.md) | An array of tasks to be populated by the template. | [optional] [default to null]
**GroupId** | **string** | The ID of the group to be populated. | [optional] [default to null]
**WorkflowStateId** | **int64** | The ID of the workflow state to be populated. | [optional] [default to null]
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | [optional] [default to null]
**OwnerIds** | **[]string** | An array of UUIDs of the owners of this story. | [optional] [default to null]
**Estimate** | **int64** | The numeric point estimate to be populated. | [optional] [default to null]
**ProjectId** | **int64** | The ID of the project the story belongs to. | [optional] [default to null]
**LinkedFileIds** | **[]int64** | An array of the linked file IDs to be populated. | [optional] [default to null]
**Deadline** | [**time.Time**](time.Time.md) | The due date of the story. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

