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

type GlipPatchTeamBody struct {
	// Team access level
	Public bool `json:"public,omitempty"`
	// Team name. Maximum number of characters supported is 250
	Name string `json:"name,omitempty"`
	// Team description. Maximum number of characters supported is 1000
	Description string `json:"description,omitempty"`
}
