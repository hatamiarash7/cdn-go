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

// checks if the AttackReportUriData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportUriData{}

// AttackReportUriData struct for AttackReportUriData
type AttackReportUriData struct {
	Data []AttackReportUri `json:"data,omitempty"`
}

// NewAttackReportUriData instantiates a new AttackReportUriData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportUriData() *AttackReportUriData {
	this := AttackReportUriData{}
	return &this
}

// NewAttackReportUriDataWithDefaults instantiates a new AttackReportUriData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportUriDataWithDefaults() *AttackReportUriData {
	this := AttackReportUriData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AttackReportUriData) GetData() []AttackReportUri {
	if o == nil || IsNil(o.Data) {
		var ret []AttackReportUri
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportUriData) GetDataOk() ([]AttackReportUri, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AttackReportUriData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AttackReportUri and assigns it to the Data field.
func (o *AttackReportUriData) SetData(v []AttackReportUri) {
	o.Data = v
}

func (o AttackReportUriData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportUriData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAttackReportUriData struct {
	value *AttackReportUriData
	isSet bool
}

func (v NullableAttackReportUriData) Get() *AttackReportUriData {
	return v.value
}

func (v *NullableAttackReportUriData) Set(val *AttackReportUriData) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportUriData) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportUriData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportUriData(val *AttackReportUriData) *NullableAttackReportUriData {
	return &NullableAttackReportUriData{value: val, isSet: true}
}

func (v NullableAttackReportUriData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportUriData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


