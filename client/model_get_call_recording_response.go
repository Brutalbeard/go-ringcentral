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

type GetCallRecordingResponse struct {
	// Internal identifier of a call recording
	Id string `json:"id,omitempty"`
	// Link to a call recording binary content
	ContentUri string `json:"contentUri,omitempty"`
	// Call recording file format. Supported format is audio/x-wav
	ContentType string `json:"contentType,omitempty"`
	// Recorded call duration
	Duration int32 `json:"duration,omitempty"`
}
