# SipInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | User credentials | [optional] 
**Password** | **string** | User password | [optional] 
**AuthorizationId** | **string** | Identifier for SIP authorization | [optional] 
**Domain** | **string** | SIP domain | [optional] 
**OutboundProxy** | **string** | SIP outbound proxy | [optional] 
**Transport** | **string** | Preferred transport. SIP info will be returned for this transport if supported | [optional] 
**Certificate** | **string** | For TLS transport only Base64 encoded certificate | [optional] 
**SwitchBackInterval** | **int32** | The interval in seconds after which the app must try to switch back to primary proxy if it was previously switched to backup. If this parameter is not returned, the app must stay on backup proxy and try to switch to primary proxy after the next SIP-provision call. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


