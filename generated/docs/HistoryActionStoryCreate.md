# HistoryActionStoryCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The application URL of the Story. | [default to null]
**Description** | **string** | The description of the Story. | [optional] [default to null]
**Started** | **bool** | Whether or not the Story has been started. | [optional] [default to null]
**EntityType** | **string** | The type of entity referenced. | [default to null]
**TaskIds** | **[]int64** | An array of Task IDs on this Story. | [optional] [default to null]
**StoryType** | **string** | The type of Story; either feature, bug, or chore. | [default to null]
**Name** | **string** | The name of the Story. | [default to null]
**Completed** | **bool** | Whether or not the Story is completed. | [optional] [default to null]
**Blocker** | **bool** | Whether or not the Story is blocking another Story. | [optional] [default to null]
**EpicId** | **int64** | The Epic ID for this Story. | [optional] [default to null]
**RequestedById** | **string** | The ID of the Member that requested the Story. | [optional] [default to null]
**IterationId** | **int64** | The Iteration ID the Story is in. | [optional] [default to null]
**LabelIds** | **[]int64** | An array of Labels IDs attached to the Story. | [optional] [default to null]
**GroupId** | **string** | The Team IDs for the followers of the Story. | [optional] [default to null]
**WorkflowStateId** | **int64** | An array of Workflow State IDs attached to the Story. | [optional] [default to null]
**ObjectStoryLinkIds** | **[]int64** | An array of Story IDs that are the object of a Story Link relationship. | [optional] [default to null]
**FollowerIds** | **[]string** | An array of Member IDs for the followers of the Story. | [optional] [default to null]
**OwnerIds** | **[]string** | An array of Member IDs that are the owners of the Story. | [optional] [default to null]
**CustomFieldValueIds** | **[]string** | An array of Custom Field Enum Value ids on this Story. | [optional] [default to null]
**Id** | **int64** | The ID of the entity referenced. | [default to null]
**Estimate** | **int64** | The estimate (or point value) for the Story. | [optional] [default to null]
**SubjectStoryLinkIds** | **[]int64** | An array of Story IDs that are the subject of a Story Link relationship. | [optional] [default to null]
**Action** | **string** | The action of the entity referenced. | [default to null]
**Blocked** | **bool** | Whether or not the Story is blocked by another Story. | [optional] [default to null]
**ProjectId** | **int64** | The Project ID of the Story is in. | [optional] [default to null]
**Deadline** | **string** | The timestamp representing the Story&#x27;s deadline. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

