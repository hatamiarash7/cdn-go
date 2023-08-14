/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.104.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the AccelerationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccelerationResponse{}

// AccelerationResponse struct for AccelerationResponse
type AccelerationResponse struct {
	Data *Acceleration `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewAccelerationResponse instantiates a new AccelerationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccelerationResponse() *AccelerationResponse {
	this := AccelerationResponse{}
	return &this
}

// NewAccelerationResponseWithDefaults instantiates a new AccelerationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccelerationResponseWithDefaults() *AccelerationResponse {
	this := AccelerationResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AccelerationResponse) GetData() Acceleration {
	if o == nil || IsNil(o.Data) {
		var ret Acceleration
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccelerationResponse) GetDataOk() (*Acceleration, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AccelerationResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Acceleration and assigns it to the Data field.
func (o *AccelerationResponse) SetData(v Acceleration) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccelerationResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccelerationResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *AccelerationResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *AccelerationResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *AccelerationResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *AccelerationResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o AccelerationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccelerationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableAccelerationResponse struct {
	value *AccelerationResponse
	isSet bool
}

func (v NullableAccelerationResponse) Get() *AccelerationResponse {
	return v.value
}

func (v *NullableAccelerationResponse) Set(val *AccelerationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccelerationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccelerationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccelerationResponse(val *AccelerationResponse) *NullableAccelerationResponse {
	return &NullableAccelerationResponse{value: val, isSet: true}
}

func (v NullableAccelerationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccelerationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


