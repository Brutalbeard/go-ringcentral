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

type GetCountryInfoNumberParser struct {
	// Internal identifier of a country
	Id string `json:"id,omitempty"`
	// Canonical URI of a country
	Uri string `json:"uri,omitempty"`
	// Country calling code defined by ITU-T recommendations E.123 and E.164, see Calling Codes
	CallingCode string `json:"callingCode,omitempty"`
	// Emergency calling feature availability/emergency address requirement indicator
	EmergencyCalling bool `json:"emergencyCalling,omitempty"`
	// Country code according to the ISO standard, see ISO 3166
	IsoCode string `json:"isoCode,omitempty"`
	// Official name of a country
	Name string `json:"name,omitempty"`
}
