# \UserSettingsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserProfileImage**](UserSettingsApi.md#CreateUserProfileImage) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | Upload User Profile Image
[**DeleteExtension**](UserSettingsApi.md#DeleteExtension) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId} | Delete Extension
[**ListExtensionGrants**](UserSettingsApi.md#ListExtensionGrants) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/grant | Get Extension Grant List
[**ReadConferencingSettings**](UserSettingsApi.md#ReadConferencingSettings) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/conferencing | Get User Conferencing Settings
[**ReadExtension**](UserSettingsApi.md#ReadExtension) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId} | Get Extension
[**ReadExtensionCallerId**](UserSettingsApi.md#ReadExtensionCallerId) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-id | Get Extension Caller ID
[**ReadNotificationSettings**](UserSettingsApi.md#ReadNotificationSettings) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/notification-settings | Get Notification Settings
[**ReadScaledPofileImage**](UserSettingsApi.md#ReadScaledPofileImage) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image/{scaleSize} | Get Scaled User Profile Image
[**ReadUserProfileImage**](UserSettingsApi.md#ReadUserProfileImage) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | Get User Profile Image
[**UpdateConferencingSettings**](UserSettingsApi.md#UpdateConferencingSettings) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/conferencing | Update User Conferencing Settings
[**UpdateExtension**](UserSettingsApi.md#UpdateExtension) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId} | Update Extension
[**UpdateExtensionCallerId**](UserSettingsApi.md#UpdateExtensionCallerId) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/caller-id | Update Extension Caller ID
[**UpdateNotificationSettings**](UserSettingsApi.md#UpdateNotificationSettings) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/notification-settings | Update Notification Settings
[**UpdateUserProfileImage**](UserSettingsApi.md#UpdateUserProfileImage) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/profile-image | Update User Profile Image



## CreateUserProfileImage

> CreateUserProfileImage(ctx, extensionId, accountId, image)
Upload User Profile Image

Uploads the extension profile image.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**|  | 
**accountId** | **string**|  | 
**image** | ***os.File*****os.File**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExtension

> DeleteExtension(ctx, extensionId, accountId)
Delete Extension

Deletes extension(s) by ID(s).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**|  | 
**accountId** | **string**|  | 

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


## ListExtensionGrants

> GetExtensionGrantListResponse ListExtensionGrants(ctx, accountId, extensionId, optional)
Get Extension Grant List

Returns the list of extension grants.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionGrantsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExtensionGrantsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.String**|  | [default to 1]
 **perPage** | **optional.String**|  | [default to 100]

### Return type

[**GetExtensionGrantListResponse**](GetExtensionGrantListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadConferencingSettings

> GetConferencingInfoResponse ReadConferencingSettings(ctx, accountId, extensionId, optional)
Get User Conferencing Settings

Returns the information on the Free Conference Calling (FCC) feature for a given extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ReadConferencingSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadConferencingSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **countryId** | **optional.String**| Internal identifier of a country. If not specified, the response is returned for the brand country | 

### Return type

[**GetConferencingInfoResponse**](GetConferencingInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadExtension

> GetExtensionInfoResponse ReadExtension(ctx, accountId, extensionId)
Get Extension

Returns basic information about a particular extension of an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**GetExtensionInfoResponse**](GetExtensionInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadExtensionCallerId

> ExtensionCallerIdInfo ReadExtensionCallerId(ctx, accountId, extensionId)
Get Extension Caller ID

Returns information on an outbound caller ID of an extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

### Return type

[**ExtensionCallerIdInfo**](ExtensionCallerIdInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNotificationSettings

> NotificationSettings ReadNotificationSettings(ctx, accountId, extensionId)
Get Notification Settings

Returns notification settings for the current extension.  <p>Knowledge Article: <a href=\"https://success.ringcentral.com/articles/RC_Knowledge_Article/9740\">User Settings - Set up Message Notifications</a></p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**NotificationSettings**](NotificationSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadScaledPofileImage

> *os.File ReadScaledPofileImage(ctx, accountId, extensionId, scaleSize)
Get Scaled User Profile Image

Returns scaled profile image of an extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**scaleSize** | **string**| Dimensions of a profile image which will be returned in response. If this path parameter is not specified in request URI then | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserProfileImage

> *os.File ReadUserProfileImage(ctx, extensionId, accountId)
Get User Profile Image

Returns a profile image of an extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**|  | 
**accountId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, image/png, image/jpeg, image/gif

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConferencingSettings

> GetConferencingInfoResponse UpdateConferencingSettings(ctx, accountId, extensionId, body)
Update User Conferencing Settings

Updates the default conferencing number for the current extension. The number can be selected from conferencing numbers of the current extension. Updates the setting, allowing participants join the conference before host.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**UpdateConferencingInfoRequest**](UpdateConferencingInfoRequest.md)| JSON body | 

### Return type

[**GetConferencingInfoResponse**](GetConferencingInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtension

> GetExtensionInfoResponse UpdateExtension(ctx, accountId, extensionId, body)
Update Extension

Updates user settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**ExtensionUpdateRequest**](ExtensionUpdateRequest.md)| JSON body | 

### Return type

[**GetExtensionInfoResponse**](GetExtensionInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExtensionCallerId

> ExtensionCallerIdInfo UpdateExtensionCallerId(ctx, accountId, extensionId, body)
Update Extension Caller ID

Updates outbound caller ID information of an extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**ExtensionCallerIdInfo**](ExtensionCallerIdInfo.md)| JSON body | 

### Return type

[**ExtensionCallerIdInfo**](ExtensionCallerIdInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationSettings

> NotificationSettings UpdateNotificationSettings(ctx, accountId, extensionId, body)
Update Notification Settings

Updates notification settings for the current extension. <p>Knowledge Article: <a href=\"https://success.ringcentral.com/articles/RC_Knowledge_Article/9740\">User Settings - Set up Message Notifications</a></p>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **int32**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**body** | [**NotificationSettingsUpdateRequest**](NotificationSettingsUpdateRequest.md)|  | 

### Return type

[**NotificationSettings**](NotificationSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserProfileImage

> UpdateUserProfileImage(ctx, extensionId, accountId, optional)
Update User Profile Image

Updates the extension profile image

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**|  | 
**accountId** | **string**|  | 
 **optional** | ***UpdateUserProfileImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserProfileImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **image** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

