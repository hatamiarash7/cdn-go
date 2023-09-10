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

// checks if the WafPackageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPackageDetails{}

// WafPackageDetails struct for WafPackageDetails
type WafPackageDetails struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Provider *WafPackageProvider `json:"provider,omitempty"`
	// JSON-schema of parameters of the package
	ParamsSchema map[string]interface{} `json:"params_schema,omitempty"`
	// It will be filled by default disabled rules when it's not provided
	DisabledRules []string `json:"disabled_rules,omitempty"`
	// It will be filled by default disabled rulesets when it's not provided
	DisabledRulesets []string `json:"disabled_rulesets,omitempty"`
	Rulesets []WafRuleset `json:"rulesets,omitempty"`
}

// NewWafPackageDetails instantiates a new WafPackageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPackageDetails() *WafPackageDetails {
	this := WafPackageDetails{}
	return &this
}

// NewWafPackageDetailsWithDefaults instantiates a new WafPackageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPackageDetailsWithDefaults() *WafPackageDetails {
	this := WafPackageDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WafPackageDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackageDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WafPackageDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WafPackageDetails) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafPackageDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackageDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafPackageDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafPackageDetails) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *WafPackageDetails) GetProvider() WafPackageProvider {
	if o == nil || IsNil(o.Provider) {
		var ret WafPackageProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackageDetails) GetProviderOk() (*WafPackageProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *WafPackageDetails) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given WafPackageProvider and assigns it to the Provider field.
func (o *WafPackageDetails) SetProvider(v WafPackageProvider) {
	o.Provider = &v
}

// GetParamsSchema returns the ParamsSchema field value if set, zero value otherwise.
func (o *WafPackageDetails) GetParamsSchema() map[string]interface{} {
	if o == nil || IsNil(o.ParamsSchema) {
		var ret map[string]interface{}
		return ret
	}
	return o.ParamsSchema
}

// GetParamsSchemaOk returns a tuple with the ParamsSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackageDetails) GetParamsSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ParamsSchema) {
		return map[string]interface{}{}, false
	}
	return o.ParamsSchema, true
}

// HasParamsSchema returns a boolean if a field has been set.
func (o *WafPackageDetails) HasParamsSchema() bool {
	if o != nil && !IsNil(o.ParamsSchema) {
		return true
	}

	return false
}

// SetParamsSchema gets a reference to the given map[string]interface{} and assigns it to the ParamsSchema field.
func (o *WafPackageDetails) SetParamsSchema(v map[string]interface{}) {
	o.ParamsSchema = v
}

// GetDisabledRules returns the DisabledRules field value if set, zero value otherwise.
func (o *WafPackageDetails) GetDisabledRules() []string {
	if o == nil || IsNil(o.DisabledRules) {
		var ret []string
		return ret
	}
	return o.DisabledRules
}

// GetDisabledRulesOk returns a tuple with the DisabledRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackageDetails) GetDisabledRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledRules) {
		return nil, false
	}
	return o.DisabledRules, true
}

// HasDisabledRules returns a boolean if a field has been set.
func (o *WafPackageDetails) HasDisabledRules() bool {
	if o != nil && !IsNil(o.DisabledRules) {
		return true
	}

	return false
}

// SetDisabledRules gets a reference to the given []string and assigns it to the DisabledRules field.
func (o *WafPackageDetails) SetDisabledRules(v []string) {
	o.DisabledRules = v
}

// GetDisabledRulesets returns the DisabledRulesets field value if set, zero value otherwise.
func (o *WafPackageDetails) GetDisabledRulesets() []string {
	if o == nil || IsNil(o.DisabledRulesets) {
		var ret []string
		return ret
	}
	return o.DisabledRulesets
}

// GetDisabledRulesetsOk returns a tuple with the DisabledRulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackageDetails) GetDisabledRulesetsOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledRulesets) {
		return nil, false
	}
	return o.DisabledRulesets, true
}

// HasDisabledRulesets returns a boolean if a field has been set.
func (o *WafPackageDetails) HasDisabledRulesets() bool {
	if o != nil && !IsNil(o.DisabledRulesets) {
		return true
	}

	return false
}

// SetDisabledRulesets gets a reference to the given []string and assigns it to the DisabledRulesets field.
func (o *WafPackageDetails) SetDisabledRulesets(v []string) {
	o.DisabledRulesets = v
}

// GetRulesets returns the Rulesets field value if set, zero value otherwise.
func (o *WafPackageDetails) GetRulesets() []WafRuleset {
	if o == nil || IsNil(o.Rulesets) {
		var ret []WafRuleset
		return ret
	}
	return o.Rulesets
}

// GetRulesetsOk returns a tuple with the Rulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPackageDetails) GetRulesetsOk() ([]WafRuleset, bool) {
	if o == nil || IsNil(o.Rulesets) {
		return nil, false
	}
	return o.Rulesets, true
}

// HasRulesets returns a boolean if a field has been set.
func (o *WafPackageDetails) HasRulesets() bool {
	if o != nil && !IsNil(o.Rulesets) {
		return true
	}

	return false
}

// SetRulesets gets a reference to the given []WafRuleset and assigns it to the Rulesets field.
func (o *WafPackageDetails) SetRulesets(v []WafRuleset) {
	o.Rulesets = v
}

func (o WafPackageDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPackageDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Rulesets) {
		toSerialize["rulesets"] = o.Rulesets
	}
	return toSerialize, nil
}

type NullableWafPackageDetails struct {
	value *WafPackageDetails
	isSet bool
}

func (v NullableWafPackageDetails) Get() *WafPackageDetails {
	return v.value
}

func (v *NullableWafPackageDetails) Set(val *WafPackageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPackageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPackageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPackageDetails(val *WafPackageDetails) *NullableWafPackageDetails {
	return &NullableWafPackageDetails{value: val, isSet: true}
}

func (v NullableWafPackageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPackageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


