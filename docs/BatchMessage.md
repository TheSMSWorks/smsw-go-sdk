# BatchMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | The sender of the message. Should be no longer than 11 characters for alphanumeric or 15 characters for numeric sender ID&#39;s. No spaces or special characters. | 
**Destinations** | **[]string** | Telephone numbers of each of the recipients | 
**Content** | **string** | Message to send to the recipient | 
**Deliveryreporturl** | Pointer to **string** | The url to which we should POST delivery reports to for this message. If none is specified, we&#39;ll use the global delivery report URL that you&#39;ve configured on your account page. | [optional] 
**Schedule** | Pointer to **string** | Date-time at which to send the batch. This is only used by the batch/schedule service. | [optional] 
**Tag** | Pointer to **string** | An identifying label for the message, which you can use to filter and report on messages you&#39;ve sent later. Ideal for campaigns. A maximum of 280 characters. | [optional] 
**Ttl** | Pointer to **float32** | The number of minutes before the delivery report is deleted. Optional. Omit to prevent delivery report deletion. Integer. | [optional] 
**Validity** | Pointer to **float32** | The optional number of minutes to attempt delivery before the message is marked as EXPIRED. Optional. The default is 2880 minutes. Integer. | [optional] 
**Ai** | Pointer to **bool** | Used to determine whether The SMS Works AI Optimiser should be used in the event that the message is just longer than the 1 or 2 credit boundary. This setting overrides the AI Optimiser configuration on your SMS Works account. | [optional] 

## Methods

### NewBatchMessage

`func NewBatchMessage(sender string, destinations []string, content string, ) *BatchMessage`

NewBatchMessage instantiates a new BatchMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchMessageWithDefaults

`func NewBatchMessageWithDefaults() *BatchMessage`

NewBatchMessageWithDefaults instantiates a new BatchMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *BatchMessage) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *BatchMessage) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *BatchMessage) SetSender(v string)`

SetSender sets Sender field to given value.


### GetDestinations

`func (o *BatchMessage) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *BatchMessage) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *BatchMessage) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.


### GetContent

`func (o *BatchMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BatchMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BatchMessage) SetContent(v string)`

SetContent sets Content field to given value.


### GetDeliveryreporturl

`func (o *BatchMessage) GetDeliveryreporturl() string`

GetDeliveryreporturl returns the Deliveryreporturl field if non-nil, zero value otherwise.

### GetDeliveryreporturlOk

`func (o *BatchMessage) GetDeliveryreporturlOk() (*string, bool)`

GetDeliveryreporturlOk returns a tuple with the Deliveryreporturl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryreporturl

`func (o *BatchMessage) SetDeliveryreporturl(v string)`

SetDeliveryreporturl sets Deliveryreporturl field to given value.

### HasDeliveryreporturl

`func (o *BatchMessage) HasDeliveryreporturl() bool`

HasDeliveryreporturl returns a boolean if a field has been set.

### GetSchedule

`func (o *BatchMessage) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BatchMessage) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BatchMessage) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *BatchMessage) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTag

`func (o *BatchMessage) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BatchMessage) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BatchMessage) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BatchMessage) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTtl

`func (o *BatchMessage) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *BatchMessage) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *BatchMessage) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *BatchMessage) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetValidity

`func (o *BatchMessage) GetValidity() float32`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *BatchMessage) GetValidityOk() (*float32, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *BatchMessage) SetValidity(v float32)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *BatchMessage) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetAi

`func (o *BatchMessage) GetAi() bool`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *BatchMessage) GetAiOk() (*bool, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *BatchMessage) SetAi(v bool)`

SetAi sets Ai field to given value.

### HasAi

`func (o *BatchMessage) HasAi() bool`

HasAi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


