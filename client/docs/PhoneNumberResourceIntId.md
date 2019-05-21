# PhoneNumberResourceIntId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Internal identifier of a phone number | [optional] 
**Country** | [**CountryResource**](CountryResource.md) |  | [optional] 
**Extension** | [**PhoneNumberExtensionInfo**](PhoneNumberExtensionInfo.md) |  | [optional] 
**Label** | **string** | Custom user name of a phone number, if any | [optional] 
**Location** | **string** | Location (City, State). Filled for local US numbers | [optional] 
**PaymentType** | **string** | Payment type. &#39;External&#39; is returned for forwarded numbers which are not terminated in the RingCentral phone system &#x3D; [&#39;External&#39;, &#39;TollFree&#39;, &#39;Local&#39;], | [optional] 
**PhoneNumber** | **string** | Phone number | [optional] 
**Status** | **string** | Status of a phone number. If the value is &#39;Normal&#39;, the phone number is ready to be used. Otherwise it is an external number not yet ported to RingCentral , | [optional] 
**UsageType** | **string** |  | [optional] 
**Type** | **string** | Type of a phone number | [optional] 
**ReservationId** | **string** | Internal identifier of a phone number reservation; encoded data including reservation type (by brand, by account, by session), particular brand/account/session data, and reservation date and time | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


