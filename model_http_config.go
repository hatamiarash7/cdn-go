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

// checks if the HttpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpConfig{}

// HttpConfig struct for HttpConfig
type HttpConfig struct {
	Method string `json:"method"`
	Port int32 `json:"port"`
	// path for test
	Path string `json:"path"`
	AllowInsecure bool `json:"allow_insecure"`
	ExpectedResponse ExpectedResponse `json:"expected_response"`
	Headers map[string]string `json:"headers"`
	FollowRedirects *bool `json:"follow_redirects,omitempty"`
	// In milliseconds
	Timeout int32 `json:"timeout"`
}

// NewHttpConfig instantiates a new HttpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpConfig(method string, port int32, path string, allowInsecure bool, expectedResponse ExpectedResponse, headers map[string]string, timeout int32) *HttpConfig {
	this := HttpConfig{}
	this.Method = method
	this.Port = port
	this.Path = path
	this.AllowInsecure = allowInsecure
	this.ExpectedResponse = expectedResponse
	this.Headers = headers
	this.Timeout = timeout
	return &this
}

// NewHttpConfigWithDefaults instantiates a new HttpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpConfigWithDefaults() *HttpConfig {
	this := HttpConfig{}
	return &this
}

// GetMethod returns the Method field value
func (o *HttpConfig) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *HttpConfig) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *HttpConfig) SetMethod(v string) {
	o.Method = v
}

// GetPort returns the Port field value
func (o *HttpConfig) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *HttpConfig) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *HttpConfig) SetPort(v int32) {
	o.Port = v
}

// GetPath returns the Path field value
func (o *HttpConfig) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *HttpConfig) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *HttpConfig) SetPath(v string) {
	o.Path = v
}

// GetAllowInsecure returns the AllowInsecure field value
func (o *HttpConfig) GetAllowInsecure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowInsecure
}

// GetAllowInsecureOk returns a tuple with the AllowInsecure field value
// and a boolean to check if the value has been set.
func (o *HttpConfig) GetAllowInsecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowInsecure, true
}

// SetAllowInsecure sets field value
func (o *HttpConfig) SetAllowInsecure(v bool) {
	o.AllowInsecure = v
}

// GetExpectedResponse returns the ExpectedResponse field value
func (o *HttpConfig) GetExpectedResponse() ExpectedResponse {
	if o == nil {
		var ret ExpectedResponse
		return ret
	}

	return o.ExpectedResponse
}

// GetExpectedResponseOk returns a tuple with the ExpectedResponse field value
// and a boolean to check if the value has been set.
func (o *HttpConfig) GetExpectedResponseOk() (*ExpectedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpectedResponse, true
}

// SetExpectedResponse sets field value
func (o *HttpConfig) SetExpectedResponse(v ExpectedResponse) {
	o.ExpectedResponse = v
}

// GetHeaders returns the Headers field value
func (o *HttpConfig) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *HttpConfig) GetHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *HttpConfig) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetFollowRedirects returns the FollowRedirects field value if set, zero value otherwise.
func (o *HttpConfig) GetFollowRedirects() bool {
	if o == nil || IsNil(o.FollowRedirects) {
		var ret bool
		return ret
	}
	return *o.FollowRedirects
}

// GetFollowRedirectsOk returns a tuple with the FollowRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpConfig) GetFollowRedirectsOk() (*bool, bool) {
	if o == nil || IsNil(o.FollowRedirects) {
		return nil, false
	}
	return o.FollowRedirects, true
}

// HasFollowRedirects returns a boolean if a field has been set.
func (o *HttpConfig) HasFollowRedirects() bool {
	if o != nil && !IsNil(o.FollowRedirects) {
		return true
	}

	return false
}

// SetFollowRedirects gets a reference to the given bool and assigns it to the FollowRedirects field.
func (o *HttpConfig) SetFollowRedirects(v bool) {
	o.FollowRedirects = &v
}

// GetTimeout returns the Timeout field value
func (o *HttpConfig) GetTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *HttpConfig) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *HttpConfig) SetTimeout(v int32) {
	o.Timeout = v
}

func (o HttpConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["method"] = o.Method
	toSerialize["port"] = o.Port
	toSerialize["path"] = o.Path
	toSerialize["allow_insecure"] = o.AllowInsecure
	toSerialize["expected_response"] = o.ExpectedResponse
	toSerialize["headers"] = o.Headers
	// skip: follow_redirects is readOnly
	toSerialize["timeout"] = o.Timeout
	return toSerialize, nil
}

type NullableHttpConfig struct {
	value *HttpConfig
	isSet bool
}

func (v NullableHttpConfig) Get() *HttpConfig {
	return v.value
}

func (v *NullableHttpConfig) Set(val *HttpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpConfig(val *HttpConfig) *NullableHttpConfig {
	return &NullableHttpConfig{value: val, isSet: true}
}

func (v NullableHttpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


