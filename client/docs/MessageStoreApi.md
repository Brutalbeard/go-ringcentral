# \MessageStoreApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessage**](MessageStoreApi.md#DeleteMessage) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | Delete Message
[**DeleteMessageByFilter**](MessageStoreApi.md#DeleteMessageByFilter) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store | Delete Conversation
[**ListMessages**](MessageStoreApi.md#ListMessages) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store | Get Message List
[**ReadMessage**](MessageStoreApi.md#ReadMessage) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | Get Message
[**ReadMessageContent**](MessageStoreApi.md#ReadMessageContent) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId}/content/{attachmentId} | Get Message Content
[**ReadMessageStoreConfiguration**](MessageStoreApi.md#ReadMessageStoreConfiguration) | **Get** /restapi/v1.0/account/{accountId}/message-store-configuration | Get Message Store Configuration
[**SyncMessages**](MessageStoreApi.md#SyncMessages) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-sync | Sync Messages
[**UpdateMessage**](MessageStoreApi.md#UpdateMessage) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/message-store/{messageId} | Update Message(s)
[**UpdateMessageStoreConfiguration**](MessageStoreApi.md#UpdateMessageStoreConfiguration) | **Put** /restapi/v1.0/account/{accountId}/message-store-configuration | Update Message Store Configuration



## DeleteMessage

> DeleteMessage(ctx, accountId, extensionId, messageId, optional)
Delete Message

Deletes message(s) by the given message ID(s). The first call of this method transfers the message to the 'Delete' status. The second call transfers the deleted message to the 'Purged' status. If it is required to make the message 'Purged' immediately (from the first call), then set the query parameter purge to 'True'.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**messageId** | **int32**| Internal identifier of a message | 
 **optional** | ***DeleteMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **purge** | **optional.Bool**| If the value is &#39;True&#39;, then the message is purged immediately with all the attachments | [default to false]
 **conversationId** | **optional.Int32**| Internal identifier of a message thread | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessageByFilter

> DeleteMessageByFilter(ctx, extensionId, accountId, optional)
Delete Conversation

Deletes conversation(s) by conversation ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**|  | 
**accountId** | **string**|  | 
 **optional** | ***DeleteMessageByFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteMessageByFilterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conversationId** | [**optional.Interface of []string**](string.md)|  | 
 **dateTo** | **optional.Time**| Messages received earlier then the date specified will be deleted. The default value is &#39;Now&#39; | 
 **type_** | **optional.String**| Type of messages to be deleted | [default to All]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessages

> GetMessageList ListMessages(ctx, accountId, extensionId, optional)
Get Message List

Returns the list of messages from an extension mailbox.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **availability** | [**optional.Interface of []string**](string.md)| Specifies the availability status for the resulting messages. Multiple values are accepted | 
 **conversationId** | **optional.Int32**| Specifies the conversation identifier for the resulting messages | 
 **dateFrom** | **optional.Time**| The start datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | 
 **dateTo** | **optional.Time**| The end datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **direction** | [**optional.Interface of []string**](string.md)| The direction for the resulting messages. If not specified, both inbound and outbound messages are returned. Multiple values are accepted | 
 **distinctConversations** | **optional.Bool**| If &#39;True&#39;, then the latest messages per every conversation ID are returned | 
 **messageType** | [**optional.Interface of []string**](string.md)| The type of the resulting messages. If not specified, all messages without message type filtering are returned. Multiple values are accepted | 
 **readStatus** | [**optional.Interface of []string**](string.md)| The read status for the resulting messages. Multiple values are accepted | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **phoneNumber** | **optional.String**| The phone number. If specified, messages are returned for this particular phone number only | 

### Return type

[**GetMessageList**](GetMessageList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMessage

> GetMessageInfoResponse ReadMessage(ctx, accountId, extensionId, messageId)
Get Message

Returns individual message record(s) by the given message ID(s). The length of inbound messages is unlimited. Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**messageId** | **int32**| Internal identifier of a message | 

### Return type

[**GetMessageInfoResponse**](GetMessageInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMessageContent

> *os.File ReadMessageContent(ctx, accountId, extensionId, attachmentId, messageId, optional)
Get Message Content

Returns a specific message attachment data as media stream.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**attachmentId** | **int32**| Internal identifier of a message attachment | 
**messageId** | **int32**| Internal identifier of a message | 
 **optional** | ***ReadMessageContentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadMessageContentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **range_** | **optional.String**|  | 
 **contentDisposition** | **optional.String**| Content disposition of a response | [default to Inline]

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, image/tiff

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMessageStoreConfiguration

> MessageStoreConfiguration ReadMessageStoreConfiguration(ctx, accountId)
Get Message Store Configuration

Returns message store settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**MessageStoreConfiguration**](MessageStoreConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncMessages

> GetMessageSyncResponse SyncMessages(ctx, accountId, extensionId, optional)
Sync Messages

Synchronizes messages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***SyncMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncMessagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **conversationId** | **optional.Int32**| Conversation identifier for the resulting messages. Meaningful for SMS and Pager messages only. | 
 **dateFrom** | **optional.Time**| The start datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | 
 **dateTo** | **optional.Time**| The end datetime for resulting messages in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **direction** | [**optional.Interface of []string**](string.md)| Direction for the resulting messages. If not specified, both inbound and outbound messages are returned. Multiple values are accepted | 
 **distinctConversations** | **optional.Bool**| If &#39;True&#39;, then the latest messages per every conversation ID are returned | 
 **messageType** | [**optional.Interface of []string**](string.md)| Type for the resulting messages. If not specified, all types of messages are returned. Multiple values are accepted | 
 **recordCount** | **optional.Int32**| Limits the number of records to be returned (works in combination with dateFrom and dateTo if specified) | 
 **syncToken** | **optional.String**| Value of syncToken property of last sync request response | 
 **syncType** | [**optional.Interface of []string**](string.md)| Type of message synchronization | 

### Return type

[**GetMessageSyncResponse**](GetMessageSyncResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> GetMessageInfoResponse UpdateMessage(ctx, accountId, extensionId, messageId, body)
Update Message(s)

Updates message(s) by ID(s). Batch request is supported, see Batch Requests for details. Currently, only the message read status updating is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**messageId** | **int32**| Internal identifier of a message | 
**body** | [**UpdateMessageRequest**](UpdateMessageRequest.md)| JSON body | 

### Return type

[**GetMessageInfoResponse**](GetMessageInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessageStoreConfiguration

> MessageStoreConfiguration UpdateMessageStoreConfiguration(ctx, accountId, body)
Update Message Store Configuration

Updates message store settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
**body** | [**MessageStoreConfiguration**](MessageStoreConfiguration.md)| JSON body | 

### Return type

[**MessageStoreConfiguration**](MessageStoreConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

