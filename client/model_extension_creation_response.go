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

type ExtensionCreationResponse struct {
	// Internal identifier of an extension
	Id int64 `json:"id,omitempty"`
	// Canonical URI of an extension
	Uri     string      `json:"uri,omitempty"`
	Contact ContactInfo `json:"contact,omitempty"`
	// Number of department extension
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	// Extension user name
	Name string `json:"name,omitempty"`
	// For Partner Applications Internal identifier of an extension created by partner. The RingCentral supports the mapping of accounts and stores the corresponding account ID/extension ID for each partner ID of a client application. In request URIs partner IDs are accepted instead of regular RingCentral native IDs as path parameters using pid = XXX clause. Though in response URIs contain the corresponding account IDs and extension IDs. In all request and response bodies these values are reflected via partnerId attributes of account and extension
	PartnerId    string               `json:"partnerId,omitempty"`
	Permissions  ExtensionPermissions `json:"permissions,omitempty"`
	ProfileImage ProfileImageInfo     `json:"profileImage,omitempty"`
	// List of non-RC internal identifiers assigned to an extension
	References       []ReferenceInfo  `json:"references,omitempty"`
	Roles            []Roles          `json:"roles,omitempty"`
	RegionalSettings RegionalSettings `json:"regionalSettings,omitempty"`
	// Extension service features returned in response only when the logged-in user requests his/her own extension info, see also Extension Service Features
	ServiceFeatures []ExtensionServiceFeatureInfo `json:"serviceFeatures,omitempty"`
	// Specifies extension configuration wizard state (web service setup). The default value is 'NotStarted'
	SetupWizardState string `json:"setupWizardState,omitempty"`
	// Extension current state. If the status is 'Unassigned'. Returned for all extensions
	Status     string              `json:"status,omitempty"`
	StatusInfo ExtensionStatusInfo `json:"statusInfo,omitempty"`
	// Extension type
	Type string `json:"type,omitempty"`
	// Hides extension from showing in company directory. Supported for extensions of User type only
	Hidden bool `json:"hidden,omitempty"`
}
