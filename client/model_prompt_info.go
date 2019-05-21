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

type PromptInfo struct {
	// Internal identifier of a prompt
	Uri string `json:"uri,omitempty"`
	// Link to a prompt metadata
	Id string `json:"id,omitempty"`
	// Type of a prompt media content
	ContentType string `json:"contentType,omitempty"`
	// Link to a prompt media content
	ContentUri string `json:"contentUri,omitempty"`
	// Name of a prompt
	Filename string `json:"filename,omitempty"`
}
