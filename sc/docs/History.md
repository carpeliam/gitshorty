# History

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangedAt** | **string** | The date when the change occurred. | [default to null]
**PrimaryId** | [***interface{}**](interface{}.md) | The ID of the primary entity that has changed, if applicable. | [optional] [default to null]
**References** | [**[]interface{}**](interface{}.md) | An array of objects affected by the change. Reference objects provide basic information for the entities reference in the history actions. Some have specific fields, but they always contain an id, entity_type, and a name. | [optional] [default to null]
**Actions** | [**[]interface{}**](interface{}.md) | An array of actions that were performed for the change. | [default to null]
**MemberId** | **string** | The ID of the member who performed the change. | [optional] [default to null]
**ExternalId** | **string** | The ID of the webhook that handled the change. | [optional] [default to null]
**Id** | **string** | The ID representing the change for the story. | [default to null]
**Version** | **string** | The version of the change format. | [default to null]
**WebhookId** | **string** | The ID of the webhook that handled the change. | [optional] [default to null]
**AutomationId** | **string** | The ID of the automation that performed the change. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

