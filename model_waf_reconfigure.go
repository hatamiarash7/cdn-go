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

// checks if the WafReconfigure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafReconfigure{}

// WafReconfigure struct for WafReconfigure
type WafReconfigure struct {
	PresetId *string `json:"preset_id,omitempty"`
}

// NewWafReconfigure instantiates a new WafReconfigure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafReconfigure() *WafReconfigure {
	this := WafReconfigure{}
	return &this
}

// NewWafReconfigureWithDefaults instantiates a new WafReconfigure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafReconfigureWithDefaults() *WafReconfigure {
	this := WafReconfigure{}
	return &this
}

// GetPresetId returns the PresetId field value if set, zero value otherwise.
func (o *WafReconfigure) GetPresetId() string {
	if o == nil || IsNil(o.PresetId) {
		var ret string
		return ret
	}
	return *o.PresetId
}

// GetPresetIdOk returns a tuple with the PresetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafReconfigure) GetPresetIdOk() (*string, bool) {
	if o == nil || IsNil(o.PresetId) {
		return nil, false
	}
	return o.PresetId, true
}

// HasPresetId returns a boolean if a field has been set.
func (o *WafReconfigure) HasPresetId() bool {
	if o != nil && !IsNil(o.PresetId) {
		return true
	}

	return false
}

// SetPresetId gets a reference to the given string and assigns it to the PresetId field.
func (o *WafReconfigure) SetPresetId(v string) {
	o.PresetId = &v
}

func (o WafReconfigure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafReconfigure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PresetId) {
		toSerialize["preset_id"] = o.PresetId
	}
	return toSerialize, nil
}

type NullableWafReconfigure struct {
	value *WafReconfigure
	isSet bool
}

func (v NullableWafReconfigure) Get() *WafReconfigure {
	return v.value
}

func (v *NullableWafReconfigure) Set(val *WafReconfigure) {
	v.value = val
	v.isSet = true
}

func (v NullableWafReconfigure) IsSet() bool {
	return v.isSet
}

func (v *NullableWafReconfigure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafReconfigure(val *WafReconfigure) *NullableWafReconfigure {
	return &NullableWafReconfigure{value: val, isSet: true}
}

func (v NullableWafReconfigure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafReconfigure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


