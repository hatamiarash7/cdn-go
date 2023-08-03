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

// checks if the PageRuleSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageRuleSummary{}

// PageRuleSummary struct for PageRuleSummary
type PageRuleSummary struct {
	Id *string `json:"id,omitempty"`
	DomainId *string `json:"domain_id,omitempty"`
	// Order of the page-rule
	Seq *int32 `json:"seq,omitempty"`
	// This flag is deprecated in favor of is_protected flag
	// Deprecated
	UrlType *string `json:"url_type,omitempty"`
	// Protected records cannot be modified or deleted by user.
	IsProtected *bool `json:"is_protected,omitempty"`
	// URL pattern of target pages
	Url *string `json:"url,omitempty"`
	CacheLevel *string `json:"cache_level,omitempty"`
	WafStatus *bool `json:"waf_status,omitempty"`
	// Shows whether firewall is enabled or not
	// Deprecated
	FwStatus *bool `json:"fw_status,omitempty"`
	Acceleration *Acceleration `json:"acceleration,omitempty"`
	// Secure link is enabled or not
	SlinkStatus *bool `json:"slink_status,omitempty"`
	// Is the page-rule enabled?
	Status *bool `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewPageRuleSummary instantiates a new PageRuleSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRuleSummary() *PageRuleSummary {
	this := PageRuleSummary{}
	var urlType string = "default"
	this.UrlType = &urlType
	var cacheLevel string = "query_string"
	this.CacheLevel = &cacheLevel
	var wafStatus bool = true
	this.WafStatus = &wafStatus
	var fwStatus bool = true
	this.FwStatus = &fwStatus
	var slinkStatus bool = false
	this.SlinkStatus = &slinkStatus
	var status bool = true
	this.Status = &status
	return &this
}

// NewPageRuleSummaryWithDefaults instantiates a new PageRuleSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRuleSummaryWithDefaults() *PageRuleSummary {
	this := PageRuleSummary{}
	var urlType string = "default"
	this.UrlType = &urlType
	var cacheLevel string = "query_string"
	this.CacheLevel = &cacheLevel
	var wafStatus bool = true
	this.WafStatus = &wafStatus
	var fwStatus bool = true
	this.FwStatus = &fwStatus
	var slinkStatus bool = false
	this.SlinkStatus = &slinkStatus
	var status bool = true
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PageRuleSummary) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PageRuleSummary) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PageRuleSummary) SetId(v string) {
	o.Id = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *PageRuleSummary) GetDomainId() string {
	if o == nil || IsNil(o.DomainId) {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetDomainIdOk() (*string, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *PageRuleSummary) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *PageRuleSummary) SetDomainId(v string) {
	o.DomainId = &v
}

// GetSeq returns the Seq field value if set, zero value otherwise.
func (o *PageRuleSummary) GetSeq() int32 {
	if o == nil || IsNil(o.Seq) {
		var ret int32
		return ret
	}
	return *o.Seq
}

// GetSeqOk returns a tuple with the Seq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetSeqOk() (*int32, bool) {
	if o == nil || IsNil(o.Seq) {
		return nil, false
	}
	return o.Seq, true
}

// HasSeq returns a boolean if a field has been set.
func (o *PageRuleSummary) HasSeq() bool {
	if o != nil && !IsNil(o.Seq) {
		return true
	}

	return false
}

// SetSeq gets a reference to the given int32 and assigns it to the Seq field.
func (o *PageRuleSummary) SetSeq(v int32) {
	o.Seq = &v
}

// GetUrlType returns the UrlType field value if set, zero value otherwise.
// Deprecated
func (o *PageRuleSummary) GetUrlType() string {
	if o == nil || IsNil(o.UrlType) {
		var ret string
		return ret
	}
	return *o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PageRuleSummary) GetUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UrlType) {
		return nil, false
	}
	return o.UrlType, true
}

// HasUrlType returns a boolean if a field has been set.
func (o *PageRuleSummary) HasUrlType() bool {
	if o != nil && !IsNil(o.UrlType) {
		return true
	}

	return false
}

// SetUrlType gets a reference to the given string and assigns it to the UrlType field.
// Deprecated
func (o *PageRuleSummary) SetUrlType(v string) {
	o.UrlType = &v
}

// GetIsProtected returns the IsProtected field value if set, zero value otherwise.
func (o *PageRuleSummary) GetIsProtected() bool {
	if o == nil || IsNil(o.IsProtected) {
		var ret bool
		return ret
	}
	return *o.IsProtected
}

// GetIsProtectedOk returns a tuple with the IsProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetIsProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProtected) {
		return nil, false
	}
	return o.IsProtected, true
}

// HasIsProtected returns a boolean if a field has been set.
func (o *PageRuleSummary) HasIsProtected() bool {
	if o != nil && !IsNil(o.IsProtected) {
		return true
	}

	return false
}

// SetIsProtected gets a reference to the given bool and assigns it to the IsProtected field.
func (o *PageRuleSummary) SetIsProtected(v bool) {
	o.IsProtected = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PageRuleSummary) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PageRuleSummary) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PageRuleSummary) SetUrl(v string) {
	o.Url = &v
}

// GetCacheLevel returns the CacheLevel field value if set, zero value otherwise.
func (o *PageRuleSummary) GetCacheLevel() string {
	if o == nil || IsNil(o.CacheLevel) {
		var ret string
		return ret
	}
	return *o.CacheLevel
}

// GetCacheLevelOk returns a tuple with the CacheLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetCacheLevelOk() (*string, bool) {
	if o == nil || IsNil(o.CacheLevel) {
		return nil, false
	}
	return o.CacheLevel, true
}

// HasCacheLevel returns a boolean if a field has been set.
func (o *PageRuleSummary) HasCacheLevel() bool {
	if o != nil && !IsNil(o.CacheLevel) {
		return true
	}

	return false
}

// SetCacheLevel gets a reference to the given string and assigns it to the CacheLevel field.
func (o *PageRuleSummary) SetCacheLevel(v string) {
	o.CacheLevel = &v
}

// GetWafStatus returns the WafStatus field value if set, zero value otherwise.
func (o *PageRuleSummary) GetWafStatus() bool {
	if o == nil || IsNil(o.WafStatus) {
		var ret bool
		return ret
	}
	return *o.WafStatus
}

// GetWafStatusOk returns a tuple with the WafStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetWafStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.WafStatus) {
		return nil, false
	}
	return o.WafStatus, true
}

// HasWafStatus returns a boolean if a field has been set.
func (o *PageRuleSummary) HasWafStatus() bool {
	if o != nil && !IsNil(o.WafStatus) {
		return true
	}

	return false
}

// SetWafStatus gets a reference to the given bool and assigns it to the WafStatus field.
func (o *PageRuleSummary) SetWafStatus(v bool) {
	o.WafStatus = &v
}

// GetFwStatus returns the FwStatus field value if set, zero value otherwise.
// Deprecated
func (o *PageRuleSummary) GetFwStatus() bool {
	if o == nil || IsNil(o.FwStatus) {
		var ret bool
		return ret
	}
	return *o.FwStatus
}

// GetFwStatusOk returns a tuple with the FwStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PageRuleSummary) GetFwStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.FwStatus) {
		return nil, false
	}
	return o.FwStatus, true
}

// HasFwStatus returns a boolean if a field has been set.
func (o *PageRuleSummary) HasFwStatus() bool {
	if o != nil && !IsNil(o.FwStatus) {
		return true
	}

	return false
}

// SetFwStatus gets a reference to the given bool and assigns it to the FwStatus field.
// Deprecated
func (o *PageRuleSummary) SetFwStatus(v bool) {
	o.FwStatus = &v
}

// GetAcceleration returns the Acceleration field value if set, zero value otherwise.
func (o *PageRuleSummary) GetAcceleration() Acceleration {
	if o == nil || IsNil(o.Acceleration) {
		var ret Acceleration
		return ret
	}
	return *o.Acceleration
}

// GetAccelerationOk returns a tuple with the Acceleration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetAccelerationOk() (*Acceleration, bool) {
	if o == nil || IsNil(o.Acceleration) {
		return nil, false
	}
	return o.Acceleration, true
}

// HasAcceleration returns a boolean if a field has been set.
func (o *PageRuleSummary) HasAcceleration() bool {
	if o != nil && !IsNil(o.Acceleration) {
		return true
	}

	return false
}

// SetAcceleration gets a reference to the given Acceleration and assigns it to the Acceleration field.
func (o *PageRuleSummary) SetAcceleration(v Acceleration) {
	o.Acceleration = &v
}

// GetSlinkStatus returns the SlinkStatus field value if set, zero value otherwise.
func (o *PageRuleSummary) GetSlinkStatus() bool {
	if o == nil || IsNil(o.SlinkStatus) {
		var ret bool
		return ret
	}
	return *o.SlinkStatus
}

// GetSlinkStatusOk returns a tuple with the SlinkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetSlinkStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.SlinkStatus) {
		return nil, false
	}
	return o.SlinkStatus, true
}

// HasSlinkStatus returns a boolean if a field has been set.
func (o *PageRuleSummary) HasSlinkStatus() bool {
	if o != nil && !IsNil(o.SlinkStatus) {
		return true
	}

	return false
}

// SetSlinkStatus gets a reference to the given bool and assigns it to the SlinkStatus field.
func (o *PageRuleSummary) SetSlinkStatus(v bool) {
	o.SlinkStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PageRuleSummary) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PageRuleSummary) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *PageRuleSummary) SetStatus(v bool) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PageRuleSummary) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PageRuleSummary) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PageRuleSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PageRuleSummary) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleSummary) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PageRuleSummary) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PageRuleSummary) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PageRuleSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageRuleSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.DomainId) {
		toSerialize["domain_id"] = o.DomainId
	}
	if !IsNil(o.Seq) {
		toSerialize["seq"] = o.Seq
	}
	if !IsNil(o.UrlType) {
		toSerialize["url_type"] = o.UrlType
	}
	// skip: is_protected is readOnly
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.CacheLevel) {
		toSerialize["cache_level"] = o.CacheLevel
	}
	if !IsNil(o.WafStatus) {
		toSerialize["waf_status"] = o.WafStatus
	}
	if !IsNil(o.FwStatus) {
		toSerialize["fw_status"] = o.FwStatus
	}
	if !IsNil(o.Acceleration) {
		toSerialize["acceleration"] = o.Acceleration
	}
	if !IsNil(o.SlinkStatus) {
		toSerialize["slink_status"] = o.SlinkStatus
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	return toSerialize, nil
}

type NullablePageRuleSummary struct {
	value *PageRuleSummary
	isSet bool
}

func (v NullablePageRuleSummary) Get() *PageRuleSummary {
	return v.value
}

func (v *NullablePageRuleSummary) Set(val *PageRuleSummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRuleSummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRuleSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRuleSummary(val *PageRuleSummary) *NullablePageRuleSummary {
	return &NullablePageRuleSummary{value: val, isSet: true}
}

func (v NullablePageRuleSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRuleSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


