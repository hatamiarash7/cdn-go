/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"time"
)

// checks if the DomainTransferData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainTransferData{}

// DomainTransferData struct for DomainTransferData
type DomainTransferData struct {
	Domain *string `json:"domain,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	OwnerName *string `json:"owner_name,omitempty"`
	OwnerId *string `json:"owner_id,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	Incoming *bool `json:"incoming,omitempty"`
}

// NewDomainTransferData instantiates a new DomainTransferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainTransferData() *DomainTransferData {
	this := DomainTransferData{}
	return &this
}

// NewDomainTransferDataWithDefaults instantiates a new DomainTransferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainTransferDataWithDefaults() *DomainTransferData {
	this := DomainTransferData{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DomainTransferData) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTransferData) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DomainTransferData) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DomainTransferData) SetDomain(v string) {
	o.Domain = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *DomainTransferData) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTransferData) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *DomainTransferData) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *DomainTransferData) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *DomainTransferData) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTransferData) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *DomainTransferData) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *DomainTransferData) SetAccountName(v string) {
	o.AccountName = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *DomainTransferData) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTransferData) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *DomainTransferData) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *DomainTransferData) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *DomainTransferData) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTransferData) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *DomainTransferData) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *DomainTransferData) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DomainTransferData) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTransferData) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DomainTransferData) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *DomainTransferData) SetTime(v time.Time) {
	o.Time = &v
}

// GetIncoming returns the Incoming field value if set, zero value otherwise.
func (o *DomainTransferData) GetIncoming() bool {
	if o == nil || IsNil(o.Incoming) {
		var ret bool
		return ret
	}
	return *o.Incoming
}

// GetIncomingOk returns a tuple with the Incoming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTransferData) GetIncomingOk() (*bool, bool) {
	if o == nil || IsNil(o.Incoming) {
		return nil, false
	}
	return o.Incoming, true
}

// HasIncoming returns a boolean if a field has been set.
func (o *DomainTransferData) HasIncoming() bool {
	if o != nil && !IsNil(o.Incoming) {
		return true
	}

	return false
}

// SetIncoming gets a reference to the given bool and assigns it to the Incoming field.
func (o *DomainTransferData) SetIncoming(v bool) {
	o.Incoming = &v
}

func (o DomainTransferData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainTransferData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.OwnerName) {
		toSerialize["owner_name"] = o.OwnerName
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Incoming) {
		toSerialize["incoming"] = o.Incoming
	}
	return toSerialize, nil
}

type NullableDomainTransferData struct {
	value *DomainTransferData
	isSet bool
}

func (v NullableDomainTransferData) Get() *DomainTransferData {
	return v.value
}

func (v *NullableDomainTransferData) Set(val *DomainTransferData) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainTransferData) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainTransferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainTransferData(val *DomainTransferData) *NullableDomainTransferData {
	return &NullableDomainTransferData{value: val, isSet: true}
}

func (v NullableDomainTransferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainTransferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


