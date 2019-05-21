# \CalendarEventsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvent**](CalendarEventsApi.md#CreateEvent) | **Post** /restapi/v1.0/glip/events | Create Event
[**CreateEventbyGroupId**](CalendarEventsApi.md#CreateEventbyGroupId) | **Post** /restapi/v1.0/glip/groups/{groupId}/events | Create Event by Group ID
[**DeleteEvent**](CalendarEventsApi.md#DeleteEvent) | **Delete** /restapi/v1.0/glip/events/{eventId} | Delete Event
[**ListGroupEvents**](CalendarEventsApi.md#ListGroupEvents) | **Get** /restapi/v1.0/glip/groups/{groupId}/events | Get Group Events
[**ReadEvent**](CalendarEventsApi.md#ReadEvent) | **Get** /restapi/v1.0/glip/events/{eventId} | Get Event
[**ReadGlipEvents**](CalendarEventsApi.md#ReadGlipEvents) | **Get** /restapi/v1.0/glip/events | Get User Events List
[**UpdateEvent**](CalendarEventsApi.md#UpdateEvent) | **Put** /restapi/v1.0/glip/events/{eventId} | Update Event



## CreateEvent

> GlipEventInfo CreateEvent(ctx, body)
Create Event

Creates a new event.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GlipEventCreate**](GlipEventCreate.md)|  | 

### Return type

[**GlipEventInfo**](GlipEventInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEventbyGroupId

> GlipEventInfo CreateEventbyGroupId(ctx, groupId, body)
Create Event by Group ID

Creates a new event by the specified group ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| Internal identifier of a group | 
**body** | [**GlipEventCreate**](GlipEventCreate.md)|  | 

### Return type

[**GlipEventInfo**](GlipEventInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEvent

> DeleteEvent(ctx, eventId)
Delete Event

Deletes an event by the specified ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| Internal identifier of an event to be deleted | 

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


## ListGroupEvents

> GlipEventInfo ListGroupEvents(ctx, groupId)
Get Group Events

Returns the list of events available for the current user by the specified group ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| Internal identifier of a group | 

### Return type

[**GlipEventInfo**](GlipEventInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEvent

> GlipEventInfo ReadEvent(ctx, eventId)
Get Event

Returns event(s) with given id(s).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | [**[]string**](string.md)| event id or comma separated list of event ids. | 

### Return type

[**GlipEventInfo**](GlipEventInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/mixed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipEvents

> GlipEventsInfo ReadGlipEvents(ctx, optional)
Get User Events List

Returns all events created by the current user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReadGlipEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadGlipEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordCount** | **optional.Int32**| Number of groups to be fetched by one request. The maximum value is 250, by default - 30. | [default to 30]
 **pageToken** | **optional.String**| Token of a page to be returned | 

### Return type

[**GlipEventsInfo**](GlipEventsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/mixed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEvent

> GlipEventInfo UpdateEvent(ctx, eventId, body)
Update Event

Updates an event by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string**| Internal identifier of an event | 
**body** | [**GlipEventCreate**](GlipEventCreate.md)|  | 

### Return type

[**GlipEventInfo**](GlipEventInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

