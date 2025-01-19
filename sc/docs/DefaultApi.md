# {{classname}}

All URIs are relative to *https://api.app.shortcut.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategory**](DefaultApi.md#CreateCategory) | **Post** /api/v3/categories | Create Category
[**CreateEntityTemplate**](DefaultApi.md#CreateEntityTemplate) | **Post** /api/v3/entity-templates | Create Entity Template
[**CreateEpic**](DefaultApi.md#CreateEpic) | **Post** /api/v3/epics | Create Epic
[**CreateEpicComment**](DefaultApi.md#CreateEpicComment) | **Post** /api/v3/epics/{epic-public-id}/comments | Create Epic Comment
[**CreateEpicCommentComment**](DefaultApi.md#CreateEpicCommentComment) | **Post** /api/v3/epics/{epic-public-id}/comments/{comment-public-id} | Create Epic Comment Comment
[**CreateGroup**](DefaultApi.md#CreateGroup) | **Post** /api/v3/groups | Create Group
[**CreateIteration**](DefaultApi.md#CreateIteration) | **Post** /api/v3/iterations | Create Iteration
[**CreateLabel**](DefaultApi.md#CreateLabel) | **Post** /api/v3/labels | Create Label
[**CreateLinkedFile**](DefaultApi.md#CreateLinkedFile) | **Post** /api/v3/linked-files | Create Linked File
[**CreateMilestone**](DefaultApi.md#CreateMilestone) | **Post** /api/v3/milestones | Create Milestone
[**CreateMultipleStories**](DefaultApi.md#CreateMultipleStories) | **Post** /api/v3/stories/bulk | Create Multiple Stories
[**CreateObjective**](DefaultApi.md#CreateObjective) | **Post** /api/v3/objectives | Create Objective
[**CreateProject**](DefaultApi.md#CreateProject) | **Post** /api/v3/projects | Create Project
[**CreateStory**](DefaultApi.md#CreateStory) | **Post** /api/v3/stories | Create Story
[**CreateStoryComment**](DefaultApi.md#CreateStoryComment) | **Post** /api/v3/stories/{story-public-id}/comments | Create Story Comment
[**CreateStoryFromTemplate**](DefaultApi.md#CreateStoryFromTemplate) | **Post** /api/v3/stories/from-template | Create Story From Template
[**CreateStoryLink**](DefaultApi.md#CreateStoryLink) | **Post** /api/v3/story-links | Create Story Link
[**CreateStoryReaction**](DefaultApi.md#CreateStoryReaction) | **Post** /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions | Create Story Reaction
[**CreateTask**](DefaultApi.md#CreateTask) | **Post** /api/v3/stories/{story-public-id}/tasks | Create Task
[**DeleteCategory**](DefaultApi.md#DeleteCategory) | **Delete** /api/v3/categories/{category-public-id} | Delete Category
[**DeleteCustomField**](DefaultApi.md#DeleteCustomField) | **Delete** /api/v3/custom-fields/{custom-field-public-id} | Delete Custom Field
[**DeleteEntityTemplate**](DefaultApi.md#DeleteEntityTemplate) | **Delete** /api/v3/entity-templates/{entity-template-public-id} | Delete Entity Template
[**DeleteEpic**](DefaultApi.md#DeleteEpic) | **Delete** /api/v3/epics/{epic-public-id} | Delete Epic
[**DeleteEpicComment**](DefaultApi.md#DeleteEpicComment) | **Delete** /api/v3/epics/{epic-public-id}/comments/{comment-public-id} | Delete Epic Comment
[**DeleteFile**](DefaultApi.md#DeleteFile) | **Delete** /api/v3/files/{file-public-id} | Delete File
[**DeleteIteration**](DefaultApi.md#DeleteIteration) | **Delete** /api/v3/iterations/{iteration-public-id} | Delete Iteration
[**DeleteLabel**](DefaultApi.md#DeleteLabel) | **Delete** /api/v3/labels/{label-public-id} | Delete Label
[**DeleteLinkedFile**](DefaultApi.md#DeleteLinkedFile) | **Delete** /api/v3/linked-files/{linked-file-public-id} | Delete Linked File
[**DeleteMilestone**](DefaultApi.md#DeleteMilestone) | **Delete** /api/v3/milestones/{milestone-public-id} | Delete Milestone
[**DeleteMultipleStories**](DefaultApi.md#DeleteMultipleStories) | **Delete** /api/v3/stories/bulk | Delete Multiple Stories
[**DeleteObjective**](DefaultApi.md#DeleteObjective) | **Delete** /api/v3/objectives/{objective-public-id} | Delete Objective
[**DeleteProject**](DefaultApi.md#DeleteProject) | **Delete** /api/v3/projects/{project-public-id} | Delete Project
[**DeleteStory**](DefaultApi.md#DeleteStory) | **Delete** /api/v3/stories/{story-public-id} | Delete Story
[**DeleteStoryComment**](DefaultApi.md#DeleteStoryComment) | **Delete** /api/v3/stories/{story-public-id}/comments/{comment-public-id} | Delete Story Comment
[**DeleteStoryLink**](DefaultApi.md#DeleteStoryLink) | **Delete** /api/v3/story-links/{story-link-public-id} | Delete Story Link
[**DeleteStoryReaction**](DefaultApi.md#DeleteStoryReaction) | **Delete** /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions | Delete Story Reaction
[**DeleteTask**](DefaultApi.md#DeleteTask) | **Delete** /api/v3/stories/{story-public-id}/tasks/{task-public-id} | Delete Task
[**DisableIterations**](DefaultApi.md#DisableIterations) | **Put** /api/v3/iterations/disable | Disable Iterations
[**DisableStoryTemplates**](DefaultApi.md#DisableStoryTemplates) | **Put** /api/v3/entity-templates/disable | Disable Story Templates
[**EnableIterations**](DefaultApi.md#EnableIterations) | **Put** /api/v3/iterations/enable | Enable Iterations
[**EnableStoryTemplates**](DefaultApi.md#EnableStoryTemplates) | **Put** /api/v3/entity-templates/enable | Enable Story Templates
[**GetCategory**](DefaultApi.md#GetCategory) | **Get** /api/v3/categories/{category-public-id} | Get Category
[**GetCurrentMemberInfo**](DefaultApi.md#GetCurrentMemberInfo) | **Get** /api/v3/member | Get Current Member Info
[**GetCustomField**](DefaultApi.md#GetCustomField) | **Get** /api/v3/custom-fields/{custom-field-public-id} | Get Custom Field
[**GetEntityTemplate**](DefaultApi.md#GetEntityTemplate) | **Get** /api/v3/entity-templates/{entity-template-public-id} | Get Entity Template
[**GetEpic**](DefaultApi.md#GetEpic) | **Get** /api/v3/epics/{epic-public-id} | Get Epic
[**GetEpicComment**](DefaultApi.md#GetEpicComment) | **Get** /api/v3/epics/{epic-public-id}/comments/{comment-public-id} | Get Epic Comment
[**GetEpicWorkflow**](DefaultApi.md#GetEpicWorkflow) | **Get** /api/v3/epic-workflow | Get Epic Workflow
[**GetExternalLinkStories**](DefaultApi.md#GetExternalLinkStories) | **Get** /api/v3/external-link/stories | Get External Link Stories
[**GetFile**](DefaultApi.md#GetFile) | **Get** /api/v3/files/{file-public-id} | Get File
[**GetGroup**](DefaultApi.md#GetGroup) | **Get** /api/v3/groups/{group-public-id} | Get Group
[**GetIteration**](DefaultApi.md#GetIteration) | **Get** /api/v3/iterations/{iteration-public-id} | Get Iteration
[**GetKeyResult**](DefaultApi.md#GetKeyResult) | **Get** /api/v3/key-results/{key-result-public-id} | Get Key Result
[**GetLabel**](DefaultApi.md#GetLabel) | **Get** /api/v3/labels/{label-public-id} | Get Label
[**GetLinkedFile**](DefaultApi.md#GetLinkedFile) | **Get** /api/v3/linked-files/{linked-file-public-id} | Get Linked File
[**GetMember**](DefaultApi.md#GetMember) | **Get** /api/v3/members/{member-public-id} | Get Member
[**GetMilestone**](DefaultApi.md#GetMilestone) | **Get** /api/v3/milestones/{milestone-public-id} | Get Milestone
[**GetObjective**](DefaultApi.md#GetObjective) | **Get** /api/v3/objectives/{objective-public-id} | Get Objective
[**GetProject**](DefaultApi.md#GetProject) | **Get** /api/v3/projects/{project-public-id} | Get Project
[**GetRepository**](DefaultApi.md#GetRepository) | **Get** /api/v3/repositories/{repo-public-id} | Get Repository
[**GetStory**](DefaultApi.md#GetStory) | **Get** /api/v3/stories/{story-public-id} | Get Story
[**GetStoryComment**](DefaultApi.md#GetStoryComment) | **Get** /api/v3/stories/{story-public-id}/comments/{comment-public-id} | Get Story Comment
[**GetStoryLink**](DefaultApi.md#GetStoryLink) | **Get** /api/v3/story-links/{story-link-public-id} | Get Story Link
[**GetTask**](DefaultApi.md#GetTask) | **Get** /api/v3/stories/{story-public-id}/tasks/{task-public-id} | Get Task
[**GetWorkflow**](DefaultApi.md#GetWorkflow) | **Get** /api/v3/workflows/{workflow-public-id} | Get Workflow
[**ListCategories**](DefaultApi.md#ListCategories) | **Get** /api/v3/categories | List Categories
[**ListCategoryMilestones**](DefaultApi.md#ListCategoryMilestones) | **Get** /api/v3/categories/{category-public-id}/milestones | List Category Milestones
[**ListCategoryObjectives**](DefaultApi.md#ListCategoryObjectives) | **Get** /api/v3/categories/{category-public-id}/objectives | List Category Objectives
[**ListCustomFields**](DefaultApi.md#ListCustomFields) | **Get** /api/v3/custom-fields | List Custom Fields
[**ListEntityTemplates**](DefaultApi.md#ListEntityTemplates) | **Get** /api/v3/entity-templates | List Entity Templates
[**ListEpicComments**](DefaultApi.md#ListEpicComments) | **Get** /api/v3/epics/{epic-public-id}/comments | List Epic Comments
[**ListEpicStories**](DefaultApi.md#ListEpicStories) | **Get** /api/v3/epics/{epic-public-id}/stories | List Epic Stories
[**ListEpics**](DefaultApi.md#ListEpics) | **Get** /api/v3/epics | List Epics
[**ListFiles**](DefaultApi.md#ListFiles) | **Get** /api/v3/files | List Files
[**ListGroupStories**](DefaultApi.md#ListGroupStories) | **Get** /api/v3/groups/{group-public-id}/stories | List Group Stories
[**ListGroups**](DefaultApi.md#ListGroups) | **Get** /api/v3/groups | List Groups
[**ListIterationStories**](DefaultApi.md#ListIterationStories) | **Get** /api/v3/iterations/{iteration-public-id}/stories | List Iteration Stories
[**ListIterations**](DefaultApi.md#ListIterations) | **Get** /api/v3/iterations | List Iterations
[**ListLabelEpics**](DefaultApi.md#ListLabelEpics) | **Get** /api/v3/labels/{label-public-id}/epics | List Label Epics
[**ListLabelStories**](DefaultApi.md#ListLabelStories) | **Get** /api/v3/labels/{label-public-id}/stories | List Label Stories
[**ListLabels**](DefaultApi.md#ListLabels) | **Get** /api/v3/labels | List Labels
[**ListLinkedFiles**](DefaultApi.md#ListLinkedFiles) | **Get** /api/v3/linked-files | List Linked Files
[**ListMembers**](DefaultApi.md#ListMembers) | **Get** /api/v3/members | List Members
[**ListMilestoneEpics**](DefaultApi.md#ListMilestoneEpics) | **Get** /api/v3/milestones/{milestone-public-id}/epics | List Milestone Epics
[**ListMilestones**](DefaultApi.md#ListMilestones) | **Get** /api/v3/milestones | List Milestones
[**ListObjectiveEpics**](DefaultApi.md#ListObjectiveEpics) | **Get** /api/v3/objectives/{objective-public-id}/epics | List Objective Epics
[**ListObjectives**](DefaultApi.md#ListObjectives) | **Get** /api/v3/objectives | List Objectives
[**ListProjects**](DefaultApi.md#ListProjects) | **Get** /api/v3/projects | List Projects
[**ListRepositories**](DefaultApi.md#ListRepositories) | **Get** /api/v3/repositories | List Repositories
[**ListStories**](DefaultApi.md#ListStories) | **Get** /api/v3/projects/{project-public-id}/stories | List Stories
[**ListStoryComment**](DefaultApi.md#ListStoryComment) | **Get** /api/v3/stories/{story-public-id}/comments | List Story Comment
[**ListWorkflows**](DefaultApi.md#ListWorkflows) | **Get** /api/v3/workflows | List Workflows
[**Search**](DefaultApi.md#Search) | **Get** /api/v3/search | Search
[**SearchEpics**](DefaultApi.md#SearchEpics) | **Get** /api/v3/search/epics | Search Epics
[**SearchIterations**](DefaultApi.md#SearchIterations) | **Get** /api/v3/search/iterations | Search Iterations
[**SearchMilestones**](DefaultApi.md#SearchMilestones) | **Get** /api/v3/search/milestones | Search Milestones
[**SearchObjectives**](DefaultApi.md#SearchObjectives) | **Get** /api/v3/search/objectives | Search Objectives
[**SearchStories**](DefaultApi.md#SearchStories) | **Get** /api/v3/search/stories | Search Stories
[**SearchStoriesOld**](DefaultApi.md#SearchStoriesOld) | **Post** /api/v3/stories/search | Search Stories (Old)
[**StoryHistory**](DefaultApi.md#StoryHistory) | **Get** /api/v3/stories/{story-public-id}/history | Story History
[**UnlinkCommentThreadFromSlack**](DefaultApi.md#UnlinkCommentThreadFromSlack) | **Post** /api/v3/stories/{story-public-id}/comments/{comment-public-id}/unlink-from-slack | Unlink Comment thread from Slack
[**UnlinkProductboardFromEpic**](DefaultApi.md#UnlinkProductboardFromEpic) | **Post** /api/v3/epics/{epic-public-id}/unlink-productboard | Unlink Productboard from Epic
[**UpdateCategory**](DefaultApi.md#UpdateCategory) | **Put** /api/v3/categories/{category-public-id} | Update Category
[**UpdateCustomField**](DefaultApi.md#UpdateCustomField) | **Put** /api/v3/custom-fields/{custom-field-public-id} | Update Custom Field
[**UpdateEntityTemplate**](DefaultApi.md#UpdateEntityTemplate) | **Put** /api/v3/entity-templates/{entity-template-public-id} | Update Entity Template
[**UpdateEpic**](DefaultApi.md#UpdateEpic) | **Put** /api/v3/epics/{epic-public-id} | Update Epic
[**UpdateEpicComment**](DefaultApi.md#UpdateEpicComment) | **Put** /api/v3/epics/{epic-public-id}/comments/{comment-public-id} | Update Epic Comment
[**UpdateFile**](DefaultApi.md#UpdateFile) | **Put** /api/v3/files/{file-public-id} | Update File
[**UpdateGroup**](DefaultApi.md#UpdateGroup) | **Put** /api/v3/groups/{group-public-id} | Update Group
[**UpdateIteration**](DefaultApi.md#UpdateIteration) | **Put** /api/v3/iterations/{iteration-public-id} | Update Iteration
[**UpdateKeyResult**](DefaultApi.md#UpdateKeyResult) | **Put** /api/v3/key-results/{key-result-public-id} | Update Key Result
[**UpdateLabel**](DefaultApi.md#UpdateLabel) | **Put** /api/v3/labels/{label-public-id} | Update Label
[**UpdateLinkedFile**](DefaultApi.md#UpdateLinkedFile) | **Put** /api/v3/linked-files/{linked-file-public-id} | Update Linked File
[**UpdateMilestone**](DefaultApi.md#UpdateMilestone) | **Put** /api/v3/milestones/{milestone-public-id} | Update Milestone
[**UpdateMultipleStories**](DefaultApi.md#UpdateMultipleStories) | **Put** /api/v3/stories/bulk | Update Multiple Stories
[**UpdateObjective**](DefaultApi.md#UpdateObjective) | **Put** /api/v3/objectives/{objective-public-id} | Update Objective
[**UpdateProject**](DefaultApi.md#UpdateProject) | **Put** /api/v3/projects/{project-public-id} | Update Project
[**UpdateStory**](DefaultApi.md#UpdateStory) | **Put** /api/v3/stories/{story-public-id} | Update Story
[**UpdateStoryComment**](DefaultApi.md#UpdateStoryComment) | **Put** /api/v3/stories/{story-public-id}/comments/{comment-public-id} | Update Story Comment
[**UpdateStoryLink**](DefaultApi.md#UpdateStoryLink) | **Put** /api/v3/story-links/{story-link-public-id} | Update Story Link
[**UpdateTask**](DefaultApi.md#UpdateTask) | **Put** /api/v3/stories/{story-public-id}/tasks/{task-public-id} | Update Task
[**UploadFiles**](DefaultApi.md#UploadFiles) | **Post** /api/v3/files | Upload Files

# **CreateCategory**
> Category CreateCategory(ctx, body)
Create Category

Create Category allows you to create a new Category in Shortcut.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCategory**](CreateCategory.md)|  | 

### Return type

[**Category**](Category.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEntityTemplate**
> EntityTemplate CreateEntityTemplate(ctx, body)
Create Entity Template

Create a new entity template for the Workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEntityTemplate**](CreateEntityTemplate.md)| Request parameters for creating an entirely new entity template. | 

### Return type

[**EntityTemplate**](EntityTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEpic**
> Epic CreateEpic(ctx, body)
Create Epic

Create Epic allows you to create a new Epic in Shortcut.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEpic**](CreateEpic.md)|  | 

### Return type

[**Epic**](Epic.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEpicComment**
> ThreadedComment CreateEpicComment(ctx, body, epicPublicId)
Create Epic Comment

This endpoint allows you to create a threaded Comment on an Epic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEpicComment**](CreateEpicComment.md)|  | 
  **epicPublicId** | **int64**| The ID of the associated Epic. | 

### Return type

[**ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEpicCommentComment**
> ThreadedComment CreateEpicCommentComment(ctx, body, epicPublicId, commentPublicId)
Create Epic Comment Comment

This endpoint allows you to create a nested Comment reply to an existing Epic Comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCommentComment**](CreateCommentComment.md)|  | 
  **epicPublicId** | **int64**| The ID of the associated Epic. | 
  **commentPublicId** | **int64**| The ID of the parent Epic Comment. | 

### Return type

[**ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateGroup**
> Group CreateGroup(ctx, body)
Create Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateGroup**](CreateGroup.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIteration**
> Iteration CreateIteration(ctx, body)
Create Iteration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateIteration**](CreateIteration.md)|  | 

### Return type

[**Iteration**](Iteration.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLabel**
> Label CreateLabel(ctx, body)
Create Label

Create Label allows you to create a new Label in Shortcut.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateLabelParams**](CreateLabelParams.md)| Request parameters for creating a Label on a Shortcut Story. | 

### Return type

[**Label**](Label.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLinkedFile**
> LinkedFile CreateLinkedFile(ctx, body)
Create Linked File

Create Linked File allows you to create a new Linked File in Shortcut.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateLinkedFile**](CreateLinkedFile.md)|  | 

### Return type

[**LinkedFile**](LinkedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMilestone**
> Milestone CreateMilestone(ctx, body)
Create Milestone

(Deprecated: Use 'Create Objective') Create Milestone allows you to create a new Milestone in Shortcut.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMilestone**](CreateMilestone.md)|  | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMultipleStories**
> []StorySlim CreateMultipleStories(ctx, body)
Create Multiple Stories

Create Multiple Stories allows you to create multiple stories in a single request using the same syntax as [Create Story](https://developer.shortcut.com/api/rest/v3#create-story).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStories**](CreateStories.md)|  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateObjective**
> Objective CreateObjective(ctx, body)
Create Objective

Create Objective allows you to create a new Objective in Shortcut.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateObjective**](CreateObjective.md)|  | 

### Return type

[**Objective**](Objective.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> Project CreateProject(ctx, body)
Create Project

Create Project is used to create a new Shortcut Project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateProject**](CreateProject.md)|  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStory**
> Story CreateStory(ctx, body)
Create Story

Create Story is used to add a new story to your Shortcut Workspace.     This endpoint requires that either **workflow_state_id** or **project_id** be provided, but will reject the request if both or neither are specified. The workflow_state_id has been marked as required and is the recommended field to specify because we are in the process of sunsetting Projects in Shortcut.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStoryParams**](CreateStoryParams.md)| Request parameters for creating a story. | 

### Return type

[**Story**](Story.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoryComment**
> StoryComment CreateStoryComment(ctx, body, storyPublicId)
Create Story Comment

Create Comment allows you to create a Comment on any Story.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStoryComment**](CreateStoryComment.md)|  | 
  **storyPublicId** | **int64**| The ID of the Story that the Comment is in. | 

### Return type

[**StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoryFromTemplate**
> Story CreateStoryFromTemplate(ctx, body)
Create Story From Template

Create Story From Template is used to add a new story derived from a template to your Shortcut Workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStoryFromTemplateParams**](CreateStoryFromTemplateParams.md)| Request parameters for creating a story from a story template. These parameters are merged with the values derived from the template. | 

### Return type

[**Story**](Story.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoryLink**
> StoryLink CreateStoryLink(ctx, body)
Create Story Link

Story Links (called Story Relationships in the UI) allow you create semantic relationships between two stories. The parameters read like an active voice grammatical sentence:  subject -> verb -> object.  The subject story acts on the object Story; the object story is the direct object of the sentence.  The subject story \"blocks\", \"duplicates\", or \"relates to\" the object story.  Examples: - \"story 5 blocks story 6” -- story 6 is now \"blocked\" until story 5 is moved to a Done workflow state. - \"story 2 duplicates story 1” -- Story 2 represents the same body of work as Story 1 (and should probably be archived). - \"story 7 relates to story 3”

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStoryLink**](CreateStoryLink.md)|  | 

### Return type

[**StoryLink**](StoryLink.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoryReaction**
> []StoryReaction CreateStoryReaction(ctx, body, storyPublicId, commentPublicId)
Create Story Reaction

Create a reaction to a story comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrDeleteStoryReaction**](CreateOrDeleteStoryReaction.md)|  | 
  **storyPublicId** | **int64**| The ID of the Story that the Comment is in. | 
  **commentPublicId** | **int64**| The ID of the Comment. | 

### Return type

[**[]StoryReaction**](StoryReaction.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTask**
> Task CreateTask(ctx, body, storyPublicId)
Create Task

Create Task is used to create a new task in a Story.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateTask**](CreateTask.md)|  | 
  **storyPublicId** | **int64**| The ID of the Story that the Task will be in. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCategory**
> DeleteCategory(ctx, categoryPublicId)
Delete Category

Delete Category can be used to delete any Category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryPublicId** | **int64**| The unique ID of the Category. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomField**
> DeleteCustomField(ctx, customFieldPublicId)
Delete Custom Field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customFieldPublicId** | [**string**](.md)| The unique ID of the CustomField. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEntityTemplate**
> DeleteEntityTemplate(ctx, entityTemplatePublicId)
Delete Entity Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityTemplatePublicId** | [**string**](.md)| The unique ID of the entity template. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEpic**
> DeleteEpic(ctx, epicPublicId)
Delete Epic

Delete Epic can be used to delete the Epic. The only required parameter is Epic ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **epicPublicId** | **int64**| The unique ID of the Epic. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEpicComment**
> DeleteEpicComment(ctx, epicPublicId, commentPublicId)
Delete Epic Comment

This endpoint allows you to delete a Comment from an Epic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **epicPublicId** | **int64**| The ID of the associated Epic. | 
  **commentPublicId** | **int64**| The ID of the Comment. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFile**
> DeleteFile(ctx, filePublicId)
Delete File

Delete File deletes a previously uploaded file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filePublicId** | **int64**| The File’s unique ID. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIteration**
> DeleteIteration(ctx, iterationPublicId)
Delete Iteration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iterationPublicId** | **int64**| The unique ID of the Iteration. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLabel**
> DeleteLabel(ctx, labelPublicId)
Delete Label

Delete Label can be used to delete any Label.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelPublicId** | **int64**| The unique ID of the Label. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLinkedFile**
> DeleteLinkedFile(ctx, linkedFilePublicId)
Delete Linked File

Delete Linked File can be used to delete any previously attached Linked-File.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkedFilePublicId** | **int64**| The unique identifier of the linked file. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMilestone**
> DeleteMilestone(ctx, milestonePublicId)
Delete Milestone

(Deprecated: Use 'Delete Objective') Delete Milestone can be used to delete any Milestone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **milestonePublicId** | **int64**| The ID of the Milestone. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMultipleStories**
> DeleteMultipleStories(ctx, body)
Delete Multiple Stories

Delete Multiple Stories allows you to delete multiple archived stories at once.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteStories**](DeleteStories.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObjective**
> DeleteObjective(ctx, objectivePublicId)
Delete Objective

Delete Objective can be used to delete any Objective.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectivePublicId** | **int64**| The ID of the Objective. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> DeleteProject(ctx, projectPublicId)
Delete Project

Delete Project can be used to delete a Project. Projects can only be deleted if all associated Stories are moved or deleted. In the case that the Project cannot be deleted, you will receive a 422 response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectPublicId** | **int64**| The unique ID of the Project. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStory**
> DeleteStory(ctx, storyPublicId)
Delete Story

Delete Story can be used to delete any Story.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The ID of the Story. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoryComment**
> DeleteStoryComment(ctx, storyPublicId, commentPublicId)
Delete Story Comment

Delete a Comment from any story.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The ID of the Story that the Comment is in. | 
  **commentPublicId** | **int64**| The ID of the Comment. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoryLink**
> DeleteStoryLink(ctx, storyLinkPublicId)
Delete Story Link

Removes the relationship between the stories for the given Story Link.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyLinkPublicId** | **int64**| The unique ID of the Story Link. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoryReaction**
> DeleteStoryReaction(ctx, body, storyPublicId, commentPublicId)
Delete Story Reaction

Delete a reaction from any story comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrDeleteStoryReaction**](CreateOrDeleteStoryReaction.md)|  | 
  **storyPublicId** | **int64**| The ID of the Story that the Comment is in. | 
  **commentPublicId** | **int64**| The ID of the Comment. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTask**
> DeleteTask(ctx, storyPublicId, taskPublicId)
Delete Task

Delete Task can be used to delete any previously created Task on a Story.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The unique ID of the Story this Task is associated with. | 
  **taskPublicId** | **int64**| The unique ID of the Task. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableIterations**
> DisableIterations(ctx, )
Disable Iterations

Disables Iterations for the current workspace

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableStoryTemplates**
> DisableStoryTemplates(ctx, )
Disable Story Templates

Disables the Story Template feature for the Workspace.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableIterations**
> EnableIterations(ctx, )
Enable Iterations

Enables Iterations for the current workspace

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableStoryTemplates**
> EnableStoryTemplates(ctx, )
Enable Story Templates

Enables the Story Template feature for the Workspace.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCategory**
> Category GetCategory(ctx, categoryPublicId)
Get Category

Get Category returns information about the selected Category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryPublicId** | **int64**| The unique ID of the Category. | 

### Return type

[**Category**](Category.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentMemberInfo**
> MemberInfo GetCurrentMemberInfo(ctx, )
Get Current Member Info

Returns information about the authenticated member.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MemberInfo**](MemberInfo.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomField**
> CustomField GetCustomField(ctx, customFieldPublicId)
Get Custom Field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customFieldPublicId** | [**string**](.md)| The unique ID of the CustomField. | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntityTemplate**
> EntityTemplate GetEntityTemplate(ctx, entityTemplatePublicId)
Get Entity Template

Get Entity Template returns information about a given entity template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityTemplatePublicId** | [**string**](.md)| The unique ID of the entity template. | 

### Return type

[**EntityTemplate**](EntityTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEpic**
> Epic GetEpic(ctx, epicPublicId)
Get Epic

Get Epic returns information about the selected Epic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **epicPublicId** | **int64**| The unique ID of the Epic. | 

### Return type

[**Epic**](Epic.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEpicComment**
> ThreadedComment GetEpicComment(ctx, epicPublicId, commentPublicId)
Get Epic Comment

This endpoint returns information about the selected Epic Comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **epicPublicId** | **int64**| The ID of the associated Epic. | 
  **commentPublicId** | **int64**| The ID of the Comment. | 

### Return type

[**ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEpicWorkflow**
> EpicWorkflow GetEpicWorkflow(ctx, )
Get Epic Workflow

Returns the Epic Workflow for the Workspace.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EpicWorkflow**](EpicWorkflow.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalLinkStories**
> []StorySlim GetExternalLinkStories(ctx, externalLink)
Get External Link Stories

Get Stories which have a given External Link associated with them.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalLink** | **string**| The external link associated with one or more stories. | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFile**
> UploadedFile GetFile(ctx, filePublicId)
Get File

Get File returns information about the selected UploadedFile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filePublicId** | **int64**| The File’s unique ID. | 

### Return type

[**UploadedFile**](UploadedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroup**
> Group GetGroup(ctx, groupPublicId)
Get Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupPublicId** | [**string**](.md)| The unique ID of the Group. | 

### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIteration**
> Iteration GetIteration(ctx, iterationPublicId)
Get Iteration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iterationPublicId** | **int64**| The unique ID of the Iteration. | 

### Return type

[**Iteration**](Iteration.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeyResult**
> KeyResult GetKeyResult(ctx, keyResultPublicId)
Get Key Result

Get Key Result returns information about a chosen Key Result.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyResultPublicId** | [**string**](.md)| The ID of the Key Result. | 

### Return type

[**KeyResult**](KeyResult.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLabel**
> Label GetLabel(ctx, labelPublicId)
Get Label

Get Label returns information about the selected Label.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelPublicId** | **int64**| The unique ID of the Label. | 

### Return type

[**Label**](Label.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinkedFile**
> LinkedFile GetLinkedFile(ctx, linkedFilePublicId)
Get Linked File

Get File returns information about the selected Linked File.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **linkedFilePublicId** | **int64**| The unique identifier of the linked file. | 

### Return type

[**LinkedFile**](LinkedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMember**
> Member GetMember(ctx, memberPublicId, optional)
Get Member

Returns information about a Member.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **memberPublicId** | [**string**](.md)| The Member&#x27;s unique ID. | 
 **optional** | ***DefaultApiGetMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgPublicId** | [**optional.Interface of string**](.md)| The unique ID of the Organization to limit the lookup to. | 

### Return type

[**Member**](Member.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMilestone**
> Milestone GetMilestone(ctx, milestonePublicId)
Get Milestone

(Deprecated: Use 'Get Objective') Get Milestone returns information about a chosen Milestone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **milestonePublicId** | **int64**| The ID of the Milestone. | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjective**
> Objective GetObjective(ctx, objectivePublicId)
Get Objective

Get Objective returns information about a chosen Objective.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectivePublicId** | **int64**| The ID of the Objective. | 

### Return type

[**Objective**](Objective.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> Project GetProject(ctx, projectPublicId)
Get Project

Get Project returns information about the selected Project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectPublicId** | **int64**| The unique ID of the Project. | 

### Return type

[**Project**](Project.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepository**
> Repository GetRepository(ctx, repoPublicId)
Get Repository

Get Repository returns information about the selected Repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoPublicId** | **int64**| The unique ID of the Repository. | 

### Return type

[**Repository**](Repository.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStory**
> Story GetStory(ctx, storyPublicId)
Get Story

Get Story returns information about a chosen Story.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The ID of the Story. | 

### Return type

[**Story**](Story.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoryComment**
> StoryComment GetStoryComment(ctx, storyPublicId, commentPublicId)
Get Story Comment

Get Comment is used to get Comment information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The ID of the Story that the Comment is in. | 
  **commentPublicId** | **int64**| The ID of the Comment. | 

### Return type

[**StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoryLink**
> StoryLink GetStoryLink(ctx, storyLinkPublicId)
Get Story Link

Returns the stories and their relationship for the given Story Link.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyLinkPublicId** | **int64**| The unique ID of the Story Link. | 

### Return type

[**StoryLink**](StoryLink.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTask**
> Task GetTask(ctx, storyPublicId, taskPublicId)
Get Task

Returns information about a chosen Task.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The unique ID of the Story this Task is associated with. | 
  **taskPublicId** | **int64**| The unique ID of the Task. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflow**
> Workflow GetWorkflow(ctx, workflowPublicId)
Get Workflow

Get Workflow returns information about a chosen Workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowPublicId** | **int64**| The ID of the Workflow. | 

### Return type

[**Workflow**](Workflow.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCategories**
> []Category ListCategories(ctx, )
List Categories

List Categories returns a list of all Categories and their attributes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Category**](Category.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCategoryMilestones**
> []Milestone ListCategoryMilestones(ctx, categoryPublicId)
List Category Milestones

List Category Milestones returns a list of all Milestones with the Category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryPublicId** | **int64**| The unique ID of the Category. | 

### Return type

[**[]Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCategoryObjectives**
> []Milestone ListCategoryObjectives(ctx, categoryPublicId)
List Category Objectives

Returns a list of all Objectives with the Category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryPublicId** | **int64**| The unique ID of the Category. | 

### Return type

[**[]Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomFields**
> []CustomField ListCustomFields(ctx, )
List Custom Fields

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CustomField**](CustomField.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEntityTemplates**
> []EntityTemplate ListEntityTemplates(ctx, )
List Entity Templates

List all the entity templates for the Workspace.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EntityTemplate**](EntityTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEpicComments**
> []ThreadedComment ListEpicComments(ctx, epicPublicId)
List Epic Comments

Get a list of all Comments on an Epic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **epicPublicId** | **int64**| The unique ID of the Epic. | 

### Return type

[**[]ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEpicStories**
> []StorySlim ListEpicStories(ctx, epicPublicId, optional)
List Epic Stories

Get a list of all Stories in an Epic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **epicPublicId** | **int64**| The unique ID of the Epic. | 
 **optional** | ***DefaultApiListEpicStoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListEpicStoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includesDescription** | **optional.Bool**| A true/false boolean indicating whether to return Stories with their descriptions. | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEpics**
> []EpicSlim ListEpics(ctx, optional)
List Epics

List Epics returns a list of all Epics and their attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiListEpicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListEpicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includesDescription** | **optional.Bool**| A true/false boolean indicating whether to return Epics with their descriptions. | 

### Return type

[**[]EpicSlim**](EpicSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFiles**
> []UploadedFile ListFiles(ctx, )
List Files

List Files returns a list of all UploadedFiles in the workspace.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UploadedFile**](UploadedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupStories**
> []StorySlim ListGroupStories(ctx, groupPublicId, optional)
List Group Stories

List the Stories assigned to the Group. (By default, limited to 1,000).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupPublicId** | [**string**](.md)| The unique ID of the Group. | 
 **optional** | ***DefaultApiListGroupStoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListGroupStoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| The maximum number of results to return. (Defaults to 1000, max 1000) | 
 **offset** | **optional.Int64**| The offset at which to begin returning results. (Defaults to 0) | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroups**
> []Group ListGroups(ctx, )
List Groups

A group in our API maps to a \"Team\" within the Shortcut Product. A Team is a collection of Users that can be associated to Stories, Epics, and Iterations within Shortcut.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIterationStories**
> []StorySlim ListIterationStories(ctx, iterationPublicId, optional)
List Iteration Stories

Get a list of all Stories in an Iteration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **iterationPublicId** | **int64**| The unique ID of the Iteration. | 
 **optional** | ***DefaultApiListIterationStoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListIterationStoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includesDescription** | **optional.Bool**| A true/false boolean indicating whether to return Stories with their descriptions. | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIterations**
> []IterationSlim ListIterations(ctx, )
List Iterations

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]IterationSlim**](IterationSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLabelEpics**
> []EpicSlim ListLabelEpics(ctx, labelPublicId)
List Label Epics

List all of the Epics with the Label.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelPublicId** | **int64**| The unique ID of the Label. | 

### Return type

[**[]EpicSlim**](EpicSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLabelStories**
> []StorySlim ListLabelStories(ctx, labelPublicId, optional)
List Label Stories

List all of the Stories with the Label.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelPublicId** | **int64**| The unique ID of the Label. | 
 **optional** | ***DefaultApiListLabelStoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListLabelStoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includesDescription** | **optional.Bool**| A true/false boolean indicating whether to return Stories with their descriptions. | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLabels**
> []Label ListLabels(ctx, optional)
List Labels

List Labels returns a list of all Labels and their attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiListLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slim** | **optional.Bool**| A true/false boolean indicating if the slim versions of the Label should be returned. | 

### Return type

[**[]Label**](Label.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLinkedFiles**
> []LinkedFile ListLinkedFiles(ctx, )
List Linked Files

List Linked Files returns a list of all Linked-Files and their attributes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LinkedFile**](LinkedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMembers**
> []Member ListMembers(ctx, optional)
List Members

Returns information about members of the Workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiListMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgPublicId** | [**optional.Interface of string**](.md)| The unique ID of the Organization to limit the list to. | 

### Return type

[**[]Member**](Member.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMilestoneEpics**
> []EpicSlim ListMilestoneEpics(ctx, milestonePublicId)
List Milestone Epics

(Deprecated: Use 'List Objective Epics') List all of the Epics within the Milestone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **milestonePublicId** | **int64**| The ID of the Milestone. | 

### Return type

[**[]EpicSlim**](EpicSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMilestones**
> []Milestone ListMilestones(ctx, )
List Milestones

(Deprecated: Use 'List Objectives') List Milestones returns a list of all Milestones and their attributes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectiveEpics**
> []EpicSlim ListObjectiveEpics(ctx, objectivePublicId)
List Objective Epics

List all of the Epics within the Objective.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectivePublicId** | **int64**| The ID of the Objective. | 

### Return type

[**[]EpicSlim**](EpicSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListObjectives**
> []Objective ListObjectives(ctx, )
List Objectives

List Objectives returns a list of all Objectives and their attributes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Objective**](Objective.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjects**
> []Project ListProjects(ctx, )
List Projects

List Projects returns a list of all Projects and their attributes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Project**](Project.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRepositories**
> []Repository ListRepositories(ctx, )
List Repositories

List Repositories returns a list of all Repositories and their attributes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Repository**](Repository.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStories**
> []StorySlim ListStories(ctx, projectPublicId, optional)
List Stories

List Stories returns a list of all Stories in a selected Project and their attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectPublicId** | **int64**| The unique ID of the Project. | 
 **optional** | ***DefaultApiListStoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiListStoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includesDescription** | **optional.Bool**| A true/false boolean indicating whether to return Stories with their descriptions. | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStoryComment**
> []StoryComment ListStoryComment(ctx, storyPublicId)
List Story Comment

Lists Comments associated with a Story

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The ID of the Story that the Comment is in. | 

### Return type

[**[]StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkflows**
> []Workflow ListWorkflows(ctx, )
List Workflows

Returns a list of all Workflows in the Workspace.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Workflow**](Workflow.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Search**
> SearchResults Search(ctx, query, optional)
Search

Search lets you search Epics and Stories based on desired parameters. Since ordering of the results can change over time (due to search ranking decay, new Epics and Stories being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators) | 
 **optional** | ***DefaultApiSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int64**| The number of search results to include in a page. Minimum of 1 and maximum of 25. | 
 **detail** | **optional.String**| The amount of detail included in each result item.    \&quot;full\&quot; will include all descriptions and comments and more fields on    related items such as pull requests, branches and tasks.    \&quot;slim\&quot; omits larger fulltext fields such as descriptions and comments    and only references related items by id.    The default is \&quot;full\&quot;. | 
 **next** | **optional.String**| The next page token. | 
 **entityTypes** | [**optional.Interface of []string**](string.md)| A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, objective, story. | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchEpics**
> EpicSearchResults SearchEpics(ctx, query, optional)
Search Epics

Search Epics lets you search Epics based on desired parameters. Since ordering of stories can change over time (due to search ranking decay, new Epics being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators) | 
 **optional** | ***DefaultApiSearchEpicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSearchEpicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int64**| The number of search results to include in a page. Minimum of 1 and maximum of 25. | 
 **detail** | **optional.String**| The amount of detail included in each result item.    \&quot;full\&quot; will include all descriptions and comments and more fields on    related items such as pull requests, branches and tasks.    \&quot;slim\&quot; omits larger fulltext fields such as descriptions and comments    and only references related items by id.    The default is \&quot;full\&quot;. | 
 **next** | **optional.String**| The next page token. | 
 **entityTypes** | [**optional.Interface of []string**](string.md)| A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, objective, story. | 

### Return type

[**EpicSearchResults**](EpicSearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchIterations**
> IterationSearchResults SearchIterations(ctx, query, optional)
Search Iterations

Search Iterations lets you search Iterations based on desired parameters. Since ordering of results can change over time (due to search ranking decay, new Iterations being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators) | 
 **optional** | ***DefaultApiSearchIterationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSearchIterationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int64**| The number of search results to include in a page. Minimum of 1 and maximum of 25. | 
 **detail** | **optional.String**| The amount of detail included in each result item.    \&quot;full\&quot; will include all descriptions and comments and more fields on    related items such as pull requests, branches and tasks.    \&quot;slim\&quot; omits larger fulltext fields such as descriptions and comments    and only references related items by id.    The default is \&quot;full\&quot;. | 
 **next** | **optional.String**| The next page token. | 
 **entityTypes** | [**optional.Interface of []string**](string.md)| A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, objective, story. | 

### Return type

[**IterationSearchResults**](IterationSearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchMilestones**
> ObjectiveSearchResults SearchMilestones(ctx, query, optional)
Search Milestones

Search Milestones lets you search Milestones based on desired parameters. Since ordering of results can change over time (due to search ranking decay, new Milestones being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators) | 
 **optional** | ***DefaultApiSearchMilestonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSearchMilestonesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int64**| The number of search results to include in a page. Minimum of 1 and maximum of 25. | 
 **detail** | **optional.String**| The amount of detail included in each result item.    \&quot;full\&quot; will include all descriptions and comments and more fields on    related items such as pull requests, branches and tasks.    \&quot;slim\&quot; omits larger fulltext fields such as descriptions and comments    and only references related items by id.    The default is \&quot;full\&quot;. | 
 **next** | **optional.String**| The next page token. | 
 **entityTypes** | [**optional.Interface of []string**](string.md)| A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, objective, story. | 

### Return type

[**ObjectiveSearchResults**](ObjectiveSearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchObjectives**
> ObjectiveSearchResults SearchObjectives(ctx, query, optional)
Search Objectives

Search Objectives lets you search Objectives based on desired parameters. Since ordering of results can change over time (due to search ranking decay, new Objectives being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators) | 
 **optional** | ***DefaultApiSearchObjectivesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSearchObjectivesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int64**| The number of search results to include in a page. Minimum of 1 and maximum of 25. | 
 **detail** | **optional.String**| The amount of detail included in each result item.    \&quot;full\&quot; will include all descriptions and comments and more fields on    related items such as pull requests, branches and tasks.    \&quot;slim\&quot; omits larger fulltext fields such as descriptions and comments    and only references related items by id.    The default is \&quot;full\&quot;. | 
 **next** | **optional.String**| The next page token. | 
 **entityTypes** | [**optional.Interface of []string**](string.md)| A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, objective, story. | 

### Return type

[**ObjectiveSearchResults**](ObjectiveSearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchStories**
> StorySearchResults SearchStories(ctx, query, optional)
Search Stories

Search Stories lets you search Stories based on desired parameters. Since ordering of stories can change over time (due to search ranking decay, new stories being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| See our help center article on [search operators](https://help.shortcut.com/hc/en-us/articles/360000046646-Search-Operators) | 
 **optional** | ***DefaultApiSearchStoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiSearchStoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int64**| The number of search results to include in a page. Minimum of 1 and maximum of 25. | 
 **detail** | **optional.String**| The amount of detail included in each result item.    \&quot;full\&quot; will include all descriptions and comments and more fields on    related items such as pull requests, branches and tasks.    \&quot;slim\&quot; omits larger fulltext fields such as descriptions and comments    and only references related items by id.    The default is \&quot;full\&quot;. | 
 **next** | **optional.String**| The next page token. | 
 **entityTypes** | [**optional.Interface of []string**](string.md)| A collection of entity_types to search. Defaults to story and epic. Supports: epic, iteration, objective, story. | 

### Return type

[**StorySearchResults**](StorySearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchStoriesOld**
> []StorySlim SearchStoriesOld(ctx, body)
Search Stories (Old)

Search Stories lets you search Stories based on desired parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchStories**](SearchStories.md)|  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoryHistory**
> []History StoryHistory(ctx, storyPublicId)
Story History

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The ID of the Story. | 

### Return type

[**[]History**](History.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkCommentThreadFromSlack**
> StoryComment UnlinkCommentThreadFromSlack(ctx, storyPublicId, commentPublicId)
Unlink Comment thread from Slack

Unlinks a Comment from its linked Slack thread (Comment replies and Slack replies will no longer be synced)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyPublicId** | **int64**| The ID of the Story to unlink. | 
  **commentPublicId** | **int64**| The ID of the Comment to unlink. | 

### Return type

[**StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkProductboardFromEpic**
> UnlinkProductboardFromEpic(ctx, epicPublicId)
Unlink Productboard from Epic

This endpoint allows you to unlink a productboard epic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **epicPublicId** | **int64**| The unique ID of the Epic. | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCategory**
> Category UpdateCategory(ctx, body, categoryPublicId)
Update Category

Update Category allows you to replace a Category name with another name. If you try to name a Category something that already exists, you will receive a 422 response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCategory**](UpdateCategory.md)|  | 
  **categoryPublicId** | **int64**| The unique ID of the Category you wish to update. | 

### Return type

[**Category**](Category.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomField**
> CustomField UpdateCustomField(ctx, body, customFieldPublicId)
Update Custom Field

Update Custom Field can be used to update the definition of a Custom Field. The order of items in the 'values' collection is interpreted to be their ascending sort order.To delete an existing enum value, simply omit it from the 'values' collection. New enum values may be created inline by including an object in the 'values' collection having a 'value' entry with no 'id' (eg. {'value': 'myNewValue', 'color_key': 'green'}).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCustomField**](UpdateCustomField.md)|  | 
  **customFieldPublicId** | [**string**](.md)| The unique ID of the CustomField. | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntityTemplate**
> EntityTemplate UpdateEntityTemplate(ctx, body, entityTemplatePublicId)
Update Entity Template

Update an entity template's name or its contents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEntityTemplate**](UpdateEntityTemplate.md)| Request parameters for changing either a template&#x27;s name or any of   the attributes it is designed to pre-populate. | 
  **entityTemplatePublicId** | [**string**](.md)| The unique ID of the template to be updated. | 

### Return type

[**EntityTemplate**](EntityTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEpic**
> Epic UpdateEpic(ctx, body, epicPublicId)
Update Epic

Update Epic can be used to update numerous fields in the Epic. The only required parameter is Epic ID, which can be found in the Shortcut UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEpic**](UpdateEpic.md)|  | 
  **epicPublicId** | **int64**| The unique ID of the Epic. | 

### Return type

[**Epic**](Epic.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEpicComment**
> ThreadedComment UpdateEpicComment(ctx, body, epicPublicId, commentPublicId)
Update Epic Comment

This endpoint allows you to update a threaded Comment on an Epic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateComment**](UpdateComment.md)|  | 
  **epicPublicId** | **int64**| The ID of the associated Epic. | 
  **commentPublicId** | **int64**| The ID of the Comment. | 

### Return type

[**ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFile**
> UploadedFile UpdateFile(ctx, body, filePublicId)
Update File

Update File updates the properties of an UploadedFile (but not its content).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFile**](UpdateFile.md)|  | 
  **filePublicId** | **int64**| The unique ID assigned to the file in Shortcut. | 

### Return type

[**UploadedFile**](UploadedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> Group UpdateGroup(ctx, body, groupPublicId)
Update Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateGroup**](UpdateGroup.md)|  | 
  **groupPublicId** | [**string**](.md)| The unique ID of the Group. | 

### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIteration**
> Iteration UpdateIteration(ctx, body, iterationPublicId)
Update Iteration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateIteration**](UpdateIteration.md)|  | 
  **iterationPublicId** | **int64**| The unique ID of the Iteration. | 

### Return type

[**Iteration**](Iteration.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateKeyResult**
> KeyResult UpdateKeyResult(ctx, body, keyResultPublicId)
Update Key Result

Update Key Result allows updating a Key Result's name or initial, observed, or target values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateKeyResult**](UpdateKeyResult.md)|  | 
  **keyResultPublicId** | [**string**](.md)| The ID of the Key Result. | 

### Return type

[**KeyResult**](KeyResult.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLabel**
> Label UpdateLabel(ctx, body, labelPublicId)
Update Label

Update Label allows you to replace a Label name with another name. If you try to name a Label something that already exists, you will receive a 422 response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateLabel**](UpdateLabel.md)|  | 
  **labelPublicId** | **int64**| The unique ID of the Label you wish to update. | 

### Return type

[**Label**](Label.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLinkedFile**
> LinkedFile UpdateLinkedFile(ctx, body, linkedFilePublicId)
Update Linked File

Updated Linked File allows you to update properties of a previously attached Linked-File.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateLinkedFile**](UpdateLinkedFile.md)|  | 
  **linkedFilePublicId** | **int64**| The unique identifier of the linked file. | 

### Return type

[**LinkedFile**](LinkedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMilestone**
> Milestone UpdateMilestone(ctx, body, milestonePublicId)
Update Milestone

(Deprecated: Use 'Update Objective') Update Milestone can be used to update Milestone properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateMilestone**](UpdateMilestone.md)|  | 
  **milestonePublicId** | **int64**| The ID of the Milestone. | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMultipleStories**
> []StorySlim UpdateMultipleStories(ctx, body)
Update Multiple Stories

Update Multiple Stories allows you to make changes to numerous stories at once.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateStories**](UpdateStories.md)|  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateObjective**
> Objective UpdateObjective(ctx, body, objectivePublicId)
Update Objective

Update Objective can be used to update Objective properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateObjective**](UpdateObjective.md)|  | 
  **objectivePublicId** | **int64**| The ID of the Objective. | 

### Return type

[**Objective**](Objective.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProject**
> Project UpdateProject(ctx, body, projectPublicId)
Update Project

Update Project can be used to change properties of a Project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateProject**](UpdateProject.md)|  | 
  **projectPublicId** | **int64**| The unique ID of the Project. | 

### Return type

[**Project**](Project.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStory**
> Story UpdateStory(ctx, body, storyPublicId)
Update Story

Update Story can be used to update Story properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateStory**](UpdateStory.md)|  | 
  **storyPublicId** | **int64**| The unique identifier of this story. | 

### Return type

[**Story**](Story.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStoryComment**
> StoryComment UpdateStoryComment(ctx, body, storyPublicId, commentPublicId)
Update Story Comment

Update Comment replaces the text of the existing Comment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateStoryComment**](UpdateStoryComment.md)|  | 
  **storyPublicId** | **int64**| The ID of the Story that the Comment is in. | 
  **commentPublicId** | **int64**| The ID of the Comment | 

### Return type

[**StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStoryLink**
> StoryLink UpdateStoryLink(ctx, body, storyLinkPublicId)
Update Story Link

Updates the stories and/or the relationship for the given Story Link.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateStoryLink**](UpdateStoryLink.md)|  | 
  **storyLinkPublicId** | **int64**| The unique ID of the Story Link. | 

### Return type

[**StoryLink**](StoryLink.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTask**
> Task UpdateTask(ctx, body, storyPublicId, taskPublicId)
Update Task

Update Task can be used to update Task properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTask**](UpdateTask.md)|  | 
  **storyPublicId** | **int64**| The unique identifier of the parent Story. | 
  **taskPublicId** | **int64**| The unique identifier of the Task you wish to update. | 

### Return type

[**Task**](Task.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFiles**
> []UploadedFile UploadFiles(ctx, storyId, file0, file1, file2, file3)
Upload Files

Upload Files uploads one or many files and optionally associates them with a story.    Use the multipart/form-data content-type to upload.    Each `file` key should contain a separate file.    Each UploadedFile's name comes from the Content-Disposition header \"filename\" directive for that field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storyId** | **int64**|  | 
  **file0** | ***os.File*****os.File**|  | 
  **file1** | ***os.File*****os.File**|  | 
  **file2** | ***os.File*****os.File**|  | 
  **file3** | ***os.File*****os.File**|  | 

### Return type

[**[]UploadedFile**](UploadedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

