# \SSLTLSApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SslCertDestroy**](SSLTLSApi.md#SslCertDestroy) | **Delete** /domains/{domain}/ssl/certificates/{id} | Delete an unused customer certificate
[**SslCertOrderIndex**](SSLTLSApi.md#SslCertOrderIndex) | **Get** /domains/{domain}/ssl/orders | Get All Managed certificate orders history
[**SslCertOrderRetry**](SSLTLSApi.md#SslCertOrderRetry) | **Post** /domains/{domain}/ssl/orders/action/retry | Retry a previously &#x60;killed&#x60; order
[**SslCertStore**](SSLTLSApi.md#SslCertStore) | **Post** /domains/{domain}/ssl/certificates | Upload Certificate
[**SslIndex**](SSLTLSApi.md#SslIndex) | **Get** /domains/{domain}/ssl | Get SSL settings
[**SslUpdate**](SSLTLSApi.md#SslUpdate) | **Patch** /domains/{domain}/ssl | Update domain&#39;s SSL settings



## SslCertDestroy

> MessageResponse SslCertDestroy(ctx, domain, id).Execute()

Delete an unused customer certificate

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
    resp, r, err := apiClient.SSLTLSApi.SslCertDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLTLSApi.SslCertDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslCertDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLTLSApi.SslCertDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslCertDestroyRequest struct via the builder pattern


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


## SslCertOrderIndex

> SslCertOrderIndex200Response SslCertOrderIndex(ctx, domain).Execute()

Get All Managed certificate orders history

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
    resp, r, err := apiClient.SSLTLSApi.SslCertOrderIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLTLSApi.SslCertOrderIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslCertOrderIndex`: SslCertOrderIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `SSLTLSApi.SslCertOrderIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslCertOrderIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SslCertOrderIndex200Response**](SslCertOrderIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## SslCertOrderRetry

> MessageResponse SslCertOrderRetry(ctx, domain).Execute()

Retry a previously `killed` order

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
    resp, r, err := apiClient.SSLTLSApi.SslCertOrderRetry(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLTLSApi.SslCertOrderRetry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslCertOrderRetry`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLTLSApi.SslCertOrderRetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslCertOrderRetryRequest struct via the builder pattern


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


## SslCertStore

> MessageResponse SslCertStore(ctx, domain).Certificate(certificate).PrivateKey(privateKey).Execute()

Upload Certificate

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
    certificate := os.NewFile(1234, "some_file") // *os.File |  (optional)
    privateKey := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSLTLSApi.SslCertStore(context.Background(), domain).Certificate(certificate).PrivateKey(privateKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLTLSApi.SslCertStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslCertStore`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLTLSApi.SslCertStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslCertStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificate** | ***os.File** |  | 
 **privateKey** | ***os.File** |  | 

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


## SslIndex

> SslResponse SslIndex(ctx, domain).Execute()

Get SSL settings

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
    resp, r, err := apiClient.SSLTLSApi.SslIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLTLSApi.SslIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslIndex`: SslResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLTLSApi.SslIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SslResponse**](SslResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## SslUpdate

> SslResponse SslUpdate(ctx, domain).SslUpdate(sslUpdate).Execute()

Update domain's SSL settings

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
    sslUpdate := *openapiclient.NewSslUpdate() // SslUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSLTLSApi.SslUpdate(context.Background(), domain).SslUpdate(sslUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSLTLSApi.SslUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SslUpdate`: SslResponse
    fmt.Fprintf(os.Stdout, "Response from `SSLTLSApi.SslUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSslUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sslUpdate** | [**SslUpdate**](SslUpdate.md) |  | 

### Return type

[**SslResponse**](SslResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

