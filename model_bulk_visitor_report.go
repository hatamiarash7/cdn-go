/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.108.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the BulkVisitorReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkVisitorReport{}

// BulkVisitorReport struct for BulkVisitorReport
type BulkVisitorReport struct {
	// Domain's ID
	Resource *string `json:"resource,omitempty"`
	Success *bool `json:"success,omitempty"`
	// The error message
	Message NullableString `json:"message,omitempty"`
	Data *BulkVisitorReportData `json:"data,omitempty"`
}

// NewBulkVisitorReport instantiates a new BulkVisitorReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkVisitorReport() *BulkVisitorReport {
	this := BulkVisitorReport{}
	return &this
}

// NewBulkVisitorReportWithDefaults instantiates a new BulkVisitorReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkVisitorReportWithDefaults() *BulkVisitorReport {
	this := BulkVisitorReport{}
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *BulkVisitorReport) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkVisitorReport) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *BulkVisitorReport) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *BulkVisitorReport) SetResource(v string) {
	o.Resource = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkVisitorReport) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkVisitorReport) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkVisitorReport) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *BulkVisitorReport) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkVisitorReport) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkVisitorReport) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *BulkVisitorReport) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *BulkVisitorReport) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *BulkVisitorReport) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *BulkVisitorReport) UnsetMessage() {
	o.Message.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BulkVisitorReport) GetData() BulkVisitorReportData {
	if o == nil || IsNil(o.Data) {
		var ret BulkVisitorReportData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkVisitorReport) GetDataOk() (*BulkVisitorReportData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BulkVisitorReport) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BulkVisitorReportData and assigns it to the Data field.
func (o *BulkVisitorReport) SetData(v BulkVisitorReportData) {
	o.Data = &v
}

func (o BulkVisitorReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkVisitorReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBulkVisitorReport struct {
	value *BulkVisitorReport
	isSet bool
}

func (v NullableBulkVisitorReport) Get() *BulkVisitorReport {
	return v.value
}

func (v *NullableBulkVisitorReport) Set(val *BulkVisitorReport) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkVisitorReport) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkVisitorReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkVisitorReport(val *BulkVisitorReport) *NullableBulkVisitorReport {
	return &NullableBulkVisitorReport{value: val, isSet: true}
}

func (v NullableBulkVisitorReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkVisitorReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


