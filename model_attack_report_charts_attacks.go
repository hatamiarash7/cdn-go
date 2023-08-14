/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.104.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
	"time"
)

// checks if the AttackReportChartsAttacks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportChartsAttacks{}

// AttackReportChartsAttacks struct for AttackReportChartsAttacks
type AttackReportChartsAttacks struct {
	Title *string `json:"title,omitempty"`
	Categories []time.Time `json:"categories,omitempty"`
	Series []AttackReportChartsAttacksSeriesInner `json:"series,omitempty"`
}

// NewAttackReportChartsAttacks instantiates a new AttackReportChartsAttacks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportChartsAttacks() *AttackReportChartsAttacks {
	this := AttackReportChartsAttacks{}
	return &this
}

// NewAttackReportChartsAttacksWithDefaults instantiates a new AttackReportChartsAttacks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportChartsAttacksWithDefaults() *AttackReportChartsAttacks {
	this := AttackReportChartsAttacks{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AttackReportChartsAttacks) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportChartsAttacks) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AttackReportChartsAttacks) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AttackReportChartsAttacks) SetTitle(v string) {
	o.Title = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *AttackReportChartsAttacks) GetCategories() []time.Time {
	if o == nil || IsNil(o.Categories) {
		var ret []time.Time
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportChartsAttacks) GetCategoriesOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *AttackReportChartsAttacks) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []time.Time and assigns it to the Categories field.
func (o *AttackReportChartsAttacks) SetCategories(v []time.Time) {
	o.Categories = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *AttackReportChartsAttacks) GetSeries() []AttackReportChartsAttacksSeriesInner {
	if o == nil || IsNil(o.Series) {
		var ret []AttackReportChartsAttacksSeriesInner
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportChartsAttacks) GetSeriesOk() ([]AttackReportChartsAttacksSeriesInner, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *AttackReportChartsAttacks) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []AttackReportChartsAttacksSeriesInner and assigns it to the Series field.
func (o *AttackReportChartsAttacks) SetSeries(v []AttackReportChartsAttacksSeriesInner) {
	o.Series = v
}

func (o AttackReportChartsAttacks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportChartsAttacks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	return toSerialize, nil
}

type NullableAttackReportChartsAttacks struct {
	value *AttackReportChartsAttacks
	isSet bool
}

func (v NullableAttackReportChartsAttacks) Get() *AttackReportChartsAttacks {
	return v.value
}

func (v *NullableAttackReportChartsAttacks) Set(val *AttackReportChartsAttacks) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportChartsAttacks) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportChartsAttacks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportChartsAttacks(val *AttackReportChartsAttacks) *NullableAttackReportChartsAttacks {
	return &NullableAttackReportChartsAttacks{value: val, isSet: true}
}

func (v NullableAttackReportChartsAttacks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportChartsAttacks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


