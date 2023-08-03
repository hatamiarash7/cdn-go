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

// checks if the TrafficCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficCharts{}

// TrafficCharts struct for TrafficCharts
type TrafficCharts struct {
	Requests *TrafficChartsRequests `json:"requests,omitempty"`
	Traffics *TrafficChartsTraffics `json:"traffics,omitempty"`
}

// NewTrafficCharts instantiates a new TrafficCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficCharts() *TrafficCharts {
	this := TrafficCharts{}
	return &this
}

// NewTrafficChartsWithDefaults instantiates a new TrafficCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficChartsWithDefaults() *TrafficCharts {
	this := TrafficCharts{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *TrafficCharts) GetRequests() TrafficChartsRequests {
	if o == nil || IsNil(o.Requests) {
		var ret TrafficChartsRequests
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharts) GetRequestsOk() (*TrafficChartsRequests, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *TrafficCharts) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given TrafficChartsRequests and assigns it to the Requests field.
func (o *TrafficCharts) SetRequests(v TrafficChartsRequests) {
	o.Requests = &v
}

// GetTraffics returns the Traffics field value if set, zero value otherwise.
func (o *TrafficCharts) GetTraffics() TrafficChartsTraffics {
	if o == nil || IsNil(o.Traffics) {
		var ret TrafficChartsTraffics
		return ret
	}
	return *o.Traffics
}

// GetTrafficsOk returns a tuple with the Traffics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficCharts) GetTrafficsOk() (*TrafficChartsTraffics, bool) {
	if o == nil || IsNil(o.Traffics) {
		return nil, false
	}
	return o.Traffics, true
}

// HasTraffics returns a boolean if a field has been set.
func (o *TrafficCharts) HasTraffics() bool {
	if o != nil && !IsNil(o.Traffics) {
		return true
	}

	return false
}

// SetTraffics gets a reference to the given TrafficChartsTraffics and assigns it to the Traffics field.
func (o *TrafficCharts) SetTraffics(v TrafficChartsTraffics) {
	o.Traffics = &v
}

func (o TrafficCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	if !IsNil(o.Traffics) {
		toSerialize["traffics"] = o.Traffics
	}
	return toSerialize, nil
}

type NullableTrafficCharts struct {
	value *TrafficCharts
	isSet bool
}

func (v NullableTrafficCharts) Get() *TrafficCharts {
	return v.value
}

func (v *NullableTrafficCharts) Set(val *TrafficCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficCharts(val *TrafficCharts) *NullableTrafficCharts {
	return &NullableTrafficCharts{value: val, isSet: true}
}

func (v NullableTrafficCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


