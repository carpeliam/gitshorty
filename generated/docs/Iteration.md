# Iteration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Iteration. | [default to null]
**Description** | **string** | The description of the iteration. | [default to null]
**EntityType** | **string** | A string description of this resource | [default to null]
**Labels** | [**[]Label**](Label.md) | An array of labels attached to the iteration. | [default to null]
**MentionIds** | **[]string** | &#x60;Deprecated:&#x60; use &#x60;member_mention_ids&#x60;. | [default to null]
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Story description. | [default to null]
**AssociatedGroups** | [**[]IterationAssociatedGroup**](IterationAssociatedGroup.md) | An array containing Group IDs and Group-owned story counts for the Iteration&#x27;s associated groups. | [default to null]
**Name** | **string** | The name of the iteration. | [default to null]
**GlobalId** | **string** |  | [default to null]
**LabelIds** | **[]int64** | An array of label ids attached to the iteration. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The instant when this iteration was last updated. | [default to null]
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Story description. | [default to null]
**EndDate** | [**time.Time**](time.Time.md) | The date this iteration ends. | [default to null]
**FollowerIds** | **[]string** | An array of UUIDs for any Members listed as Followers. | [default to null]
**GroupIds** | **[]string** | An array of UUIDs for any Groups you want to add as Followers. Currently, only one Group association is presented in our web UI. | [default to null]
**StartDate** | [**time.Time**](time.Time.md) | The date this iteration begins. | [default to null]
**Status** | **string** | The status of the iteration. Values are either \&quot;unstarted\&quot;, \&quot;started\&quot;, or \&quot;done\&quot;. | [default to null]
**Id** | **int64** | The ID of the iteration. | [default to null]
**Stats** | [***IterationStats**](IterationStats.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The instant when this iteration was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

