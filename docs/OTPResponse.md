# OTPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messageid** | Pointer to **string** | The messageid of the SMS used to send the OTP. Save this in your application to use when verifying passcodes. | [optional] 
**Status** | Pointer to **string** | The initial status of the OTP message. | [optional] 
**Credits** | Pointer to **float32** | The credit balance on your account | [optional] 
**CreditsUsed** | Pointer to **float32** | The number of credits used to send this message | [optional] 
**Messageparts** | Pointer to **float32** | The number of message parts used to send this message | [optional] 

## Methods

### NewOTPResponse

`func NewOTPResponse() *OTPResponse`

NewOTPResponse instantiates a new OTPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTPResponseWithDefaults

`func NewOTPResponseWithDefaults() *OTPResponse`

NewOTPResponseWithDefaults instantiates a new OTPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageid

`func (o *OTPResponse) GetMessageid() string`

GetMessageid returns the Messageid field if non-nil, zero value otherwise.

### GetMessageidOk

`func (o *OTPResponse) GetMessageidOk() (*string, bool)`

GetMessageidOk returns a tuple with the Messageid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageid

`func (o *OTPResponse) SetMessageid(v string)`

SetMessageid sets Messageid field to given value.

### HasMessageid

`func (o *OTPResponse) HasMessageid() bool`

HasMessageid returns a boolean if a field has been set.

### GetStatus

`func (o *OTPResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OTPResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OTPResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OTPResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCredits

`func (o *OTPResponse) GetCredits() float32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *OTPResponse) GetCreditsOk() (*float32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *OTPResponse) SetCredits(v float32)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *OTPResponse) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetCreditsUsed

`func (o *OTPResponse) GetCreditsUsed() float32`

GetCreditsUsed returns the CreditsUsed field if non-nil, zero value otherwise.

### GetCreditsUsedOk

`func (o *OTPResponse) GetCreditsUsedOk() (*float32, bool)`

GetCreditsUsedOk returns a tuple with the CreditsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsUsed

`func (o *OTPResponse) SetCreditsUsed(v float32)`

SetCreditsUsed sets CreditsUsed field to given value.

### HasCreditsUsed

`func (o *OTPResponse) HasCreditsUsed() bool`

HasCreditsUsed returns a boolean if a field has been set.

### GetMessageparts

`func (o *OTPResponse) GetMessageparts() float32`

GetMessageparts returns the Messageparts field if non-nil, zero value otherwise.

### GetMessagepartsOk

`func (o *OTPResponse) GetMessagepartsOk() (*float32, bool)`

GetMessagepartsOk returns a tuple with the Messageparts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageparts

`func (o *OTPResponse) SetMessageparts(v float32)`

SetMessageparts sets Messageparts field to given value.

### HasMessageparts

`func (o *OTPResponse) HasMessageparts() bool`

HasMessageparts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


