/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.103.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
	"time"
)

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain struct for Domain
type Domain struct {
	Id *string `json:"id,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	// Deprecated in favor of name attribute
	// Deprecated
	Domain *string `json:"domain,omitempty"`
	Name *string `json:"name,omitempty"`
	// - `0` - Traffic - `1` - Basic - `2` - Growth - `3` - Professional - `4` - Enterprise 
	PlanLevel *int32 `json:"plan_level,omitempty"`
	// Desired NS values for the domain
	NsKeys []string `json:"ns_keys,omitempty"`
	// Current NS values for the domain
	CurrentNs []string `json:"current_ns,omitempty"`
	// Current record for CNAME Setup of the domain
	TargetCname NullableString `json:"target_cname,omitempty"`
	// Domain's custom record for CNAME Setup
	CustomCname NullableString `json:"custom_cname,omitempty"`
	// Partial domain is using CNAME Setup and full domain is using NS Setup
	Type *string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
	DnsCloud *bool `json:"dns_cloud,omitempty"`
	Restriction []string `json:"restriction,omitempty"`
	Transfer *DomainTransferData `json:"transfer,omitempty"`
	FingerprintStatus *bool `json:"fingerprint_status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	var dnsCloud bool = false
	this.DnsCloud = &dnsCloud
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	var dnsCloud bool = false
	this.DnsCloud = &dnsCloud
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Domain) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Domain) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Domain) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Domain) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Domain) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Domain) SetUserId(v string) {
	o.UserId = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
// Deprecated
func (o *Domain) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Domain) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Domain) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
// Deprecated
func (o *Domain) SetDomain(v string) {
	o.Domain = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Domain) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Domain) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Domain) SetName(v string) {
	o.Name = &v
}

// GetPlanLevel returns the PlanLevel field value if set, zero value otherwise.
func (o *Domain) GetPlanLevel() int32 {
	if o == nil || IsNil(o.PlanLevel) {
		var ret int32
		return ret
	}
	return *o.PlanLevel
}

// GetPlanLevelOk returns a tuple with the PlanLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetPlanLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.PlanLevel) {
		return nil, false
	}
	return o.PlanLevel, true
}

// HasPlanLevel returns a boolean if a field has been set.
func (o *Domain) HasPlanLevel() bool {
	if o != nil && !IsNil(o.PlanLevel) {
		return true
	}

	return false
}

// SetPlanLevel gets a reference to the given int32 and assigns it to the PlanLevel field.
func (o *Domain) SetPlanLevel(v int32) {
	o.PlanLevel = &v
}

// GetNsKeys returns the NsKeys field value if set, zero value otherwise.
func (o *Domain) GetNsKeys() []string {
	if o == nil || IsNil(o.NsKeys) {
		var ret []string
		return ret
	}
	return o.NsKeys
}

// GetNsKeysOk returns a tuple with the NsKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetNsKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.NsKeys) {
		return nil, false
	}
	return o.NsKeys, true
}

// HasNsKeys returns a boolean if a field has been set.
func (o *Domain) HasNsKeys() bool {
	if o != nil && !IsNil(o.NsKeys) {
		return true
	}

	return false
}

// SetNsKeys gets a reference to the given []string and assigns it to the NsKeys field.
func (o *Domain) SetNsKeys(v []string) {
	o.NsKeys = v
}

// GetCurrentNs returns the CurrentNs field value if set, zero value otherwise.
func (o *Domain) GetCurrentNs() []string {
	if o == nil || IsNil(o.CurrentNs) {
		var ret []string
		return ret
	}
	return o.CurrentNs
}

// GetCurrentNsOk returns a tuple with the CurrentNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCurrentNsOk() ([]string, bool) {
	if o == nil || IsNil(o.CurrentNs) {
		return nil, false
	}
	return o.CurrentNs, true
}

// HasCurrentNs returns a boolean if a field has been set.
func (o *Domain) HasCurrentNs() bool {
	if o != nil && !IsNil(o.CurrentNs) {
		return true
	}

	return false
}

// SetCurrentNs gets a reference to the given []string and assigns it to the CurrentNs field.
func (o *Domain) SetCurrentNs(v []string) {
	o.CurrentNs = v
}

// GetTargetCname returns the TargetCname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Domain) GetTargetCname() string {
	if o == nil || IsNil(o.TargetCname.Get()) {
		var ret string
		return ret
	}
	return *o.TargetCname.Get()
}

// GetTargetCnameOk returns a tuple with the TargetCname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetTargetCnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetCname.Get(), o.TargetCname.IsSet()
}

// HasTargetCname returns a boolean if a field has been set.
func (o *Domain) HasTargetCname() bool {
	if o != nil && o.TargetCname.IsSet() {
		return true
	}

	return false
}

// SetTargetCname gets a reference to the given NullableString and assigns it to the TargetCname field.
func (o *Domain) SetTargetCname(v string) {
	o.TargetCname.Set(&v)
}
// SetTargetCnameNil sets the value for TargetCname to be an explicit nil
func (o *Domain) SetTargetCnameNil() {
	o.TargetCname.Set(nil)
}

// UnsetTargetCname ensures that no value is present for TargetCname, not even an explicit nil
func (o *Domain) UnsetTargetCname() {
	o.TargetCname.Unset()
}

// GetCustomCname returns the CustomCname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Domain) GetCustomCname() string {
	if o == nil || IsNil(o.CustomCname.Get()) {
		var ret string
		return ret
	}
	return *o.CustomCname.Get()
}

