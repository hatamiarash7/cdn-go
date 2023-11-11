# LoadBalancerUpdatePoolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origins** | Pointer to [**[]LoadBalancerOriginStore**](LoadBalancerOriginStore.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **int32** | Zero means the default pool | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Keepalive** | Pointer to **string** |  | [optional] [default to "off"]
**NextUpstreamTcp** | Pointer to [**NextUpstreamTcp**](NextUpstreamTcp.md) |  | [optional] [default to FALSE]
**NextUpstreamTcpCodes** | Pointer to [**NextUpstreamTcpCodes**](NextUpstreamTcpCodes.md) |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**MonitoringStatus** | Pointer to [**NullableMonitoringStatus**](MonitoringStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewLoadBalancerUpdatePoolsInner

`func NewLoadBalancerUpdatePoolsInner() *LoadBalancerUpdatePoolsInner`

NewLoadBalancerUpdatePoolsInner instantiates a new LoadBalancerUpdatePoolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerUpdatePoolsInnerWithDefaults

`func NewLoadBalancerUpdatePoolsInnerWithDefaults() *LoadBalancerUpdatePoolsInner`

NewLoadBalancerUpdatePoolsInnerWithDefaults instantiates a new LoadBalancerUpdatePoolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigins

`func (o *LoadBalancerUpdatePoolsInner) GetOrigins() []LoadBalancerOriginStore`

GetOrigins returns the Origins field if non-nil, zero value otherwise.

### GetOriginsOk

`func (o *LoadBalancerUpdatePoolsInner) GetOriginsOk() (*[]LoadBalancerOriginStore, bool)`

GetOriginsOk returns a tuple with the Origins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigins

`func (o *LoadBalancerUpdatePoolsInner) SetOrigins(v []LoadBalancerOriginStore)`

SetOrigins sets Origins field to given value.

### HasOrigins

`func (o *LoadBalancerUpdatePoolsInner) HasOrigins() bool`

HasOrigins returns a boolean if a field has been set.

### GetId

`func (o *LoadBalancerUpdatePoolsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerUpdatePoolsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerUpdatePoolsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerUpdatePoolsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancerUpdatePoolsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerUpdatePoolsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerUpdatePoolsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerUpdatePoolsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LoadBalancerUpdatePoolsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancerUpdatePoolsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancerUpdatePoolsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancerUpdatePoolsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *LoadBalancerUpdatePoolsInner) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoadBalancerUpdatePoolsInner) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoadBalancerUpdatePoolsInner) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LoadBalancerUpdatePoolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *LoadBalancerUpdatePoolsInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LoadBalancerUpdatePoolsInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LoadBalancerUpdatePoolsInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LoadBalancerUpdatePoolsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetMethod

`func (o *LoadBalancerUpdatePoolsInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoadBalancerUpdatePoolsInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoadBalancerUpdatePoolsInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoadBalancerUpdatePoolsInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetKeepalive

`func (o *LoadBalancerUpdatePoolsInner) GetKeepalive() string`

GetKeepalive returns the Keepalive field if non-nil, zero value otherwise.

### GetKeepaliveOk

`func (o *LoadBalancerUpdatePoolsInner) GetKeepaliveOk() (*string, bool)`

GetKeepaliveOk returns a tuple with the Keepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepalive

`func (o *LoadBalancerUpdatePoolsInner) SetKeepalive(v string)`

SetKeepalive sets Keepalive field to given value.

### HasKeepalive

`func (o *LoadBalancerUpdatePoolsInner) HasKeepalive() bool`

HasKeepalive returns a boolean if a field has been set.

### GetNextUpstreamTcp

`func (o *LoadBalancerUpdatePoolsInner) GetNextUpstreamTcp() NextUpstreamTcp`

GetNextUpstreamTcp returns the NextUpstreamTcp field if non-nil, zero value otherwise.

### GetNextUpstreamTcpOk

`func (o *LoadBalancerUpdatePoolsInner) GetNextUpstreamTcpOk() (*NextUpstreamTcp, bool)`

GetNextUpstreamTcpOk returns a tuple with the NextUpstreamTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpstreamTcp

`func (o *LoadBalancerUpdatePoolsInner) SetNextUpstreamTcp(v NextUpstreamTcp)`

SetNextUpstreamTcp sets NextUpstreamTcp field to given value.

### HasNextUpstreamTcp

`func (o *LoadBalancerUpdatePoolsInner) HasNextUpstreamTcp() bool`

HasNextUpstreamTcp returns a boolean if a field has been set.

### GetNextUpstreamTcpCodes

`func (o *LoadBalancerUpdatePoolsInner) GetNextUpstreamTcpCodes() NextUpstreamTcpCodes`

GetNextUpstreamTcpCodes returns the NextUpstreamTcpCodes field if non-nil, zero value otherwise.

### GetNextUpstreamTcpCodesOk

`func (o *LoadBalancerUpdatePoolsInner) GetNextUpstreamTcpCodesOk() (*NextUpstreamTcpCodes, bool)`

GetNextUpstreamTcpCodesOk returns a tuple with the NextUpstreamTcpCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpstreamTcpCodes

`func (o *LoadBalancerUpdatePoolsInner) SetNextUpstreamTcpCodes(v NextUpstreamTcpCodes)`

SetNextUpstreamTcpCodes sets NextUpstreamTcpCodes field to given value.

### HasNextUpstreamTcpCodes

`func (o *LoadBalancerUpdatePoolsInner) HasNextUpstreamTcpCodes() bool`

HasNextUpstreamTcpCodes returns a boolean if a field has been set.

### GetRegions

`func (o *LoadBalancerUpdatePoolsInner) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *LoadBalancerUpdatePoolsInner) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *LoadBalancerUpdatePoolsInner) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *LoadBalancerUpdatePoolsInner) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetMonitoringStatus

`func (o *LoadBalancerUpdatePoolsInner) GetMonitoringStatus() MonitoringStatus`

GetMonitoringStatus returns the MonitoringStatus field if non-nil, zero value otherwise.

### GetMonitoringStatusOk

`func (o *LoadBalancerUpdatePoolsInner) GetMonitoringStatusOk() (*MonitoringStatus, bool)`

GetMonitoringStatusOk returns a tuple with the MonitoringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStatus

`func (o *LoadBalancerUpdatePoolsInner) SetMonitoringStatus(v MonitoringStatus)`

SetMonitoringStatus sets MonitoringStatus field to given value.

### HasMonitoringStatus

`func (o *LoadBalancerUpdatePoolsInner) HasMonitoringStatus() bool`

HasMonitoringStatus returns a boolean if a field has been set.

### SetMonitoringStatusNil

`func (o *LoadBalancerUpdatePoolsInner) SetMonitoringStatusNil(b bool)`

 SetMonitoringStatusNil sets the value for MonitoringStatus to be an explicit nil

### UnsetMonitoringStatus
`func (o *LoadBalancerUpdatePoolsInner) UnsetMonitoringStatus()`

UnsetMonitoringStatus ensures that no value is present for MonitoringStatus, not even an explicit nil
### GetCreatedAt

`func (o *LoadBalancerUpdatePoolsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadBalancerUpdatePoolsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadBalancerUpdatePoolsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadBalancerUpdatePoolsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LoadBalancerUpdatePoolsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoadBalancerUpdatePoolsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoadBalancerUpdatePoolsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoadBalancerUpdatePoolsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


