# Task

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Full text of the Task. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**StoryId** | **int64** | The unique identifier of the parent Story. | [default to null]
**MentionIds** | **[]string** | &#x60;Deprecated:&#x60; use &#x60;member_mention_ids&#x60;. | [default to null]
**MemberMentionIds** | **[]string** | An array of UUIDs of Members mentioned in this Task. | [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) | The time/date the Task was completed. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Task was updated. | [default to null]
**GroupMentionIds** | **[]string** | An array of UUIDs of Groups mentioned in this Task. | [default to null]
**OwnerIds** | **[]string** | An array of UUIDs of the Owners of this Task. | [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Task has been imported from another tool, the ID in the other tool can be indicated here. | [default to null]
**Id** | **int64** | The unique ID of the Task. | [default to null]
**Position** | **int64** | The number corresponding to the Task&#x27;s position within a list of Tasks on a Story. | [default to null]
**Complete** | **bool** | True/false boolean indicating whether the Task has been completed. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Task was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

