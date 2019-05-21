# \MeetingManagementApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMeeting**](MeetingManagementApi.md#CreateMeeting) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting | Create Meetings
[**DeleteMeeting**](MeetingManagementApi.md#DeleteMeeting) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | Delete Meeting
[**EndMeeting**](MeetingManagementApi.md#EndMeeting) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId}/end | End Meeting
[**ListMeetings**](MeetingManagementApi.md#ListMeetings) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting | Get Scheduled Meetings
[**ReadMeeting**](MeetingManagementApi.md#ReadMeeting) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | Get Meeting Info
[**UpdateMeeting**](MeetingManagementApi.md#UpdateMeeting) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/meeting/{meetingId} | Update Meeting



## CreateMeeting

> MeetingResponseResource CreateMeeting(ctx, extensionId, accountId, optional)
Create Meetings

Creates a new meeting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **optional** | ***CreateMeetingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMeetingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MeetingRequestResource**](MeetingRequestResource.md)|  | 

### Return type

[**MeetingResponseResource**](MeetingResponseResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMeeting

> DeleteMeeting(ctx, meetingId, extensionId, accountId)
Delete Meeting

Deletes a scheduled meeting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string**| Internal identifier of a RingCentral meeting | 
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

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


## EndMeeting

> EndMeeting(ctx, meetingId, extensionId, accountId)
End Meeting

Ends a meetings which is in progress.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string**| Internal identifier of a RingCentral meeting | 
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
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


## ListMeetings

> MeetingsResource ListMeetings(ctx, extensionId, accountId)
Get Scheduled Meetings

Returns a list of meetings for a particular extension. The list of meetings does not include meetings of 'Instant' type.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**MeetingsResource**](MeetingsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMeeting

> MeetingResponseResource ReadMeeting(ctx, meetingId, extensionId, accountId)
Get Meeting Info

Returns a particular meetings details by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string**| Internal identifier of a RingCentral meeting | 
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 

### Return type

[**MeetingResponseResource**](MeetingResponseResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeeting

> MeetingResponseResource UpdateMeeting(ctx, meetingId, extensionId, accountId, optional)
Update Meeting

Modifies a particular meeting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingId** | **string**| Internal identifier of a RingCentral meeting | 
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | 
 **optional** | ***UpdateMeetingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMeetingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of MeetingRequestResource**](MeetingRequestResource.md)|  | 

### Return type

[**MeetingResponseResource**](MeetingResponseResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

