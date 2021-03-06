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

type IvrMenuInfo struct {
	// Internal identifier of an IVR Menu extension
	Id string `json:"id,omitempty"`
	// Link to an IVR Menu extension resource
	Uri string `json:"uri,omitempty"`
	// First name of an IVR Menu user
	Name string `json:"name,omitempty"`
	// Number of an IVR Menu extension
	ExtensionNumber string            `json:"extensionNumber,omitempty"`
	Prompt          IvrMenuPromptInfo `json:"prompt,omitempty"`
	// Keys handling settings
	Actions []IvrMenuActionsInfo `json:"actions,omitempty"`
}
