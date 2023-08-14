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

// checks if the ApplicationCategoryApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCategoryApplicationsInner{}

// ApplicationCategoryApplicationsInner struct for ApplicationCategoryApplicationsInner
type ApplicationCategoryApplicationsInner struct {
	Id *string `json:"id,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Logo *string `json:"logo,omitempty"`
}

// NewApplicationCategoryApplicationsInner instantiates a new ApplicationCategoryApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCategoryApplicationsInner() *ApplicationCategoryApplicationsInner {
	this := ApplicationCategoryApplicationsInner{}
	return &this
}

// NewApplicationCategoryApplicationsInnerWithDefaults instantiates a new ApplicationCategoryApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCategoryApplicationsInnerWithDefaults() *ApplicationCategoryApplicationsInner {
	this := ApplicationCategoryApplicationsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationCategoryApplicationsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategoryApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationCategoryApplicationsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationCategoryApplicationsInner) SetId(v string) {
	o.Id = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *ApplicationCategoryApplicationsInner) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategoryApplicationsInner) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *ApplicationCategoryApplicationsInner) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *ApplicationCategoryApplicationsInner) SetSlug(v string) {
	o.Slug = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationCategoryApplicationsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategoryApplicationsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationCategoryApplicationsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationCategoryApplicationsInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationCategoryApplicationsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategoryApplicationsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationCategoryApplicationsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationCategoryApplicationsInner) SetDescription(v string) {
	o.Description = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ApplicationCategoryApplicationsInner) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategoryApplicationsInner) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ApplicationCategoryApplicationsInner) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ApplicationCategoryApplicationsInner) SetLogo(v string) {
	o.Logo = &v
}

func (o ApplicationCategoryApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCategoryApplicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

type NullableApplicationCategoryApplicationsInner struct {
	value *ApplicationCategoryApplicationsInner
	isSet bool
}

func (v NullableApplicationCategoryApplicationsInner) Get() *ApplicationCategoryApplicationsInner {
	return v.value
}

func (v *NullableApplicationCategoryApplicationsInner) Set(val *ApplicationCategoryApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCategoryApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCategoryApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCategoryApplicationsInner(val *ApplicationCategoryApplicationsInner) *NullableApplicationCategoryApplicationsInner {
	return &NullableApplicationCategoryApplicationsInner{value: val, isSet: true}
}

func (v NullableApplicationCategoryApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCategoryApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


