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

type CallParty struct {
	// Internal identifier of a party
	Id     string         `json:"id,omitempty"`
	Status CallStatusInfo `json:"status,omitempty"`
	// Specifies if a call participant is muted or not
	Muted bool `json:"muted,omitempty"`
	// True if party is not connected to a session voice conference. False - otherwise
	StandAlone bool      `json:"standAlone,omitempty"`
	Park       ParkInfo  `json:"park,omitempty"`
	From       PartyInfo `json:"from,omitempty"`
	To         PartyInfo `json:"to,omitempty"`
	Owner      OwnerInfo `json:"owner,omitempty"`
	// Direction of a call
	Direction string `json:"direction,omitempty"`
	// A party's role in the conference scenarios. For calls of 'Conference' type only
	ConferenceRole string `json:"conferenceRole,omitempty"`
	// A party's role in 'Ring Me'/'RingOut' scenarios. For calls of 'Ringout' type only
	RingOutRole string `json:"ringOutRole,omitempty"`
	// A party's role in 'Ring Me'/'RingOut' scenarios. For calls of 'Ringme' type only
	RingMeRole string `json:"ringMeRole,omitempty"`
	// Active recordings list
	Recordings []RecordingInfo `json:"recordings,omitempty"`
}
