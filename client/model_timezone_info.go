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

type TimezoneInfo struct {
	// Internal identifier of a timezone
	Id string `json:"id,omitempty"`
	// Canonical URI of a timezone
	Uri string `json:"uri,omitempty"`
	// Short name of a timezone
	Name string `json:"name,omitempty"`
	// Meaningful description of the timezone
	Description string `json:"description,omitempty"`
}
