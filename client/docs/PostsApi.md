# \PostsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlipCard**](PostsApi.md#CreateGlipCard) | **Post** /restapi/v1.0/glip/cards | Create Card
[**CreateGlipGroupPost**](PostsApi.md#CreateGlipGroupPost) | **Post** /restapi/v1.0/glip/groups/{groupId}/posts | Create Post in Group
[**CreateGlipPost**](PostsApi.md#CreateGlipPost) | **Post** /restapi/v1.0/glip/chats/{chatId}/posts | Create Post
[**CreatePost**](PostsApi.md#CreatePost) | **Post** /restapi/v1.0/glip/posts | Create Post
[**DeleteGlipCard**](PostsApi.md#DeleteGlipCard) | **Delete** /restapi/v1.0/glip/cards/{cardId} | Delete Card
[**DeleteGlipPost**](PostsApi.md#DeleteGlipPost) | **Delete** /restapi/v1.0/glip/chats/{chatId}/posts/{postId} | Delete Post
[**ListGlipGroupPosts**](PostsApi.md#ListGlipGroupPosts) | **Get** /restapi/v1.0/glip/groups/{groupId}/posts | Get Group Posts
[**ListGlipPosts**](PostsApi.md#ListGlipPosts) | **Get** /restapi/v1.0/glip/posts | Get Posts
[**PatchGlipPost**](PostsApi.md#PatchGlipPost) | **Patch** /restapi/v1.0/glip/chats/{chatId}/posts/{postId} | Update Post
[**ReadGlipCard**](PostsApi.md#ReadGlipCard) | **Get** /restapi/v1.0/glip/cards/{cardId} | Get Card
[**ReadGlipPost**](PostsApi.md#ReadGlipPost) | **Get** /restapi/v1.0/glip/chats/{chatId}/posts/{postId} | Get Post
[**ReadGlipPosts**](PostsApi.md#ReadGlipPosts) | **Get** /restapi/v1.0/glip/chats/{chatId}/posts | Get Posts
[**UpdateGlipCard**](PostsApi.md#UpdateGlipCard) | **Put** /restapi/v1.0/glip/cards/{cardId} | Update Card
[**UpdateGlipPostText**](PostsApi.md#UpdateGlipPostText) | **Put** /restapi/v1.0/glip/groups/{groupId}/posts/{postId}/text | Update Post



## CreateGlipCard

> GlipMessageAttachmentInfo CreateGlipCard(ctx, body, optional)
Create Card

Creates a new message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GlipMessageAttachmentInfoRequest**](GlipMessageAttachmentInfoRequest.md)|  | 
 **optional** | ***CreateGlipCardOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateGlipCardOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupId** | **optional.Int32**| Internal identifier of a group where to create a post with the card | 

### Return type

[**GlipMessageAttachmentInfo**](GlipMessageAttachmentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlipGroupPost

> GlipPostInfo CreateGlipGroupPost(ctx, groupId, body)
Create Post in Group

Creates a new post in a group specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| Internal identifier of a group. | 
**body** | [**GlipCreatePost**](GlipCreatePost.md)|  | 

### Return type

[**GlipPostInfo**](GlipPostInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlipPost

> GlipPostInfo CreateGlipPost(ctx, chatId, body)
Create Post

Create a Post

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a chat. | 
**body** | [**GlipPostPostBody**](GlipPostPostBody.md)| JSON body | 

### Return type

[**GlipPostInfo**](GlipPostInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePost

> GlipPostInfo CreatePost(ctx, body)
Create Post

Creates a post.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**GlipCreatePost**](GlipCreatePost.md)| JSON body | 

### Return type

[**GlipPostInfo**](GlipPostInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlipCard

> DeleteGlipCard(ctx, cardId)
Delete Card

Deletes a card by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string**| Card ID to be deleted. | 

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


## DeleteGlipPost

> DeleteGlipPost(ctx, chatId, postId)
Delete Post

Deletes a post.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a chat. | 
**postId** | **string**| Internal identifier of a post to be deleted. | 

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


## ListGlipGroupPosts

> GlipPosts ListGlipGroupPosts(ctx, groupId, optional)
Get Group Posts

Returns posts which are available for the current user by group ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| Internal identifier of a group | 
 **optional** | ***ListGlipGroupPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGlipGroupPostsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordCount** | **optional.Int32**| Max number of records to be returned | [default to 30]
 **pageToken** | **optional.String**| Pagination token | 

### Return type

[**GlipPosts**](GlipPosts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlipPosts

> GlipPosts ListGlipPosts(ctx, optional)
Get Posts

Returns posts available for the current user by group ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListGlipPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGlipPostsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.String**| Identifier of a group to filter posts | 
 **pageToken** | **optional.String**| Token of a page to be returned | 
 **recordCount** | **optional.Int64**| Number of records to be returned. The maximum value is 250, by default - 30 | [default to 30]

### Return type

[**GlipPosts**](GlipPosts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGlipPost

> GlipPostInfo PatchGlipPost(ctx, chatId, postId, body)
Update Post

Updates the post.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a chat. | 
**postId** | **string**| Internal identifier of a post to be updated. | 
**body** | [**GlipPatchPostBody**](GlipPatchPostBody.md)| JSON body | 

### Return type

[**GlipPostInfo**](GlipPostInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipCard

> GlipMessageAttachmentInfo ReadGlipCard(ctx, cardId)
Get Card

Returns card(s) with given id(s).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | [**[]string**](string.md)| Internal identifier of a card or comma separated list of card IDs. | 

### Return type

[**GlipMessageAttachmentInfo**](GlipMessageAttachmentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/mixed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipPost

> GlipPostInfo ReadGlipPost(ctx, chatId, postId)
Get Post

Returns information about a post by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a chat. | 
**postId** | **string**| Internal identifier of a post. | 

### Return type

[**GlipPostInfo**](GlipPostInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadGlipPosts

> GlipPostsList ReadGlipPosts(ctx, chatId, optional)
Get Posts

Returns a list of posts from the chat specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string**| Internal identifier of a chat. | 
 **optional** | ***ReadGlipPostsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadGlipPostsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordCount** | **optional.Int32**| Max number of posts to be fetched by one request (Not more than 250). | [default to 30]
 **pageToken** | **optional.String**| Pagination token. | 

### Return type

[**GlipPostsList**](GlipPostsList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlipCard

> UpdateGlipCard(ctx, cardId, body)
Update Card

Updates a card.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardId** | **string**| Internal identifier of a card | 
**body** | [**GlipMessageAttachmentInfoRequest**](GlipMessageAttachmentInfoRequest.md)|  | 

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


## UpdateGlipPostText

> string UpdateGlipPostText(ctx, groupId, postId, text)
Update Post

Modifies text of a post.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| Internal identifier of a group | 
**postId** | **string**| Internal identifier of a post | 
**text** | **string**|  | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: text/_*
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

