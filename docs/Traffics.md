# Traffics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**TrafficStatistics**](TrafficStatistics.md) |  | [optional] 
**Charts** | Pointer to [**TrafficCharts**](TrafficCharts.md) |  | [optional] 

## Methods

### NewTraffics

`func NewTraffics() *Traffics`

NewTraffics instantiates a new Traffics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficsWithDefaults

`func NewTrafficsWithDefaults() *Traffics`

NewTrafficsWithDefaults instantiates a new Traffics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *Traffics) GetStatistics() TrafficStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Traffics) GetStatisticsOk() (*TrafficStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Traffics) SetStatistics(v TrafficStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Traffics) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *Traffics) GetCharts() TrafficCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *Traffics) GetChartsOk() (*TrafficCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *Traffics) SetCharts(v TrafficCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *Traffics) HasCharts() bool`

HasCharts returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


