# OTPVerify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passcode** | Pointer to **string** | One-Time Passcode submitted to your application | [optional] 

## Methods

### NewOTPVerify

`func NewOTPVerify() *OTPVerify`

NewOTPVerify instantiates a new OTPVerify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTPVerifyWithDefaults

`func NewOTPVerifyWithDefaults() *OTPVerify`

NewOTPVerifyWithDefaults instantiates a new OTPVerify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasscode

`func (o *OTPVerify) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *OTPVerify) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *OTPVerify) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *OTPVerify) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


