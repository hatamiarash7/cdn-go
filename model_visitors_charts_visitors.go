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

// checks if the VisitorsChartsVisitors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisitorsChartsVisitors{}

// VisitorsChartsVisitors struct for VisitorsChartsVisitors
type VisitorsChartsVisitors struct {
	Title *string `json:"title,omitempty"`
	Categories []time.Time `json:"categories,omitempty"`
	Series []VisitorsChartsVisitorsSeriesInner `json:"series,omitempty"`
}

// NewVisitorsChartsVisitors instantiates a new VisitorsChartsVisitors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisitorsChartsVisitors() *VisitorsChartsVisitors {
	this := VisitorsChartsVisitors{}
	return &this
}

// NewVisitorsChartsVisitorsWithDefaults instantiates a new VisitorsChartsVisitors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisitorsChartsVisitorsWithDefaults() *VisitorsChartsVisitors {
	this := VisitorsChartsVisitors{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *VisitorsChartsVisitors) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitorsChartsVisitors) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *VisitorsChartsVisitors) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *VisitorsChartsVisitors) SetTitle(v string) {
	o.Title = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *VisitorsChartsVisitors) GetCategories() []time.Time {
	if o == nil || IsNil(o.Categories) {
		var ret []time.Time
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitorsChartsVisitors) GetCategoriesOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *VisitorsChartsVisitors) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []time.Time and assigns it to the Categories field.
func (o *VisitorsChartsVisitors) SetCategories(v []time.Time) {
	o.Categories = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *VisitorsChartsVisitors) GetSeries() []VisitorsChartsVisitorsSeriesInner {
	if o == nil || IsNil(o.Series) {
		var ret []VisitorsChartsVisitorsSeriesInner
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitorsChartsVisitors) GetSeriesOk() ([]VisitorsChartsVisitorsSeriesInner, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *VisitorsChartsVisitors) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []VisitorsChartsVisitorsSeriesInner and assigns it to the Series field.
func (o *VisitorsChartsVisitors) SetSeries(v []VisitorsChartsVisitorsSeriesInner) {
	o.Series = v
}

func (o VisitorsChartsVisitors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisitorsChartsVisitors) ToMap() (map[string]interface{}, error) {
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

type NullableVisitorsChartsVisitors struct {
	value *VisitorsChartsVisitors
	isSet bool
}

func (v NullableVisitorsChartsVisitors) Get() *VisitorsChartsVisitors {
	return v.value
}

func (v *NullableVisitorsChartsVisitors) Set(val *VisitorsChartsVisitors) {
	v.value = val
	v.isSet = true
}

func (v NullableVisitorsChartsVisitors) IsSet() bool {
	return v.isSet
}

func (v *NullableVisitorsChartsVisitors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisitorsChartsVisitors(val *VisitorsChartsVisitors) *NullableVisitorsChartsVisitors {
	return &NullableVisitorsChartsVisitors{value: val, isSet: true}
}

func (v NullableVisitorsChartsVisitors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisitorsChartsVisitors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


