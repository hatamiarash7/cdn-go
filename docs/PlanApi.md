# \PlanApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsPlans**](PlanApi.md#DomainsPlans) | **Get** /domains/{domain}/plans | Get the list of feature defintions for plans based on different sets
[**DomainsPlansUpdate**](PlanApi.md#DomainsPlansUpdate) | **Put** /domains/{domain}/plan | Update the domain&#39;s plan
[**DomainsPlansUsages**](PlanApi.md#DomainsPlansUsages) | **Get** /domains/{domain}/plan/usages | Get usages based on features and an estimated cost
[**DomainsPlansViolations**](PlanApi.md#DomainsPlansViolations) | **Get** /domains/{domain}/plan/violations | Get violations based on plans
[**PlansIndex**](PlanApi.md#PlansIndex) | **Get** /plans | Get the list of feature defintions for plans based on different sets



## DomainsPlans

> PlanResponse DomainsPlans(ctx, domain).IgnoredPlans(ignoredPlans).Execute()

Get the list of feature defintions for plans based on different sets

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
    ignoredPlans := "0,1" // string | Comma separaterd plan levels to ignore (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.DomainsPlans(context.Background(), domain).IgnoredPlans(ignoredPlans).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.DomainsPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsPlans`: PlanResponse
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.DomainsPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ignoredPlans** | **string** | Comma separaterd plan levels to ignore | 

### Return type

[**PlanResponse**](PlanResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsPlansUpdate

> MessageResponse DomainsPlansUpdate(ctx, domain).PlanUpdate(planUpdate).Execute()

Update the domain's plan

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
    planUpdate := *openapiclient.NewPlanUpdate(int32(123)) // PlanUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.DomainsPlansUpdate(context.Background(), domain).PlanUpdate(planUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.DomainsPlansUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsPlansUpdate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.DomainsPlansUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsPlansUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **planUpdate** | [**PlanUpdate**](PlanUpdate.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsPlansUsages

> DomainsPlansUsages200Response DomainsPlansUsages(ctx, domain).TargetPlan(targetPlan).Execute()

Get usages based on features and an estimated cost

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
    targetPlan := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.DomainsPlansUsages(context.Background(), domain).TargetPlan(targetPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.DomainsPlansUsages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsPlansUsages`: DomainsPlansUsages200Response
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.DomainsPlansUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsPlansUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetPlan** | **int32** |  | 

### Return type

[**DomainsPlansUsages200Response**](DomainsPlansUsages200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsPlansViolations

> DomainsPlansViolations200Response DomainsPlansViolations(ctx, domain).Execute()

Get violations based on plans

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.DomainsPlansViolations(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.DomainsPlansViolations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsPlansViolations`: DomainsPlansViolations200Response
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.DomainsPlansViolations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsPlansViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainsPlansViolations200Response**](DomainsPlansViolations200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## PlansIndex

> PlanResponse PlansIndex(ctx).Domain(domain).IgnoredPlans(ignoredPlans).Execute()

Get the list of feature defintions for plans based on different sets

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
    domain := *openapiclient.NewPlansIndexDomainParameter() // PlansIndexDomainParameter | Domain name or id (optional)
    ignoredPlans := "0,1" // string | Comma separaterd plan levels to ignore (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.PlansIndex(context.Background()).Domain(domain).IgnoredPlans(ignoredPlans).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.PlansIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlansIndex`: PlanResponse
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.PlansIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlansIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | [**PlansIndexDomainParameter**](PlansIndexDomainParameter.md) | Domain name or id | 
 **ignoredPlans** | **string** | Comma separaterd plan levels to ignore | 

### Return type

[**PlanResponse**](PlanResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

