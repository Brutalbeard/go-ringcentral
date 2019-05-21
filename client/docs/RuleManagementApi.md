# \RuleManagementApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnsweringRule**](RuleManagementApi.md#CreateAnsweringRule) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule | Create Call Handling Rule
[**CreateCompanyAnsweringRule**](RuleManagementApi.md#CreateCompanyAnsweringRule) | **Post** /restapi/v1.0/account/{accountId}/answering-rule | Create Company Call Handling Rule
[**CreateCompanyGreeting**](RuleManagementApi.md#CreateCompanyGreeting) | **Post** /restapi/v1.0/account/{accountId}/greeting | Create Company Greeting
[**CreateCustomUserGreeting**](RuleManagementApi.md#CreateCustomUserGreeting) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/greeting | Create Custom User Greeting
[**DeleteAnsweringRule**](RuleManagementApi.md#DeleteAnsweringRule) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{ruleId} | Delete Call Handling Rule
[**DeleteCallRecordingCustomGreeting**](RuleManagementApi.md#DeleteCallRecordingCustomGreeting) | **Delete** /restapi/v1.0/account/{accountId}/call-recording/custom-greetings/{greetingId} | Delete Call Recording Custom Greeting
[**DeleteCallRecordingCustomGreetingList**](RuleManagementApi.md#DeleteCallRecordingCustomGreetingList) | **Delete** /restapi/v1.0/account/{accountId}/call-recording/custom-greetings | Delete Call Recording Custom Greeting List
[**DeleteCompanyAnsweringRule**](RuleManagementApi.md#DeleteCompanyAnsweringRule) | **Delete** /restapi/v1.0/account/{accountId}/answering-rule/{ruleId} | Delete Company Call Handling Rule
[**ListAnsweringRules**](RuleManagementApi.md#ListAnsweringRules) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule | Get Call Handling Rules
[**ListCallRecordingCustomGreetings**](RuleManagementApi.md#ListCallRecordingCustomGreetings) | **Get** /restapi/v1.0/account/{accountId}/call-recording/custom-greetings | Get Call Recording Custom Greeting List
[**ListCallRecordingExtensions**](RuleManagementApi.md#ListCallRecordingExtensions) | **Get** /restapi/v1.0/account/{accountId}/call-recording/extensions | Get Call Recording Extension List
[**ListCompanyAnsweringRules**](RuleManagementApi.md#ListCompanyAnsweringRules) | **Get** /restapi/v1.0/account/{accountId}/answering-rule | Get Company Call Handling Rule List
[**ListStandardGreetings**](RuleManagementApi.md#ListStandardGreetings) | **Get** /restapi/v1.0/dictionary/greeting | Get Standard Greeting List
[**ReadAnsweringRule**](RuleManagementApi.md#ReadAnsweringRule) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{ruleId} | Get Call Handling Rule
[**ReadCallRecordingSettings**](RuleManagementApi.md#ReadCallRecordingSettings) | **Get** /restapi/v1.0/account/{accountId}/call-recording | Get Call Recording Settings
[**ReadCompanyAnsweringRule**](RuleManagementApi.md#ReadCompanyAnsweringRule) | **Get** /restapi/v1.0/account/{accountId}/answering-rule/{ruleId} | Get Company Call Handling Rule
[**ReadCustomGreeting**](RuleManagementApi.md#ReadCustomGreeting) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/greeting/{greetingId} | Get Custom Greeting
[**ReadStandardGreeting**](RuleManagementApi.md#ReadStandardGreeting) | **Get** /restapi/v1.0/dictionary/greeting/{greetingId} | Get Standard Greeting
[**UpdateAnsweringRule**](RuleManagementApi.md#UpdateAnsweringRule) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/answering-rule/{ruleId} | Update Call Handling Rule
[**UpdateCallRecordingExtensionList**](RuleManagementApi.md#UpdateCallRecordingExtensionList) | **Post** /restapi/v1.0/account/{accountId}/call-recording/bulk-assign | Update Call Recording Extension List
[**UpdateCallRecordingSettings**](RuleManagementApi.md#UpdateCallRecordingSettings) | **Put** /restapi/v1.0/account/{accountId}/call-recording | Update Call Recording Settings
[**UpdateCompanyAnsweringRule**](RuleManagementApi.md#UpdateCompanyAnsweringRule) | **Put** /restapi/v1.0/account/{accountId}/answering-rule/{ruleId} | Update Company Call Handling Rule



## CreateAnsweringRule

> AnsweringRuleInfo CreateAnsweringRule(ctx, accountId, extensionId, body)
Create Call Handling Rule

Creates a custom answering rule for a particular caller ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**body** | [**CreateAnsweringRuleRequest**](CreateAnsweringRuleRequest.md)| JSON body | 

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompanyAnsweringRule

> CompanyAnsweringRuleInfo CreateCompanyAnsweringRule(ctx, accountId, body)
Create Company Call Handling Rule

Creates a company answering rule for a particular caller ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**body** | [**CompanyAnsweringRuleRequest**](CompanyAnsweringRuleRequest.md)| JSON body | 

### Return type

[**CompanyAnsweringRuleInfo**](CompanyAnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompanyGreeting

> CustomCompanyGreetingInfo CreateCompanyGreeting(ctx, accountId, uNKNOWNBASETYPE)
Create Company Greeting

Creates a custom company greeting.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**CustomCompanyGreetingInfo**](CustomCompanyGreetingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/mixed, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomUserGreeting

> CustomUserGreetingInfo CreateCustomUserGreeting(ctx, accountId, extensionId, uNKNOWNBASETYPE)
Create Custom User Greeting

Creates custom greeting for an extension user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**CustomUserGreetingInfo**](CustomUserGreetingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/mixed, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAnsweringRule

> DeleteAnsweringRule(ctx, accountId, extensionId, ruleId)
Delete Call Handling Rule

Deletes a custom answering rule by a particular ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**ruleId** | **string**| Internal identifier of an answering rule | 

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


## DeleteCallRecordingCustomGreeting

> DeleteCallRecordingCustomGreeting(ctx, accountId, greetingId)
Delete Call Recording Custom Greeting

Deletes call recording custom greeting(s).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**greetingId** | **string**|  | 

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


## DeleteCallRecordingCustomGreetingList

> DeleteCallRecordingCustomGreetingList(ctx, accountId)
Delete Call Recording Custom Greeting List

Deletes call recording custom greetings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

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


## DeleteCompanyAnsweringRule

> DeleteCompanyAnsweringRule(ctx, accountId, ruleId)
Delete Company Call Handling Rule

Deletes a company custom answering rule by a particular ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**ruleId** | **string**| Internal identifier of an answering rule | 

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


## ListAnsweringRules

> UserAnsweringRuleList ListAnsweringRules(ctx, accountId, extensionId, optional)
Get Call Handling Rules

Returns the extension answering rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListAnsweringRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAnsweringRulesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **view** | **optional.String**|  | [default to Simple]
 **enabledOnly** | **optional.String**|  | [default to false]
 **page** | **optional.String**|  | [default to 1]
 **perPage** | **optional.String**|  | [default to 100]

### Return type

[**UserAnsweringRuleList**](UserAnsweringRuleList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallRecordingCustomGreetings

> CallRecordingCustomGreetings ListCallRecordingCustomGreetings(ctx, accountId, optional)
Get Call Recording Custom Greeting List

Returns call recording custom greetings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***ListCallRecordingCustomGreetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallRecordingCustomGreetingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | 

### Return type

[**CallRecordingCustomGreetings**](CallRecordingCustomGreetings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallRecordingExtensions

> CallRecordingExtensions ListCallRecordingExtensions(ctx, accountId)
Get Call Recording Extension List

Returns the list of extensions to be recorded.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**CallRecordingExtensions**](CallRecordingExtensions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyAnsweringRules

> CompanyAnsweringRuleList ListCompanyAnsweringRules(ctx, accountId)
Get Company Call Handling Rule List

Returns a list of company answering rules.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]

### Return type

[**CompanyAnsweringRuleList**](CompanyAnsweringRuleList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStandardGreetings

> DictionaryGreetingList ListStandardGreetings(ctx, optional)
Get Standard Greeting List

Returns the list of predefined standard greetings. Custom greetings recorded by user are not returned in response to this request. See Get Extension Custom Greetings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListStandardGreetingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListStandardGreetingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted. | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items). | [default to 100]
 **type_** | **optional.String**| Type of a greeting, specifying the case when the greeting is played | 
 **usageType** | **optional.String**| Usage type of a greeting, specifying if the greeting is applied for user extension or department extension | 

### Return type

[**DictionaryGreetingList**](DictionaryGreetingList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnsweringRule

> AnsweringRuleInfo ReadAnsweringRule(ctx, accountId, extensionId, ruleId, optional)
Get Call Handling Rule

Returns an answering rule by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**ruleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either &#39;business-hours-rule&#39; or &#39;after-hours-rule&#39; | 
 **optional** | ***ReadAnsweringRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadAnsweringRuleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **showInactiveNumbers** | **optional.Bool**| Indicates whether inactive numbers should be returned or not | [default to false]

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCallRecordingSettings

> CallRecordingSettingsResource ReadCallRecordingSettings(ctx, accountId)
Get Call Recording Settings

Returns call recording settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 

### Return type

[**CallRecordingSettingsResource**](CallRecordingSettingsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCompanyAnsweringRule

> CompanyAnsweringRuleInfo ReadCompanyAnsweringRule(ctx, accountId, ruleId)
Get Company Call Handling Rule

Returns a company answering rule by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**ruleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either &#39;business-hours-rule&#39; or &#39;after-hours-rule&#39; | 

### Return type

[**CompanyAnsweringRuleInfo**](CompanyAnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCustomGreeting

> CustomUserGreetingInfo ReadCustomGreeting(ctx, accountId, extensionId, greetingId)
Get Custom Greeting

Returns a custom user greeting by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**greetingId** | **int32**| Internal identifier of a greeting | 

### Return type

[**CustomUserGreetingInfo**](CustomUserGreetingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadStandardGreeting

> DictionaryGreetingInfo ReadStandardGreeting(ctx, greetingId)
Get Standard Greeting

Returns a standard greeting by ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**greetingId** | **string**|  | 

### Return type

[**DictionaryGreetingInfo**](DictionaryGreetingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnsweringRule

> AnsweringRuleInfo UpdateAnsweringRule(ctx, accountId, extensionId, ruleId, body)
Update Call Handling Rule

Updates a custom answering rule for a particular caller ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**ruleId** | **string**| Internal identifier of an answering rule | 
**body** | [**UpdateAnsweringRuleRequest**](UpdateAnsweringRuleRequest.md)| JSON body | 

### Return type

[**AnsweringRuleInfo**](AnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallRecordingExtensionList

> UpdateCallRecordingExtensionList(ctx, accountId, body)
Update Call Recording Extension List

Creates or updates the list of extensions to be recorded.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
**body** | [**BulkAccountCallRecordingsResource**](BulkAccountCallRecordingsResource.md)|  | 

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


## UpdateCallRecordingSettings

> CallRecordingSettingsResource UpdateCallRecordingSettings(ctx, accountId, optional)
Update Call Recording Settings

Updates current call recording settings.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**|  | 
 **optional** | ***UpdateCallRecordingSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCallRecordingSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CallRecordingSettingsResource**](CallRecordingSettingsResource.md)|  | 

### Return type

[**CallRecordingSettingsResource**](CallRecordingSettingsResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyAnsweringRule

> CompanyAnsweringRuleInfo UpdateCompanyAnsweringRule(ctx, accountId, ruleId, body)
Update Company Call Handling Rule

Updates a company answering rule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**ruleId** | **string**| Internal identifier of an answering rule. The value can be standard digital ID or specific ID - either &#39;business-hours-rule&#39; or &#39;after-hours-rule&#39; | 
**body** | [**CompanyAnsweringRuleUpdate**](CompanyAnsweringRuleUpdate.md)| JSON body | 

### Return type

[**CompanyAnsweringRuleInfo**](CompanyAnsweringRuleInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

