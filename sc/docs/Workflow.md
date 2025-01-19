# Workflow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description of the workflow. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**ProjectIds** | **[]float64** | An array of IDs of projects within the Workflow. | [default to null]
**States** | [**[]WorkflowState**](WorkflowState.md) | A map of the states in this Workflow. | [default to null]
**Name** | **string** | The name of the workflow. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date the Workflow was updated. | [default to null]
**AutoAssignOwner** | **bool** | Indicates if an owner is automatically assigned when an unowned story is started. | [default to null]
**Id** | **int64** | The unique ID of the Workflow. | [default to null]
**TeamId** | **int64** | The ID of the team the workflow belongs to. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date the Workflow was created. | [default to null]
**DefaultStateId** | **int64** | The unique ID of the default state that new Stories are entered into. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

