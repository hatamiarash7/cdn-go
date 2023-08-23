# \PageRuleApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PageRulesDestroy**](PageRuleApi.md#PageRulesDestroy) | **Delete** /domains/{domain}/page-rules/{id} | Delete the page-rule
[**PageRulesDiffShow**](PageRuleApi.md#PageRulesDiffShow) | **Get** /domains/{domain}/page-rules/{id}/diff | Get the page-rule&#39;s exceptions
[**PageRulesDiffUpdate**](PageRuleApi.md#PageRulesDiffUpdate) | **Patch** /domains/{domain}/page-rules/{id}/diff | Update the page-rule&#39;s exceptions
[**PageRulesIndex**](PageRuleApi.md#PageRulesIndex) | **Get** /domains/{domain}/page-rules | Get list of page-rules
[**PageRulesPurge**](PageRuleApi.md#PageRulesPurge) | **Delete** /domains/{domain}/page-rules/{id}/purge | Purge the page-rule
[**PageRulesShow**](PageRuleApi.md#PageRulesShow) | **Get** /domains/{domain}/page-rules/{id} | Get the page-rule&#39;s information
[**PageRulesStatusUpdate**](PageRuleApi.md#PageRulesStatusUpdate) | **Patch** /domains/{domain}/page-rules/{id} | Toggle status of the page-rule
[**PageRulesStore**](PageRuleApi.md#PageRulesStore) | **Post** /domains/{domain}/page-rules | Create new page-rule
[**PageRulesUpdate**](PageRuleApi.md#PageRulesUpdate) | **Put** /domains/{domain}/page-rules/{id} | Update the page-rule



## PageRulesDestroy

> MessageResponse PageRulesDestroy(ctx, domain, id).Execute()

Delete the page-rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesDestroyRequest struct via the builder pattern


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


## PageRulesDiffShow

> PageRuleDiffData PageRulesDiffShow(ctx, domain, id).Execute()

Get the page-rule's exceptions

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesDiffShow(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesDiffShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesDiffShow`: PageRuleDiffData
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesDiffShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesDiffShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PageRuleDiffData**](PageRuleDiffData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## PageRulesDiffUpdate

> PageRulesDiffUpdate200Response PageRulesDiffUpdate(ctx, domain, id).PageRule(pageRule).Execute()

Update the page-rule's exceptions

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    pageRule := *openapiclient.NewPageRule() // PageRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesDiffUpdate(context.Background(), domain, id).PageRule(pageRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesDiffUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesDiffUpdate`: PageRulesDiffUpdate200Response
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesDiffUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesDiffUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRule** | [**PageRule**](PageRule.md) |  | 

### Return type

[**PageRulesDiffUpdate200Response**](PageRulesDiffUpdate200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## PageRulesIndex

> PageRulesIndex200Response PageRulesIndex(ctx, domain).Search(search).PerPage(perPage).Page(page).Order(order).Execute()

Get list of page-rules

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
    search := "search_example" // string | Search term (optional)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)
    order := "order_example" // string | Sort page rules in ascending or descending order base on seq (optional) (default to "desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesIndex(context.Background(), domain).Search(search).PerPage(perPage).Page(page).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesIndex`: PageRulesIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Search term | 
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]
 **order** | **string** | Sort page rules in ascending or descending order base on seq | [default to &quot;desc&quot;]

### Return type

[**PageRulesIndex200Response**](PageRulesIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## PageRulesPurge

> MessageResponse PageRulesPurge(ctx, domain, id).Execute()

Purge the page-rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesPurge(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesPurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesPurge`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesPurgeRequest struct via the builder pattern


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


## PageRulesShow

> PageRuleData PageRulesShow(ctx, domain, id).Execute()

Get the page-rule's information

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesShow(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesShow`: PageRuleData
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PageRuleData**](PageRuleData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## PageRulesStatusUpdate

> MessageResponse PageRulesStatusUpdate(ctx, domain, id).UpdateBooleanStatus(updateBooleanStatus).Execute()

Toggle status of the page-rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    updateBooleanStatus := *openapiclient.NewUpdateBooleanStatus() // UpdateBooleanStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesStatusUpdate(context.Background(), domain, id).UpdateBooleanStatus(updateBooleanStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesStatusUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesStatusUpdate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesStatusUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesStatusUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateBooleanStatus** | [**UpdateBooleanStatus**](UpdateBooleanStatus.md) |  | 

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


## PageRulesStore

> PageRuleResponse PageRulesStore(ctx, domain).PageRule(pageRule).Execute()

Create new page-rule

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
    pageRule := *openapiclient.NewPageRule() // PageRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesStore(context.Background(), domain).PageRule(pageRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesStore`: PageRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRule** | [**PageRule**](PageRule.md) |  | 

### Return type

[**PageRuleResponse**](PageRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## PageRulesUpdate

> PageRuleResponse PageRulesUpdate(ctx, domain, id).PageRule(pageRule).Execute()

Update the page-rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    pageRule := *openapiclient.NewPageRule() // PageRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PageRuleApi.PageRulesUpdate(context.Background(), domain, id).PageRule(pageRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageRuleApi.PageRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageRulesUpdate`: PageRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `PageRuleApi.PageRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageRule** | [**PageRule**](PageRule.md) |  | 

### Return type

[**PageRuleResponse**](PageRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

