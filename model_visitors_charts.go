/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.108.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the VisitorsCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisitorsCharts{}

// VisitorsCharts struct for VisitorsCharts
type VisitorsCharts struct {
	Visitors *VisitorsChartsVisitors `json:"visitors,omitempty"`
}

// NewVisitorsCharts instantiates a new VisitorsCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisitorsCharts() *VisitorsCharts {
	this := VisitorsCharts{}
	return &this
}

// NewVisitorsChartsWithDefaults instantiates a new VisitorsCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisitorsChartsWithDefaults() *VisitorsCharts {
	this := VisitorsCharts{}
	return &this
}

// GetVisitors returns the Visitors field value if set, zero value otherwise.
func (o *VisitorsCharts) GetVisitors() VisitorsChartsVisitors {
	if o == nil || IsNil(o.Visitors) {
		var ret VisitorsChartsVisitors
		return ret
	}
	return *o.Visitors
}

// GetVisitorsOk returns a tuple with the Visitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitorsCharts) GetVisitorsOk() (*VisitorsChartsVisitors, bool) {
	if o == nil || IsNil(o.Visitors) {
		return nil, false
	}
	return o.Visitors, true
}

// HasVisitors returns a boolean if a field has been set.
func (o *VisitorsCharts) HasVisitors() bool {
	if o != nil && !IsNil(o.Visitors) {
		return true
	}

	return false
}

// SetVisitors gets a reference to the given VisitorsChartsVisitors and assigns it to the Visitors field.
func (o *VisitorsCharts) SetVisitors(v VisitorsChartsVisitors) {
	o.Visitors = &v
}

func (o VisitorsCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisitorsCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Visitors) {
		toSerialize["visitors"] = o.Visitors
	}
	return toSerialize, nil
}

type NullableVisitorsCharts struct {
	value *VisitorsCharts
	isSet bool
}

func (v NullableVisitorsCharts) Get() *VisitorsCharts {
	return v.value
}

func (v *NullableVisitorsCharts) Set(val *VisitorsCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableVisitorsCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableVisitorsCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisitorsCharts(val *VisitorsCharts) *NullableVisitorsCharts {
	return &NullableVisitorsCharts{value: val, isSet: true}
}

func (v NullableVisitorsCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisitorsCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


