# \CallQueuesApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignMultipleCallQueueMembers**](CallQueuesApi.md#AssignMultipleCallQueueMembers) | **Post** /restapi/v1.0/account/{accountId}/call-queues/{groupId}/bulk-assign | Assign Multiple Call Queue Members
[**AssignMultipleDepartmentMembers**](CallQueuesApi.md#AssignMultipleDepartmentMembers) | **Post** /restapi/v1.0/account/{accountId}/department/bulk-assign | Assign Multiple Department Members
[**ListCallQueueMembers**](CallQueuesApi.md#ListCallQueueMembers) | **Get** /restapi/v1.0/account/{accountId}/call-queues/{groupId}/members | Get Call Queue Members
[**ListCallQueues**](CallQueuesApi.md#ListCallQueues) | **Get** /restapi/v1.0/account/{accountId}/call-queues | Get Call Queues
[**ListDepartmentMembers**](CallQueuesApi.md#ListDepartmentMembers) | **Get** /restapi/v1.0/account/{accountId}/department/{departmentId}/members | Get Department Member List



## AssignMultipleCallQueueMembers

> AssignMultipleCallQueueMembers(ctx, accountId, groupId, body)
Assign Multiple Call Queue Members

Updates a call queue group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**groupId** | **string**|  | 
**body** | [**CallQueueBulkAssignResource**](CallQueueBulkAssignResource.md)| Changes for the given group | 

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


## AssignMultipleDepartmentMembers

> AssignMultipleDepartmentMembers(ctx, accountId, optional)
Assign Multiple Department Members

Adds and/or removes multiple call queue members

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | [default to ~]
 **optional** | ***AssignMultipleDepartmentMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssignMultipleDepartmentMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DepartmentBulkAssignResource**](DepartmentBulkAssignResource.md)|  | 

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


## ListCallQueueMembers

> CallQueueMembers ListCallQueueMembers(ctx, accountId, groupId, optional)
Get Call Queue Members

Returns call queue group members.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**groupId** | **string**|  | 
 **optional** | ***ListCallQueueMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallQueueMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]

### Return type

[**CallQueueMembers**](CallQueueMembers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallQueues

> CallQueues ListCallQueues(ctx, accountId, optional)
Get Call Queues

Returns call queue group list

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***ListCallQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallQueuesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **memberExtensionId** | **optional.String**| Internal identifier of an extension that is a member of every group within the result | 

### Return type

[**CallQueues**](CallQueues.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDepartmentMembers

> DepartmentMemberList ListDepartmentMembers(ctx, accountId, departmentId, optional)
Get Department Member List

Viewing user account info (including name, business name, address and phone number/account number)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**departmentId** | **int32**| Internal identifier of a Department extension (same as extensionId but only the ID of a department extension is valid) | 
 **optional** | ***ListDepartmentMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDepartmentMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]

### Return type

[**DepartmentMemberList**](DepartmentMemberList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

