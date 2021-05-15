/*
 * The SMS Works API
 *
 * The SMS Works provides a low-cost, reliable SMS API for developers. Pay only for delivered texts, all failed messages are refunded.
 *
 * API version: 1.5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// SMS Message Batch
type BatchMessage struct {
	// The sender of the message. Should be no longer than 11 characters for alphanumeric or 15 characters for numeric sender ID's. No spaces or special characters.
	Sender string `json:"sender"`
	// Telephone numbers of each of the recipients
	Destinations []string `json:"destinations"`
	// Message to send to the recipient
	Content string `json:"content"`
	// The url to which we should POST delivery reports to for this message. If none is specified, we'll use the global delivery report URL that you've configured on your account page.
	Deliveryreporturl string `json:"deliveryreporturl,omitempty"`
	// Date-time at which to send the batch. This is only used by the batch/schedule service.
	Schedule string `json:"schedule,omitempty"`
	// An identifying label for the message, which you can use to filter and report on messages you've sent later. Ideal for campaigns. A maximum of 280 characters.
	Tag string `json:"tag,omitempty"`
	// The number of minutes before the message is deleted. Optional. Omit to prevent delivery report deletion. Integer.
	Ttl float64 `json:"ttl,omitempty"`
}
