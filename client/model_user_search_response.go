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

type UserSearchResponse struct {
	// user list
	Resources    []UserResponse `json:"Resources,omitempty"`
	ItemsPerPage int64          `json:"itemsPerPage,omitempty"`
	Schemas      []string       `json:"schemas,omitempty"`
	StartIndex   int64          `json:"startIndex,omitempty"`
	TotalResults int64          `json:"totalResults,omitempty"`
}
