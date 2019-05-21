# \SubscriptionsApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsApi.md#CreateSubscription) | **Post** /restapi/v1.0/subscription | Create Subscription
[**DeleteSubscription**](SubscriptionsApi.md#DeleteSubscription) | **Delete** /restapi/v1.0/subscription/{subscriptionId} | Cancel Subscription
[**ListSubscriptions**](SubscriptionsApi.md#ListSubscriptions) | **Get** /restapi/v1.0/subscription | Get Subscriptions
[**ReadSubscription**](SubscriptionsApi.md#ReadSubscription) | **Get** /restapi/v1.0/subscription/{subscriptionId} | Get Subscription
[**RenewSubscription**](SubscriptionsApi.md#RenewSubscription) | **Post** /restapi/v1.0/subscription/{subscriptionId}/renew | Renew Subscription
[**UpdateSubscription**](SubscriptionsApi.md#UpdateSubscription) | **Put** /restapi/v1.0/subscription/{subscriptionId} | Renew Subscription / Update Event Filters



## CreateSubscription

> SubscriptionInfo CreateSubscription(ctx, body)
Create Subscription

Creates a new subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CreateSubscriptionRequest**](CreateSubscriptionRequest.md)| JSON body | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, subscriptionId)
Cancel Subscription

Cancels the existent subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| Internal identifier of a subscription | 

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


## ListSubscriptions

> RecordsCollectionResourceSubscriptionResponse ListSubscriptions(ctx, )
Get Subscriptions

Returns a list of subscriptions created by a particular user on a particular client app.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**RecordsCollectionResourceSubscriptionResponse**](RecordsCollectionResourceSubscriptionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadSubscription

> SubscriptionInfo ReadSubscription(ctx, subscriptionId)
Get Subscription

Returns the requested subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **int32**| Internal identifier of a subscription | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewSubscription

> SubscriptionInfo RenewSubscription(ctx, subscriptionId)
Renew Subscription

Renews an existent subscription by ID by posting request with an empty body.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**|  | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> SubscriptionInfo UpdateSubscription(ctx, subscriptionId, body, optional)
Renew Subscription / Update Event Filters

Renews the existent subscription if the request body is empty. If event filters are specified, calling this method modifies the event filters for the existing subscription. The client application can extend or narrow the events for which it receives notifications in the frame of one subscription.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string**| Internal identifier of a subscription | 
**body** | [**ModifySubscriptionRequest**](ModifySubscriptionRequest.md)| JSON body | 
 **optional** | ***UpdateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSubscriptionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aggregated** | **optional.Bool**| If &#39;True&#39; then aggregated presence status is returned in a notification payload | 

### Return type

[**SubscriptionInfo**](SubscriptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

