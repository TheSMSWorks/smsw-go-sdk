# ScheduledMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the scheduled message (either &#39;SCHEDULED&#39;, &#39;PROCESSED&#39; or &#39;CANCELLED&#39;) | [optional] 
**Id** | Pointer to **string** | The scheduled message ID | [optional] 
**Batch** | Pointer to **bool** | Describes whether the a batch of messages has been scheduled, or just a single message | [optional] 
**Message** | Pointer to [**ScheduledMessagesResponseMessage**](ScheduledMessagesResponseMessage.md) |  | [optional] 

## Methods

### NewScheduledMessagesResponse

`func NewScheduledMessagesResponse() *ScheduledMessagesResponse`

NewScheduledMessagesResponse instantiates a new ScheduledMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledMessagesResponseWithDefaults

`func NewScheduledMessagesResponseWithDefaults() *ScheduledMessagesResponse`

NewScheduledMessagesResponseWithDefaults instantiates a new ScheduledMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ScheduledMessagesResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledMessagesResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledMessagesResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduledMessagesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetId

`func (o *ScheduledMessagesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledMessagesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledMessagesResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduledMessagesResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBatch

`func (o *ScheduledMessagesResponse) GetBatch() bool`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *ScheduledMessagesResponse) GetBatchOk() (*bool, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *ScheduledMessagesResponse) SetBatch(v bool)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *ScheduledMessagesResponse) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetMessage

`func (o *ScheduledMessagesResponse) GetMessage() ScheduledMessagesResponseMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScheduledMessagesResponse) GetMessageOk() (*ScheduledMessagesResponseMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScheduledMessagesResponse) SetMessage(v ScheduledMessagesResponseMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScheduledMessagesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


