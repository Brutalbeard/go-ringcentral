# \RingOutApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRingOutCall**](RingOutApi.md#CreateRingOutCall) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ring-out | Make RingOut Call
[**CreateRingOutCallDeprecated**](RingOutApi.md#CreateRingOutCallDeprecated) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout | Make RingOut Call
[**DeleteRingOutCall**](RingOutApi.md#DeleteRingOutCall) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ring-out/{ringoutId} | Cancel RingOut Call
[**DeleteRingOutCallDeprecated**](RingOutApi.md#DeleteRingOutCallDeprecated) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId} | Cancel RingOut Call
[**ReadRingOutCallStatus**](RingOutApi.md#ReadRingOutCallStatus) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ring-out/{ringoutId} | Get RingOut Call Status
[**ReadRingOutCallStatusDeprecated**](RingOutApi.md#ReadRingOutCallStatusDeprecated) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/ringout/{ringoutId} | Get RingOut Call Status



## CreateRingOutCall

> GetRingOutStatusResponse CreateRingOutCall(ctx, accountId, extensionId, body)
Make RingOut Call

Makes a 2-leg RingOut call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**MakeRingOutRequest**](MakeRingOutRequest.md)| JSON body | 

### Return type

[**GetRingOutStatusResponse**](GetRingOutStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRingOutCallDeprecated

> GetRingOutStatusResponseIntId CreateRingOutCallDeprecated(ctx, accountId, extensionId, body)
Make RingOut Call

Makes a 2-leg RingOut call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**MakeRingOutRequest**](MakeRingOutRequest.md)| JSON body | 

### Return type

[**GetRingOutStatusResponseIntId**](GetRingOutStatusResponseIntId.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRingOutCall

> DeleteRingOutCall(ctx, accountId, extensionId, ringoutId)
Cancel RingOut Call

Cancels a 2-leg RingOut call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**ringoutId** | **string**| Internal identifier of a RingOut call | 

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


## DeleteRingOutCallDeprecated

> DeleteRingOutCallDeprecated(ctx, accountId, extensionId, ringoutId)
Cancel RingOut Call

Cancels a 2-leg RingOut call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**ringoutId** | **int32**| Internal identifier of a RingOut call | 

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


## ReadRingOutCallStatus

> GetRingOutStatusResponse ReadRingOutCallStatus(ctx, accountId, extensionId, ringoutId)
Get RingOut Call Status

Returns the status of a 2-leg RingOut call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**ringoutId** | **string**| Internal identifier of a RingOut call | 

### Return type

[**GetRingOutStatusResponse**](GetRingOutStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadRingOutCallStatusDeprecated

> GetRingOutStatusResponse ReadRingOutCallStatusDeprecated(ctx, accountId, extensionId, ringoutId)
Get RingOut Call Status

Returns status of a 2-leg RingOut call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**ringoutId** | **int32**| Internal identifier of a RingOut call | 

### Return type

[**GetRingOutStatusResponse**](GetRingOutStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

