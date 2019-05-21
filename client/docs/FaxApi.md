# \FaxApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFaxMessage**](FaxApi.md#CreateFaxMessage) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/fax | Create Fax Message
[**ListFaxCoverPages**](FaxApi.md#ListFaxCoverPages) | **Get** /restapi/v1.0/dictionary/fax-cover-page | Get Fax Cover Page List



## CreateFaxMessage

> FaxResponse CreateFaxMessage(ctx, accountId, extensionId, attachment, to, optional)
Create Fax Message

Creates and sends/resends a new fax message. Resend can be done if sending failed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account (integer) or tilde (~) to indicate the account which was logged-in within the current session. | 
**extensionId** | **string**| Internal identifier of an extension (integer) or tilde (~) to indicate the extension assigned to the account logged-in within the current session | 
**attachment** | **[]*os.File**| File to upload | 
**to** | [**[]string**](string.md)| To Phone Number | 
 **optional** | ***CreateFaxMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFaxMessageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **faxResolution** | **optional.String**| Resolution of Fax | 
 **sendTime** | **optional.Time**| Timestamp to send fax at. If not specified (current or the past), the fax is sent immediately | 
 **isoCode** | **optional.String**| ISO Code. e.g UK | 
 **coverIndex** | **optional.Int32**| Cover page identifier. For the list of available cover page identifiers please call the method Fax Cover Pages. If not specified, the default cover page which is configured in &#39;Outbound Fax Settings&#39; is attached | 
 **coverPageText** | **optional.String**| Cover page text, entered by the fax sender and printed on the cover page. Maximum length is limited to 1024 symbols | 

### Return type

[**FaxResponse**](FaxResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFaxCoverPages

> ListFaxCoverPagesResponse ListFaxCoverPages(ctx, optional)
Get Fax Cover Page List

Returns fax cover pages available for the current extension.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFaxCoverPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFaxCoverPagesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]

### Return type

[**ListFaxCoverPagesResponse**](ListFaxCoverPagesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

