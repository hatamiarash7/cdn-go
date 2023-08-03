/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.103.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the LogForwarder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarder{}

// LogForwarder struct for LogForwarder
type LogForwarder struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	ConnectionType string `json:"connection_type"`
	DataFormat LogForwarderDataFormat `json:"data_format"`
	Settings LogForwarderSetting `json:"settings"`
	Status bool `json:"status"`
}

// NewLogForwarder instantiates a new LogForwarder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarder(name string, description string, type_ string, connectionType string, dataFormat LogForwarderDataFormat, settings LogForwarderSetting, status bool) *LogForwarder {
	this := LogForwarder{}
	this.Name = name
	this.Description = description
	this.Type = type_
	this.ConnectionType = connectionType
	this.DataFormat = dataFormat
	this.Settings = settings
	this.Status = status
	return &this
}

// NewLogForwarderWithDefaults instantiates a new LogForwarder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderWithDefaults() *LogForwarder {
	this := LogForwarder{}
	return &this
}

// GetName returns the Name field value
func (o *LogForwarder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogForwarder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LogForwarder) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *LogForwarder) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LogForwarder) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LogForwarder) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *LogForwarder) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogForwarder) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogForwarder) SetType(v string) {
	o.Type = v
}

// GetConnectionType returns the ConnectionType field value
func (o *LogForwarder) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *LogForwarder) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *LogForwarder) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetDataFormat returns the DataFormat field value
func (o *LogForwarder) GetDataFormat() LogForwarderDataFormat {
	if o == nil {
		var ret LogForwarderDataFormat
		return ret
	}

	return o.DataFormat
}

// GetDataFormatOk returns a tuple with the DataFormat field value
// and a boolean to check if the value has been set.
func (o *LogForwarder) GetDataFormatOk() (*LogForwarderDataFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataFormat, true
}

// SetDataFormat sets field value
func (o *LogForwarder) SetDataFormat(v LogForwarderDataFormat) {
	o.DataFormat = v
}

// GetSettings returns the Settings field value
func (o *LogForwarder) GetSettings() LogForwarderSetting {
	if o == nil {
		var ret LogForwarderSetting
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *LogForwarder) GetSettingsOk() (*LogForwarderSetting, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *LogForwarder) SetSettings(v LogForwarderSetting) {
	o.Settings = v
}

// GetStatus returns the Status field value
func (o *LogForwarder) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LogForwarder) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LogForwarder) SetStatus(v bool) {
	o.Status = v
}

func (o LogForwarder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	toSerialize["connection_type"] = o.ConnectionType
	toSerialize["data_format"] = o.DataFormat
	toSerialize["settings"] = o.Settings
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableLogForwarder struct {
	value *LogForwarder
	isSet bool
}

func (v NullableLogForwarder) Get() *LogForwarder {
	return v.value
}

func (v *NullableLogForwarder) Set(val *LogForwarder) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarder) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarder(val *LogForwarder) *NullableLogForwarder {
	return &NullableLogForwarder{value: val, isSet: true}
}

func (v NullableLogForwarder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


