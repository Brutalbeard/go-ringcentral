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

type AccountPresenceEventBody struct {
	// Internal identifier of an extension. Optional parameter
	ExtensionId string `json:"extensionId,omitempty"`
	// Telephony presence status. Returned if telephony status is changed.
	TelephonyStatus string `json:"telephonyStatus,omitempty"`
	// Order number of a notification to state the chronology
	Sequence int32 `json:"sequence,omitempty"`
	// Aggregated presence status, calculated from a number of sources
	PresenceStatus string `json:"presenceStatus,omitempty"`
	// User-defined presence status (as previously published by the user)
	UserStatus string `json:"userStatus,omitempty"`
	// Extended DnD (Do not Disturb) status
	DndStatus string `json:"dndStatus,omitempty"`
	// If 'True' enables other extensions to see the extension presence status
	AllowSeeMyPresence bool `json:"allowSeeMyPresence,omitempty"`
	// If 'True' enables to ring extension phone, if any user monitored by this extension is ringing
	RingOnMonitoredCall bool `json:"ringOnMonitoredCall,omitempty"`
	// If 'True' enables the extension user to pick up a monitored line on hold
	PickUpCallsOnHold bool `json:"pickUpCallsOnHold,omitempty"`
	// Internal identifier of a subscription owner extension
	OwnerId string `json:"ownerId,omitempty"`
}
