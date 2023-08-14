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

// checks if the ResponseTimeChartsIr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTimeChartsIr{}

// ResponseTimeChartsIr struct for ResponseTimeChartsIr
type ResponseTimeChartsIr struct {
	Title *string `json:"title,omitempty"`
	Categories []time.Time `json:"categories,omitempty"`
	Series []ResponseTimeChartsIrSeriesInner `json:"series,omitempty"`
}

// NewResponseTimeChartsIr instantiates a new ResponseTimeChartsIr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTimeChartsIr() *ResponseTimeChartsIr {
	this := ResponseTimeChartsIr{}
	return &this
}

// NewResponseTimeChartsIrWithDefaults instantiates a new ResponseTimeChartsIr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTimeChartsIrWithDefaults() *ResponseTimeChartsIr {
	this := ResponseTimeChartsIr{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ResponseTimeChartsIr) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeChartsIr) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ResponseTimeChartsIr) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ResponseTimeChartsIr) SetTitle(v string) {
	o.Title = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *ResponseTimeChartsIr) GetCategories() []time.Time {
	if o == nil || IsNil(o.Categories) {
		var ret []time.Time
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeChartsIr) GetCategoriesOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *ResponseTimeChartsIr) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []time.Time and assigns it to the Categories field.
func (o *ResponseTimeChartsIr) SetCategories(v []time.Time) {
	o.Categories = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *ResponseTimeChartsIr) GetSeries() []ResponseTimeChartsIrSeriesInner {
	if o == nil || IsNil(o.Series) {
		var ret []ResponseTimeChartsIrSeriesInner
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeChartsIr) GetSeriesOk() ([]ResponseTimeChartsIrSeriesInner, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *ResponseTimeChartsIr) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []ResponseTimeChartsIrSeriesInner and assigns it to the Series field.
func (o *ResponseTimeChartsIr) SetSeries(v []ResponseTimeChartsIrSeriesInner) {
	o.Series = v
}

func (o ResponseTimeChartsIr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTimeChartsIr) ToMap() (map[string]interface{}, error) {
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

type NullableResponseTimeChartsIr struct {
	value *ResponseTimeChartsIr
	isSet bool
}

func (v NullableResponseTimeChartsIr) Get() *ResponseTimeChartsIr {
	return v.value
}

func (v *NullableResponseTimeChartsIr) Set(val *ResponseTimeChartsIr) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTimeChartsIr) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTimeChartsIr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTimeChartsIr(val *ResponseTimeChartsIr) *NullableResponseTimeChartsIr {
	return &NullableResponseTimeChartsIr{value: val, isSet: true}
}

func (v NullableResponseTimeChartsIr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTimeChartsIr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


