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

type ExtensionInfo struct {
	// Internal identifier of an extension
	Id string `json:"id,omitempty"`
	// Canonical URI of an extension
	Uri string `json:"uri,omitempty"`
	// Number of department extension
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	// For Partner Applications Internal identifier of an extension created by partner. The RingCentral supports the mapping of accounts and stores the corresponding account ID/extension ID for each partner ID of a client application. In request URIs partner IDs are accepted instead of regular RingCentral native IDs as path parameters using pid = XXX clause. Though in response URIs contain the corresponding account IDs and extension IDs. In all request and response bodies these values are reflected via partnerId attributes of account and extension
	PartnerId string `json:"partnerId,omitempty"`
}
