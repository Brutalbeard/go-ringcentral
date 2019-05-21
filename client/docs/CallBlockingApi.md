# \CallBlockingApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlockedAllowedNumber**](CallBlockingApi.md#CreateBlockedAllowedNumber) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-blocking/phone-numbers | Add Blocked/Allowed Number
[**DeleteBlockedAllowedNumber**](CallBlockingApi.md#DeleteBlockedAllowedNumber) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-blocking/phone-numbers/{blockedNumberId} | Delete Blocked/Allowed Number
[**ListBlockedAllowedNumbers**](CallBlockingApi.md#ListBlockedAllowedNumbers) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-blocking/phone-numbers | Get Blocked/Allowed Phone Numbers
[**ReadBlockedAllowedNumber**](CallBlockingApi.md#ReadBlockedAllowedNumber) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-blocking/phone-numbers/{blockedNumberId} | Get Blocked/Allowed Number
[**ReadCallerBlockingSettings**](CallBlockingApi.md#ReadCallerBlockingSettings) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-blocking | Get Caller Blocking Settings
[**UpdateBlockedAllowedNumber**](CallBlockingApi.md#UpdateBlockedAllowedNumber) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-blocking/phone-numbers/{blockedNumberId} | Update Blocked/Allowed Number
[**UpdateCallerBlockingSettings**](CallBlockingApi.md#UpdateCallerBlockingSettings) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-blocking | Update Caller Blocking Settings



## CreateBlockedAllowedNumber

> BlockedAllowedPhoneNumberInfo CreateBlockedAllowedNumber(ctx, accountId, extensionId, body)
Add Blocked/Allowed Number

Updates either blocked or allowed phone number list with a new phone number.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**extensionId** | **string**|  | 
**body** | [**AddBlockedAllowedPhoneNumber**](AddBlockedAllowedPhoneNumber.md)| JSON body | 

### Return type

[**BlockedAllowedPhoneNumberInfo**](BlockedAllowedPhoneNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlockedAllowedNumber

> DeleteBlockedAllowedNumber(ctx, accountId, extensionId, blockedNumberId)
Delete Blocked/Allowed Number

Deletes blocked or allowed phone number(s) by their ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**extensionId** | **string**|  | 
**blockedNumberId** | **string**|  | 

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


## ListBlockedAllowedNumbers

> BlockedAllowedPhoneNumbersList ListBlockedAllowedNumbers(ctx, accountId, extensionId, optional)
Get Blocked/Allowed Phone Numbers

Returns the lists of blocked and allowed phone numbers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**extensionId** | **string**|  | 
 **optional** | ***ListBlockedAllowedNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBlockedAllowedNumbersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**|  | 
 **perPage** | **optional.Int32**|  | 
 **status** | **optional.String**|  | 

### Return type

[**BlockedAllowedPhoneNumbersList**](BlockedAllowedPhoneNumbersList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadBlockedAllowedNumber

> BlockedAllowedPhoneNumberInfo ReadBlockedAllowedNumber(ctx, accountId, extensionId, blockedNumberId)
Get Blocked/Allowed Number

Returns blocked or allowed phone number(s) by their ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**extensionId** | **string**|  | 
**blockedNumberId** | **string**|  | 

### Return type

[**BlockedAllowedPhoneNumberInfo**](BlockedAllowedPhoneNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCallerBlockingSettings

> CallerBlockingSettings ReadCallerBlockingSettings(ctx, accountId, extensionId)
Get Caller Blocking Settings

Returns the current caller blocking settings of a user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**extensionId** | **string**|  | 

### Return type

[**CallerBlockingSettings**](CallerBlockingSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlockedAllowedNumber

> BlockedAllowedPhoneNumberInfo UpdateBlockedAllowedNumber(ctx, accountId, extensionId, blockedNumberId, optional)
Update Blocked/Allowed Number

Updates blocked or allowed phone number(s) by their ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**extensionId** | **string**|  | 
**blockedNumberId** | **string**|  | 
 **optional** | ***UpdateBlockedAllowedNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBlockedAllowedNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of AddBlockedAllowedPhoneNumber**](AddBlockedAllowedPhoneNumber.md)|  | 

### Return type

[**BlockedAllowedPhoneNumberInfo**](BlockedAllowedPhoneNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallerBlockingSettings

> CallerBlockingSettings UpdateCallerBlockingSettings(ctx, accountId, extensionId, optional)
Update Caller Blocking Settings

Updates the current caller blocking settings of a user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**extensionId** | **string**|  | 
 **optional** | ***UpdateCallerBlockingSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCallerBlockingSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CallerBlockingSettingsUpdate**](CallerBlockingSettingsUpdate.md)|  | 

### Return type

[**CallerBlockingSettings**](CallerBlockingSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

