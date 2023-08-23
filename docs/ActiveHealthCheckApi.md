# \ActiveHealthCheckApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveHealthCheckDestroy**](ActiveHealthCheckApi.md#ActiveHealthCheckDestroy) | **Delete** /domains/{domain}/health-checks/{healthcheck} | Delete healthcheck
[**ActiveHealthCheckIndex**](ActiveHealthCheckApi.md#ActiveHealthCheckIndex) | **Get** /domains/{domain}/health-checks | Get Defined HealthCheck
[**ActiveHealthCheckReportsDetails**](ActiveHealthCheckApi.md#ActiveHealthCheckReportsDetails) | **Get** /domains/{domain}/health-checks/reports/details | Get reports of a single healthcheck monitoring
[**ActiveHealthCheckReportsSummary**](ActiveHealthCheckApi.md#ActiveHealthCheckReportsSummary) | **Get** /domains/{domain}/health-checks/reports/summary | Get summary reports of a single healthcheck monitoring
[**ActiveHealthCheckShow**](ActiveHealthCheckApi.md#ActiveHealthCheckShow) | **Get** /domains/{domain}/health-checks/{healthcheck} | Get a single healthcheck
[**ActiveHealthCheckStore**](ActiveHealthCheckApi.md#ActiveHealthCheckStore) | **Post** /domains/{domain}/health-checks | Store a new HealthCheck
[**ActiveHealthCheckUpdate**](ActiveHealthCheckApi.md#ActiveHealthCheckUpdate) | **Patch** /domains/{domain}/health-checks/{healthcheck} | Update Health check
[**HealthChecksZonesIndex**](ActiveHealthCheckApi.md#HealthChecksZonesIndex) | **Get** /health-checks/zones | Get list of all health-check zones



## ActiveHealthCheckDestroy

> MessageResponse ActiveHealthCheckDestroy(ctx, domain, healthcheck).Execute()

Delete healthcheck

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn/go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    healthcheck := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveHealthCheckApi.ActiveHealthCheckDestroy(context.Background(), domain, healthcheck).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveHealthCheckApi.ActiveHealthCheckDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActiveHealthCheckDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ActiveHealthCheckApi.ActiveHealthCheckDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**healthcheck** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveHealthCheckDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ActiveHealthCheckIndex

> ActiveHealthCheckIndex200Response ActiveHealthCheckIndex(ctx, domain).Execute()

Get Defined HealthCheck

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn/go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveHealthCheckApi.ActiveHealthCheckIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveHealthCheckApi.ActiveHealthCheckIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActiveHealthCheckIndex`: ActiveHealthCheckIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `ActiveHealthCheckApi.ActiveHealthCheckIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveHealthCheckIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActiveHealthCheckIndex200Response**](ActiveHealthCheckIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ActiveHealthCheckReportsDetails

> ActiveHealthCheckReportsDetails200Response ActiveHealthCheckReportsDetails(ctx, domain).Name(name).Upstream(upstream).Type_(type_).Period(period).Since(since).Until(until).Direction(direction).PerPage(perPage).Page(page).Execute()

Get reports of a single healthcheck monitoring

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn/go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    name := "name_example" // string | Name of the health-check
    upstream := "upstream_example" // string | Upstream of the health-check monitoring
    type_ := "type__example" // string | Type of reports for the health check (optional) (default to "all")
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    direction := "direction_example" // string | Set the direction of sorting (optional)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveHealthCheckApi.ActiveHealthCheckReportsDetails(context.Background(), domain).Name(name).Upstream(upstream).Type_(type_).Period(period).Since(since).Until(until).Direction(direction).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveHealthCheckApi.ActiveHealthCheckReportsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActiveHealthCheckReportsDetails`: ActiveHealthCheckReportsDetails200Response
    fmt.Fprintf(os.Stdout, "Response from `ActiveHealthCheckApi.ActiveHealthCheckReportsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveHealthCheckReportsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the health-check | 
 **upstream** | **string** | Upstream of the health-check monitoring | 
 **type_** | **string** | Type of reports for the health check | [default to &quot;all&quot;]
 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **direction** | **string** | Set the direction of sorting | 
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**ActiveHealthCheckReportsDetails200Response**](ActiveHealthCheckReportsDetails200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ActiveHealthCheckReportsSummary

> ActiveHealthCheckReportsSummary200Response ActiveHealthCheckReportsSummary(ctx, domain).Name(name).Upstream(upstream).Period(period).Since(since).Until(until).Direction(direction).Execute()

Get summary reports of a single healthcheck monitoring

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "git.arvancloud.ir/arvancloud/cdn/go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    name := "name_example" // string | Name of the health-check
    upstream := "upstream_example" // string | Upstream of the health-check monitoring
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    direction := "direction_example" // string | Set the direction of sorting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveHealthCheckApi.ActiveHealthCheckReportsSummary(context.Background(), domain).Name(name).Upstream(upstream).Period(period).Since(since).Until(until).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveHealthCheckApi.ActiveHealthCheckReportsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActiveHealthCheckReportsSummary`: ActiveHealthCheckReportsSummary200Response
    fmt.Fprintf(os.Stdout, "Response from `ActiveHealthCheckApi.ActiveHealthCheckReportsSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveHealthCheckReportsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the health-check | 
 **upstream** | **string** | Upstream of the health-check monitoring | 
 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **direction** | **string** | Set the direction of sorting | 

### Return type

[**ActiveHealthCheckReportsSummary200Response**](ActiveHealthCheckReportsSummary200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ActiveHealthCheckShow

> HealthCheckResponse ActiveHealthCheckShow(ctx, domain, healthcheck).Execute()

Get a single healthcheck

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn/go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    healthcheck := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveHealthCheckApi.ActiveHealthCheckShow(context.Background(), domain, healthcheck).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveHealthCheckApi.ActiveHealthCheckShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActiveHealthCheckShow`: HealthCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ActiveHealthCheckApi.ActiveHealthCheckShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**healthcheck** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveHealthCheckShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HealthCheckResponse**](HealthCheckResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ActiveHealthCheckStore

> HealthCheckResponse ActiveHealthCheckStore(ctx, domain).HealthCheck(healthCheck).Execute()

Store a new HealthCheck

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn/go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    healthCheck := *openapiclient.NewHealthCheck() // HealthCheck |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveHealthCheckApi.ActiveHealthCheckStore(context.Background(), domain).HealthCheck(healthCheck).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveHealthCheckApi.ActiveHealthCheckStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActiveHealthCheckStore`: HealthCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ActiveHealthCheckApi.ActiveHealthCheckStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveHealthCheckStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **healthCheck** | [**HealthCheck**](HealthCheck.md) |  | 

### Return type

[**HealthCheckResponse**](HealthCheckResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ActiveHealthCheckUpdate

> HealthCheckResponse ActiveHealthCheckUpdate(ctx, domain, healthcheck).HealthCheck(healthCheck).Execute()

Update Health check

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn/go-sdk"
)

func main() {
    domain := "example.com" // string | Domain name
    healthcheck := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    healthCheck := *openapiclient.NewHealthCheck() // HealthCheck |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveHealthCheckApi.ActiveHealthCheckUpdate(context.Background(), domain, healthcheck).HealthCheck(healthCheck).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveHealthCheckApi.ActiveHealthCheckUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActiveHealthCheckUpdate`: HealthCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ActiveHealthCheckApi.ActiveHealthCheckUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**healthcheck** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActiveHealthCheckUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **healthCheck** | [**HealthCheck**](HealthCheck.md) |  | 

### Return type

[**HealthCheckResponse**](HealthCheckResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## HealthChecksZonesIndex

> HealthChecksZonesIndex200Response HealthChecksZonesIndex(ctx).Execute()

Get list of all health-check zones

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "git.arvancloud.ir/arvancloud/cdn/go-sdk"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveHealthCheckApi.HealthChecksZonesIndex(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveHealthCheckApi.HealthChecksZonesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthChecksZonesIndex`: HealthChecksZonesIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `ActiveHealthCheckApi.HealthChecksZonesIndex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthChecksZonesIndexRequest struct via the builder pattern


### Return type

[**HealthChecksZonesIndex200Response**](HealthChecksZonesIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

