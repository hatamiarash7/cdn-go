/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the AttackReportMapCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportMapCharts{}

// AttackReportMapCharts struct for AttackReportMapCharts
type AttackReportMapCharts struct {
	Attacks *map[string]AttackReportMapChartsAttacksValue `json:"attacks,omitempty"`
}

// NewAttackReportMapCharts instantiates a new AttackReportMapCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportMapCharts() *AttackReportMapCharts {
	this := AttackReportMapCharts{}
	return &this
}

// NewAttackReportMapChartsWithDefaults instantiates a new AttackReportMapCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportMapChartsWithDefaults() *AttackReportMapCharts {
	this := AttackReportMapCharts{}
	return &this
}

// GetAttacks returns the Attacks field value if set, zero value otherwise.
func (o *AttackReportMapCharts) GetAttacks() map[string]AttackReportMapChartsAttacksValue {
	if o == nil || IsNil(o.Attacks) {
		var ret map[string]AttackReportMapChartsAttacksValue
		return ret
	}
	return *o.Attacks
}

// GetAttacksOk returns a tuple with the Attacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapCharts) GetAttacksOk() (*map[string]AttackReportMapChartsAttacksValue, bool) {
	if o == nil || IsNil(o.Attacks) {
		return nil, false
	}
	return o.Attacks, true
}

// HasAttacks returns a boolean if a field has been set.
func (o *AttackReportMapCharts) HasAttacks() bool {
	if o != nil && !IsNil(o.Attacks) {
		return true
	}

	return false
}

// SetAttacks gets a reference to the given map[string]AttackReportMapChartsAttacksValue and assigns it to the Attacks field.
func (o *AttackReportMapCharts) SetAttacks(v map[string]AttackReportMapChartsAttacksValue) {
	o.Attacks = &v
}

func (o AttackReportMapCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportMapCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attacks) {
		toSerialize["attacks"] = o.Attacks
	}
	return toSerialize, nil
}

type NullableAttackReportMapCharts struct {
	value *AttackReportMapCharts
	isSet bool
}

func (v NullableAttackReportMapCharts) Get() *AttackReportMapCharts {
	return v.value
}

func (v *NullableAttackReportMapCharts) Set(val *AttackReportMapCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportMapCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportMapCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportMapCharts(val *AttackReportMapCharts) *NullableAttackReportMapCharts {
	return &NullableAttackReportMapCharts{value: val, isSet: true}
}

func (v NullableAttackReportMapCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportMapCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


