/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.108.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"time"
)

// checks if the CNAMERecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CNAMERecord{}

// CNAMERecord struct for CNAMERecord
type CNAMERecord struct {
	Value *CNAMERecordValue `json:"value,omitempty"`
	Type *string `json:"type,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
	Cloud *bool `json:"cloud,omitempty"`
	UpstreamHttps *string `json:"upstream_https,omitempty"`
	IpFilterMode *DnsRecordIpFilterMode `json:"ip_filter_mode,omitempty"`
	// Protected records cannot be modified or deleted by user.
	IsProtected *bool `json:"is_protected,omitempty"`
	Usage []string `json:"usage,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewCNAMERecord instantiates a new CNAMERecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCNAMERecord() *CNAMERecord {
	this := CNAMERecord{}
	var cloud bool = false
	this.Cloud = &cloud
	return &this
}

// NewCNAMERecordWithDefaults instantiates a new CNAMERecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCNAMERecordWithDefaults() *CNAMERecord {
	this := CNAMERecord{}
	var type_ string = "cname"
	this.Type = &type_
	var cloud bool = false
	this.Cloud = &cloud
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CNAMERecord) GetValue() CNAMERecordValue {
	if o == nil || IsNil(o.Value) {
		var ret CNAMERecordValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetValueOk() (*CNAMERecordValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CNAMERecord) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CNAMERecordValue and assigns it to the Value field.
func (o *CNAMERecord) SetValue(v CNAMERecordValue) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CNAMERecord) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CNAMERecord) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CNAMERecord) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CNAMERecord) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CNAMERecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CNAMERecord) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CNAMERecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CNAMERecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CNAMERecord) SetName(v string) {
	o.Name = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CNAMERecord) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CNAMERecord) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *CNAMERecord) SetTtl(v int32) {
	o.Ttl = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *CNAMERecord) GetCloud() bool {
	if o == nil || IsNil(o.Cloud) {
		var ret bool
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetCloudOk() (*bool, bool) {
	if o == nil || IsNil(o.Cloud) {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *CNAMERecord) HasCloud() bool {
	if o != nil && !IsNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given bool and assigns it to the Cloud field.
func (o *CNAMERecord) SetCloud(v bool) {
	o.Cloud = &v
}

// GetUpstreamHttps returns the UpstreamHttps field value if set, zero value otherwise.
func (o *CNAMERecord) GetUpstreamHttps() string {
	if o == nil || IsNil(o.UpstreamHttps) {
		var ret string
		return ret
	}
	return *o.UpstreamHttps
}

// GetUpstreamHttpsOk returns a tuple with the UpstreamHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetUpstreamHttpsOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamHttps) {
		return nil, false
	}
	return o.UpstreamHttps, true
}

// HasUpstreamHttps returns a boolean if a field has been set.
func (o *CNAMERecord) HasUpstreamHttps() bool {
	if o != nil && !IsNil(o.UpstreamHttps) {
		return true
	}

	return false
}

// SetUpstreamHttps gets a reference to the given string and assigns it to the UpstreamHttps field.
func (o *CNAMERecord) SetUpstreamHttps(v string) {
	o.UpstreamHttps = &v
}

// GetIpFilterMode returns the IpFilterMode field value if set, zero value otherwise.
func (o *CNAMERecord) GetIpFilterMode() DnsRecordIpFilterMode {
	if o == nil || IsNil(o.IpFilterMode) {
		var ret DnsRecordIpFilterMode
		return ret
	}
	return *o.IpFilterMode
}

// GetIpFilterModeOk returns a tuple with the IpFilterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetIpFilterModeOk() (*DnsRecordIpFilterMode, bool) {
	if o == nil || IsNil(o.IpFilterMode) {
		return nil, false
	}
	return o.IpFilterMode, true
}

// HasIpFilterMode returns a boolean if a field has been set.
func (o *CNAMERecord) HasIpFilterMode() bool {
	if o != nil && !IsNil(o.IpFilterMode) {
		return true
	}

	return false
}

// SetIpFilterMode gets a reference to the given DnsRecordIpFilterMode and assigns it to the IpFilterMode field.
func (o *CNAMERecord) SetIpFilterMode(v DnsRecordIpFilterMode) {
	o.IpFilterMode = &v
}

// GetIsProtected returns the IsProtected field value if set, zero value otherwise.
func (o *CNAMERecord) GetIsProtected() bool {
	if o == nil || IsNil(o.IsProtected) {
		var ret bool
		return ret
	}
	return *o.IsProtected
}

// GetIsProtectedOk returns a tuple with the IsProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetIsProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProtected) {
		return nil, false
	}
	return o.IsProtected, true
}

// HasIsProtected returns a boolean if a field has been set.
func (o *CNAMERecord) HasIsProtected() bool {
	if o != nil && !IsNil(o.IsProtected) {
		return true
	}

	return false
}

// SetIsProtected gets a reference to the given bool and assigns it to the IsProtected field.
func (o *CNAMERecord) SetIsProtected(v bool) {
	o.IsProtected = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *CNAMERecord) GetUsage() []string {
	if o == nil || IsNil(o.Usage) {
		var ret []string
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetUsageOk() ([]string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *CNAMERecord) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []string and assigns it to the Usage field.
func (o *CNAMERecord) SetUsage(v []string) {
	o.Usage = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CNAMERecord) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CNAMERecord) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CNAMERecord) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CNAMERecord) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CNAMERecord) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CNAMERecord) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CNAMERecord) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CNAMERecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CNAMERecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Cloud) {
		toSerialize["cloud"] = o.Cloud
	}
	if !IsNil(o.UpstreamHttps) {
		toSerialize["upstream_https"] = o.UpstreamHttps
	}
	if !IsNil(o.IpFilterMode) {
		toSerialize["ip_filter_mode"] = o.IpFilterMode
	}
	// skip: is_protected is readOnly
	// skip: usage is readOnly
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	return toSerialize, nil
}

type NullableCNAMERecord struct {
	value *CNAMERecord
	isSet bool
}

func (v NullableCNAMERecord) Get() *CNAMERecord {
	return v.value
}

func (v *NullableCNAMERecord) Set(val *CNAMERecord) {
	v.value = val
	v.isSet = true
}

func (v NullableCNAMERecord) IsSet() bool {
	return v.isSet
}

func (v *NullableCNAMERecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCNAMERecord(val *CNAMERecord) *NullableCNAMERecord {
	return &NullableCNAMERecord{value: val, isSet: true}
}

func (v NullableCNAMERecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCNAMERecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


