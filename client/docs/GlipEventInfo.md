# GlipEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an event | [optional] 
**CreatorId** | **string** | Internal identifier of a person created an event | [optional] 
**Title** | **string** | Event title | [optional] 
**StartTime** | **string** | Datetime of starting an event | [optional] 
**EndTime** | **string** | Datetime of ending an event | [optional] 
**AllDay** | **bool** | Indicates whether an event has some specific time slot or lasts for the whole day(s) | [optional] [default to false]
**Recurrence** | **string** | Event recurrence settings | [optional] 
**EndingCondition** | **string** | Condition of ending | [optional] 
**EndingAfter** | **int32** | Count of iterations. For periodic events only | [optional] 
**EndingOn** | **string** | Iterations end datetime for periodic events | [optional] [default to ENDING_ON_NONE]
**Color** | **string** | Color of Event title (including its presentation in Calendar) | [optional] [default to COLOR_BLACK]
**Location** | **string** | Event location | [optional] 
**Description** | **string** | Event details | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


