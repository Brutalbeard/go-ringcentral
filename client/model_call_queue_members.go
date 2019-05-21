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

type CallQueueMembers struct {
	// Link to a call queue members resource
	Uri string `json:"uri"`
	// List of call queue members
	Records    []CallQueueMemberInfo      `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}