# SubscriptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a subscription | [optional] 
**Uri** | **string** | Canonical URI of a subscription | [optional] 
**EventFilters** | **[]string** | Collection of API resources (message-store/presence/detailed presence) corresponding to events the user is subscribed to | [optional] 
**DisabledFilters** | [**[]DisabledFilterInfo**](DisabledFilterInfo.md) | Collection of API resources (message-store/presence/detailed presence) corresponding to events the user is not subscribed to due to certain limitations | [optional] 
**ExpirationTime** | [**time.Time**](time.Time.md) | Subscription expiration datetime in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**ExpiresIn** | **int32** | Subscription lifetime in seconds | [optional] 
**Status** | **string** | Subscription status | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | Subscription creation datetime in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format including timezone, for example 2016-03-10T18:07:52.534Z | [optional] 
**DeliveryMode** | [**NotificationDeliveryMode**](NotificationDeliveryMode.md) |  | 
**BlacklistedData** | [**NotificationBlacklistedData**](NotificationBlacklistedData.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


