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

type CallLogRecord struct {
	// Internal identifier of a cal log record
	Id string `json:"id,omitempty"`
	// Canonical URI of a call log record
	Uri string `json:"uri,omitempty"`
	// Internal identifier of a call session
	SessionId string            `json:"sessionId,omitempty"`
	From      CallLogCallerInfo `json:"from,omitempty"`
	To        CallLogCallerInfo `json:"to,omitempty"`
	// Call type
	Type string `json:"type,omitempty"`
	// Call direction
	Direction string `json:"direction,omitempty"`
	// Action description of the call operation
	Action string `json:"action,omitempty"`
	// Status description of the call operation
	Result string `json:"result,omitempty"`
	// Reason of a call result:  * `Accepted` - The call was connected to and accepted by this number  * `Connected` - The call was answered, but there was no response on how to handle the call (for example, a voice mail system answered the call and did not push \"1\" to accept)  * `Line Busy` - The phone number you dialed was busy  * `Not Answered` - The phone number you dialed was not answered  * `No Answer` - You did not answer the call  * `Hang Up` - The caller hung up before the call was answered  * `Stopped` - This attempt was stopped because the call was answered by another phone  * `Internal Error` - An internal error occurred when making the call. Please try again  * `No Credit` - There was not enough Calling Credit on your account to make this call  * `Restricted Number` - The number you dialed is restricted by RingCentral  * `Wrong Number` - The number you dialed has either been disconnected or is not a valid phone number. Please check the number and try again  * `International Disabled` - International calling is not enabled on your account. Contact customer service to activate International Calling  * `International Restricted` - The country and/or area you attempted to call has been prohibited by your administrator  * `Bad Number` - An error occurred when making the call. Please check the number before trying again  * `Info 411 Restricted` - Calling to 411 Information Services is restricted  * `Customer 611 Restricted` - 611 customer service is not supported. Please contact customer service at <(888) 555-1212>  * `No Digital Line` - This DigitalLine was either not plugged in or did not have an internet connection  * `Failed Try Again` - Call failed. Please try again  * `Max Call Limit` - The number of simultaneous calls to your account has reached its limit  * `Too Many Calls` - The number of simultaneous calls for per DigitalLine associated with Other Phone has reached its limit. Please contact customer service  * `Calls Not Accepted` - Your account was not accepting calls at this time  * `Number Not Allowed` - The number that was dialed to access your account is not allowed  * `Number Blocked` - This number is in your Blocked Numbers list  * `Number Disabled` - The phone number and/or area you attempted to call has been prohibited by your administrator  * `Resource Error` - An error occurred when making the call. Please try again  * `Call Loop` - A call loop occurred due to an incorrect call forwarding configuration. Please check that you are not forwarding calls back to your own account  * `Fax Not Received` - An incoming fax could not be received because a proper connection with the sender's fax machine could not be established  * `Fax Partially Sent` - The fax was only partially sent. Possible explanations include phone line quality to poor to maintain the connection or the call was dropped  * `Fax Not Sent` - An attempt to send the fax was made, but could not connect with the receiving fax machine  * `Fax Poor Line` - An attempt to send the fax was made, but the phone line quality was too poor to send the fax  * `Fax Prepare Error` - An internal error occurred when preparing the fax. Please try again  * `Fax Save Error` - An internal error occurred when saving the fax. Please try again  * `Fax Send Error` - An error occurred when sending the fax. Please try again
	Reason string `json:"reason,omitempty"`
	// The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	StartTime time.Time `json:"startTime,omitempty"`
	// Call duration in seconds
	Duration  int32                `json:"duration,omitempty"`
	Recording CallLogRecordingInfo `json:"recording,omitempty"`
	// For 'Detailed' view only. The datetime when the call log record was modified in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`
	// For 'Detailed' view only. Call transport
	Transport string                         `json:"transport,omitempty"`
	Extension ActiveCallsRecordExtensionInfo `json:"extension,omitempty"`
	Delegate  DelegateInfo                   `json:"delegate,omitempty"`
	// For 'Detailed' view only. Leg description
	Legs    []CallLogRecordLegInfo `json:"legs"`
	Message CallLogRecordMessage   `json:"message,omitempty"`
	// Returned only if this call was deleted. Value is set to 'True' in this case
	Deleted bool `json:"deleted,omitempty"`
}
