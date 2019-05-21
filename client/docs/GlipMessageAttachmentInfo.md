# GlipMessageAttachmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an attachment | [optional] 
**Type** | **string** | Type of an attachment | [optional] [default to TYPE_CARD]
**Fallback** | **string** | A string of default text that will be rendered in the case that the client does not support Interactive Messages | [optional] 
**Intro** | **string** | A pretext to the message | [optional] 
**Author** | [**GlipMessageAttachmentAuthorInfo**](GlipMessageAttachmentAuthorInfo.md) |  | [optional] 
**Title** | **string** | Message title | [optional] 
**Text** | **string** | A large string field (up to 1000 chars) to be displayed as the body of a message (Supports GlipDown) | [optional] 
**ImageUri** | **string** | url used to display a single image at the bottom of a message | [optional] 
**ThumbnailUri** | **string** | url used to display a thumbnail to the right of a message (82x82) | [optional] 
**Fields** | [**[]GlipMessageAttachmentFieldsInfo**](GlipMessageAttachmentFieldsInfo.md) | Information on navigation | [optional] 
**Footnote** | [**GlipMessageAttachmentFootnoteInfo**](GlipMessageAttachmentFootnoteInfo.md) |  | [optional] 
**CreatorId** | **string** | Internal identifier of a person created an event | [optional] 
**StartTime** | **string** | Datetime of starting an event | [optional] 
**EndTime** | **string** | Datetime of ending an event | [optional] 
**AllDay** | **bool** | Indicates whether an event has some specific time slot or lasts for the whole day(s) | [optional] [default to false]
**Recurrence** | **string** | Event recurrence settings. | [optional] 
**EndingCondition** | **string** | Condition of ending | [optional] 
**EndingAfter** | **int32** | Count of iterations. For periodic events only | [optional] 
**EndingOn** | **string** | Iterations end datetime for periodic events | [optional] [default to ENDING_ON_NONE]
**Color** | **string** | Color of Event title, including its presentation in Calendar; or the color of the side border of an interactive message of a Card | [optional] [default to COLOR_BLACK]
**Location** | **string** | Event location | [optional] 
**Description** | **string** | Event details | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


