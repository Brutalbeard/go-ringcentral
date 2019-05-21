# GlipWebhookInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a webhook | [optional] 
**CreatorId** | **string** | Internal identifier of the user who created a webhook | [optional] 
**GroupsId** | **[]string** | Internal identifiers of groups where a webhook has been created | [optional] 
**CreationTime** | [**time.Time**](time.Time.md) | Webhook creation time in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format | [optional] 
**LastModifiedTime** | [**time.Time**](time.Time.md) | Webhook last update time in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format | [optional] 
**Uri** | **string** | Public link to send a webhook payload | [optional] 
**Status** | **string** | Current status of a webhook | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


