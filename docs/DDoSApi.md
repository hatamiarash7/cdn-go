# \DDoSApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DdosReprioritize**](DDoSApi.md#DdosReprioritize) | **Post** /domains/{domain}/ddos/actions/reprioritize | Change priority of ddos rules
[**DdosRulesDestroy**](DDoSApi.md#DdosRulesDestroy) | **Delete** /domains/{domain}/ddos/rules/{id} | Delete DDoS protection rule
[**DdosRulesIndex**](DDoSApi.md#DdosRulesIndex) | **Get** /domains/{domain}/ddos/rules | Get DDoS Protection Rules
[**DdosRulesShow**](DDoSApi.md#DdosRulesShow) | **Get** /domains/{domain}/ddos/rules/{id} | Get DDoS protection&#39;s rule information
[**DdosRulesStore**](DDoSApi.md#DdosRulesStore) | **Post** /domains/{domain}/ddos/rules | Create new DDoS protection rule
[**DdosRulesUpdate**](DDoSApi.md#DdosRulesUpdate) | **Patch** /domains/{domain}/ddos/rules/{id} | Update the DDoS protection rule
[**DdosSettingsIndex**](DDoSApi.md#DdosSettingsIndex) | **Get** /domains/{domain}/ddos/settings | Get DDoS protection settings
[**DdosSettingsUpdate**](DDoSApi.md#DdosSettingsUpdate) | **Patch** /domains/{domain}/ddos/settings | Update domain&#39;s DDoS protection configuration



## DdosReprioritize

> MessageResponse DdosReprioritize(ctx, domain).ReprioritizeRuleRequest(reprioritizeRuleRequest).Execute()

Change priority of ddos rules



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
    reprioritizeRuleRequest := *openapiclient.NewReprioritizeRuleRequest("RuleId_example") // ReprioritizeRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DDoSApi.DdosReprioritize(context.Background(), domain).ReprioritizeRuleRequest(reprioritizeRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DDoSApi.DdosReprioritize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosReprioritize`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DDoSApi.DdosReprioritize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosReprioritizeRequest struct via the builder pattern


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


## DdosRulesDestroy

> MessageResponse DdosRulesDestroy(ctx, domain, id).Execute()

Delete DDoS protection rule

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
    resp, r, err := apiClient.DDoSApi.DdosRulesDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DDoSApi.DdosRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosRulesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DDoSApi.DdosRulesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosRulesDestroyRequest struct via the builder pattern


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


## DdosRulesIndex

> DdosRulesIndex200Response DdosRulesIndex(ctx, domain).PerPage(perPage).OrderBy(orderBy).OrderType(orderType).Search(search).Execute()

Get DDoS Protection Rules

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
    perPage := int32(56) // int32 |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    orderType := "orderType_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DDoSApi.DdosRulesIndex(context.Background(), domain).PerPage(perPage).OrderBy(orderBy).OrderType(orderType).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DDoSApi.DdosRulesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosRulesIndex`: DdosRulesIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `DDoSApi.DdosRulesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosRulesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **orderType** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**DdosRulesIndex200Response**](DdosRulesIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DdosRulesShow

> DdosRuleData DdosRulesShow(ctx, domain, id).Execute()

Get DDoS protection's rule information

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
    resp, r, err := apiClient.DDoSApi.DdosRulesShow(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DDoSApi.DdosRulesShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosRulesShow`: DdosRuleData
    fmt.Fprintf(os.Stdout, "Response from `DDoSApi.DdosRulesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosRulesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DdosRuleData**](DdosRuleData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DdosRulesStore

> DdosRuleResponse DdosRulesStore(ctx, domain).DdosRule(ddosRule).Execute()

Create new DDoS protection rule

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
    ddosRule := *openapiclient.NewDdosRule() // DdosRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DDoSApi.DdosRulesStore(context.Background(), domain).DdosRule(ddosRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DDoSApi.DdosRulesStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosRulesStore`: DdosRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DDoSApi.DdosRulesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosRulesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ddosRule** | [**DdosRule**](DdosRule.md) |  | 

### Return type

[**DdosRuleResponse**](DdosRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DdosRulesUpdate

> DdosRuleResponse DdosRulesUpdate(ctx, domain, id).DdosRule(ddosRule).Execute()

Update the DDoS protection rule

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
    ddosRule := *openapiclient.NewDdosRule() // DdosRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DDoSApi.DdosRulesUpdate(context.Background(), domain, id).DdosRule(ddosRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DDoSApi.DdosRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosRulesUpdate`: DdosRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `DDoSApi.DdosRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ddosRule** | [**DdosRule**](DdosRule.md) |  | 

### Return type

[**DdosRuleResponse**](DdosRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DdosSettingsIndex

> DdosSettingsData DdosSettingsIndex(ctx, domain).Execute()

Get DDoS protection settings

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
    resp, r, err := apiClient.DDoSApi.DdosSettingsIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DDoSApi.DdosSettingsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosSettingsIndex`: DdosSettingsData
    fmt.Fprintf(os.Stdout, "Response from `DDoSApi.DdosSettingsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosSettingsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DdosSettingsData**](DdosSettingsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DdosSettingsUpdate

> DdosSettingsUpdate200Response DdosSettingsUpdate(ctx, domain).DdosSettings(ddosSettings).Execute()

Update domain's DDoS protection configuration

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
    ddosSettings := *openapiclient.NewDdosSettings() // DdosSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DDoSApi.DdosSettingsUpdate(context.Background(), domain).DdosSettings(ddosSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DDoSApi.DdosSettingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DdosSettingsUpdate`: DdosSettingsUpdate200Response
    fmt.Fprintf(os.Stdout, "Response from `DDoSApi.DdosSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdosSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ddosSettings** | [**DdosSettings**](DdosSettings.md) |  | 

### Return type

[**DdosSettingsUpdate200Response**](DdosSettingsUpdate200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

