# GlipChatInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a chat | [optional] 
**Type** | **string** | Type of a chat | [optional] 
**Public** | **bool** | For &#39;Team&#39; chat type only. Team access level. | [optional] 
**Name** | **string** | For &#39;Team&#39;,&#39;Everyone&#39; chats types only. Chat name. | [optional] 
**Description** | **string** | For &#39;Team&#39;,&#39;Everyone&#39; chats types only. Chat description. | [optional] 
**Status** | **string** | For &#39;Team&#39; chat type only. Team status. | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | Chat creation datetime in ISO 8601 format | [optional] 
**LastModifiedTime** | [**time.Time**](time.Time.md) | Chat last change datetime in ISO 8601 format | [optional] 
**Members** | [**[]GlipChatMemberInfo**](GlipChatMemberInfo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


