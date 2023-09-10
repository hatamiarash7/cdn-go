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

// checks if the WafRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafRule{}

// WafRule struct for WafRule
type WafRule struct {
	Id *string `json:"id,omitempty"`
	// - `?` matches any single character. - `*` matches any (possibly empty) sequence of characters. - `**` matches the current directory and arbitrary subdirectories. This sequence must form a single path component, so both `**a` and `b**` are invalid and will result in an error. A sequence of more than two consecutive `*` characters is also invalid. - `[...]` matches any character inside the brackets. Character sequences can also specify ranges of characters, as ordered by Unicode, so e.g. `[0-9]` specifies any character between 0 and 9 inclusive. An unclosed bracket is invalid. - `[!...]` is the negation of `[...]`, i.e. it matches any characters not in the brackets. - The metacharacters `?`, `*`, `[`, `] `can be matched by using brackets (e.g. `[?]`). When a `]` occurs immediately following `[` or `[!` then it is interpreted as being part of, rather then ending, the character set, so `]` and NOT `]` can be matched by `[]]` and `[!]]` respectively. The - character can be specified inside a character sequence pattern by placing it at the start or the end, e.g. `[abc-]`. 
	UrlPattern *string `json:"url_pattern,omitempty"`
	Sources []string `json:"sources,omitempty"`
	Action *string `json:"action,omitempty"`
	Description *string `json:"description,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
}

// NewWafRule instantiates a new WafRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRule() *WafRule {
	this := WafRule{}
	return &this
}

// NewWafRuleWithDefaults instantiates a new WafRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleWithDefaults() *WafRule {
	this := WafRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WafRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WafRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WafRule) SetId(v string) {
	o.Id = &v
}

// GetUrlPattern returns the UrlPattern field value if set, zero value otherwise.
func (o *WafRule) GetUrlPattern() string {
	if o == nil || IsNil(o.UrlPattern) {
		var ret string
		return ret
	}
	return *o.UrlPattern
}

// GetUrlPatternOk returns a tuple with the UrlPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetUrlPatternOk() (*string, bool) {
	if o == nil || IsNil(o.UrlPattern) {
		return nil, false
	}
	return o.UrlPattern, true
}

// HasUrlPattern returns a boolean if a field has been set.
func (o *WafRule) HasUrlPattern() bool {
	if o != nil && !IsNil(o.UrlPattern) {
		return true
	}

	return false
}

// SetUrlPattern gets a reference to the given string and assigns it to the UrlPattern field.
func (o *WafRule) SetUrlPattern(v string) {
	o.UrlPattern = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *WafRule) GetSources() []string {
	if o == nil || IsNil(o.Sources) {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *WafRule) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *WafRule) SetSources(v []string) {
	o.Sources = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WafRule) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WafRule) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *WafRule) SetAction(v string) {
	o.Action = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WafRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WafRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WafRule) SetDescription(v string) {
	o.Description = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *WafRule) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *WafRule) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *WafRule) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o WafRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.UrlPattern) {
		toSerialize["url_pattern"] = o.UrlPattern
	}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	return toSerialize, nil
}

type NullableWafRule struct {
	value *WafRule
	isSet bool
}

func (v NullableWafRule) Get() *WafRule {
	return v.value
}

func (v *NullableWafRule) Set(val *WafRule) {
	v.value = val
	v.isSet = true
}

func (v NullableWafRule) IsSet() bool {
	return v.isSet
}

func (v *NullableWafRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafRule(val *WafRule) *NullableWafRule {
	return &NullableWafRule{value: val, isSet: true}
}

func (v NullableWafRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


