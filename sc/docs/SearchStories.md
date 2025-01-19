# SearchStories

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | A true/false boolean indicating whether the Story is in archived state. | [optional] [default to null]
**OwnerId** | **string** | An array of UUIDs for any Users who may be Owners of the Stories. | [optional] [default to null]
**StoryType** | **string** | The type of Stories that you want returned. | [optional] [default to null]
**EpicIds** | **[]int64** | The Epic IDs that may be associated with the Stories. | [optional] [default to null]
**ProjectIds** | **[]int64** | The IDs for the Projects the Stories may be assigned to. | [optional] [default to null]
**UpdatedAtEnd** | [**time.Time**](time.Time.md) | Stories should have been updated on or before this date. | [optional] [default to null]
**CompletedAtEnd** | [**time.Time**](time.Time.md) | Stories should have been completed on or before this date. | [optional] [default to null]
**WorkflowStateTypes** | **[]string** | The type of Workflow State the Stories may be in. | [optional] [default to null]
**DeadlineEnd** | [**time.Time**](time.Time.md) | Stories should have a deadline on or before this date. | [optional] [default to null]
**CreatedAtStart** | [**time.Time**](time.Time.md) | Stories should have been created on or after this date. | [optional] [default to null]
**EpicId** | **int64** | The Epic IDs that may be associated with the Stories. | [optional] [default to null]
**LabelName** | **string** | The name of any associated Labels. | [optional] [default to null]
**RequestedById** | **string** | The UUID of any Users who may have requested the Stories. | [optional] [default to null]
**IterationId** | **int64** | The Iteration ID that may be associated with the Stories. | [optional] [default to null]
**LabelIds** | **[]int64** | The Label IDs that may be associated with the Stories. | [optional] [default to null]
**GroupId** | **string** | The Group ID that is associated with the Stories | [optional] [default to null]
**WorkflowStateId** | **int64** | The unique IDs of the specific Workflow States that the Stories should be in. | [optional] [default to null]
**IterationIds** | **[]int64** | The Iteration IDs that may be associated with the Stories. | [optional] [default to null]
**CreatedAtEnd** | [**time.Time**](time.Time.md) | Stories should have been created on or before this date. | [optional] [default to null]
**DeadlineStart** | [**time.Time**](time.Time.md) | Stories should have a deadline on or after this date. | [optional] [default to null]
**GroupIds** | **[]string** | The Group IDs that are associated with the Stories | [optional] [default to null]
**OwnerIds** | **[]string** | An array of UUIDs for any Users who may be Owners of the Stories. | [optional] [default to null]
**ExternalId** | **string** | An ID or URL that references an external resource. Useful during imports. | [optional] [default to null]
**IncludesDescription** | **bool** | Whether to include the story description in the response. | [optional] [default to null]
**Estimate** | **int64** | The number of estimate points associate with the Stories. | [optional] [default to null]
**ProjectId** | **int64** | The IDs for the Projects the Stories may be assigned to. | [optional] [default to null]
**CompletedAtStart** | [**time.Time**](time.Time.md) | Stories should have been completed on or after this date. | [optional] [default to null]
**UpdatedAtStart** | [**time.Time**](time.Time.md) | Stories should have been updated on or after this date. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

