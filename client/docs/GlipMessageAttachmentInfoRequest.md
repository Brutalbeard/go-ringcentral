# GlipMessageAttachmentInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of attachment | [optional] [default to TYPE_CARD]
**Title** | **string** | Attachment title | [optional] 
**Fallback** | **string** | Default message returned in case the client does not support interactive messages | [optional] 
**Color** | **string** | Color of Event title, including its presentation in Calendar; or the color of the side border of an interactive message of a Card. The default color is &#39;Black&#39; | [optional] 
**Intro** | **string** | Introductory text displayed directly above a message | [optional] 
**Author** | [**GlipMessageAttachmentAuthorInfo**](GlipMessageAttachmentAuthorInfo.md) |  | [optional] 
**Text** | **string** | Text of attachment (up to 1000 chars), supports GlipDown | [optional] 
**ImageUri** | **string** | Link to an image displayed at the bottom of a message | [optional] 
**ThumbnailUri** | **string** | Link to an image preview displayed to the right of a message (82x82) | [optional] 
**Fields** | [**[]GlipMessageAttachmentFieldsInfo**](GlipMessageAttachmentFieldsInfo.md) | Individual subsections within a message | [optional] 
**Footnote** | [**GlipMessageAttachmentFootnoteInfo**](GlipMessageAttachmentFootnoteInfo.md) |  | [optional] 
**StartTime** | **string** | Datetime of starting an event | [optional] 
**EndTime** | **string** | Datetime of ending an event | [optional] 
**AllDay** | **bool** | Indicates whether an event has some specific time slot or lasts for the whole day(s) | [optional] [default to false]
**Recurrence** | **string** | Event recurrence settings. For non-periodic events the value is &#39;None&#39;. Must be greater or equal to event duration: 1- Day/Weekday; 7 - Week; 28 - Month; 365 - Year | [optional] 
**EndingCondition** | **string** | Condition of ending an event | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


