# GlipPostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a post | [optional] 
**GroupId** | **string** | Internal identifier of a group a post belongs to | [optional] 
**Type** | **string** | Type of a post | [optional] 
**Text** | **string** | For &#39;TextMessage&#39; post type only. Text of a message | [optional] 
**CreatorId** | **string** | Internal identifier of a user - author of a post | [optional] 
**AddedPersonIds** | **[]string** | For &#39;PersonsAdded&#39; post type only. Identifiers of persons added to a group | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | Post creation datetime in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format | [optional] 
**LastModifiedTime** | [**time.Time**](time.Time.md) | Post last modification datetime in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format | [optional] 
**Attachments** | [**[]GlipMessageAttachmentInfo**](GlipMessageAttachmentInfo.md) | List of posted attachments | [optional] 
**Mentions** | [**[]GlipMentionsInfo**](GlipMentionsInfo.md) |  | [optional] 
**Activity** | **string** | Label of activity type | [optional] 
**Title** | **string** | Title of a message. (Can be set for bot&#39;s messages only) | [optional] 
**IconUri** | **string** | Link to an image used as an icon for this message | [optional] 
**IconEmoji** | **string** | Emoji used as an icon for this message | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


