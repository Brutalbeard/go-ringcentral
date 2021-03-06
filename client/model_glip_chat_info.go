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

import (
	"time"
)

type GlipChatInfo struct {
	// Internal identifier of a chat
	Id string `json:"id,omitempty"`
	// Type of a chat
	Type string `json:"type,omitempty"`
	// For 'Team' chat type only. Team access level.
	Public bool `json:"public,omitempty"`
	// For 'Team','Everyone' chats types only. Chat name.
	Name string `json:"name,omitempty"`
	// For 'Team','Everyone' chats types only. Chat description.
	Description string `json:"description,omitempty"`
	// For 'Team' chat type only. Team status.
	Status string `json:"status,omitempty"`
	// Chat creation datetime in ISO 8601 format
	CreationTime time.Time `json:"creationTime,omitempty"`
	// Chat last change datetime in ISO 8601 format
	LastModifiedTime time.Time            `json:"lastModifiedTime,omitempty"`
	Members          []GlipChatMemberInfo `json:"members,omitempty"`
}
