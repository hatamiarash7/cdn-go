# \LoadBalancingApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadBalancersDestroy**](LoadBalancingApi.md#LoadBalancersDestroy) | **Delete** /domains/{domain}/load-balancers/{loadBalancerId} | Remove a load balancer
[**LoadBalancersIndex**](LoadBalancingApi.md#LoadBalancersIndex) | **Get** /domains/{domain}/load-balancers | Get list of load balancers
[**LoadBalancersPoolsDestroy**](LoadBalancingApi.md#LoadBalancersPoolsDestroy) | **Delete** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId} | Remove a load balancer pool
[**LoadBalancersPoolsIndex**](LoadBalancingApi.md#LoadBalancersPoolsIndex) | **Get** /domains/{domain}/load-balancers/{loadBalancerId}/pools | Get the list of pools of a load balancers
[**LoadBalancersPoolsOriginsDestroy**](LoadBalancingApi.md#LoadBalancersPoolsOriginsDestroy) | **Delete** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins/{loadBalancerPoolOriginId} | Remove an origin from the pool of the load balancer
[**LoadBalancersPoolsOriginsIndex**](LoadBalancingApi.md#LoadBalancersPoolsOriginsIndex) | **Get** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins | Get list of origins of a pool
[**LoadBalancersPoolsOriginsShow**](LoadBalancingApi.md#LoadBalancersPoolsOriginsShow) | **Get** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins/{loadBalancerPoolOriginId} | Get load balancer origin information
[**LoadBalancersPoolsOriginsStore**](LoadBalancingApi.md#LoadBalancersPoolsOriginsStore) | **Post** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins | Create a new origin in the pool of the load balancer
[**LoadBalancersPoolsOriginsUpdate**](LoadBalancingApi.md#LoadBalancersPoolsOriginsUpdate) | **Patch** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins/{loadBalancerPoolOriginId} | Update an existing origin of the pool
[**LoadBalancersPoolsShow**](LoadBalancingApi.md#LoadBalancersPoolsShow) | **Get** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId} | Get load balancer pool information
[**LoadBalancersPoolsStore**](LoadBalancingApi.md#LoadBalancersPoolsStore) | **Post** /domains/{domain}/load-balancers/{loadBalancerId}/pools | Create a new pool for the load balancer
[**LoadBalancersPoolsUpdate**](LoadBalancingApi.md#LoadBalancersPoolsUpdate) | **Put** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId} | Update an existing load balancer pool with origins
[**LoadBalancersPoolsUpdatePool**](LoadBalancingApi.md#LoadBalancersPoolsUpdatePool) | **Patch** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId} | Update an existing load balancer pool
[**LoadBalancersPrioritizePool**](LoadBalancingApi.md#LoadBalancersPrioritizePool) | **Post** /domains/{domain}/load-balancers/{loadBalancerId}/prioritize | Reorder the priority of load balancer pools
[**LoadBalancersRegionsIndex**](LoadBalancingApi.md#LoadBalancersRegionsIndex) | **Get** /load-balancers/regions | Get list of regions for load balancers
[**LoadBalancersSettingsShow**](LoadBalancingApi.md#LoadBalancersSettingsShow) | **Get** /domains/{domain}/load-balancers/settings | Get list of domain load balancer global settings
[**LoadBalancersSettingsUpdate**](LoadBalancingApi.md#LoadBalancersSettingsUpdate) | **Patch** /domains/{domain}/load-balancers/settings | Update domain&#39;s global load balancer settings
[**LoadBalancersShow**](LoadBalancingApi.md#LoadBalancersShow) | **Get** /domains/{domain}/load-balancers/{loadBalancerId} | Get load balancer information
[**LoadBalancersStore**](LoadBalancingApi.md#LoadBalancersStore) | **Post** /domains/{domain}/load-balancers | Create a new load balancer
[**LoadBalancersUpdate**](LoadBalancingApi.md#LoadBalancersUpdate) | **Patch** /domains/{domain}/load-balancers/{loadBalancerId} | Update a load balancer
[**LoadBalancersUpdateWithPools**](LoadBalancingApi.md#LoadBalancersUpdateWithPools) | **Put** /domains/{domain}/load-balancers/{loadBalancerId} | Update a load balancer



## LoadBalancersDestroy

> MessageResponse LoadBalancersDestroy(ctx, domain, loadBalancerId).Execute()

Remove a load balancer

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersDestroy(context.Background(), domain, loadBalancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersDestroyRequest struct via the builder pattern


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


## LoadBalancersIndex

> LoadBalancersIndex200Response LoadBalancersIndex(ctx, domain).Execute()

Get list of load balancers

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
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersIndex(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersIndex`: LoadBalancersIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoadBalancersIndex200Response**](LoadBalancersIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsDestroy

> MessageResponse LoadBalancersPoolsDestroy(ctx, domain, loadBalancerId, loadBalancerPoolId).Execute()

Remove a load balancer pool

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsDestroy(context.Background(), domain, loadBalancerId, loadBalancerPoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsDestroyRequest struct via the builder pattern


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


## LoadBalancersPoolsIndex

> LoadBalancersPoolsIndex200Response LoadBalancersPoolsIndex(ctx, domain, loadBalancerId).Execute()

Get the list of pools of a load balancers

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsIndex(context.Background(), domain, loadBalancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsIndex`: LoadBalancersPoolsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoadBalancersPoolsIndex200Response**](LoadBalancersPoolsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsOriginsDestroy

> MessageResponse LoadBalancersPoolsOriginsDestroy(ctx, domain, loadBalancerId, loadBalancerPoolId, loadBalancerPoolOriginId).Execute()

Remove an origin from the pool of the load balancer

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer
    loadBalancerPoolOriginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of an origin of the pool in the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsOriginsDestroy(context.Background(), domain, loadBalancerId, loadBalancerPoolId, loadBalancerPoolOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsOriginsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsOriginsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsOriginsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 
**loadBalancerPoolOriginId** | **string** | ID of an origin of the pool in the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsOriginsDestroyRequest struct via the builder pattern


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


## LoadBalancersPoolsOriginsIndex

> LoadBalancersPoolsOriginsIndex200Response LoadBalancersPoolsOriginsIndex(ctx, domain, loadBalancerId, loadBalancerPoolId).Execute()

Get list of origins of a pool

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsOriginsIndex(context.Background(), domain, loadBalancerId, loadBalancerPoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsOriginsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsOriginsIndex`: LoadBalancersPoolsOriginsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsOriginsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsOriginsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LoadBalancersPoolsOriginsIndex200Response**](LoadBalancersPoolsOriginsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsOriginsShow

> LoadBalancerOriginData LoadBalancersPoolsOriginsShow(ctx, domain, loadBalancerId, loadBalancerPoolId, loadBalancerPoolOriginId).Execute()

Get load balancer origin information

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer
    loadBalancerPoolOriginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of an origin of the pool in the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsOriginsShow(context.Background(), domain, loadBalancerId, loadBalancerPoolId, loadBalancerPoolOriginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsOriginsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsOriginsShow`: LoadBalancerOriginData
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsOriginsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 
**loadBalancerPoolOriginId** | **string** | ID of an origin of the pool in the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsOriginsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**LoadBalancerOriginData**](LoadBalancerOriginData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsOriginsStore

> LoadBalancerOriginResponse LoadBalancersPoolsOriginsStore(ctx, domain, loadBalancerId, loadBalancerPoolId).LoadBalancerOriginStore(loadBalancerOriginStore).Execute()

Create a new origin in the pool of the load balancer

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer
    loadBalancerOriginStore := *openapiclient.NewLoadBalancerOriginStore(false, "Address_example", int32(123), int32(123), "Protocol_example") // LoadBalancerOriginStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsOriginsStore(context.Background(), domain, loadBalancerId, loadBalancerPoolId).LoadBalancerOriginStore(loadBalancerOriginStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsOriginsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsOriginsStore`: LoadBalancerOriginResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsOriginsStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsOriginsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **loadBalancerOriginStore** | [**LoadBalancerOriginStore**](LoadBalancerOriginStore.md) |  | 

### Return type

[**LoadBalancerOriginResponse**](LoadBalancerOriginResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsOriginsUpdate

> LoadBalancerOriginResponse LoadBalancersPoolsOriginsUpdate(ctx, domain, loadBalancerId, loadBalancerPoolId, loadBalancerPoolOriginId).LoadBalancerOrigin(loadBalancerOrigin).Execute()

Update an existing origin of the pool

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer
    loadBalancerPoolOriginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of an origin of the pool in the load balancer
    loadBalancerOrigin := *openapiclient.NewLoadBalancerOrigin() // LoadBalancerOrigin |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsOriginsUpdate(context.Background(), domain, loadBalancerId, loadBalancerPoolId, loadBalancerPoolOriginId).LoadBalancerOrigin(loadBalancerOrigin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsOriginsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsOriginsUpdate`: LoadBalancerOriginResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsOriginsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 
**loadBalancerPoolOriginId** | **string** | ID of an origin of the pool in the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsOriginsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **loadBalancerOrigin** | [**LoadBalancerOrigin**](LoadBalancerOrigin.md) |  | 

### Return type

[**LoadBalancerOriginResponse**](LoadBalancerOriginResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsShow

> LoadBalancerPoolData LoadBalancersPoolsShow(ctx, domain, loadBalancerId, loadBalancerPoolId).Execute()

Get load balancer pool information

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsShow(context.Background(), domain, loadBalancerId, loadBalancerPoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsShow`: LoadBalancerPoolData
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LoadBalancerPoolData**](LoadBalancerPoolData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsStore

> LoadBalancerPoolResponse LoadBalancersPoolsStore(ctx, domain, loadBalancerId).LoadBalancerPoolStore(loadBalancerPoolStore).Execute()

Create a new pool for the load balancer

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolStore := *openapiclient.NewLoadBalancerPoolStore("Name_example", false, "Method_example", "Keepalive_example", "NextUpstreamTcp_example") // LoadBalancerPoolStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsStore(context.Background(), domain, loadBalancerId).LoadBalancerPoolStore(loadBalancerPoolStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsStore`: LoadBalancerPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **loadBalancerPoolStore** | [**LoadBalancerPoolStore**](LoadBalancerPoolStore.md) |  | 

### Return type

[**LoadBalancerPoolResponse**](LoadBalancerPoolResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsUpdate

> LoadBalancerPoolResponse LoadBalancersPoolsUpdate(ctx, domain, loadBalancerId, loadBalancerPoolId).LoadBalancerPoolStore(loadBalancerPoolStore).Execute()

Update an existing load balancer pool with origins

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer
    loadBalancerPoolStore := *openapiclient.NewLoadBalancerPoolStore("Name_example", false, "Method_example", "Keepalive_example", "NextUpstreamTcp_example") // LoadBalancerPoolStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsUpdate(context.Background(), domain, loadBalancerId, loadBalancerPoolId).LoadBalancerPoolStore(loadBalancerPoolStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsUpdate`: LoadBalancerPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **loadBalancerPoolStore** | [**LoadBalancerPoolStore**](LoadBalancerPoolStore.md) |  | 

### Return type

[**LoadBalancerPoolResponse**](LoadBalancerPoolResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPoolsUpdatePool

> LoadBalancerPoolResponse LoadBalancersPoolsUpdatePool(ctx, domain, loadBalancerId, loadBalancerPoolId).LoadBalancerPool(loadBalancerPool).Execute()

Update an existing load balancer pool

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerPoolId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of a pool of the load balancer
    loadBalancerPool := *openapiclient.NewLoadBalancerPool() // LoadBalancerPool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPoolsUpdatePool(context.Background(), domain, loadBalancerId, loadBalancerPoolId).LoadBalancerPool(loadBalancerPool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPoolsUpdatePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPoolsUpdatePool`: LoadBalancerPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPoolsUpdatePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 
**loadBalancerPoolId** | **string** | ID of a pool of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPoolsUpdatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **loadBalancerPool** | [**LoadBalancerPool**](LoadBalancerPool.md) |  | 

### Return type

[**LoadBalancerPoolResponse**](LoadBalancerPoolResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersPrioritizePool

> LoadBalancerResponse LoadBalancersPrioritizePool(ctx, domain, loadBalancerId).PrioritizePool(prioritizePool).Execute()

Reorder the priority of load balancer pools

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    prioritizePool := *openapiclient.NewPrioritizePool("PoolId_example", "AfterPoolId_example", "BeforePoolId_example") // PrioritizePool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersPrioritizePool(context.Background(), domain, loadBalancerId).PrioritizePool(prioritizePool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersPrioritizePool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersPrioritizePool`: LoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersPrioritizePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersPrioritizePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **prioritizePool** | [**PrioritizePool**](PrioritizePool.md) |  | 

### Return type

[**LoadBalancerResponse**](LoadBalancerResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersRegionsIndex

> LoadBalancersRegionsIndex200Response LoadBalancersRegionsIndex(ctx).Execute()

Get list of regions for load balancers

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
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersRegionsIndex(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersRegionsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersRegionsIndex`: LoadBalancersRegionsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersRegionsIndex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersRegionsIndexRequest struct via the builder pattern


### Return type

[**LoadBalancersRegionsIndex200Response**](LoadBalancersRegionsIndex200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersSettingsShow

> LoadBalancerSettingsData LoadBalancersSettingsShow(ctx, domain).Execute()

Get list of domain load balancer global settings

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
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersSettingsShow(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersSettingsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersSettingsShow`: LoadBalancerSettingsData
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersSettingsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersSettingsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoadBalancerSettingsData**](LoadBalancerSettingsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersSettingsUpdate

> LoadBalancersSettingsUpdate200Response LoadBalancersSettingsUpdate(ctx, domain).LoadBalancerSetting(loadBalancerSetting).Execute()

Update domain's global load balancer settings

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
    loadBalancerSetting := *openapiclient.NewLoadBalancerSetting() // LoadBalancerSetting |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersSettingsUpdate(context.Background(), domain).LoadBalancerSetting(loadBalancerSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersSettingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersSettingsUpdate`: LoadBalancersSettingsUpdate200Response
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadBalancerSetting** | [**LoadBalancerSetting**](LoadBalancerSetting.md) |  | 

### Return type

[**LoadBalancersSettingsUpdate200Response**](LoadBalancersSettingsUpdate200Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersShow

> LoadBalancerData LoadBalancersShow(ctx, domain, loadBalancerId).Execute()

Get load balancer information

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersShow(context.Background(), domain, loadBalancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersShow`: LoadBalancerData
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoadBalancerData**](LoadBalancerData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersStore

> LoadBalancerResponse LoadBalancersStore(ctx, domain).LoadBalancerStore(loadBalancerStore).Execute()

Create a new load balancer

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
    loadBalancerStore := *openapiclient.NewLoadBalancerStore("lb1", false, "Method_example") // LoadBalancerStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersStore(context.Background(), domain).LoadBalancerStore(loadBalancerStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersStore`: LoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loadBalancerStore** | [**LoadBalancerStore**](LoadBalancerStore.md) |  | 

### Return type

[**LoadBalancerResponse**](LoadBalancerResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersUpdate

> LoadBalancerResponse LoadBalancersUpdate(ctx, domain, loadBalancerId).LoadBalancer(loadBalancer).Execute()

Update a load balancer

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancer := *openapiclient.NewLoadBalancer() // LoadBalancer |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersUpdate(context.Background(), domain, loadBalancerId).LoadBalancer(loadBalancer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersUpdate`: LoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **loadBalancer** | [**LoadBalancer**](LoadBalancer.md) |  | 

### Return type

[**LoadBalancerResponse**](LoadBalancerResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## LoadBalancersUpdateWithPools

> LoadBalancerResponse LoadBalancersUpdateWithPools(ctx, domain, loadBalancerId).LoadBalancerUpdate(loadBalancerUpdate).Execute()

Update a load balancer

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
    loadBalancerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the load balancer
    loadBalancerUpdate := *openapiclient.NewLoadBalancerUpdate() // LoadBalancerUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancingApi.LoadBalancersUpdateWithPools(context.Background(), domain, loadBalancerId).LoadBalancerUpdate(loadBalancerUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancingApi.LoadBalancersUpdateWithPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadBalancersUpdateWithPools`: LoadBalancerResponse
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancingApi.LoadBalancersUpdateWithPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadBalancersUpdateWithPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **loadBalancerUpdate** | [**LoadBalancerUpdate**](LoadBalancerUpdate.md) |  | 

### Return type

[**LoadBalancerResponse**](LoadBalancerResponse.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

