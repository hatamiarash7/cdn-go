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

// checks if the EmailForwardingStatsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingStatsData{}

// EmailForwardingStatsData struct for EmailForwardingStatsData
type EmailForwardingStatsData struct {
	Data *EmailForwardingStats `json:"data,omitempty"`
}

// NewEmailForwardingStatsData instantiates a new EmailForwardingStatsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingStatsData() *EmailForwardingStatsData {
	this := EmailForwardingStatsData{}
	return &this
}

// NewEmailForwardingStatsDataWithDefaults instantiates a new EmailForwardingStatsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingStatsDataWithDefaults() *EmailForwardingStatsData {
	this := EmailForwardingStatsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailForwardingStatsData) GetData() EmailForwardingStats {
	if o == nil || IsNil(o.Data) {
		var ret EmailForwardingStats
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingStatsData) GetDataOk() (*EmailForwardingStats, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailForwardingStatsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EmailForwardingStats and assigns it to the Data field.
func (o *EmailForwardingStatsData) SetData(v EmailForwardingStats) {
	o.Data = &v
}

func (o EmailForwardingStatsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingStatsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEmailForwardingStatsData struct {
	value *EmailForwardingStatsData
	isSet bool
}

func (v NullableEmailForwardingStatsData) Get() *EmailForwardingStatsData {
	return v.value
}

func (v *NullableEmailForwardingStatsData) Set(val *EmailForwardingStatsData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingStatsData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingStatsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingStatsData(val *EmailForwardingStatsData) *NullableEmailForwardingStatsData {
	return &NullableEmailForwardingStatsData{value: val, isSet: true}
}

func (v NullableEmailForwardingStatsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingStatsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


