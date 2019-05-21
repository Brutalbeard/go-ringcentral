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

type RegionalSettingsInfo struct {
	Timezone         TimezoneResource `json:"timezone,omitempty"`
	HomeCountry      CountryResource  `json:"homeCountry,omitempty"`
	Language         LanguageResource `json:"language,omitempty"`
	GreetingLanguage LanguageResource `json:"greetingLanguage,omitempty"`
	FormattingLocale LanguageResource `json:"formattingLocale,omitempty"`
}