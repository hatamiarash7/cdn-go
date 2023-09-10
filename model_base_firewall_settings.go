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

// checks if the BaseFirewallSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseFirewallSettings{}

// BaseFirewallSettings struct for BaseFirewallSettings
type BaseFirewallSettings struct {
	IsEnabled *bool `json:"is_enabled,omitempty"`
	DefaultAction *string `json:"default_action,omitempty"`
	// True to verify that SNI and hostname are equal
	VerifySni *bool `json:"verify_sni,omitempty"`
}

// NewBaseFirewallSettings instantiates a new BaseFirewallSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseFirewallSettings() *BaseFirewallSettings {
	this := BaseFirewallSettings{}
	var verifySni bool = false
	this.VerifySni = &verifySni
	return &this
}

// NewBaseFirewallSettingsWithDefaults instantiates a new BaseFirewallSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseFirewallSettingsWithDefaults() *BaseFirewallSettings {
	this := BaseFirewallSettings{}
	var verifySni bool = false
	this.VerifySni = &verifySni
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *BaseFirewallSettings) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallSettings) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *BaseFirewallSettings) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *BaseFirewallSettings) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDefaultAction returns the DefaultAction field value if set, zero value otherwise.
func (o *BaseFirewallSettings) GetDefaultAction() string {
	if o == nil || IsNil(o.DefaultAction) {
		var ret string
		return ret
	}
	return *o.DefaultAction
}

// GetDefaultActionOk returns a tuple with the DefaultAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallSettings) GetDefaultActionOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAction) {
		return nil, false
	}
	return o.DefaultAction, true
}

// HasDefaultAction returns a boolean if a field has been set.
func (o *BaseFirewallSettings) HasDefaultAction() bool {
	if o != nil && !IsNil(o.DefaultAction) {
		return true
	}

	return false
}

// SetDefaultAction gets a reference to the given string and assigns it to the DefaultAction field.
func (o *BaseFirewallSettings) SetDefaultAction(v string) {
	o.DefaultAction = &v
}

// GetVerifySni returns the VerifySni field value if set, zero value otherwise.
func (o *BaseFirewallSettings) GetVerifySni() bool {
	if o == nil || IsNil(o.VerifySni) {
		var ret bool
		return ret
	}
	return *o.VerifySni
}

// GetVerifySniOk returns a tuple with the VerifySni field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallSettings) GetVerifySniOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySni) {
		return nil, false
	}
	return o.VerifySni, true
}

// HasVerifySni returns a boolean if a field has been set.
func (o *BaseFirewallSettings) HasVerifySni() bool {
	if o != nil && !IsNil(o.VerifySni) {
		return true
	}

	return false
}

// SetVerifySni gets a reference to the given bool and assigns it to the VerifySni field.
func (o *BaseFirewallSettings) SetVerifySni(v bool) {
	o.VerifySni = &v
}

func (o BaseFirewallSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseFirewallSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: is_enabled is readOnly
	if !IsNil(o.DefaultAction) {
		toSerialize["default_action"] = o.DefaultAction
	}
	if !IsNil(o.VerifySni) {
		toSerialize["verify_sni"] = o.VerifySni
	}
	return toSerialize, nil
}

type NullableBaseFirewallSettings struct {
	value *BaseFirewallSettings
	isSet bool
}

func (v NullableBaseFirewallSettings) Get() *BaseFirewallSettings {
	return v.value
}

func (v *NullableBaseFirewallSettings) Set(val *BaseFirewallSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseFirewallSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseFirewallSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseFirewallSettings(val *BaseFirewallSettings) *NullableBaseFirewallSettings {
	return &NullableBaseFirewallSettings{value: val, isSet: true}
}

func (v NullableBaseFirewallSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseFirewallSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


