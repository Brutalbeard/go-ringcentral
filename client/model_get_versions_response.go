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

type GetVersionsResponse struct {
	// Canonical URI of the API version
	Uri string `json:"uri,omitempty"`
	// Full API version information: uri, number, release date
	ApiVersions []VersionInfo `json:"apiVersions,omitempty"`
	// Server version
	ServerVersion string `json:"serverVersion,omitempty"`
	// Server revision
	ServerRevision string `json:"serverRevision,omitempty"`
}
