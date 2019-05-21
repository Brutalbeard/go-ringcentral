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

type VersionInfo struct {
	// Canonical URI of API versions
	Uri string `json:"uri,omitempty"`
	// Version of the RingCentral REST API
	VersionString string `json:"versionString,omitempty"`
	// Release date of this version
	ReleaseDate string `json:"releaseDate,omitempty"`
	// URI part determining the current version
	UriString string `json:"uriString,omitempty"`
}
