# ContactInfoCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | For User extension type only. Extension user first name | 
**LastName** | **string** | For User extension type only. Extension user last name | 
**Company** | **string** | Extension user company name | [optional] 
**JobTitle** | **string** |  | [optional] 
**Email** | **string** | Email of extension user | 
**BusinessPhone** | **string** | Extension user contact phone number in [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) format | [optional] 
**MobilePhone** | **string** | Extension user mobile (**non** Toll Free) phone number in [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) (with &#39;+&#39; sign) format | [optional] 
**BusinessAddress** | [**ContactBusinessAddressInfo**](ContactBusinessAddressInfo.md) |  | [optional] 
**EmailAsLoginName** | **bool** |  If &#39;True&#39; then contact email is enabled as login name for this user. Please note that email should be unique in this case. The default value is &#39;False&#39; | [optional] 
**PronouncedName** | [**PronouncedNameInfo**](PronouncedNameInfo.md) |  | [optional] 
**Department** | **string** | Extension user department, if any | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


