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

type GetMessageSyncResponse struct {
	// List of message records with synchronization information
	Records  []GetMessageInfoResponse `json:"records"`
	SyncInfo SyncInfoMessages         `json:"syncInfo"`
}
