# \InternalContactsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDirectoryEntries**](InternalContactsApi.md#ListDirectoryEntries) | **Get** /restapi/v1.0/account/{accountId}/directory/entries | Get Company Directory Entries
[**ReadAccountFederation**](InternalContactsApi.md#ReadAccountFederation) | **Get** /restapi/v1.0/account/{accountId}/directory/federation | Get Account Federation
[**ReadDirectoryEntry**](InternalContactsApi.md#ReadDirectoryEntry) | **Get** /restapi/v1.0/account/{accountId}/directory/entries/{entryId} | Get Corporate Directory Entry
[**SearchDirectoryEntries**](InternalContactsApi.md#SearchDirectoryEntries) | **Post** /restapi/v1.0/account/{accountId}/directory/entries/search | Search Company Directory Entries



## ListDirectoryEntries

> DirectoryResource ListDirectoryEntries(ctx, accountId, optional)
Get Company Directory Entries

Returns contact information on corporate users of federated accounts. Please note: 1. `User`, `DigitalUser`, `VirtualUser` and `FaxUser` types are returned as `User` type. 2. `ApplicationExtension` type is not returned. 3. Only extensions in `Enabled`, `Disabled` and `NotActivated` state are returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **optional** | ***ListDirectoryEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDirectoryEntriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showFederated** | **optional.Bool**| If &#39;True&#39; then contacts of all accounts in federation are returned. If &#39;False&#39; then only contacts of the current account are returned, and account section is eliminated in this case | [default to true]
 **type_** | **optional.String**| Type of an extension | 
 **page** | **optional.String**| Page number | [default to 1]
 **perPage** | **optional.Int32**| Records count to be returned per one page. The default value is 1000. Specific keyword values: &#x60;all&#x60; - all records are returned in one page; &#x60;max&#x60; - maximum count of records that can be returned in one page | [default to 1000]
 **siteId** | **optional.String**| Internal identifier of the business site to which extensions belong | 
 **ifNoneMatch** | **optional.String**| If-None-Match | 

### Return type

[**DirectoryResource**](DirectoryResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAccountFederation

> FederationResource ReadAccountFederation(ctx, accountId, optional)
Get Account Federation

Returns information on a federation and associated accounts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **optional** | ***ReadAccountFederationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadAccountFederationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rCExtensionId** | **optional.String**| RCExtensionId | 

### Return type

[**FederationResource**](FederationResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDirectoryEntry

> ContactResource ReadDirectoryEntry(ctx, entryId, accountId)
Get Corporate Directory Entry

Returns contact information on a particular corporate user of a federated account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string**| Internal identifier of extension to read information for | 
**accountId** | **string**| Internal identifier of owning account | 

### Return type

[**ContactResource**](ContactResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDirectoryEntries

> DirectoryResource SearchDirectoryEntries(ctx, accountId, optional)
Search Company Directory Entries

Returns contact information on corporate users of federated accounts according to the specified filtering and ordering.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **optional** | ***SearchDirectoryEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchDirectoryEntriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SearchDirectoryEntriesRequest**](SearchDirectoryEntriesRequest.md)|  | 

### Return type

[**DirectoryResource**](DirectoryResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

