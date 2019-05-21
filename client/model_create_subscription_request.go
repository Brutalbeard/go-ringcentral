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

type CreateSubscriptionRequest struct {
	// Collection of URIs to API resources. For APNS transport type only message event filter is available
	EventFilters []string                        `json:"eventFilters"`
	DeliveryMode NotificationDeliveryModeRequest `json:"deliveryMode"`
	// Subscription lifetime in seconds. Max value is 7 days (604800 sec). For *WebHook* transport type max value might be set up to 630720000 seconds (20 years)
	ExpiresIn int32 `json:"expiresIn,omitempty"`
}
