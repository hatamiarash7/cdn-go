/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.104.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the RateLimitingRulesUpdate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimitingRulesUpdate200Response{}

// RateLimitingRulesUpdate200Response struct for RateLimitingRulesUpdate200Response
type RateLimitingRulesUpdate200Response struct {
	Data *RateLimitRuleView `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewRateLimitingRulesUpdate200Response instantiates a new RateLimitingRulesUpdate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitingRulesUpdate200Response() *RateLimitingRulesUpdate200Response {
	this := RateLimitingRulesUpdate200Response{}
	return &this
}

// NewRateLimitingRulesUpdate200ResponseWithDefaults instantiates a new RateLimitingRulesUpdate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitingRulesUpdate200ResponseWithDefaults() *RateLimitingRulesUpdate200Response {
	this := RateLimitingRulesUpdate200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RateLimitingRulesUpdate200Response) GetData() RateLimitRuleView {
	if o == nil || IsNil(o.Data) {
		var ret RateLimitRuleView
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitingRulesUpdate200Response) GetDataOk() (*RateLimitRuleView, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RateLimitingRulesUpdate200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RateLimitRuleView and assigns it to the Data field.
func (o *RateLimitingRulesUpdate200Response) SetData(v RateLimitRuleView) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RateLimitingRulesUpdate200Response) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RateLimitingRulesUpdate200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *RateLimitingRulesUpdate200Response) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *RateLimitingRulesUpdate200Response) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *RateLimitingRulesUpdate200Response) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *RateLimitingRulesUpdate200Response) UnsetMessage() {
	o.Message.Unset()
}

func (o RateLimitingRulesUpdate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitingRulesUpdate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableRateLimitingRulesUpdate200Response struct {
	value *RateLimitingRulesUpdate200Response
	isSet bool
}

func (v NullableRateLimitingRulesUpdate200Response) Get() *RateLimitingRulesUpdate200Response {
	return v.value
}

func (v *NullableRateLimitingRulesUpdate200Response) Set(val *RateLimitingRulesUpdate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitingRulesUpdate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitingRulesUpdate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitingRulesUpdate200Response(val *RateLimitingRulesUpdate200Response) *NullableRateLimitingRulesUpdate200Response {
	return &NullableRateLimitingRulesUpdate200Response{value: val, isSet: true}
}

func (v NullableRateLimitingRulesUpdate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitingRulesUpdate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


