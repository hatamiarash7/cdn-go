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

// checks if the ReprioritizeRuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReprioritizeRuleRequest{}

// ReprioritizeRuleRequest struct for ReprioritizeRuleRequest
type ReprioritizeRuleRequest struct {
	// ID of the rule you want to move
	RuleId string `json:"rule_id"`
	// ID of the rule you want to be prior to the selected rule
	AfterRuleId *string `json:"after_rule_id,omitempty"`
	// ID of the rule you want to follow the selected rule
	BeforeRuleId *string `json:"before_rule_id,omitempty"`
}

// NewReprioritizeRuleRequest instantiates a new ReprioritizeRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReprioritizeRuleRequest(ruleId string) *ReprioritizeRuleRequest {
	this := ReprioritizeRuleRequest{}
	this.RuleId = ruleId
	return &this
}

// NewReprioritizeRuleRequestWithDefaults instantiates a new ReprioritizeRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReprioritizeRuleRequestWithDefaults() *ReprioritizeRuleRequest {
	this := ReprioritizeRuleRequest{}
	return &this
}

// GetRuleId returns the RuleId field value
func (o *ReprioritizeRuleRequest) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *ReprioritizeRuleRequest) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *ReprioritizeRuleRequest) SetRuleId(v string) {
	o.RuleId = v
}

// GetAfterRuleId returns the AfterRuleId field value if set, zero value otherwise.
func (o *ReprioritizeRuleRequest) GetAfterRuleId() string {
	if o == nil || IsNil(o.AfterRuleId) {
		var ret string
		return ret
	}
	return *o.AfterRuleId
}

// GetAfterRuleIdOk returns a tuple with the AfterRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprioritizeRuleRequest) GetAfterRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfterRuleId) {
		return nil, false
	}
	return o.AfterRuleId, true
}

// HasAfterRuleId returns a boolean if a field has been set.
func (o *ReprioritizeRuleRequest) HasAfterRuleId() bool {
	if o != nil && !IsNil(o.AfterRuleId) {
		return true
	}

	return false
}

// SetAfterRuleId gets a reference to the given string and assigns it to the AfterRuleId field.
func (o *ReprioritizeRuleRequest) SetAfterRuleId(v string) {
	o.AfterRuleId = &v
}

// GetBeforeRuleId returns the BeforeRuleId field value if set, zero value otherwise.
func (o *ReprioritizeRuleRequest) GetBeforeRuleId() string {
	if o == nil || IsNil(o.BeforeRuleId) {
		var ret string
		return ret
	}
	return *o.BeforeRuleId
}

// GetBeforeRuleIdOk returns a tuple with the BeforeRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReprioritizeRuleRequest) GetBeforeRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BeforeRuleId) {
		return nil, false
	}
	return o.BeforeRuleId, true
}

// HasBeforeRuleId returns a boolean if a field has been set.
func (o *ReprioritizeRuleRequest) HasBeforeRuleId() bool {
	if o != nil && !IsNil(o.BeforeRuleId) {
		return true
	}

	return false
}

// SetBeforeRuleId gets a reference to the given string and assigns it to the BeforeRuleId field.
func (o *ReprioritizeRuleRequest) SetBeforeRuleId(v string) {
	o.BeforeRuleId = &v
}

func (o ReprioritizeRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReprioritizeRuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rule_id"] = o.RuleId
	if !IsNil(o.AfterRuleId) {
		toSerialize["after_rule_id"] = o.AfterRuleId
	}
	if !IsNil(o.BeforeRuleId) {
		toSerialize["before_rule_id"] = o.BeforeRuleId
	}
	return toSerialize, nil
}

type NullableReprioritizeRuleRequest struct {
	value *ReprioritizeRuleRequest
	isSet bool
}

func (v NullableReprioritizeRuleRequest) Get() *ReprioritizeRuleRequest {
	return v.value
}

func (v *NullableReprioritizeRuleRequest) Set(val *ReprioritizeRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReprioritizeRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReprioritizeRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReprioritizeRuleRequest(val *ReprioritizeRuleRequest) *NullableReprioritizeRuleRequest {
	return &NullableReprioritizeRuleRequest{value: val, isSet: true}
}

func (v NullableReprioritizeRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReprioritizeRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


