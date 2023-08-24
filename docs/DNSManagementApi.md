# \DNSManagementApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsRecordsCloud**](DNSManagementApi.md#DnsRecordsCloud) | **Put** /domains/{domain}/dns-records/{id}/cloud | Toggle cloud status (To proxy or not proxy, that&#39;s the question!)
[**DnsRecordsDestroy**](DNSManagementApi.md#DnsRecordsDestroy) | **Delete** /domains/{domain}/dns-records/{id} | Remove a DNS record
[**DnsRecordsDnsSecShow**](DNSManagementApi.md#DnsRecordsDnsSecShow) | **Get** /domains/{domain}/dns-records/dnssec | Get status of DNSSEC
[**DnsRecordsDnsSecUpdate**](DNSManagementApi.md#DnsRecordsDnsSecUpdate) | **Put** /domains/{domain}/dns-records/dnssec/actions | Update DNSSEC status
[**DnsRecordsImport**](DNSManagementApi.md#DnsRecordsImport) | **Post** /domains/{domain}/dns-records/import | Import DNS records using BIND file
[**DnsRecordsIndex**](DNSManagementApi.md#DnsRecordsIndex) | **Get** /domains/{domain}/dns-records | Get list of DNS records
[**DnsRecordsShow**](DNSManagementApi.md#DnsRecordsShow) | **Get** /domains/{domain}/dns-records/{id} | Get information of a record
[**DnsRecordsStore**](DNSManagementApi.md#DnsRecordsStore) | **Post** /domains/{domain}/dns-records | Create new DNS record
[**DnsRecordsUpdate**](DNSManagementApi.md#DnsRecordsUpdate) | **Put** /domains/{domain}/dns-records/{id} | Update a DNS record



## DnsRecordsCloud

> DnsRecordResponse DnsRecordsCloud(ctx, domain, id).DnsRecordCloud(dnsRecordCloud).Execute()

Toggle cloud status (To proxy or not proxy, that's the question!)

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the DNS record
    dnsRecordCloud := *openapiclient.NewDnsRecordCloud(false) // DnsRecordCloud |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsCloud(context.Background(), domain, id).DnsRecordCloud(dnsRecordCloud).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsCloud``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsCloud`: DnsRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** | ID of the DNS record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsCloudRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsRecordCloud** | [**DnsRecordCloud**](DnsRecordCloud.md) |  | 

### Return type

[**DnsRecordResponse**](DnsRecordResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DnsRecordsDestroy

> MessageResponse DnsRecordsDestroy(ctx, domain, id).Execute()

Remove a DNS record

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the DNS record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsDestroy(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** | ID of the DNS record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsDestroyRequest struct via the builder pattern


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


## DnsRecordsDnsSecShow

> DnsSecData DnsRecordsDnsSecShow(ctx, domain).Execute()

Get status of DNSSEC

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
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsDnsSecShow(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsDnsSecShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsDnsSecShow`: DnsSecData
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsDnsSecShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsDnsSecShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnsSecData**](DnsSecData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DnsRecordsDnsSecUpdate

> DnsSecData DnsRecordsDnsSecUpdate(ctx, domain).DnsSecStatus(dnsSecStatus).Execute()

Update DNSSEC status

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
    dnsSecStatus := *openapiclient.NewDnsSecStatus(false) // DnsSecStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsDnsSecUpdate(context.Background(), domain).DnsSecStatus(dnsSecStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsDnsSecUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsDnsSecUpdate`: DnsSecData
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsDnsSecUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsDnsSecUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsSecStatus** | [**DnsSecStatus**](DnsSecStatus.md) |  | 

### Return type

[**DnsSecData**](DnsSecData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DnsRecordsImport

> MessageResponse DnsRecordsImport(ctx, domain).FZoneFile(fZoneFile).Execute()

Import DNS records using BIND file

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
    fZoneFile := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsImport(context.Background(), domain).FZoneFile(fZoneFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsImport`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fZoneFile** | ***os.File** |  | 

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


## DnsRecordsIndex

> DnsRecordsIndex200Response DnsRecordsIndex(ctx, domain).Search(search).Type_(type_).Page(page).PerPage(perPage).Execute()

Get list of DNS records

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
    search := "search_example" // string | Search term (optional)
    type_ := "a,cname,txt" // string | Type of a dns record. To filter multiple types separate them using a comma (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsIndex(context.Background(), domain).Search(search).Type_(type_).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsIndex`: DnsRecordsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** | Search term | 
 **type_** | **string** | Type of a dns record. To filter multiple types separate them using a comma | 
 **page** | **int32** | Set the desired page number | [default to 1]
 **perPage** | **int32** | Set how many items returned per page | 

### Return type

[**DnsRecordsIndex200Response**](DnsRecordsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DnsRecordsShow

> DnsRecordData DnsRecordsShow(ctx, domain, id).Execute()

Get information of a record

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the DNS record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsShow(context.Background(), domain, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsShow`: DnsRecordData
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** | ID of the DNS record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DnsRecordData**](DnsRecordData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DnsRecordsStore

> DnsRecordResponse DnsRecordsStore(ctx, domain).DnsRecord(dnsRecord).Execute()

Create new DNS record

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
    dnsRecord := openapiclient.DnsRecord{AAAARecord: openapiclient.NewAAAARecord()} // DnsRecord |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsStore(context.Background(), domain).DnsRecord(dnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsStore`: DnsRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsRecord** | [**DnsRecord**](DnsRecord.md) |  | 

### Return type

[**DnsRecordResponse**](DnsRecordResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## DnsRecordsUpdate

> DnsRecordResponse DnsRecordsUpdate(ctx, domain, id).DnsRecord(dnsRecord).Execute()

Update a DNS record

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the DNS record
    dnsRecord := openapiclient.DnsRecord{AAAARecord: openapiclient.NewAAAARecord()} // DnsRecord |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSManagementApi.DnsRecordsUpdate(context.Background(), domain, id).DnsRecord(dnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSManagementApi.DnsRecordsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRecordsUpdate`: DnsRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSManagementApi.DnsRecordsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**id** | **string** | ID of the DNS record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsRecordsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsRecord** | [**DnsRecord**](DnsRecord.md) |  | 

### Return type

[**DnsRecordResponse**](DnsRecordResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

