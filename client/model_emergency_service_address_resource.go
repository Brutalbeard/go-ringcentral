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

type EmergencyServiceAddressResource struct {
	Street       string `json:"street,omitempty"`
	Street2      string `json:"street2,omitempty"`
	City         string `json:"city,omitempty"`
	Zip          string `json:"zip,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	// State/province name
	State string `json:"state,omitempty"`
	// Internal identifier of a state
	StateId string `json:"stateId,omitempty"`
	// ISO code of a state
	StateIsoCode string `json:"stateIsoCode,omitempty"`
	// Full name of a state
	StateName string `json:"stateName,omitempty"`
	// Internal identifier of a country
	CountryId string `json:"countryId,omitempty"`
	// ISO code of a country
	CountryIsoCode string `json:"countryIsoCode,omitempty"`
	// Country name
	Country string `json:"country,omitempty"`
	// Full name of a country
	CountryName string `json:"countryName,omitempty"`
}
