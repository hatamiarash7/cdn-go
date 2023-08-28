/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the WafPreset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPreset{}

// WafPreset struct for WafPreset
type WafPreset struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Packages []WafPresetPackagesInner `json:"packages,omitempty"`
}

// NewWafPreset instantiates a new WafPreset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPreset() *WafPreset {
	this := WafPreset{}
	return &this
}

// NewWafPresetWithDefaults instantiates a new WafPreset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPresetWithDefaults() *WafPreset {
	this := WafPreset{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WafPreset) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPreset) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WafPreset) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WafPreset) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafPreset) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPreset) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafPreset) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafPreset) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WafPreset) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPreset) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WafPreset) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WafPreset) SetDescription(v string) {
	o.Description = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *WafPreset) GetPackages() []WafPresetPackagesInner {
	if o == nil || IsNil(o.Packages) {
		var ret []WafPresetPackagesInner
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPreset) GetPackagesOk() ([]WafPresetPackagesInner, bool) {
	if o == nil || IsNil(o.Packages) {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *WafPreset) HasPackages() bool {
	if o != nil && !IsNil(o.Packages) {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []WafPresetPackagesInner and assigns it to the Packages field.
func (o *WafPreset) SetPackages(v []WafPresetPackagesInner) {
	o.Packages = v
}

func (o WafPreset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPreset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Packages) {
		toSerialize["packages"] = o.Packages
	}
	return toSerialize, nil
}

type NullableWafPreset struct {
	value *WafPreset
	isSet bool
}

func (v NullableWafPreset) Get() *WafPreset {
	return v.value
}

func (v *NullableWafPreset) Set(val *WafPreset) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPreset) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPreset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPreset(val *WafPreset) *NullableWafPreset {
	return &NullableWafPreset{value: val, isSet: true}
}

func (v NullableWafPreset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPreset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


