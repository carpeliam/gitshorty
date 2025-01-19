# UpdateStories

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | If the Stories should be archived or not. | [optional] [default to null]
**StoryIds** | **[]int64** | The Ids of the Stories you wish to update. | [default to null]
**StoryType** | **string** | The type of story (feature, bug, chore). | [optional] [default to null]
**MoveTo** | **string** | One of \&quot;first\&quot; or \&quot;last\&quot;. This can be used to move the given story to the first or last position in the workflow state. | [optional] [default to null]
**FollowerIdsAdd** | **[]string** | The UUIDs of the new followers to be added. | [optional] [default to null]
**EpicId** | **int64** | The ID of the epic the story belongs to. | [optional] [default to null]
**ExternalLinks** | **[]string** | An array of External Links associated with this story. | [optional] [default to null]
**FollowerIdsRemove** | **[]string** | The UUIDs of the followers to be removed. | [optional] [default to null]
**RequestedById** | **string** | The ID of the member that requested the story. | [optional] [default to null]
**IterationId** | **int64** | The ID of the iteration the story belongs to. | [optional] [default to null]
**CustomFieldsRemove** | [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] [default to null]
**LabelsAdd** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be added. | [optional] [default to null]
**GroupId** | **string** | The Id of the Group the Stories should belong to. | [optional] [default to null]
**WorkflowStateId** | **int64** | The ID of the workflow state to put the stories in. | [optional] [default to null]
**BeforeId** | **int64** | The ID of the story that the stories are to be moved before. | [optional] [default to null]
**Estimate** | **int64** | The numeric point estimate of the story. Can also be null, which means unestimated. | [optional] [default to null]
**AfterId** | **int64** | The ID of the story that the stories are to be moved below. | [optional] [default to null]
**OwnerIdsRemove** | **[]string** | The UUIDs of the owners to be removed. | [optional] [default to null]
**CustomFieldsAdd** | [**[]CustomFieldValueParams**](CustomFieldValueParams.md) | A map specifying a CustomField ID and CustomFieldEnumValue ID that represents an assertion of some value for a CustomField. | [optional] [default to null]
**ProjectId** | **int64** | The ID of the Project the Stories should belong to. | [optional] [default to null]
**LabelsRemove** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of labels to be removed. | [optional] [default to null]
**Deadline** | [**time.Time**](time.Time.md) | The due date of the story. | [optional] [default to null]
**OwnerIdsAdd** | **[]string** | The UUIDs of the new owners to be added. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

