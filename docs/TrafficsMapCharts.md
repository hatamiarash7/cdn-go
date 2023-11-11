# TrafficsMapCharts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to [**map[string]CountryRequestChart**](CountryRequestChart.md) |  | [optional] 
**Traffics** | Pointer to [**map[string]CountryTrafficChart**](CountryTrafficChart.md) |  | [optional] 

## Methods

### NewTrafficsMapCharts

`func NewTrafficsMapCharts() *TrafficsMapCharts`

NewTrafficsMapCharts instantiates a new TrafficsMapCharts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficsMapChartsWithDefaults

`func NewTrafficsMapChartsWithDefaults() *TrafficsMapCharts`

NewTrafficsMapChartsWithDefaults instantiates a new TrafficsMapCharts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *TrafficsMapCharts) GetRequests() map[string]CountryRequestChart`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TrafficsMapCharts) GetRequestsOk() (*map[string]CountryRequestChart, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TrafficsMapCharts) SetRequests(v map[string]CountryRequestChart)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *TrafficsMapCharts) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTraffics

`func (o *TrafficsMapCharts) GetTraffics() map[string]CountryTrafficChart`

GetTraffics returns the Traffics field if non-nil, zero value otherwise.

### GetTrafficsOk

`func (o *TrafficsMapCharts) GetTrafficsOk() (*map[string]CountryTrafficChart, bool)`

GetTrafficsOk returns a tuple with the Traffics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffics

`func (o *TrafficsMapCharts) SetTraffics(v map[string]CountryTrafficChart)`

SetTraffics sets Traffics field to given value.

### HasTraffics

`func (o *TrafficsMapCharts) HasTraffics() bool`

HasTraffics returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


