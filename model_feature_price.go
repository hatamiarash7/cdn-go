/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.103.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the FeaturePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturePrice{}

// FeaturePrice struct for FeaturePrice
type FeaturePrice struct {
	MetricKey *string `json:"metric_key,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

// NewFeaturePrice instantiates a new FeaturePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturePrice() *FeaturePrice {
	this := FeaturePrice{}
	return &this
}

// NewFeaturePriceWithDefaults instantiates a new FeaturePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturePriceWithDefaults() *FeaturePrice {
	this := FeaturePrice{}
	return &this
}

// GetMetricKey returns the MetricKey field value if set, zero value otherwise.
func (o *FeaturePrice) GetMetricKey() string {
	if o == nil || IsNil(o.MetricKey) {
		var ret string
		return ret
	}
	return *o.MetricKey
}

// GetMetricKeyOk returns a tuple with the MetricKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturePrice) GetMetricKeyOk() (*string, bool) {
	if o == nil || IsNil(o.MetricKey) {
		return nil, false
	}
	return o.MetricKey, true
}

// HasMetricKey returns a boolean if a field has been set.
func (o *FeaturePrice) HasMetricKey() bool {
	if o != nil && !IsNil(o.MetricKey) {
		return true
	}

	return false
}

// SetMetricKey gets a reference to the given string and assigns it to the MetricKey field.
func (o *FeaturePrice) SetMetricKey(v string) {
	o.MetricKey = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *FeaturePrice) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturePrice) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *FeaturePrice) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *FeaturePrice) SetCurrency(v string) {
	o.Currency = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FeaturePrice) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturePrice) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FeaturePrice) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *FeaturePrice) SetValue(v float32) {
	o.Value = &v
}

func (o FeaturePrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeaturePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetricKey) {
		toSerialize["metric_key"] = o.MetricKey
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFeaturePrice struct {
	value *FeaturePrice
	isSet bool
}

func (v NullableFeaturePrice) Get() *FeaturePrice {
	return v.value
}

func (v *NullableFeaturePrice) Set(val *FeaturePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturePrice(val *FeaturePrice) *NullableFeaturePrice {
	return &NullableFeaturePrice{value: val, isSet: true}
}

func (v NullableFeaturePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


