# \CDNAppsApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsCategoryIndex**](CDNAppsApi.md#AppsCategoryIndex) | **Get** /apps/category | Get the list of application categories
[**AppsCategoryShow**](CDNAppsApi.md#AppsCategoryShow) | **Get** /apps/category/{application-category} | Get an existing application category
[**AppsIndex**](CDNAppsApi.md#AppsIndex) | **Get** /apps | Get list of all available cdn-apps
[**AppsLike**](CDNAppsApi.md#AppsLike) | **Post** /apps/{id} | Expressing like and dislike about a single cdn-app
[**AppsShow**](CDNAppsApi.md#AppsShow) | **Get** /apps/{id} | Get a single cdn-app
[**DomainsAppsDestroy**](CDNAppsApi.md#DomainsAppsDestroy) | **Delete** /domains/{domain}/apps/{id} | Uninstall the application from domain
[**DomainsAppsIndex**](CDNAppsApi.md#DomainsAppsIndex) | **Get** /domains/{domain}/apps | Get list of all applications installed on a domain
[**DomainsAppsInstalled**](CDNAppsApi.md#DomainsAppsInstalled) | **Get** /domains/{domain}/apps/{id} | Check the application is installed on the domain
[**DomainsAppsStore**](CDNAppsApi.md#DomainsAppsStore) | **Post** /domains/{domain}/apps/{id} | Install the application on the domain
[**DomainsAppsTriggerWebhook**](CDNAppsApi.md#DomainsAppsTriggerWebhook) | **Post** /domains/{domain}/apps/{id}/actions/trigger_webhook | trigger webhook event



## AppsCategoryIndex

> AppsCategoryIndex200Response AppsCategoryIndex(ctx).Categories(categories).Execute()

Get the list of application categories

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
    categories := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAppsApi.AppsCategoryIndex(context.Background()).Categories(categories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.AppsCategoryIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsCategoryIndex`: AppsCategoryIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.AppsCategoryIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsCategoryIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categories** | **[]string** |  | 

### Return type

[**AppsCategoryIndex200Response**](AppsCategoryIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## AppsCategoryShow

> AppsCategoryShow200Response AppsCategoryShow(ctx, applicationCategory).Execute()

Get an existing application category

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
    applicationCategory := "applicationCategory_example" // string | The id of the category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAppsApi.AppsCategoryShow(context.Background(), applicationCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.AppsCategoryShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsCategoryShow`: AppsCategoryShow200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.AppsCategoryShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationCategory** | **string** | The id of the category | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsCategoryShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppsCategoryShow200Response**](AppsCategoryShow200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## AppsIndex

> AppsIndex200Response AppsIndex(ctx).CategoryId(categoryId).Search(search).PerPage(perPage).Page(page).SortBy(sortBy).Direction(direction).Execute()

Get list of all available cdn-apps

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
    categoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter apps by category (optional)
    search := "search_example" // string | Search term (optional)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)
    sortBy := "sortBy_example" // string |  (optional)
    direction := "direction_example" // string | Set the direction of sorting (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAppsApi.AppsIndex(context.Background()).CategoryId(categoryId).Search(search).PerPage(perPage).Page(page).SortBy(sortBy).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.AppsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsIndex`: AppsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.AppsIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **string** | Filter apps by category | 
 **search** | **string** | Search term | 
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]
 **sortBy** | **string** |  | 
 **direction** | **string** | Set the direction of sorting | 

### Return type

[**AppsIndex200Response**](AppsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## AppsLike

> CdnAppLikeStatsData AppsLike(ctx, id).CdnAppLike(cdnAppLike).Execute()

Expressing like and dislike about a single cdn-app

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    cdnAppLike := *openapiclient.NewCdnAppLike() // CdnAppLike |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAppsApi.AppsLike(context.Background(), id).CdnAppLike(cdnAppLike).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.AppsLike``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsLike`: CdnAppLikeStatsData
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.AppsLike`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsLikeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnAppLike** | [**CdnAppLike**](CdnAppLike.md) |  | 

### Return type

[**CdnAppLikeStatsData**](CdnAppLikeStatsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## AppsShow

> CdnAppData AppsShow(ctx, id).Execute()

Get a single cdn-app

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAppsApi.AppsShow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.AppsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsShow`: CdnAppData
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.AppsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CdnAppData**](CdnAppData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsAppsDestroy

> MessageResponse DomainsAppsDestroy(ctx, domain, id).Execute()

Uninstall the application from domain

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
    resp, r, err := apiClient.CDNAppsApi.DomainsAppsDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.DomainsAppsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsAppsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.DomainsAppsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsAppsDestroyRequest struct via the builder pattern


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


## DomainsAppsIndex

> AppsIndex200Response DomainsAppsIndex(ctx, domain).Execute()

Get list of all applications installed on a domain

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
    resp, r, err := apiClient.CDNAppsApi.DomainsAppsIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.DomainsAppsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsAppsIndex`: AppsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.DomainsAppsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsAppsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppsIndex200Response**](AppsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsAppsInstalled

> CdnAppInstall DomainsAppsInstalled(ctx, domain, id).Execute()

Check the application is installed on the domain

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
    resp, r, err := apiClient.CDNAppsApi.DomainsAppsInstalled(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.DomainsAppsInstalled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsAppsInstalled`: CdnAppInstall
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.DomainsAppsInstalled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsAppsInstalledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CdnAppInstall**](CdnAppInstall.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsAppsStore

> DomainsAppsStore200Response DomainsAppsStore(ctx, domain, id).Body(body).Execute()

Install the application on the domain

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAppsApi.DomainsAppsStore(context.Background(), domain, id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.DomainsAppsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsAppsStore`: DomainsAppsStore200Response
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.DomainsAppsStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsAppsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**DomainsAppsStore200Response**](DomainsAppsStore200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsAppsTriggerWebhook

> MessageResponse DomainsAppsTriggerWebhook(ctx, domain, id).CdnAppTriggerWebhook(cdnAppTriggerWebhook).Execute()

trigger webhook event

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
    cdnAppTriggerWebhook := *openapiclient.NewCdnAppTriggerWebhook("Event_example", map[string]interface{}(123)) // CdnAppTriggerWebhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CDNAppsApi.DomainsAppsTriggerWebhook(context.Background(), domain, id).CdnAppTriggerWebhook(cdnAppTriggerWebhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CDNAppsApi.DomainsAppsTriggerWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsAppsTriggerWebhook`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CDNAppsApi.DomainsAppsTriggerWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsAppsTriggerWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cdnAppTriggerWebhook** | [**CdnAppTriggerWebhook**](CdnAppTriggerWebhook.md) |  | 

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

