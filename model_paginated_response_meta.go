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

// checks if the PaginatedResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseMeta{}

// PaginatedResponseMeta struct for PaginatedResponseMeta
type PaginatedResponseMeta struct {
	CurrentPage *int32 `json:"current_page,omitempty"`
	From *int32 `json:"from,omitempty"`
	LastPage *int32 `json:"last_page,omitempty"`
	Path *string `json:"path,omitempty"`
	PerPage *int32 `json:"per_page,omitempty"`
	To *int32 `json:"to,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewPaginatedResponseMeta instantiates a new PaginatedResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseMeta() *PaginatedResponseMeta {
	this := PaginatedResponseMeta{}
	return &this
}

// NewPaginatedResponseMetaWithDefaults instantiates a new PaginatedResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseMetaWithDefaults() *PaginatedResponseMeta {
	this := PaginatedResponseMeta{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetCurrentPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasCurrentPage() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *PaginatedResponseMeta) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *PaginatedResponseMeta) SetFrom(v int32) {
	o.From = &v
}

// GetLastPage returns the LastPage field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetLastPage() int32 {
	if o == nil || IsNil(o.LastPage) {
		var ret int32
		return ret
	}
	return *o.LastPage
}

// GetLastPageOk returns a tuple with the LastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetLastPageOk() (*int32, bool) {
	if o == nil || IsNil(o.LastPage) {
		return nil, false
	}
	return o.LastPage, true
}

// HasLastPage returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasLastPage() bool {
	if o != nil && !IsNil(o.LastPage) {
		return true
	}

	return false
}

// SetLastPage gets a reference to the given int32 and assigns it to the LastPage field.
func (o *PaginatedResponseMeta) SetLastPage(v int32) {
	o.LastPage = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PaginatedResponseMeta) SetPath(v string) {
	o.Path = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetPerPage() int32 {
	if o == nil || IsNil(o.PerPage) {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasPerPage() bool {
	if o != nil && !IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *PaginatedResponseMeta) SetPerPage(v int32) {
	o.PerPage = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetTo() int32 {
	if o == nil || IsNil(o.To) {
		var ret int32
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetToOk() (*int32, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given int32 and assigns it to the To field.
func (o *PaginatedResponseMeta) SetTo(v int32) {
	o.To = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PaginatedResponseMeta) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseMeta) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PaginatedResponseMeta) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PaginatedResponseMeta) SetTotal(v int32) {
	o.Total = &v
}

func (o PaginatedResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentPage) {
		toSerialize["current_page"] = o.CurrentPage
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.LastPage) {
		toSerialize["last_page"] = o.LastPage
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullablePaginatedResponseMeta struct {
	value *PaginatedResponseMeta
	isSet bool
}

func (v NullablePaginatedResponseMeta) Get() *PaginatedResponseMeta {
	return v.value
}

func (v *NullablePaginatedResponseMeta) Set(val *PaginatedResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseMeta(val *PaginatedResponseMeta) *NullablePaginatedResponseMeta {
	return &NullablePaginatedResponseMeta{value: val, isSet: true}
}

func (v NullablePaginatedResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


