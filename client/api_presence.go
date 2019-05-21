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
	"fmt"
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

type PresenceApiService service

/*
PresenceApiService Get User Presence Status List
Returns presence status of all extensions of an account. Please note: The presenceStatus is returned as Offline (the parameters telephonyStatus, message, userStatus and dndStatus are not returned at all) for the following extension types: Department, Announcement Only, Voicemail (Take Messages Only), Fax User, Paging Only Group, Shared Lines Group, IVR Menu, Application Extension.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param optional nil or *ReadAccountPresenceOpts - Optional Parameters:
 * @param "DetailedTelephonyState" (optional.Bool) -  Whether to return detailed telephony state
 * @param "SipData" (optional.Bool) -  Whether to return SIP data
 * @param "Page" (optional.Int32) -  Page number for account presence information
 * @param "PerPage" (optional.Int32) -  Number for account presence information items per page
@return AccountPresenceInfo
*/

type ReadAccountPresenceOpts struct {
	DetailedTelephonyState optional.Bool
	SipData                optional.Bool
	Page                   optional.Int32
	PerPage                optional.Int32
}

func (a *PresenceApiService) ReadAccountPresence(ctx context.Context, accountId string, localVarOptionals *ReadAccountPresenceOpts) (AccountPresenceInfo, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccountPresenceInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/presence"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.DetailedTelephonyState.IsSet() {
		localVarQueryParams.Add("detailedTelephonyState", parameterToString(localVarOptionals.DetailedTelephonyState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SipData.IsSet() {
		localVarQueryParams.Add("sipData", parameterToString(localVarOptionals.SipData.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("perPage", parameterToString(localVarOptionals.PerPage.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
			var v AccountPresenceInfo
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
PresenceApiService Get User Presence Status
Returns presence status of an extension or several extensions by their ID(s). Batch request is supported. The &#39;presenceStatus&#39; is returned as Offline (the parameters &#39;telephonyStatus&#39;, &#39;message&#39;, &#39;userStatus&#39; and &#39;dndStatus&#39; are not returned at all) for the following extension types: Department/Announcement Only/Take Messages Only (Voicemail)/Fax User/Paging Only Group/Shared Lines Group/IVR Menu/Application Extension/Park Location.If the user requests his/her own presence status, the response contains actual presence status even if the status publication is turned off. Batch request is supported. For batch requests the number of extensions in one request is limited to 30. If more extensions are included in the request, the error code 400 Bad Request is returned with the logical error code InvalidMultipartRequest and the corresponding message &#39;Extension Presence Info multipart request is limited to 30 extensions&#39;.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param optional nil or *ReadUserPresenceStatusOpts - Optional Parameters:
 * @param "DetailedTelephonyState" (optional.Bool) -  Whether to return detailed telephony state
 * @param "SipData" (optional.Bool) -  Whether to return SIP data
@return GetPresenceInfo
*/

type ReadUserPresenceStatusOpts struct {
	DetailedTelephonyState optional.Bool
	SipData                optional.Bool
}

func (a *PresenceApiService) ReadUserPresenceStatus(ctx context.Context, accountId string, extensionId string, localVarOptionals *ReadUserPresenceStatusOpts) (GetPresenceInfo, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetPresenceInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/presence"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.DetailedTelephonyState.IsSet() {
		localVarQueryParams.Add("detailedTelephonyState", parameterToString(localVarOptionals.DetailedTelephonyState.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SipData.IsSet() {
		localVarQueryParams.Add("sipData", parameterToString(localVarOptionals.SipData.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
			var v GetPresenceInfo
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
PresenceApiService Update User Presence Status
Updates user-defined extension presence status, status message and DnD status by extension ID. Supported for regular User extensions only. The extension types listed do not support presence status: Department, Announcement Only, Take Messages Only (Voicemail), Fax User, Paging Only Group, Shared Lines Group, IVR Menu, Application Extension.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param optional nil or *UpdateUserPresenceStatusOpts - Optional Parameters:
 * @param "Body" (optional.Interface of PresenceInfoResource) -
@return PresenceInfoResource
*/

type UpdateUserPresenceStatusOpts struct {
	Body optional.Interface
}

func (a *PresenceApiService) UpdateUserPresenceStatus(ctx context.Context, extensionId string, accountId string, localVarOptionals *UpdateUserPresenceStatusOpts) (PresenceInfoResource, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PresenceInfoResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/presence"
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarOptionalBody, localVarOptionalBodyok := localVarOptionals.Body.Value().(PresenceInfoResource)
		if !localVarOptionalBodyok {
			return localVarReturnValue, nil, reportError("body should be PresenceInfoResource")
		}
		localVarPostBody = &localVarOptionalBody
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
			var v PresenceInfoResource
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
