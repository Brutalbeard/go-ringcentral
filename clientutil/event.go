package clientutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	rc "github.com/brutalbeard/go-ringcentral/client"
)

const (
	ThisDir = "src/github.com/brutalbeard/go-ringcentral/clientutil"
)

// EventSimple is a event for Instant SMS for testing purposes
type EventSimple struct {
	UUID           string    `json:"uuid,omitempty"`
	Event          string    `json:"event,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	SubscriptionId string    `json:"subscriptionId,omitempty"`
	OwnerId        string    `json:"ownerId,omitempty"`
	Body           EventBody `json:"body,omitempty"`
}

// Event is a generic event.
type Event struct {
	UUID           string           `json:"uuid,omitempty"`
	Event          string           `json:"event,omitempty"`
	Timestamp      time.Time        `json:"timestamp,omitempty"`
	SubscriptionId string           `json:"subscriptionId,omitempty"`
	OwnerId        string           `json:"ownerId,omitempty"`
	Body           EventBodyWrapper `json:"body,omitempty"`
}

func (evt *Event) IsEventType(eventType EventType) bool {
	parsedEvtType, err := ParseEventTypeForFilter(evt.Event)
	if err != nil {
		return false
	}
	if eventType == parsedEvtType {
		fmt.Println(evt.Event)
		fmt.Println(parsedEvtType.String())
		//panic("AA")
		return true
	}
	return false
}

type EventBodyWrapper struct {
	Raw    string      `json:"raw"`
	Parsed interface{} `json:"parsed"`
}

func (w *EventBodyWrapper) UnmarshalJSON(data []byte) error {
	w.Raw = string(data)
	return nil
}

func EventParseBytes(data []byte) (*Event, error) {
	evt := &Event{}
	err := json.Unmarshal(data, evt)
	return evt, err
}

func (evt *Event) GetInstantMessageBody() (*rc.InstantMessageEvent, error) {
	evtType, err := ParseEventTypeForFilter(evt.Event)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse event: %v", evt.Event)
	} else if evtType != InstantMessageEvent {
		return nil, fmt.Errorf("Incorrect event type: %v", evtType.String())
	}
	body := &rc.InstantMessageEvent{}
	err = json.Unmarshal([]byte(evt.Body.Raw), body)
	return body, err
}

func (evt *Event) GetGlipPostEventBody() (*rc.GlipPostEvent, error) {
	evtType, err := ParseEventTypeForFilter(evt.Event)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse event: %v", evt.Event)
	} else if evtType != GlipPostEvent {
		return nil, fmt.Errorf("Incorrect event type: %v", evtType.String())
	}
	body := &rc.GlipPostEvent{}
	err = json.Unmarshal([]byte(evt.Body.Raw), body)
	return body, err
}

func GetFileBytesForEventType(eventType string) ([]byte, error) {
	file := GetFilePathForEventType(eventType)
	return ioutil.ReadFile(file)
}

func GetFilePathForEventType(eventType string) string {
	filename := GetFileNameForEventType(eventType)
	return filepath.Join(os.Getenv("GOPATH"), ThisDir, filename)
}

func GetFileNameForEventType(eventType string) string {
	return fmt.Sprintf("example_event_%s.json", strings.ToLower(eventType))
}

type EventBody struct {
	ExpiresIn            int       `json:"expiresIn,omitempty"`
	CreationTime         time.Time `json:"creationTime,omitempty"`
	Direction            string    `json:"direction,omitempty"`
	LastModificationTime time.Time `json:"lastModificationTime,omitempty"`
	Subject              string    `json:"subject,omitempty"`
	Type                 string    `json:"type,omitempty"`
}

/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
/*
type InstantMessageEvent struct {

	// Internal identifier of an message
	Id string `json:"id,omitempty"`

	// Message receiver(s) information
	To []NotificationRecipientInfo `json:"to,omitempty"`

	// Extension Type
	From *SenderInfo `json:"from,omitempty"`

	// Type of a message. The default value is 'SMS'
	Type_ string `json:"type,omitempty"`

	// Message creation datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	CreationTime time.Time `json:"creationTime,omitempty"`

	// Datetime when the message was modified in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`

	// Status of a message. The default value is 'Unread'
	ReadStatus string `json:"readStatus,omitempty"`

	// The default value is 'Normal'
	Priority string `json:"priority,omitempty"`

	// Message attachment data
	Attachments []MessageAttachmentInfo `json:"attachments,omitempty"`

	// Message direction. The default value is 'Inbound'
	Direction string `json:"direction,omitempty"`

	// Message availability status. The default value is 'Alive'
	Availability string `json:"availability,omitempty"`

	// Message subject. It replicates message text which is also returned as an attachment
	Subject string `json:"subject,omitempty"`

	// Status of a message. The default value is 'Received'
	MessageStatus string `json:"messageStatus,omitempty"`

	// Identifier of the conversation the message belongs to
	ConversationId string `json:"conversationId,omitempty"`
}
*/
