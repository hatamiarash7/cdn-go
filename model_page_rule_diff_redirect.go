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

// checks if the PageRuleDiffRedirect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageRuleDiffRedirect{}

// PageRuleDiffRedirect struct for PageRuleDiffRedirect
type PageRuleDiffRedirect struct {
	Enable *bool `json:"enable,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
	Url NullableString `json:"url,omitempty"`
}

// NewPageRuleDiffRedirect instantiates a new PageRuleDiffRedirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRuleDiffRedirect() *PageRuleDiffRedirect {
	this := PageRuleDiffRedirect{}
	return &this
}

// NewPageRuleDiffRedirectWithDefaults instantiates a new PageRuleDiffRedirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRuleDiffRedirectWithDefaults() *PageRuleDiffRedirect {
	this := PageRuleDiffRedirect{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *PageRuleDiffRedirect) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleDiffRedirect) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *PageRuleDiffRedirect) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *PageRuleDiffRedirect) SetEnable(v bool) {
	o.Enable = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *PageRuleDiffRedirect) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleDiffRedirect) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *PageRuleDiffRedirect) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *PageRuleDiffRedirect) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PageRuleDiffRedirect) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PageRuleDiffRedirect) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PageRuleDiffRedirect) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PageRuleDiffRedirect) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PageRuleDiffRedirect) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PageRuleDiffRedirect) UnsetUrl() {
	o.Url.Unset()
}

func (o PageRuleDiffRedirect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageRuleDiffRedirect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullablePageRuleDiffRedirect struct {
	value *PageRuleDiffRedirect
	isSet bool
}

func (v NullablePageRuleDiffRedirect) Get() *PageRuleDiffRedirect {
	return v.value
}

func (v *NullablePageRuleDiffRedirect) Set(val *PageRuleDiffRedirect) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRuleDiffRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRuleDiffRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRuleDiffRedirect(val *PageRuleDiffRedirect) *NullablePageRuleDiffRedirect {
	return &NullablePageRuleDiffRedirect{value: val, isSet: true}
}

func (v NullablePageRuleDiffRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRuleDiffRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


