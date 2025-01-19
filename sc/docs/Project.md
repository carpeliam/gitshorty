# Project

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Project. | [default to null]
**Description** | **string** | The description of the Project. | [default to null]
**Archived** | **bool** | True/false boolean indicating whether the Project is in an Archived state. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**DaysToThermometer** | **int64** | The number of days before the thermometer appears in the Story summary. | [default to null]
**Color** | **string** | The color associated with the Project in the Shortcut member interface. | [default to null]
**WorkflowId** | **int64** | The ID of the workflow the project belongs to. | [default to null]
**Name** | **string** | The name of the Project | [default to null]
**GlobalId** | **string** | The Global ID of the Project. | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The date at which the Project was started. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date that the Project was last updated. | [default to null]
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here. | [default to null]
**Id** | **int64** | The unique ID of the Project. | [default to null]
**ShowThermometer** | **bool** | Configuration to enable or disable thermometers in the Story summary. | [default to null]
**TeamId** | **int64** | The ID of the team the project belongs to. | [default to null]
**IterationLength** | **int64** | The number of weeks per iteration in this Project. | [default to null]
**Abbreviation** | **string** | The Project abbreviation used in Story summaries. Should be kept to 3 characters at most. | [default to null]
**Stats** | [***ProjectStats**](ProjectStats.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date that the Project was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

