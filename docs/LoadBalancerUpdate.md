# LoadBalancerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**TimeSlice** | Pointer to **string** | Human friendly time duration for which a pool will uninterruptedly be selected in cluster_rr strategy, i.e. pools will switch once every time slice. | [optional] [default to "0s"]
**Pools** | Pointer to [**[]LoadBalancerUpdatePoolsInner**](LoadBalancerUpdatePoolsInner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewLoadBalancerUpdate

`func NewLoadBalancerUpdate() *LoadBalancerUpdate`

NewLoadBalancerUpdate instantiates a new LoadBalancerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerUpdateWithDefaults

`func NewLoadBalancerUpdateWithDefaults() *LoadBalancerUpdate`

NewLoadBalancerUpdateWithDefaults instantiates a new LoadBalancerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancerUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LoadBalancerUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancerUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancerUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancerUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *LoadBalancerUpdate) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoadBalancerUpdate) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoadBalancerUpdate) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LoadBalancerUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMethod

`func (o *LoadBalancerUpdate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoadBalancerUpdate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoadBalancerUpdate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoadBalancerUpdate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetTimeSlice

`func (o *LoadBalancerUpdate) GetTimeSlice() string`

GetTimeSlice returns the TimeSlice field if non-nil, zero value otherwise.

### GetTimeSliceOk

`func (o *LoadBalancerUpdate) GetTimeSliceOk() (*string, bool)`

GetTimeSliceOk returns a tuple with the TimeSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlice

`func (o *LoadBalancerUpdate) SetTimeSlice(v string)`

SetTimeSlice sets TimeSlice field to given value.

### HasTimeSlice

`func (o *LoadBalancerUpdate) HasTimeSlice() bool`

HasTimeSlice returns a boolean if a field has been set.

### GetPools

`func (o *LoadBalancerUpdate) GetPools() []LoadBalancerUpdatePoolsInner`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *LoadBalancerUpdate) GetPoolsOk() (*[]LoadBalancerUpdatePoolsInner, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *LoadBalancerUpdate) SetPools(v []LoadBalancerUpdatePoolsInner)`

SetPools sets Pools field to given value.

### HasPools

`func (o *LoadBalancerUpdate) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoadBalancerUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadBalancerUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadBalancerUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadBalancerUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LoadBalancerUpdate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoadBalancerUpdate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoadBalancerUpdate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoadBalancerUpdate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


