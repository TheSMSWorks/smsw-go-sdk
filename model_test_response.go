/*
The SMS Works API

The SMS Works provides a low-cost, reliable SMS API for developers. Pay only for delivered texts, all failed UK messages are refunded.

API version: 1.9.0
Contact: support@thesmsworks.co.uk
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the TestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestResponse{}

// TestResponse struct for TestResponse
type TestResponse struct {
	Message string `json:"message"`
}

type _TestResponse TestResponse

// NewTestResponse instantiates a new TestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResponse(message string) *TestResponse {
	this := TestResponse{}
	this.Message = message
	return &this
}

// NewTestResponseWithDefaults instantiates a new TestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResponseWithDefaults() *TestResponse {
	this := TestResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *TestResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TestResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TestResponse) SetMessage(v string) {
	o.Message = v
}

func (o TestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *TestResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTestResponse := _TestResponse{}

	err = json.Unmarshal(bytes, &varTestResponse)

	if err != nil {
		return err
	}

	*o = TestResponse(varTestResponse)

	return err
}

type NullableTestResponse struct {
	value *TestResponse
	isSet bool
}

func (v NullableTestResponse) Get() *TestResponse {
	return v.value
}

func (v *NullableTestResponse) Set(val *TestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResponse(val *TestResponse) *NullableTestResponse {
	return &NullableTestResponse{value: val, isSet: true}
}

func (v NullableTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


