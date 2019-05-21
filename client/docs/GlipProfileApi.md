# \GlipProfileApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadGlipCompany**](GlipProfileApi.md#ReadGlipCompany) | **Get** /restapi/v1.0/glip/companies/{companyId} | Get Company Info
[**ReadGlipPerson**](GlipProfileApi.md#ReadGlipPerson) | **Get** /restapi/v1.0/glip/persons/{personId} | Get Person
[**ReadGlipPreferences**](GlipProfileApi.md#ReadGlipPreferences) | **Get** /restapi/v1.0/glip/preferences | Get Preferences



## ReadGlipCompany

> GlipCompany ReadGlipCompany(ctx, companyId)
Get Company Info

Returns information about a company or multiple companies by ID(s)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string**| Internal identifier of an RC account/Glip company, or tilde (~) to indicate a company the current user belongs to. | 

### Return type

[**GlipCompany**](GlipCompany.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/mixed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipPerson

> GlipPersonInfo ReadGlipPerson(ctx, personId)
Get Person

Returns a user or multiple users by their ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string**| Internal identifier of a user to be returned, the maximum number of IDs is 30 | 

### Return type

[**GlipPersonInfo**](GlipPersonInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipPreferences

> GlipPreferencesInfo ReadGlipPreferences(ctx, )
Get Preferences

Returns information about user preferences.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GlipPreferencesInfo**](GlipPreferencesInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

