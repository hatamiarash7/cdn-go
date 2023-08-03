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

// checks if the BaseDnsRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseDnsRecord{}

// BaseDnsRecord struct for BaseDnsRecord
type BaseDnsRecord struct {
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

// NewBaseDnsRecord instantiates a new BaseDnsRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseDnsRecord() *BaseDnsRecord {
	this := BaseDnsRecord{}
	var cloud bool = false
	this.Cloud = &cloud
	return &this
}

// NewBaseDnsRecordWithDefaults instantiates a new BaseDnsRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseDnsRecordWithDefaults() *BaseDnsRecord {
	this := BaseDnsRecord{}
	var cloud bool = false
	this.Cloud = &cloud
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseDnsRecord) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseDnsRecord) SetName(v string) {
	o.Name = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *BaseDnsRecord) SetTtl(v int32) {
	o.Ttl = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetCloud() bool {
	if o == nil || IsNil(o.Cloud) {
		var ret bool
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetCloudOk() (*bool, bool) {
	if o == nil || IsNil(o.Cloud) {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasCloud() bool {
	if o != nil && !IsNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given bool and assigns it to the Cloud field.
func (o *BaseDnsRecord) SetCloud(v bool) {
	o.Cloud = &v
}

// GetUpstreamHttps returns the UpstreamHttps field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetUpstreamHttps() string {
	if o == nil || IsNil(o.UpstreamHttps) {
		var ret string
		return ret
	}
	return *o.UpstreamHttps
}

// GetUpstreamHttpsOk returns a tuple with the UpstreamHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetUpstreamHttpsOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamHttps) {
		return nil, false
	}
	return o.UpstreamHttps, true
}

// HasUpstreamHttps returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasUpstreamHttps() bool {
	if o != nil && !IsNil(o.UpstreamHttps) {
		return true
	}

	return false
}

// SetUpstreamHttps gets a reference to the given string and assigns it to the UpstreamHttps field.
func (o *BaseDnsRecord) SetUpstreamHttps(v string) {
	o.UpstreamHttps = &v
}

// GetIpFilterMode returns the IpFilterMode field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetIpFilterMode() DnsRecordIpFilterMode {
	if o == nil || IsNil(o.IpFilterMode) {
		var ret DnsRecordIpFilterMode
		return ret
	}
	return *o.IpFilterMode
}

// GetIpFilterModeOk returns a tuple with the IpFilterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetIpFilterModeOk() (*DnsRecordIpFilterMode, bool) {
	if o == nil || IsNil(o.IpFilterMode) {
		return nil, false
	}
	return o.IpFilterMode, true
}

// HasIpFilterMode returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasIpFilterMode() bool {
	if o != nil && !IsNil(o.IpFilterMode) {
		return true
	}

	return false
}

// SetIpFilterMode gets a reference to the given DnsRecordIpFilterMode and assigns it to the IpFilterMode field.
func (o *BaseDnsRecord) SetIpFilterMode(v DnsRecordIpFilterMode) {
	o.IpFilterMode = &v
}

// GetIsProtected returns the IsProtected field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetIsProtected() bool {
	if o == nil || IsNil(o.IsProtected) {
		var ret bool
		return ret
	}
	return *o.IsProtected
}

// GetIsProtectedOk returns a tuple with the IsProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetIsProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProtected) {
		return nil, false
	}
	return o.IsProtected, true
}

// HasIsProtected returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasIsProtected() bool {
	if o != nil && !IsNil(o.IsProtected) {
		return true
	}

	return false
}

// SetIsProtected gets a reference to the given bool and assigns it to the IsProtected field.
func (o *BaseDnsRecord) SetIsProtected(v bool) {
	o.IsProtected = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetUsage() []string {
	if o == nil || IsNil(o.Usage) {
		var ret []string
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetUsageOk() ([]string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []string and assigns it to the Usage field.
func (o *BaseDnsRecord) SetUsage(v []string) {
	o.Usage = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BaseDnsRecord) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BaseDnsRecord) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDnsRecord) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BaseDnsRecord) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *BaseDnsRecord) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o BaseDnsRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseDnsRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableBaseDnsRecord struct {
	value *BaseDnsRecord
	isSet bool
}

func (v NullableBaseDnsRecord) Get() *BaseDnsRecord {
	return v.value
}

func (v *NullableBaseDnsRecord) Set(val *BaseDnsRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseDnsRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseDnsRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseDnsRecord(val *BaseDnsRecord) *NullableBaseDnsRecord {
	return &NullableBaseDnsRecord{value: val, isSet: true}
}

func (v NullableBaseDnsRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseDnsRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


