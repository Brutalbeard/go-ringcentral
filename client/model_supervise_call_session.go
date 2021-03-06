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

type SuperviseCallSession struct {
	From PartyInfo `json:"from,omitempty"`
	To   PartyInfo `json:"to,omitempty"`
	// Direction of a call
	Direction string `json:"direction,omitempty"`
	// Internal identifier of a party that monitors a call
	Id string `json:"id,omitempty"`
	// Specifies if a call party is muted
	Muted bool      `json:"muted,omitempty"`
	Owner OwnerInfo `json:"owner,omitempty"`
	// Specifies if a device is stand-alone
	StandAlone bool           `json:"standAlone,omitempty"`
	Status     CallStatusInfo `json:"status,omitempty"`
}
