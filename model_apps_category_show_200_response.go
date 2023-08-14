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

// checks if the AppsCategoryShow200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsCategoryShow200Response{}

// AppsCategoryShow200Response struct for AppsCategoryShow200Response
type AppsCategoryShow200Response struct {
	Data *ApplicationCategory `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewAppsCategoryShow200Response instantiates a new AppsCategoryShow200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsCategoryShow200Response() *AppsCategoryShow200Response {
	this := AppsCategoryShow200Response{}
	return &this
}

// NewAppsCategoryShow200ResponseWithDefaults instantiates a new AppsCategoryShow200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsCategoryShow200ResponseWithDefaults() *AppsCategoryShow200Response {
	this := AppsCategoryShow200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppsCategoryShow200Response) GetData() ApplicationCategory {
	if o == nil || IsNil(o.Data) {
		var ret ApplicationCategory
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsCategoryShow200Response) GetDataOk() (*ApplicationCategory, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppsCategoryShow200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ApplicationCategory and assigns it to the Data field.
func (o *AppsCategoryShow200Response) SetData(v ApplicationCategory) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppsCategoryShow200Response) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppsCategoryShow200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *AppsCategoryShow200Response) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *AppsCategoryShow200Response) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *AppsCategoryShow200Response) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *AppsCategoryShow200Response) UnsetMessage() {
	o.Message.Unset()
}

func (o AppsCategoryShow200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsCategoryShow200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableAppsCategoryShow200Response struct {
	value *AppsCategoryShow200Response
	isSet bool
}

func (v NullableAppsCategoryShow200Response) Get() *AppsCategoryShow200Response {
	return v.value
}

func (v *NullableAppsCategoryShow200Response) Set(val *AppsCategoryShow200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsCategoryShow200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsCategoryShow200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsCategoryShow200Response(val *AppsCategoryShow200Response) *NullableAppsCategoryShow200Response {
	return &NullableAppsCategoryShow200Response{value: val, isSet: true}
}

func (v NullableAppsCategoryShow200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsCategoryShow200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


