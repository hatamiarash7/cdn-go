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

// checks if the StatusCodeReportStatisticsStatusCodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeReportStatisticsStatusCodes{}

// StatusCodeReportStatisticsStatusCodes struct for StatusCodeReportStatisticsStatusCodes
type StatusCodeReportStatisticsStatusCodes struct {
	Var2xxSum *int32 `json:"2xx_sum,omitempty"`
	Var3xxSum *int32 `json:"3xx_sum,omitempty"`
	Var4xxSum *int32 `json:"4xx_sum,omitempty"`
	Var5xxSum *int32 `json:"5xx_sum,omitempty"`
}

// NewStatusCodeReportStatisticsStatusCodes instantiates a new StatusCodeReportStatisticsStatusCodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeReportStatisticsStatusCodes() *StatusCodeReportStatisticsStatusCodes {
	this := StatusCodeReportStatisticsStatusCodes{}
	return &this
}

// NewStatusCodeReportStatisticsStatusCodesWithDefaults instantiates a new StatusCodeReportStatisticsStatusCodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeReportStatisticsStatusCodesWithDefaults() *StatusCodeReportStatisticsStatusCodes {
	this := StatusCodeReportStatisticsStatusCodes{}
	return &this
}

// GetVar2xxSum returns the Var2xxSum field value if set, zero value otherwise.
func (o *StatusCodeReportStatisticsStatusCodes) GetVar2xxSum() int32 {
	if o == nil || IsNil(o.Var2xxSum) {
		var ret int32
		return ret
	}
	return *o.Var2xxSum
}

// GetVar2xxSumOk returns a tuple with the Var2xxSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportStatisticsStatusCodes) GetVar2xxSumOk() (*int32, bool) {
	if o == nil || IsNil(o.Var2xxSum) {
		return nil, false
	}
	return o.Var2xxSum, true
}

// HasVar2xxSum returns a boolean if a field has been set.
func (o *StatusCodeReportStatisticsStatusCodes) HasVar2xxSum() bool {
	if o != nil && !IsNil(o.Var2xxSum) {
		return true
	}

	return false
}

// SetVar2xxSum gets a reference to the given int32 and assigns it to the Var2xxSum field.
func (o *StatusCodeReportStatisticsStatusCodes) SetVar2xxSum(v int32) {
	o.Var2xxSum = &v
}

// GetVar3xxSum returns the Var3xxSum field value if set, zero value otherwise.
func (o *StatusCodeReportStatisticsStatusCodes) GetVar3xxSum() int32 {
	if o == nil || IsNil(o.Var3xxSum) {
		var ret int32
		return ret
	}
	return *o.Var3xxSum
}

// GetVar3xxSumOk returns a tuple with the Var3xxSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportStatisticsStatusCodes) GetVar3xxSumOk() (*int32, bool) {
	if o == nil || IsNil(o.Var3xxSum) {
		return nil, false
	}
	return o.Var3xxSum, true
}

// HasVar3xxSum returns a boolean if a field has been set.
func (o *StatusCodeReportStatisticsStatusCodes) HasVar3xxSum() bool {
	if o != nil && !IsNil(o.Var3xxSum) {
		return true
	}

	return false
}

// SetVar3xxSum gets a reference to the given int32 and assigns it to the Var3xxSum field.
func (o *StatusCodeReportStatisticsStatusCodes) SetVar3xxSum(v int32) {
	o.Var3xxSum = &v
}

// GetVar4xxSum returns the Var4xxSum field value if set, zero value otherwise.
func (o *StatusCodeReportStatisticsStatusCodes) GetVar4xxSum() int32 {
	if o == nil || IsNil(o.Var4xxSum) {
		var ret int32
		return ret
	}
	return *o.Var4xxSum
}

// GetVar4xxSumOk returns a tuple with the Var4xxSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportStatisticsStatusCodes) GetVar4xxSumOk() (*int32, bool) {
	if o == nil || IsNil(o.Var4xxSum) {
		return nil, false
	}
	return o.Var4xxSum, true
}

// HasVar4xxSum returns a boolean if a field has been set.
func (o *StatusCodeReportStatisticsStatusCodes) HasVar4xxSum() bool {
	if o != nil && !IsNil(o.Var4xxSum) {
		return true
	}

	return false
}

// SetVar4xxSum gets a reference to the given int32 and assigns it to the Var4xxSum field.
func (o *StatusCodeReportStatisticsStatusCodes) SetVar4xxSum(v int32) {
	o.Var4xxSum = &v
}

// GetVar5xxSum returns the Var5xxSum field value if set, zero value otherwise.
func (o *StatusCodeReportStatisticsStatusCodes) GetVar5xxSum() int32 {
	if o == nil || IsNil(o.Var5xxSum) {
		var ret int32
		return ret
	}
	return *o.Var5xxSum
}

// GetVar5xxSumOk returns a tuple with the Var5xxSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportStatisticsStatusCodes) GetVar5xxSumOk() (*int32, bool) {
	if o == nil || IsNil(o.Var5xxSum) {
		return nil, false
	}
	return o.Var5xxSum, true
}

// HasVar5xxSum returns a boolean if a field has been set.
func (o *StatusCodeReportStatisticsStatusCodes) HasVar5xxSum() bool {
	if o != nil && !IsNil(o.Var5xxSum) {
		return true
	}

	return false
}

// SetVar5xxSum gets a reference to the given int32 and assigns it to the Var5xxSum field.
func (o *StatusCodeReportStatisticsStatusCodes) SetVar5xxSum(v int32) {
	o.Var5xxSum = &v
}

func (o StatusCodeReportStatisticsStatusCodes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeReportStatisticsStatusCodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var2xxSum) {
		toSerialize["2xx_sum"] = o.Var2xxSum
	}
	if !IsNil(o.Var3xxSum) {
		toSerialize["3xx_sum"] = o.Var3xxSum
	}
	if !IsNil(o.Var4xxSum) {
		toSerialize["4xx_sum"] = o.Var4xxSum
	}
	if !IsNil(o.Var5xxSum) {
		toSerialize["5xx_sum"] = o.Var5xxSum
	}
	return toSerialize, nil
}

type NullableStatusCodeReportStatisticsStatusCodes struct {
	value *StatusCodeReportStatisticsStatusCodes
	isSet bool
}

func (v NullableStatusCodeReportStatisticsStatusCodes) Get() *StatusCodeReportStatisticsStatusCodes {
	return v.value
}

func (v *NullableStatusCodeReportStatisticsStatusCodes) Set(val *StatusCodeReportStatisticsStatusCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeReportStatisticsStatusCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeReportStatisticsStatusCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeReportStatisticsStatusCodes(val *StatusCodeReportStatisticsStatusCodes) *NullableStatusCodeReportStatisticsStatusCodes {
	return &NullableStatusCodeReportStatisticsStatusCodes{value: val, isSet: true}
}

func (v NullableStatusCodeReportStatisticsStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeReportStatisticsStatusCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


