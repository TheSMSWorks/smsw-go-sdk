/*
 * The SMS Works API
 *
 * The SMS Works provides a low-cost, reliable SMS API for developers. Pay only for delivered texts, all failed messages are refunded.
 *
 * API version: 1.5.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An array of objects containing metadata key/value pairs that have been saved on messages.
type QueryMetadata struct {
	Schema *MetaData `json:"schema,omitempty"`
}