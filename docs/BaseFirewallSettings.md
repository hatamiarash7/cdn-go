# BaseFirewallSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**DefaultAction** | Pointer to **string** |  | [optional] 
**VerifySni** | Pointer to **bool** | True to verify that SNI and hostname are equal | [optional] [default to false]
**SkipGlobalWhitelist** | Pointer to **NullableBool** | Shows hether global whitelist should be skipped for the domain or not | [optional] [default to false]
**SkipGlobalFirewall** | Pointer to **NullableBool** | Shows whether global firewall should be skipped for the domain or not | [optional] [default to false]

## Methods

### NewBaseFirewallSettings

`func NewBaseFirewallSettings() *BaseFirewallSettings`

NewBaseFirewallSettings instantiates a new BaseFirewallSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFirewallSettingsWithDefaults

`func NewBaseFirewallSettingsWithDefaults() *BaseFirewallSettings`

NewBaseFirewallSettingsWithDefaults instantiates a new BaseFirewallSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *BaseFirewallSettings) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BaseFirewallSettings) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BaseFirewallSettings) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *BaseFirewallSettings) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDefaultAction

`func (o *BaseFirewallSettings) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *BaseFirewallSettings) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *BaseFirewallSettings) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *BaseFirewallSettings) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetVerifySni

`func (o *BaseFirewallSettings) GetVerifySni() bool`

GetVerifySni returns the VerifySni field if non-nil, zero value otherwise.

### GetVerifySniOk

`func (o *BaseFirewallSettings) GetVerifySniOk() (*bool, bool)`

GetVerifySniOk returns a tuple with the VerifySni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySni

`func (o *BaseFirewallSettings) SetVerifySni(v bool)`

SetVerifySni sets VerifySni field to given value.

### HasVerifySni

`func (o *BaseFirewallSettings) HasVerifySni() bool`

HasVerifySni returns a boolean if a field has been set.

### GetSkipGlobalWhitelist

`func (o *BaseFirewallSettings) GetSkipGlobalWhitelist() bool`

GetSkipGlobalWhitelist returns the SkipGlobalWhitelist field if non-nil, zero value otherwise.

### GetSkipGlobalWhitelistOk

`func (o *BaseFirewallSettings) GetSkipGlobalWhitelistOk() (*bool, bool)`

GetSkipGlobalWhitelistOk returns a tuple with the SkipGlobalWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGlobalWhitelist

`func (o *BaseFirewallSettings) SetSkipGlobalWhitelist(v bool)`

SetSkipGlobalWhitelist sets SkipGlobalWhitelist field to given value.

### HasSkipGlobalWhitelist

`func (o *BaseFirewallSettings) HasSkipGlobalWhitelist() bool`

HasSkipGlobalWhitelist returns a boolean if a field has been set.

### SetSkipGlobalWhitelistNil

`func (o *BaseFirewallSettings) SetSkipGlobalWhitelistNil(b bool)`

 SetSkipGlobalWhitelistNil sets the value for SkipGlobalWhitelist to be an explicit nil

### UnsetSkipGlobalWhitelist
`func (o *BaseFirewallSettings) UnsetSkipGlobalWhitelist()`

UnsetSkipGlobalWhitelist ensures that no value is present for SkipGlobalWhitelist, not even an explicit nil
### GetSkipGlobalFirewall

`func (o *BaseFirewallSettings) GetSkipGlobalFirewall() bool`

GetSkipGlobalFirewall returns the SkipGlobalFirewall field if non-nil, zero value otherwise.

### GetSkipGlobalFirewallOk

`func (o *BaseFirewallSettings) GetSkipGlobalFirewallOk() (*bool, bool)`

GetSkipGlobalFirewallOk returns a tuple with the SkipGlobalFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGlobalFirewall

`func (o *BaseFirewallSettings) SetSkipGlobalFirewall(v bool)`

SetSkipGlobalFirewall sets SkipGlobalFirewall field to given value.

### HasSkipGlobalFirewall

`func (o *BaseFirewallSettings) HasSkipGlobalFirewall() bool`

HasSkipGlobalFirewall returns a boolean if a field has been set.

### SetSkipGlobalFirewallNil

`func (o *BaseFirewallSettings) SetSkipGlobalFirewallNil(b bool)`

 SetSkipGlobalFirewallNil sets the value for SkipGlobalFirewall to be an explicit nil

### UnsetSkipGlobalFirewall
`func (o *BaseFirewallSettings) UnsetSkipGlobalFirewall()`

UnsetSkipGlobalFirewall ensures that no value is present for SkipGlobalFirewall, not even an explicit nil

[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


