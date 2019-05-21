# \CallMonitoringGroupsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallMonitoringGroup**](CallMonitoringGroupsApi.md#CreateCallMonitoringGroup) | **Post** /restapi/v1.0/account/{accountId}/call-monitoring-groups | Create Call Monitoring Group
[**DeleteCallMonitoringGroup**](CallMonitoringGroupsApi.md#DeleteCallMonitoringGroup) | **Delete** /restapi/v1.0/account/{accountId}/call-monitoring-groups/{groupId} | Delete Call Monitoring Group
[**ListCallMonitoringGroupMembers**](CallMonitoringGroupsApi.md#ListCallMonitoringGroupMembers) | **Get** /restapi/v1.0/account/{accountId}/call-monitoring-groups/{groupId}/members | Get Call Monitoring Group Member List
[**ListCallMonitoringGroups**](CallMonitoringGroupsApi.md#ListCallMonitoringGroups) | **Get** /restapi/v1.0/account/{accountId}/call-monitoring-groups | Get Call Monitoring Groups List
[**UpdateCallMonitoringGroup**](CallMonitoringGroupsApi.md#UpdateCallMonitoringGroup) | **Put** /restapi/v1.0/account/{accountId}/call-monitoring-groups/{groupId} | Updates Call Monitoring Group
[**UpdateCallMonitoringGroupList**](CallMonitoringGroupsApi.md#UpdateCallMonitoringGroupList) | **Post** /restapi/v1.0/account/{accountId}/call-monitoring-groups/{groupId}/bulk-assign | Update Call Monitoring Group List



## CreateCallMonitoringGroup

> CallMonitoringGroup CreateCallMonitoringGroup(ctx, accountId, body)
Create Call Monitoring Group

Creates a new call monitoring group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**body** | [**CreateCallMonitoringGroupRequest**](CreateCallMonitoringGroupRequest.md)| Parameters of a call monitoring group that will be created | 

### Return type

[**CallMonitoringGroup**](CallMonitoringGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallMonitoringGroup

> DeleteCallMonitoringGroup(ctx, accountId, groupId)
Delete Call Monitoring Group

Remove infromation about the given call monitoring group.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**groupId** | **string**|  | 

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


## ListCallMonitoringGroupMembers

> CallMonitoringGroupMemberList ListCallMonitoringGroupMembers(ctx, accountId, groupId, optional)
Get Call Monitoring Group Member List

Returns call monitoring group members.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**groupId** | **string**|  | 
 **optional** | ***ListCallMonitoringGroupMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallMonitoringGroupMembersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]

### Return type

[**CallMonitoringGroupMemberList**](CallMonitoringGroupMemberList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallMonitoringGroups

> CallMonitoringGroups ListCallMonitoringGroups(ctx, accountId, optional)
Get Call Monitoring Groups List

Returns call monitoring groups that can be filtered by some extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***ListCallMonitoringGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallMonitoringGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **memberExtensionId** | **optional.String**| Internal identifier of an extension that is a member of every group within the result | 

### Return type

[**CallMonitoringGroups**](CallMonitoringGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallMonitoringGroup

> CallMonitoringGroup UpdateCallMonitoringGroup(ctx, accountId, groupId, body)
Updates Call Monitoring Group

Updates call monitoring group name by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**groupId** | **string**| Internal identifier of a call monitoring group | 
**body** | [**CreateCallMonitoringGroupRequest**](CreateCallMonitoringGroupRequest.md)| Parameters of a call monitoring group that will be updated | 

### Return type

[**CallMonitoringGroup**](CallMonitoringGroup.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallMonitoringGroupList

> UpdateCallMonitoringGroupList(ctx, accountId, groupId, body)
Update Call Monitoring Group List

Updates call monitoring groups.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**groupId** | **string**|  | 
**body** | [**CallMonitoringBulkAssign**](CallMonitoringBulkAssign.md)| Changes for the given group | 

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

