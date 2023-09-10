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

// checks if the SavedTrafficsStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedTrafficsStatistics{}

// SavedTrafficsStatistics struct for SavedTrafficsStatistics
type SavedTrafficsStatistics struct {
	Traffic *SavedTrafficsStatisticsTraffic `json:"traffic,omitempty"`
	Request *SavedTrafficsStatisticsTraffic `json:"request,omitempty"`
}

// NewSavedTrafficsStatistics instantiates a new SavedTrafficsStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedTrafficsStatistics() *SavedTrafficsStatistics {
	this := SavedTrafficsStatistics{}
	return &this
}

// NewSavedTrafficsStatisticsWithDefaults instantiates a new SavedTrafficsStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedTrafficsStatisticsWithDefaults() *SavedTrafficsStatistics {
	this := SavedTrafficsStatistics{}
	return &this
}

// GetTraffic returns the Traffic field value if set, zero value otherwise.
func (o *SavedTrafficsStatistics) GetTraffic() SavedTrafficsStatisticsTraffic {
	if o == nil || IsNil(o.Traffic) {
		var ret SavedTrafficsStatisticsTraffic
		return ret
	}
	return *o.Traffic
}

// GetTrafficOk returns a tuple with the Traffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsStatistics) GetTrafficOk() (*SavedTrafficsStatisticsTraffic, bool) {
	if o == nil || IsNil(o.Traffic) {
		return nil, false
	}
	return o.Traffic, true
}

// HasTraffic returns a boolean if a field has been set.
func (o *SavedTrafficsStatistics) HasTraffic() bool {
	if o != nil && !IsNil(o.Traffic) {
		return true
	}

	return false
}

// SetTraffic gets a reference to the given SavedTrafficsStatisticsTraffic and assigns it to the Traffic field.
func (o *SavedTrafficsStatistics) SetTraffic(v SavedTrafficsStatisticsTraffic) {
	o.Traffic = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SavedTrafficsStatistics) GetRequest() SavedTrafficsStatisticsTraffic {
	if o == nil || IsNil(o.Request) {
		var ret SavedTrafficsStatisticsTraffic
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrafficsStatistics) GetRequestOk() (*SavedTrafficsStatisticsTraffic, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SavedTrafficsStatistics) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given SavedTrafficsStatisticsTraffic and assigns it to the Request field.
func (o *SavedTrafficsStatistics) SetRequest(v SavedTrafficsStatisticsTraffic) {
	o.Request = &v
}

func (o SavedTrafficsStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedTrafficsStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Traffic) {
		toSerialize["traffic"] = o.Traffic
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return toSerialize, nil
}

type NullableSavedTrafficsStatistics struct {
	value *SavedTrafficsStatistics
	isSet bool
}

func (v NullableSavedTrafficsStatistics) Get() *SavedTrafficsStatistics {
	return v.value
}

func (v *NullableSavedTrafficsStatistics) Set(val *SavedTrafficsStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedTrafficsStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedTrafficsStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedTrafficsStatistics(val *SavedTrafficsStatistics) *NullableSavedTrafficsStatistics {
	return &NullableSavedTrafficsStatistics{value: val, isSet: true}
}

func (v NullableSavedTrafficsStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedTrafficsStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


