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

type GlipPostTeamBody struct {
	// Team access level.
	Public bool `json:"public,omitempty"`
	// Team name.
	Name string `json:"name"`
	// Team description.
	Description string `json:"description,omitempty"`
	// Identifier(s) of team members.
	Members []CreateGlipConversationRequestMembers `json:"members,omitempty"`
}