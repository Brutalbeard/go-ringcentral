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

type CreateDataExportTaskRequest struct {
	// Starting time for data collection. The default value is current time minus 24 hours. If `dateTo` is not specified then its value is set to `dateFrom` plus 24 hours. The specified time range should not be greater than 7 days
	DateFrom string `json:"dateFrom,omitempty"`
	// Ending time for data collection. The default value is current time. If `dateFrom` is not specified then its value is set to `dateTo` minus 24 hours. The specified time range should not be greater than 7 days
	DateTo string `json:"dateTo,omitempty"`
	// List of users which data is collected. The following data will be exported: posts, tasks, events, etc. posted by the user(s); posts addressing the user(s) via direct and @Mentions; tasks assigned to the listed user(s). The list of 30 users per request is supported.
	UserIds []string `json:"userIds,omitempty"`
	// List of chats from which the data (posts, files, tasks, events, notes, etc.) will be collected
	ChatIds []string `json:"chatIds,omitempty"`
}
