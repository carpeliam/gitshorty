# PullRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | **string** | A string description of this resource. | [default to null]
**Closed** | **bool** | True/False boolean indicating whether the VCS pull request has been closed. | [default to null]
**Merged** | **bool** | True/False boolean indicating whether the VCS pull request has been merged. | [default to null]
**NumAdded** | **int64** | Number of lines added in the pull request, according to VCS. | [default to null]
**BranchId** | **int64** | The ID of the branch for the particular pull request. | [default to null]
**OverlappingStories** | **[]int64** | An array of Story ids that have Pull Requests that change at least one of the same lines this Pull Request changes. | [optional] [default to null]
**Number** | **int64** | The pull request&#x27;s unique number ID in VCS. | [default to null]
**BranchName** | **string** | The name of the branch for the particular pull request. | [default to null]
**TargetBranchName** | **string** | The name of the target branch for the particular pull request. | [default to null]
**NumCommits** | **int64** | The number of commits on the pull request. | [default to null]
**Title** | **string** | The title of the pull request. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the pull request was created. | [default to null]
**HasOverlappingStories** | **bool** | Boolean indicating that the Pull Request has Stories that have Pull Requests that change at least one of the same lines this Pull Request changes. | [default to null]
**Draft** | **bool** | True/False boolean indicating whether the VCS pull request is in the draft state. | [default to null]
**Id** | **int64** | The unique ID associated with the pull request in Shortcut. | [default to null]
**VcsLabels** | [**[]PullRequestLabel**](PullRequestLabel.md) | An array of PullRequestLabels attached to the PullRequest. | [optional] [default to null]
**Url** | **string** | The URL for the pull request. | [default to null]
**NumRemoved** | **int64** | Number of lines removed in the pull request, according to VCS. | [default to null]
**ReviewStatus** | **string** | The status of the review for the pull request. | [optional] [default to null]
**NumModified** | **int64** | Number of lines modified in the pull request, according to VCS. | [default to null]
**BuildStatus** | **string** | The status of the Continuous Integration workflow for the pull request. | [optional] [default to null]
**TargetBranchId** | **int64** | The ID of the target branch for the particular pull request. | [default to null]
**RepositoryId** | **int64** | The ID of the repository for the particular pull request. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the pull request was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

