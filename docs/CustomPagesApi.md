# \CustomPagesApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomPagesGet**](CustomPagesApi.md#CustomPagesGet) | **Get** /domains/{domain}/custom-pages | Get list of custom pages
[**CustomPagesUpdate**](CustomPagesApi.md#CustomPagesUpdate) | **Post** /domains/{domain}/custom-pages | Update custom page



## CustomPagesGet

> CustomPagesData CustomPagesGet(ctx, domain).Execute()

Get list of custom pages

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
    resp, r, err := apiClient.CustomPagesApi.CustomPagesGet(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesApi.CustomPagesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomPagesGet`: CustomPagesData
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesApi.CustomPagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomPagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomPagesData**](CustomPagesData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## CustomPagesUpdate

> MessageResponse CustomPagesUpdate(ctx, domain).Type_(type_).Page(page).Url(url).File(file).Execute()

Update custom page

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
    type_ := "type__example" // string |  (optional)
    page := "page_example" // string | ddos_js and ddos_captcha can only be used with file type (optional)
    url := "url_example" // string |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPagesApi.CustomPagesUpdate(context.Background(), domain).Type_(type_).Page(page).Url(url).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPagesApi.CustomPagesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomPagesUpdate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomPagesApi.CustomPagesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomPagesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** |  | 
 **page** | **string** | ddos_js and ddos_captcha can only be used with file type | 
 **url** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

