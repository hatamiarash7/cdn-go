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

// checks if the LogForwarderAccessLogType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderAccessLogType{}

// LogForwarderAccessLogType Access log type
type LogForwarderAccessLogType struct {
	Method *bool `json:"method,omitempty"`
	Scheme *bool `json:"scheme,omitempty"`
	Domain *bool `json:"domain,omitempty"`
	Uri *bool `json:"uri,omitempty"`
	QueryString *bool `json:"query_string,omitempty"`
	Referer *bool `json:"referer,omitempty"`
	Ip *bool `json:"ip,omitempty"`
	Ua *bool `json:"ua,omitempty"`
	Country *bool `json:"country,omitempty"`
	Asn *bool `json:"asn,omitempty"`
	ContentType *bool `json:"content_type,omitempty"`
	Status *bool `json:"status,omitempty"`
	TlsFingerprint *bool `json:"tls_fingerprint,omitempty"`
	ServerPort *bool `json:"server_port,omitempty"`
	BytesSent *bool `json:"bytes_sent,omitempty"`
	BytesReceived *bool `json:"bytes_received,omitempty"`
	UpstreamTime *bool `json:"upstream_time,omitempty"`
	Cache *bool `json:"cache,omitempty"`
	RequestId *bool `json:"request_id,omitempty"`
	Timestamp *bool `json:"timestamp,omitempty"`
}

// NewLogForwarderAccessLogType instantiates a new LogForwarderAccessLogType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderAccessLogType() *LogForwarderAccessLogType {
	this := LogForwarderAccessLogType{}
	return &this
}

