/*
 * The SMS Works API
 *
 * The SMS Works provides a low-cost, reliable SMS API for developers. Pay only for delivered texts, all failed messages are refunded.
 *
 * API version: 1.5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// search parameters for querying the message database
type Query struct {
	// The status of the messages you would like returned (either 'SENT', 'DELIVERED', 'EXPIRED', 'UNDELIVERABLE', 'REJECTED' or 'INCOMING')
	Status string `json:"status,omitempty"`
	// The number of credits used on the message. Floating point number.
	Credits float64 `json:"credits,omitempty"`
	// The phone number of the recipient. Start UK numbers with 44 and drop the leading 0.
	Destination string `json:"destination,omitempty"`
	// The sender of the message (this can be the configured sender name for an outbound message or the senders phone number for an inbound message).
	Sender string `json:"sender,omitempty"`
	// The keyword used in the inbound message
	Keyword string `json:"keyword,omitempty"`
	// The date-time from which you would like matching messages
	From string `json:"from,omitempty"`
	// The date-time to which you would like matching messages
	To string `json:"to,omitempty"`
	// In queries for incoming messages ('status' is 'INCOMING'), specify whether you explicitly want unread messages (true) or read messages (false). Omit this parameter in other circumstances.
	Unread bool `json:"unread,omitempty"`
	Metadata *QueryMetadata `json:"metadata,omitempty"`
}
