# \TroubleshootApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TroubleshootsIndex**](TroubleshootApi.md#TroubleshootsIndex) | **Get** /domains/{domain}/troubleshoots | Show list of troubleshoots for given domain
[**TroubleshootsLatest**](TroubleshootApi.md#TroubleshootsLatest) | **Get** /domains/{domain}/troubleshoots/latest | Show the latest troubleshoot for given domain
[**TroubleshootsStore**](TroubleshootApi.md#TroubleshootsStore) | **Post** /domains/{domain}/troubleshoots | Create new troubleshoot



## TroubleshootsIndex

> TroubleshootsIndex200Response TroubleshootsIndex(ctx, domain).PerPage(perPage).Page(page).Execute()

Show list of troubleshoots for given domain

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
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TroubleshootApi.TroubleshootsIndex(context.Background(), domain).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootApi.TroubleshootsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TroubleshootsIndex`: TroubleshootsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootApi.TroubleshootsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTroubleshootsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**TroubleshootsIndex200Response**](TroubleshootsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## TroubleshootsLatest

> TroubleshootsStore201Response TroubleshootsLatest(ctx, domain).Execute()

Show the latest troubleshoot for given domain

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
    resp, r, err := apiClient.TroubleshootApi.TroubleshootsLatest(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootApi.TroubleshootsLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TroubleshootsLatest`: TroubleshootsStore201Response
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootApi.TroubleshootsLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTroubleshootsLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TroubleshootsStore201Response**](TroubleshootsStore201Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## TroubleshootsStore

> TroubleshootsStore201Response TroubleshootsStore(ctx, domain).Execute()

Create new troubleshoot

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
    resp, r, err := apiClient.TroubleshootApi.TroubleshootsStore(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootApi.TroubleshootsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TroubleshootsStore`: TroubleshootsStore201Response
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootApi.TroubleshootsStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTroubleshootsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TroubleshootsStore201Response**](TroubleshootsStore201Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

