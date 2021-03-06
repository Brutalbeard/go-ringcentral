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

type ModelInfo struct {
	// Addon identifier. For HardPhones of certain types, which are compatible with this addon identifier
	Id string `json:"id,omitempty"`
	// Device name
	Name string `json:"name,omitempty"`
	// Addons description
	Addons []AddonInfo `json:"addons"`
	// Device feature or multiple features supported
	Features []string `json:"features,omitempty"`
}
