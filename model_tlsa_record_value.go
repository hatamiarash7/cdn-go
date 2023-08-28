/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the TLSARecordValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TLSARecordValue{}

// TLSARecordValue struct for TLSARecordValue
type TLSARecordValue struct {
	Usage string `json:"usage"`
	Selector string `json:"selector"`
	MatchingType string `json:"matching_type"`
	Certificate string `json:"certificate"`
}

// NewTLSARecordValue instantiates a new TLSARecordValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSARecordValue(usage string, selector string, matchingType string, certificate string) *TLSARecordValue {
	this := TLSARecordValue{}
	this.Usage = usage
	this.Selector = selector
	this.MatchingType = matchingType
	this.Certificate = certificate
	return &this
}

// NewTLSARecordValueWithDefaults instantiates a new TLSARecordValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSARecordValueWithDefaults() *TLSARecordValue {
	this := TLSARecordValue{}
	return &this
}

// GetUsage returns the Usage field value
func (o *TLSARecordValue) GetUsage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *TLSARecordValue) GetUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *TLSARecordValue) SetUsage(v string) {
	o.Usage = v
}

// GetSelector returns the Selector field value
func (o *TLSARecordValue) GetSelector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *TLSARecordValue) GetSelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *TLSARecordValue) SetSelector(v string) {
	o.Selector = v
}

// GetMatchingType returns the MatchingType field value
func (o *TLSARecordValue) GetMatchingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchingType
}

// GetMatchingTypeOk returns a tuple with the MatchingType field value
// and a boolean to check if the value has been set.
func (o *TLSARecordValue) GetMatchingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchingType, true
}

// SetMatchingType sets field value
func (o *TLSARecordValue) SetMatchingType(v string) {
	o.MatchingType = v
}

// GetCertificate returns the Certificate field value
func (o *TLSARecordValue) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *TLSARecordValue) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *TLSARecordValue) SetCertificate(v string) {
	o.Certificate = v
}

func (o TLSARecordValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TLSARecordValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["usage"] = o.Usage
	toSerialize["selector"] = o.Selector
	toSerialize["matching_type"] = o.MatchingType
	toSerialize["certificate"] = o.Certificate
	return toSerialize, nil
}

type NullableTLSARecordValue struct {
	value *TLSARecordValue
	isSet bool
}

func (v NullableTLSARecordValue) Get() *TLSARecordValue {
	return v.value
}

func (v *NullableTLSARecordValue) Set(val *TLSARecordValue) {
	v.value = val
	v.isSet = true
}

func (v NullableTLSARecordValue) IsSet() bool {
	return v.isSet
}

func (v *NullableTLSARecordValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTLSARecordValue(val *TLSARecordValue) *NullableTLSARecordValue {
	return &NullableTLSARecordValue{value: val, isSet: true}
}

func (v NullableTLSARecordValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTLSARecordValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


