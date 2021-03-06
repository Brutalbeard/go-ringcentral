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

type ConferencePhoneNumberInfo struct {
	// Dial-in phone number to connect to a conference
	PhoneNumber string `json:"phoneNumber,omitempty"`
	// 'True' if the number is default for the conference. Default conference number is a domestic number that can be set by user (otherwise it is set by the system). Only one default number per country is allowed
	Default bool `json:"default,omitempty"`
}
