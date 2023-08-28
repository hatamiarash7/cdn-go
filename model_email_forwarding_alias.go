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

// checks if the EmailForwardingAlias type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingAlias{}

// EmailForwardingAlias struct for EmailForwardingAlias
type EmailForwardingAlias struct {
	Id *string `json:"id,omitempty"`
	LocalPart *string `json:"local_part,omitempty"`
	Email *string `json:"email,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
}

// NewEmailForwardingAlias instantiates a new EmailForwardingAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingAlias() *EmailForwardingAlias {
	this := EmailForwardingAlias{}
	return &this
}

// NewEmailForwardingAliasWithDefaults instantiates a new EmailForwardingAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingAliasWithDefaults() *EmailForwardingAlias {
	this := EmailForwardingAlias{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailForwardingAlias) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAlias) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailForwardingAlias) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailForwardingAlias) SetId(v string) {
	o.Id = &v
}

// GetLocalPart returns the LocalPart field value if set, zero value otherwise.
func (o *EmailForwardingAlias) GetLocalPart() string {
	if o == nil || IsNil(o.LocalPart) {
		var ret string
		return ret
	}
	return *o.LocalPart
}

// GetLocalPartOk returns a tuple with the LocalPart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAlias) GetLocalPartOk() (*string, bool) {
	if o == nil || IsNil(o.LocalPart) {
		return nil, false
	}
	return o.LocalPart, true
}

// HasLocalPart returns a boolean if a field has been set.
func (o *EmailForwardingAlias) HasLocalPart() bool {
	if o != nil && !IsNil(o.LocalPart) {
		return true
	}

	return false
}

// SetLocalPart gets a reference to the given string and assigns it to the LocalPart field.
func (o *EmailForwardingAlias) SetLocalPart(v string) {
	o.LocalPart = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmailForwardingAlias) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAlias) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailForwardingAlias) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EmailForwardingAlias) SetEmail(v string) {
	o.Email = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *EmailForwardingAlias) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAlias) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *EmailForwardingAlias) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *EmailForwardingAlias) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o EmailForwardingAlias) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingAlias) ToMap() (map[string]interface{}, error) {
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

type NullableEmailForwardingAlias struct {
	value *EmailForwardingAlias
	isSet bool
}

func (v NullableEmailForwardingAlias) Get() *EmailForwardingAlias {
	return v.value
}

func (v *NullableEmailForwardingAlias) Set(val *EmailForwardingAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingAlias(val *EmailForwardingAlias) *NullableEmailForwardingAlias {
	return &NullableEmailForwardingAlias{value: val, isSet: true}
}

func (v NullableEmailForwardingAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


