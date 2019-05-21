# GetAccountInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an account | [optional] 
**Uri** | **string** | Canonical URI of an account | [optional] 
**MainNumber** | **string** | Main phone number of the current account | [optional] 
**Operator** | [**GetExtensionInfoResponse**](GetExtensionInfoResponse.md) |  | [optional] 
**PartnerId** | **string** | Additional account identifier, developed and applied by the client | [optional] 
**ServiceInfo** | [**ServiceInfo**](ServiceInfo.md) |  | [optional] 
**SetupWizardState** | **string** | Specifies account configuration wizard state (web service setup). The default value is &#39;NotStarted&#39; | [optional] 
**Status** | **string** | Status of the current account | [optional] 
**StatusInfo** | [**AccountStatusInfo**](AccountStatusInfo.md) |  | [optional] 
**RegionalSettings** | [**RegionalSettings**](RegionalSettings.md) |  | [optional] 
**Federated** | **bool** | Specifies whether an account is included into any federation of accounts or not | [optional] 
**OutboundCallPrefix** | **int32** | If outbound call prefix is not specified, or set to null (0), then the parameter is not returned; the supported value range is 2-9 | [optional] 
**Cfid** | **string** | Customer facing identifier. Returned for accounts with the turned off PBX features. Equals to main company number in [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) (without &#39;+&#39; sign)format  | [optional] 
**Limits** | [**AccountLimits**](AccountLimits.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


