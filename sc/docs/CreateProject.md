# CreateProject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The Project description. | [optional] [default to null]
**Color** | **string** | The color you wish to use for the Project in the system. | [optional] [default to null]
**Name** | **string** | The name of the Project. | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The date at which the Project was started. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Defaults to the time/date it is created but can be set to reflect another date. | [optional] [default to null]
**FollowerIds** | **[]string** | An array of UUIDs for any members you want to add as Owners on this new Epic. | [optional] [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Project has been imported from another tool, the ID in the other tool can be indicated here. | [optional] [default to null]
**TeamId** | **int64** | The ID of the team the project belongs to. | [default to null]
**IterationLength** | **int64** | The number of weeks per iteration in this Project. | [optional] [default to null]
**Abbreviation** | **string** | The Project abbreviation used in Story summaries. Should be kept to 3 characters at most. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Defaults to the time/date it is created but can be set to reflect another date. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

