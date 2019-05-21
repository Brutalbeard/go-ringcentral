# \ConversationsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlipConversation**](ConversationsApi.md#CreateGlipConversation) | **Post** /restapi/v1.0/glip/conversations | Create/Open Conversation
[**ListGlipConversations**](ConversationsApi.md#ListGlipConversations) | **Get** /restapi/v1.0/glip/conversations | Get Conversations
[**ReadGlipConversation**](ConversationsApi.md#ReadGlipConversation) | **Get** /restapi/v1.0/glip/conversations/{chatId} | Get Conversation



## CreateGlipConversation

> GlipConversationInfo CreateGlipConversation(ctx, body)
Create/Open Conversation

Creates a new conversation or opens an existing one.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateGlipConversationRequest**](CreateGlipConversationRequest.md)| JSON body | 

### Return type

[**GlipConversationInfo**](GlipConversationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlipConversations

> GlipConversationsList ListGlipConversations(ctx, optional)
Get Conversations

Returns the list of conversations where the user is a member. All records in response are sorted by creation time of a chat in ascending order.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListGlipConversationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGlipConversationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordCount** | **optional.Int32**| Number of conversations to be fetched by one request. The maximum value is 250, by default - 30 | [default to 30]
 **pageToken** | **optional.String**| Pagination token. | 

### Return type

[**GlipConversationsList**](GlipConversationsList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipConversation

> GlipConversationInfo ReadGlipConversation(ctx, chatId)
Get Conversation

Returns information about a conversation by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a conversation to be returned. | 

### Return type

[**GlipConversationInfo**](GlipConversationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

