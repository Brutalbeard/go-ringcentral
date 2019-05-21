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

type CompanyBusinessHours struct {
	// Canonical URI of a business-hours resource
	Uri      string                           `json:"uri,omitempty"`
	Schedule CompanyBusinessHoursScheduleInfo `json:"schedule,omitempty"`
}
