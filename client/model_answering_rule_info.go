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

type AnsweringRuleInfo struct {
	// Canonical URI to an answering rule resource
	Uri string `json:"uri,omitempty"`
	// Internal identifier of an answering rule
	Id string `json:"id,omitempty"`
	// Type of an answering rule
	Type string `json:"type,omitempty"`
	// Name of an answering rule specified by user
	Name string `json:"name,omitempty"`
	// Specifies if an answering rule is active or inactive
	Enabled  bool         `json:"enabled,omitempty"`
	Schedule ScheduleInfo `json:"schedule,omitempty"`
	// Answering rules are applied when calling to selected number(s)
	CalledNumbers []CalledNumberInfo `json:"calledNumbers,omitempty"`
	// Answering rules are applied when calls are received from specified caller(s)
	Callers []CallersInfo `json:"callers,omitempty"`
	// Specifies how incoming calls are forwarded
	CallHandlingAction      string                      `json:"callHandlingAction,omitempty"`
	Forwarding              ForwardingInfo              `json:"forwarding,omitempty"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding,omitempty"`
	Queue                   QueueInfo                   `json:"queue,omitempty"`
	Transfer                TransferredExtensionInfo    `json:"transfer,omitempty"`
	Voicemail               VoicemailInfo               `json:"voicemail,omitempty"`
	// Greetings applied for an answering rule; only predefined greetings can be applied, see Dictionary Greeting List
	Greetings []GreetingInfo `json:"greetings,omitempty"`
	// Call screening status. 'Off' - no call screening; 'NoCallerId' - if caller ID is missing, then callers are asked to say their name before connecting; 'UnknownCallerId' - if caller ID is not in contact list, then callers are asked to say their name before connecting; 'Always' - the callers are always asked to say their name before connecting. The default value is 'Off'
	Screening string `json:"screening,omitempty"`
}
