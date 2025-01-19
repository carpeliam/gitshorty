# LinkedFile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the file. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**StoryIds** | **[]int64** | The IDs of the stories this file is attached to. | [default to null]
**MentionIds** | **[]string** | &#x60;Deprecated:&#x60; use &#x60;member_mention_ids&#x60;. | [default to null]
**MemberMentionIds** | **[]string** | The members that are mentioned in the description of the file. | [default to null]
**Name** | **string** | The name of the linked file. | [default to null]
**ThumbnailUrl** | **string** | The URL of the file thumbnail, if the integration provided it. | [default to null]
**Type_** | **string** | The integration type (e.g. google, dropbox, box). | [default to null]
**Size** | **int64** | The filesize, if the integration provided it. | [default to null]
**UploaderId** | **string** | The UUID of the member that uploaded the file. | [default to null]
**ContentType** | **string** | The content type of the image (e.g. txt/plain). | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the LinkedFile was updated. | [default to null]
**GroupMentionIds** | **[]string** | The groups that are mentioned in the description of the file. | [default to null]
**Id** | **int64** | The unique identifier for the file. | [default to null]
**Url** | **string** | The URL of the file. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the LinkedFile was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

