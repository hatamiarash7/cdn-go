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

// checks if the LoadBalancersPoolsIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancersPoolsIndex200Response{}

// LoadBalancersPoolsIndex200Response struct for LoadBalancersPoolsIndex200Response
type LoadBalancersPoolsIndex200Response struct {
	Data []LoadBalancerPool `json:"data,omitempty"`
}

// NewLoadBalancersPoolsIndex200Response instantiates a new LoadBalancersPoolsIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancersPoolsIndex200Response() *LoadBalancersPoolsIndex200Response {
	this := LoadBalancersPoolsIndex200Response{}
	return &this
}

// NewLoadBalancersPoolsIndex200ResponseWithDefaults instantiates a new LoadBalancersPoolsIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancersPoolsIndex200ResponseWithDefaults() *LoadBalancersPoolsIndex200Response {
	this := LoadBalancersPoolsIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancersPoolsIndex200Response) GetData() []LoadBalancerPool {
	if o == nil || IsNil(o.Data) {
		var ret []LoadBalancerPool
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancersPoolsIndex200Response) GetDataOk() ([]LoadBalancerPool, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancersPoolsIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []LoadBalancerPool and assigns it to the Data field.
func (o *LoadBalancersPoolsIndex200Response) SetData(v []LoadBalancerPool) {
	o.Data = v
}

func (o LoadBalancersPoolsIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancersPoolsIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLoadBalancersPoolsIndex200Response struct {
	value *LoadBalancersPoolsIndex200Response
	isSet bool
}

func (v NullableLoadBalancersPoolsIndex200Response) Get() *LoadBalancersPoolsIndex200Response {
	return v.value
}

func (v *NullableLoadBalancersPoolsIndex200Response) Set(val *LoadBalancersPoolsIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancersPoolsIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancersPoolsIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancersPoolsIndex200Response(val *LoadBalancersPoolsIndex200Response) *NullableLoadBalancersPoolsIndex200Response {
	return &NullableLoadBalancersPoolsIndex200Response{value: val, isSet: true}
}

func (v NullableLoadBalancersPoolsIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancersPoolsIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


