# \SCIMApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckHealth**](SCIMApi.md#CheckHealth) | **Get** /scim/health | check health
[**CheckHealth2**](SCIMApi.md#CheckHealth2) | **Get** /scim/v2/health | check health
[**CreateUser**](SCIMApi.md#CreateUser) | **Post** /scim/Users | create a user
[**CreateUser2**](SCIMApi.md#CreateUser2) | **Post** /scim/v2/Users | Create User
[**DeleteUser2**](SCIMApi.md#DeleteUser2) | **Delete** /scim/v2/Users/{id} | Delete User
[**PatchUser2**](SCIMApi.md#PatchUser2) | **Patch** /scim/v2/Users/{id} | Update/Patch User
[**ReadServiceProviderConfig**](SCIMApi.md#ReadServiceProviderConfig) | **Get** /scim/ServiceProviderConfig | Get Service Provider Config
[**ReadServiceProviderConfig2**](SCIMApi.md#ReadServiceProviderConfig2) | **Get** /scim/v2/ServiceProviderConfig | Get Service Provider Config
[**ReadUser2**](SCIMApi.md#ReadUser2) | **Get** /scim/v2/Users/{id} | Get User
[**ReplaceUser2**](SCIMApi.md#ReplaceUser2) | **Put** /scim/v2/Users/{id} | Update/Replace User
[**SearchViaGet**](SCIMApi.md#SearchViaGet) | **Get** /scim/Users | search or list users
[**SearchViaGet2**](SCIMApi.md#SearchViaGet2) | **Get** /scim/v2/Users | Search or List Users
[**SearchViaPost2**](SCIMApi.md#SearchViaPost2) | **Post** /scim/v2/Users/.search | Search or List Users



## CheckHealth

> CheckHealth(ctx, )
check health

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


## CheckHealth2

> CheckHealth2(ctx, )
check health

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


## CreateUser

> UserResponse CreateUser(ctx, optional)
create a user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of User**](User.md)| a new user without &#39;id&#39; | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser2

> UserResponse CreateUser2(ctx, optional)
Create User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateUser2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUser2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of User**](User.md)| a new user without &#39;id&#39; | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser2

> DeleteUser2(ctx, id)
Delete User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| user id | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUser2

> UserResponse PatchUser2(ctx, id, optional)
Update/Patch User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| user id | 
 **optional** | ***PatchUser2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PatchUser2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserPatch**](UserPatch.md)| patch operations list | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadServiceProviderConfig

> ServiceProviderConfig ReadServiceProviderConfig(ctx, )
Get Service Provider Config

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceProviderConfig**](ServiceProviderConfig.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadServiceProviderConfig2

> ServiceProviderConfig ReadServiceProviderConfig2(ctx, )
Get Service Provider Config

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ServiceProviderConfig**](ServiceProviderConfig.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUser2

> UserResponse ReadUser2(ctx, id)
Get User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| user id | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceUser2

> UserResponse ReplaceUser2(ctx, id, optional)
Update/Replace User

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| user id | 
 **optional** | ***ReplaceUser2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplaceUser2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of User**](User.md)| an existing user | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchViaGet

> UserSearchResponse SearchViaGet(ctx, optional)
search or list users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchViaGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchViaGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| only support &#39;userName&#39; or &#39;email&#39; filter expressions for now | 
 **startIndex** | **optional.Int32**| start index (1-based) | [default to 1]
 **count** | **optional.Int32**| page size | [default to 100]

### Return type

[**UserSearchResponse**](UserSearchResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchViaGet2

> UserSearchResponse SearchViaGet2(ctx, optional)
Search or List Users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchViaGet2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchViaGet2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| only support &#39;userName&#39; or &#39;email&#39; filter expressions for now | 
 **startIndex** | **optional.Int32**| start index (1-based) | [default to 1]
 **count** | **optional.Int32**| page size | [default to 100]

### Return type

[**UserSearchResponse**](UserSearchResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchViaPost2

> UserSearchResponse SearchViaPost2(ctx, optional)
Search or List Users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchViaPost2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchViaPost2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SearchRequest**](SearchRequest.md)| search parameters | 

### Return type

[**UserSearchResponse**](UserSearchResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

