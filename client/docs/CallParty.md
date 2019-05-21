# CallParty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a party | [optional] 
**Status** | [**CallStatusInfo**](CallStatusInfo.md) |  | [optional] 
**Muted** | **bool** | Specifies if a call participant is muted or not | [optional] 
**StandAlone** | **bool** | True if party is not connected to a session voice conference. False - otherwise | [optional] 
**Park** | [**ParkInfo**](ParkInfo.md) |  | [optional] 
**From** | [**PartyInfo**](PartyInfo.md) |  | [optional] 
**To** | [**PartyInfo**](PartyInfo.md) |  | [optional] 
**Owner** | [**OwnerInfo**](OwnerInfo.md) |  | [optional] 
**Direction** | **string** | Direction of a call | [optional] 
**ConferenceRole** | **string** | A party&#39;s role in the conference scenarios. For calls of &#39;Conference&#39; type only | [optional] 
**RingOutRole** | **string** | A party&#39;s role in &#39;Ring Me&#39;/&#39;RingOut&#39; scenarios. For calls of &#39;Ringout&#39; type only | [optional] 
**RingMeRole** | **string** | A party&#39;s role in &#39;Ring Me&#39;/&#39;RingOut&#39; scenarios. For calls of &#39;Ringme&#39; type only | [optional] 
**Recordings** | [**[]RecordingInfo**](RecordingInfo.md) | Active recordings list | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


