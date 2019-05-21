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

type CallMonitoringGroups struct {
	// Link to a call monitoring groups resource
	Uri string `json:"uri"`
	// List of call monitoring group members
	Records    []CallMonitoringGroup      `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging     ProvisioningPagingInfo     `json:"paging"`
}