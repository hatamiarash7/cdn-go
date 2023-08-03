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

// checks if the BaseFirewallRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseFirewallRule{}

// BaseFirewallRule struct for BaseFirewallRule
type BaseFirewallRule struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Wireshark-like filter expression
	FilterExpr *string `json:"filter_expr,omitempty"`
	Action *string `json:"action,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	Note *string `json:"note,omitempty"`
}

// NewBaseFirewallRule instantiates a new BaseFirewallRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseFirewallRule() *BaseFirewallRule {
	this := BaseFirewallRule{}
	return &this
}

// NewBaseFirewallRuleWithDefaults instantiates a new BaseFirewallRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseFirewallRuleWithDefaults() *BaseFirewallRule {
	this := BaseFirewallRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseFirewallRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseFirewallRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseFirewallRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseFirewallRule) SetName(v string) {
	o.Name = &v
}

// GetFilterExpr returns the FilterExpr field value if set, zero value otherwise.
func (o *BaseFirewallRule) GetFilterExpr() string {
	if o == nil || IsNil(o.FilterExpr) {
		var ret string
		return ret
	}
	return *o.FilterExpr
}

// GetFilterExprOk returns a tuple with the FilterExpr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallRule) GetFilterExprOk() (*string, bool) {
	if o == nil || IsNil(o.FilterExpr) {
		return nil, false
	}
	return o.FilterExpr, true
}

// HasFilterExpr returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasFilterExpr() bool {
	if o != nil && !IsNil(o.FilterExpr) {
		return true
	}

	return false
}

// SetFilterExpr gets a reference to the given string and assigns it to the FilterExpr field.
func (o *BaseFirewallRule) SetFilterExpr(v string) {
	o.FilterExpr = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BaseFirewallRule) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallRule) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BaseFirewallRule) SetAction(v string) {
	o.Action = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *BaseFirewallRule) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallRule) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *BaseFirewallRule) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *BaseFirewallRule) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseFirewallRule) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *BaseFirewallRule) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *BaseFirewallRule) SetNote(v string) {
	o.Note = &v
}

func (o BaseFirewallRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseFirewallRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FilterExpr) {
		toSerialize["filter_expr"] = o.FilterExpr
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

type NullableBaseFirewallRule struct {
	value *BaseFirewallRule
	isSet bool
}

func (v NullableBaseFirewallRule) Get() *BaseFirewallRule {
	return v.value
}

func (v *NullableBaseFirewallRule) Set(val *BaseFirewallRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseFirewallRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseFirewallRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseFirewallRule(val *BaseFirewallRule) *NullableBaseFirewallRule {
	return &NullableBaseFirewallRule{value: val, isSet: true}
}

func (v NullableBaseFirewallRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseFirewallRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


