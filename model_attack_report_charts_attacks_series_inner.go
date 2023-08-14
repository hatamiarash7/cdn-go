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

// checks if the AttackReportChartsAttacksSeriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportChartsAttacksSeriesInner{}

// AttackReportChartsAttacksSeriesInner struct for AttackReportChartsAttacksSeriesInner
type AttackReportChartsAttacksSeriesInner struct {
	Name *string `json:"name,omitempty"`
	Data []int32 `json:"data,omitempty"`
}

// NewAttackReportChartsAttacksSeriesInner instantiates a new AttackReportChartsAttacksSeriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportChartsAttacksSeriesInner() *AttackReportChartsAttacksSeriesInner {
	this := AttackReportChartsAttacksSeriesInner{}
	return &this
}

// NewAttackReportChartsAttacksSeriesInnerWithDefaults instantiates a new AttackReportChartsAttacksSeriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportChartsAttacksSeriesInnerWithDefaults() *AttackReportChartsAttacksSeriesInner {
	this := AttackReportChartsAttacksSeriesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttackReportChartsAttacksSeriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportChartsAttacksSeriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttackReportChartsAttacksSeriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttackReportChartsAttacksSeriesInner) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AttackReportChartsAttacksSeriesInner) GetData() []int32 {
	if o == nil || IsNil(o.Data) {
		var ret []int32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportChartsAttacksSeriesInner) GetDataOk() ([]int32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AttackReportChartsAttacksSeriesInner) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []int32 and assigns it to the Data field.
func (o *AttackReportChartsAttacksSeriesInner) SetData(v []int32) {
	o.Data = v
}

func (o AttackReportChartsAttacksSeriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportChartsAttacksSeriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAttackReportChartsAttacksSeriesInner struct {
	value *AttackReportChartsAttacksSeriesInner
	isSet bool
}

func (v NullableAttackReportChartsAttacksSeriesInner) Get() *AttackReportChartsAttacksSeriesInner {
	return v.value
}

func (v *NullableAttackReportChartsAttacksSeriesInner) Set(val *AttackReportChartsAttacksSeriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportChartsAttacksSeriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportChartsAttacksSeriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportChartsAttacksSeriesInner(val *AttackReportChartsAttacksSeriesInner) *NullableAttackReportChartsAttacksSeriesInner {
	return &NullableAttackReportChartsAttacksSeriesInner{value: val, isSet: true}
}

func (v NullableAttackReportChartsAttacksSeriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportChartsAttacksSeriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


