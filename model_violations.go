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

// checks if the Violations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Violations{}

// Violations struct for Violations
type Violations struct {
	Violations *ViolationsViolations `json:"violations,omitempty"`
}

// NewViolations instantiates a new Violations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViolations() *Violations {
	this := Violations{}
	return &this
}

// NewViolationsWithDefaults instantiates a new Violations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViolationsWithDefaults() *Violations {
	this := Violations{}
	return &this
}

// GetViolations returns the Violations field value if set, zero value otherwise.
func (o *Violations) GetViolations() ViolationsViolations {
	if o == nil || IsNil(o.Violations) {
		var ret ViolationsViolations
		return ret
	}
	return *o.Violations
}

// GetViolationsOk returns a tuple with the Violations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Violations) GetViolationsOk() (*ViolationsViolations, bool) {
	if o == nil || IsNil(o.Violations) {
		return nil, false
	}
	return o.Violations, true
}

// HasViolations returns a boolean if a field has been set.
func (o *Violations) HasViolations() bool {
	if o != nil && !IsNil(o.Violations) {
		return true
	}

	return false
}

// SetViolations gets a reference to the given ViolationsViolations and assigns it to the Violations field.
func (o *Violations) SetViolations(v ViolationsViolations) {
	o.Violations = &v
}

func (o Violations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Violations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Violations) {
		toSerialize["violations"] = o.Violations
	}
	return toSerialize, nil
}

type NullableViolations struct {
	value *Violations
	isSet bool
}

func (v NullableViolations) Get() *Violations {
	return v.value
}

func (v *NullableViolations) Set(val *Violations) {
	v.value = val
	v.isSet = true
}

func (v NullableViolations) IsSet() bool {
	return v.isSet
}

func (v *NullableViolations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViolations(val *Violations) *NullableViolations {
	return &NullableViolations{value: val, isSet: true}
}

func (v NullableViolations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViolations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


