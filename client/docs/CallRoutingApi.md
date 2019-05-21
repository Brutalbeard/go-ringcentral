# \CallRoutingApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIVRMenu**](CallRoutingApi.md#CreateIVRMenu) | **Post** /restapi/v1.0/account/{accountId}/ivr-menus | Create IVR Menu
[**CreateIVRPrompt**](CallRoutingApi.md#CreateIVRPrompt) | **Post** /restapi/v1.0/account/{accountId}/ivr-prompts | Create IVR Prompts
[**DeleteIVRPrompt**](CallRoutingApi.md#DeleteIVRPrompt) | **Delete** /restapi/v1.0/account/{accountId}/ivr-prompts/{promptId} | Delete IVR Prompt
[**ListIVRPrompts**](CallRoutingApi.md#ListIVRPrompts) | **Get** /restapi/v1.0/account/{accountId}/ivr-prompts | Get IVR Prompt List
[**ReadIVRMenu**](CallRoutingApi.md#ReadIVRMenu) | **Get** /restapi/v1.0/account/{accountId}/ivr-menus/{ivrMenuId} | Get IVR Menu
[**ReadIVRPrompt**](CallRoutingApi.md#ReadIVRPrompt) | **Get** /restapi/v1.0/account/{accountId}/ivr-prompts/{promptId} | Get IVR Prompt
[**ReadIVRPromptContent**](CallRoutingApi.md#ReadIVRPromptContent) | **Get** /restapi/v1.0/account/{accountId}/ivr-prompts/{promptId}/content | Get IVR Prompt Content
[**UpdateIVRMenu**](CallRoutingApi.md#UpdateIVRMenu) | **Put** /restapi/v1.0/account/{accountId}/ivr-menus/{ivrMenuId} | Update IVR Menu
[**UpdateIVRPrompt**](CallRoutingApi.md#UpdateIVRPrompt) | **Put** /restapi/v1.0/account/{accountId}/ivr-prompts/{promptId} | Update IVR Prompt



## CreateIVRMenu

> IvrMenuInfo CreateIVRMenu(ctx, accountId, body)
Create IVR Menu

Creates a company IVR menu.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**body** | [**IvrMenuInfo**](IvrMenuInfo.md)| JSON body | 

### Return type

[**IvrMenuInfo**](IVRMenuInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIVRPrompt

> PromptInfo CreateIVRPrompt(ctx, accountId, attachment, optional)
Create IVR Prompts

Creates an IVR prompt.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**attachment** | ***os.File*****os.File**| Audio file that will be used as a prompt. Attachment cannot be empty, only audio files are supported | 
 **optional** | ***CreateIVRPromptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIVRPromptOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**| Description of file contents. | 

### Return type

[**PromptInfo**](PromptInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIVRPrompt

> DeleteIVRPrompt(ctx, accountId, promptId)
Delete IVR Prompt

Deletes an IVR prompt by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**promptId** | **string**|  | 

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


## ListIVRPrompts

> IvrPrompts ListIVRPrompts(ctx, accountId)
Get IVR Prompt List

Returns the list of IVR prompts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**IvrPrompts**](IVRPrompts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIVRMenu

> IvrMenuInfo ReadIVRMenu(ctx, accountId, ivrMenuId)
Get IVR Menu

Returns a company IVR menu by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**ivrMenuId** | **string**|  | 

### Return type

[**IvrMenuInfo**](IVRMenuInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIVRPrompt

> PromptInfo ReadIVRPrompt(ctx, accountId, promptId)
Get IVR Prompt

Returns an IVR prompt by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**promptId** | **string**|  | 

### Return type

[**PromptInfo**](PromptInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIVRPromptContent

> ReadIVRPromptContent(ctx, accountId, promptId)
Get IVR Prompt Content

Returns media content of an IVR prompt by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**promptId** | **string**|  | 

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


## UpdateIVRMenu

> IvrMenuInfo UpdateIVRMenu(ctx, accountId, ivrMenuId, body)
Update IVR Menu

Updates a company IVR menu by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**ivrMenuId** | **string**|  | 
**body** | [**IvrMenuInfo**](IvrMenuInfo.md)| JSON body | 

### Return type

[**IvrMenuInfo**](IVRMenuInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIVRPrompt

> UpdateIVRPrompt(ctx, accountId, promptId)
Update IVR Prompt

Updates an IVR prompt by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**promptId** | **string**|  | 

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

