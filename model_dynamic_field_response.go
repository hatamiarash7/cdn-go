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

// checks if the DynamicFieldResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicFieldResponse{}

// DynamicFieldResponse struct for DynamicFieldResponse
type DynamicFieldResponse struct {
	Data *DynamicField `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewDynamicFieldResponse instantiates a new DynamicFieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicFieldResponse() *DynamicFieldResponse {
	this := DynamicFieldResponse{}
	return &this
}

// NewDynamicFieldResponseWithDefaults instantiates a new DynamicFieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicFieldResponseWithDefaults() *DynamicFieldResponse {
	this := DynamicFieldResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DynamicFieldResponse) GetData() DynamicField {
	if o == nil || IsNil(o.Data) {
		var ret DynamicField
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicFieldResponse) GetDataOk() (*DynamicField, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DynamicFieldResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DynamicField and assigns it to the Data field.
func (o *DynamicFieldResponse) SetData(v DynamicField) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DynamicFieldResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DynamicFieldResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DynamicFieldResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *DynamicFieldResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *DynamicFieldResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *DynamicFieldResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o DynamicFieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicFieldResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableDynamicFieldResponse struct {
	value *DynamicFieldResponse
	isSet bool
}

func (v NullableDynamicFieldResponse) Get() *DynamicFieldResponse {
	return v.value
}

func (v *NullableDynamicFieldResponse) Set(val *DynamicFieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicFieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicFieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicFieldResponse(val *DynamicFieldResponse) *NullableDynamicFieldResponse {
	return &NullableDynamicFieldResponse{value: val, isSet: true}
}

func (v NullableDynamicFieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicFieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


