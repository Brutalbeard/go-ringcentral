# \RegionalSettingsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCountries**](RegionalSettingsApi.md#ListCountries) | **Get** /restapi/v1.0/dictionary/country | Get Country List
[**ListLanguages**](RegionalSettingsApi.md#ListLanguages) | **Get** /restapi/v1.0/dictionary/language | Get Language List
[**ListLocations**](RegionalSettingsApi.md#ListLocations) | **Get** /restapi/v1.0/dictionary/location | Get Location List
[**ListStates**](RegionalSettingsApi.md#ListStates) | **Get** /restapi/v1.0/dictionary/state | Get States List
[**ListTimezones**](RegionalSettingsApi.md#ListTimezones) | **Get** /restapi/v1.0/dictionary/timezone | Get Timezone List
[**ReadCountry**](RegionalSettingsApi.md#ReadCountry) | **Get** /restapi/v1.0/dictionary/country/{countryId} | Get Country
[**ReadLanguage**](RegionalSettingsApi.md#ReadLanguage) | **Get** /restapi/v1.0/dictionary/language/{languageId} | Get Language
[**ReadState**](RegionalSettingsApi.md#ReadState) | **Get** /restapi/v1.0/dictionary/state/{stateId} | Get State
[**ReadTimezone**](RegionalSettingsApi.md#ReadTimezone) | **Get** /restapi/v1.0/dictionary/timezone/{timezoneId} | Get Timezone



## ListCountries

> GetCountryListResponse ListCountries(ctx, optional)
Get Country List

Returns all the countries available for calling.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCountriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCountriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginAllowed** | **optional.Bool**| Specifies whether login with the phone numbers of this country is enabled or not | 
 **signupAllowed** | **optional.Bool**| Indicates whether signup/billing is allowed for a country. If not specified all countries are returned (according to other filters specified if any) | 
 **numberSelling** | **optional.Bool**| Specifies if RingCentral sells phone numbers of this country | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **freeSoftphoneLine** | **optional.Bool**| Specifies if free phone line for softphone is available for a country or not | 

### Return type

[**GetCountryListResponse**](GetCountryListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLanguages

> LanguageList ListLanguages(ctx, )
Get Language List

Returns the information about supported languages.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LanguageList**](LanguageList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocations

> GetLocationListResponse ListLocations(ctx, optional)
Get Location List

Returns all available locations for a certain state.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListLocationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **optional.String**| Sorts results by the property specified | [default to City]
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **stateId** | **optional.String**| Internal identifier of a state | 
 **withNxx** | **optional.Bool**| Specifies if nxx codes are returned | 

### Return type

[**GetLocationListResponse**](GetLocationListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStates

> GetStateListResponse ListStates(ctx, optional)
Get States List

Returns all the states of a certain country

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListStatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListStatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allCountries** | **optional.Bool**| If set to &#39;True&#39; then states for all countries are returned and &#x60;countryId&#x60; is ignored, even if specified. If the value is empty then the parameter is ignored | 
 **countryId** | **optional.Int32**| Internal identifier of a country | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **withPhoneNumbers** | **optional.Bool**| If &#39;True&#39;, the list of states with phone numbers available for buying is returned | [default to false]

### Return type

[**GetStateListResponse**](GetStateListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTimezones

> GetTimezoneListResponse ListTimezones(ctx, optional)
Get Timezone List

Returns all available timezones.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListTimezonesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTimezonesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.String**| Indicates the page size (number of items) | [default to 100]

### Return type

[**GetTimezoneListResponse**](GetTimezoneListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCountry

> GetCountryInfoDictionaryResponse ReadCountry(ctx, countryId)
Get Country

Returns the information on a specific country.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **int32**| Internal identifier of a country | 

### Return type

[**GetCountryInfoDictionaryResponse**](GetCountryInfoDictionaryResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadLanguage

> LanguageInfo ReadLanguage(ctx, languageId)
Get Language

Returns language by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **int32**| Internal identifier of a language | 

### Return type

[**LanguageInfo**](LanguageInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadState

> GetStateInfoResponse ReadState(ctx, stateId)
Get State

Returns the information on a specific state.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateId** | **int32**| Internal identifier of a state | 

### Return type

[**GetStateInfoResponse**](GetStateInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTimezone

> GetTimezoneInfoResponse ReadTimezone(ctx, timezoneId, optional)
Get Timezone

Returns the information on a certain timezone.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timezoneId** | **int32**| Internal identifier of a timezone | 
 **optional** | ***ReadTimezoneOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadTimezoneOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.String**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.String**| Indicates the page size (number of items) | [default to 100]

### Return type

[**GetTimezoneInfoResponse**](GetTimezoneInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

