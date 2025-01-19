# WorkflowState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of what sort of Stories belong in that Workflow state. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**Color** | **string** | The hex color for this Workflow State. | [optional] [default to null]
**Verb** | **string** | The verb that triggers a move to that Workflow State when making VCS commits. | [default to null]
**Name** | **string** | The Workflow State&#x27;s name. | [default to null]
**GlobalId** | **string** |  | [default to null]
**NumStories** | **int64** | The number of Stories currently in that Workflow State. | [default to null]
**Type_** | **string** | The type of Workflow State (Unstarted, Started, or Finished) | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When the Workflow State was last updated. | [default to null]
**Id** | **int64** | The unique ID of the Workflow State. | [default to null]
**NumStoryTemplates** | **int64** | The number of Story Templates associated with that Workflow State. | [default to null]
**Position** | **int64** | The position that the Workflow State is in, starting with 0 at the left. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Workflow State was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

