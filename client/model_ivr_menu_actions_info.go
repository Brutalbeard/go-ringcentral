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

type IvrMenuActionsInfo struct {
	// Key. The following values are supported: numeric: '1' to '9' Star Hash NoInput
	Input string `json:"input,omitempty"`
	// Internal identifier of an answering rule
	Action    string               `json:"action,omitempty"`
	Extension IvrMenuExtensionInfo `json:"extension,omitempty"`
	// For 'Transfer' action only. PSTN number in E.164 format
	PhoneNumber string `json:"phoneNumber,omitempty"`
}
