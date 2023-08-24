# \RateLimitingApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RateLimitingActionsReprioritize**](RateLimitingApi.md#RateLimitingActionsReprioritize) | **Post** /domains/{domain}/rate-limit/actions/reprioritize | Change priority of Rate limiting&#39;s rules
[**RateLimitingRulesCreate**](RateLimitingApi.md#RateLimitingRulesCreate) | **Post** /domains/{domain}/rate-limit/rules | Store new Rate limiting rule
[**RateLimitingRulesDestroy**](RateLimitingApi.md#RateLimitingRulesDestroy) | **Delete** /domains/{domain}/rate-limit/rules/{id} | Delete Rate limiting&#39;s rule
[**RateLimitingRulesIndex**](RateLimitingApi.md#RateLimitingRulesIndex) | **Get** /domains/{domain}/rate-limit/rules | Get Rate limiting rules
[**RateLimitingRulesShow**](RateLimitingApi.md#RateLimitingRulesShow) | **Get** /domains/{domain}/rate-limit/rules/{id} | Get Rate limiting&#39;s rule information
[**RateLimitingRulesUpdate**](RateLimitingApi.md#RateLimitingRulesUpdate) | **Patch** /domains/{domain}/rate-limit/rules/{id} | Update the Rate limiting&#39;s rule
[**RateLimitingSettingsIndex**](RateLimitingApi.md#RateLimitingSettingsIndex) | **Get** /domains/{domain}/rate-limit/settings | Get Rate limiting settings
[**RateLimitingSettingsUpdate**](RateLimitingApi.md#RateLimitingSettingsUpdate) | **Patch** /domains/{domain}/rate-limit/settings | Update domain&#39;s Rate limiting configuration



## RateLimitingActionsReprioritize

> MessageResponse RateLimitingActionsReprioritize(ctx, domain).ReprioritizeRuleRequest(reprioritizeRuleRequest).Execute()

Change priority of Rate limiting's rules



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
    reprioritizeRuleRequest := *openapiclient.NewReprioritizeRuleRequest("RuleId_example") // ReprioritizeRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitingApi.RateLimitingActionsReprioritize(context.Background(), domain).ReprioritizeRuleRequest(reprioritizeRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.RateLimitingActionsReprioritize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateLimitingActionsReprioritize`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.RateLimitingActionsReprioritize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitingActionsReprioritizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reprioritizeRuleRequest** | [**ReprioritizeRuleRequest**](ReprioritizeRuleRequest.md) |  | 

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


## RateLimitingRulesCreate

> RateLimitRuleData RateLimitingRulesCreate(ctx, domain).RateLimitRule(rateLimitRule).Execute()

Store new Rate limiting rule

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
    rateLimitRule := *openapiclient.NewRateLimitRule("UrlPattern_example", int32(123), int32(123)) // RateLimitRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitingApi.RateLimitingRulesCreate(context.Background(), domain).RateLimitRule(rateLimitRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.RateLimitingRulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateLimitingRulesCreate`: RateLimitRuleData
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.RateLimitingRulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitingRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rateLimitRule** | [**RateLimitRule**](RateLimitRule.md) |  | 

### Return type

[**RateLimitRuleData**](RateLimitRuleData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## RateLimitingRulesDestroy

> MessageResponse RateLimitingRulesDestroy(ctx, domain, id).Execute()

Delete Rate limiting's rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitingApi.RateLimitingRulesDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.RateLimitingRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateLimitingRulesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.RateLimitingRulesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitingRulesDestroyRequest struct via the builder pattern


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


## RateLimitingRulesIndex

> RateLimitingRulesIndex200Response RateLimitingRulesIndex(ctx, domain).PerPage(perPage).OrderBy(orderBy).OrderType(orderType).Search(search).Execute()

Get Rate limiting rules

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
    perPage := int32(56) // int32 |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    orderType := "orderType_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitingApi.RateLimitingRulesIndex(context.Background(), domain).PerPage(perPage).OrderBy(orderBy).OrderType(orderType).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.RateLimitingRulesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateLimitingRulesIndex`: RateLimitingRulesIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.RateLimitingRulesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitingRulesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **orderType** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**RateLimitingRulesIndex200Response**](RateLimitingRulesIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## RateLimitingRulesShow

> RateLimitRuleData RateLimitingRulesShow(ctx, domain, id).Execute()

Get Rate limiting's rule information

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitingApi.RateLimitingRulesShow(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.RateLimitingRulesShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateLimitingRulesShow`: RateLimitRuleData
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.RateLimitingRulesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitingRulesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RateLimitRuleData**](RateLimitRuleData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## RateLimitingRulesUpdate

> RateLimitingRulesUpdate200Response RateLimitingRulesUpdate(ctx, domain, id).RateLimitRule(rateLimitRule).Execute()

Update the Rate limiting's rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    rateLimitRule := *openapiclient.NewRateLimitRule("UrlPattern_example", int32(123), int32(123)) // RateLimitRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitingApi.RateLimitingRulesUpdate(context.Background(), domain, id).RateLimitRule(rateLimitRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.RateLimitingRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateLimitingRulesUpdate`: RateLimitingRulesUpdate200Response
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.RateLimitingRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitingRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rateLimitRule** | [**RateLimitRule**](RateLimitRule.md) |  | 

### Return type

[**RateLimitingRulesUpdate200Response**](RateLimitingRulesUpdate200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## RateLimitingSettingsIndex

> RateLimitSettingsData RateLimitingSettingsIndex(ctx, domain).Execute()

Get Rate limiting settings

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
    resp, r, err := apiClient.RateLimitingApi.RateLimitingSettingsIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.RateLimitingSettingsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateLimitingSettingsIndex`: RateLimitSettingsData
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.RateLimitingSettingsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitingSettingsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RateLimitSettingsData**](RateLimitSettingsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## RateLimitingSettingsUpdate

> RateLimitingSettingsUpdate200Response RateLimitingSettingsUpdate(ctx, domain).RateLimitSettings(rateLimitSettings).Execute()

Update domain's Rate limiting configuration

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
    rateLimitSettings := *openapiclient.NewRateLimitSettings() // RateLimitSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RateLimitingApi.RateLimitingSettingsUpdate(context.Background(), domain).RateLimitSettings(rateLimitSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitingApi.RateLimitingSettingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RateLimitingSettingsUpdate`: RateLimitingSettingsUpdate200Response
    fmt.Fprintf(os.Stdout, "Response from `RateLimitingApi.RateLimitingSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRateLimitingSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rateLimitSettings** | [**RateLimitSettings**](RateLimitSettings.md) |  | 

### Return type

[**RateLimitingSettingsUpdate200Response**](RateLimitingSettingsUpdate200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

