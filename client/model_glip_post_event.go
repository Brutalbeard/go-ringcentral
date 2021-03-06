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

type GlipPostEvent struct {
	// Internal identifier of a post
	Id string `json:"id,omitempty"`
	// Type of a post event
	EventType string `json:"eventType,omitempty"`
	// Internal identifier of a group a post belongs to
	GroupId string `json:"groupId,omitempty"`
	// Type of a post. 'TextMessage' - an incoming text message; 'PersonJoined' - a message notifying that person has joined a conversation; 'PersonsAdded' - a message notifying that a person(s) were added to a conversation
	Type string `json:"type,omitempty"`
	// For 'TextMessage' post type only. Message text
	Text string `json:"text,omitempty"`
	// Internal identifier of a user - author of a post
	CreatorId string `json:"creatorId,omitempty"`
	// For PersonsAdded post type only. Identifiers of persons added to a group
	AddedPersonIds []string `json:"addedPersonIds,omitempty"`
	// For PersonsRemoved post type only. Identifiers of persons removed from a group
	RemovedPersonIds []string `json:"removedPersonIds,omitempty"`
	// List of at mentions in post text with names.
	Mentions []GlipMentionsInfo `json:"mentions,omitempty"`
	// Post creation datetime in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format
	CreationTime time.Time `json:"creationTime,omitempty"`
	// Post last change datetime in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`
}
