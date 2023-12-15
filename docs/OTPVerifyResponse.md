# OTPVerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | The mobile number that the OTP was sent to | [optional] 
**Status** | Pointer to **string** | The status of the OTP. If the passcode is used within the validity period then this will be &#39;VERIFIED&#39;, otherwise it will be &#39;EXPIRED&#39; | [optional] 
**Passcode** | Pointer to **float32** | The passcode used. | [optional] 
**Validity** | Pointer to **float32** | The length of time in seconds for which the generated passcode is valid. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | A JSON object storing data supplied when this passcode was generated, for use in your application. | [optional] 
**Created** | Pointer to **string** | The ISO 8601 date/time at which this OTP was created | [optional] 
**Expires** | Pointer to **string** | The ISO 8601 date/time at which this OTP expires | [optional] 
**Modified** | Pointer to **string** | The ISO 8601 date/time at which this OTP was modified (typically when it was verified) | [optional] 

## Methods

### NewOTPVerifyResponse

`func NewOTPVerifyResponse() *OTPVerifyResponse`

NewOTPVerifyResponse instantiates a new OTPVerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTPVerifyResponseWithDefaults

`func NewOTPVerifyResponseWithDefaults() *OTPVerifyResponse`

NewOTPVerifyResponseWithDefaults instantiates a new OTPVerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *OTPVerifyResponse) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *OTPVerifyResponse) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *OTPVerifyResponse) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *OTPVerifyResponse) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetStatus

`func (o *OTPVerifyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OTPVerifyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OTPVerifyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OTPVerifyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPasscode

`func (o *OTPVerifyResponse) GetPasscode() float32`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *OTPVerifyResponse) GetPasscodeOk() (*float32, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *OTPVerifyResponse) SetPasscode(v float32)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *OTPVerifyResponse) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### GetValidity

`func (o *OTPVerifyResponse) GetValidity() float32`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *OTPVerifyResponse) GetValidityOk() (*float32, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *OTPVerifyResponse) SetValidity(v float32)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *OTPVerifyResponse) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetMetadata

`func (o *OTPVerifyResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OTPVerifyResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OTPVerifyResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OTPVerifyResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreated

`func (o *OTPVerifyResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OTPVerifyResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OTPVerifyResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OTPVerifyResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpires

`func (o *OTPVerifyResponse) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *OTPVerifyResponse) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *OTPVerifyResponse) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *OTPVerifyResponse) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetModified

`func (o *OTPVerifyResponse) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *OTPVerifyResponse) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *OTPVerifyResponse) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *OTPVerifyResponse) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


