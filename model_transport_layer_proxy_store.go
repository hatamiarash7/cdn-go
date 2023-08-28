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

// checks if the TransportLayerProxyStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxyStore{}

// TransportLayerProxyStore struct for TransportLayerProxyStore
type TransportLayerProxyStore struct {
	AppName string `json:"app_name"`
	Description *string `json:"description,omitempty"`
	Domain string `json:"domain"`
	Port int32 `json:"port"`
	ProxyProtocol string `json:"proxy_protocol"`
	BalanceAlgorithm string `json:"balance_algorithm"`
	Servers []TransportLayerProxyServer `json:"servers,omitempty"`
	FirewallDefaultAction string `json:"firewall_default_action"`
	Firewalls []TransportLayerProxyFirewall `json:"firewalls,omitempty"`
}

// NewTransportLayerProxyStore instantiates a new TransportLayerProxyStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxyStore(appName string, domain string, port int32, proxyProtocol string, balanceAlgorithm string, firewallDefaultAction string) *TransportLayerProxyStore {
	this := TransportLayerProxyStore{}
	this.AppName = appName
	this.Domain = domain
	this.Port = port
	this.ProxyProtocol = proxyProtocol
	this.BalanceAlgorithm = balanceAlgorithm
	this.FirewallDefaultAction = firewallDefaultAction
	return &this
}

// NewTransportLayerProxyStoreWithDefaults instantiates a new TransportLayerProxyStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyStoreWithDefaults() *TransportLayerProxyStore {
	this := TransportLayerProxyStore{}
	return &this
}

// GetAppName returns the AppName field value
func (o *TransportLayerProxyStore) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *TransportLayerProxyStore) SetAppName(v string) {
	o.AppName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransportLayerProxyStore) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransportLayerProxyStore) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransportLayerProxyStore) SetDescription(v string) {
	o.Description = &v
}

// GetDomain returns the Domain field value
func (o *TransportLayerProxyStore) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *TransportLayerProxyStore) SetDomain(v string) {
	o.Domain = v
}

// GetPort returns the Port field value
func (o *TransportLayerProxyStore) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *TransportLayerProxyStore) SetPort(v int32) {
	o.Port = v
}

// GetProxyProtocol returns the ProxyProtocol field value
func (o *TransportLayerProxyStore) GetProxyProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProxyProtocol
}

// GetProxyProtocolOk returns a tuple with the ProxyProtocol field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetProxyProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProxyProtocol, true
}

// SetProxyProtocol sets field value
func (o *TransportLayerProxyStore) SetProxyProtocol(v string) {
	o.ProxyProtocol = v
}

// GetBalanceAlgorithm returns the BalanceAlgorithm field value
func (o *TransportLayerProxyStore) GetBalanceAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceAlgorithm
}

// GetBalanceAlgorithmOk returns a tuple with the BalanceAlgorithm field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetBalanceAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceAlgorithm, true
}

// SetBalanceAlgorithm sets field value
func (o *TransportLayerProxyStore) SetBalanceAlgorithm(v string) {
	o.BalanceAlgorithm = v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *TransportLayerProxyStore) GetServers() []TransportLayerProxyServer {
	if o == nil || IsNil(o.Servers) {
		var ret []TransportLayerProxyServer
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetServersOk() ([]TransportLayerProxyServer, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *TransportLayerProxyStore) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []TransportLayerProxyServer and assigns it to the Servers field.
func (o *TransportLayerProxyStore) SetServers(v []TransportLayerProxyServer) {
	o.Servers = v
}

// GetFirewallDefaultAction returns the FirewallDefaultAction field value
func (o *TransportLayerProxyStore) GetFirewallDefaultAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirewallDefaultAction
}

// GetFirewallDefaultActionOk returns a tuple with the FirewallDefaultAction field value
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetFirewallDefaultActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirewallDefaultAction, true
}

// SetFirewallDefaultAction sets field value
func (o *TransportLayerProxyStore) SetFirewallDefaultAction(v string) {
	o.FirewallDefaultAction = v
}

// GetFirewalls returns the Firewalls field value if set, zero value otherwise.
func (o *TransportLayerProxyStore) GetFirewalls() []TransportLayerProxyFirewall {
	if o == nil || IsNil(o.Firewalls) {
		var ret []TransportLayerProxyFirewall
		return ret
	}
	return o.Firewalls
}

// GetFirewallsOk returns a tuple with the Firewalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyStore) GetFirewallsOk() ([]TransportLayerProxyFirewall, bool) {
	if o == nil || IsNil(o.Firewalls) {
		return nil, false
	}
	return o.Firewalls, true
}

// HasFirewalls returns a boolean if a field has been set.
func (o *TransportLayerProxyStore) HasFirewalls() bool {
	if o != nil && !IsNil(o.Firewalls) {
		return true
	}

	return false
}

// SetFirewalls gets a reference to the given []TransportLayerProxyFirewall and assigns it to the Firewalls field.
func (o *TransportLayerProxyStore) SetFirewalls(v []TransportLayerProxyFirewall) {
	o.Firewalls = v
}

func (o TransportLayerProxyStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxyStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_name"] = o.AppName
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["domain"] = o.Domain
	toSerialize["port"] = o.Port
	toSerialize["proxy_protocol"] = o.ProxyProtocol
	toSerialize["balance_algorithm"] = o.BalanceAlgorithm
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	toSerialize["firewall_default_action"] = o.FirewallDefaultAction
	if !IsNil(o.Firewalls) {
		toSerialize["firewalls"] = o.Firewalls
	}
	return toSerialize, nil
}

type NullableTransportLayerProxyStore struct {
	value *TransportLayerProxyStore
	isSet bool
}

func (v NullableTransportLayerProxyStore) Get() *TransportLayerProxyStore {
	return v.value
}

func (v *NullableTransportLayerProxyStore) Set(val *TransportLayerProxyStore) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyStore) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyStore(val *TransportLayerProxyStore) *NullableTransportLayerProxyStore {
	return &NullableTransportLayerProxyStore{value: val, isSet: true}
}

func (v NullableTransportLayerProxyStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


