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

// checks if the HealthCheckZoneName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckZoneName{}

// HealthCheckZoneName struct for HealthCheckZoneName
type HealthCheckZoneName struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewHealthCheckZoneName instantiates a new HealthCheckZoneName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckZoneName() *HealthCheckZoneName {
	this := HealthCheckZoneName{}
	return &this
}

// NewHealthCheckZoneNameWithDefaults instantiates a new HealthCheckZoneName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckZoneNameWithDefaults() *HealthCheckZoneName {
	this := HealthCheckZoneName{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HealthCheckZoneName) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckZoneName) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HealthCheckZoneName) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HealthCheckZoneName) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HealthCheckZoneName) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckZoneName) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HealthCheckZoneName) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HealthCheckZoneName) SetName(v string) {
	o.Name = &v
}

func (o HealthCheckZoneName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckZoneName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableHealthCheckZoneName struct {
	value *HealthCheckZoneName
	isSet bool
}

func (v NullableHealthCheckZoneName) Get() *HealthCheckZoneName {
	return v.value
}

func (v *NullableHealthCheckZoneName) Set(val *HealthCheckZoneName) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckZoneName) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckZoneName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckZoneName(val *HealthCheckZoneName) *NullableHealthCheckZoneName {
	return &NullableHealthCheckZoneName{value: val, isSet: true}
}

func (v NullableHealthCheckZoneName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckZoneName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


