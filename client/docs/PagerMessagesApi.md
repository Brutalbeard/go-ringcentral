# \PagerMessagesApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternalTextMessage**](PagerMessagesApi.md#CreateInternalTextMessage) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/company-pager | Create Internal Text Message



## CreateInternalTextMessage

> GetMessageInfoResponse CreateInternalTextMessage(ctx, accountId, extensionId, body)
Create Internal Text Message

Creates and sends an internal text message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**CreateInternalTextMessageRequest**](CreateInternalTextMessageRequest.md)| JSON body | 

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

