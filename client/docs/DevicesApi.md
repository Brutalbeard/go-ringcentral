# \DevicesApi

All URIs are relative to *https://platform.ringcentral.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExtensionDevices**](DevicesApi.md#ListExtensionDevices) | **Get** /restapi/v1.0/account/{accountId}/extension/{extensionId}/device | Get Extension Device List
[**ReadDevice**](DevicesApi.md#ReadDevice) | **Get** /restapi/v1.0/account/{accountId}/device/{deviceId} | Get Device
[**UpdateDevice**](DevicesApi.md#UpdateDevice) | **Put** /restapi/v1.0/account/{accountId}/device/{deviceId} | Update Device



## ListExtensionDevices

> GetExtensionDevicesResponse ListExtensionDevices(ctx, accountId, extensionId, optional)
Get Extension Device List

Returns devices of the extension(s) by their ID(s). Batch request is supported

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**extensionId** | **string**| Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session | [default to ~]
 **optional** | ***ListExtensionDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExtensionDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linePooling** | **optional.String**| Pooling type of a device | 
 **feature** | **optional.String**| Device feature or multiple features supported | 

### Return type

[**GetExtensionDevicesResponse**](GetExtensionDevicesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDevice

> GetDeviceInfoResponse ReadDevice(ctx, accountId, deviceId, optional)
Get Device

Returns account device(s) by their ID(s).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string**| Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session | [default to ~]
**deviceId** | **int32**| Internal identifier of a device | 
 **optional** | ***ReadDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReadDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syncEmergencyAddress** | **optional.Bool**| Specifies if emergency address should be synchronized or not | [default to false]

### Return type

[**GetDeviceInfoResponse**](GetDeviceInfoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> DeviceResource UpdateDevice(ctx, deviceId, accountId, optional)
Update Device

Updates account device(s) by their ID(s).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**|  | 
**accountId** | **string**|  | 
 **optional** | ***UpdateDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AccountDeviceUpdate**](AccountDeviceUpdate.md)|  | 

### Return type

[**DeviceResource**](DeviceResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

