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

type ExtensionInfoEventBody struct {
	// Internal identifier of an extension
	ExtensionId string `json:"extensionId,omitempty"`
	// Type of extension info change
	EventType string `json:"eventType,omitempty"`
	// Returned for 'Update' event type only. The possible values are: /nAccountInfo - change of account parameters/nExtensionInfo - change of contact info, service features, departments, status/nPhoneNumber - change of phone numbers/nRole - change of permissions/nProfileImage - change of profile image
	Hints []string `json:"hints,omitempty"`
	// Internal identifier of a subscription owner extension
	OwnerId string `json:"ownerId,omitempty"`
}
