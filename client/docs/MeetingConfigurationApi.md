# \MeetingConfigurationApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAssistants**](MeetingConfigurationApi.md#ReadAssistants) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meetings-configuration/assistants | Get Assistants
[**ReadAssistedUsers**](MeetingConfigurationApi.md#ReadAssistedUsers) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meetings-configuration/assisted | Get Assisted Users
[**ReadMeetingServiceInfo**](MeetingConfigurationApi.md#ReadMeetingServiceInfo) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/service-info | Get Meeting Service Info



## ReadAssistants

> AssistantsResource ReadAssistants(ctx, accountId, extensionId)
Get Assistants

Returns assistants information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**AssistantsResource**](AssistantsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAssistedUsers

> AssistedUsersResource ReadAssistedUsers(ctx, accountId, extensionId)
Get Assisted Users

Returns assisted users information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 

### Return type

[**AssistedUsersResource**](AssistedUsersResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMeetingServiceInfo

> MeetingServiceInfoResource ReadMeetingServiceInfo(ctx, extensionId, accountId)
Get Meeting Service Info

Returns information on dial-in numbers for meetings, support and international dial-in numbers URIs and meeting account information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**MeetingServiceInfoResource**](MeetingServiceInfoResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

