# \CallLogApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserCallLog**](CallLogApi.md#DeleteUserCallLog) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log | Delete User Call Log
[**ListCompanyActiveCalls**](CallLogApi.md#ListCompanyActiveCalls) | **Get** /restapi/v1.0/account/{accountId}/active-calls | Get Company Active Calls
[**ListExtensionActiveCalls**](CallLogApi.md#ListExtensionActiveCalls) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/active-calls | Get User Active Calls
[**ReadCompanyCallLog**](CallLogApi.md#ReadCompanyCallLog) | **Get** /restapi/v1.0/account/{accountId}/call-log | Get Company Call Log Records
[**ReadCompanyCallRecord**](CallLogApi.md#ReadCompanyCallRecord) | **Get** /restapi/v1.0/account/{accountId}/call-log/{callRecordId} | Get Company Call Log Record(s)
[**ReadUserCallLog**](CallLogApi.md#ReadUserCallLog) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log | Get User Call Log Records
[**ReadUserCallRecord**](CallLogApi.md#ReadUserCallRecord) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log/{callRecordId} | Get User Call Record
[**SyncUserCallLog**](CallLogApi.md#SyncUserCallLog) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/call-log-sync | Sync User Call Log



## DeleteUserCallLog

> DeleteUserCallLog(ctx, accountId, extensionId, optional)
Delete User Call Log

Deletes filtered call log records.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***DeleteUserCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteUserCallLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dateTo** | **optional.Time**| The end datetime for records deletion in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **phoneNumber** | **optional.String**|  | 
 **extensionNumber** | **optional.String**|  | 
 **type_** | [**optional.Interface of []string**](string.md)|  | 
 **direction** | [**optional.Interface of []string**](string.md)|  | 
 **dateFrom** | **optional.Time**|  | 

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


## ListCompanyActiveCalls

> ActiveCallsResponse ListCompanyActiveCalls(ctx, accountId, optional)
Get Company Active Calls

Returns records of all calls that are in progress, ordered by start time in descending order.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***ListCompanyActiveCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCompanyActiveCallsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | [**optional.Interface of []string**](string.md)| The direction for the result records. If not specified, both inbound and outbound records are returned. Multiple values are accepted | 
 **view** | **optional.String**| View of call records. The same view parameter specified for FSync will be applied for ISync, the view cannot be changed for ISync | [default to Simple]
 **type_** | [**optional.Interface of []string**](string.md)| Call type of a record. If not specified, all call types are returned. Multiple values are accepted | 
 **transport** | [**optional.Interface of []string**](string.md)| Call transport type. &#39;PSTN&#39; specifies that a call leg is initiated from the PSTN network provider; &#39;VoIP&#39; - from an RC phone. By default this filter is disabled | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]

### Return type

[**ActiveCallsResponse**](ActiveCallsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensionActiveCalls

> ActiveCallsResponse ListExtensionActiveCalls(ctx, accountId, extensionId, optional)
Get User Active Calls

Returns records of all extension calls that are in progress, ordered by start time in descending order.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionActiveCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExtensionActiveCallsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **direction** | [**optional.Interface of []string**](string.md)| The direction for the result records. If not specified, both inbound and outbound records are returned. Multiple values are accepted | 
 **view** | **optional.String**| View of call records. The same view parameter specified for FSync will be applied for ISync, the view cannot be changed for ISync | [default to Simple]
 **type_** | [**optional.Interface of []string**](string.md)| Call type of a record. If not specified, all call types are returned. Multiple values are accepted | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]

### Return type

[**ActiveCallsResponse**](ActiveCallsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCompanyCallLog

> AccountCallLogResponse ReadCompanyCallLog(ctx, accountId, optional)
Get Company Call Log Records

Returns call log records filtered by parameters specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***ReadCompanyCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadCompanyCallLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionNumber** | **optional.String**| Extension number of a user. If specified, returns call log for a particular extension only | 
 **phoneNumber** | **optional.String**| Phone number of a caller/call recipient. If specified, returns all calls (both incoming and outcoming) with the phone number specified. Cannot be specified together with the extensionNumber filter | 
 **direction** | [**optional.Interface of []string**](string.md)| The direction for the result records. If not specified, both inbound and outbound records are returned. Multiple values are accepted | 
 **type_** | [**optional.Interface of []string**](string.md)| Call type of a record. If not specified, all call types are returned. Multiple values are accepted | 
 **view** | **optional.String**| View of call records. The same view parameter specified for FSync will be applied for ISync, the view cannot be changed for ISync | [default to Simple]
 **withRecording** | **optional.Bool**| **Deprecated**. Supported for compatibility reasons only. &#x60;true&#x60; if only recorded calls are returned. The default value is &#x60;false&#x60;. If both &#x60;withRecording&#x60; and &#x60;recordingType&#x60; are specified, &#x60;withRecording&#x60; is ignored | 
 **recordingType** | **optional.String**| Type of a call recording. If not specified, then calls without recordings are also returned | 
 **dateFrom** | **optional.Time**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | 
 **dateTo** | **optional.Time**| The end datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **sessionId** | **optional.String**| Internal identifier of a session | 

### Return type

[**AccountCallLogResponse**](AccountCallLogResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCompanyCallRecord

> CompanyCallLogRecord ReadCompanyCallRecord(ctx, accountId, callRecordId)
Get Company Call Log Record(s)

Returns individual call log record(s) by ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**callRecordId** | **string**| Internal identifier of a call log record | 

### Return type

[**CompanyCallLogRecord**](CompanyCallLogRecord.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserCallLog

> UserCallLogResponse ReadUserCallLog(ctx, accountId, extensionId, optional)
Get User Call Log Records

Returns call log records filtered by parameters specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ReadUserCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadUserCallLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extensionNumber** | **optional.String**| Extension number of a user. If specified, returns call log for a particular extension only | 
 **showBlocked** | **optional.Bool**| If &#39;True&#39; then calls from/to blocked numbers are returned | [default to true]
 **phoneNumber** | **optional.String**| Phone number of a caller/callee. If specified, returns all calls (both incoming and outcoming) with the phone number specified | 
 **direction** | [**optional.Interface of []string**](string.md)| The direction for the resulting records. If not specified, both inbound and outbound records are returned. Multiple values are accepted | 
 **sessionId** | **optional.String**| Internal identifier of a session | 
 **type_** | [**optional.Interface of []string**](string.md)| Call type of a record. It is allowed to specify more than one type. If not specified, all call types are returned. Multiple values are accepted | 
 **transport** | [**optional.Interface of []string**](string.md)| Call transport type. &#39;PSTN&#39; specifies that a call leg is initiated from the PSTN network provider; &#39;VoIP&#39; - from an RC phone. By default this filter is disabled | 
 **view** | **optional.String**| View of call records. The same view parameter specified for FSync will be applied for ISync, the view cannot be changed for ISync | [default to Simple]
 **withRecording** | **optional.Bool**| **Deprecated**. Supported for compatibility reasons. &#x60;True&#x60; if only recorded calls are returned. If both &#x60;withRecording&#x60; and &#x60;recordingType&#x60; are specified, then &#x60;withRecording&#x60; is ignored | [default to false]
 **recordingType** | **optional.String**| Type of a call recording. If not specified, then calls without recordings are also returned | 
 **dateTo** | **optional.Time**| The end datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is current time | 
 **dateFrom** | **optional.Time**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is dateTo minus 24 hours | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are allowed | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **showDeleted** | **optional.Bool**| If &#39;True&#39; then deleted calls are returned | [default to false]

### Return type

[**UserCallLogResponse**](UserCallLogResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserCallRecord

> UserCallLogRecord ReadUserCallRecord(ctx, callRecordId, extensionId, accountId, optional)
Get User Call Record

Returns call log records by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callRecordId** | [**[]string**](string.md)|  | 
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **optional** | ***ReadUserCallRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadUserCallRecordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **view** | **optional.String**| View of call records. The same view parameter specified for FSync will be applied for ISync, the view cannot be changed for ISync | [default to Simple]

### Return type

[**UserCallLogRecord**](UserCallLogRecord.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncUserCallLog

> CallLogSync SyncUserCallLog(ctx, accountId, extensionId, optional)
Sync User Call Log

Synchronizes call log records

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***SyncUserCallLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncUserCallLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syncType** | [**optional.Interface of []string**](string.md)| Type of synchronization | 
 **syncToken** | **optional.String**| Value of syncToken property of last sync request response | 
 **dateFrom** | **optional.Time**| The start datetime for resulting records in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. The default value is the current moment | 
 **recordCount** | **optional.Int32**| For &#39;FSync&#39; the parameter is mandatory, it limits the number of records to be returned in response. For &#39;ISync&#39; it specifies with how many records to extend sync Frame to the past, the maximum number of records is 250 | 
 **statusGroup** | [**optional.Interface of []string**](string.md)| Type of calls to be returned. The default value is &#39;All&#39; | 
 **view** | **optional.String**| View of call records. The same view parameter specified for FSync will be applied for ISync, the view cannot be changed for ISync | [default to Simple]
 **showDeleted** | **optional.Bool**| Supproted for ISync. If &#39;True&#39; then deleted call records are returned | [default to false]

### Return type

[**CallLogSync**](CallLogSync.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

