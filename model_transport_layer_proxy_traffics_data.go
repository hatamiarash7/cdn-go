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

// checks if the TransportLayerProxyTrafficsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxyTrafficsData{}

// TransportLayerProxyTrafficsData struct for TransportLayerProxyTrafficsData
type TransportLayerProxyTrafficsData struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewTransportLayerProxyTrafficsData instantiates a new TransportLayerProxyTrafficsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxyTrafficsData() *TransportLayerProxyTrafficsData {
	this := TransportLayerProxyTrafficsData{}
	return &this
}

// NewTransportLayerProxyTrafficsDataWithDefaults instantiates a new TransportLayerProxyTrafficsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyTrafficsDataWithDefaults() *TransportLayerProxyTrafficsData {
	this := TransportLayerProxyTrafficsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TransportLayerProxyTrafficsData) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyTrafficsData) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TransportLayerProxyTrafficsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *TransportLayerProxyTrafficsData) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o TransportLayerProxyTrafficsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxyTrafficsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTransportLayerProxyTrafficsData struct {
	value *TransportLayerProxyTrafficsData
	isSet bool
}

func (v NullableTransportLayerProxyTrafficsData) Get() *TransportLayerProxyTrafficsData {
	return v.value
}

func (v *NullableTransportLayerProxyTrafficsData) Set(val *TransportLayerProxyTrafficsData) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyTrafficsData) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyTrafficsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyTrafficsData(val *TransportLayerProxyTrafficsData) *NullableTransportLayerProxyTrafficsData {
	return &NullableTransportLayerProxyTrafficsData{value: val, isSet: true}
}

func (v NullableTransportLayerProxyTrafficsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyTrafficsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


