# UploadedFile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the file. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**StoryIds** | **[]int64** | The unique IDs of the Stories associated with this file. | [default to null]
**MentionIds** | **[]string** | &#x60;Deprecated:&#x60; use &#x60;member_mention_ids&#x60;. | [default to null]
**MemberMentionIds** | **[]string** | The unique IDs of the Members who are mentioned in the file description. | [default to null]
**Name** | **string** | The optional User-specified name of the file. | [default to null]
**ThumbnailUrl** | **string** | The url where the thumbnail of the file can be found in Shortcut. | [default to null]
**Size** | **int64** | The size of the file. | [default to null]
**UploaderId** | **string** | The unique ID of the Member who uploaded the file. | [default to null]
**ContentType** | **string** | Free form string corresponding to a text or image file. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date that the file was updated. | [default to null]
**Filename** | **string** | The name assigned to the file in Shortcut upon upload. | [default to null]
**GroupMentionIds** | **[]string** | The unique IDs of the Groups who are mentioned in the file description. | [default to null]
**ExternalId** | **string** | This field can be set to another unique ID. In the case that the File has been imported from another tool, the ID in the other tool can be indicated here. | [default to null]
**Id** | **int64** | The unique ID for the file. | [default to null]
**Url** | **string** | The URL for the file. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date that the file was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

