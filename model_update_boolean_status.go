/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.103.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the UpdateBooleanStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBooleanStatus{}

// UpdateBooleanStatus struct for UpdateBooleanStatus
type UpdateBooleanStatus struct {
	Status *bool `json:"status,omitempty"`
}

// NewUpdateBooleanStatus instantiates a new UpdateBooleanStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBooleanStatus() *UpdateBooleanStatus {
	this := UpdateBooleanStatus{}
	return &this
}

// NewUpdateBooleanStatusWithDefaults instantiates a new UpdateBooleanStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBooleanStatusWithDefaults() *UpdateBooleanStatus {
	this := UpdateBooleanStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateBooleanStatus) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateBooleanStatus) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateBooleanStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *UpdateBooleanStatus) SetStatus(v bool) {
	o.Status = &v
}

func (o UpdateBooleanStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBooleanStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUpdateBooleanStatus struct {
	value *UpdateBooleanStatus
	isSet bool
}

func (v NullableUpdateBooleanStatus) Get() *UpdateBooleanStatus {
	return v.value
}

func (v *NullableUpdateBooleanStatus) Set(val *UpdateBooleanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBooleanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBooleanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBooleanStatus(val *UpdateBooleanStatus) *NullableUpdateBooleanStatus {
	return &NullableUpdateBooleanStatus{value: val, isSet: true}
}

func (v NullableUpdateBooleanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBooleanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


