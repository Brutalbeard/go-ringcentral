# CreateDataExportTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFrom** | **string** | Starting time for data collection. The default value is current time minus 24 hours. If &#x60;dateTo&#x60; is not specified then its value is set to &#x60;dateFrom&#x60; plus 24 hours. The specified time range should not be greater than 7 days | [optional] 
**DateTo** | **string** | Ending time for data collection. The default value is current time. If &#x60;dateFrom&#x60; is not specified then its value is set to &#x60;dateTo&#x60; minus 24 hours. The specified time range should not be greater than 7 days | [optional] 
**UserIds** | **[]string** | List of users which data is collected. The following data will be exported: posts, tasks, events, etc. posted by the user(s); posts addressing the user(s) via direct and @Mentions; tasks assigned to the listed user(s). The list of 30 users per request is supported. | [optional] 
**ChatIds** | **[]string** | List of chats from which the data (posts, files, tasks, events, notes, etc.) will be collected | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


