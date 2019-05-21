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

type AccountStatusInfo struct {
	// A free-form user comment, describing the status change reason
	Comment string `json:"comment,omitempty"`
	// Type of suspension
	Reason string `json:"reason,omitempty"`
	// Date until which the account will get deleted. The default value is 30 days since current date
	Till string `json:"till,omitempty"`
}
