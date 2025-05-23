# EpicSlim

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUrl** | **string** | The Shortcut application url for the Epic. | [default to null]
**Description** | **string** | The Epic&#x27;s description. | [optional] [default to null]
**Archived** | **bool** | True/false boolean that indicates whether the Epic is archived or not. | [default to null]
**Started** | **bool** | A true/false boolean indicating if the Epic has been started. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**Labels** | [**[]LabelSlim**](LabelSlim.md) | An array of Labels attached to the Epic. | [default to null]
**MentionIds** | **[]string** | &#x60;Deprecated:&#x60; use &#x60;member_mention_ids&#x60;. | [default to null]
**MemberMentionIds** | **[]string** | An array of Member IDs that have been mentioned in the Epic description. | [default to null]
**AssociatedGroups** | [**[]EpicAssociatedGroup**](EpicAssociatedGroup.md) | An array containing Group IDs and Group-owned story counts for the Epic&#x27;s associated groups. | [default to null]
**ProjectIds** | **[]int64** | The IDs of Projects related to this Epic. | [default to null]
**StoriesWithoutProjects** | **int64** | The number of stories in this epic which are not associated with a project. | [default to null]
**CompletedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Epic was completed. | [default to null]
**ProductboardPluginId** | **string** | The ID of the associated productboard integration. | [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | The time/date the Epic was started. | [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) | The time/date the Epic was completed. | [default to null]
**ObjectiveIds** | **[]int64** | An array of IDs for Objectives to which this epic is related. | [default to null]
**Name** | **string** | The name of the Epic. | [default to null]
**GlobalId** | **string** |  | [default to null]
**Completed** | **bool** | A true/false boolean indicating if the Epic has been completed. | [default to null]
**ProductboardUrl** | **string** | The URL of the associated productboard feature. | [default to null]
**PlannedStartDate** | [**time.Time**](time.Time.md) | The Epic&#x27;s planned start date. | [default to null]
**State** | **string** | &#x60;Deprecated&#x60; The workflow state that the Epic is in. | [default to null]
**MilestoneId** | **int64** | &#x60;Deprecated&#x60; The ID of the Objective this Epic is related to. Use &#x60;objective_ids&#x60;. | [default to null]
**RequestedById** | **string** | The ID of the Member that requested the epic. | [default to null]
**EpicStateId** | **int64** | The ID of the Epic State. | [default to null]
**LabelIds** | **[]int64** | An array of Label ids attached to the Epic. | [default to null]
**StartedAtOverride** | [**time.Time**](time.Time.md) | A manual override for the time/date the Epic was started. | [default to null]
**GroupId** | **string** | &#x60;Deprecated&#x60; The ID of the group to associate with the epic. Use &#x60;group_ids&#x60;. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Epic was updated. | [default to null]
**GroupMentionIds** | **[]string** | An array of Group IDs that have been mentioned in the Epic description. | [default to null]
**ProductboardId** | **string** | The ID of the associated productboard feature. | [default to null]
**FollowerIds** | **[]string** | An array of UUIDs for any Members you want to add as Followers on this Epic. | [default to null]
**GroupIds** | **[]string** | An array of UUIDS for Groups to which this Epic is related. | [default to null]
**OwnerIds** | **[]string** | An array of UUIDs for any members you want to add as Owners on this new Epic. | [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the Epic has been imported from another tool, the ID in the other tool can be indicated here. | [default to null]
**Id** | **int64** | The unique ID of the Epic. | [default to null]
**Position** | **int64** | The Epic&#x27;s relative position in the Epic workflow state. | [default to null]
**ProductboardName** | **string** | The name of the associated productboard feature. | [default to null]
**Deadline** | [**time.Time**](time.Time.md) | The Epic&#x27;s deadline. | [default to null]
**Stats** | [***EpicStats**](EpicStats.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Epic was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

