# \SIPApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSIPRegistration**](SIPApi.md#CreateSIPRegistration) | **Post** /restapi/v1.0/client-info/sip-provision | Register SIP Device



## CreateSIPRegistration

> CreateSipRegistrationResponse CreateSIPRegistration(ctx, body)
Register SIP Device

Creates SIP registration of a device/application (WebPhone, Mobile, softphone)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateSipRegistrationRequest**](CreateSipRegistrationRequest.md)| JSON body | 

### Return type

[**CreateSipRegistrationResponse**](CreateSipRegistrationResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

