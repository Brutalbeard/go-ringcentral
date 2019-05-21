# \UserPermissionsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUserPermission**](UserPermissionsApi.md#CheckUserPermission) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile/check | Check User Permission
[**ReadAuthorizationProfile**](UserPermissionsApi.md#ReadAuthorizationProfile) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/authz-profile | Get Authorization Profile



## CheckUserPermission

> AuthProfileCheckResource CheckUserPermission(ctx, extensionId, accountId, optional)
Check User Permission

Checks if a certain user permission is activated for a particular extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**|  | 
**accountId** | **string**|  | 
 **optional** | ***CheckUserPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CheckUserPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **permissionId** | **optional.String**|  | 
 **targetExtensionId** | **optional.String**|  | 

### Return type

[**AuthProfileCheckResource**](AuthProfileCheckResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAuthorizationProfile

> AuthProfileResource ReadAuthorizationProfile(ctx, extensionId, accountId)
Get Authorization Profile

Returns a list of user permissions granted at authorization procedure. Please note: Some permissions may be restricted by extension type.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**|  | 
**accountId** | **string**|  | 

### Return type

[**AuthProfileResource**](AuthProfileResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

