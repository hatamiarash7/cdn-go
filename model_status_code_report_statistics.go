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

// checks if the StatusCodeReportStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeReportStatistics{}

// StatusCodeReportStatistics struct for StatusCodeReportStatistics
type StatusCodeReportStatistics struct {
	StatusCodes *StatusCodeReportStatisticsStatusCodes `json:"status_codes,omitempty"`
}

// NewStatusCodeReportStatistics instantiates a new StatusCodeReportStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeReportStatistics() *StatusCodeReportStatistics {
	this := StatusCodeReportStatistics{}
	return &this
}

// NewStatusCodeReportStatisticsWithDefaults instantiates a new StatusCodeReportStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeReportStatisticsWithDefaults() *StatusCodeReportStatistics {
	this := StatusCodeReportStatistics{}
	return &this
}

// GetStatusCodes returns the StatusCodes field value if set, zero value otherwise.
func (o *StatusCodeReportStatistics) GetStatusCodes() StatusCodeReportStatisticsStatusCodes {
	if o == nil || IsNil(o.StatusCodes) {
		var ret StatusCodeReportStatisticsStatusCodes
		return ret
	}
	return *o.StatusCodes
}

// GetStatusCodesOk returns a tuple with the StatusCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportStatistics) GetStatusCodesOk() (*StatusCodeReportStatisticsStatusCodes, bool) {
	if o == nil || IsNil(o.StatusCodes) {
		return nil, false
	}
	return o.StatusCodes, true
}

// HasStatusCodes returns a boolean if a field has been set.
func (o *StatusCodeReportStatistics) HasStatusCodes() bool {
	if o != nil && !IsNil(o.StatusCodes) {
		return true
	}

	return false
}

// SetStatusCodes gets a reference to the given StatusCodeReportStatisticsStatusCodes and assigns it to the StatusCodes field.
func (o *StatusCodeReportStatistics) SetStatusCodes(v StatusCodeReportStatisticsStatusCodes) {
	o.StatusCodes = &v
}

func (o StatusCodeReportStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeReportStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusCodes) {
		toSerialize["status_codes"] = o.StatusCodes
	}
	return toSerialize, nil
}

type NullableStatusCodeReportStatistics struct {
	value *StatusCodeReportStatistics
	isSet bool
}

func (v NullableStatusCodeReportStatistics) Get() *StatusCodeReportStatistics {
	return v.value
}

func (v *NullableStatusCodeReportStatistics) Set(val *StatusCodeReportStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeReportStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeReportStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeReportStatistics(val *StatusCodeReportStatistics) *NullableStatusCodeReportStatistics {
	return &NullableStatusCodeReportStatistics{value: val, isSet: true}
}

func (v NullableStatusCodeReportStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeReportStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


