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

type PhoneNumberExtensionInfo struct {
	Uri              string                       `json:"uri,omitempty"`
	Id               string                       `json:"id,omitempty"`
	PartnerId        string                       `json:"partnerId,omitempty"`
	ExtensionNumber  string                       `json:"extensionNumber,omitempty"`
	LoginName        string                       `json:"loginName,omitempty"`
	Contact          ExtensionContactInfo         `json:"contact,omitempty"`
	References       []Reference                  `json:"references,omitempty"`
	Name             string                       `json:"name,omitempty"`
	Type             string                       `json:"type,omitempty"`
	Status           string                       `json:"status,omitempty"`
	StatusInfo       StatusInfo                   `json:"statusInfo,omitempty"`
	Departments      []DepartmentResource         `json:"departments,omitempty"`
	ServiceFeatures  []ServiceFeatureValue        `json:"serviceFeatures,omitempty"`
	RegionalSettings RegionalSettingsInfo         `json:"regionalSettings,omitempty"`
	SetupWizardState string                       `json:"setupWizardState,omitempty"`
	Permissions      ExtensionPermissionsResource `json:"permissions,omitempty"`
	Password         string                       `json:"password,omitempty"`
	IvrPin           string                       `json:"ivrPin,omitempty"`
	ProfileImage     ProfileImageResource         `json:"profileImage,omitempty"`
}