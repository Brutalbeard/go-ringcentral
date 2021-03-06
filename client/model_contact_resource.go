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

type ContactResource struct {
	Account         AccountResource `json:"account,omitempty"`
	Department      string          `json:"department,omitempty"`
	Email           string          `json:"email,omitempty"`
	ExtensionNumber string          `json:"extensionNumber,omitempty"`
	// First name of a contact, for user extensions only
	FirstName string `json:"firstName,omitempty"`
	// Last name of a contact, for user extensions only
	LastName string `json:"lastName,omitempty"`
	// Name of a contact, for non-user extensions
	Name         string                               `json:"name,omitempty"`
	Id           string                               `json:"id,omitempty"`
	JobTitle     string                               `json:"jobTitle,omitempty"`
	PhoneNumbers []PhoneNumberResource                `json:"phoneNumbers,omitempty"`
	ProfileImage AccountDirectoryProfileImageResource `json:"profileImage,omitempty"`
	Site         BusinessSiteResource                 `json:"site,omitempty"`
	Status       string                               `json:"status,omitempty"`
	Type         string                               `json:"type,omitempty"`
}
