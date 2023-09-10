# \FirewallApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FirewallReprioritize**](FirewallApi.md#FirewallReprioritize) | **Post** /domains/{domain}/firewall/actions/reprioritize | Change priority of firewall rules
[**FirewallRulesDestroy**](FirewallApi.md#FirewallRulesDestroy) | **Delete** /domains/{domain}/firewall/rules/{id} | Delete firewall rule
[**FirewallRulesIndex**](FirewallApi.md#FirewallRulesIndex) | **Get** /domains/{domain}/firewall/rules | Get domain&#39;s firewall rules
[**FirewallRulesShow**](FirewallApi.md#FirewallRulesShow) | **Get** /domains/{domain}/firewall/rules/{id} | Get firewall rule information
[**FirewallRulesStore**](FirewallApi.md#FirewallRulesStore) | **Post** /domains/{domain}/firewall/rules | Create new firewall rule
[**FirewallRulesUpdate**](FirewallApi.md#FirewallRulesUpdate) | **Patch** /domains/{domain}/firewall/rules/{id} | Update the firewall rule
[**FirewallSettingsIndex**](FirewallApi.md#FirewallSettingsIndex) | **Get** /domains/{domain}/firewall/settings | Get domain&#39;s firewall configuration
[**FirewallSettingsUpdate**](FirewallApi.md#FirewallSettingsUpdate) | **Patch** /domains/{domain}/firewall/settings | Update domain&#39;s firewall configuration



## FirewallReprioritize

> MessageResponse FirewallReprioritize(ctx, domain).ReprioritizeRuleRequest(reprioritizeRuleRequest).Execute()

Change priority of firewall rules



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
    resp, r, err := apiClient.FirewallApi.FirewallReprioritize(context.Background(), domain).ReprioritizeRuleRequest(reprioritizeRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.FirewallReprioritize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirewallReprioritize`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.FirewallReprioritize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirewallReprioritizeRequest struct via the builder pattern


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


## FirewallRulesDestroy

> MessageResponse FirewallRulesDestroy(ctx, domain, id).Execute()

Delete firewall rule

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
    resp, r, err := apiClient.FirewallApi.FirewallRulesDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.FirewallRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirewallRulesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.FirewallRulesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirewallRulesDestroyRequest struct via the builder pattern


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


## FirewallRulesIndex

> FirewallRulesIndex200Response FirewallRulesIndex(ctx, domain).PerPage(perPage).OrderBy(orderBy).OrderType(orderType).Search(search).Execute()

Get domain's firewall rules

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
    resp, r, err := apiClient.FirewallApi.FirewallRulesIndex(context.Background(), domain).PerPage(perPage).OrderBy(orderBy).OrderType(orderType).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.FirewallRulesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirewallRulesIndex`: FirewallRulesIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.FirewallRulesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirewallRulesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **orderType** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**FirewallRulesIndex200Response**](FirewallRulesIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## FirewallRulesShow

> FirewallRuleResponse FirewallRulesShow(ctx, domain, id).Execute()

Get firewall rule information

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
    resp, r, err := apiClient.FirewallApi.FirewallRulesShow(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.FirewallRulesShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirewallRulesShow`: FirewallRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.FirewallRulesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirewallRulesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FirewallRuleResponse**](FirewallRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## FirewallRulesStore

> FirewallRuleResponse FirewallRulesStore(ctx, domain).FirewallRule(firewallRule).Execute()

Create new firewall rule

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
    firewallRule := *openapiclient.NewFirewallRule("Name_example", "ip.geoip.country in {"IR" "TH" "US"} and ssl", "Action_example") // FirewallRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.FirewallRulesStore(context.Background(), domain).FirewallRule(firewallRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.FirewallRulesStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirewallRulesStore`: FirewallRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.FirewallRulesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirewallRulesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallRule** | [**FirewallRule**](FirewallRule.md) |  | 

### Return type

[**FirewallRuleResponse**](FirewallRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## FirewallRulesUpdate

> FirewallRuleResponse FirewallRulesUpdate(ctx, domain, id).FirewallRule(firewallRule).Execute()

Update the firewall rule

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
    firewallRule := *openapiclient.NewFirewallRule("Name_example", "ip.geoip.country in {"IR" "TH" "US"} and ssl", "Action_example") // FirewallRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.FirewallRulesUpdate(context.Background(), domain, id).FirewallRule(firewallRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.FirewallRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirewallRulesUpdate`: FirewallRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.FirewallRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirewallRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firewallRule** | [**FirewallRule**](FirewallRule.md) |  | 

### Return type

[**FirewallRuleResponse**](FirewallRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## FirewallSettingsIndex

> FirewallSettingsIndex200Response FirewallSettingsIndex(ctx, domain).Execute()

Get domain's firewall configuration

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
    resp, r, err := apiClient.FirewallApi.FirewallSettingsIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.FirewallSettingsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirewallSettingsIndex`: FirewallSettingsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.FirewallSettingsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirewallSettingsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallSettingsIndex200Response**](FirewallSettingsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## FirewallSettingsUpdate

> FirewallSettingsIndex200Response FirewallSettingsUpdate(ctx, domain).FirewallSettings(firewallSettings).Execute()

Update domain's firewall configuration

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
    firewallSettings := *openapiclient.NewFirewallSettings() // FirewallSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.FirewallSettingsUpdate(context.Background(), domain).FirewallSettings(firewallSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.FirewallSettingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FirewallSettingsUpdate`: FirewallSettingsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.FirewallSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiFirewallSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallSettings** | [**FirewallSettings**](FirewallSettings.md) |  | 

### Return type

[**FirewallSettingsIndex200Response**](FirewallSettingsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

