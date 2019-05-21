# \AuthenticationApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](AuthenticationApi.md#Authorize) | **Get** /restapi/oauth/authorize | Authorize
[**GetToken**](AuthenticationApi.md#GetToken) | **Post** /restapi/oauth/token | Get Token
[**RevokeToken**](AuthenticationApi.md#RevokeToken) | **Post** /restapi/oauth/revoke | Revoke Token



## Authorize

> Authorize(ctx, responseType, redirectUri, clientId, optional)
Authorize

Returns link to a login page location.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**responseType** | **string**| Determines authorization flow: **code** - Authorization Code, **token** - Implicit Grant | 
**redirectUri** | **string**| This is a callback URI which determines where the response is sent. The value of this parameter must exactly match one of the URIs you have provided for your app upon registration | 
**clientId** | **string**| Identifier (key) of a client application | 
 **optional** | ***AuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorizeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **state** | **optional.String**| Client state. Returned back to the client at the end of the flow | 
 **brandId** | **optional.String**| Brand identifier. If it is not provided in request, server will try to determine brand from client app profile. The default value is &#39;1210&#39; - RingCentral US | 
 **display** | **optional.String**| Style of login form. The default value is &#39;page&#39;. The &#39;popup&#39; and &#39;touch&#39; values are featured for mobile applications | 
 **prompt** | **optional.String**| Specifies which login form will be displayed. Space-separated set of the following values: &#39;login&#39; - official RingCentral login form, &#39;sso&#39; - Single Sign-On login form, &#39;consent&#39; - form to show the requested scope and prompt user for consent. Either &#39;login&#39; or &#39;sso&#39; (or both) must be specified. The default value is &#39;login&amp;sso&#39; | 
 **localeId** | **optional.String**| Localization code of a language. Overwrites &#39;Accept-Language&#39; header value | 
 **uiLocales** | **optional.String**| Localization code of a language. Overwrites &#39;localeId&#39; parameter value | 
 **uiOptions** | **optional.String**| User interface options data | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> TokenInfo GetToken(ctx, optional)
Get Token

Returns access tokens for making API requests

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Phone number linked to an account or extension in E.164 format with or without leading &#39;+&#39; sign | 
 **password** | **optional.String**| User&#39;s password | 
 **extension** | **optional.String**| Optional. Extension short number. If company number is specified as a username, and extension is not specified, the server will attempt to authenticate client as main company administrator | 
 **grantType** | **optional.String**| Grant type | [default to password]
 **code** | **optional.String**| Authorization code | 
 **redirectUri** | **optional.String**| This is a callback URI which determines where the response is sent. The value of this parameter must exactly match one of the URIs you have provided for your app upon registration | 
 **accessTokenTtl** | **optional.Int64**| Access token lifetime in seconds | [default to 3600]
 **refreshTokenTtl** | **optional.Int64**| Refresh token lifetime in seconds | [default to 604800]
 **scope** | **optional.String**| List of API permissions to be used with access token. Can be omitted when requesting all permissions defined during the application registration phase | 
 **refreshToken** | **optional.String**| Previously issued refresh token. This is the only formData field used for the Refresh Token Flow. | 
 **endpointId** | **optional.String**| The unique identifier of a client application. If not specified, the previously specified or auto generated value is used by default | 

### Return type

[**TokenInfo**](TokenInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeToken

> RevokeToken(ctx, token)
Revoke Token

Revokes access/refresh token. Requests to this endpoint must be authenticated with HTTP Basic scheme using the application key and application secret as login and password, correspondingly.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**| Active access or refresh token to be revoked | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

