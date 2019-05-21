# \PresenceApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAccountPresence**](PresenceApi.md#ReadAccountPresence) | **Get** /restapi/v1.0/account/{accountId}/presence | Get User Presence Status List
[**ReadUserPresenceStatus**](PresenceApi.md#ReadUserPresenceStatus) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence | Get User Presence Status
[**UpdateUserPresenceStatus**](PresenceApi.md#UpdateUserPresenceStatus) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/presence | Update User Presence Status



## ReadAccountPresence

> AccountPresenceInfo ReadAccountPresence(ctx, accountId, optional)
Get User Presence Status List

Returns presence status of all extensions of an account. Please note: The presenceStatus is returned as Offline (the parameters telephonyStatus, message, userStatus and dndStatus are not returned at all) for the following extension types: Department, Announcement Only, Voicemail (Take Messages Only), Fax User, Paging Only Group, Shared Lines Group, IVR Menu, Application Extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***ReadAccountPresenceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadAccountPresenceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detailedTelephonyState** | **optional.Bool**| Whether to return detailed telephony state | 
 **sipData** | **optional.Bool**| Whether to return SIP data | 
 **page** | **optional.Int32**| Page number for account presence information | 
 **perPage** | **optional.Int32**| Number for account presence information items per page | 

### Return type

[**AccountPresenceInfo**](AccountPresenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserPresenceStatus

> GetPresenceInfo ReadUserPresenceStatus(ctx, accountId, extensionId, optional)
Get User Presence Status

Returns presence status of an extension or several extensions by their ID(s). Batch request is supported. The 'presenceStatus' is returned as Offline (the parameters 'telephonyStatus', 'message', 'userStatus' and 'dndStatus' are not returned at all) for the following extension types: Department/Announcement Only/Take Messages Only (Voicemail)/Fax User/Paging Only Group/Shared Lines Group/IVR Menu/Application Extension/Park Location.If the user requests his/her own presence status, the response contains actual presence status even if the status publication is turned off. Batch request is supported. For batch requests the number of extensions in one request is limited to 30. If more extensions are included in the request, the error code 400 Bad Request is returned with the logical error code InvalidMultipartRequest and the corresponding message 'Extension Presence Info multipart request is limited to 30 extensions'.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ReadUserPresenceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadUserPresenceStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detailedTelephonyState** | **optional.Bool**| Whether to return detailed telephony state | 
 **sipData** | **optional.Bool**| Whether to return SIP data | 

### Return type

[**GetPresenceInfo**](GetPresenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPresenceStatus

> PresenceInfoResource UpdateUserPresenceStatus(ctx, extensionId, accountId, optional)
Update User Presence Status

Updates user-defined extension presence status, status message and DnD status by extension ID. Supported for regular User extensions only. The extension types listed do not support presence status: Department, Announcement Only, Take Messages Only (Voicemail), Fax User, Paging Only Group, Shared Lines Group, IVR Menu, Application Extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
 **optional** | ***UpdateUserPresenceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserPresenceStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PresenceInfoResource**](PresenceInfoResource.md)|  | 

### Return type

[**PresenceInfoResource**](PresenceInfoResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

