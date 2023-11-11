# TrafficsMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**[]CountryStatistics**](CountryStatistics.md) |  | [optional] 
**Charts** | Pointer to [**TrafficsMapCharts**](TrafficsMapCharts.md) |  | [optional] 
**Lists** | Pointer to [**[]CountryList**](CountryList.md) |  | [optional] 

## Methods

### NewTrafficsMap

`func NewTrafficsMap() *TrafficsMap`

NewTrafficsMap instantiates a new TrafficsMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficsMapWithDefaults

`func NewTrafficsMapWithDefaults() *TrafficsMap`

NewTrafficsMapWithDefaults instantiates a new TrafficsMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *TrafficsMap) GetStatistics() []CountryStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *TrafficsMap) GetStatisticsOk() (*[]CountryStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *TrafficsMap) SetStatistics(v []CountryStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *TrafficsMap) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *TrafficsMap) GetCharts() TrafficsMapCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *TrafficsMap) GetChartsOk() (*TrafficsMapCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *TrafficsMap) SetCharts(v TrafficsMapCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *TrafficsMap) HasCharts() bool`

HasCharts returns a boolean if a field has been set.

### GetLists

`func (o *TrafficsMap) GetLists() []CountryList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *TrafficsMap) GetListsOk() (*[]CountryList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *TrafficsMap) SetLists(v []CountryList)`

SetLists sets Lists field to given value.

### HasLists

`func (o *TrafficsMap) HasLists() bool`

HasLists returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


