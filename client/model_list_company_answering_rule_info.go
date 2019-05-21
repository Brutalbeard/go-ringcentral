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

type ListCompanyAnsweringRuleInfo struct {
	// Internal identifier of an answering rule
	Id string `json:"id,omitempty"`
	// Canonical URI of an answering rule
	Uri string `json:"uri,omitempty"`
	// Specifies if the rule is active or inactive. The default value is 'True'
	Enabled bool `json:"enabled,omitempty"`
	// Type of an answering rule, the default value is 'Custom' = ['BusinessHours', 'AfterHours', 'Custom']
	Type string `json:"type,omitempty"`
	// Name of an answering rule specified by user. Max number of symbols is 30. The default value is 'My Rule N' where 'N' is the first free number
	Name string `json:"name,omitempty"`
}
