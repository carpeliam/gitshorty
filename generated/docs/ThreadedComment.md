# ThreadedComment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Comment. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**Deleted** | **bool** | True/false boolean indicating whether the Comment is deleted. | [default to null]
**MentionIds** | **[]string** | &#x60;Deprecated:&#x60; use &#x60;member_mention_ids&#x60;. | [default to null]
**AuthorId** | **string** | The unique ID of the Member that authored the Comment. | [default to null]
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in this Comment. | [default to null]
**Comments** | [**[]ThreadedComment**](ThreadedComment.md) | A nested array of threaded comments. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Comment was updated. | [default to null]
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in this Comment. | [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here. | [default to null]
**Id** | **int64** | The unique ID of the Comment. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Comment was created. | [default to null]
**Text** | **string** | The text of the Comment. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

