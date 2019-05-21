# TokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access token to pass to subsequent API requests | [optional] 
**ExpiresIn** | **int32** | Issued access token TTL (time to live), in seconds | [optional] 
**RefreshToken** | **string** | Refresh token to get a new access token, when the issued one expires | [optional] 
**RefreshTokenExpiresIn** | **int32** | Issued refresh token TTL (time to live), in seconds | [optional] 
**Scope** | **string** | List of permissions allowed with this access token, white-space separated | [optional] 
**TokenType** | **string** | Type of token. The only possible value supported is Bearer. This value should be used when specifying access token in Authorization header of subsequent API requests | [optional] 
**OwnerId** | **string** | Extension identifier | [optional] 
**EndpointId** | **string** | Application instance identifier | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


