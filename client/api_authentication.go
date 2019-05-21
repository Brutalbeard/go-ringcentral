/*
 * RingCentral Connect Platform API
 *
 * RingCentral Connect Platform API
 *
 * API version: 1.0.36
 * Contact: platform@ringcentral.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ringcentral

import (
	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type AuthenticationApiService service

/*
AuthenticationApiService Authorize
Returns link to a login page location.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param responseType Determines authorization flow: **code** - Authorization Code, **token** - Implicit Grant
 * @param redirectUri This is a callback URI which determines where the response is sent. The value of this parameter must exactly match one of the URIs you have provided for your app upon registration
 * @param clientId Identifier (key) of a client application
 * @param optional nil or *AuthorizeOpts - Optional Parameters:
 * @param "State" (optional.String) -  Client state. Returned back to the client at the end of the flow
 * @param "BrandId" (optional.String) -  Brand identifier. If it is not provided in request, server will try to determine brand from client app profile. The default value is '1210' - RingCentral US
 * @param "Display" (optional.String) -  Style of login form. The default value is 'page'. The 'popup' and 'touch' values are featured for mobile applications
 * @param "Prompt" (optional.String) -  Specifies which login form will be displayed. Space-separated set of the following values: 'login' - official RingCentral login form, 'sso' - Single Sign-On login form, 'consent' - form to show the requested scope and prompt user for consent. Either 'login' or 'sso' (or both) must be specified. The default value is 'login&sso'
 * @param "LocaleId" (optional.String) -  Localization code of a language. Overwrites 'Accept-Language' header value
 * @param "UiLocales" (optional.String) -  Localization code of a language. Overwrites 'localeId' parameter value
 * @param "UiOptions" (optional.String) -  User interface options data
*/

type AuthorizeOpts struct {
	State     optional.String
	BrandId   optional.String
	Display   optional.String
	Prompt    optional.String
	LocaleId  optional.String
	UiLocales optional.String
	UiOptions optional.String
}

func (a *AuthenticationApiService) Authorize(ctx context.Context, responseType string, redirectUri string, clientId string, localVarOptionals *AuthorizeOpts) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/oauth/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("response_type", parameterToString(responseType, ""))
	localVarFormParams.Add("redirect_uri", parameterToString(redirectUri, ""))
	localVarFormParams.Add("client_id", parameterToString(clientId, ""))
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarFormParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.BrandId.IsSet() {
		localVarFormParams.Add("brand_id", parameterToString(localVarOptionals.BrandId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Display.IsSet() {
		localVarFormParams.Add("display", parameterToString(localVarOptionals.Display.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prompt.IsSet() {
		localVarFormParams.Add("prompt", parameterToString(localVarOptionals.Prompt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.LocaleId.IsSet() {
		localVarFormParams.Add("localeId", parameterToString(localVarOptionals.LocaleId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UiLocales.IsSet() {
		localVarFormParams.Add("ui_locales", parameterToString(localVarOptionals.UiLocales.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UiOptions.IsSet() {
		localVarFormParams.Add("ui_options", parameterToString(localVarOptionals.UiOptions.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
AuthenticationApiService Get Token
Returns access tokens for making API requests
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetTokenOpts - Optional Parameters:
 * @param "Username" (optional.String) -  Phone number linked to an account or extension in E.164 format with or without leading '+' sign
 * @param "Password" (optional.String) -  User's password
 * @param "Extension" (optional.String) -  Optional. Extension short number. If company number is specified as a username, and extension is not specified, the server will attempt to authenticate client as main company administrator
 * @param "GrantType" (optional.String) -  Grant type
 * @param "Code" (optional.String) -  Authorization code
 * @param "RedirectUri" (optional.String) -  This is a callback URI which determines where the response is sent. The value of this parameter must exactly match one of the URIs you have provided for your app upon registration
 * @param "AccessTokenTtl" (optional.Int64) -  Access token lifetime in seconds
 * @param "RefreshTokenTtl" (optional.Int64) -  Refresh token lifetime in seconds
 * @param "Scope" (optional.String) -  List of API permissions to be used with access token. Can be omitted when requesting all permissions defined during the application registration phase
 * @param "RefreshToken" (optional.String) -  Previously issued refresh token. This is the only formData field used for the Refresh Token Flow.
 * @param "EndpointId" (optional.String) -  The unique identifier of a client application. If not specified, the previously specified or auto generated value is used by default
@return TokenInfo
*/

type GetTokenOpts struct {
	Username        optional.String
	Password        optional.String
	Extension       optional.String
	GrantType       optional.String
	Code            optional.String
	RedirectUri     optional.String
	AccessTokenTtl  optional.Int64
	RefreshTokenTtl optional.Int64
	Scope           optional.String
	RefreshToken    optional.String
	EndpointId      optional.String
}

func (a *AuthenticationApiService) GetToken(ctx context.Context, localVarOptionals *GetTokenOpts) (TokenInfo, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TokenInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/oauth/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.Username.IsSet() {
		localVarFormParams.Add("username", parameterToString(localVarOptionals.Username.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Password.IsSet() {
		localVarFormParams.Add("password", parameterToString(localVarOptionals.Password.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Extension.IsSet() {
		localVarFormParams.Add("extension", parameterToString(localVarOptionals.Extension.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GrantType.IsSet() {
		localVarFormParams.Add("grant_type", parameterToString(localVarOptionals.GrantType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Code.IsSet() {
		localVarFormParams.Add("code", parameterToString(localVarOptionals.Code.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarFormParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccessTokenTtl.IsSet() {
		localVarFormParams.Add("access_token_ttl", parameterToString(localVarOptionals.AccessTokenTtl.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RefreshTokenTtl.IsSet() {
		localVarFormParams.Add("refresh_token_ttl", parameterToString(localVarOptionals.RefreshTokenTtl.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarFormParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RefreshToken.IsSet() {
		localVarFormParams.Add("refresh_token", parameterToString(localVarOptionals.RefreshToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndpointId.IsSet() {
		localVarFormParams.Add("endpoint_id", parameterToString(localVarOptionals.EndpointId.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v TokenInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AuthenticationApiService Revoke Token
Revokes access/refresh token. Requests to this endpoint must be authenticated with HTTP Basic scheme using the application key and application secret as login and password, correspondingly.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param token Active access or refresh token to be revoked
*/
func (a *AuthenticationApiService) RevokeToken(ctx context.Context, token string) (*http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/oauth/revoke"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("token", parameterToString(token, ""))
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}