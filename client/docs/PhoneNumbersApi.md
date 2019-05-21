# \PhoneNumbersApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountPhoneNumbers**](PhoneNumbersApi.md#ListAccountPhoneNumbers) | **Get** /restapi/v1.0/account/{accountId}/phone-number | Get Company Phone Number List
[**ListExtensionPhoneNumbers**](PhoneNumbersApi.md#ListExtensionPhoneNumbers) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/phone-number | Get Extension Phone Number List
[**ParsePhoneNumber**](PhoneNumbersApi.md#ParsePhoneNumber) | **Post** /restapi/v1.0/number-parser/parse | Parse Phone Number
[**ReadAccountPhoneNumber**](PhoneNumbersApi.md#ReadAccountPhoneNumber) | **Get** /restapi/v1.0/account/{accountId}/phone-number/{phoneNumberId} | Get Phone Number



## ListAccountPhoneNumbers

> AccountPhoneNumbers ListAccountPhoneNumbers(ctx, accountId, optional)
Get Company Phone Number List

Returns the list of phone numbers assigned to RingCentral customer account. Both company-level and extension-level numbers are returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***ListAccountPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAccountPhoneNumbersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **usageType** | [**optional.Interface of []string**](string.md)| Usage type of a phone number | 

### Return type

[**AccountPhoneNumbers**](AccountPhoneNumbers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensionPhoneNumbers

> GetExtensionPhoneNumbersResponse ListExtensionPhoneNumbers(ctx, accountId, extensionId, optional)
Get Extension Phone Number List

Returns the list of phone numbers that are used by a particular extension, and can be filtered by the phone number type. The returned list contains all numbers which are directly mapped to a given extension plus the features and also company-level numbers which may be used when performing different operations on behalf of this extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionPhoneNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExtensionPhoneNumbersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **usageType** | [**optional.Interface of []string**](string.md)| Usage type of a phone number | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39; | 
 **perPage** | **optional.Int32**| Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default | 

### Return type

[**GetExtensionPhoneNumbersResponse**](GetExtensionPhoneNumbersResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParsePhoneNumber

> ParsePhoneNumberResponse ParsePhoneNumber(ctx, body, optional)
Parse Phone Number

Returns one or more parsed and/or formatted phone numbers that are passed as a string.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**ParsePhoneNumberRequest**](ParsePhoneNumberRequest.md)| JSON body | 
 **optional** | ***ParsePhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ParsePhoneNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **homeCountry** | **optional.String**| Internal identifier of a home country. The default value is ISO code (ISO 3166) of the user&#39;s home country or brand country, if the user is undefined | 
 **nationalAsPriority** | **optional.Bool**| The default value is &#39;False&#39;. If &#39;True&#39;, the numbers that are closer to the home country are given higher priority | 

### Return type

[**ParsePhoneNumberResponse**](ParsePhoneNumberResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccountPhoneNumber

> CompanyPhoneNumberInfo ReadAccountPhoneNumber(ctx, accountId, phoneNumberId)
Get Phone Number

Returns the phone number(s) belonging to a certain account or extension by phoneNumberId(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**phoneNumberId** | **int32**| Internal identifier of a phone number | 

### Return type

[**CompanyPhoneNumberInfo**](CompanyPhoneNumberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

