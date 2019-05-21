# \CallRecordingsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCallRecordingData**](CallRecordingsApi.md#ListCallRecordingData) | **Get** /restapi/v1.0/account/{accountId}/recording/{recordingId}/content | Get Call Recordings Data
[**ReadCallRecording**](CallRecordingsApi.md#ReadCallRecording) | **Get** /restapi/v1.0/account/{accountId}/recording/{recordingId} | Get Call Recording



## ListCallRecordingData

> *os.File ListCallRecordingData(ctx, accountId, recordingId)
Get Call Recordings Data

Returns media content of a call recording.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**recordingId** | **string**| Internal identifier of a recording (returned in Call Log) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/x-wav, audio/mpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCallRecording

> GetCallRecordingResponse ReadCallRecording(ctx, accountId, recordingId)
Get Call Recording

Returns call recordings by ID(s).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**recordingId** | **string**| Internal identifier of a recording (returned in Call Log) | 

### Return type

[**GetCallRecordingResponse**](GetCallRecordingResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

