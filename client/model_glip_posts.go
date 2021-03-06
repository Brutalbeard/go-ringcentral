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

type GlipPosts struct {
	// List of posts
	Records    []GlipPostInfo     `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation,omitempty"`
}
