# ExtendedErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ErrorCode** | **float32** | Numeric code used to identify the error. Integer. | 
**Status** | **string** |  | 
**Permanent** | Pointer to **bool** |  | [optional] 

## Methods

### NewExtendedErrorModel

`func NewExtendedErrorModel(message string, errorCode float32, status string, ) *ExtendedErrorModel`

NewExtendedErrorModel instantiates a new ExtendedErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedErrorModelWithDefaults

`func NewExtendedErrorModelWithDefaults() *ExtendedErrorModel`

NewExtendedErrorModelWithDefaults instantiates a new ExtendedErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ExtendedErrorModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExtendedErrorModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExtendedErrorModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorCode

`func (o *ExtendedErrorModel) GetErrorCode() float32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ExtendedErrorModel) GetErrorCodeOk() (*float32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ExtendedErrorModel) SetErrorCode(v float32)`

SetErrorCode sets ErrorCode field to given value.


### GetStatus

`func (o *ExtendedErrorModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtendedErrorModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtendedErrorModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPermanent

`func (o *ExtendedErrorModel) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *ExtendedErrorModel) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *ExtendedErrorModel) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.

### HasPermanent

`func (o *ExtendedErrorModel) HasPermanent() bool`

HasPermanent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


