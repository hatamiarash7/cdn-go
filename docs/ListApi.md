# \ListApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListsDestroy**](ListApi.md#ListsDestroy) | **Delete** /dynamic-fields/{id} | Delete List
[**ListsIndex**](ListApi.md#ListsIndex) | **Get** /dynamic-fields | Get the list of Lists
[**ListsShow**](ListApi.md#ListsShow) | **Get** /dynamic-fields/{id} | Get an existing List
[**ListsStore**](ListApi.md#ListsStore) | **Post** /dynamic-fields | Store new List
[**ListsUpdate**](ListApi.md#ListsUpdate) | **Patch** /dynamic-fields/{id} | Update an existing List



## ListsDestroy

> MessageResponse ListsDestroy(ctx, id).Execute()

Delete List

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListApi.ListsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListApi.ListsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ListApi.ListsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsDestroyRequest struct via the builder pattern


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


## ListsIndex

> ListsIndex200Response ListsIndex(ctx).PerPage(perPage).Page(page).Scope(scope).Type_(type_).Name(name).Execute()

Get the list of Lists

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
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)
    scope := "scope_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListApi.ListsIndex(context.Background()).PerPage(perPage).Page(page).Scope(scope).Type_(type_).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListApi.ListsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsIndex`: ListsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `ListApi.ListsIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]
 **scope** | **string** |  | 
 **type_** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**ListsIndex200Response**](ListsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ListsShow

> DynamicFieldData ListsShow(ctx, id).Execute()

Get an existing List

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListApi.ListsShow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListApi.ListsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsShow`: DynamicFieldData
    fmt.Fprintf(os.Stdout, "Response from `ListApi.ListsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DynamicFieldData**](DynamicFieldData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ListsStore

> DynamicFieldResponse ListsStore(ctx).DynamicField(dynamicField).Execute()

Store new List

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
    dynamicField := *openapiclient.NewDynamicField("Name_example", "Type_example", []openapiclient.DynamicFieldValue{*openapiclient.NewDynamicFieldValue()}) // DynamicField |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListApi.ListsStore(context.Background()).DynamicField(dynamicField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListApi.ListsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsStore`: DynamicFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `ListApi.ListsStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dynamicField** | [**DynamicField**](DynamicField.md) |  | 

### Return type

[**DynamicFieldResponse**](DynamicFieldResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## ListsUpdate

> DynamicFieldResponse ListsUpdate(ctx, id).DynamicField(dynamicField).Execute()

Update an existing List

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    dynamicField := *openapiclient.NewDynamicField("Name_example", "Type_example", []openapiclient.DynamicFieldValue{*openapiclient.NewDynamicFieldValue()}) // DynamicField |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListApi.ListsUpdate(context.Background(), id).DynamicField(dynamicField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListApi.ListsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsUpdate`: DynamicFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `ListApi.ListsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicField** | [**DynamicField**](DynamicField.md) |  | 

### Return type

[**DynamicFieldResponse**](DynamicFieldResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

