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

type GlipEventInfo struct {
	// Internal identifier of an event
	Id string `json:"id,omitempty"`
	// Internal identifier of a person created an event
	CreatorId string `json:"creatorId,omitempty"`
	// Event title
	Title string `json:"title,omitempty"`
	// Datetime of starting an event
	StartTime string `json:"startTime,omitempty"`
	// Datetime of ending an event
	EndTime string `json:"endTime,omitempty"`
	// Indicates whether an event has some specific time slot or lasts for the whole day(s)
	AllDay bool `json:"allDay,omitempty"`
	// Event recurrence settings
	Recurrence string `json:"recurrence,omitempty"`
	// Condition of ending
	EndingCondition string `json:"endingCondition,omitempty"`
	// Count of iterations. For periodic events only
	EndingAfter int32 `json:"endingAfter,omitempty"`
	// Iterations end datetime for periodic events
	EndingOn string `json:"endingOn,omitempty"`
	// Color of Event title (including its presentation in Calendar)
	Color string `json:"color,omitempty"`
	// Event location
	Location string `json:"location,omitempty"`
	// Event details
	Description string `json:"description,omitempty"`
}