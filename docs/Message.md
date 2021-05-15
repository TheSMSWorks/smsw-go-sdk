# Message

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | The sender of the message. Should be no longer than 11 characters for alphanumeric or 15 characters for numeric sender ID&#x27;s. No spaces or special characters. | [default to null]
**Destination** | **string** | Telephone number of the recipient | [default to null]
**Content** | **string** | Message to send to the recipient. Content can be up to 1280 characters in length. Messages of 160 characters or fewer are charged 1 credit. If your message is longer than 160 characters then it will be broken down in to chunks of 153 characters before being sent to the recipient&#x27;s handset, and you will be charged 1 credit for each 153 characters. Messages sent to numbers registered outside the UK will be typically charged double credits, but for certain countries may be charged fractions of credits (e.g. 2.5). Please contact us for rates for each country. | [default to null]
**Deliveryreporturl** | **string** | The url to which we should POST delivery reports to for this message. If none is specified, we&#x27;ll use the global delivery report URL that you&#x27;ve configured on your account page. | [optional] [default to null]
**Schedule** | **string** | Date at which to send the message. This is only used by the message/schedule service and can be left empty for other services. | [optional] [default to null]
**Tag** | **string** | An identifying label for the message, which you can use to filter and report on messages you&#x27;ve sent later. Ideal for campaigns. A maximum of 280 characters. | [optional] [default to null]
**Ttl** | **float64** | The optional number of minutes before the message is deleted. Optional. Omit to prevent delivery report deletion. Integer. | [optional] [default to null]
**Responseemail** | **[]string** | An optional list of email addresses to forward responses to this specific message to. An SMS Works Reply Number is required to use this feature. | [optional] [default to null]
**Metadata** | [***MessageMetadata**](Message_metadata.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

