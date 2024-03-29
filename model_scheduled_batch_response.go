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

// checks if the ScheduledBatchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledBatchResponse{}

// ScheduledBatchResponse struct for ScheduledBatchResponse
type ScheduledBatchResponse struct {
	Batchid string `json:"batchid"`
	Status string `json:"status"`
}

type _ScheduledBatchResponse ScheduledBatchResponse

// NewScheduledBatchResponse instantiates a new ScheduledBatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledBatchResponse(batchid string, status string) *ScheduledBatchResponse {
	this := ScheduledBatchResponse{}
	this.Batchid = batchid
	this.Status = status
	return &this
}

// NewScheduledBatchResponseWithDefaults instantiates a new ScheduledBatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledBatchResponseWithDefaults() *ScheduledBatchResponse {
	this := ScheduledBatchResponse{}
	return &this
}

// GetBatchid returns the Batchid field value
func (o *ScheduledBatchResponse) GetBatchid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Batchid
}

// GetBatchidOk returns a tuple with the Batchid field value
// and a boolean to check if the value has been set.
func (o *ScheduledBatchResponse) GetBatchidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Batchid, true
}

// SetBatchid sets field value
func (o *ScheduledBatchResponse) SetBatchid(v string) {
	o.Batchid = v
}

// GetStatus returns the Status field value
func (o *ScheduledBatchResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ScheduledBatchResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ScheduledBatchResponse) SetStatus(v string) {
	o.Status = v
}

func (o ScheduledBatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledBatchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["batchid"] = o.Batchid
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ScheduledBatchResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"batchid",
		"status",
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

	varScheduledBatchResponse := _ScheduledBatchResponse{}

	err = json.Unmarshal(bytes, &varScheduledBatchResponse)

	if err != nil {
		return err
	}

	*o = ScheduledBatchResponse(varScheduledBatchResponse)

	return err
}

type NullableScheduledBatchResponse struct {
	value *ScheduledBatchResponse
	isSet bool
}

func (v NullableScheduledBatchResponse) Get() *ScheduledBatchResponse {
	return v.value
}

func (v *NullableScheduledBatchResponse) Set(val *ScheduledBatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledBatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledBatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledBatchResponse(val *ScheduledBatchResponse) *NullableScheduledBatchResponse {
	return &NullableScheduledBatchResponse{value: val, isSet: true}
}

func (v NullableScheduledBatchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledBatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


