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

// checks if the ReportsErrorLogDetails200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsErrorLogDetails200Response{}

// ReportsErrorLogDetails200Response struct for ReportsErrorLogDetails200Response
type ReportsErrorLogDetails200Response struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewReportsErrorLogDetails200Response instantiates a new ReportsErrorLogDetails200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsErrorLogDetails200Response() *ReportsErrorLogDetails200Response {
	this := ReportsErrorLogDetails200Response{}
	return &this
}

// NewReportsErrorLogDetails200ResponseWithDefaults instantiates a new ReportsErrorLogDetails200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsErrorLogDetails200ResponseWithDefaults() *ReportsErrorLogDetails200Response {
	this := ReportsErrorLogDetails200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReportsErrorLogDetails200Response) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsErrorLogDetails200Response) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReportsErrorLogDetails200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *ReportsErrorLogDetails200Response) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o ReportsErrorLogDetails200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsErrorLogDetails200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableReportsErrorLogDetails200Response struct {
	value *ReportsErrorLogDetails200Response
	isSet bool
}

func (v NullableReportsErrorLogDetails200Response) Get() *ReportsErrorLogDetails200Response {
	return v.value
}

func (v *NullableReportsErrorLogDetails200Response) Set(val *ReportsErrorLogDetails200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsErrorLogDetails200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsErrorLogDetails200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsErrorLogDetails200Response(val *ReportsErrorLogDetails200Response) *NullableReportsErrorLogDetails200Response {
	return &NullableReportsErrorLogDetails200Response{value: val, isSet: true}
}

func (v NullableReportsErrorLogDetails200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsErrorLogDetails200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