// NewLogForwarderAccessLogTypeWithDefaults instantiates a new LogForwarderAccessLogType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderAccessLogTypeWithDefaults() *LogForwarderAccessLogType {
	this := LogForwarderAccessLogType{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetMethod() bool {
	if o == nil || IsNil(o.Method) {
		var ret bool
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetMethodOk() (*bool, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given bool and assigns it to the Method field.
func (o *LogForwarderAccessLogType) SetMethod(v bool) {
	o.Method = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetScheme() bool {
	if o == nil || IsNil(o.Scheme) {
		var ret bool
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetSchemeOk() (*bool, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given bool and assigns it to the Scheme field.
func (o *LogForwarderAccessLogType) SetScheme(v bool) {
	o.Scheme = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetDomain() bool {
	if o == nil || IsNil(o.Domain) {
		var ret bool
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetDomainOk() (*bool, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given bool and assigns it to the Domain field.
func (o *LogForwarderAccessLogType) SetDomain(v bool) {
	o.Domain = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetUri() bool {
	if o == nil || IsNil(o.Uri) {
		var ret bool
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetUriOk() (*bool, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given bool and assigns it to the Uri field.
func (o *LogForwarderAccessLogType) SetUri(v bool) {
	o.Uri = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetQueryString() bool {
	if o == nil || IsNil(o.QueryString) {
		var ret bool
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetQueryStringOk() (*bool, bool) {
	if o == nil || IsNil(o.QueryString) {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasQueryString() bool {
	if o != nil && !IsNil(o.QueryString) {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given bool and assigns it to the QueryString field.
func (o *LogForwarderAccessLogType) SetQueryString(v bool) {
	o.QueryString = &v
}

// GetReferer returns the Referer field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetReferer() bool {
	if o == nil || IsNil(o.Referer) {
		var ret bool
		return ret
	}
	return *o.Referer
}

// GetRefererOk returns a tuple with the Referer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetRefererOk() (*bool, bool) {
	if o == nil || IsNil(o.Referer) {
		return nil, false
	}
	return o.Referer, true
}

// HasReferer returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasReferer() bool {
	if o != nil && !IsNil(o.Referer) {
		return true
	}

	return false
}

// SetReferer gets a reference to the given bool and assigns it to the Referer field.
func (o *LogForwarderAccessLogType) SetReferer(v bool) {
	o.Referer = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetIp() bool {
	if o == nil || IsNil(o.Ip) {
		var ret bool
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetIpOk() (*bool, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given bool and assigns it to the Ip field.
func (o *LogForwarderAccessLogType) SetIp(v bool) {
	o.Ip = &v
}

// GetUa returns the Ua field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetUa() bool {
	if o == nil || IsNil(o.Ua) {
		var ret bool
		return ret
	}
	return *o.Ua
}

// GetUaOk returns a tuple with the Ua field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetUaOk() (*bool, bool) {
	if o == nil || IsNil(o.Ua) {
		return nil, false
	}
	return o.Ua, true
}

// HasUa returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasUa() bool {
	if o != nil && !IsNil(o.Ua) {
		return true
	}

	return false
}

// SetUa gets a reference to the given bool and assigns it to the Ua field.
func (o *LogForwarderAccessLogType) SetUa(v bool) {
	o.Ua = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetCountry() bool {
	if o == nil || IsNil(o.Country) {
		var ret bool
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetCountryOk() (*bool, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given bool and assigns it to the Country field.
func (o *LogForwarderAccessLogType) SetCountry(v bool) {
	o.Country = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetAsn() bool {
	if o == nil || IsNil(o.Asn) {
		var ret bool
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetAsnOk() (*bool, bool) {
	if o == nil || IsNil(o.Asn) {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasAsn() bool {
	if o != nil && !IsNil(o.Asn) {
		return true
	}

	return false
}

// SetAsn gets a reference to the given bool and assigns it to the Asn field.
func (o *LogForwarderAccessLogType) SetAsn(v bool) {
	o.Asn = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetContentType() bool {
	if o == nil || IsNil(o.ContentType) {
		var ret bool
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetContentTypeOk() (*bool, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given bool and assigns it to the ContentType field.
func (o *LogForwarderAccessLogType) SetContentType(v bool) {
	o.ContentType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *LogForwarderAccessLogType) SetStatus(v bool) {
	o.Status = &v
}

// GetTlsFingerprint returns the TlsFingerprint field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetTlsFingerprint() bool {
	if o == nil || IsNil(o.TlsFingerprint) {
		var ret bool
		return ret
	}
	return *o.TlsFingerprint
}

// GetTlsFingerprintOk returns a tuple with the TlsFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetTlsFingerprintOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsFingerprint) {
		return nil, false
	}
	return o.TlsFingerprint, true
}

// HasTlsFingerprint returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasTlsFingerprint() bool {
	if o != nil && !IsNil(o.TlsFingerprint) {
		return true
	}

	return false
}

// SetTlsFingerprint gets a reference to the given bool and assigns it to the TlsFingerprint field.
func (o *LogForwarderAccessLogType) SetTlsFingerprint(v bool) {
	o.TlsFingerprint = &v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetServerPort() bool {
	if o == nil || IsNil(o.ServerPort) {
		var ret bool
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetServerPortOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerPort) {
		return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasServerPort() bool {
	if o != nil && !IsNil(o.ServerPort) {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given bool and assigns it to the ServerPort field.
func (o *LogForwarderAccessLogType) SetServerPort(v bool) {
	o.ServerPort = &v
}

// GetBytesSent returns the BytesSent field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetBytesSent() bool {
	if o == nil || IsNil(o.BytesSent) {
		var ret bool
		return ret
	}
	return *o.BytesSent
}

// GetBytesSentOk returns a tuple with the BytesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetBytesSentOk() (*bool, bool) {
	if o == nil || IsNil(o.BytesSent) {
		return nil, false
	}
	return o.BytesSent, true
}

// HasBytesSent returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasBytesSent() bool {
	if o != nil && !IsNil(o.BytesSent) {
		return true
	}

	return false
}

// SetBytesSent gets a reference to the given bool and assigns it to the BytesSent field.
func (o *LogForwarderAccessLogType) SetBytesSent(v bool) {
	o.BytesSent = &v
}

// GetBytesReceived returns the BytesReceived field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetBytesReceived() bool {
	if o == nil || IsNil(o.BytesReceived) {
		var ret bool
		return ret
	}
	return *o.BytesReceived
}

// GetBytesReceivedOk returns a tuple with the BytesReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetBytesReceivedOk() (*bool, bool) {
	if o == nil || IsNil(o.BytesReceived) {
		return nil, false
	}
	return o.BytesReceived, true
}

// HasBytesReceived returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasBytesReceived() bool {
	if o != nil && !IsNil(o.BytesReceived) {
		return true
	}

	return false
}

// SetBytesReceived gets a reference to the given bool and assigns it to the BytesReceived field.
func (o *LogForwarderAccessLogType) SetBytesReceived(v bool) {
	o.BytesReceived = &v
}

// GetUpstreamTime returns the UpstreamTime field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetUpstreamTime() bool {
	if o == nil || IsNil(o.UpstreamTime) {
		var ret bool
		return ret
	}
	return *o.UpstreamTime
}

// GetUpstreamTimeOk returns a tuple with the UpstreamTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetUpstreamTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.UpstreamTime) {
		return nil, false
	}
	return o.UpstreamTime, true
}

// HasUpstreamTime returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasUpstreamTime() bool {
	if o != nil && !IsNil(o.UpstreamTime) {
		return true
	}

	return false
}

// SetUpstreamTime gets a reference to the given bool and assigns it to the UpstreamTime field.
func (o *LogForwarderAccessLogType) SetUpstreamTime(v bool) {
	o.UpstreamTime = &v
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetCache() bool {
	if o == nil || IsNil(o.Cache) {
		var ret bool
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.Cache) {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasCache() bool {
	if o != nil && !IsNil(o.Cache) {
		return true
	}

	return false
}

// SetCache gets a reference to the given bool and assigns it to the Cache field.
func (o *LogForwarderAccessLogType) SetCache(v bool) {
	o.Cache = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetRequestId() bool {
	if o == nil || IsNil(o.RequestId) {
		var ret bool
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetRequestIdOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given bool and assigns it to the RequestId field.
func (o *LogForwarderAccessLogType) SetRequestId(v bool) {
	o.RequestId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LogForwarderAccessLogType) GetTimestamp() bool {
	if o == nil || IsNil(o.Timestamp) {
		var ret bool
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderAccessLogType) GetTimestampOk() (*bool, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LogForwarderAccessLogType) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given bool and assigns it to the Timestamp field.
func (o *LogForwarderAccessLogType) SetTimestamp(v bool) {
	o.Timestamp = &v
}

func (o LogForwarderAccessLogType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderAccessLogType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.QueryString) {
		toSerialize["query_string"] = o.QueryString
	}
	if !IsNil(o.Referer) {
		toSerialize["referer"] = o.Referer
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Ua) {
		toSerialize["ua"] = o.Ua
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Asn) {
		toSerialize["asn"] = o.Asn
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TlsFingerprint) {
		toSerialize["tls_fingerprint"] = o.TlsFingerprint
	}
	if !IsNil(o.ServerPort) {
		toSerialize["server_port"] = o.ServerPort
	}
	if !IsNil(o.BytesSent) {
		toSerialize["bytes_sent"] = o.BytesSent
	}
	if !IsNil(o.BytesReceived) {
		toSerialize["bytes_received"] = o.BytesReceived
	}
	if !IsNil(o.UpstreamTime) {
		toSerialize["upstream_time"] = o.UpstreamTime
	}
	if !IsNil(o.Cache) {
		toSerialize["cache"] = o.Cache
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableLogForwarderAccessLogType struct {
	value *LogForwarderAccessLogType
	isSet bool
}

func (v NullableLogForwarderAccessLogType) Get() *LogForwarderAccessLogType {
	return v.value
}

func (v *NullableLogForwarderAccessLogType) Set(val *LogForwarderAccessLogType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderAccessLogType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderAccessLogType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderAccessLogType(val *LogForwarderAccessLogType) *NullableLogForwarderAccessLogType {
	return &NullableLogForwarderAccessLogType{value: val, isSet: true}
}

func (v NullableLogForwarderAccessLogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderAccessLogType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


