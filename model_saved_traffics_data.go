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

// checks if the SavedTrafficsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedTrafficsData{}

// SavedTrafficsData struct for SavedTrafficsData
type SavedTrafficsData struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewSavedTrafficsData instantiates a new SavedTrafficsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedTrafficsData() *SavedTrafficsData {
	this := SavedTrafficsData{}
	return &this
}

// NewSavedTrafficsDataWithDefaults instantiates a new SavedTrafficsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedTrafficsDataWithDefaults() *SavedTrafficsData {
	this := SavedTrafficsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SavedTrafficsData) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsData) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SavedTrafficsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *SavedTrafficsData) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o SavedTrafficsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedTrafficsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSavedTrafficsData struct {
	value *SavedTrafficsData
	isSet bool
}

func (v NullableSavedTrafficsData) Get() *SavedTrafficsData {
	return v.value
}

func (v *NullableSavedTrafficsData) Set(val *SavedTrafficsData) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedTrafficsData) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedTrafficsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedTrafficsData(val *SavedTrafficsData) *NullableSavedTrafficsData {
	return &NullableSavedTrafficsData{value: val, isSet: true}
}

func (v NullableSavedTrafficsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedTrafficsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


