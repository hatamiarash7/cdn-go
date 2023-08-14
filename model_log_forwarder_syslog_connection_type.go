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

// checks if the LogForwarderSyslogConnectionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderSyslogConnectionType{}

// LogForwarderSyslogConnectionType Syslog connection
type LogForwarderSyslogConnectionType struct {
	SampleRate *int32 `json:"sample_rate,omitempty"`
	Logtype *string `json:"logtype,omitempty"`
	Host *string `json:"host,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Tls *bool `json:"tls,omitempty"`
	Cert *string `json:"cert,omitempty"`
	RetryTime *int32 `json:"retry_time,omitempty"`
}

// NewLogForwarderSyslogConnectionType instantiates a new LogForwarderSyslogConnectionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderSyslogConnectionType() *LogForwarderSyslogConnectionType {
	this := LogForwarderSyslogConnectionType{}
	return &this
}

// NewLogForwarderSyslogConnectionTypeWithDefaults instantiates a new LogForwarderSyslogConnectionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderSyslogConnectionTypeWithDefaults() *LogForwarderSyslogConnectionType {
	this := LogForwarderSyslogConnectionType{}
	return &this
}

// GetSampleRate returns the SampleRate field value if set, zero value otherwise.
func (o *LogForwarderSyslogConnectionType) GetSampleRate() int32 {
	if o == nil || IsNil(o.SampleRate) {
		var ret int32
		return ret
	}
	return *o.SampleRate
}

// GetSampleRateOk returns a tuple with the SampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSyslogConnectionType) GetSampleRateOk() (*int32, bool) {
	if o == nil || IsNil(o.SampleRate) {
		return nil, false
	}
	return o.SampleRate, true
}

// HasSampleRate returns a boolean if a field has been set.
func (o *LogForwarderSyslogConnectionType) HasSampleRate() bool {
	if o != nil && !IsNil(o.SampleRate) {
		return true
	}

	return false
}

// SetSampleRate gets a reference to the given int32 and assigns it to the SampleRate field.
func (o *LogForwarderSyslogConnectionType) SetSampleRate(v int32) {
	o.SampleRate = &v
}

// GetLogtype returns the Logtype field value if set, zero value otherwise.
func (o *LogForwarderSyslogConnectionType) GetLogtype() string {
	if o == nil || IsNil(o.Logtype) {
		var ret string
		return ret
	}
	return *o.Logtype
}

// GetLogtypeOk returns a tuple with the Logtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSyslogConnectionType) GetLogtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Logtype) {
		return nil, false
	}
	return o.Logtype, true
}

// HasLogtype returns a boolean if a field has been set.
func (o *LogForwarderSyslogConnectionType) HasLogtype() bool {
	if o != nil && !IsNil(o.Logtype) {
		return true
	}

	return false
}

// SetLogtype gets a reference to the given string and assigns it to the Logtype field.
func (o *LogForwarderSyslogConnectionType) SetLogtype(v string) {
	o.Logtype = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *LogForwarderSyslogConnectionType) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSyslogConnectionType) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *LogForwarderSyslogConnectionType) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *LogForwarderSyslogConnectionType) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LogForwarderSyslogConnectionType) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSyslogConnectionType) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LogForwarderSyslogConnectionType) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LogForwarderSyslogConnectionType) SetPort(v int32) {
	o.Port = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *LogForwarderSyslogConnectionType) GetTls() bool {
	if o == nil || IsNil(o.Tls) {
		var ret bool
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSyslogConnectionType) GetTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *LogForwarderSyslogConnectionType) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given bool and assigns it to the Tls field.
func (o *LogForwarderSyslogConnectionType) SetTls(v bool) {
	o.Tls = &v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *LogForwarderSyslogConnectionType) GetCert() string {
	if o == nil || IsNil(o.Cert) {
		var ret string
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSyslogConnectionType) GetCertOk() (*string, bool) {
	if o == nil || IsNil(o.Cert) {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *LogForwarderSyslogConnectionType) HasCert() bool {
	if o != nil && !IsNil(o.Cert) {
		return true
	}

	return false
}

// SetCert gets a reference to the given string and assigns it to the Cert field.
func (o *LogForwarderSyslogConnectionType) SetCert(v string) {
	o.Cert = &v
}

// GetRetryTime returns the RetryTime field value if set, zero value otherwise.
func (o *LogForwarderSyslogConnectionType) GetRetryTime() int32 {
	if o == nil || IsNil(o.RetryTime) {
		var ret int32
		return ret
	}
	return *o.RetryTime
}

// GetRetryTimeOk returns a tuple with the RetryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSyslogConnectionType) GetRetryTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryTime) {
		return nil, false
	}
	return o.RetryTime, true
}

// HasRetryTime returns a boolean if a field has been set.
func (o *LogForwarderSyslogConnectionType) HasRetryTime() bool {
	if o != nil && !IsNil(o.RetryTime) {
		return true
	}

	return false
}

// SetRetryTime gets a reference to the given int32 and assigns it to the RetryTime field.
func (o *LogForwarderSyslogConnectionType) SetRetryTime(v int32) {
	o.RetryTime = &v
}

func (o LogForwarderSyslogConnectionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderSyslogConnectionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SampleRate) {
		toSerialize["sample_rate"] = o.SampleRate
	}
	if !IsNil(o.Logtype) {
		toSerialize["logtype"] = o.Logtype
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	if !IsNil(o.Cert) {
		toSerialize["cert"] = o.Cert
	}
	if !IsNil(o.RetryTime) {
		toSerialize["retry_time"] = o.RetryTime
	}
	return toSerialize, nil
}

type NullableLogForwarderSyslogConnectionType struct {
	value *LogForwarderSyslogConnectionType
	isSet bool
}

func (v NullableLogForwarderSyslogConnectionType) Get() *LogForwarderSyslogConnectionType {
	return v.value
}

func (v *NullableLogForwarderSyslogConnectionType) Set(val *LogForwarderSyslogConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderSyslogConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderSyslogConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderSyslogConnectionType(val *LogForwarderSyslogConnectionType) *NullableLogForwarderSyslogConnectionType {
	return &NullableLogForwarderSyslogConnectionType{value: val, isSet: true}
}

func (v NullableLogForwarderSyslogConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderSyslogConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


