# GrantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Canonical URI of a grant | [optional] 
**Extension** | [**ExtensionInfoGrants**](ExtensionInfoGrants.md) |  | [optional] 
**CallPickup** | **bool** | Specifies if picking up of other extensions&#39; calls is allowed for the extension. If &#39;Presence&#39; feature is disabled for the given extension, the flag is not returned | [optional] 
**CallMonitoring** | **bool** | Specifies if monitoring of other extensions&#39; calls is allowed for the extension. If &#39;CallMonitoring&#39; feature is disabled for the given extension, the flag is not returned | [optional] 
**CallOnBehalfOf** | **bool** | Specifies whether the current extension is able to make or receive calls on behalf of the user referenced in extension object | [optional] 
**CallDelegation** | **bool** | Specifies whether the current extension can delegate a call to the user referenced in extension object | [optional] 
**GroupPaging** | **bool** | Specifies whether the current extension is allowed to call Paging Only group referenced to in extension object | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


