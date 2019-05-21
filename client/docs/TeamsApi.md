# \TeamsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGlipTeamMembers**](TeamsApi.md#AddGlipTeamMembers) | **Post** /restapi/v1.0/glip/teams/{chatId}/add | Add Team Members
[**ArchiveGlipTeam**](TeamsApi.md#ArchiveGlipTeam) | **Post** /restapi/v1.0/glip/teams/{chatId}/archive | Archive Team
[**CreateGlipTeam**](TeamsApi.md#CreateGlipTeam) | **Post** /restapi/v1.0/glip/teams | Create Team
[**DeleteGlipTeam**](TeamsApi.md#DeleteGlipTeam) | **Delete** /restapi/v1.0/glip/teams/{chatId} | Delete Team
[**JoinGlipTeam**](TeamsApi.md#JoinGlipTeam) | **Post** /restapi/v1.0/glip/teams/{chatId}/join | Join Team
[**LeaveGlipTeam**](TeamsApi.md#LeaveGlipTeam) | **Post** /restapi/v1.0/glip/teams/{chatId}/leave | Leave Team
[**ListGlipTeams**](TeamsApi.md#ListGlipTeams) | **Get** /restapi/v1.0/glip/teams | Get Teams
[**PatchGlipEveryone**](TeamsApi.md#PatchGlipEveryone) | **Patch** /restapi/v1.0/glip/everyone | Update Everyone Сhat
[**PatchGlipTeam**](TeamsApi.md#PatchGlipTeam) | **Patch** /restapi/v1.0/glip/teams/{chatId} | Update Team
[**ReadGlipEveryone**](TeamsApi.md#ReadGlipEveryone) | **Get** /restapi/v1.0/glip/everyone | Get Everyone Chat
[**ReadGlipTeam**](TeamsApi.md#ReadGlipTeam) | **Get** /restapi/v1.0/glip/teams/{chatId} | Get Team
[**RemoveGlipTeamMembers**](TeamsApi.md#RemoveGlipTeamMembers) | **Post** /restapi/v1.0/glip/teams/{chatId}/remove | Remove Team Members
[**UnarchiveGlipTeam**](TeamsApi.md#UnarchiveGlipTeam) | **Post** /restapi/v1.0/glip/teams/{chatId}/unarchive | Unarchive Team



## AddGlipTeamMembers

> AddGlipTeamMembers(ctx, chatId, body)
Add Team Members

Adds members to the team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to add members to. | 
**body** | [**GlipPostMembersListBody**](GlipPostMembersListBody.md)| JSON body | 

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


## ArchiveGlipTeam

> ArchiveGlipTeam(ctx, chatId)
Archive Team

Archives a team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to be archived. | 

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


## CreateGlipTeam

> GlipTeamInfo CreateGlipTeam(ctx, body)
Create Team

Create a Team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GlipPostTeamBody**](GlipPostTeamBody.md)| JSON body | 

### Return type

[**GlipTeamInfo**](GlipTeamInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlipTeam

> DeleteGlipTeam(ctx, chatId)
Delete Team

Deletes the team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to be deleted. | 

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


## JoinGlipTeam

> JoinGlipTeam(ctx, chatId)
Join Team

Join to the team

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to be joined. | 

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


## LeaveGlipTeam

> LeaveGlipTeam(ctx, chatId)
Leave Team

The user leaves the team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to be left. | 

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


## ListGlipTeams

> GlipTeamsList ListGlipTeams(ctx, optional)
Get Teams

Returns the list of teams where the user is a member and public teams that can be joined. All records in response are sorted by creation time of a chat in ascending order.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListGlipTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGlipTeamsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordCount** | **optional.Int32**| Number of teams to be fetched by one request. The maximum value is 250, by default - 30 | [default to 30]
 **pageToken** | **optional.String**| Pagination token. | 

### Return type

[**GlipTeamsList**](GlipTeamsList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGlipEveryone

> GlipEveryoneInfo PatchGlipEveryone(ctx, optional)
Update Everyone Сhat

Updates Everyone chat information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PatchGlipEveryoneOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PatchGlipEveryoneOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateGlipEveryoneRequest**](UpdateGlipEveryoneRequest.md)| JSON body | 

### Return type

[**GlipEveryoneInfo**](GlipEveryoneInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGlipTeam

> GlipTeamInfo PatchGlipTeam(ctx, chatId, body)
Update Team

Updates a team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to be updated. | 
**body** | [**GlipPatchTeamBody**](GlipPatchTeamBody.md)| JSON body | 

### Return type

[**GlipTeamInfo**](GlipTeamInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipEveryone

> GlipEveryoneInfo ReadGlipEveryone(ctx, )
Get Everyone Chat

Returns information about Everyone chat.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GlipEveryoneInfo**](GlipEveryoneInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipTeam

> GlipTeamInfo ReadGlipTeam(ctx, chatId)
Get Team

Returns information about a team by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to be returned. | 

### Return type

[**GlipTeamInfo**](GlipTeamInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGlipTeamMembers

> RemoveGlipTeamMembers(ctx, chatId, body)
Remove Team Members

Removes members from a team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to remove members from. | 
**body** | [**GlipPostMembersIdsListBody**](GlipPostMembersIdsListBody.md)| JSON body | 

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


## UnarchiveGlipTeam

> UnarchiveGlipTeam(ctx, chatId)
Unarchive Team

Unarchives a team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a team to be made active. | 

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

