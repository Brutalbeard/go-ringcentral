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

type PresenceNavigationInfo struct {
	FirstPage    PresenceNavigationInfoUri `json:"firstPage,omitempty"`
	NextPage     PresenceNavigationInfoUri `json:"nextPage,omitempty"`
	PreviousPage PresenceNavigationInfoUri `json:"previousPage,omitempty"`
	LastPage     PresenceNavigationInfoUri `json:"lastPage,omitempty"`
}
