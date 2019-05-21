# \GlipComplianceExportsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataExportTask**](GlipComplianceExportsApi.md#CreateDataExportTask) | **Post** /restapi/v1.0/glip/data-export | Create Data Export Task
[**ReadComplianceArchive**](GlipComplianceExportsApi.md#ReadComplianceArchive) | **Get** /restapi/v1.0/glip/data-export/{taskId}/archive/{archiveId} | Get Glip Data Archive
[**ReadDataExportTask**](GlipComplianceExportsApi.md#ReadDataExportTask) | **Get** /restapi/v1.0/glip/data-export/{taskId} | Get Data Export Task



## CreateDataExportTask

> DataExportTask CreateDataExportTask(ctx, optional)
Create Data Export Task

Creates a new task for Glip data export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDataExportTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDataExportTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreateDataExportTaskRequest**](CreateDataExportTaskRequest.md)| JSON body | 

### Return type

[**DataExportTask**](DataExportTask.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadComplianceArchive

> ReadComplianceArchive(ctx, taskId, archiveId)
Get Glip Data Archive

Returns data archive by ID in application/zip format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Internal identifier of a data export task | 
**archiveId** | **string**|  | 

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


## ReadDataExportTask

> DataExportTask ReadDataExportTask(ctx, taskId)
Get Data Export Task

Returns data export results.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string**| Internal identifier of a task to be retrieved | 

### Return type

[**DataExportTask**](DataExportTask.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

