# InstantMessageEventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a message | [optional] 
**To** | [**[]NotificationRecipientInfo**](NotificationRecipientInfo.md) | Message receiver(s) information | [optional] 
**From** | [**SenderInfo**](SenderInfo.md) |  | [optional] 
**Type** | **string** | Type of a message. The default value is &#39;SMS&#39; | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | Message creation datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**LastModifiedTime** | [**time.Time**](time.Time.md) | Datetime when the message was modified in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**ReadStatus** | **string** | Status of a message | [optional] [default to Unread]
**Priority** | **string** |  | [optional] [default to Normal]
**Attachments** | [**[]MessageAttachmentInfo**](MessageAttachmentInfo.md) | Message attachment data | [optional] 
**Direction** | **string** | Message direction | [optional] [default to Inbound]
**Availability** | **string** | Message availability status | [optional] [default to Alive]
**Subject** | **string** | Message subject. It replicates message text which is also returned as an attachment | [optional] 
**MessageStatus** | **string** | Status of a message | [optional] [default to Received]
**ConversationId** | **string** | Deprecated. Identifier of a conversation the message belongs to | [optional] 
**Conversation** | [**ConversationInfo**](ConversationInfo.md) |  | [optional] 
**OwnerId** | **string** | Internal identifier of a subscription owner extension | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


