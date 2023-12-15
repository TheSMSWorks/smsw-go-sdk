# MessageResponseFailurereason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **float32** | Numeric code that defines the error. Integer. | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Permanent** | Pointer to **bool** |  | [optional] 

## Methods

### NewMessageResponseFailurereason

`func NewMessageResponseFailurereason() *MessageResponseFailurereason`

NewMessageResponseFailurereason instantiates a new MessageResponseFailurereason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageResponseFailurereasonWithDefaults

`func NewMessageResponseFailurereasonWithDefaults() *MessageResponseFailurereason`

NewMessageResponseFailurereasonWithDefaults instantiates a new MessageResponseFailurereason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MessageResponseFailurereason) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MessageResponseFailurereason) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MessageResponseFailurereason) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *MessageResponseFailurereason) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *MessageResponseFailurereason) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MessageResponseFailurereason) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MessageResponseFailurereason) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MessageResponseFailurereason) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetPermanent

`func (o *MessageResponseFailurereason) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *MessageResponseFailurereason) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *MessageResponseFailurereason) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.

### HasPermanent

`func (o *MessageResponseFailurereason) HasPermanent() bool`

HasPermanent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


