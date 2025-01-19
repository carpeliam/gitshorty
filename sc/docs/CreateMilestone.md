# CreateMilestone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Milestone. | [default to null]
**Description** | **string** | The Milestone&#x27;s description. | [optional] [default to null]
**State** | **string** | The workflow state that the Milestone is in. | [optional] [default to null]
**StartedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Milestone was started. | [optional] [default to null]
**CompletedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Milestone was completed. | [optional] [default to null]
**Categories** | [**[]CreateCategoryParams**](CreateCategoryParams.md) | An array of IDs of Categories attached to the Milestone. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

