# \WAFApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalWafIndex**](WAFApi.md#GlobalWafIndex) | **Get** /waf | Get WAF presets
[**GlobalWafShowPackage**](WAFApi.md#GlobalWafShowPackage) | **Get** /waf/packages/{packageId} | Get WAF package details
[**WafPackageReprioritize**](WAFApi.md#WafPackageReprioritize) | **Post** /domains/{domain}/waf/actions/reprioritize-package | Change priority of WAF packages
[**WafPackagesDestroy**](WAFApi.md#WafPackagesDestroy) | **Delete** /domains/{domain}/waf/packages/{id} | Delete WAF package from domain
[**WafPackagesIndex**](WAFApi.md#WafPackagesIndex) | **Get** /domains/{domain}/waf/packages | Get WAF packages
[**WafPackagesShow**](WAFApi.md#WafPackagesShow) | **Get** /domains/{domain}/waf/packages/{id} | Get WAF package information
[**WafPackagesStore**](WAFApi.md#WafPackagesStore) | **Post** /domains/{domain}/waf/packages | Add new WAF package to domain
[**WafPackagesUpdate**](WAFApi.md#WafPackagesUpdate) | **Patch** /domains/{domain}/waf/packages/{id} | Update the WAF package
[**WafReconfigure**](WAFApi.md#WafReconfigure) | **Post** /domains/{domain}/waf/actions/reconfigure | Reconfigure WAF module using a preset
[**WafReprioritize**](WAFApi.md#WafReprioritize) | **Post** /domains/{domain}/waf/actions/reprioritize | Change priority of WAF rules
[**WafRulesDestroy**](WAFApi.md#WafRulesDestroy) | **Delete** /domains/{domain}/waf/rules/{id} | Delete WAF rule
[**WafRulesIndex**](WAFApi.md#WafRulesIndex) | **Get** /domains/{domain}/waf/rules | Get WAF Rules
[**WafRulesShow**](WAFApi.md#WafRulesShow) | **Get** /domains/{domain}/waf/rules/{id} | Get WAF rule information
[**WafRulesStore**](WAFApi.md#WafRulesStore) | **Post** /domains/{domain}/waf/rules | Create new WAF rule
[**WafRulesUpdate**](WAFApi.md#WafRulesUpdate) | **Patch** /domains/{domain}/waf/rules/{id} | Update the WAF rule
[**WafSettingsIndex**](WAFApi.md#WafSettingsIndex) | **Get** /domains/{domain}/waf/settings | Get WAF configuration
[**WafSettingsUpdate**](WAFApi.md#WafSettingsUpdate) | **Patch** /domains/{domain}/waf/settings | Configure WAF module of the domain



## GlobalWafIndex

> WafPresetsData GlobalWafIndex(ctx).Execute()

Get WAF presets

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.GlobalWafIndex(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.GlobalWafIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GlobalWafIndex`: WafPresetsData
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.GlobalWafIndex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalWafIndexRequest struct via the builder pattern


### Return type

[**WafPresetsData**](WafPresetsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## GlobalWafShowPackage

> WafPackageDetailsData GlobalWafShowPackage(ctx, packageId).Execute()

Get WAF package details

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
    packageId := "packageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.GlobalWafShowPackage(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.GlobalWafShowPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GlobalWafShowPackage`: WafPackageDetailsData
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.GlobalWafShowPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalWafShowPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WafPackageDetailsData**](WafPackageDetailsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafPackageReprioritize

> MessageResponse WafPackageReprioritize(ctx, domain).WafReprioritize(wafReprioritize).Execute()

Change priority of WAF packages



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
    wafReprioritize := *openapiclient.NewWafReprioritize("PackageId_example") // WafReprioritize |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafPackageReprioritize(context.Background(), domain).WafReprioritize(wafReprioritize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafPackageReprioritize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafPackageReprioritize`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafPackageReprioritize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafPackageReprioritizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wafReprioritize** | [**WafReprioritize**](WafReprioritize.md) |  | 

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


## WafPackagesDestroy

> MessageResponse WafPackagesDestroy(ctx, domain, id).Execute()

Delete WAF package from domain

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafPackagesDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafPackagesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafPackagesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafPackagesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafPackagesDestroyRequest struct via the builder pattern


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


## WafPackagesIndex

> DomainWafPackagesData WafPackagesIndex(ctx, domain).Available(available).Execute()

Get WAF packages



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
    available := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafPackagesIndex(context.Background(), domain).Available(available).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafPackagesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafPackagesIndex`: DomainWafPackagesData
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafPackagesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafPackagesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **available** | **bool** |  | [default to false]

### Return type

[**DomainWafPackagesData**](DomainWafPackagesData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafPackagesShow

> DomainWafPackageDetailsData WafPackagesShow(ctx, domain, id).Execute()

Get WAF package information

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafPackagesShow(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafPackagesShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafPackagesShow`: DomainWafPackageDetailsData
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafPackagesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafPackagesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DomainWafPackageDetailsData**](DomainWafPackageDetailsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafPackagesStore

> WafPackagesStore200Response WafPackagesStore(ctx, domain).DomainWafPackageStore(domainWafPackageStore).Execute()

Add new WAF package to domain

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
    domainWafPackageStore := *openapiclient.NewDomainWafPackageStore("Id_example") // DomainWafPackageStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafPackagesStore(context.Background(), domain).DomainWafPackageStore(domainWafPackageStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafPackagesStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafPackagesStore`: WafPackagesStore200Response
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafPackagesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafPackagesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainWafPackageStore** | [**DomainWafPackageStore**](DomainWafPackageStore.md) |  | 

### Return type

[**WafPackagesStore200Response**](WafPackagesStore200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafPackagesUpdate

> WafPackagesUpdate200Response WafPackagesUpdate(ctx, domain, id).DomainWafPackage(domainWafPackage).Execute()

Update the WAF package

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
    id := "id_example" // string | 
    domainWafPackage := *openapiclient.NewDomainWafPackage() // DomainWafPackage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafPackagesUpdate(context.Background(), domain, id).DomainWafPackage(domainWafPackage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafPackagesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafPackagesUpdate`: WafPackagesUpdate200Response
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafPackagesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafPackagesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **domainWafPackage** | [**DomainWafPackage**](DomainWafPackage.md) |  | 

### Return type

[**WafPackagesUpdate200Response**](WafPackagesUpdate200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafReconfigure

> MessageResponse WafReconfigure(ctx, domain).WafReconfigure(wafReconfigure).Execute()

Reconfigure WAF module using a preset



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
    wafReconfigure := *openapiclient.NewWafReconfigure() // WafReconfigure |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafReconfigure(context.Background(), domain).WafReconfigure(wafReconfigure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafReconfigure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafReconfigure`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafReconfigure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafReconfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wafReconfigure** | [**WafReconfigure**](WafReconfigure.md) |  | 

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


## WafReprioritize

> MessageResponse WafReprioritize(ctx, domain).ReprioritizeRuleRequest(reprioritizeRuleRequest).Execute()

Change priority of WAF rules



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
    resp, r, err := apiClient.WAFApi.WafReprioritize(context.Background(), domain).ReprioritizeRuleRequest(reprioritizeRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafReprioritize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafReprioritize`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafReprioritize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafReprioritizeRequest struct via the builder pattern


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


## WafRulesDestroy

> MessageResponse WafRulesDestroy(ctx, domain, id).Execute()

Delete WAF rule

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
    resp, r, err := apiClient.WAFApi.WafRulesDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafRulesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafRulesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafRulesDestroyRequest struct via the builder pattern


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


## WafRulesIndex

> WafRulesIndex200Response WafRulesIndex(ctx, domain).PerPage(perPage).OrderBy(orderBy).OrderType(orderType).Search(search).Execute()

Get WAF Rules

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
    resp, r, err := apiClient.WAFApi.WafRulesIndex(context.Background(), domain).PerPage(perPage).OrderBy(orderBy).OrderType(orderType).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafRulesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafRulesIndex`: WafRulesIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafRulesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafRulesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **orderBy** | **string** |  | 
 **orderType** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**WafRulesIndex200Response**](WafRulesIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafRulesShow

> WafRuleResponse WafRulesShow(ctx, domain, id).Execute()

Get WAF rule information

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
    resp, r, err := apiClient.WAFApi.WafRulesShow(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafRulesShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafRulesShow`: WafRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafRulesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafRulesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WafRuleResponse**](WafRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafRulesStore

> WafRuleResponse WafRulesStore(ctx, domain).WafRule(wafRule).Execute()

Create new WAF rule

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
    wafRule := *openapiclient.NewWafRule() // WafRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafRulesStore(context.Background(), domain).WafRule(wafRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafRulesStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafRulesStore`: WafRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafRulesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafRulesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wafRule** | [**WafRule**](WafRule.md) |  | 

### Return type

[**WafRuleResponse**](WafRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafRulesUpdate

> WafRuleResponse WafRulesUpdate(ctx, domain, id).WafRule(wafRule).Execute()

Update the WAF rule

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
    wafRule := *openapiclient.NewWafRule() // WafRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafRulesUpdate(context.Background(), domain, id).WafRule(wafRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafRulesUpdate`: WafRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wafRule** | [**WafRule**](WafRule.md) |  | 

### Return type

[**WafRuleResponse**](WafRuleResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafSettingsIndex

> WafSettingsData WafSettingsIndex(ctx, domain).Execute()

Get WAF configuration

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
    resp, r, err := apiClient.WAFApi.WafSettingsIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafSettingsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafSettingsIndex`: WafSettingsData
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafSettingsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafSettingsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WafSettingsData**](WafSettingsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## WafSettingsUpdate

> WafSettingsData WafSettingsUpdate(ctx, domain).WafSettings(wafSettings).Execute()

Configure WAF module of the domain

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
    wafSettings := *openapiclient.NewWafSettings() // WafSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.WafSettingsUpdate(context.Background(), domain).WafSettings(wafSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.WafSettingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WafSettingsUpdate`: WafSettingsData
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.WafSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiWafSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wafSettings** | [**WafSettings**](WafSettings.md) |  | 

### Return type

[**WafSettingsData**](WafSettingsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

