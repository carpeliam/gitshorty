# CreateObjective

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Objective. | [default to null]
**Description** | **string** | The Objective&#x27;s description. | [optional] [default to null]
**State** | **string** | The workflow state that the Objective is in. | [optional] [default to null]
**StartedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Objective was started. | [optional] [default to null]
**CompletedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Objective was completed. | [optional] [default to null]
**Categories** | [**[]CreateCategoryParams**](CreateCategoryParams.md) | An array of IDs of Categories attached to the Objective. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

