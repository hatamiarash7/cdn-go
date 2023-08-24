# \CachingApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CachingDeprecatedPurge**](CachingApi.md#CachingDeprecatedPurge) | **Delete** /domains/{domain}/caching | Purge CDN Cache
[**CachingIndex**](CachingApi.md#CachingIndex) | **Get** /domains/{domain}/caching | Get caching settings
[**CachingPurge**](CachingApi.md#CachingPurge) | **Post** /domains/{domain}/caching/purge | Purge CDN Cache
[**CachingUpdate**](CachingApi.md#CachingUpdate) | **Patch** /domains/{domain}/caching | Update caching settings
[**PurgeTagsDestroy**](CachingApi.md#PurgeTagsDestroy) | **Delete** /domains/{domain}/purge-tags | Delete a Domain&#39;s Purge tag
[**PurgeTagsIndex**](CachingApi.md#PurgeTagsIndex) | **Get** /domains/{domain}/purge-tags | Get domain&#39;s Purge tags



## CachingDeprecatedPurge

> MessageResponse CachingDeprecatedPurge(ctx, domain).Purge(purge).PurgeUrls(purgeUrls).PurgeTags(purgeTags).Execute()

Purge CDN Cache



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
    purge := "purge_example" // string | 
    purgeUrls := []string{"Inner_example"} // []string | URLs to be purged from cache. Required if purge value is set to individual. (optional)
    purgeTags := []string{"Inner_example"} // []string | Tags to be purged from cache. Required if purge value is set to tags. Each tag must be 32 characters or less. Only ASCII characters are acceptable.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CachingApi.CachingDeprecatedPurge(context.Background(), domain).Purge(purge).PurgeUrls(purgeUrls).PurgeTags(purgeTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CachingApi.CachingDeprecatedPurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CachingDeprecatedPurge`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CachingApi.CachingDeprecatedPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachingDeprecatedPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purge** | **string** |  | 
 **purgeUrls** | **[]string** | URLs to be purged from cache. Required if purge value is set to individual. | 
 **purgeTags** | **[]string** | Tags to be purged from cache. Required if purge value is set to tags. Each tag must be 32 characters or less. Only ASCII characters are acceptable.  | 

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


## CachingIndex

> CacheSettingsData CachingIndex(ctx, domain).Execute()

Get caching settings

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
    resp, r, err := apiClient.CachingApi.CachingIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CachingApi.CachingIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CachingIndex`: CacheSettingsData
    fmt.Fprintf(os.Stdout, "Response from `CachingApi.CachingIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachingIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CacheSettingsData**](CacheSettingsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## CachingPurge

> MessageResponse CachingPurge(ctx, domain).CachingPurge(cachingPurge).Execute()

Purge CDN Cache



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
    cachingPurge := *openapiclient.NewCachingPurge("Purge_example") // CachingPurge |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CachingApi.CachingPurge(context.Background(), domain).CachingPurge(cachingPurge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CachingApi.CachingPurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CachingPurge`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CachingApi.CachingPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachingPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cachingPurge** | [**CachingPurge**](CachingPurge.md) |  | 

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


## CachingUpdate

> MessageResponse CachingUpdate(ctx, domain).CacheSettings(cacheSettings).Execute()

Update caching settings

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
    cacheSettings := *openapiclient.NewCacheSettings() // CacheSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CachingApi.CachingUpdate(context.Background(), domain).CacheSettings(cacheSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CachingApi.CachingUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CachingUpdate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CachingApi.CachingUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCachingUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cacheSettings** | [**CacheSettings**](CacheSettings.md) |  | 

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


## PurgeTagsDestroy

> MessageResponse PurgeTagsDestroy(ctx, domain).Tag(tag).Execute()

Delete a Domain's Purge tag

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
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CachingApi.PurgeTagsDestroy(context.Background(), domain).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CachingApi.PurgeTagsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurgeTagsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CachingApi.PurgeTagsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeTagsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **string** |  | 

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


## PurgeTagsIndex

> PurgeTagsIndex200Response PurgeTagsIndex(ctx, domain).Execute()

Get domain's Purge tags



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
    resp, r, err := apiClient.CachingApi.PurgeTagsIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CachingApi.PurgeTagsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurgeTagsIndex`: PurgeTagsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `CachingApi.PurgeTagsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeTagsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurgeTagsIndex200Response**](PurgeTagsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

