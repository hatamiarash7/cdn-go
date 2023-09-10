/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.108.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"time"
)

// checks if the DnsRequestReportStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRequestReportStatistics{}

// DnsRequestReportStatistics struct for DnsRequestReportStatistics
type DnsRequestReportStatistics struct {
	Total *int32 `json:"total,omitempty"`
	Top *time.Time `json:"top,omitempty"`
}

// NewDnsRequestReportStatistics instantiates a new DnsRequestReportStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRequestReportStatistics() *DnsRequestReportStatistics {
	this := DnsRequestReportStatistics{}
	return &this
}

// NewDnsRequestReportStatisticsWithDefaults instantiates a new DnsRequestReportStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRequestReportStatisticsWithDefaults() *DnsRequestReportStatistics {
	this := DnsRequestReportStatistics{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DnsRequestReportStatistics) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRequestReportStatistics) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DnsRequestReportStatistics) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *DnsRequestReportStatistics) SetTotal(v int32) {
	o.Total = &v
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *DnsRequestReportStatistics) GetTop() time.Time {
	if o == nil || IsNil(o.Top) {
		var ret time.Time
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRequestReportStatistics) GetTopOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Top) {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *DnsRequestReportStatistics) HasTop() bool {
	if o != nil && !IsNil(o.Top) {
		return true
	}

	return false
}

// SetTop gets a reference to the given time.Time and assigns it to the Top field.
func (o *DnsRequestReportStatistics) SetTop(v time.Time) {
	o.Top = &v
}

func (o DnsRequestReportStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRequestReportStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Top) {
		toSerialize["top"] = o.Top
	}
	return toSerialize, nil
}

type NullableDnsRequestReportStatistics struct {
	value *DnsRequestReportStatistics
	isSet bool
}

func (v NullableDnsRequestReportStatistics) Get() *DnsRequestReportStatistics {
	return v.value
}

func (v *NullableDnsRequestReportStatistics) Set(val *DnsRequestReportStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRequestReportStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRequestReportStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRequestReportStatistics(val *DnsRequestReportStatistics) *NullableDnsRequestReportStatistics {
	return &NullableDnsRequestReportStatistics{value: val, isSet: true}
}

func (v NullableDnsRequestReportStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRequestReportStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


