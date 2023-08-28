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

// checks if the FirewallSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallSettings{}

// FirewallSettings struct for FirewallSettings
type FirewallSettings struct {
	DefaultActionDetails NullableFirewallActionDetails `json:"default_action_details,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	DefaultAction *string `json:"default_action,omitempty"`
	// True to verify that SNI and hostname are equal
	VerifySni *bool `json:"verify_sni,omitempty"`
}

// NewFirewallSettings instantiates a new FirewallSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallSettings() *FirewallSettings {
	this := FirewallSettings{}
	var verifySni bool = false
	this.VerifySni = &verifySni
	return &this
}

// NewFirewallSettingsWithDefaults instantiates a new FirewallSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallSettingsWithDefaults() *FirewallSettings {
	this := FirewallSettings{}
	var verifySni bool = false
	this.VerifySni = &verifySni
	return &this
}

// GetDefaultActionDetails returns the DefaultActionDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallSettings) GetDefaultActionDetails() FirewallActionDetails {
	if o == nil || IsNil(o.DefaultActionDetails.Get()) {
		var ret FirewallActionDetails
		return ret
	}
	return *o.DefaultActionDetails.Get()
}

// GetDefaultActionDetailsOk returns a tuple with the DefaultActionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallSettings) GetDefaultActionDetailsOk() (*FirewallActionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultActionDetails.Get(), o.DefaultActionDetails.IsSet()
}

// HasDefaultActionDetails returns a boolean if a field has been set.
func (o *FirewallSettings) HasDefaultActionDetails() bool {
	if o != nil && o.DefaultActionDetails.IsSet() {
		return true
	}

	return false
}

// SetDefaultActionDetails gets a reference to the given NullableFirewallActionDetails and assigns it to the DefaultActionDetails field.
func (o *FirewallSettings) SetDefaultActionDetails(v FirewallActionDetails) {
	o.DefaultActionDetails.Set(&v)
}
// SetDefaultActionDetailsNil sets the value for DefaultActionDetails to be an explicit nil
func (o *FirewallSettings) SetDefaultActionDetailsNil() {
	o.DefaultActionDetails.Set(nil)
}

// UnsetDefaultActionDetails ensures that no value is present for DefaultActionDetails, not even an explicit nil
func (o *FirewallSettings) UnsetDefaultActionDetails() {
	o.DefaultActionDetails.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *FirewallSettings) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSettings) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *FirewallSettings) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *FirewallSettings) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDefaultAction returns the DefaultAction field value if set, zero value otherwise.
func (o *FirewallSettings) GetDefaultAction() string {
	if o == nil || IsNil(o.DefaultAction) {
		var ret string
		return ret
	}
	return *o.DefaultAction
}

// GetDefaultActionOk returns a tuple with the DefaultAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSettings) GetDefaultActionOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAction) {
		return nil, false
	}
	return o.DefaultAction, true
}

// HasDefaultAction returns a boolean if a field has been set.
func (o *FirewallSettings) HasDefaultAction() bool {
	if o != nil && !IsNil(o.DefaultAction) {
		return true
	}

	return false
}

// SetDefaultAction gets a reference to the given string and assigns it to the DefaultAction field.
func (o *FirewallSettings) SetDefaultAction(v string) {
	o.DefaultAction = &v
}

// GetVerifySni returns the VerifySni field value if set, zero value otherwise.
func (o *FirewallSettings) GetVerifySni() bool {
	if o == nil || IsNil(o.VerifySni) {
		var ret bool
		return ret
	}
	return *o.VerifySni
}

// GetVerifySniOk returns a tuple with the VerifySni field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSettings) GetVerifySniOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySni) {
		return nil, false
	}
	return o.VerifySni, true
}

// HasVerifySni returns a boolean if a field has been set.
func (o *FirewallSettings) HasVerifySni() bool {
	if o != nil && !IsNil(o.VerifySni) {
		return true
	}

	return false
}

// SetVerifySni gets a reference to the given bool and assigns it to the VerifySni field.
func (o *FirewallSettings) SetVerifySni(v bool) {
	o.VerifySni = &v
}

func (o FirewallSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultActionDetails.IsSet() {
		toSerialize["default_action_details"] = o.DefaultActionDetails.Get()
	}
	// skip: is_enabled is readOnly
	if !IsNil(o.DefaultAction) {
		toSerialize["default_action"] = o.DefaultAction
	}
	if !IsNil(o.VerifySni) {
		toSerialize["verify_sni"] = o.VerifySni
	}
	return toSerialize, nil
}

type NullableFirewallSettings struct {
	value *FirewallSettings
	isSet bool
}

func (v NullableFirewallSettings) Get() *FirewallSettings {
	return v.value
}

func (v *NullableFirewallSettings) Set(val *FirewallSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallSettings(val *FirewallSettings) *NullableFirewallSettings {
	return &NullableFirewallSettings{value: val, isSet: true}
}

func (v NullableFirewallSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


