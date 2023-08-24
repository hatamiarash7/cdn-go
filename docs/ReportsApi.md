# \ReportsApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkReportsTrafficsTotal**](ReportsApi.md#BulkReportsTrafficsTotal) | **Post** /bulk/reports/traffics | Get traffic report for multiple domains
[**BulkReportsVisitorsTotal**](ReportsApi.md#BulkReportsVisitorsTotal) | **Post** /bulk/reports/visitors | Get visitor report for multiple domains
[**ReportsAttacksAttackers**](ReportsApi.md#ReportsAttacksAttackers) | **Get** /domains/{domain}/reports/attacks/attackers | Get list of attackers information
[**ReportsAttacksIndex**](ReportsApi.md#ReportsAttacksIndex) | **Get** /domains/{domain}/reports/attacks/list | Get list of attacks details
[**ReportsAttacksMap**](ReportsApi.md#ReportsAttacksMap) | **Get** /domains/{domain}/reports/attacks/map | Get geo-map of attacks
[**ReportsAttacksShow**](ReportsApi.md#ReportsAttacksShow) | **Get** /domains/{domain}/reports/attacks | Get report of attacks
[**ReportsAttacksUri**](ReportsApi.md#ReportsAttacksUri) | **Get** /domains/{domain}/reports/attacks/uri | Get list of URLs under attack
[**ReportsDnsGeo**](ReportsApi.md#ReportsDnsGeo) | **Get** /domains/{domain}/reports/dns-geo | Get DNS requests as geo-map
[**ReportsDnsRequests**](ReportsApi.md#ReportsDnsRequests) | **Get** /domains/{domain}/reports/dns-requests | Get response time report
[**ReportsErrorLogDetails**](ReportsApi.md#ReportsErrorLogDetails) | **Get** /domains/{domain}/reports/error-log-details | Get detail of an error
[**ReportsErrorLogs**](ReportsApi.md#ReportsErrorLogs) | **Get** /domains/{domain}/reports/error-logs | Get list of errors
[**ReportsErrorLogsChart**](ReportsApi.md#ReportsErrorLogsChart) | **Get** /domains/{domain}/reports/error-logs/chart | Get chart view of errors
[**ReportsResponseTimeIndex**](ReportsApi.md#ReportsResponseTimeIndex) | **Get** /domains/{domain}/reports/response-time | Get response time report
[**ReportsStatusIndex**](ReportsApi.md#ReportsStatusIndex) | **Get** /domains/{domain}/reports/status | Get status codes pie chart
[**ReportsStatusSummary**](ReportsApi.md#ReportsStatusSummary) | **Get** /domains/{domain}/reports/status/summary | Get an overview of status codes pie chart
[**ReportsTrafficsMap**](ReportsApi.md#ReportsTrafficsMap) | **Get** /domains/{domain}/reports/traffics/map | Get traffic as geo-map
[**ReportsTrafficsSaved**](ReportsApi.md#ReportsTrafficsSaved) | **Get** /domains/{domain}/reports/traffics/saved | Get traffic saved to total pie chart
[**ReportsTrafficsTotal**](ReportsApi.md#ReportsTrafficsTotal) | **Get** /domains/{domain}/reports/traffics | Get traffic report for domain
[**ReportsTransportLayerProxiesTraffics**](ReportsApi.md#ReportsTransportLayerProxiesTraffics) | **Get** /domains/{domain}/reports/transport-layer-proxies/{transportLayerProxyId}/traffics | Get traffic report for transport layer proxy
[**ReportsVisitorsHighRequestIps**](ReportsApi.md#ReportsVisitorsHighRequestIps) | **Get** /domains/{domain}/reports/high-request-ips | Get report of IPs with highest number of requests
[**ReportsVisitorsIndex**](ReportsApi.md#ReportsVisitorsIndex) | **Get** /domains/{domain}/reports/visitors | Get report of visitors for domain



## BulkReportsTrafficsTotal

> BulkTrafficReportData BulkReportsTrafficsTotal(ctx).Period(period).Since(since).Until(until).BulkReportsTrafficsTotalRequest(bulkReportsTrafficsTotalRequest).Execute()

Get traffic report for multiple domains

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    bulkReportsTrafficsTotalRequest := *openapiclient.NewBulkReportsTrafficsTotalRequest() // BulkReportsTrafficsTotalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.BulkReportsTrafficsTotal(context.Background()).Period(period).Since(since).Until(until).BulkReportsTrafficsTotalRequest(bulkReportsTrafficsTotalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.BulkReportsTrafficsTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkReportsTrafficsTotal`: BulkTrafficReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.BulkReportsTrafficsTotal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkReportsTrafficsTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **bulkReportsTrafficsTotalRequest** | [**BulkReportsTrafficsTotalRequest**](BulkReportsTrafficsTotalRequest.md) |  | 

### Return type

[**BulkTrafficReportData**](BulkTrafficReportData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## BulkReportsVisitorsTotal

> BulkVisitorReportData BulkReportsVisitorsTotal(ctx).Period(period).Since(since).Until(until).BulkReportsVisitorsTotalRequest(bulkReportsVisitorsTotalRequest).Execute()

Get visitor report for multiple domains

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    bulkReportsVisitorsTotalRequest := *openapiclient.NewBulkReportsVisitorsTotalRequest() // BulkReportsVisitorsTotalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.BulkReportsVisitorsTotal(context.Background()).Period(period).Since(since).Until(until).BulkReportsVisitorsTotalRequest(bulkReportsVisitorsTotalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.BulkReportsVisitorsTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkReportsVisitorsTotal`: BulkVisitorReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.BulkReportsVisitorsTotal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkReportsVisitorsTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **bulkReportsVisitorsTotalRequest** | [**BulkReportsVisitorsTotalRequest**](BulkReportsVisitorsTotalRequest.md) |  | 

### Return type

[**BulkVisitorReportData**](BulkVisitorReportData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsAttacksAttackers

> ReportsAttacksAttackers200Response ReportsAttacksAttackers(ctx, domain).Period(period).Execute()

Get list of attackers information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsAttacksAttackers(context.Background(), domain).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsAttacksAttackers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksAttackers`: ReportsAttacksAttackers200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsAttacksAttackers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksAttackersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 

### Return type

[**ReportsAttacksAttackers200Response**](ReportsAttacksAttackers200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsAttacksIndex

> ReportsAttacksIndex200Response ReportsAttacksIndex(ctx, domain).Period(period).PerPage(perPage).Page(page).Execute()

Get list of attacks details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsAttacksIndex(context.Background(), domain).Period(period).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsAttacksIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksIndex`: ReportsAttacksIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsAttacksIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**ReportsAttacksIndex200Response**](ReportsAttacksIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsAttacksMap

> AttackReportMapData ReportsAttacksMap(ctx, domain).Period(period).Execute()

Get geo-map of attacks

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsAttacksMap(context.Background(), domain).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsAttacksMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksMap`: AttackReportMapData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsAttacksMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 

### Return type

[**AttackReportMapData**](AttackReportMapData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsAttacksShow

> ReportsAttacksShow200Response ReportsAttacksShow(ctx, domain).Period(period).Execute()

Get report of attacks

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsAttacksShow(context.Background(), domain).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsAttacksShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksShow`: ReportsAttacksShow200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsAttacksShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 

### Return type

[**ReportsAttacksShow200Response**](ReportsAttacksShow200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsAttacksUri

> AttackReportUriData ReportsAttacksUri(ctx, domain).Period(period).Execute()

Get list of URLs under attack

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsAttacksUri(context.Background(), domain).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsAttacksUri``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksUri`: AttackReportUriData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsAttacksUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 

### Return type

[**AttackReportUriData**](AttackReportUriData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsDnsGeo

> DnsGeoReportData ReportsDnsGeo(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get DNS requests as geo-map

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsDnsGeo(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsDnsGeo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsDnsGeo`: DnsGeoReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsDnsGeo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsDnsGeoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**DnsGeoReportData**](DnsGeoReportData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsDnsRequests

> DnsRequestReportData ReportsDnsRequests(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get response time report

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsDnsRequests(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsDnsRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsDnsRequests`: DnsRequestReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsDnsRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsDnsRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**DnsRequestReportData**](DnsRequestReportData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsErrorLogDetails

> ReportsErrorLogDetails200Response ReportsErrorLogDetails(ctx, domain).Period(period).Since(since).Until(until).Error_(error_).Execute()

Get detail of an error

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    error_ := "error__example" // string | Error message to search for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsErrorLogDetails(context.Background(), domain).Period(period).Since(since).Until(until).Error_(error_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsErrorLogDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsErrorLogDetails`: ReportsErrorLogDetails200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsErrorLogDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsErrorLogDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **error_** | **string** | Error message to search for | 

### Return type

[**ReportsErrorLogDetails200Response**](ReportsErrorLogDetails200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsErrorLogs

> ErrorLogsData ReportsErrorLogs(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get list of errors

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsErrorLogs(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsErrorLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsErrorLogs`: ErrorLogsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsErrorLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsErrorLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**ErrorLogsData**](ErrorLogsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsErrorLogsChart

> ErrorLogChartData ReportsErrorLogsChart(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get chart view of errors

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsErrorLogsChart(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsErrorLogsChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsErrorLogsChart`: ErrorLogChartData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsErrorLogsChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsErrorLogsChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**ErrorLogChartData**](ErrorLogChartData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsResponseTimeIndex

> ResponseTimeData ReportsResponseTimeIndex(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get response time report

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsResponseTimeIndex(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsResponseTimeIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsResponseTimeIndex`: ResponseTimeData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsResponseTimeIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsResponseTimeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**ResponseTimeData**](ResponseTimeData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsStatusIndex

> StatusCodeReportData ReportsStatusIndex(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get status codes pie chart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsStatusIndex(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsStatusIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsStatusIndex`: StatusCodeReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsStatusIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsStatusIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**StatusCodeReportData**](StatusCodeReportData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsStatusSummary

> StatusCodeSummaryData ReportsStatusSummary(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get an overview of status codes pie chart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsStatusSummary(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsStatusSummary`: StatusCodeSummaryData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsStatusSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsStatusSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**StatusCodeSummaryData**](StatusCodeSummaryData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsTrafficsMap

> MapTrafficsData ReportsTrafficsMap(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get traffic as geo-map

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsTrafficsMap(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsTrafficsMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsTrafficsMap`: MapTrafficsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsTrafficsMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsTrafficsMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**MapTrafficsData**](MapTrafficsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsTrafficsSaved

> SavedTrafficsData ReportsTrafficsSaved(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get traffic saved to total pie chart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsTrafficsSaved(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsTrafficsSaved``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsTrafficsSaved`: SavedTrafficsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsTrafficsSaved`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsTrafficsSavedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**SavedTrafficsData**](SavedTrafficsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsTrafficsTotal

> TrafficsData ReportsTrafficsTotal(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get traffic report for domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsTrafficsTotal(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsTrafficsTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsTrafficsTotal`: TrafficsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsTrafficsTotal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsTrafficsTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**TrafficsData**](TrafficsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsTransportLayerProxiesTraffics

> TransportLayerProxyTrafficsData ReportsTransportLayerProxiesTraffics(ctx, domain, transportLayerProxyId).Period(period).Since(since).Until(until).Execute()

Get traffic report for transport layer proxy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    transportLayerProxyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport layer proxy id
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsTransportLayerProxiesTraffics(context.Background(), domain, transportLayerProxyId).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsTransportLayerProxiesTraffics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsTransportLayerProxiesTraffics`: TransportLayerProxyTrafficsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsTransportLayerProxiesTraffics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**transportLayerProxyId** | **string** | Transport layer proxy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsTransportLayerProxiesTrafficsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**TransportLayerProxyTrafficsData**](TransportLayerProxyTrafficsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsVisitorsHighRequestIps

> ReportsVisitorsHighRequestIps200Response ReportsVisitorsHighRequestIps(ctx, domain).Period(period).Since(since).Until(until).PerPage(perPage).Page(page).Execute()

Get report of IPs with highest number of requests

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsVisitorsHighRequestIps(context.Background(), domain).Period(period).Since(since).Until(until).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsVisitorsHighRequestIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsVisitorsHighRequestIps`: ReportsVisitorsHighRequestIps200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsVisitorsHighRequestIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsVisitorsHighRequestIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**ReportsVisitorsHighRequestIps200Response**](ReportsVisitorsHighRequestIps200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ReportsVisitorsIndex

> VisitorsData ReportsVisitorsIndex(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get report of visitors for domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn-go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsApi.ReportsVisitorsIndex(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsVisitorsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsVisitorsIndex`: VisitorsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsVisitorsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsVisitorsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**VisitorsData**](VisitorsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

