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

// checks if the ReportsAttacksShow200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsAttacksShow200Response{}

// ReportsAttacksShow200Response struct for ReportsAttacksShow200Response
type ReportsAttacksShow200Response struct {
	Data *AttackReport `json:"data,omitempty"`
}

// NewReportsAttacksShow200Response instantiates a new ReportsAttacksShow200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsAttacksShow200Response() *ReportsAttacksShow200Response {
	this := ReportsAttacksShow200Response{}
	return &this
}

// NewReportsAttacksShow200ResponseWithDefaults instantiates a new ReportsAttacksShow200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsAttacksShow200ResponseWithDefaults() *ReportsAttacksShow200Response {
	this := ReportsAttacksShow200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReportsAttacksShow200Response) GetData() AttackReport {
	if o == nil || IsNil(o.Data) {
		var ret AttackReport
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsAttacksShow200Response) GetDataOk() (*AttackReport, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReportsAttacksShow200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AttackReport and assigns it to the Data field.
func (o *ReportsAttacksShow200Response) SetData(v AttackReport) {
	o.Data = &v
}

func (o ReportsAttacksShow200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsAttacksShow200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableReportsAttacksShow200Response struct {
	value *ReportsAttacksShow200Response
	isSet bool
}

func (v NullableReportsAttacksShow200Response) Get() *ReportsAttacksShow200Response {
	return v.value
}

func (v *NullableReportsAttacksShow200Response) Set(val *ReportsAttacksShow200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsAttacksShow200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsAttacksShow200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsAttacksShow200Response(val *ReportsAttacksShow200Response) *NullableReportsAttacksShow200Response {
	return &NullableReportsAttacksShow200Response{value: val, isSet: true}
}

func (v NullableReportsAttacksShow200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsAttacksShow200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


