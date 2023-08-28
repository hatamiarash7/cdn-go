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

// checks if the AttackReportItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportItem{}

// AttackReportItem struct for AttackReportItem
type AttackReportItem struct {
	AttackerIp *string `json:"attacker_ip,omitempty"`
	AttackerCountry *string `json:"attacker_country,omitempty"`
	Method *string `json:"method,omitempty"`
	Uri *string `json:"uri,omitempty"`
	Host []string `json:"host,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	UriArgs *string `json:"uri_args,omitempty"`
	Cookie []string `json:"cookie,omitempty"`
	Alerts []string `json:"alerts,omitempty"`
	UserAgent []string `json:"user_agent,omitempty"`
}

// NewAttackReportItem instantiates a new AttackReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportItem() *AttackReportItem {
	this := AttackReportItem{}
	return &this
}

// NewAttackReportItemWithDefaults instantiates a new AttackReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportItemWithDefaults() *AttackReportItem {
	this := AttackReportItem{}
	return &this
}

// GetAttackerIp returns the AttackerIp field value if set, zero value otherwise.
func (o *AttackReportItem) GetAttackerIp() string {
	if o == nil || IsNil(o.AttackerIp) {
		var ret string
		return ret
	}
	return *o.AttackerIp
}

// GetAttackerIpOk returns a tuple with the AttackerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetAttackerIpOk() (*string, bool) {
	if o == nil || IsNil(o.AttackerIp) {
		return nil, false
	}
	return o.AttackerIp, true
}

// HasAttackerIp returns a boolean if a field has been set.
func (o *AttackReportItem) HasAttackerIp() bool {
	if o != nil && !IsNil(o.AttackerIp) {
		return true
	}

	return false
}

// SetAttackerIp gets a reference to the given string and assigns it to the AttackerIp field.
func (o *AttackReportItem) SetAttackerIp(v string) {
	o.AttackerIp = &v
}

// GetAttackerCountry returns the AttackerCountry field value if set, zero value otherwise.
func (o *AttackReportItem) GetAttackerCountry() string {
	if o == nil || IsNil(o.AttackerCountry) {
		var ret string
		return ret
	}
	return *o.AttackerCountry
}

// GetAttackerCountryOk returns a tuple with the AttackerCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetAttackerCountryOk() (*string, bool) {
	if o == nil || IsNil(o.AttackerCountry) {
		return nil, false
	}
	return o.AttackerCountry, true
}

// HasAttackerCountry returns a boolean if a field has been set.
func (o *AttackReportItem) HasAttackerCountry() bool {
	if o != nil && !IsNil(o.AttackerCountry) {
		return true
	}

	return false
}

// SetAttackerCountry gets a reference to the given string and assigns it to the AttackerCountry field.
func (o *AttackReportItem) SetAttackerCountry(v string) {
	o.AttackerCountry = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *AttackReportItem) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *AttackReportItem) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *AttackReportItem) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *AttackReportItem) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *AttackReportItem) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *AttackReportItem) SetUri(v string) {
	o.Uri = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *AttackReportItem) GetHost() []string {
	if o == nil || IsNil(o.Host) {
		var ret []string
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetHostOk() ([]string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *AttackReportItem) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given []string and assigns it to the Host field.
func (o *AttackReportItem) SetHost(v []string) {
	o.Host = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AttackReportItem) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AttackReportItem) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *AttackReportItem) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetUriArgs returns the UriArgs field value if set, zero value otherwise.
func (o *AttackReportItem) GetUriArgs() string {
	if o == nil || IsNil(o.UriArgs) {
		var ret string
		return ret
	}
	return *o.UriArgs
}

// GetUriArgsOk returns a tuple with the UriArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetUriArgsOk() (*string, bool) {
	if o == nil || IsNil(o.UriArgs) {
		return nil, false
	}
	return o.UriArgs, true
}

// HasUriArgs returns a boolean if a field has been set.
func (o *AttackReportItem) HasUriArgs() bool {
	if o != nil && !IsNil(o.UriArgs) {
		return true
	}

	return false
}

// SetUriArgs gets a reference to the given string and assigns it to the UriArgs field.
func (o *AttackReportItem) SetUriArgs(v string) {
	o.UriArgs = &v
}

// GetCookie returns the Cookie field value if set, zero value otherwise.
func (o *AttackReportItem) GetCookie() []string {
	if o == nil || IsNil(o.Cookie) {
		var ret []string
		return ret
	}
	return o.Cookie
}

// GetCookieOk returns a tuple with the Cookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetCookieOk() ([]string, bool) {
	if o == nil || IsNil(o.Cookie) {
		return nil, false
	}
	return o.Cookie, true
}

// HasCookie returns a boolean if a field has been set.
func (o *AttackReportItem) HasCookie() bool {
	if o != nil && !IsNil(o.Cookie) {
		return true
	}

	return false
}

// SetCookie gets a reference to the given []string and assigns it to the Cookie field.
func (o *AttackReportItem) SetCookie(v []string) {
	o.Cookie = v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *AttackReportItem) GetAlerts() []string {
	if o == nil || IsNil(o.Alerts) {
		var ret []string
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetAlertsOk() ([]string, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *AttackReportItem) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []string and assigns it to the Alerts field.
func (o *AttackReportItem) SetAlerts(v []string) {
	o.Alerts = v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *AttackReportItem) GetUserAgent() []string {
	if o == nil || IsNil(o.UserAgent) {
		var ret []string
		return ret
	}
	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportItem) GetUserAgentOk() ([]string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *AttackReportItem) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given []string and assigns it to the UserAgent field.
func (o *AttackReportItem) SetUserAgent(v []string) {
	o.UserAgent = v
}

func (o AttackReportItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttackerIp) {
		toSerialize["attacker_ip"] = o.AttackerIp
	}
	if !IsNil(o.AttackerCountry) {
		toSerialize["attacker_country"] = o.AttackerCountry
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.UriArgs) {
		toSerialize["uri_args"] = o.UriArgs
	}
	if !IsNil(o.Cookie) {
		toSerialize["cookie"] = o.Cookie
	}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.UserAgent) {
		toSerialize["user_agent"] = o.UserAgent
	}
	return toSerialize, nil
}

type NullableAttackReportItem struct {
	value *AttackReportItem
	isSet bool
}

func (v NullableAttackReportItem) Get() *AttackReportItem {
	return v.value
}

func (v *NullableAttackReportItem) Set(val *AttackReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportItem(val *AttackReportItem) *NullableAttackReportItem {
	return &NullableAttackReportItem{value: val, isSet: true}
}

func (v NullableAttackReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


