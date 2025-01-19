# CreateEpic

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The Epic&#x27;s description. | [optional] [default to null]
**Labels** | [**[]CreateLabelParams**](CreateLabelParams.md) | An array of Labels attached to the Epic. | [optional] [default to null]
**CompletedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Epic was completed. | [optional] [default to null]
**ObjectiveIds** | **[]int64** | An array of IDs for Objectives to which this Epic is related. | [optional] [default to null]
**Name** | **string** | The Epic&#x27;s name. | [default to null]
**PlannedStartDate** | [**time.Time**](time.Time.md) | The Epic&#x27;s planned start date. | [optional] [default to null]
**State** | **string** | &#x60;Deprecated&#x60; The Epic&#x27;s state (to do, in progress, or done); will be ignored when &#x60;epic_state_id&#x60; is set. | [optional] [default to null]
**MilestoneId** | **int64** | &#x60;Deprecated&#x60; The ID of the Milestone this Epic is related to. Use &#x60;objective_ids&#x60;. | [optional] [default to null]
**RequestedById** | **string** | The ID of the member that requested the epic. | [optional] [default to null]
**EpicStateId** | **int64** | The ID of the Epic State. | [optional] [default to null]
**StartedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Epic was started. | [optional] [default to null]
**GroupId** | **string** | &#x60;Deprecated&#x60; The ID of the group to associate with the epic. Use &#x60;group_ids&#x60;. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Defaults to the time/date it is created but can be set to reflect another date. | [optional] [default to null]
**FollowerIds** | **[]string** | An array of UUIDs for any Members you want to add as Followers on this new Epic. | [optional] [default to null]
**GroupIds** | **[]string** | An array of UUIDS for Groups to which this Epic is related. | [optional] [default to null]
**OwnerIds** | **[]string** | An array of UUIDs for any members you want to add as Owners on this new Epic. | [optional] [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here. | [optional] [default to null]
**Deadline** | [**time.Time**](time.Time.md) | The Epic&#x27;s deadline. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Defaults to the time/date it is created but can be set to reflect another date. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

