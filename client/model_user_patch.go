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

type UserPatch struct {
	// patch operations list
	Operations []PatchOperation `json:"Operations"`
	Schemas    []string         `json:"schemas"`
}
