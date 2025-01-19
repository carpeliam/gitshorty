# UpdateMilestone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The Milestone&#x27;s description. | [optional] [default to null]
**Archived** | **bool** | A boolean indicating whether the Milestone is archived or not | [optional] [default to null]
**CompletedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Milestone was completed. | [optional] [default to null]
**Name** | **string** | The name of the Milestone. | [optional] [default to null]
**State** | **string** | The workflow state that the Milestone is in. | [optional] [default to null]
**StartedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Milestone was started. | [optional] [default to null]
**Categories** | [**[]CreateCategoryParams**](CreateCategoryParams.md) | An array of IDs of Categories attached to the Milestone. | [optional] [default to null]
**BeforeId** | **int64** | The ID of the Milestone we want to move this Milestone before. | [optional] [default to null]
**AfterId** | **int64** | The ID of the Milestone we want to move this Milestone after. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

