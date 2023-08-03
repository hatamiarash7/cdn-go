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

// checks if the MapTrafficsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapTrafficsData{}

// MapTrafficsData struct for MapTrafficsData
type MapTrafficsData struct {
	Data *map[string]interface{} `json:"data,omitempty"`
	Lists []CountryList `json:"lists,omitempty"`
}

// NewMapTrafficsData instantiates a new MapTrafficsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapTrafficsData() *MapTrafficsData {
	this := MapTrafficsData{}
	return &this
}

// NewMapTrafficsDataWithDefaults instantiates a new MapTrafficsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapTrafficsDataWithDefaults() *MapTrafficsData {
	this := MapTrafficsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MapTrafficsData) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapTrafficsData) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MapTrafficsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *MapTrafficsData) SetData(v map[string]interface{}) {
	o.Data = &v
}

// GetLists returns the Lists field value if set, zero value otherwise.
func (o *MapTrafficsData) GetLists() []CountryList {
	if o == nil || IsNil(o.Lists) {
		var ret []CountryList
		return ret
	}
	return o.Lists
}

// GetListsOk returns a tuple with the Lists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MapTrafficsData) GetListsOk() ([]CountryList, bool) {
	if o == nil || IsNil(o.Lists) {
		return nil, false
	}
	return o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *MapTrafficsData) HasLists() bool {
	if o != nil && !IsNil(o.Lists) {
		return true
	}

	return false
}

// SetLists gets a reference to the given []CountryList and assigns it to the Lists field.
func (o *MapTrafficsData) SetLists(v []CountryList) {
	o.Lists = v
}

func (o MapTrafficsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapTrafficsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Lists) {
		toSerialize["lists"] = o.Lists
	}
	return toSerialize, nil
}

type NullableMapTrafficsData struct {
	value *MapTrafficsData
	isSet bool
}

func (v NullableMapTrafficsData) Get() *MapTrafficsData {
	return v.value
}

func (v *NullableMapTrafficsData) Set(val *MapTrafficsData) {
	v.value = val
	v.isSet = true
}

func (v NullableMapTrafficsData) IsSet() bool {
	return v.isSet
}

func (v *NullableMapTrafficsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapTrafficsData(val *MapTrafficsData) *NullableMapTrafficsData {
	return &NullableMapTrafficsData{value: val, isSet: true}
}

func (v NullableMapTrafficsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapTrafficsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


