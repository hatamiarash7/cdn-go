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

// checks if the WafPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPackage{}

// WafPackage struct for WafPackage
type WafPackage struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Provider *WafPackageProvider `json:"provider,omitempty"`
	// JSON-schema of parameters of the package
	ParamsSchema map[string]interface{} `json:"params_schema,omitempty"`
	// It will be filled by default disabled rules when it's not provided
	DisabledRules []string `json:"disabled_rules,omitempty"`
	// It will be filled by default disabled rulesets when it's not provided
	DisabledRulesets []string `json:"disabled_rulesets,omitempty"`
}

// NewWafPackage instantiates a new WafPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPackage() *WafPackage {
	this := WafPackage{}
	return &this
}

// NewWafPackageWithDefaults instantiates a new WafPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPackageWithDefaults() *WafPackage {
	this := WafPackage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WafPackage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WafPackage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WafPackage) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafPackage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafPackage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafPackage) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *WafPackage) GetProvider() WafPackageProvider {
	if o == nil || IsNil(o.Provider) {
		var ret WafPackageProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackage) GetProviderOk() (*WafPackageProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *WafPackage) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given WafPackageProvider and assigns it to the Provider field.
func (o *WafPackage) SetProvider(v WafPackageProvider) {
	o.Provider = &v
}

// GetParamsSchema returns the ParamsSchema field value if set, zero value otherwise.
func (o *WafPackage) GetParamsSchema() map[string]interface{} {
	if o == nil || IsNil(o.ParamsSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.ParamsSchema
}

// GetParamsSchemaOk returns a tuple with the ParamsSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackage) GetParamsSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ParamsSchema) {
		return map[string]interface{}{}, false
	}
	return o.ParamsSchema, true
}

// HasParamsSchema returns a boolean if a field has been set.
func (o *WafPackage) HasParamsSchema() bool {
	if o != nil && !IsNil(o.ParamsSchema) {
		return true
	}

	return false
}

// SetParamsSchema gets a reference to the given map[string]interface{} and assigns it to the ParamsSchema field.
func (o *WafPackage) SetParamsSchema(v map[string]interface{}) {
	o.ParamsSchema = v
}

// GetDisabledRules returns the DisabledRules field value if set, zero value otherwise.
func (o *WafPackage) GetDisabledRules() []string {
	if o == nil || IsNil(o.DisabledRules) {
		var ret []string
		return ret
	}
	return o.DisabledRules
}

// GetDisabledRulesOk returns a tuple with the DisabledRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackage) GetDisabledRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledRules) {
		return nil, false
	}
	return o.DisabledRules, true
}

// HasDisabledRules returns a boolean if a field has been set.
func (o *WafPackage) HasDisabledRules() bool {
	if o != nil && !IsNil(o.DisabledRules) {
		return true
	}

	return false
}

// SetDisabledRules gets a reference to the given []string and assigns it to the DisabledRules field.
func (o *WafPackage) SetDisabledRules(v []string) {
	o.DisabledRules = v
}

// GetDisabledRulesets returns the DisabledRulesets field value if set, zero value otherwise.
func (o *WafPackage) GetDisabledRulesets() []string {
	if o == nil || IsNil(o.DisabledRulesets) {
		var ret []string
		return ret
	}
	return o.DisabledRulesets
}

// GetDisabledRulesetsOk returns a tuple with the DisabledRulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackage) GetDisabledRulesetsOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledRulesets) {
		return nil, false
	}
	return o.DisabledRulesets, true
}

// HasDisabledRulesets returns a boolean if a field has been set.
func (o *WafPackage) HasDisabledRulesets() bool {
	if o != nil && !IsNil(o.DisabledRulesets) {
		return true
	}

	return false
}

// SetDisabledRulesets gets a reference to the given []string and assigns it to the DisabledRulesets field.
func (o *WafPackage) SetDisabledRulesets(v []string) {
	o.DisabledRulesets = v
}

func (o WafPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: name is readOnly
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.ParamsSchema) {
		toSerialize["params_schema"] = o.ParamsSchema
	}
	if !IsNil(o.DisabledRules) {
		toSerialize["disabled_rules"] = o.DisabledRules
	}
	if !IsNil(o.DisabledRulesets) {
		toSerialize["disabled_rulesets"] = o.DisabledRulesets
	}
	return toSerialize, nil
}

type NullableWafPackage struct {
	value *WafPackage
	isSet bool
}

func (v NullableWafPackage) Get() *WafPackage {
	return v.value
}

func (v *NullableWafPackage) Set(val *WafPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPackage(val *WafPackage) *NullableWafPackage {
	return &NullableWafPackage{value: val, isSet: true}
}

func (v NullableWafPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


