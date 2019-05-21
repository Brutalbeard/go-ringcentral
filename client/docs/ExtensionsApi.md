# \ExtensionsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtension**](ExtensionsApi.md#CreateExtension) | **Post** /restapi/v1.0/account/{accountId}/extension | Create Extension
[**ListExtensions**](ExtensionsApi.md#ListExtensions) | **Get** /restapi/v1.0/account/{accountId}/extension | Get Extension List
[**ListUserTemplates**](ExtensionsApi.md#ListUserTemplates) | **Get** /restapi/v1.0/account/{accountId}/templates | Get User Template List
[**ReadUserTemplate**](ExtensionsApi.md#ReadUserTemplate) | **Get** /restapi/v1.0/account/{accountId}/templates/{templateId} | Get User Template



## CreateExtension

> ExtensionCreationResponse CreateExtension(ctx, accountId, body)
Create Extension

Creates an extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**body** | [**ExtensionCreationRequest**](ExtensionCreationRequest.md)| JSON body | 

### Return type

[**ExtensionCreationResponse**](ExtensionCreationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensions

> GetExtensionListResponse ListExtensions(ctx, accountId, optional)
Get Extension List

Returns the list of extensions created for a particular account. All types of extensions are included in this list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExtensionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionId** | **optional.String**| Extension number to retrieve | 
 **email** | **optional.String**| Extension email address | 
 **page** | **optional.Int64**| Indicates the page number to retrieve. Only positive number values are allowed | [default to 1]
 **perPage** | **optional.Int64**| Indicates the page size (number of items) | [default to 100]
 **status** | [**optional.Interface of []string**](string.md)| Extension current state. Multiple values are supported. If &#39;Unassigned&#39; is specified, then extensions without extensionNumber are returned. If not specified, then all extensions are returned. | 
 **type_** | [**optional.Interface of []string**](string.md)| Extension type. Multiple values are supported | 

### Return type

[**GetExtensionListResponse**](GetExtensionListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserTemplates

> UserTemplates ListUserTemplates(ctx, accountId, optional)
Get User Template List

Returns the list of user templates for the current account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***ListUserTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | 
 **page** | **optional.String**|  | 
 **perPage** | **optional.String**|  | 

### Return type

[**UserTemplates**](UserTemplates.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserTemplate

> TemplateInfo ReadUserTemplate(ctx, accountId, templateId)
Get User Template

Returns the user template by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**templateId** | **string**|  | 

### Return type

[**TemplateInfo**](TemplateInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

