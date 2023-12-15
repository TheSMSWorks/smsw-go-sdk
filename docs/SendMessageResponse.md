# SendMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messageid** | **string** |  | 
**Status** | **string** |  | 
**Credits** | **float32** | The number of remaining credits on your SMS Works account. Floating point number. | 
**CreditsUsed** | **float32** | The number of credits used to send the message. Floating point number. | 

## Methods

### NewSendMessageResponse

`func NewSendMessageResponse(messageid string, status string, credits float32, creditsUsed float32, ) *SendMessageResponse`

NewSendMessageResponse instantiates a new SendMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageResponseWithDefaults

`func NewSendMessageResponseWithDefaults() *SendMessageResponse`

NewSendMessageResponseWithDefaults instantiates a new SendMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageid

`func (o *SendMessageResponse) GetMessageid() string`

GetMessageid returns the Messageid field if non-nil, zero value otherwise.

### GetMessageidOk

`func (o *SendMessageResponse) GetMessageidOk() (*string, bool)`

GetMessageidOk returns a tuple with the Messageid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageid

`func (o *SendMessageResponse) SetMessageid(v string)`

SetMessageid sets Messageid field to given value.


### GetStatus

`func (o *SendMessageResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendMessageResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendMessageResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCredits

`func (o *SendMessageResponse) GetCredits() float32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *SendMessageResponse) GetCreditsOk() (*float32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *SendMessageResponse) SetCredits(v float32)`

SetCredits sets Credits field to given value.


### GetCreditsUsed

`func (o *SendMessageResponse) GetCreditsUsed() float32`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *SendMessageResponse) GetCreditsUsedOk() (*float32, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *SendMessageResponse) SetCreditsUsed(v float32)`

SetCreditsUsed sets CreditsUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


