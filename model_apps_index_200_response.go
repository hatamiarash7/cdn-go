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

// checks if the AppsIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsIndex200Response{}

// AppsIndex200Response struct for AppsIndex200Response
type AppsIndex200Response struct {
	Data []CdnApp `json:"data,omitempty"`
	Links *PaginatedResponseLinks `json:"links,omitempty"`
	Meta *PaginatedResponseMeta `json:"meta,omitempty"`
}

// NewAppsIndex200Response instantiates a new AppsIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsIndex200Response() *AppsIndex200Response {
	this := AppsIndex200Response{}
	return &this
}

// NewAppsIndex200ResponseWithDefaults instantiates a new AppsIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsIndex200ResponseWithDefaults() *AppsIndex200Response {
	this := AppsIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppsIndex200Response) GetData() []CdnApp {
	if o == nil || IsNil(o.Data) {
		var ret []CdnApp
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsIndex200Response) GetDataOk() ([]CdnApp, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppsIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CdnApp and assigns it to the Data field.
func (o *AppsIndex200Response) SetData(v []CdnApp) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppsIndex200Response) GetLinks() PaginatedResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsIndex200Response) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppsIndex200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedResponseLinks and assigns it to the Links field.
func (o *AppsIndex200Response) SetLinks(v PaginatedResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppsIndex200Response) GetMeta() PaginatedResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatedResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsIndex200Response) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppsIndex200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatedResponseMeta and assigns it to the Meta field.
func (o *AppsIndex200Response) SetMeta(v PaginatedResponseMeta) {
	o.Meta = &v
}

func (o AppsIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableAppsIndex200Response struct {
	value *AppsIndex200Response
	isSet bool
}

func (v NullableAppsIndex200Response) Get() *AppsIndex200Response {
	return v.value
}

func (v *NullableAppsIndex200Response) Set(val *AppsIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsIndex200Response(val *AppsIndex200Response) *NullableAppsIndex200Response {
	return &NullableAppsIndex200Response{value: val, isSet: true}
}

func (v NullableAppsIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


