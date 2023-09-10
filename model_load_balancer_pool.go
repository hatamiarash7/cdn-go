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

// checks if the LoadBalancerPool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPool{}

// LoadBalancerPool struct for LoadBalancerPool
type LoadBalancerPool struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *bool `json:"status,omitempty"`
	// Zero means the default pool
	Priority *int32 `json:"priority,omitempty"`
	Method *string `json:"method,omitempty"`
	Keepalive *string `json:"keepalive,omitempty"`
	NextUpstreamTcp *string `json:"next_upstream_tcp,omitempty"`
	Regions []string `json:"regions,omitempty"`
	Origins []LoadBalancerOrigin `json:"origins,omitempty"`
	MonitoringStatus NullableMonitoringStatus `json:"monitoring_status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewLoadBalancerPool instantiates a new LoadBalancerPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPool() *LoadBalancerPool {
	this := LoadBalancerPool{}
	var keepalive string = "off"
	this.Keepalive = &keepalive
	var nextUpstreamTcp string = "off"
	this.NextUpstreamTcp = &nextUpstreamTcp
	return &this
}

// NewLoadBalancerPoolWithDefaults instantiates a new LoadBalancerPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPoolWithDefaults() *LoadBalancerPool {
	this := LoadBalancerPool{}
	var keepalive string = "off"
	this.Keepalive = &keepalive
	var nextUpstreamTcp string = "off"
	this.NextUpstreamTcp = &nextUpstreamTcp
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoadBalancerPool) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancerPool) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadBalancerPool) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *LoadBalancerPool) SetStatus(v bool) {
	o.Status = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *LoadBalancerPool) SetPriority(v int32) {
	o.Priority = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoadBalancerPool) SetMethod(v string) {
	o.Method = &v
}

// GetKeepalive returns the Keepalive field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetKeepalive() string {
	if o == nil || IsNil(o.Keepalive) {
		var ret string
		return ret
	}
	return *o.Keepalive
}

// GetKeepaliveOk returns a tuple with the Keepalive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetKeepaliveOk() (*string, bool) {
	if o == nil || IsNil(o.Keepalive) {
		return nil, false
	}
	return o.Keepalive, true
}

// HasKeepalive returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasKeepalive() bool {
	if o != nil && !IsNil(o.Keepalive) {
		return true
	}

	return false
}

// SetKeepalive gets a reference to the given string and assigns it to the Keepalive field.
func (o *LoadBalancerPool) SetKeepalive(v string) {
	o.Keepalive = &v
}

// GetNextUpstreamTcp returns the NextUpstreamTcp field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetNextUpstreamTcp() string {
	if o == nil || IsNil(o.NextUpstreamTcp) {
		var ret string
		return ret
	}
	return *o.NextUpstreamTcp
}

// GetNextUpstreamTcpOk returns a tuple with the NextUpstreamTcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetNextUpstreamTcpOk() (*string, bool) {
	if o == nil || IsNil(o.NextUpstreamTcp) {
		return nil, false
	}
	return o.NextUpstreamTcp, true
}

// HasNextUpstreamTcp returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasNextUpstreamTcp() bool {
	if o != nil && !IsNil(o.NextUpstreamTcp) {
		return true
	}

	return false
}

// SetNextUpstreamTcp gets a reference to the given string and assigns it to the NextUpstreamTcp field.
func (o *LoadBalancerPool) SetNextUpstreamTcp(v string) {
	o.NextUpstreamTcp = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *LoadBalancerPool) SetRegions(v []string) {
	o.Regions = v
}

// GetOrigins returns the Origins field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetOrigins() []LoadBalancerOrigin {
	if o == nil || IsNil(o.Origins) {
		var ret []LoadBalancerOrigin
		return ret
	}
	return o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetOriginsOk() ([]LoadBalancerOrigin, bool) {
	if o == nil || IsNil(o.Origins) {
		return nil, false
	}
	return o.Origins, true
}

// HasOrigins returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasOrigins() bool {
	if o != nil && !IsNil(o.Origins) {
		return true
	}

	return false
}

// SetOrigins gets a reference to the given []LoadBalancerOrigin and assigns it to the Origins field.
func (o *LoadBalancerPool) SetOrigins(v []LoadBalancerOrigin) {
	o.Origins = v
}

// GetMonitoringStatus returns the MonitoringStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadBalancerPool) GetMonitoringStatus() MonitoringStatus {
	if o == nil || IsNil(o.MonitoringStatus.Get()) {
		var ret MonitoringStatus
		return ret
	}
	return *o.MonitoringStatus.Get()
}

// GetMonitoringStatusOk returns a tuple with the MonitoringStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerPool) GetMonitoringStatusOk() (*MonitoringStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonitoringStatus.Get(), o.MonitoringStatus.IsSet()
}

// HasMonitoringStatus returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasMonitoringStatus() bool {
	if o != nil && o.MonitoringStatus.IsSet() {
		return true
	}

	return false
}

// SetMonitoringStatus gets a reference to the given NullableMonitoringStatus and assigns it to the MonitoringStatus field.
func (o *LoadBalancerPool) SetMonitoringStatus(v MonitoringStatus) {
	o.MonitoringStatus.Set(&v)
}
// SetMonitoringStatusNil sets the value for MonitoringStatus to be an explicit nil
func (o *LoadBalancerPool) SetMonitoringStatusNil() {
	o.MonitoringStatus.Set(nil)
}

// UnsetMonitoringStatus ensures that no value is present for MonitoringStatus, not even an explicit nil
func (o *LoadBalancerPool) UnsetMonitoringStatus() {
	o.MonitoringStatus.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LoadBalancerPool) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LoadBalancerPool) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPool) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoadBalancerPool) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *LoadBalancerPool) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o LoadBalancerPool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Keepalive) {
		toSerialize["keepalive"] = o.Keepalive
	}
	if !IsNil(o.NextUpstreamTcp) {
		toSerialize["next_upstream_tcp"] = o.NextUpstreamTcp
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	// skip: origins is readOnly
	if o.MonitoringStatus.IsSet() {
		toSerialize["monitoring_status"] = o.MonitoringStatus.Get()
	}
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	return toSerialize, nil
}

type NullableLoadBalancerPool struct {
	value *LoadBalancerPool
	isSet bool
}

func (v NullableLoadBalancerPool) Get() *LoadBalancerPool {
	return v.value
}

func (v *NullableLoadBalancerPool) Set(val *LoadBalancerPool) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPool) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPool(val *LoadBalancerPool) *NullableLoadBalancerPool {
	return &NullableLoadBalancerPool{value: val, isSet: true}
}

func (v NullableLoadBalancerPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


