# MessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batchid** | Pointer to **string** |  | [optional] 
**Content** | **string** |  | 
**Created** | **string** |  | 
**Customerid** | **string** |  | 
**Deliveryreporturl** | Pointer to **string** |  | [optional] 
**Destination** | **float32** |  | 
**Failurereason** | Pointer to [**MessageResponseFailurereason**](MessageResponseFailurereason.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Keyword** | Pointer to **string** |  | [optional] 
**Messageid** | **string** |  | 
**Modified** | **string** |  | 
**Schedule** | **string** |  | 
**Status** | **string** |  | 
**Sender** | **string** |  | 
**Tag** | **string** |  | 

## Methods

### NewMessageResponse

`func NewMessageResponse(content string, created string, customerid string, destination float32, messageid string, modified string, schedule string, status string, sender string, tag string, ) *MessageResponse`

NewMessageResponse instantiates a new MessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageResponseWithDefaults

`func NewMessageResponseWithDefaults() *MessageResponse`

NewMessageResponseWithDefaults instantiates a new MessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchid

`func (o *MessageResponse) GetBatchid() string`

GetBatchid returns the Batchid field if non-nil, zero value otherwise.

### GetBatchidOk

`func (o *MessageResponse) GetBatchidOk() (*string, bool)`

GetBatchidOk returns a tuple with the Batchid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchid

`func (o *MessageResponse) SetBatchid(v string)`

SetBatchid sets Batchid field to given value.

### HasBatchid

`func (o *MessageResponse) HasBatchid() bool`

HasBatchid returns a boolean if a field has been set.

### GetContent

`func (o *MessageResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreated

`func (o *MessageResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MessageResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MessageResponse) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetCustomerid

`func (o *MessageResponse) GetCustomerid() string`

GetCustomerid returns the Customerid field if non-nil, zero value otherwise.

### GetCustomeridOk

`func (o *MessageResponse) GetCustomeridOk() (*string, bool)`

GetCustomeridOk returns a tuple with the Customerid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerid

`func (o *MessageResponse) SetCustomerid(v string)`

SetCustomerid sets Customerid field to given value.


### GetDeliveryreporturl

`func (o *MessageResponse) GetDeliveryreporturl() string`

GetDeliveryreporturl returns the Deliveryreporturl field if non-nil, zero value otherwise.

### GetDeliveryreporturlOk

`func (o *MessageResponse) GetDeliveryreporturlOk() (*string, bool)`

GetDeliveryreporturlOk returns a tuple with the Deliveryreporturl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryreporturl

`func (o *MessageResponse) SetDeliveryreporturl(v string)`

SetDeliveryreporturl sets Deliveryreporturl field to given value.

### HasDeliveryreporturl

`func (o *MessageResponse) HasDeliveryreporturl() bool`

HasDeliveryreporturl returns a boolean if a field has been set.

### GetDestination

`func (o *MessageResponse) GetDestination() float32`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MessageResponse) GetDestinationOk() (*float32, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MessageResponse) SetDestination(v float32)`

SetDestination sets Destination field to given value.


### GetFailurereason

`func (o *MessageResponse) GetFailurereason() MessageResponseFailurereason`

GetFailurereason returns the Failurereason field if non-nil, zero value otherwise.

### GetFailurereasonOk

`func (o *MessageResponse) GetFailurereasonOk() (*MessageResponseFailurereason, bool)`

GetFailurereasonOk returns a tuple with the Failurereason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurereason

`func (o *MessageResponse) SetFailurereason(v MessageResponseFailurereason)`

SetFailurereason sets Failurereason field to given value.

### HasFailurereason

`func (o *MessageResponse) HasFailurereason() bool`

HasFailurereason returns a boolean if a field has been set.

### GetId

`func (o *MessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessageResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *MessageResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MessageResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MessageResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MessageResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetKeyword

`func (o *MessageResponse) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *MessageResponse) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *MessageResponse) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *MessageResponse) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetMessageid

`func (o *MessageResponse) GetMessageid() string`

GetMessageid returns the Messageid field if non-nil, zero value otherwise.

### GetMessageidOk

`func (o *MessageResponse) GetMessageidOk() (*string, bool)`

GetMessageidOk returns a tuple with the Messageid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageid

`func (o *MessageResponse) SetMessageid(v string)`

SetMessageid sets Messageid field to given value.


### GetModified

`func (o *MessageResponse) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *MessageResponse) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *MessageResponse) SetModified(v string)`

SetModified sets Modified field to given value.


### GetSchedule

`func (o *MessageResponse) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MessageResponse) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MessageResponse) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetStatus

`func (o *MessageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSender

`func (o *MessageResponse) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MessageResponse) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MessageResponse) SetSender(v string)`

SetSender sets Sender field to given value.


### GetTag

`func (o *MessageResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *MessageResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *MessageResponse) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


