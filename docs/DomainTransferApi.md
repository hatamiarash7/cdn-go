# \DomainTransferApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsTransferIndex**](DomainTransferApi.md#DomainsTransferIndex) | **Get** /domains/transfer | Get the list of pending transfers
[**DomainsTransferStore**](DomainTransferApi.md#DomainsTransferStore) | **Post** /domains/{domain}/transfer | Transfer domain to another account
[**DomainsTransferUpdate**](DomainTransferApi.md#DomainsTransferUpdate) | **Post** /domains/transfer/change-status | Accept or cancel transferring a domain



## DomainsTransferIndex

> DomainsTransferIndex200Response DomainsTransferIndex(ctx).PerPage(perPage).Page(page).Type_(type_).Execute()

Get the list of pending transfers

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
    type_ := "type__example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainTransferApi.DomainsTransferIndex(context.Background()).PerPage(perPage).Page(page).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainTransferApi.DomainsTransferIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsTransferIndex`: DomainsTransferIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainTransferApi.DomainsTransferIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsTransferIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]
 **type_** | **string** |  | 

### Return type

[**DomainsTransferIndex200Response**](DomainsTransferIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsTransferStore

> DomainTransferData DomainsTransferStore(ctx, domain).TransferDomain(transferDomain).Execute()

Transfer domain to another account

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
    transferDomain := *openapiclient.NewTransferDomain("AccountId_example") // TransferDomain |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainTransferApi.DomainsTransferStore(context.Background(), domain).TransferDomain(transferDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainTransferApi.DomainsTransferStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsTransferStore`: DomainTransferData
    fmt.Fprintf(os.Stdout, "Response from `DomainTransferApi.DomainsTransferStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsTransferStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferDomain** | [**TransferDomain**](TransferDomain.md) |  | 

### Return type

[**DomainTransferData**](DomainTransferData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DomainsTransferUpdate

> MessageResponse DomainsTransferUpdate(ctx).TransferDomainChangeStatus(transferDomainChangeStatus).Execute()

Accept or cancel transferring a domain

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
    transferDomainChangeStatus := *openapiclient.NewTransferDomainChangeStatus("example.com", "Status_example") // TransferDomainChangeStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainTransferApi.DomainsTransferUpdate(context.Background()).TransferDomainChangeStatus(transferDomainChangeStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainTransferApi.DomainsTransferUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsTransferUpdate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainTransferApi.DomainsTransferUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsTransferUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferDomainChangeStatus** | [**TransferDomainChangeStatus**](TransferDomainChangeStatus.md) |  | 

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

