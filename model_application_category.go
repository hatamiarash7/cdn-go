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

// checks if the ApplicationCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCategory{}

// ApplicationCategory struct for ApplicationCategory
type ApplicationCategory struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Active *bool `json:"active,omitempty"`
	Order *int32 `json:"order,omitempty"`
	NameTranslation NullableApplicationCategoryNameTranslation `json:"name_translation,omitempty"`
	Applications []ApplicationCategoryApplicationsInner `json:"applications,omitempty"`
}

// NewApplicationCategory instantiates a new ApplicationCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCategory() *ApplicationCategory {
	this := ApplicationCategory{}
	return &this
}

// NewApplicationCategoryWithDefaults instantiates a new ApplicationCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCategoryWithDefaults() *ApplicationCategory {
	this := ApplicationCategory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationCategory) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategory) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationCategory) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationCategory) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationCategory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationCategory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationCategory) SetName(v string) {
	o.Name = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApplicationCategory) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategory) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApplicationCategory) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApplicationCategory) SetActive(v bool) {
	o.Active = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ApplicationCategory) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategory) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ApplicationCategory) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ApplicationCategory) SetOrder(v int32) {
	o.Order = &v
}

// GetNameTranslation returns the NameTranslation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCategory) GetNameTranslation() ApplicationCategoryNameTranslation {
	if o == nil || IsNil(o.NameTranslation.Get()) {
		var ret ApplicationCategoryNameTranslation
		return ret
	}
	return *o.NameTranslation.Get()
}

// GetNameTranslationOk returns a tuple with the NameTranslation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCategory) GetNameTranslationOk() (*ApplicationCategoryNameTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.NameTranslation.Get(), o.NameTranslation.IsSet()
}

// HasNameTranslation returns a boolean if a field has been set.
func (o *ApplicationCategory) HasNameTranslation() bool {
	if o != nil && o.NameTranslation.IsSet() {
		return true
	}

	return false
}

// SetNameTranslation gets a reference to the given NullableApplicationCategoryNameTranslation and assigns it to the NameTranslation field.
func (o *ApplicationCategory) SetNameTranslation(v ApplicationCategoryNameTranslation) {
	o.NameTranslation.Set(&v)
}
// SetNameTranslationNil sets the value for NameTranslation to be an explicit nil
func (o *ApplicationCategory) SetNameTranslationNil() {
	o.NameTranslation.Set(nil)
}

// UnsetNameTranslation ensures that no value is present for NameTranslation, not even an explicit nil
func (o *ApplicationCategory) UnsetNameTranslation() {
	o.NameTranslation.Unset()
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *ApplicationCategory) GetApplications() []ApplicationCategoryApplicationsInner {
	if o == nil || IsNil(o.Applications) {
		var ret []ApplicationCategoryApplicationsInner
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategory) GetApplicationsOk() ([]ApplicationCategoryApplicationsInner, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *ApplicationCategory) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []ApplicationCategoryApplicationsInner and assigns it to the Applications field.
func (o *ApplicationCategory) SetApplications(v []ApplicationCategoryApplicationsInner) {
	o.Applications = v
}

func (o ApplicationCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.NameTranslation.IsSet() {
		toSerialize["name_translation"] = o.NameTranslation.Get()
	}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	return toSerialize, nil
}

type NullableApplicationCategory struct {
	value *ApplicationCategory
	isSet bool
}

func (v NullableApplicationCategory) Get() *ApplicationCategory {
	return v.value
}

func (v *NullableApplicationCategory) Set(val *ApplicationCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCategory(val *ApplicationCategory) *NullableApplicationCategory {
	return &NullableApplicationCategory{value: val, isSet: true}
}

func (v NullableApplicationCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


