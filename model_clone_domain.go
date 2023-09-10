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

// checks if the CloneDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloneDomain{}

// CloneDomain struct for CloneDomain
type CloneDomain struct {
	From string `json:"from"`
}

// NewCloneDomain instantiates a new CloneDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneDomain(from string) *CloneDomain {
	this := CloneDomain{}
	this.From = from
	return &this
}

// NewCloneDomainWithDefaults instantiates a new CloneDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneDomainWithDefaults() *CloneDomain {
	this := CloneDomain{}
	return &this
}

// GetFrom returns the From field value
func (o *CloneDomain) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CloneDomain) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CloneDomain) SetFrom(v string) {
	o.From = v
}

func (o CloneDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloneDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	return toSerialize, nil
}

type NullableCloneDomain struct {
	value *CloneDomain
	isSet bool
}

func (v NullableCloneDomain) Get() *CloneDomain {
	return v.value
}

func (v *NullableCloneDomain) Set(val *CloneDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneDomain(val *CloneDomain) *NullableCloneDomain {
	return &NullableCloneDomain{value: val, isSet: true}
}

func (v NullableCloneDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


