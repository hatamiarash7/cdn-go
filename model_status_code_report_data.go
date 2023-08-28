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

// checks if the StatusCodeReportData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeReportData{}

// StatusCodeReportData struct for StatusCodeReportData
type StatusCodeReportData struct {
	Data *StatusCodeReport `json:"data,omitempty"`
}

// NewStatusCodeReportData instantiates a new StatusCodeReportData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeReportData() *StatusCodeReportData {
	this := StatusCodeReportData{}
	return &this
}

// NewStatusCodeReportDataWithDefaults instantiates a new StatusCodeReportData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeReportDataWithDefaults() *StatusCodeReportData {
	this := StatusCodeReportData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StatusCodeReportData) GetData() StatusCodeReport {
	if o == nil || IsNil(o.Data) {
		var ret StatusCodeReport
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportData) GetDataOk() (*StatusCodeReport, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StatusCodeReportData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given StatusCodeReport and assigns it to the Data field.
func (o *StatusCodeReportData) SetData(v StatusCodeReport) {
	o.Data = &v
}

func (o StatusCodeReportData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeReportData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableStatusCodeReportData struct {
	value *StatusCodeReportData
	isSet bool
}

func (v NullableStatusCodeReportData) Get() *StatusCodeReportData {
	return v.value
}

func (v *NullableStatusCodeReportData) Set(val *StatusCodeReportData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeReportData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeReportData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeReportData(val *StatusCodeReportData) *NullableStatusCodeReportData {
	return &NullableStatusCodeReportData{value: val, isSet: true}
}

func (v NullableStatusCodeReportData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeReportData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


