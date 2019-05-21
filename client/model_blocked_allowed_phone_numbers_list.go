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

// List of blocked or allowed phone numbers
type BlockedAllowedPhoneNumbersList struct {
	// Link to a list of blocked/allowed phone numbers resource
	Uri        string                          `json:"uri,omitempty"`
	Records    []BlockedAllowedPhoneNumberInfo `json:"records,omitempty"`
	Navigation CallHandlingNavigationInfo      `json:"navigation,omitempty"`
	Paging     CallHandlingPagingInfo          `json:"paging,omitempty"`
}