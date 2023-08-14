/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.104.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
	"time"
)

// checks if the BaseHealthCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseHealthCheck{}

// BaseHealthCheck struct for BaseHealthCheck
type BaseHealthCheck struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// can be IP/Host when type is `upstream`, otherwise it must be a valid record ID
	Origin *string `json:"origin,omitempty"`
	OriginType *string `json:"origin_type,omitempty"`
	Upstreams []string `json:"upstreams,omitempty"`
	// In milliseconds
	Interval *int32 `json:"interval,omitempty"`
	Threshold *int32 `json:"threshold,omitempty"`
	Type *string `json:"type,omitempty"`
	// The health-check is off or on
	Status *bool `json:"status,omitempty"`
	// Number of immediate retries in case of a timeout
	Retries *int32 `json:"retries,omitempty"`
	Zones []HealthCheckZone `json:"zones,omitempty"`
	MonitoringUpdatedAt NullableTime `json:"monitoring_updated_at,omitempty"`
}

// NewBaseHealthCheck instantiates a new BaseHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseHealthCheck() *BaseHealthCheck {
	this := BaseHealthCheck{}
	var status bool = true
	this.Status = &status
	return &this
}

// NewBaseHealthCheckWithDefaults instantiates a new BaseHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseHealthCheckWithDefaults() *BaseHealthCheck {
	this := BaseHealthCheck{}
	var status bool = true
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseHealthCheck) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseHealthCheck) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseHealthCheck) SetDescription(v string) {
	o.Description = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *BaseHealthCheck) SetOrigin(v string) {
	o.Origin = &v
}

// GetOriginType returns the OriginType field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetOriginType() string {
	if o == nil || IsNil(o.OriginType) {
		var ret string
		return ret
	}
	return *o.OriginType
}

// GetOriginTypeOk returns a tuple with the OriginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetOriginTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OriginType) {
		return nil, false
	}
	return o.OriginType, true
}

// HasOriginType returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasOriginType() bool {
	if o != nil && !IsNil(o.OriginType) {
		return true
	}

	return false
}

// SetOriginType gets a reference to the given string and assigns it to the OriginType field.
func (o *BaseHealthCheck) SetOriginType(v string) {
	o.OriginType = &v
}

// GetUpstreams returns the Upstreams field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetUpstreams() []string {
	if o == nil || IsNil(o.Upstreams) {
		var ret []string
		return ret
	}
	return o.Upstreams
}

// GetUpstreamsOk returns a tuple with the Upstreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetUpstreamsOk() ([]string, bool) {
	if o == nil || IsNil(o.Upstreams) {
		return nil, false
	}
	return o.Upstreams, true
}

// HasUpstreams returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasUpstreams() bool {
	if o != nil && !IsNil(o.Upstreams) {
		return true
	}

	return false
}

// SetUpstreams gets a reference to the given []string and assigns it to the Upstreams field.
func (o *BaseHealthCheck) SetUpstreams(v []string) {
	o.Upstreams = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *BaseHealthCheck) SetInterval(v int32) {
	o.Interval = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *BaseHealthCheck) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BaseHealthCheck) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *BaseHealthCheck) SetStatus(v bool) {
	o.Status = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetRetries() int32 {
	if o == nil || IsNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasRetries() bool {
	if o != nil && !IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *BaseHealthCheck) SetRetries(v int32) {
	o.Retries = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *BaseHealthCheck) GetZones() []HealthCheckZone {
	if o == nil || IsNil(o.Zones) {
		var ret []HealthCheckZone
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseHealthCheck) GetZonesOk() ([]HealthCheckZone, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []HealthCheckZone and assigns it to the Zones field.
func (o *BaseHealthCheck) SetZones(v []HealthCheckZone) {
	o.Zones = v
}

// GetMonitoringUpdatedAt returns the MonitoringUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseHealthCheck) GetMonitoringUpdatedAt() time.Time {
	if o == nil || IsNil(o.MonitoringUpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MonitoringUpdatedAt.Get()
}

// GetMonitoringUpdatedAtOk returns a tuple with the MonitoringUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseHealthCheck) GetMonitoringUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitoringUpdatedAt.Get(), o.MonitoringUpdatedAt.IsSet()
}

// HasMonitoringUpdatedAt returns a boolean if a field has been set.
func (o *BaseHealthCheck) HasMonitoringUpdatedAt() bool {
	if o != nil && o.MonitoringUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetMonitoringUpdatedAt gets a reference to the given NullableTime and assigns it to the MonitoringUpdatedAt field.
func (o *BaseHealthCheck) SetMonitoringUpdatedAt(v time.Time) {
	o.MonitoringUpdatedAt.Set(&v)
}
// SetMonitoringUpdatedAtNil sets the value for MonitoringUpdatedAt to be an explicit nil
func (o *BaseHealthCheck) SetMonitoringUpdatedAtNil() {
	o.MonitoringUpdatedAt.Set(nil)
}

// UnsetMonitoringUpdatedAt ensures that no value is present for MonitoringUpdatedAt, not even an explicit nil
func (o *BaseHealthCheck) UnsetMonitoringUpdatedAt() {
	o.MonitoringUpdatedAt.Unset()
}

func (o BaseHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseHealthCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.OriginType) {
		toSerialize["origin_type"] = o.OriginType
	}
	if !IsNil(o.Upstreams) {
		toSerialize["upstreams"] = o.Upstreams
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	if o.MonitoringUpdatedAt.IsSet() {
		toSerialize["monitoring_updated_at"] = o.MonitoringUpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullableBaseHealthCheck struct {
	value *BaseHealthCheck
	isSet bool
}

func (v NullableBaseHealthCheck) Get() *BaseHealthCheck {
	return v.value
}

func (v *NullableBaseHealthCheck) Set(val *BaseHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseHealthCheck(val *BaseHealthCheck) *NullableBaseHealthCheck {
	return &NullableBaseHealthCheck{value: val, isSet: true}
}

func (v NullableBaseHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


