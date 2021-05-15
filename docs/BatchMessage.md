# BatchMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | The sender of the message. Should be no longer than 11 characters for alphanumeric or 15 characters for numeric sender ID&#x27;s. No spaces or special characters. | [default to null]
**Destinations** | **[]string** | Telephone numbers of each of the recipients | [default to null]
**Content** | **string** | Message to send to the recipient | [default to null]
**Deliveryreporturl** | **string** | The url to which we should POST delivery reports to for this message. If none is specified, we&#x27;ll use the global delivery report URL that you&#x27;ve configured on your account page. | [optional] [default to null]
**Schedule** | **string** | Date-time at which to send the batch. This is only used by the batch/schedule service. | [optional] [default to null]
**Tag** | **string** | An identifying label for the message, which you can use to filter and report on messages you&#x27;ve sent later. Ideal for campaigns. A maximum of 280 characters. | [optional] [default to null]
**Ttl** | **float64** | The number of minutes before the message is deleted. Optional. Omit to prevent delivery report deletion. Integer. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

