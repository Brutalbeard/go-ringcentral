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

type NotificationSettings struct {
	// Canonical URI of notifications settings resource
	Uri string `json:"uri,omitempty"`
	// List of notification recipient email addresses
	EmailAddresses []string `json:"emailAddresses,omitempty"`
	// List of notification recipient email addresses
	SmsEmailAddresses []string `json:"smsEmailAddresses,omitempty"`
	// Specifies notifications settings mode. If 'True' then advanced mode is on, it allows using different emails and/or phone numbers for each notification type. If 'False' then basic mode is on. Advanced mode settings are returned in both modes, if specified once, but if basic mode is switched on, they are not applied
	AdvancedMode  bool              `json:"advancedMode,omitempty"`
	Voicemails    VoicemailsInfo    `json:"voicemails,omitempty"`
	InboundFaxes  InboundFaxesInfo  `json:"inboundFaxes,omitempty"`
	OutboundFaxes OutboundFaxesInfo `json:"outboundFaxes,omitempty"`
	InboundTexts  InboundTextsInfo  `json:"inboundTexts,omitempty"`
	MissedCalls   MissedCallsInfo   `json:"missedCalls,omitempty"`
}
