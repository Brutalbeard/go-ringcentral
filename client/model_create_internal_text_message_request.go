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

type CreateInternalTextMessageRequest struct {
	From PagerCallerInfoRequest `json:"from"`
	// Internal identifier of a message this message replies to
	ReplyOn int32 `json:"replyOn,omitempty"`
	// Text of a pager message. Max length is 1024 symbols (2-byte UTF-16 encoded). If a character is encoded in 4 bytes in UTF-16 it is treated as 2 characters, thus restricting the maximum message length to 512 symbols
	Text string `json:"text"`
	// Optional if replyOn parameter is specified. Receiver of a pager message.
	To []PagerCallerInfoRequest `json:"to,omitempty"`
}
