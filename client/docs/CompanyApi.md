# \CompanyApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAccountBusinessAddress**](CompanyApi.md#ReadAccountBusinessAddress) | **Get** /restapi/v1.0/account/{accountId}/business-address | Get Account Business Address
[**ReadAccountInfo**](CompanyApi.md#ReadAccountInfo) | **Get** /restapi/v1.0/account/{accountId} | Get Account Info
[**ReadAccountServiceInfo**](CompanyApi.md#ReadAccountServiceInfo) | **Get** /restapi/v1.0/account/{accountId}/service-info | Get Account Service Info
[**UpdateAccountBusinessAddress**](CompanyApi.md#UpdateAccountBusinessAddress) | **Put** /restapi/v1.0/account/{accountId}/business-address | Update Company Business Address



## ReadAccountBusinessAddress

> AccountBusinessAddressResource ReadAccountBusinessAddress(ctx, accountId)
Get Account Business Address

Returns business address of a company.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**AccountBusinessAddressResource**](AccountBusinessAddressResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccountInfo

> GetAccountInfoResponse ReadAccountInfo(ctx, accountId)
Get Account Info

Returns basic information about a particular RingCentral customer account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**GetAccountInfoResponse**](GetAccountInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccountServiceInfo

> GetServiceInfoResponse ReadAccountServiceInfo(ctx, accountId)
Get Account Service Info

Returns the information about service plan, available features and limitations for a particular RingCentral customer account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**GetServiceInfoResponse**](GetServiceInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountBusinessAddress

> AccountBusinessAddressResource UpdateAccountBusinessAddress(ctx, accountId, body)
Update Company Business Address

Updates the business address of a company that account is linked to. Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**body** | [**ModifyAccountBusinessAddressRequest**](ModifyAccountBusinessAddressRequest.md)| JSON body | 

### Return type

[**AccountBusinessAddressResource**](AccountBusinessAddressResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

