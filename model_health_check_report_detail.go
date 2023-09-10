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

// checks if the HealthCheckReportDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckReportDetail{}

// HealthCheckReportDetail struct for HealthCheckReportDetail
type HealthCheckReportDetail struct {
	Date *string `json:"date,omitempty"`
	Zone *string `json:"zone,omitempty"`
	Upstream *string `json:"upstream,omitempty"`
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewHealthCheckReportDetail instantiates a new HealthCheckReportDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckReportDetail() *HealthCheckReportDetail {
	this := HealthCheckReportDetail{}
	return &this
}

// NewHealthCheckReportDetailWithDefaults instantiates a new HealthCheckReportDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckReportDetailWithDefaults() *HealthCheckReportDetail {
	this := HealthCheckReportDetail{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *HealthCheckReportDetail) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportDetail) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *HealthCheckReportDetail) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *HealthCheckReportDetail) SetDate(v string) {
	o.Date = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *HealthCheckReportDetail) GetZone() string {
	if o == nil || IsNil(o.Zone) {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportDetail) GetZoneOk() (*string, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *HealthCheckReportDetail) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *HealthCheckReportDetail) SetZone(v string) {
	o.Zone = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *HealthCheckReportDetail) GetUpstream() string {
	if o == nil || IsNil(o.Upstream) {
		var ret string
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportDetail) GetUpstreamOk() (*string, bool) {
	if o == nil || IsNil(o.Upstream) {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *HealthCheckReportDetail) HasUpstream() bool {
	if o != nil && !IsNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given string and assigns it to the Upstream field.
func (o *HealthCheckReportDetail) SetUpstream(v string) {
	o.Upstream = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthCheckReportDetail) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportDetail) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthCheckReportDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *HealthCheckReportDetail) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HealthCheckReportDetail) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportDetail) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HealthCheckReportDetail) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HealthCheckReportDetail) SetMessage(v string) {
	o.Message = &v
}

func (o HealthCheckReportDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckReportDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableHealthCheckReportDetail struct {
	value *HealthCheckReportDetail
	isSet bool
}

func (v NullableHealthCheckReportDetail) Get() *HealthCheckReportDetail {
	return v.value
}

func (v *NullableHealthCheckReportDetail) Set(val *HealthCheckReportDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckReportDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckReportDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckReportDetail(val *HealthCheckReportDetail) *NullableHealthCheckReportDetail {
	return &NullableHealthCheckReportDetail{value: val, isSet: true}
}

func (v NullableHealthCheckReportDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckReportDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


