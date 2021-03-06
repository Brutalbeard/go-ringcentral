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

type PhoneNumbersApiService service

/*
PhoneNumbersApiService Get Company Phone Number List
Returns the list of phone numbers assigned to RingCentral customer account. Both company-level and extension-level numbers are returned.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param optional nil or *ListAccountPhoneNumbersOpts - Optional Parameters:
 * @param "Page" (optional.Int32) -  Indicates the page number to retrieve. Only positive number values are accepted
 * @param "PerPage" (optional.Int32) -  Indicates the page size (number of items)
 * @param "UsageType" (optional.Interface of []string) -  Usage type of a phone number
@return AccountPhoneNumbers
*/

type ListAccountPhoneNumbersOpts struct {
	Page      optional.Int32
	PerPage   optional.Int32
	UsageType optional.Interface
}

func (a *PhoneNumbersApiService) ListAccountPhoneNumbers(ctx context.Context, accountId string, localVarOptionals *ListAccountPhoneNumbersOpts) (AccountPhoneNumbers, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccountPhoneNumbers
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/phone-number"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PerPage.IsSet() {
		localVarQueryParams.Add("perPage", parameterToString(localVarOptionals.PerPage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UsageType.IsSet() {
		localVarQueryParams.Add("usageType", parameterToString(localVarOptionals.UsageType.Value(), "multi"))
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
		if localVarHttpResponse.StatusCode == 0 {
			var v AccountPhoneNumbers
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
PhoneNumbersApiService Get Extension Phone Number List
Returns the list of phone numbers that are used by a particular extension, and can be filtered by the phone number type. The returned list contains all numbers which are directly mapped to a given extension plus the features and also company-level numbers which may be used when performing different operations on behalf of this extension.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param optional nil or *ListExtensionPhoneNumbersOpts - Optional Parameters:
 * @param "UsageType" (optional.Interface of []string) -  Usage type of a phone number
 * @param "Page" (optional.Int32) -  Indicates the page number to retrieve. Only positive number values are allowed. Default value is '1'
 * @param "PerPage" (optional.Int32) -  Indicates the page size (number of items). If not specified, the value is '100' by default
@return GetExtensionPhoneNumbersResponse
*/

type ListExtensionPhoneNumbersOpts struct {
	UsageType optional.Interface
	Page      optional.Int32
	PerPage   optional.Int32
}

func (a *PhoneNumbersApiService) ListExtensionPhoneNumbers(ctx context.Context, accountId string, extensionId string, localVarOptionals *ListExtensionPhoneNumbersOpts) (GetExtensionPhoneNumbersResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetExtensionPhoneNumbersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/phone-number"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.UsageType.IsSet() {
		localVarQueryParams.Add("usageType", parameterToString(localVarOptionals.UsageType.Value(), "multi"))
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
			var v GetExtensionPhoneNumbersResponse
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
PhoneNumbersApiService Parse Phone Number
Returns one or more parsed and/or formatted phone numbers that are passed as a string.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body JSON body
 * @param optional nil or *ParsePhoneNumberOpts - Optional Parameters:
 * @param "HomeCountry" (optional.String) -  Internal identifier of a home country. The default value is ISO code (ISO 3166) of the user's home country or brand country, if the user is undefined
 * @param "NationalAsPriority" (optional.Bool) -  The default value is 'False'. If 'True', the numbers that are closer to the home country are given higher priority
@return ParsePhoneNumberResponse
*/

type ParsePhoneNumberOpts struct {
	HomeCountry        optional.String
	NationalAsPriority optional.Bool
}

func (a *PhoneNumbersApiService) ParsePhoneNumber(ctx context.Context, body ParsePhoneNumberRequest, localVarOptionals *ParsePhoneNumberOpts) (ParsePhoneNumberResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ParsePhoneNumberResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/number-parser/parse"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.HomeCountry.IsSet() {
		localVarQueryParams.Add("homeCountry", parameterToString(localVarOptionals.HomeCountry.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NationalAsPriority.IsSet() {
		localVarQueryParams.Add("nationalAsPriority", parameterToString(localVarOptionals.NationalAsPriority.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

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
	localVarPostBody = &body
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
		if localVarHttpResponse.StatusCode == 0 {
			var v ParsePhoneNumberResponse
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
PhoneNumbersApiService Get Phone Number
Returns the phone number(s) belonging to a certain account or extension by phoneNumberId(s). Batch request is supported.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param phoneNumberId Internal identifier of a phone number
@return CompanyPhoneNumberInfo
*/
func (a *PhoneNumbersApiService) ReadAccountPhoneNumber(ctx context.Context, accountId string, phoneNumberId int32) (CompanyPhoneNumberInfo, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CompanyPhoneNumberInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/phone-number/{phoneNumberId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneNumberId"+"}", fmt.Sprintf("%v", phoneNumberId), -1)

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
		if localVarHttpResponse.StatusCode == 0 {
			var v CompanyPhoneNumberInfo
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
