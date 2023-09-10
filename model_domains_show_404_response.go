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

// checks if the DomainsShow404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainsShow404Response{}

// DomainsShow404Response struct for DomainsShow404Response
type DomainsShow404Response struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewDomainsShow404Response instantiates a new DomainsShow404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainsShow404Response() *DomainsShow404Response {
	this := DomainsShow404Response{}
	return &this
}

// NewDomainsShow404ResponseWithDefaults instantiates a new DomainsShow404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainsShow404ResponseWithDefaults() *DomainsShow404Response {
	this := DomainsShow404Response{}
	var status bool = false
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DomainsShow404Response) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainsShow404Response) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DomainsShow404Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *DomainsShow404Response) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DomainsShow404Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainsShow404Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DomainsShow404Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DomainsShow404Response) SetMessage(v string) {
	o.Message = &v
}

func (o DomainsShow404Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainsShow404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableDomainsShow404Response struct {
	value *DomainsShow404Response
	isSet bool
}

func (v NullableDomainsShow404Response) Get() *DomainsShow404Response {
	return v.value
}

func (v *NullableDomainsShow404Response) Set(val *DomainsShow404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainsShow404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainsShow404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainsShow404Response(val *DomainsShow404Response) *NullableDomainsShow404Response {
	return &NullableDomainsShow404Response{value: val, isSet: true}
}

func (v NullableDomainsShow404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainsShow404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


