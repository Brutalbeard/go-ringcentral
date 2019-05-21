# CreateSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventFilters** | **[]string** | Collection of URIs to API resources. For APNS transport type only message event filter is available | 
**DeliveryMode** | [**NotificationDeliveryModeRequest**](NotificationDeliveryModeRequest.md) |  | 
**ExpiresIn** | **int32** | Subscription lifetime in seconds. Max value is 7 days (604800 sec). For *WebHook* transport type max value might be set up to 630720000 seconds (20 years) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


