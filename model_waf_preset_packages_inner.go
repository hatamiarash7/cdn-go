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

// checks if the WafPresetPackagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPresetPackagesInner{}

// WafPresetPackagesInner struct for WafPresetPackagesInner
type WafPresetPackagesInner struct {
	Name *string `json:"name,omitempty"`
	Provider *WafPresetPackagesInnerProvider `json:"provider,omitempty"`
}

// NewWafPresetPackagesInner instantiates a new WafPresetPackagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPresetPackagesInner() *WafPresetPackagesInner {
	this := WafPresetPackagesInner{}
	return &this
}

// NewWafPresetPackagesInnerWithDefaults instantiates a new WafPresetPackagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPresetPackagesInnerWithDefaults() *WafPresetPackagesInner {
	this := WafPresetPackagesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafPresetPackagesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPresetPackagesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafPresetPackagesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafPresetPackagesInner) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *WafPresetPackagesInner) GetProvider() WafPresetPackagesInnerProvider {
	if o == nil || IsNil(o.Provider) {
		var ret WafPresetPackagesInnerProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPresetPackagesInner) GetProviderOk() (*WafPresetPackagesInnerProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *WafPresetPackagesInner) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given WafPresetPackagesInnerProvider and assigns it to the Provider field.
func (o *WafPresetPackagesInner) SetProvider(v WafPresetPackagesInnerProvider) {
	o.Provider = &v
}

func (o WafPresetPackagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPresetPackagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableWafPresetPackagesInner struct {
	value *WafPresetPackagesInner
	isSet bool
}

func (v NullableWafPresetPackagesInner) Get() *WafPresetPackagesInner {
	return v.value
}

func (v *NullableWafPresetPackagesInner) Set(val *WafPresetPackagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPresetPackagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPresetPackagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPresetPackagesInner(val *WafPresetPackagesInner) *NullableWafPresetPackagesInner {
	return &NullableWafPresetPackagesInner{value: val, isSet: true}
}

func (v NullableWafPresetPackagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPresetPackagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


