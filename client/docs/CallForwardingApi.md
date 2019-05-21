# \CallForwardingApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForwardingNumber**](CallForwardingApi.md#CreateForwardingNumber) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number | Create Forwarding Number
[**DeleteForwardingNumber**](CallForwardingApi.md#DeleteForwardingNumber) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number/{forwardingNumberId} | Delete Forwarding Number
[**ListForwardingNumbers**](CallForwardingApi.md#ListForwardingNumbers) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number | Get Forwarding Number List
[**ReadForwardingNumber**](CallForwardingApi.md#ReadForwardingNumber) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number/{forwardingNumberId} | Get Forwarding Number
[**UpdateForwardingNumber**](CallForwardingApi.md#UpdateForwardingNumber) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/forwarding-number/{forwardingNumberId} | Update Forwarding Number



## CreateForwardingNumber

> ForwardingNumberInfo CreateForwardingNumber(ctx, accountId, extensionId, body)
Create Forwarding Number

Adds a new forwarding number to the forwarding number list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**CreateForwardingNumberRequest**](CreateForwardingNumberRequest.md)| JSON body | 

### Return type

[**ForwardingNumberInfo**](ForwardingNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForwardingNumber

> DeleteForwardingNumber(ctx, accountId, extensionId, forwardingNumberId)
Delete Forwarding Number

Deletes a forwarding number from the forwarding number list by its ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**forwardingNumberId** | **string**| Internal identifier of a forwarding number | 

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


## ListForwardingNumbers

> GetExtensionForwardingNumberListResponse ListForwardingNumbers(ctx, accountId, extensionId, optional)
Get Forwarding Number List

Returns the list of extension phone numbers used for call forwarding and call flip. The returned list contains all the extension phone numbers used for call forwarding and call flip.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListForwardingNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListForwardingNumbersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted. | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items). | [default to 100]

### Return type

[**GetExtensionForwardingNumberListResponse**](GetExtensionForwardingNumberListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadForwardingNumber

> ForwardingNumberResource ReadForwardingNumber(ctx, forwardingNumberId, extensionId, accountId)
Get Forwarding Number

Returns a specific forwarding number.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forwardingNumberId** | **string**|  | 
**extensionId** | **string**|  | 
**accountId** | **string**|  | 

### Return type

[**ForwardingNumberResource**](ForwardingNumberResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateForwardingNumber

> ForwardingNumberInfo UpdateForwardingNumber(ctx, accountId, extensionId, forwardingNumberId, body)
Update Forwarding Number

Updates the existing forwarding number from the forwarding number list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**forwardingNumberId** | **string**| Internal identifier of a forwarding number; returned in response in the &#39;id&#39; field | 
**body** | [**UpdateForwardingNumberRequest**](UpdateForwardingNumberRequest.md)| JSON body | 

### Return type

[**ForwardingNumberInfo**](ForwardingNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

