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

// checks if the DnsSecStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsSecStatus{}

// DnsSecStatus struct for DnsSecStatus
type DnsSecStatus struct {
	Enable bool `json:"enable"`
}

// NewDnsSecStatus instantiates a new DnsSecStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsSecStatus(enable bool) *DnsSecStatus {
	this := DnsSecStatus{}
	this.Enable = enable
	return &this
}

// NewDnsSecStatusWithDefaults instantiates a new DnsSecStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsSecStatusWithDefaults() *DnsSecStatus {
	this := DnsSecStatus{}
	return &this
}

// GetEnable returns the Enable field value
func (o *DnsSecStatus) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *DnsSecStatus) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *DnsSecStatus) SetEnable(v bool) {
	o.Enable = v
}

func (o DnsSecStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsSecStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable"] = o.Enable
	return toSerialize, nil
}

type NullableDnsSecStatus struct {
	value *DnsSecStatus
	isSet bool
}

func (v NullableDnsSecStatus) Get() *DnsSecStatus {
	return v.value
}

func (v *NullableDnsSecStatus) Set(val *DnsSecStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsSecStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsSecStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsSecStatus(val *DnsSecStatus) *NullableDnsSecStatus {
	return &NullableDnsSecStatus{value: val, isSet: true}
}

func (v NullableDnsSecStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsSecStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


