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

type ListFaxCoverPagesResponse struct {
	Uri        string                  `json:"uri,omitempty"`
	Records    []FaxCoverPageInfo      `json:"records,omitempty"`
	Navigation MessagingNavigationInfo `json:"navigation,omitempty"`
	Paging     MessagingPagingInfo     `json:"paging,omitempty"`
}