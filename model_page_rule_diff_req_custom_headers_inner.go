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

// checks if the PageRuleDiffReqCustomHeadersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageRuleDiffReqCustomHeadersInner{}

// PageRuleDiffReqCustomHeadersInner struct for PageRuleDiffReqCustomHeadersInner
type PageRuleDiffReqCustomHeadersInner struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	IsVar *bool `json:"is_var,omitempty"`
}

// NewPageRuleDiffReqCustomHeadersInner instantiates a new PageRuleDiffReqCustomHeadersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRuleDiffReqCustomHeadersInner() *PageRuleDiffReqCustomHeadersInner {
	this := PageRuleDiffReqCustomHeadersInner{}
	return &this
}

// NewPageRuleDiffReqCustomHeadersInnerWithDefaults instantiates a new PageRuleDiffReqCustomHeadersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRuleDiffReqCustomHeadersInnerWithDefaults() *PageRuleDiffReqCustomHeadersInner {
	this := PageRuleDiffReqCustomHeadersInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PageRuleDiffReqCustomHeadersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleDiffReqCustomHeadersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PageRuleDiffReqCustomHeadersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PageRuleDiffReqCustomHeadersInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PageRuleDiffReqCustomHeadersInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleDiffReqCustomHeadersInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PageRuleDiffReqCustomHeadersInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PageRuleDiffReqCustomHeadersInner) SetValue(v string) {
	o.Value = &v
}

// GetIsVar returns the IsVar field value if set, zero value otherwise.
func (o *PageRuleDiffReqCustomHeadersInner) GetIsVar() bool {
	if o == nil || IsNil(o.IsVar) {
		var ret bool
		return ret
	}
	return *o.IsVar
}

// GetIsVarOk returns a tuple with the IsVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleDiffReqCustomHeadersInner) GetIsVarOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVar) {
		return nil, false
	}
	return o.IsVar, true
}

// HasIsVar returns a boolean if a field has been set.
func (o *PageRuleDiffReqCustomHeadersInner) HasIsVar() bool {
	if o != nil && !IsNil(o.IsVar) {
		return true
	}

	return false
}

// SetIsVar gets a reference to the given bool and assigns it to the IsVar field.
func (o *PageRuleDiffReqCustomHeadersInner) SetIsVar(v bool) {
	o.IsVar = &v
}

func (o PageRuleDiffReqCustomHeadersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageRuleDiffReqCustomHeadersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.IsVar) {
		toSerialize["is_var"] = o.IsVar
	}
	return toSerialize, nil
}

type NullablePageRuleDiffReqCustomHeadersInner struct {
	value *PageRuleDiffReqCustomHeadersInner
	isSet bool
}

func (v NullablePageRuleDiffReqCustomHeadersInner) Get() *PageRuleDiffReqCustomHeadersInner {
	return v.value
}

func (v *NullablePageRuleDiffReqCustomHeadersInner) Set(val *PageRuleDiffReqCustomHeadersInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRuleDiffReqCustomHeadersInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRuleDiffReqCustomHeadersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRuleDiffReqCustomHeadersInner(val *PageRuleDiffReqCustomHeadersInner) *NullablePageRuleDiffReqCustomHeadersInner {
	return &NullablePageRuleDiffReqCustomHeadersInner{value: val, isSet: true}
}

func (v NullablePageRuleDiffReqCustomHeadersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRuleDiffReqCustomHeadersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


