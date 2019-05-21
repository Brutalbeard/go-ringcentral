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

type ActiveCallsRecordExtensionInfo struct {
	// Link to an extension on whose behalf a call is initiated
	Uri string `json:"uri,omitempty"`
	// Internal identifier of an extension on whose behalf a call is initiated
	Id string `json:"id,omitempty"`
}
