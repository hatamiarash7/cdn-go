# FirewallSettingsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultActionDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**DefaultAction** | Pointer to **string** |  | [optional] 
**VerifySni** | Pointer to **bool** | True to verify that SNI and hostname are equal | [optional] [default to false]
**SkipGlobalWhitelist** | Pointer to **NullableBool** | Shows hether global whitelist should be skipped for the domain or not | [optional] [default to false]
**SkipGlobalFirewall** | Pointer to **NullableBool** | Shows whether global firewall should be skipped for the domain or not | [optional] [default to false]

## Methods

### NewFirewallSettingsView

`func NewFirewallSettingsView() *FirewallSettingsView`

NewFirewallSettingsView instantiates a new FirewallSettingsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSettingsViewWithDefaults

`func NewFirewallSettingsViewWithDefaults() *FirewallSettingsView`

NewFirewallSettingsViewWithDefaults instantiates a new FirewallSettingsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultActionDetails

`func (o *FirewallSettingsView) GetDefaultActionDetails() map[string]interface{}`

GetDefaultActionDetails returns the DefaultActionDetails field if non-nil, zero value otherwise.

### GetDefaultActionDetailsOk

`func (o *FirewallSettingsView) GetDefaultActionDetailsOk() (*map[string]interface{}, bool)`

GetDefaultActionDetailsOk returns a tuple with the DefaultActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultActionDetails

`func (o *FirewallSettingsView) SetDefaultActionDetails(v map[string]interface{})`

SetDefaultActionDetails sets DefaultActionDetails field to given value.

### HasDefaultActionDetails

`func (o *FirewallSettingsView) HasDefaultActionDetails() bool`

HasDefaultActionDetails returns a boolean if a field has been set.

### SetDefaultActionDetailsNil

`func (o *FirewallSettingsView) SetDefaultActionDetailsNil(b bool)`

 SetDefaultActionDetailsNil sets the value for DefaultActionDetails to be an explicit nil

### UnsetDefaultActionDetails
`func (o *FirewallSettingsView) UnsetDefaultActionDetails()`

UnsetDefaultActionDetails ensures that no value is present for DefaultActionDetails, not even an explicit nil
### GetIsEnabled

`func (o *FirewallSettingsView) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FirewallSettingsView) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FirewallSettingsView) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FirewallSettingsView) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDefaultAction

`func (o *FirewallSettingsView) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *FirewallSettingsView) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *FirewallSettingsView) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *FirewallSettingsView) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetVerifySni

`func (o *FirewallSettingsView) GetVerifySni() bool`

GetVerifySni returns the VerifySni field if non-nil, zero value otherwise.

### GetVerifySniOk

`func (o *FirewallSettingsView) GetVerifySniOk() (*bool, bool)`

GetVerifySniOk returns a tuple with the VerifySni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySni

`func (o *FirewallSettingsView) SetVerifySni(v bool)`

SetVerifySni sets VerifySni field to given value.

### HasVerifySni

`func (o *FirewallSettingsView) HasVerifySni() bool`

HasVerifySni returns a boolean if a field has been set.

### GetSkipGlobalWhitelist

`func (o *FirewallSettingsView) GetSkipGlobalWhitelist() bool`

GetSkipGlobalWhitelist returns the SkipGlobalWhitelist field if non-nil, zero value otherwise.

### GetSkipGlobalWhitelistOk

`func (o *FirewallSettingsView) GetSkipGlobalWhitelistOk() (*bool, bool)`

GetSkipGlobalWhitelistOk returns a tuple with the SkipGlobalWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGlobalWhitelist

`func (o *FirewallSettingsView) SetSkipGlobalWhitelist(v bool)`

SetSkipGlobalWhitelist sets SkipGlobalWhitelist field to given value.

### HasSkipGlobalWhitelist

`func (o *FirewallSettingsView) HasSkipGlobalWhitelist() bool`

HasSkipGlobalWhitelist returns a boolean if a field has been set.

### SetSkipGlobalWhitelistNil

`func (o *FirewallSettingsView) SetSkipGlobalWhitelistNil(b bool)`

 SetSkipGlobalWhitelistNil sets the value for SkipGlobalWhitelist to be an explicit nil

### UnsetSkipGlobalWhitelist
`func (o *FirewallSettingsView) UnsetSkipGlobalWhitelist()`

UnsetSkipGlobalWhitelist ensures that no value is present for SkipGlobalWhitelist, not even an explicit nil
### GetSkipGlobalFirewall

`func (o *FirewallSettingsView) GetSkipGlobalFirewall() bool`

GetSkipGlobalFirewall returns the SkipGlobalFirewall field if non-nil, zero value otherwise.

### GetSkipGlobalFirewallOk

`func (o *FirewallSettingsView) GetSkipGlobalFirewallOk() (*bool, bool)`

GetSkipGlobalFirewallOk returns a tuple with the SkipGlobalFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGlobalFirewall

`func (o *FirewallSettingsView) SetSkipGlobalFirewall(v bool)`

SetSkipGlobalFirewall sets SkipGlobalFirewall field to given value.

### HasSkipGlobalFirewall

`func (o *FirewallSettingsView) HasSkipGlobalFirewall() bool`

HasSkipGlobalFirewall returns a boolean if a field has been set.

### SetSkipGlobalFirewallNil

`func (o *FirewallSettingsView) SetSkipGlobalFirewallNil(b bool)`

 SetSkipGlobalFirewallNil sets the value for SkipGlobalFirewall to be an explicit nil

### UnsetSkipGlobalFirewall
`func (o *FirewallSettingsView) UnsetSkipGlobalFirewall()`

UnsetSkipGlobalFirewall ensures that no value is present for SkipGlobalFirewall, not even an explicit nil

[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


