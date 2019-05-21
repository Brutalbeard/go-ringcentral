/*
 * RingCentral Connect Platform API
 *
 * RingCentral Connect Platform API
 *
 * API version: 1.0.36
 * Contact: platform@ringcentral.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ringcentral

type GlipMessageAttachmentInfo struct {
	// Internal identifier of an attachment
	Id string `json:"id,omitempty"`
	// Type of an attachment
	Type string `json:"type,omitempty"`
	// A string of default text that will be rendered in the case that the client does not support Interactive Messages
	Fallback string `json:"fallback,omitempty"`
	// A pretext to the message
	Intro  string                          `json:"intro,omitempty"`
	Author GlipMessageAttachmentAuthorInfo `json:"author,omitempty"`
	// Message title
	Title string `json:"title,omitempty"`
	// A large string field (up to 1000 chars) to be displayed as the body of a message (Supports GlipDown)
	Text string `json:"text,omitempty"`
	// url used to display a single image at the bottom of a message
	ImageUri string `json:"imageUri,omitempty"`
	// url used to display a thumbnail to the right of a message (82x82)
	ThumbnailUri string `json:"thumbnailUri,omitempty"`
	// Information on navigation
	Fields   []GlipMessageAttachmentFieldsInfo `json:"fields,omitempty"`
	Footnote GlipMessageAttachmentFootnoteInfo `json:"footnote,omitempty"`
	// Internal identifier of a person created an event
	CreatorId string `json:"creatorId,omitempty"`
	// Datetime of starting an event
	StartTime string `json:"startTime,omitempty"`
	// Datetime of ending an event
	EndTime string `json:"endTime,omitempty"`
	// Indicates whether an event has some specific time slot or lasts for the whole day(s)
	AllDay bool `json:"allDay,omitempty"`
	// Event recurrence settings.
	Recurrence string `json:"recurrence,omitempty"`
	// Condition of ending
	EndingCondition string `json:"endingCondition,omitempty"`
	// Count of iterations. For periodic events only
	EndingAfter int32 `json:"endingAfter,omitempty"`
	// Iterations end datetime for periodic events
	EndingOn string `json:"endingOn,omitempty"`
	// Color of Event title, including its presentation in Calendar; or the color of the side border of an interactive message of a Card
	Color string `json:"color,omitempty"`
	// Event location
	Location string `json:"location,omitempty"`
	// Event details
	Description string `json:"description,omitempty"`
}
