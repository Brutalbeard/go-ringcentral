# \ExternalContactsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContact**](ExternalContactsApi.md#CreateContact) | **Post** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact | Create Contact
[**DeleteContact**](ExternalContactsApi.md#DeleteContact) | **Delete** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | Delete Contact
[**ListContacts**](ExternalContactsApi.md#ListContacts) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact | Get Contact List
[**ListFavoriteContacts**](ExternalContactsApi.md#ListFavoriteContacts) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/favorite | Get Favorite Contact List
[**ReadContact**](ExternalContactsApi.md#ReadContact) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | Get Contact
[**SyncAddressBook**](ExternalContactsApi.md#SyncAddressBook) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book-sync | Address Book Synchronization
[**UpdateContact**](ExternalContactsApi.md#UpdateContact) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId} | Update Contact
[**UpdateFavoriteContactList**](ExternalContactsApi.md#UpdateFavoriteContactList) | **Put** /restapi/v1.0/account/{accountId}/extension/{extensionId}/favorite | Update Favorite Contact List



## CreateContact

> PersonalContactResource CreateContact(ctx, accountId, extensionId, optional)
Create Contact

Creates personal user contact.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***CreateContactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateContactOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dialingPlan** | **optional.String**| A country code value complying with the [ISO 3166-1 alpha-2](https://ru.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. The default value is home country of the current extension | 
 **body** | [**optional.Interface of PersonalContactRequest**](PersonalContactRequest.md)|  | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContact

> DeleteContact(ctx, accountId, extensionId, contactId)
Delete Contact

Deletes contact(s) by ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 

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


## ListContacts

> ContactList ListContacts(ctx, accountId, extensionId, optional)
Get Contact List

Returns user personal contacts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startsWith** | **optional.String**| If specified, only contacts whose First name or Last name start with the mentioned substring are returned. Case-insensitive | 
 **sortBy** | [**optional.Interface of []string**](string.md)| Sorts results by the specified property | 
 **page** | **optional.Int32**| Indicates the page number to retrieve. Only positive number values are accepted | [default to 1]
 **perPage** | **optional.Int32**| Indicates the page size (number of items) | [default to 100]
 **phoneNumber** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**ContactList**](ContactList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFavoriteContacts

> ListFavoriteContacts(ctx, accountId, extensionId)
Get Favorite Contact List

Returns the list of favorite contacts of the current extension. Favorite contacts include both company contacts (extensions) and personal contacts (address book records).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]

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


## ReadContact

> PersonalContactResource ReadContact(ctx, accountId, extensionId, contactId)
Get Contact

Returns contact(s) by ID(s). Batch request is supported.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAddressBook

> AddressBookSync SyncAddressBook(ctx, accountId, extensionId, optional)
Address Book Synchronization

Synchronizes user contacts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***SyncAddressBookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SyncAddressBookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syncType** | [**optional.Interface of []string**](string.md)| Type of synchronization | 
 **syncToken** | **optional.String**| Value of syncToken property of the last sync request response | 
 **perPage** | **optional.Int32**| Number of records per page to be returned. The max number of records is 250, which is also the default. For &#39;FSync&#39; if the number of records exceeds the parameter value (either specified or default), all of the pages can be retrieved in several requests. For &#39;ISync&#39; if the number of records exceeds the page size, the number of incoming changes to this number is limited | 
 **pageId** | **optional.Int32**| Internal identifier of a page. It can be obtained from the &#39;nextPageId&#39; parameter passed in response body | 

### Return type

[**AddressBookSync**](AddressBookSync.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> PersonalContactResource UpdateContact(ctx, accountId, extensionId, contactId, optional)
Update Contact

Updates personal contact information by contact ID(s). Batch request is supported

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
**contactId** | **int32**| Internal identifier of a contact record in the RingCentral database | 
 **optional** | ***UpdateContactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateContactOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dialingPlan** | **optional.String**| A country code value complying with the [ISO 3166-1 alpha-2](https://ru.wikipedia.org/wiki/ISO_3166-1_alpha-2) format. The default value is home country of the current extension | 
 **body** | [**optional.Interface of PersonalContactRequest**](PersonalContactRequest.md)|  | 

### Return type

[**PersonalContactResource**](PersonalContactResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFavoriteContactList

> UpdateFavoriteContactList(ctx, accountId, extensionId, optional)
Update Favorite Contact List

Updates the list of favorite contacts of the current extension. Favorite contacts include both company contacts (extensions) and personal contacts (address book records).**Please note**: currently personal address book size is limited to 10 000 contacts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***UpdateFavoriteContactListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFavoriteContactListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of FavoriteCollection**](FavoriteCollection.md)|  | 

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

