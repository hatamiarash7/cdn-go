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

// checks if the FirewallRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallRule{}

// FirewallRule struct for FirewallRule
type FirewallRule struct {
	ActionDetails NullableFirewallActionDetails `json:"action_details,omitempty"`
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	// Wireshark-like filter expression
	FilterExpr string `json:"filter_expr"`
	Action string `json:"action"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	Note *string `json:"note,omitempty"`
}

// NewFirewallRule instantiates a new FirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRule(name string, filterExpr string, action string) *FirewallRule {
	this := FirewallRule{}
	this.Name = name
	this.FilterExpr = filterExpr
	this.Action = action
	return &this
}

// NewFirewallRuleWithDefaults instantiates a new FirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRuleWithDefaults() *FirewallRule {
	this := FirewallRule{}
	return &this
}

// GetActionDetails returns the ActionDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallRule) GetActionDetails() FirewallActionDetails {
	if o == nil || IsNil(o.ActionDetails.Get()) {
		var ret FirewallActionDetails
		return ret
	}
	return *o.ActionDetails.Get()
}

// GetActionDetailsOk returns a tuple with the ActionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallRule) GetActionDetailsOk() (*FirewallActionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionDetails.Get(), o.ActionDetails.IsSet()
}

// HasActionDetails returns a boolean if a field has been set.
func (o *FirewallRule) HasActionDetails() bool {
	if o != nil && o.ActionDetails.IsSet() {
		return true
	}

	return false
}

// SetActionDetails gets a reference to the given NullableFirewallActionDetails and assigns it to the ActionDetails field.
func (o *FirewallRule) SetActionDetails(v FirewallActionDetails) {
	o.ActionDetails.Set(&v)
}
// SetActionDetailsNil sets the value for ActionDetails to be an explicit nil
func (o *FirewallRule) SetActionDetailsNil() {
	o.ActionDetails.Set(nil)
}

// UnsetActionDetails ensures that no value is present for ActionDetails, not even an explicit nil
func (o *FirewallRule) UnsetActionDetails() {
	o.ActionDetails.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FirewallRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FirewallRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FirewallRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *FirewallRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FirewallRule) SetName(v string) {
	o.Name = v
}

// GetFilterExpr returns the FilterExpr field value
func (o *FirewallRule) GetFilterExpr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterExpr
}

// GetFilterExprOk returns a tuple with the FilterExpr field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetFilterExprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterExpr, true
}

// SetFilterExpr sets field value
func (o *FirewallRule) SetFilterExpr(v string) {
	o.FilterExpr = v
}

// GetAction returns the Action field value
func (o *FirewallRule) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *FirewallRule) SetAction(v string) {
	o.Action = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *FirewallRule) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *FirewallRule) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *FirewallRule) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *FirewallRule) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRule) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *FirewallRule) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *FirewallRule) SetNote(v string) {
	o.Note = &v
}

func (o FirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionDetails.IsSet() {
		toSerialize["action_details"] = o.ActionDetails.Get()
	}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	toSerialize["filter_expr"] = o.FilterExpr
	toSerialize["action"] = o.Action
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

type NullableFirewallRule struct {
	value *FirewallRule
	isSet bool
}

func (v NullableFirewallRule) Get() *FirewallRule {
	return v.value
}

func (v *NullableFirewallRule) Set(val *FirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRule(val *FirewallRule) *NullableFirewallRule {
	return &NullableFirewallRule{value: val, isSet: true}
}

func (v NullableFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


