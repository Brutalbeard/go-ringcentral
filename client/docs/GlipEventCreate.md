# GlipEventCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an event | [optional] 
**CreatorId** | **string** | Internal identifier of a person created an event | [optional] 
**Title** | **string** | Event title | 
**StartTime** | **string** | Datetime of starting an event | 
**EndTime** | **string** | Datetime of ending an event | 
**AllDay** | **bool** | Indicates whether event has some specific time slot or lasts for whole day(s) | [optional] [default to false]
**Recurrence** | **string** | Event recurrence settings. For non-periodic events the value is &#39;None&#39;. Must be greater or equal to event duration: 1- Day/Weekday; 7 - Week; 28 - Month; 365 - Year | [optional] 
**EndingCondition** | **string** | Condition of ending | [optional] 
**EndingAfter** | **int32** | Count of iterations. For periodic events only. Value range is 1 - 10. Must be specified if &#39;endingCondition&#39; is &#39;Count&#39; | [optional] 
**EndingOn** | **string** | Iterations end datetime for periodic events. | [optional] [default to ENDING_ON_NONE]
**Color** | **string** | Color of Event title (including its presentation in Calendar) | [optional] [default to COLOR_BLACK]
**Location** | **string** | Event location | [optional] 
**Description** | **string** | Event details | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


