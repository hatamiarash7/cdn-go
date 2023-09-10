/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.108.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the CustomPagesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomPagesData{}

// CustomPagesData struct for CustomPagesData
type CustomPagesData struct {
	Data *CustomPages `json:"data,omitempty"`
}

// NewCustomPagesData instantiates a new CustomPagesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPagesData() *CustomPagesData {
	this := CustomPagesData{}
	return &this
}

// NewCustomPagesDataWithDefaults instantiates a new CustomPagesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPagesDataWithDefaults() *CustomPagesData {
	this := CustomPagesData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomPagesData) GetData() CustomPages {
	if o == nil || IsNil(o.Data) {
		var ret CustomPages
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPagesData) GetDataOk() (*CustomPages, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomPagesData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomPages and assigns it to the Data field.
func (o *CustomPagesData) SetData(v CustomPages) {
	o.Data = &v
}

func (o CustomPagesData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomPagesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCustomPagesData struct {
	value *CustomPagesData
	isSet bool
}

func (v NullableCustomPagesData) Get() *CustomPagesData {
	return v.value
}

func (v *NullableCustomPagesData) Set(val *CustomPagesData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPagesData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPagesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPagesData(val *CustomPagesData) *NullableCustomPagesData {
	return &NullableCustomPagesData{value: val, isSet: true}
}

func (v NullableCustomPagesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPagesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


