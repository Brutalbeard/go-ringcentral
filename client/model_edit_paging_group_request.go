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

type EditPagingGroupRequest struct {
	// List of users that will be allowed to page a group specified
	AddedUserIds []string `json:"addedUserIds,omitempty"`
	// List of users that will be unallowed to page a group specified
	RemovedUserIds []string `json:"removedUserIds,omitempty"`
	// List of account devices that will be assigned to a paging group specified
	AddedDeviceIds []string `json:"addedDeviceIds,omitempty"`
	// List of account devices that will be unassigned from a paging group specified
	RemovedDeviceIds []string `json:"removedDeviceIds,omitempty"`
}
