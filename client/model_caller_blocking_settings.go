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

// Returns the lists of blocked and allowed phone numbers
type CallerBlockingSettings struct {
	// Call blocking options: either specific or all calls and faxes
	Mode string `json:"mode,omitempty"`
	//  Determines how to handle calls with no caller ID in 'Specific' mode
	NoCallerId string `json:"noCallerId,omitempty"`
	// Blocking settings for pay phones
	PayPhones string `json:"payPhones,omitempty"`
	// List of greetings played for blocked callers
	Greetings []BlockedCallerGreetingInfo `json:"greetings,omitempty"`
}