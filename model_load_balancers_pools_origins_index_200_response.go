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

// checks if the LoadBalancersPoolsOriginsIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancersPoolsOriginsIndex200Response{}

// LoadBalancersPoolsOriginsIndex200Response struct for LoadBalancersPoolsOriginsIndex200Response
type LoadBalancersPoolsOriginsIndex200Response struct {
	Data []LoadBalancerOrigin `json:"data,omitempty"`
}

// NewLoadBalancersPoolsOriginsIndex200Response instantiates a new LoadBalancersPoolsOriginsIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancersPoolsOriginsIndex200Response() *LoadBalancersPoolsOriginsIndex200Response {
	this := LoadBalancersPoolsOriginsIndex200Response{}
	return &this
}

// NewLoadBalancersPoolsOriginsIndex200ResponseWithDefaults instantiates a new LoadBalancersPoolsOriginsIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancersPoolsOriginsIndex200ResponseWithDefaults() *LoadBalancersPoolsOriginsIndex200Response {
	this := LoadBalancersPoolsOriginsIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancersPoolsOriginsIndex200Response) GetData() []LoadBalancerOrigin {
	if o == nil || IsNil(o.Data) {
		var ret []LoadBalancerOrigin
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancersPoolsOriginsIndex200Response) GetDataOk() ([]LoadBalancerOrigin, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancersPoolsOriginsIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []LoadBalancerOrigin and assigns it to the Data field.
func (o *LoadBalancersPoolsOriginsIndex200Response) SetData(v []LoadBalancerOrigin) {
	o.Data = v
}

func (o LoadBalancersPoolsOriginsIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancersPoolsOriginsIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLoadBalancersPoolsOriginsIndex200Response struct {
	value *LoadBalancersPoolsOriginsIndex200Response
	isSet bool
}

func (v NullableLoadBalancersPoolsOriginsIndex200Response) Get() *LoadBalancersPoolsOriginsIndex200Response {
	return v.value
}

func (v *NullableLoadBalancersPoolsOriginsIndex200Response) Set(val *LoadBalancersPoolsOriginsIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancersPoolsOriginsIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancersPoolsOriginsIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancersPoolsOriginsIndex200Response(val *LoadBalancersPoolsOriginsIndex200Response) *NullableLoadBalancersPoolsOriginsIndex200Response {
	return &NullableLoadBalancersPoolsOriginsIndex200Response{value: val, isSet: true}
}

func (v NullableLoadBalancersPoolsOriginsIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancersPoolsOriginsIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


