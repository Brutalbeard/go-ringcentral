# \ChatsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignGlipGroupMembers**](ChatsApi.md#AssignGlipGroupMembers) | **Post** /restapi/v1.0/glip/groups/{groupId}/bulk-assign | Edit Group Members
[**CreateGlipGroup**](ChatsApi.md#CreateGlipGroup) | **Post** /restapi/v1.0/glip/groups | Create Group
[**FavoriteGlipChat**](ChatsApi.md#FavoriteGlipChat) | **Post** /restapi/v1.0/glip/chats/{chatId}/favorite | Add Chat to Favorites
[**ListFavoriteChats**](ChatsApi.md#ListFavoriteChats) | **Get** /restapi/v1.0/glip/favorites | Get Favorite Chats
[**ListGlipChats**](ChatsApi.md#ListGlipChats) | **Get** /restapi/v1.0/glip/chats | Get Chats
[**ListGlipGroups**](ChatsApi.md#ListGlipGroups) | **Get** /restapi/v1.0/glip/groups | Get User Groups
[**ListRecentChats**](ChatsApi.md#ListRecentChats) | **Get** /restapi/v1.0/glip/recent/chats | Get Recent Chats
[**MarkChatRead**](ChatsApi.md#MarkChatRead) | **Post** /restapi/v1.0/glip/chats/{chatId}/read | Mark Chat as Read
[**MarkChatUnread**](ChatsApi.md#MarkChatUnread) | **Post** /restapi/v1.0/glip/chats/{chatId}/unread | Mark Chat as Unread
[**ReadGlipChat**](ChatsApi.md#ReadGlipChat) | **Get** /restapi/v1.0/glip/chats/{chatId} | Get Chat
[**ReadGlipGroup**](ChatsApi.md#ReadGlipGroup) | **Get** /restapi/v1.0/glip/groups/{groupId} | Get Group
[**UnfavoriteGlipChat**](ChatsApi.md#UnfavoriteGlipChat) | **Post** /restapi/v1.0/glip/chats/{chatId}/unfavorite | Remove Chat from Favorites



## AssignGlipGroupMembers

> GlipGroupInfo AssignGlipGroupMembers(ctx, groupId, body)
Edit Group Members

Updates group members. **Please note:** Only groups of 'Team' type can be updated. Currently only one operation at a time (either adding or removal) is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| Internal identifier of a group | 
**body** | [**EditGroupRequest**](EditGroupRequest.md)| JSON body | 

### Return type

[**GlipGroupInfo**](GlipGroupInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlipGroup

> GlipGroupInfo CreateGlipGroup(ctx, body)
Create Group

Creates a new private chat/team.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GlipCreateGroup**](GlipCreateGroup.md)| JSON body | 

### Return type

[**GlipGroupInfo**](GlipGroupInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FavoriteGlipChat

> FavoriteGlipChat(ctx, chatId)
Add Chat to Favorites

Adds chat to the users's favorite list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a chat to add to favorite list. | 

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


## ListFavoriteChats

> GlipChatsListWithoutNavigation ListFavoriteChats(ctx, optional)
Get Favorite Chats

Returns user's favorite chats.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFavoriteChatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFavoriteChatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordCount** | **optional.Int32**| Max number of chats to be fetched by one request (Not more than 250). | [default to 30]

### Return type

[**GlipChatsListWithoutNavigation**](GlipChatsListWithoutNavigation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlipChats

> GlipChatsList ListGlipChats(ctx, optional)
Get Chats

Returns the list of chats where the user is a member and also public teams that can be joined. All records in response are sorted by creation time of a chat in ascending order.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListGlipChatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGlipChatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of []string**](string.md)| Type of chats to be fetched. By default all type of chats will be fetched | 
 **recordCount** | **optional.Int32**| Number of chats to be fetched by one request. The maximum value is 250, by default - 30. | [default to 30]
 **pageToken** | **optional.String**| Pagination token. | 

### Return type

[**GlipChatsList**](GlipChatsList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlipGroups

> GlipGroupList ListGlipGroups(ctx, optional)
Get User Groups

Returns the list of groups where the user is a member.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListGlipGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGlipGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Type of groups to be fetched (by default all type of groups will be fetched) | 
 **recordCount** | **optional.Int32**| Number of groups to be fetched by one request. The maximum value is 250, by default - 30 | [default to 30]
 **pageToken** | **optional.String**| Pagination token. | 

### Return type

[**GlipGroupList**](GlipGroupList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecentChats

> GlipChatsListWithoutNavigation ListRecentChats(ctx, optional)
Get Recent Chats

Returns recent chats where the user is a member. All records in response are sorted by the `lastModifiedTime` in descending order (the latest changed chat is displayed first on page)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRecentChatsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRecentChatsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of []string**](string.md)| Type of chats to be fetched. By default all chat types are returned | 
 **recordCount** | **optional.Int32**| Max number of chats to be fetched by one request (Not more than 250). | [default to 30]

### Return type

[**GlipChatsListWithoutNavigation**](GlipChatsListWithoutNavigation.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkChatRead

> MarkChatRead(ctx, chatId)
Mark Chat as Read

Marks chat as read.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Id of chat to be marked | 

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


## MarkChatUnread

> MarkChatUnread(ctx, chatId)
Mark Chat as Unread

Marks chat as unread.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Id of chat to be marked | 

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


## ReadGlipChat

> GlipChatInfo ReadGlipChat(ctx, chatId)
Get Chat

Returns information about a chat by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a chat. If tilda (~) is specified, then &#x60;/me&#x60; (Personal) chat will be returned | 

### Return type

[**GlipChatInfo**](GlipChatInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipGroup

> GlipGroupInfo ReadGlipGroup(ctx, groupId)
Get Group

Returns information about a group or multiple groups by their ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | [**[]string**](string.md)| Internal identifier of a group to be returned, the maximum number of IDs is 30 | 

### Return type

[**GlipGroupInfo**](GlipGroupInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/mixed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfavoriteGlipChat

> UnfavoriteGlipChat(ctx, chatId)
Remove Chat from Favorites

Removes chat from users's favorite list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a chat to remove from favorite list. | 

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

