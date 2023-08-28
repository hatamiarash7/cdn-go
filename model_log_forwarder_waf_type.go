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

// checks if the LogForwarderWAFType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderWAFType{}

// LogForwarderWAFType Waf log type
type LogForwarderWAFType struct {
	Product *bool `json:"product,omitempty"`
	Timestamp *bool `json:"timestamp,omitempty"`
	RemoteAddress *bool `json:"remote_address,omitempty"`
	Domain *bool `json:"domain,omitempty"`
	Data *bool `json:"data,omitempty"`
}

// NewLogForwarderWAFType instantiates a new LogForwarderWAFType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderWAFType() *LogForwarderWAFType {
	this := LogForwarderWAFType{}
	return &this
}

// NewLogForwarderWAFTypeWithDefaults instantiates a new LogForwarderWAFType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderWAFTypeWithDefaults() *LogForwarderWAFType {
	this := LogForwarderWAFType{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *LogForwarderWAFType) GetProduct() bool {
	if o == nil || IsNil(o.Product) {
		var ret bool
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderWAFType) GetProductOk() (*bool, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *LogForwarderWAFType) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given bool and assigns it to the Product field.
func (o *LogForwarderWAFType) SetProduct(v bool) {
	o.Product = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LogForwarderWAFType) GetTimestamp() bool {
	if o == nil || IsNil(o.Timestamp) {
		var ret bool
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderWAFType) GetTimestampOk() (*bool, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LogForwarderWAFType) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given bool and assigns it to the Timestamp field.
func (o *LogForwarderWAFType) SetTimestamp(v bool) {
	o.Timestamp = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *LogForwarderWAFType) GetRemoteAddress() bool {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret bool
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderWAFType) GetRemoteAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *LogForwarderWAFType) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given bool and assigns it to the RemoteAddress field.
func (o *LogForwarderWAFType) SetRemoteAddress(v bool) {
	o.RemoteAddress = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LogForwarderWAFType) GetDomain() bool {
	if o == nil || IsNil(o.Domain) {
		var ret bool
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderWAFType) GetDomainOk() (*bool, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LogForwarderWAFType) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given bool and assigns it to the Domain field.
func (o *LogForwarderWAFType) SetDomain(v bool) {
	o.Domain = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LogForwarderWAFType) GetData() bool {
	if o == nil || IsNil(o.Data) {
		var ret bool
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderWAFType) GetDataOk() (*bool, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LogForwarderWAFType) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given bool and assigns it to the Data field.
func (o *LogForwarderWAFType) SetData(v bool) {
	o.Data = &v
}

func (o LogForwarderWAFType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderWAFType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.RemoteAddress) {
		toSerialize["remote_address"] = o.RemoteAddress
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLogForwarderWAFType struct {
	value *LogForwarderWAFType
	isSet bool
}

func (v NullableLogForwarderWAFType) Get() *LogForwarderWAFType {
	return v.value
}

func (v *NullableLogForwarderWAFType) Set(val *LogForwarderWAFType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderWAFType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderWAFType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderWAFType(val *LogForwarderWAFType) *NullableLogForwarderWAFType {
	return &NullableLogForwarderWAFType{value: val, isSet: true}
}

func (v NullableLogForwarderWAFType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderWAFType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


