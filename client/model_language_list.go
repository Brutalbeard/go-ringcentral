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

type LanguageList struct {
	// Canonical URI of the language list resource
	Uri string `json:"uri,omitempty"`
	// Language data
	Records    []LanguageInfo             `json:"records,omitempty"`
	Navigation ProvisioningNavigationInfo `json:"navigation,omitempty"`
	Paging     ProvisioningPagingInfo     `json:"paging,omitempty"`
}
