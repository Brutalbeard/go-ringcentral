# \MessageExportsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageStoreReport**](MessageExportsApi.md#CreateMessageStoreReport) | **Post** /restapi/v1.0/account/{accountId}/message-store-report | Create Message Store Report
[**ReadMessageStoreReportArchive**](MessageExportsApi.md#ReadMessageStoreReportArchive) | **Get** /restapi/v1.0/account/{accountId}/message-store-report/{taskId}/archive | Get Message Store Report Archive
[**ReadMessageStoreReportArchiveContent**](MessageExportsApi.md#ReadMessageStoreReportArchiveContent) | **Get** /restapi/v1.0/account/{accountId}/message-store-report/{taskId}/archive/{archiveId} | Get Message Store Report Archive Content
[**ReadMessageStoreReportTask**](MessageExportsApi.md#ReadMessageStoreReportTask) | **Get** /restapi/v1.0/account/{accountId}/message-store-report/{taskId} | Get Message Store Report Task



## CreateMessageStoreReport

> MessageStoreReport CreateMessageStoreReport(ctx, accountId, optional)
Create Message Store Report

Creates a task to collect all account messages for a specified time interval.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***CreateMessageStoreReportOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMessageStoreReportOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateMessageStoreReportRequest**](CreateMessageStoreReportRequest.md)|  | 

### Return type

[**MessageStoreReport**](MessageStoreReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMessageStoreReportArchive

> MessageStoreReportArchive ReadMessageStoreReportArchive(ctx, accountId, taskId)
Get Message Store Report Archive

Returns the created report with message data not including attachments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**taskId** | **string**|  | 

### Return type

[**MessageStoreReportArchive**](MessageStoreReportArchive.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMessageStoreReportArchiveContent

> *os.File ReadMessageStoreReportArchiveContent(ctx, accountId, taskId, archiveId)
Get Message Store Report Archive Content

Returns one of the report archives with message contents in application/zip format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**taskId** | **string**|  | 
**archiveId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMessageStoreReportTask

> MessageStoreReport ReadMessageStoreReportTask(ctx, accountId, taskId)
Get Message Store Report Task

Returns the current status of a task on report creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**taskId** | **string**|  | 

### Return type

[**MessageStoreReport**](MessageStoreReport.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

