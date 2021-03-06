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

type GetCountryInfoDictionaryResponse struct {
	// Internal identifier of a country
	Id string `json:"id,omitempty"`
	// Canonical URI of a country
	Uri string `json:"uri,omitempty"`
	// Country calling code defined by ITU-T recommendations [E.123](https://www.itu.int/rec/T-REC-E.123-200102-I/en) and [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I)
	CallingCode string `json:"callingCode,omitempty"`
	// Emergency calling feature availability/emergency address requirement indicator
	EmergencyCalling bool `json:"emergencyCalling,omitempty"`
	// Country code according to the ISO standard, see [ISO 3166](https://www.iso.org/iso-3166-country-codes.html)
	IsoCode string `json:"isoCode,omitempty"`
	// Official name of a country
	Name string `json:"name,omitempty"`
	// Determines whether phone numbers are available for a country
	NumberSelling bool `json:"numberSelling,omitempty"`
	// Specifies whether login with the phone numbers of this country is enabled or not
	LoginAllowed bool `json:"loginAllowed,omitempty"`
	// Indicates whether signup/billing is allowed for a country
	SignupAllowed bool `json:"signupAllowed,omitempty"`
	// Specifies if free phone line for softphone is available for a country or not
	FreeSoftphoneLine bool `json:"freeSoftphoneLine,omitempty"`
}
