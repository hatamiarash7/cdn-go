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

// checks if the EmailForwardingAliasesListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingAliasesListInner{}

// EmailForwardingAliasesListInner struct for EmailForwardingAliasesListInner
type EmailForwardingAliasesListInner struct {
	Id *string `json:"id,omitempty"`
	LocalPart *string `json:"local_part,omitempty"`
	Email *string `json:"email,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
}

// NewEmailForwardingAliasesListInner instantiates a new EmailForwardingAliasesListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingAliasesListInner() *EmailForwardingAliasesListInner {
	this := EmailForwardingAliasesListInner{}
	return &this
}

// NewEmailForwardingAliasesListInnerWithDefaults instantiates a new EmailForwardingAliasesListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingAliasesListInnerWithDefaults() *EmailForwardingAliasesListInner {
	this := EmailForwardingAliasesListInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailForwardingAliasesListInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAliasesListInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailForwardingAliasesListInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailForwardingAliasesListInner) SetId(v string) {
	o.Id = &v
}

// GetLocalPart returns the LocalPart field value if set, zero value otherwise.
func (o *EmailForwardingAliasesListInner) GetLocalPart() string {
	if o == nil || IsNil(o.LocalPart) {
		var ret string
		return ret
	}
	return *o.LocalPart
}

// GetLocalPartOk returns a tuple with the LocalPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAliasesListInner) GetLocalPartOk() (*string, bool) {
	if o == nil || IsNil(o.LocalPart) {
		return nil, false
	}
	return o.LocalPart, true
}

// HasLocalPart returns a boolean if a field has been set.
func (o *EmailForwardingAliasesListInner) HasLocalPart() bool {
	if o != nil && !IsNil(o.LocalPart) {
		return true
	}

	return false
}

// SetLocalPart gets a reference to the given string and assigns it to the LocalPart field.
func (o *EmailForwardingAliasesListInner) SetLocalPart(v string) {
	o.LocalPart = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmailForwardingAliasesListInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAliasesListInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailForwardingAliasesListInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EmailForwardingAliasesListInner) SetEmail(v string) {
	o.Email = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *EmailForwardingAliasesListInner) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAliasesListInner) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *EmailForwardingAliasesListInner) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *EmailForwardingAliasesListInner) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o EmailForwardingAliasesListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingAliasesListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LocalPart) {
		toSerialize["local_part"] = o.LocalPart
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	return toSerialize, nil
}

type NullableEmailForwardingAliasesListInner struct {
	value *EmailForwardingAliasesListInner
	isSet bool
}

func (v NullableEmailForwardingAliasesListInner) Get() *EmailForwardingAliasesListInner {
	return v.value
}

func (v *NullableEmailForwardingAliasesListInner) Set(val *EmailForwardingAliasesListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingAliasesListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingAliasesListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingAliasesListInner(val *EmailForwardingAliasesListInner) *NullableEmailForwardingAliasesListInner {
	return &NullableEmailForwardingAliasesListInner{value: val, isSet: true}
}

func (v NullableEmailForwardingAliasesListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingAliasesListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


