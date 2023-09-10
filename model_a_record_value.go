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

// checks if the ARecordValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARecordValue{}

// ARecordValue struct for ARecordValue
type ARecordValue struct {
	Ip *string `json:"ip,omitempty"`
	Port NullableInt32 `json:"port,omitempty"`
	Weight NullableInt32 `json:"weight,omitempty"`
	// This key shows itself here if the weight have been changed by Health Check.
	OriginalWeight *int32 `json:"original_weight,omitempty"`
	// ISO 3166 alpha-2 country code
	Country NullableString `json:"country,omitempty"`
}

// NewARecordValue instantiates a new ARecordValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARecordValue() *ARecordValue {
	this := ARecordValue{}
	return &this
}

// NewARecordValueWithDefaults instantiates a new ARecordValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARecordValueWithDefaults() *ARecordValue {
	this := ARecordValue{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ARecordValue) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARecordValue) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ARecordValue) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *ARecordValue) SetIp(v string) {
	o.Ip = &v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ARecordValue) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ARecordValue) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *ARecordValue) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *ARecordValue) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *ARecordValue) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *ARecordValue) UnsetPort() {
	o.Port.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ARecordValue) GetWeight() int32 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret int32
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ARecordValue) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *ARecordValue) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableInt32 and assigns it to the Weight field.
func (o *ARecordValue) SetWeight(v int32) {
	o.Weight.Set(&v)
}
// SetWeightNil sets the value for Weight to be an explicit nil
func (o *ARecordValue) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *ARecordValue) UnsetWeight() {
	o.Weight.Unset()
}

// GetOriginalWeight returns the OriginalWeight field value if set, zero value otherwise.
func (o *ARecordValue) GetOriginalWeight() int32 {
	if o == nil || IsNil(o.OriginalWeight) {
		var ret int32
		return ret
	}
	return *o.OriginalWeight
}

// GetOriginalWeightOk returns a tuple with the OriginalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARecordValue) GetOriginalWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalWeight) {
		return nil, false
	}
	return o.OriginalWeight, true
}

// HasOriginalWeight returns a boolean if a field has been set.
func (o *ARecordValue) HasOriginalWeight() bool {
	if o != nil && !IsNil(o.OriginalWeight) {
		return true
	}

	return false
}

// SetOriginalWeight gets a reference to the given int32 and assigns it to the OriginalWeight field.
func (o *ARecordValue) SetOriginalWeight(v int32) {
	o.OriginalWeight = &v
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ARecordValue) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ARecordValue) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *ARecordValue) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *ARecordValue) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *ARecordValue) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *ARecordValue) UnsetCountry() {
	o.Country.Unset()
}

func (o ARecordValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARecordValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	// skip: original_weight is readOnly
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	return toSerialize, nil
}

type NullableARecordValue struct {
	value *ARecordValue
	isSet bool
}

func (v NullableARecordValue) Get() *ARecordValue {
	return v.value
}

func (v *NullableARecordValue) Set(val *ARecordValue) {
	v.value = val
	v.isSet = true
}

func (v NullableARecordValue) IsSet() bool {
	return v.isSet
}

func (v *NullableARecordValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARecordValue(val *ARecordValue) *NullableARecordValue {
	return &NullableARecordValue{value: val, isSet: true}
}

func (v NullableARecordValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARecordValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


