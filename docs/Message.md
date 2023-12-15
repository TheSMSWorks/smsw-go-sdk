# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** | The sender of the message. Should be no longer than 11 characters for alphanumeric or 15 characters for numeric sender ID&#39;s. No spaces or special characters. | 
**Destination** | **string** | Telephone number of the recipient | 
**Content** | **string** | Message to send to the recipient. Content can be up to 1280 characters in length. Messages of 160 characters or fewer are charged 1 credit. If your message is longer than 160 characters then it will be broken down in to chunks of 153 characters before being sent to the recipient&#39;s handset, and you will be charged 1 credit for each 153 characters. Messages sent to numbers registered outside the UK will be typically charged double credits, but for certain countries may be charged fractions of credits (e.g. 2.5). Please contact us for rates for each country. | 
**Deliveryreporturl** | Pointer to **string** | The url to which we should POST delivery reports to for this message. If none is specified, we&#39;ll use the global delivery report URL that you&#39;ve configured on your account page. | [optional] 
**Schedule** | Pointer to **string** | Date at which to send the message. This is only used by the message/schedule service and can be left empty for other services. | [optional] 
**Tag** | Pointer to **string** | An identifying label for the message, which you can use to filter and report on messages you&#39;ve sent later. Ideal for campaigns. A maximum of 280 characters. | [optional] 
**Ttl** | Pointer to **float32** | The optional number of minutes before the delivery report is deleted. Optional. Omit to prevent delivery report deletion. Integer. | [optional] 
**Responseemail** | Pointer to **[]string** | An optional list of email addresses to forward responses to this specific message to. An SMS Works Reply Number is required to use this feature. | [optional] 
**Metadata** | Pointer to [**MessageMetadata**](MessageMetadata.md) |  | [optional] 
**Validity** | Pointer to **float32** | The optional number of minutes to attempt delivery before the message is marked as EXPIRED. Optional. The default is 2880 minutes. Integer. | [optional] 
**Ai** | Pointer to **bool** | Used to determine whether The SMS Works AI Optimiser should be used in the event that the message is just longer than the 1 or 2 credit boundary. This setting overrides the AI Optimiser configuration on your SMS Works account. | [optional] 

## Methods

### NewMessage

`func NewMessage(sender string, destination string, content string, ) *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *Message) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Message) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Message) SetSender(v string)`

SetSender sets Sender field to given value.


### GetDestination

`func (o *Message) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Message) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Message) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetContent

`func (o *Message) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Message) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Message) SetContent(v string)`

SetContent sets Content field to given value.


### GetDeliveryreporturl

`func (o *Message) GetDeliveryreporturl() string`

GetDeliveryreporturl returns the Deliveryreporturl field if non-nil, zero value otherwise.

### GetDeliveryreporturlOk

`func (o *Message) GetDeliveryreporturlOk() (*string, bool)`

GetDeliveryreporturlOk returns a tuple with the Deliveryreporturl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryreporturl

`func (o *Message) SetDeliveryreporturl(v string)`

SetDeliveryreporturl sets Deliveryreporturl field to given value.

### HasDeliveryreporturl

`func (o *Message) HasDeliveryreporturl() bool`

HasDeliveryreporturl returns a boolean if a field has been set.

### GetSchedule

`func (o *Message) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Message) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Message) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Message) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTag

`func (o *Message) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Message) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Message) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Message) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTtl

`func (o *Message) GetTtl() float32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Message) GetTtlOk() (*float32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Message) SetTtl(v float32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Message) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetResponseemail

`func (o *Message) GetResponseemail() []string`

GetResponseemail returns the Responseemail field if non-nil, zero value otherwise.

### GetResponseemailOk

`func (o *Message) GetResponseemailOk() (*[]string, bool)`

GetResponseemailOk returns a tuple with the Responseemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseemail

`func (o *Message) SetResponseemail(v []string)`

SetResponseemail sets Responseemail field to given value.

### HasResponseemail

`func (o *Message) HasResponseemail() bool`

HasResponseemail returns a boolean if a field has been set.

### GetMetadata

`func (o *Message) GetMetadata() MessageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Message) GetMetadataOk() (*MessageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Message) SetMetadata(v MessageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Message) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetValidity

`func (o *Message) GetValidity() float32`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *Message) GetValidityOk() (*float32, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *Message) SetValidity(v float32)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *Message) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetAi

`func (o *Message) GetAi() bool`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *Message) GetAiOk() (*bool, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *Message) SetAi(v bool)`

SetAi sets Ai field to given value.

### HasAi

`func (o *Message) HasAi() bool`

HasAi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


