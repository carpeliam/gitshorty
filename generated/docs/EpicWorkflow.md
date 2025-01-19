# EpicWorkflow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | [default to null]
**Id** | **int64** | The unique ID of the Epic Workflow. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date the Epic Workflow was created. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date the Epic Workflow was updated. | [default to null]
**DefaultEpicStateId** | **int64** | The unique ID of the default Epic State that new Epics are assigned by default. | [default to null]
**EpicStates** | [**[]EpicState**](EpicState.md) | A map of the Epic States in this Epic Workflow. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

