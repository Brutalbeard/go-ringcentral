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

type CustomUserGreetingInfo struct {
	// Link to a custom user greeting
	Uri string `json:"uri,omitempty"`
	// Internal identifier of a custom user greeting
	Id string `json:"id,omitempty"`
	// Type of a custom user greeting
	Type string `json:"type,omitempty"`
	// Content media type in WAV/MP3 format
	ContentType string `json:"contentType,omitempty"`
	// Link to a greeting content (audio file)
	ContentUri    string                          `json:"contentUri,omitempty"`
	AnsweringRule CustomGreetingAnsweringRuleInfo `json:"answeringRule,omitempty"`
}
