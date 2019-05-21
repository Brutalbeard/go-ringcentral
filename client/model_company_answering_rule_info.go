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

type CompanyAnsweringRuleInfo struct {
	// Internal identifier of an answering rule
	Id string `json:"id,omitempty"`
	// Canonical URI of an answering rule
	Uri string `json:"uri,omitempty"`
	// Specifies if the rule is active or inactive. The default value is 'True'
	Enabled bool `json:"enabled,omitempty"`
	// Type of an answering rule
	Type string `json:"type,omitempty"`
	// Name of an answering rule specified by user. Max number of symbols is 30. The default value is 'My Rule N' where 'N' is the first free number
	Name string `json:"name,omitempty"`
	// Answering rule will be applied when calls are received from the specified caller(s)
	Callers []CompanyAnsweringRuleCallersInfoRequest `json:"callers,omitempty"`
	// Answering rule will be applied when calling the specified number(s)
	CalledNumbers []CompanyAnsweringRuleCalledNumberInfoRequest `json:"calledNumbers,omitempty"`
	Schedule      CompanyAnsweringRuleScheduleInfo              `json:"schedule,omitempty"`
	// Specifies how incoming calls are forwarded. The default value is 'Operator' 'Operator' - play company greeting and forward to operator extension 'Disconnect' - play company greeting and disconnect 'Bypass' - bypass greeting to go to selected extension = ['Operator', 'Disconnect', 'Bypass']
	CallHandlingAction string                                 `json:"callHandlingAction,omitempty"`
	Extension          CompanyAnsweringRuleCallersInfoRequest `json:"extension,omitempty"`
	// Greetings applied for an answering rule; only predefined greetings can be applied, see Dictionary Greeting List
	Greetings []GreetingInfo `json:"greetings,omitempty"`
}
