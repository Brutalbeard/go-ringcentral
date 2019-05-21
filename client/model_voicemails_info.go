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

type VoicemailsInfo struct {
	// Email notification flag
	NotifyByEmail bool `json:"notifyByEmail,omitempty"`
	// SMS notification flag
	NotifyBySms bool `json:"notifyBySms,omitempty"`
	// List of recipient email addresses for voicemail notifications. Returned if specified, in both modes (advanced/basic). Applied in advanced mode only
	AdvancedEmailAddresses []string `json:"advancedEmailAddresses,omitempty"`
	// List of recipient phone numbers for voicemail notifications. Returned if specified, in both modes (advanced/basic). Applied in advanced mode only
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses,omitempty"`
	// Indicates whether voicemail should be attached to email
	IncludeAttachment bool `json:"includeAttachment,omitempty"`
	// Indicates whether email should be automatically marked as read
	MarkAsRead bool `json:"markAsRead,omitempty"`
}
