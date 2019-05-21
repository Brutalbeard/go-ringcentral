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

type MessageEvent struct {
	// Universally unique identifier of a notification
	Uuid string `json:"uuid,omitempty"`
	// Event filter URI
	Event string `json:"event,omitempty"`
	// Datetime of sending a notification in [ISO 8601](shttps://en.wikipedia.org/wiki/ISO_8601) format including timezone, for example *2016-03-10T18:07:52.534Z*
	Timestamp string `json:"timestamp,omitempty"`
	// Internal identifier of a subscription
	SubscriptionId string           `json:"subscriptionId,omitempty"`
	Body           MessageEventBody `json:"body,omitempty"`
}
