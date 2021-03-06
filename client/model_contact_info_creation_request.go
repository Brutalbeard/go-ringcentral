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

type ContactInfoCreationRequest struct {
	// For User extension type only. Extension user first name
	FirstName string `json:"firstName"`
	// For User extension type only. Extension user last name
	LastName string `json:"lastName"`
	// Extension user company name
	Company  string `json:"company,omitempty"`
	JobTitle string `json:"jobTitle,omitempty"`
	// Email of extension user
	Email string `json:"email"`
	// Extension user contact phone number in [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) format
	BusinessPhone string `json:"businessPhone,omitempty"`
	// Extension user mobile (**non** Toll Free) phone number in [E.164](https://www.itu.int/rec/T-REC-E.164-201011-I) (with '+' sign) format
	MobilePhone     string                     `json:"mobilePhone,omitempty"`
	BusinessAddress ContactBusinessAddressInfo `json:"businessAddress,omitempty"`
	//  If 'True' then contact email is enabled as login name for this user. Please note that email should be unique in this case. The default value is 'False'
	EmailAsLoginName bool               `json:"emailAsLoginName,omitempty"`
	PronouncedName   PronouncedNameInfo `json:"pronouncedName,omitempty"`
	// Extension user department, if any
	Department string `json:"department,omitempty"`
}
