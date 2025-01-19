# Member

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | The Member&#x27;s role in the Workspace. | [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**Disabled** | **bool** | True/false boolean indicating whether the Member has been disabled within the Workspace. | [default to null]
**GlobalId** | **string** |  | [default to null]
**State** | **string** | The user state, one of partial, full, disabled, or imported.  A partial user is disabled, has no means to log in, and is not an import user.  A full user is enabled and has a means to log in.  A disabled user is disabled and has a means to log in.  An import user is disabled, has no means to log in, and is marked as an import user. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The time/date the Member was last updated. | [default to null]
**CreatedWithoutInvite** | **bool** | Whether this member was created as a placeholder entity. | [default to null]
**GroupIds** | **[]string** | The Member&#x27;s group ids | [default to null]
**Id** | **string** | The Member&#x27;s ID in Shortcut. | [default to null]
**Profile** | [***Profile**](Profile.md) |  | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time/date the Member was created. | [default to null]
**ReplacedBy** | **string** | The id of the member that replaces this one when merged. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

