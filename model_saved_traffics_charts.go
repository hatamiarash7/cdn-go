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

// checks if the SavedTrafficsCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedTrafficsCharts{}

// SavedTrafficsCharts struct for SavedTrafficsCharts
type SavedTrafficsCharts struct {
	Request []SavedTrafficsChartsRequestInner `json:"request,omitempty"`
	Traffic []SavedTrafficsChartsTrafficInner `json:"traffic,omitempty"`
}

// NewSavedTrafficsCharts instantiates a new SavedTrafficsCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedTrafficsCharts() *SavedTrafficsCharts {
	this := SavedTrafficsCharts{}
	return &this
}

// NewSavedTrafficsChartsWithDefaults instantiates a new SavedTrafficsCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedTrafficsChartsWithDefaults() *SavedTrafficsCharts {
	this := SavedTrafficsCharts{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SavedTrafficsCharts) GetRequest() []SavedTrafficsChartsRequestInner {
	if o == nil || IsNil(o.Request) {
		var ret []SavedTrafficsChartsRequestInner
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsCharts) GetRequestOk() ([]SavedTrafficsChartsRequestInner, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SavedTrafficsCharts) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given []SavedTrafficsChartsRequestInner and assigns it to the Request field.
func (o *SavedTrafficsCharts) SetRequest(v []SavedTrafficsChartsRequestInner) {
	o.Request = v
}

// GetTraffic returns the Traffic field value if set, zero value otherwise.
func (o *SavedTrafficsCharts) GetTraffic() []SavedTrafficsChartsTrafficInner {
	if o == nil || IsNil(o.Traffic) {
		var ret []SavedTrafficsChartsTrafficInner
		return ret
	}
	return o.Traffic
}

// GetTrafficOk returns a tuple with the Traffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsCharts) GetTrafficOk() ([]SavedTrafficsChartsTrafficInner, bool) {
	if o == nil || IsNil(o.Traffic) {
		return nil, false
	}
	return o.Traffic, true
}

// HasTraffic returns a boolean if a field has been set.
func (o *SavedTrafficsCharts) HasTraffic() bool {
	if o != nil && !IsNil(o.Traffic) {
		return true
	}

	return false
}

// SetTraffic gets a reference to the given []SavedTrafficsChartsTrafficInner and assigns it to the Traffic field.
func (o *SavedTrafficsCharts) SetTraffic(v []SavedTrafficsChartsTrafficInner) {
	o.Traffic = v
}

func (o SavedTrafficsCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedTrafficsCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Traffic) {
		toSerialize["traffic"] = o.Traffic
	}
	return toSerialize, nil
}

type NullableSavedTrafficsCharts struct {
	value *SavedTrafficsCharts
	isSet bool
}

func (v NullableSavedTrafficsCharts) Get() *SavedTrafficsCharts {
	return v.value
}

func (v *NullableSavedTrafficsCharts) Set(val *SavedTrafficsCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedTrafficsCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedTrafficsCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedTrafficsCharts(val *SavedTrafficsCharts) *NullableSavedTrafficsCharts {
	return &NullableSavedTrafficsCharts{value: val, isSet: true}
}

func (v NullableSavedTrafficsCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedTrafficsCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


