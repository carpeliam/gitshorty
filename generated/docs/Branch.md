# Branch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | [default to null]
**Deleted** | **bool** | A true/false boolean indicating if the Branch has been deleted. | [default to null]
**Name** | **string** | The name of the Branch. | [default to null]
**Persistent** | **bool** | This field is deprecated, and will always be false. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Branch was updated. | [default to null]
**PullRequests** | [**[]PullRequest**](PullRequest.md) | An array of PullRequests attached to the Branch (there is usually only one). | [default to null]
**MergedBranchIds** | **[]int64** | The IDs of the Branches the Branch has been merged into. | [default to null]
**Id** | **int64** | The unique ID of the Branch. | [default to null]
**Url** | **string** | The URL of the Branch. | [default to null]
**RepositoryId** | **int64** | The ID of the Repository that contains the Branch. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Branch was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

