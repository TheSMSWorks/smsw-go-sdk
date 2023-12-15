# ScheduledMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | Pointer to **string** | The sender of the message. Should be no longer than 11 characters for alphanumeric or 15 characters for numeric sender ID&#39;s. No spaces or special characters. | [optional] 
**Content** | Pointer to **string** | Message to be sent to the recipient | [optional] 
**Destination** | Pointer to **string** | For single scheduled messages, the mobile number of the recipient | [optional] 
**Destinations** | Pointer to **[]string** | For batch messages, the mobile numbers of each of the recipients | [optional] 
**Schedule** | Pointer to **string** | Date-time at which to send the batch. This is only used by the batch/schedule service. | [optional] 

## Methods

### NewScheduledMessage

`func NewScheduledMessage() *ScheduledMessage`

NewScheduledMessage instantiates a new ScheduledMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledMessageWithDefaults

`func NewScheduledMessageWithDefaults() *ScheduledMessage`

NewScheduledMessageWithDefaults instantiates a new ScheduledMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *ScheduledMessage) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ScheduledMessage) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ScheduledMessage) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *ScheduledMessage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetContent

`func (o *ScheduledMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ScheduledMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ScheduledMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ScheduledMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDestination

`func (o *ScheduledMessage) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ScheduledMessage) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ScheduledMessage) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ScheduledMessage) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinations

`func (o *ScheduledMessage) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *ScheduledMessage) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *ScheduledMessage) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *ScheduledMessage) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetSchedule

`func (o *ScheduledMessage) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduledMessage) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduledMessage) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ScheduledMessage) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


