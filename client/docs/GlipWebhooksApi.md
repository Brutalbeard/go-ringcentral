# \GlipWebhooksApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateGlipWebhook**](GlipWebhooksApi.md#ActivateGlipWebhook) | **Post** /restapi/v1.0/glip/webhooks/{webhookId}/activate | Activate Webhook
[**CreateGlipGroupWebhook**](GlipWebhooksApi.md#CreateGlipGroupWebhook) | **Post** /restapi/v1.0/glip/groups/{groupId}/webhooks | Create Webhook in Group
[**DeleteGlipWebhook**](GlipWebhooksApi.md#DeleteGlipWebhook) | **Delete** /restapi/v1.0/glip/webhooks/{webhookId} | Delete Webhook
[**ListGlipGroupWebhooks**](GlipWebhooksApi.md#ListGlipGroupWebhooks) | **Get** /restapi/v1.0/glip/groups/{groupId}/webhooks | Get Webhooks in Group
[**ListGlipWebhooks**](GlipWebhooksApi.md#ListGlipWebhooks) | **Get** /restapi/v1.0/glip/webhooks | Get Webhooks
[**ReadGlipWebhook**](GlipWebhooksApi.md#ReadGlipWebhook) | **Get** /restapi/v1.0/glip/webhooks/{webhookId} | Get Webhook
[**SuspendGlipWebhook**](GlipWebhooksApi.md#SuspendGlipWebhook) | **Post** /restapi/v1.0/glip/webhooks/{webhookId}/suspend | Suspend Webhook



## ActivateGlipWebhook

> ActivateGlipWebhook(ctx, webhookId)
Activate Webhook

Activates webhook by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string**| Internal identifier of a webhook | 

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


## CreateGlipGroupWebhook

> GlipWebhookInfo CreateGlipGroupWebhook(ctx, groupId)
Create Webhook in Group

Creates a new webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| Internal identifier of a group | 

### Return type

[**GlipWebhookInfo**](GlipWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlipWebhook

> DeleteGlipWebhook(ctx, webhookId)
Delete Webhook

Deletes the webhook by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string**| Internal identifier of a webhook | 

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


## ListGlipGroupWebhooks

> GlipWebhookList ListGlipGroupWebhooks(ctx, groupId)
Get Webhooks in Group

Returns webhooks which are available for the current user by group ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| Internal identifier of a group | 

### Return type

[**GlipWebhookList**](GlipWebhookList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlipWebhooks

> GlipWebhookList ListGlipWebhooks(ctx, )
Get Webhooks

Returns all webhooks.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GlipWebhookList**](GlipWebhookList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/mixed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipWebhook

> GlipWebhookList ReadGlipWebhook(ctx, webhookId)
Get Webhook

Returns webhooks(s) with the specified id(s).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**[]string**](string.md)| Internal identifier of a webhook or comma separated list of webhooks IDs | 

### Return type

[**GlipWebhookList**](GlipWebhookList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/mixed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendGlipWebhook

> SuspendGlipWebhook(ctx, webhookId)
Suspend Webhook

Suspends webhook by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string**| Internal identifier of a webhook | 

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

