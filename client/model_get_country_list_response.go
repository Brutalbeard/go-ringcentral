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

type GetCountryListResponse struct {
	// List of countries with the country data
	Records    []GetCountryInfoDictionaryResponse `json:"records"`
	Navigation ProvisioningNavigationInfo         `json:"navigation"`
	Paging     ProvisioningPagingInfo             `json:"paging"`
}
