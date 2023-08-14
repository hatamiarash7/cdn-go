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

// checks if the TransportLayerProxyFirewallsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxyFirewallsInner{}

// TransportLayerProxyFirewallsInner struct for TransportLayerProxyFirewallsInner
type TransportLayerProxyFirewallsInner struct {
	Access *string `json:"access,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Match *TransportLayerProxyMatch `json:"match,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewTransportLayerProxyFirewallsInner instantiates a new TransportLayerProxyFirewallsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxyFirewallsInner() *TransportLayerProxyFirewallsInner {
	this := TransportLayerProxyFirewallsInner{}
	return &this
}

// NewTransportLayerProxyFirewallsInnerWithDefaults instantiates a new TransportLayerProxyFirewallsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyFirewallsInnerWithDefaults() *TransportLayerProxyFirewallsInner {
	this := TransportLayerProxyFirewallsInner{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewallsInner) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewallsInner) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewallsInner) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *TransportLayerProxyFirewallsInner) SetAccess(v string) {
	o.Access = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewallsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewallsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewallsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransportLayerProxyFirewallsInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewallsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewallsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewallsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransportLayerProxyFirewallsInner) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewallsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewallsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewallsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransportLayerProxyFirewallsInner) SetType(v string) {
	o.Type = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewallsInner) GetMatch() TransportLayerProxyMatch {
	if o == nil || IsNil(o.Match) {
		var ret TransportLayerProxyMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewallsInner) GetMatchOk() (*TransportLayerProxyMatch, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewallsInner) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given TransportLayerProxyMatch and assigns it to the Match field.
func (o *TransportLayerProxyFirewallsInner) SetMatch(v TransportLayerProxyMatch) {
	o.Match = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewallsInner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewallsInner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewallsInner) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *TransportLayerProxyFirewallsInner) SetActive(v bool) {
	o.Active = &v
}

func (o TransportLayerProxyFirewallsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxyFirewallsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableTransportLayerProxyFirewallsInner struct {
	value *TransportLayerProxyFirewallsInner
	isSet bool
}

func (v NullableTransportLayerProxyFirewallsInner) Get() *TransportLayerProxyFirewallsInner {
	return v.value
}

func (v *NullableTransportLayerProxyFirewallsInner) Set(val *TransportLayerProxyFirewallsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyFirewallsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyFirewallsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyFirewallsInner(val *TransportLayerProxyFirewallsInner) *NullableTransportLayerProxyFirewallsInner {
	return &NullableTransportLayerProxyFirewallsInner{value: val, isSet: true}
}

func (v NullableTransportLayerProxyFirewallsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyFirewallsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


