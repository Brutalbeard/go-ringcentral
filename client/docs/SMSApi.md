# \SMSApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSMSMessage**](SMSApi.md#CreateSMSMessage) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/sms | Create SMS/MMS Message



## CreateSMSMessage

> GetMessageInfoResponse CreateSMSMessage(ctx, accountId, extensionId, body)
Create SMS/MMS Message

Creates and sends new SMS message. Sending SMS messages simultaneously to different recipients is limited up to 50 requests per minute; relevant for all client applications. Note: Sending and receiving SMS is now available for Toll-Free Numbers (within the USA)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**CreateSmsMessage**](CreateSmsMessage.md)| JSON body | 

### Return type

[**GetMessageInfoResponse**](GetMessageInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/mixed, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

