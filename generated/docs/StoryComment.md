# StoryComment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Comment. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**Deleted** | **bool** | True/false boolean indicating whether the Comment has been deleted. | [default to null]
**StoryId** | **int64** | The ID of the Story on which the Comment appears. | [default to null]
**MentionIds** | **[]string** | &#x60;Deprecated:&#x60; use &#x60;member_mention_ids&#x60;. | [default to null]
**AuthorId** | **string** | The unique ID of the Member who is the Comment&#x27;s author. | [default to null]
**MemberMentionIds** | **[]string** | The unique IDs of the Member who are mentioned in the Comment. | [default to null]
**Blocker** | **bool** | Marks the comment as a blocker that can be surfaced to permissions or teams mentioned in the comment. Can only be used on a top-level comment. | [optional] [default to null]
**LinkedToSlack** | **bool** | Whether the Comment is currently the root of a thread that is linked to Slack. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date when the Comment was updated. | [default to null]
**GroupMentionIds** | **[]string** | The unique IDs of the Group who are mentioned in the Comment. | [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Comment has been imported from another tool, the ID in the other tool can be indicated here. | [default to null]
**ParentId** | **int64** | The ID of the parent Comment this Comment is threaded under. | [optional] [default to null]
**Id** | **int64** | The unique ID of the Comment. | [default to null]
**Position** | **int64** | The Comments numerical position in the list from oldest to newest. | [default to null]
**UnblocksParent** | **bool** | Marks the comment as an unblocker to its  blocker parent. Can only be set on a threaded comment who has a parent with &#x60;blocker&#x60; set. | [optional] [default to null]
**Reactions** | [**[]StoryReaction**](StoryReaction.md) | A set of Reactions to this Comment. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date when the Comment was created. | [default to null]
**Text** | **string** | The text of the Comment. In the case that the Comment has been deleted, this field can be set to nil. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

