# CustomField

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A string description of the CustomField | [optional] [default to null]
**IconSetIdentifier** | **string** | A string that represents the icon that corresponds to this custom field. | [optional] [default to null]
**EntityType** | **string** | A string description of this resource. | [default to null]
**StoryTypes** | **[]string** | The types of stories this CustomField is scoped to. | [optional] [default to null]
**Name** | **string** | The name of the Custom Field. | [default to null]
**FixedPosition** | **bool** | When true, the CustomFieldEnumValues may not be reordered. | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The instant when this CustomField was last updated. | [default to null]
**Id** | **string** | The unique public ID for the CustomField. | [default to null]
**Values** | [**[]CustomFieldEnumValue**](CustomFieldEnumValue.md) | A collection of legal values for a CustomField. | [optional] [default to null]
**FieldType** | **string** | The type of Custom Field, eg. &#x27;enum&#x27;. | [default to null]
**Position** | **int64** | An integer indicating the position of this Custom Field with respect to the other CustomField | [default to null]
**CanonicalName** | **string** | The canonical name for a Shortcut-defined field. | [optional] [default to null]
**Enabled** | **bool** | When true, the CustomField can be applied to entities in the Workspace. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The instant when this CustomField was created. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

