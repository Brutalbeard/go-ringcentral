# \BusinessHoursApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadCompanyBusinessHours**](BusinessHoursApi.md#ReadCompanyBusinessHours) | **Get** /restapi/v1.0/account/{accountId}/business-hours | Get Company Business Hours
[**ReadUserBusinessHours**](BusinessHoursApi.md#ReadUserBusinessHours) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/business-hours | Get User Business Hours
[**UpdateCompanyBusinessHours**](BusinessHoursApi.md#UpdateCompanyBusinessHours) | **Put** /restapi/v1.0/account/{accountId}/business-hours | Update Company Business Hours
[**UpdateUserBusinessHours**](BusinessHoursApi.md#UpdateUserBusinessHours) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/business-hours | Update User Business Hours



## ReadCompanyBusinessHours

> CompanyBusinessHours ReadCompanyBusinessHours(ctx, accountId)
Get Company Business Hours

Returns company hours when answering rules are to be applied.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**CompanyBusinessHours**](CompanyBusinessHours.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserBusinessHours

> GetUserBusinessHoursResponse ReadUserBusinessHours(ctx, accountId, extensionId)
Get User Business Hours

Returns the user hours when the call handling rules are applied. **Please note:** If user hours are set to 'Custom hours' then a particular schedule is returned; however if set to '24 hours/7 days a week' an empty schedule is returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**GetUserBusinessHoursResponse**](GetUserBusinessHoursResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyBusinessHours

> CompanyBusinessHours UpdateCompanyBusinessHours(ctx, accountId, body)
Update Company Business Hours

Updates company hours when answering rules are to be applied.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**body** | [**CompanyBusinessHoursUpdateRequest**](CompanyBusinessHoursUpdateRequest.md)| JSON body | 

### Return type

[**CompanyBusinessHours**](CompanyBusinessHours.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserBusinessHours

> UserBusinessHoursUpdateResponse UpdateUserBusinessHours(ctx, accountId, extensionId, body)
Update User Business Hours

Updates the extension user hours when answering rules are to be applied.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**UserBusinessHoursUpdateRequest**](UserBusinessHoursUpdateRequest.md)| JSON body | 

### Return type

[**UserBusinessHoursUpdateResponse**](UserBusinessHoursUpdateResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

