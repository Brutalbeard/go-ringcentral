# \PagingOnlyGroupsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignMultiplePagingGroupUsersDevices**](PagingOnlyGroupsApi.md#AssignMultiplePagingGroupUsersDevices) | **Post** /restapi/v1.0/account/{accountId}/paging-only-groups/{pagingOnlyGroupId}/bulk-assign | Assign Paging Group Users and Devices
[**ListPagingGroupDevices**](PagingOnlyGroupsApi.md#ListPagingGroupDevices) | **Get** /restapi/v1.0/account/{accountId}/paging-only-groups/{pagingOnlyGroupId}/devices | Get Paging Group Devices
[**ListPagingGroupUsers**](PagingOnlyGroupsApi.md#ListPagingGroupUsers) | **Get** /restapi/v1.0/account/{accountId}/paging-only-groups/{pagingOnlyGroupId}/users | Get Paging Group Users



## AssignMultiplePagingGroupUsersDevices

> AssignMultiplePagingGroupUsersDevices(ctx, accountId, pagingOnlyGroupId, optional)
Assign Paging Group Users and Devices

Adds and/or removes paging group users and devices.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**pagingOnlyGroupId** | **string**|  | 
 **optional** | ***AssignMultiplePagingGroupUsersDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssignMultiplePagingGroupUsersDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EditPagingGroupRequest**](EditPagingGroupRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPagingGroupDevices

> PagingOnlyGroupDevices ListPagingGroupDevices(ctx, accountId, pagingOnlyGroupId, optional)
Get Paging Group Devices

Returns the list of paging devices assigned to this group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**pagingOnlyGroupId** | **string**|  | 
 **optional** | ***ListPagingGroupDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPagingGroupDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]

### Return type

[**PagingOnlyGroupDevices**](PagingOnlyGroupDevices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPagingGroupUsers

> PagingOnlyGroupUsers ListPagingGroupUsers(ctx, accountId, pagingOnlyGroupId, optional)
Get Paging Group Users

Returns the list of users allowed to page this group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**pagingOnlyGroupId** | **string**|  | 
 **optional** | ***ListPagingGroupUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPagingGroupUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]

### Return type

[**PagingOnlyGroupUsers**](PagingOnlyGroupUsers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

