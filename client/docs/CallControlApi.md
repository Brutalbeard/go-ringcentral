# \CallControlApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallFlipParty**](CallControlApi.md#CallFlipParty) | **Post** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId}/flip | Call Flip on Party
[**CreateCallOutCallSession**](CallControlApi.md#CreateCallOutCallSession) | **Post** /restapi/v1.0/account/{accountId}/telephony/call-out | Create CallOut Call Session
[**DeleteCallSession**](CallControlApi.md#DeleteCallSession) | **Delete** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId} | Drop Call Session
[**ForwardCallParty**](CallControlApi.md#ForwardCallParty) | **Post** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId}/forward | Forward Call Party
[**HoldCallParty**](CallControlApi.md#HoldCallParty) | **Post** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId}/hold | Hold Call Party
[**PauseResumeCallRecording**](CallControlApi.md#PauseResumeCallRecording) | **Patch** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId}/recordings/{recordingId} | Pause/Resume Recording
[**ReadCallPartyStatus**](CallControlApi.md#ReadCallPartyStatus) | **Get** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId} | Get Call Party Status
[**ReadCallSessionStatus**](CallControlApi.md#ReadCallSessionStatus) | **Get** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId} | Get Call Session Status
[**RejectParty**](CallControlApi.md#RejectParty) | **Post** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId}/reject | Reject Call Party
[**StartCallRecording**](CallControlApi.md#StartCallRecording) | **Post** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId}/recordings | Create Recording
[**SuperviseCallSession**](CallControlApi.md#SuperviseCallSession) | **Post** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/supervise | Supervise Call
[**TransferCallParty**](CallControlApi.md#TransferCallParty) | **Post** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId}/transfer | Transfer Call Party
[**UnholdCallParty**](CallControlApi.md#UnholdCallParty) | **Post** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId}/unhold | Unhold Call Party
[**UpdateCallParty**](CallControlApi.md#UpdateCallParty) | **Patch** /restapi/v1.0/account/{accountId}/telephony/sessions/{sessionId}/parties/{partyId} | Update Call Party



## CallFlipParty

> CallFlipParty(ctx, accountId, sessionId, partyId, body)
Call Flip on Party

Performs call flip procedure by holding opposite party and calling to the specified target

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 
**body** | [**CallPartyFlip**](CallPartyFlip.md)| JSON body | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCallOutCallSession

> CallSession CreateCallOutCallSession(ctx, accountId, body)
Create CallOut Call Session

Creates a CallOut Call session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**body** | [**MakeCallOutRequest**](MakeCallOutRequest.md)| JSON body | 

### Return type

[**CallSession**](CallSession.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallSession

> DeleteCallSession(ctx, accountId, sessionId)
Drop Call Session

Drops a call session.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 

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


## ForwardCallParty

> CallParty ForwardCallParty(ctx, accountId, sessionId, partyId, body)
Forward Call Party

Distributes a non-answered call to the defined target. Applicable for \"Setup\" or \"Proceeding\" states

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 
**body** | [**ForwardTarget**](ForwardTarget.md)| Distributes a non-answered call to the defined target. Only a single target is allowed | 

### Return type

[**CallParty**](CallParty.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HoldCallParty

> CallParty HoldCallParty(ctx, accountId, sessionId, partyId)
Hold Call Party

Puts the party to stand-alone mode and starts to play Hold Music according to configuration & state to peers. There is a known limitation for Hold API - hold via REST API doesn't work with hold placed via RingCentral apps or HardPhone. It means that if you muted participant via Call Control API and RingCentral Desktop app, then you need to unhold both endpoints to remove Hold Music and bring media back.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 

### Return type

[**CallParty**](CallParty.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseResumeCallRecording

> CallRecording PauseResumeCallRecording(ctx, brandId, accountId, sessionId, partyId, recordingId, body)
Pause/Resume Recording

Pause/resume recording

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string**| Identifies a brand of a logged in user or a brand of a sign-up session | [default to ~]
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 
**recordingId** | **string**| Internal identifier of a recording | 
**body** | [**CallRecordingUpdate**](CallRecordingUpdate.md)| JSON body | 

### Return type

[**CallRecording**](CallRecording.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCallPartyStatus

> CallParty ReadCallPartyStatus(ctx, accountId, sessionId, partyId)
Get Call Party Status

Returns a party status of a call session by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 

### Return type

[**CallParty**](CallParty.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCallSessionStatus

> CallSessionObject ReadCallSessionStatus(ctx, accountId, sessionId, optional)
Get Call Session Status

Returns the status of a call session by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
 **optional** | ***ReadCallSessionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadCallSessionStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timestamp** | **optional.String**| The date and time of a call session latest change | 
 **timeout** | **optional.String**| The time frame of awaiting for a status change before sending the resulting one in response | 

### Return type

[**CallSessionObject**](CallSessionObject.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectParty

> RejectParty(ctx, accountId, sessionId, partyId)
Reject Call Party

Rejects an inbound call in a \"Setup\" or \"Proceeding\" state

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 

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


## StartCallRecording

> StartCallRecording(ctx, accountId, sessionId, partyId)
Create Recording

Starts a new call recording for the party

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 

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


## SuperviseCallSession

> SuperviseCallSession SuperviseCallSession(ctx, accountId, sessionId, body)
Supervise Call

Allows to monitor a call in 'Listen' mode. Allows to monitor a call in `Listen` mode. Input parameters should contain extension number of a monitored user and internal identifier of a supervisor's device. Call session should be specified in path.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**body** | [**SuperviseCallSessionRequest**](SuperviseCallSessionRequest.md)| JSON body | 

### Return type

[**SuperviseCallSession**](SuperviseCallSession.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferCallParty

> CallParty TransferCallParty(ctx, accountId, sessionId, partyId, body)
Transfer Call Party

Transfers a party by placing a new call to the specified target

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 
**body** | [**TransferTarget**](TransferTarget.md)| Defines a target for a party transfer. Only a single target is allowed | 

### Return type

[**CallParty**](CallParty.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnholdCallParty

> CallParty UnholdCallParty(ctx, accountId, sessionId, partyId)
Unhold Call Party

Brings a party back into a call and stops to play Hold Music. There is a known limitation for Hold API - hold via REST API doesn't work with hold placed via RingCentral apps or HardPhone. It means that if you muted participant via Call Control API and RingCentral Desktop app, then you need to unhold both endpoints to remove Hold Music and bring media back.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 

### Return type

[**CallParty**](CallParty.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallParty

> CallParty UpdateCallParty(ctx, accountId, sessionId, partyId, body)
Update Call Party

Modify the party of a call session by ID. There is a known limitation for Mute scenario - mute via REST API doesn't work with mute placed via RingCentral apps or HardPhone. It means that if you muted participant via Call Control API and Ringcentral Desktop app you need to unmute both endpoints to bring media back. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**sessionId** | **string**| Internal identifier of a call session | 
**partyId** | **string**| Internal identifier of a call party | 
**body** | [**PartyUpdateRequest**](PartyUpdateRequest.md)| JSON body | 

### Return type

[**CallParty**](CallParty.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

