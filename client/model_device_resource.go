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

type DeviceResource struct {
	// Internal identifier of a device
	Id string `json:"id,omitempty"`
	// Canonical URI of a device
	Uri string `json:"uri,omitempty"`
	// Device identification number (stock keeping unit) in the format TP-ID [-AT-AC], where TP is device type (HP for RC HardPhone, DV for all other devices including softphone); ID - device model ID; AT -addon type ID; AC - addon count (if any). For example 'HP-56-2-2'
	Sku string `json:"sku,omitempty"`
	// Device type
	Type string `json:"type,omitempty"`
	// Status of a device
	Status string `json:"status,omitempty"`
	// Device name. Mandatory if ordering SoftPhone or OtherPhone . Optional for HardPhone . If not specified for HardPhone, then device model name is used as device name
	Name string `json:"name,omitempty"`
	// Serial number for HardPhone (is returned only when the phone is shipped and provisioned); endpoint_id for softphone and mobile applications
	Serial string `json:"serial,omitempty"`
	// PC name for softphone
	ComputerName string                 `json:"computerName,omitempty"`
	Model        ModelInfo              `json:"model,omitempty"`
	Extension    ExtensionResourceIntId `json:"extension,omitempty"`
	// Phone lines information
	PhoneLines              []PhoneLineResource             `json:"phoneLines,omitempty"`
	EmergencyServiceAddress EmergencyServiceAddressResource `json:"emergencyServiceAddress,omitempty"`
	Shipping                ShippingResource                `json:"shipping,omitempty"`
	// Box billing identifier of a device. Applicable only for HardPhones. It is an alternative way to identify the device to be ordered. Either  model structure, or boxBillingId must be specified for HardPhone
	BoxBillingId int32 `json:"boxBillingId,omitempty"`
	// Pooling type of a deviceHost - device with standalone paid phone line which can be linked to Glip/Softphone instanceGuest - device with a linked phone lineNone - device without a phone line or with specific line (free, BLA, etc.) = ['Host', 'Guest', 'None']
	LinePooling string `json:"linePooling,omitempty"`
	// Supported only for devices assigned to Limited extensions. If true, enables users to log in to this phone as a common phone.
	UseAsCommonPhone bool `json:"useAsCommonPhone,omitempty"`
}