// GetCustomCnameOk returns a tuple with the CustomCname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetCustomCnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomCname.Get(), o.CustomCname.IsSet()
}

// HasCustomCname returns a boolean if a field has been set.
func (o *Domain) HasCustomCname() bool {
	if o != nil && o.CustomCname.IsSet() {
		return true
	}

	return false
}

// SetCustomCname gets a reference to the given NullableString and assigns it to the CustomCname field.
func (o *Domain) SetCustomCname(v string) {
	o.CustomCname.Set(&v)
}
// SetCustomCnameNil sets the value for CustomCname to be an explicit nil
func (o *Domain) SetCustomCnameNil() {
	o.CustomCname.Set(nil)
}

// UnsetCustomCname ensures that no value is present for CustomCname, not even an explicit nil
func (o *Domain) UnsetCustomCname() {
	o.CustomCname.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Domain) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Domain) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Domain) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Domain) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Domain) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Domain) SetStatus(v string) {
	o.Status = &v
}

// GetDnsCloud returns the DnsCloud field value if set, zero value otherwise.
func (o *Domain) GetDnsCloud() bool {
	if o == nil || IsNil(o.DnsCloud) {
		var ret bool
		return ret
	}
	return *o.DnsCloud
}

// GetDnsCloudOk returns a tuple with the DnsCloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDnsCloudOk() (*bool, bool) {
	if o == nil || IsNil(o.DnsCloud) {
		return nil, false
	}
	return o.DnsCloud, true
}

// HasDnsCloud returns a boolean if a field has been set.
func (o *Domain) HasDnsCloud() bool {
	if o != nil && !IsNil(o.DnsCloud) {
		return true
	}

	return false
}

// SetDnsCloud gets a reference to the given bool and assigns it to the DnsCloud field.
func (o *Domain) SetDnsCloud(v bool) {
	o.DnsCloud = &v
}

// GetRestriction returns the Restriction field value if set, zero value otherwise.
func (o *Domain) GetRestriction() []string {
	if o == nil || IsNil(o.Restriction) {
		var ret []string
		return ret
	}
	return o.Restriction
}

// GetRestrictionOk returns a tuple with the Restriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRestrictionOk() ([]string, bool) {
	if o == nil || IsNil(o.Restriction) {
		return nil, false
	}
	return o.Restriction, true
}

// HasRestriction returns a boolean if a field has been set.
func (o *Domain) HasRestriction() bool {
	if o != nil && !IsNil(o.Restriction) {
		return true
	}

	return false
}

// SetRestriction gets a reference to the given []string and assigns it to the Restriction field.
func (o *Domain) SetRestriction(v []string) {
	o.Restriction = v
}

// GetTransfer returns the Transfer field value if set, zero value otherwise.
func (o *Domain) GetTransfer() DomainTransferData {
	if o == nil || IsNil(o.Transfer) {
		var ret DomainTransferData
		return ret
	}
	return *o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetTransferOk() (*DomainTransferData, bool) {
	if o == nil || IsNil(o.Transfer) {
		return nil, false
	}
	return o.Transfer, true
}

// HasTransfer returns a boolean if a field has been set.
func (o *Domain) HasTransfer() bool {
	if o != nil && !IsNil(o.Transfer) {
		return true
	}

	return false
}

// SetTransfer gets a reference to the given DomainTransferData and assigns it to the Transfer field.
func (o *Domain) SetTransfer(v DomainTransferData) {
	o.Transfer = &v
}

// GetFingerprintStatus returns the FingerprintStatus field value if set, zero value otherwise.
func (o *Domain) GetFingerprintStatus() bool {
	if o == nil || IsNil(o.FingerprintStatus) {
		var ret bool
		return ret
	}
	return *o.FingerprintStatus
}

// GetFingerprintStatusOk returns a tuple with the FingerprintStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetFingerprintStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.FingerprintStatus) {
		return nil, false
	}
	return o.FingerprintStatus, true
}

// HasFingerprintStatus returns a boolean if a field has been set.
func (o *Domain) HasFingerprintStatus() bool {
	if o != nil && !IsNil(o.FingerprintStatus) {
		return true
	}

	return false
}

// SetFingerprintStatus gets a reference to the given bool and assigns it to the FingerprintStatus field.
func (o *Domain) SetFingerprintStatus(v bool) {
	o.FingerprintStatus = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Domain) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Domain) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Domain) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Domain) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Domain) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Domain) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlanLevel) {
		toSerialize["plan_level"] = o.PlanLevel
	}
	if !IsNil(o.NsKeys) {
		toSerialize["ns_keys"] = o.NsKeys
	}
	if !IsNil(o.CurrentNs) {
		toSerialize["current_ns"] = o.CurrentNs
	}
	if o.TargetCname.IsSet() {
		toSerialize["target_cname"] = o.TargetCname.Get()
	}
	if o.CustomCname.IsSet() {
		toSerialize["custom_cname"] = o.CustomCname.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.DnsCloud) {
		toSerialize["dns_cloud"] = o.DnsCloud
	}
	if !IsNil(o.Restriction) {
		toSerialize["restriction"] = o.Restriction
	}
	if !IsNil(o.Transfer) {
		toSerialize["transfer"] = o.Transfer
	}
	// skip: fingerprint_status is readOnly
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	return toSerialize, nil
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


