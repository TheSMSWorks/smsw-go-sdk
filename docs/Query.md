# Query

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the messages you would like returned (either &#x27;SENT&#x27;, &#x27;DELIVERED&#x27;, &#x27;EXPIRED&#x27;, &#x27;UNDELIVERABLE&#x27;, &#x27;REJECTED&#x27; or &#x27;INCOMING&#x27;) | [optional] [default to null]
**Credits** | **float64** | The number of credits used on the message. Floating point number. | [optional] [default to null]
**Destination** | **string** | The phone number of the recipient. Start UK numbers with 44 and drop the leading 0. | [optional] [default to null]
**Sender** | **string** | The sender of the message (this can be the configured sender name for an outbound message or the senders phone number for an inbound message). | [optional] [default to null]
**Keyword** | **string** | The keyword used in the inbound message | [optional] [default to null]
**From** | **string** | The date-time from which you would like matching messages | [optional] [default to null]
**To** | **string** | The date-time to which you would like matching messages | [optional] [default to null]
**Unread** | **bool** | In queries for incoming messages (&#x27;status&#x27; is &#x27;INCOMING&#x27;), specify whether you explicitly want unread messages (true) or read messages (false). Omit this parameter in other circumstances. | [optional] [default to null]
**Metadata** | [***QueryMetadata**](Query_metadata.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

