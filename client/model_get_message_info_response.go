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

type GetMessageInfoResponse struct {
	// Internal identifier of a message
	Id string `json:"id,omitempty"`
	// Canonical URI of a message
	Uri string `json:"uri,omitempty"`
	// The list of message attachments
	Attachments []MessageAttachmentInfo `json:"attachments,omitempty"`
	// Message availability status. Message in 'Deleted' state is still preserved with all its attachments and can be restored. 'Purged' means that all attachments are already deleted and the message itself is about to be physically deleted shortly
	Availability string `json:"availability,omitempty"`
	// SMS and Pager only. Identifier of a conversation the message belongs to
	ConversationId int32            `json:"conversationId,omitempty"`
	Conversation   ConversationInfo `json:"conversation,omitempty"`
	// Message creation datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	CreationTime time.Time `json:"creationTime,omitempty"`
	// SMS only. Delivery error code returned by gateway
	DeliveryErrorCode string `json:"deliveryErrorCode,omitempty"`
	// Message direction. Note that for some message types not all directions are allowed. For example voicemail messages can be only inbound
	Direction string `json:"direction,omitempty"`
	// Fax only. Page count in a fax message
	FaxPageCount int32 `json:"faxPageCount,omitempty"`
	// Fax only. Resolution of a fax message. 'High' for black and white image scanned at 200 dpi, 'Low' for black and white image scanned at 100 dpi
	FaxResolution string                         `json:"faxResolution,omitempty"`
	From          MessageStoreCallerInfoResponse `json:"from,omitempty"`
	// The datetime when the message was modified on server in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`
	// Message status. Different message types may have different allowed status values.For outbound faxes the aggregated message status is returned: If status for at least one recipient is 'Queued', then 'Queued' value is returned If status for at least one recipient is 'SendingFailed', then 'SendingFailed' value is returned In other cases Sent status is returned
	MessageStatus string `json:"messageStatus,omitempty"`
	// 'Pager' only. 'True' if at least one of the message recipients is 'Department' extension
	PgToDepartment bool `json:"pgToDepartment,omitempty"`
	// Message priority
	Priority string `json:"priority,omitempty"`
	// Message read status
	ReadStatus string `json:"readStatus,omitempty"`
	// SMS only. The datetime when outbound SMS was delivered to recipient's handset in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. It is filled only if the carrier sends a delivery receipt to RingCentral
	SmsDeliveryTime time.Time `json:"smsDeliveryTime,omitempty"`
	// SMS only. Number of attempts made to send an outbound SMS to the gateway (if gateway is temporary unavailable)
	SmsSendingAttemptsCount int32 `json:"smsSendingAttemptsCount,omitempty"`
	// Message subject. For SMS and Pager messages it replicates message text which is also returned as an attachment
	Subject string `json:"subject,omitempty"`
	// Recipient information
	To []MessageStoreCallerInfoResponse `json:"to,omitempty"`
	// Message type
	Type string `json:"type,omitempty"`
	// Voicemail only. Status of voicemail to text transcription. If VoicemailToText feature is not activated for account, the 'NotAvailable' value is returned
	VmTranscriptionStatus string `json:"vmTranscriptionStatus,omitempty"`
}
