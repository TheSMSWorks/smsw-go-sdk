# OTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | Pointer to **string** | The sender of the message. Should be no longer than 11 characters for alphanumeric or 15 characters for numeric sender ID&#39;s. No spaces or special characters. | [optional] 
**Destination** | Pointer to **string** | The phone number of the recipient. | [optional] 
**Length** | Pointer to **map[string]interface{}** | The length of the generated passcode. The default length is 6 characters, which will apply if this parameter is omitted. All generated passcodes are numeric. Optional. | [optional] 
**Template** | Pointer to **string** | A template to use as the content for the message. You must include the &#39;{{passcode}}&#39; placeholder, which will be replaced by the generated passcode when the message is sent. Optional. | [optional] 
**Validity** | Pointer to **float32** | The length of time in seconds for which the generated passcode should be valid. Optional. | [optional] 
**Passcode** | Pointer to **string** | A passcode you supply for use in the message template. This will be stored on the OTP record in our system for later verification. Optional. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | A JSON object of no longer than 1024 bytes, containing as many parameters as you wish, to store data for use in your application. This will be returned when you verify the passcode. | [optional] 

## Methods

### NewOTP

`func NewOTP() *OTP`

NewOTP instantiates a new OTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTPWithDefaults

`func NewOTPWithDefaults() *OTP`

NewOTPWithDefaults instantiates a new OTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *OTP) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *OTP) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *OTP) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *OTP) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetDestination

`func (o *OTP) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *OTP) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *OTP) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *OTP) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetLength

`func (o *OTP) GetLength() map[string]interface{}`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *OTP) GetLengthOk() (*map[string]interface{}, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *OTP) SetLength(v map[string]interface{})`

SetLength sets Length field to given value.

### HasLength

`func (o *OTP) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetTemplate

`func (o *OTP) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *OTP) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *OTP) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *OTP) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetValidity

`func (o *OTP) GetValidity() float32`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *OTP) GetValidityOk() (*float32, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *OTP) SetValidity(v float32)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *OTP) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetPasscode

`func (o *OTP) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *OTP) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *OTP) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *OTP) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### GetMetadata

`func (o *OTP) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OTP) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OTP) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OTP) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


