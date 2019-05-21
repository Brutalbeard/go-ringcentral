# ExtensionCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | [**ContactInfoCreationRequest**](ContactInfoCreationRequest.md) |  | 
**ExtensionNumber** | **string** | Number of extension | [optional] 
**Password** | **string** | Password for extension. If not specified, the password is auto-generated | [optional] 
**References** | [**[]ReferenceInfo**](ReferenceInfo.md) | List of non-RC internal identifiers assigned to an extension | [optional] 
**Roles** | [**[]Roles**](Roles.md) |  | [optional] 
**RegionalSettings** | [**RegionalSettings**](RegionalSettings.md) |  | [optional] 
**SetupWizardState** | **string** | Specifies extension configuration wizard state (web service setup). | [optional] [default to SETUP_WIZARD_STATE_NOT_STARTED]
**Status** | **string** | Extension current state | [optional] 
**StatusInfo** | [**ExtensionStatusInfo**](ExtensionStatusInfo.md) |  | [optional] 
**Type** | **string** | Extension type | 
**Hidden** | **bool** | Hides extension from showing in company directory. Supported for extensions of User type only. For unassigned extensions the value is set to &#39;True&#39; by default. For assigned extensions the value is set to &#39;False&#39; by default | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


