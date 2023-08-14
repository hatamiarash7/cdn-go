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

// checks if the PageRuleDiffData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageRuleDiffData{}

// PageRuleDiffData struct for PageRuleDiffData
type PageRuleDiffData struct {
	Data *PageRuleDiff `json:"data,omitempty"`
}

// NewPageRuleDiffData instantiates a new PageRuleDiffData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRuleDiffData() *PageRuleDiffData {
	this := PageRuleDiffData{}
	return &this
}

// NewPageRuleDiffDataWithDefaults instantiates a new PageRuleDiffData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRuleDiffDataWithDefaults() *PageRuleDiffData {
	this := PageRuleDiffData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PageRuleDiffData) GetData() PageRuleDiff {
	if o == nil || IsNil(o.Data) {
		var ret PageRuleDiff
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleDiffData) GetDataOk() (*PageRuleDiff, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PageRuleDiffData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PageRuleDiff and assigns it to the Data field.
func (o *PageRuleDiffData) SetData(v PageRuleDiff) {
	o.Data = &v
}

func (o PageRuleDiffData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageRuleDiffData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePageRuleDiffData struct {
	value *PageRuleDiffData
	isSet bool
}

func (v NullablePageRuleDiffData) Get() *PageRuleDiffData {
	return v.value
}

func (v *NullablePageRuleDiffData) Set(val *PageRuleDiffData) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRuleDiffData) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRuleDiffData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRuleDiffData(val *PageRuleDiffData) *NullablePageRuleDiffData {
	return &NullablePageRuleDiffData{value: val, isSet: true}
}

func (v NullablePageRuleDiffData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRuleDiffData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


