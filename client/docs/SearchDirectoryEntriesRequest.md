# SearchDirectoryEntriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchString** | **string** | String value to filter the contacts. The value specified is searched through the following fields: &#x60;firstName&#x60;, &#x60;lastName&#x60;, &#x60;extensionNumber&#x60;, &#x60;phoneNumber&#x60;, &#x60;email&#x60; | [optional] 
**ShowFederated** | **bool** | If &#39;True&#39; then contacts of all accounts in federation are returned. If &#39;False&#39; then only contacts of the current account are returned, and account section is eliminated in this case | [optional] [default to true]
**ExtensionType** | **string** | Type of extension to filter the contacts | [optional] 
**OrderBy** | [**[]OrderBy**](OrderBy.md) | Sorting settings | [optional] 
**Page** | **int32** |  | [optional] 
**PerPage** | **int32** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


