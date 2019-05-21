# \APIInfoApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAPIStatus**](APIInfoApi.md#ReadAPIStatus) | **Get** /restapi/v1.0/status | Get Service Status
[**ReadAPIVersion**](APIInfoApi.md#ReadAPIVersion) | **Get** /restapi/{apiVersion} | Get Version Info
[**ReadAPIVersions**](APIInfoApi.md#ReadAPIVersions) | **Get** /restapi | Get API Versions



## ReadAPIStatus

> ReadAPIStatus(ctx, )
Get Service Status

Returns current PAS service status.

### Required Parameters

This endpoint does not need any parameter.

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


## ReadAPIVersion

> GetVersionResponse ReadAPIVersion(ctx, apiVersion)
Get Version Info

Returns current API version info by apiVersion.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string**| API version to be requested, for example &#39;v1.0&#39; | 

### Return type

[**GetVersionResponse**](GetVersionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAPIVersions

> GetVersionsResponse ReadAPIVersions(ctx, )
Get API Versions

Returns current API version(s) and server info.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetVersionsResponse**](GetVersionsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

