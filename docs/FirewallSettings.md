# FirewallSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultActionDetails** | Pointer to [**NullableFirewallActionDetails**](FirewallActionDetails.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**DefaultAction** | Pointer to **string** |  | [optional] 
**VerifySni** | Pointer to **bool** | True to verify that SNI and hostname are equal | [optional] [default to false]
**SkipGlobalWhitelist** | Pointer to **NullableBool** | Shows hether global whitelist should be skipped for the domain or not | [optional] [default to false]
**SkipGlobalFirewall** | Pointer to **NullableBool** | Shows whether global firewall should be skipped for the domain or not | [optional] [default to false]

## Methods

### NewFirewallSettings

`func NewFirewallSettings() *FirewallSettings`

NewFirewallSettings instantiates a new FirewallSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSettingsWithDefaults

`func NewFirewallSettingsWithDefaults() *FirewallSettings`

NewFirewallSettingsWithDefaults instantiates a new FirewallSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultActionDetails

`func (o *FirewallSettings) GetDefaultActionDetails() FirewallActionDetails`

GetDefaultActionDetails returns the DefaultActionDetails field if non-nil, zero value otherwise.

### GetDefaultActionDetailsOk

`func (o *FirewallSettings) GetDefaultActionDetailsOk() (*FirewallActionDetails, bool)`

GetDefaultActionDetailsOk returns a tuple with the DefaultActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultActionDetails

`func (o *FirewallSettings) SetDefaultActionDetails(v FirewallActionDetails)`

SetDefaultActionDetails sets DefaultActionDetails field to given value.

### HasDefaultActionDetails

`func (o *FirewallSettings) HasDefaultActionDetails() bool`

HasDefaultActionDetails returns a boolean if a field has been set.

### SetDefaultActionDetailsNil

`func (o *FirewallSettings) SetDefaultActionDetailsNil(b bool)`

 SetDefaultActionDetailsNil sets the value for DefaultActionDetails to be an explicit nil

### UnsetDefaultActionDetails
`func (o *FirewallSettings) UnsetDefaultActionDetails()`

UnsetDefaultActionDetails ensures that no value is present for DefaultActionDetails, not even an explicit nil
### GetIsEnabled

`func (o *FirewallSettings) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FirewallSettings) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FirewallSettings) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FirewallSettings) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDefaultAction

`func (o *FirewallSettings) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *FirewallSettings) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *FirewallSettings) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *FirewallSettings) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetVerifySni

`func (o *FirewallSettings) GetVerifySni() bool`

GetVerifySni returns the VerifySni field if non-nil, zero value otherwise.

### GetVerifySniOk

`func (o *FirewallSettings) GetVerifySniOk() (*bool, bool)`

GetVerifySniOk returns a tuple with the VerifySni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySni

`func (o *FirewallSettings) SetVerifySni(v bool)`

SetVerifySni sets VerifySni field to given value.

### HasVerifySni

`func (o *FirewallSettings) HasVerifySni() bool`

HasVerifySni returns a boolean if a field has been set.

### GetSkipGlobalWhitelist

`func (o *FirewallSettings) GetSkipGlobalWhitelist() bool`

GetSkipGlobalWhitelist returns the SkipGlobalWhitelist field if non-nil, zero value otherwise.

### GetSkipGlobalWhitelistOk

`func (o *FirewallSettings) GetSkipGlobalWhitelistOk() (*bool, bool)`

GetSkipGlobalWhitelistOk returns a tuple with the SkipGlobalWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGlobalWhitelist

`func (o *FirewallSettings) SetSkipGlobalWhitelist(v bool)`

SetSkipGlobalWhitelist sets SkipGlobalWhitelist field to given value.

### HasSkipGlobalWhitelist

`func (o *FirewallSettings) HasSkipGlobalWhitelist() bool`

HasSkipGlobalWhitelist returns a boolean if a field has been set.

### SetSkipGlobalWhitelistNil

`func (o *FirewallSettings) SetSkipGlobalWhitelistNil(b bool)`

 SetSkipGlobalWhitelistNil sets the value for SkipGlobalWhitelist to be an explicit nil

### UnsetSkipGlobalWhitelist
`func (o *FirewallSettings) UnsetSkipGlobalWhitelist()`

UnsetSkipGlobalWhitelist ensures that no value is present for SkipGlobalWhitelist, not even an explicit nil
### GetSkipGlobalFirewall

`func (o *FirewallSettings) GetSkipGlobalFirewall() bool`

GetSkipGlobalFirewall returns the SkipGlobalFirewall field if non-nil, zero value otherwise.

### GetSkipGlobalFirewallOk

`func (o *FirewallSettings) GetSkipGlobalFirewallOk() (*bool, bool)`

GetSkipGlobalFirewallOk returns a tuple with the SkipGlobalFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipGlobalFirewall

`func (o *FirewallSettings) SetSkipGlobalFirewall(v bool)`

SetSkipGlobalFirewall sets SkipGlobalFirewall field to given value.

### HasSkipGlobalFirewall

`func (o *FirewallSettings) HasSkipGlobalFirewall() bool`

HasSkipGlobalFirewall returns a boolean if a field has been set.

### SetSkipGlobalFirewallNil

`func (o *FirewallSettings) SetSkipGlobalFirewallNil(b bool)`

 SetSkipGlobalFirewallNil sets the value for SkipGlobalFirewall to be an explicit nil

### UnsetSkipGlobalFirewall
`func (o *FirewallSettings) UnsetSkipGlobalFirewall()`

UnsetSkipGlobalFirewall ensures that no value is present for SkipGlobalFirewall, not even an explicit nil

[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


