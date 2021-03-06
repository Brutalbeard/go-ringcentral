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

type ActivePermissionResource struct {
	Permission    PermissionIdResource `json:"permission,omitempty"`
	EffectiveRole RoleIdResource       `json:"effectiveRole,omitempty"`
	Scopes        []string             `json:"scopes,omitempty"`
}
