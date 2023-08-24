# \EmailForwardingApi

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailForwardingsActivate**](EmailForwardingApi.md#EmailForwardingsActivate) | **Post** /domains/{domain}/email-forwardings/activate | Activate email forwarding
[**EmailForwardingsAliasesDestroy**](EmailForwardingApi.md#EmailForwardingsAliasesDestroy) | **Delete** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId} | Delete an alias
[**EmailForwardingsAliasesIndex**](EmailForwardingApi.md#EmailForwardingsAliasesIndex) | **Get** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases | List of email forwarding aliases for given domain
[**EmailForwardingsAliasesStore**](EmailForwardingApi.md#EmailForwardingsAliasesStore) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases | Create new alias
[**EmailForwardingsAliasesToggleActivation**](EmailForwardingApi.md#EmailForwardingsAliasesToggleActivation) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId}/toggle-activation | Toggle alias activation
[**EmailForwardingsAliasesUpdateRecipients**](EmailForwardingApi.md#EmailForwardingsAliasesUpdateRecipients) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId}/recipients | Update alias recipients
[**EmailForwardingsCatchAllActivate**](EmailForwardingApi.md#EmailForwardingsCatchAllActivate) | **Post** /domains/{domain}/email-forwardings/catch-all/activate | Activate email forwarding catch all
[**EmailForwardingsCatchAllDeactivate**](EmailForwardingApi.md#EmailForwardingsCatchAllDeactivate) | **Post** /domains/{domain}/email-forwardings/catch-all/deactivate | Deactivate email forwarding catch all
[**EmailForwardingsDeactivate**](EmailForwardingApi.md#EmailForwardingsDeactivate) | **Post** /domains/{domain}/email-forwardings/deactivate | Deactivate email forwarding
[**EmailForwardingsRecipientsDestroy**](EmailForwardingApi.md#EmailForwardingsRecipientsDestroy) | **Delete** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId} | Delete a recipient
[**EmailForwardingsRecipientsIndex**](EmailForwardingApi.md#EmailForwardingsRecipientsIndex) | **Get** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients | List recipients of an email forwarding
[**EmailForwardingsRecipientsResendVerification**](EmailForwardingApi.md#EmailForwardingsRecipientsResendVerification) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/resend-verification | Resend Verification
[**EmailForwardingsRecipientsSetDefault**](EmailForwardingApi.md#EmailForwardingsRecipientsSetDefault) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/set-default | Set default recipient
[**EmailForwardingsRecipientsStore**](EmailForwardingApi.md#EmailForwardingsRecipientsStore) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients | Create new recipient
[**EmailForwardingsRecipientsVerify**](EmailForwardingApi.md#EmailForwardingsRecipientsVerify) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/verify | Verify recipient
[**EmailForwardingsStats**](EmailForwardingApi.md#EmailForwardingsStats) | **Get** /domains/{domain}/email-forwardings/stats | Show stats of domain&#39;s email forwarding



## EmailForwardingsActivate

> MessageResponse EmailForwardingsActivate(ctx, domain).Execute()

Activate email forwarding

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
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsActivate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsActivate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsActivateRequest struct via the builder pattern


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


## EmailForwardingsAliasesDestroy

> MessageResponse EmailForwardingsAliasesDestroy(ctx, domain, emailForwardingId, emailForwardingAliasId).Execute()

Delete an alias

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingAliasId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding alias id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsAliasesDestroy(context.Background(), domain, emailForwardingId, emailForwardingAliasId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsAliasesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsAliasesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingAliasId** | **string** | Email forwarding alias id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesDestroyRequest struct via the builder pattern


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


## EmailForwardingsAliasesIndex

> EmailForwardingAliasesListData EmailForwardingsAliasesIndex(ctx, domain, emailForwardingId).PerPage(perPage).Page(page).Execute()

List of email forwarding aliases for given domain

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsAliasesIndex(context.Background(), domain, emailForwardingId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsAliasesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesIndex`: EmailForwardingAliasesListData
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsAliasesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**EmailForwardingAliasesListData**](EmailForwardingAliasesListData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## EmailForwardingsAliasesStore

> EmailForwardingsAliasesStore201Response EmailForwardingsAliasesStore(ctx, domain, emailForwardingId).EmailForwardingAliasesStore(emailForwardingAliasesStore).Execute()

Create new alias

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingAliasesStore := *openapiclient.NewEmailForwardingAliasesStore("LocalPart_example", []string{"Recipients_example"}) // EmailForwardingAliasesStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsAliasesStore(context.Background(), domain, emailForwardingId).EmailForwardingAliasesStore(emailForwardingAliasesStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsAliasesStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesStore`: EmailForwardingsAliasesStore201Response
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsAliasesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailForwardingAliasesStore** | [**EmailForwardingAliasesStore**](EmailForwardingAliasesStore.md) |  | 

### Return type

[**EmailForwardingsAliasesStore201Response**](EmailForwardingsAliasesStore201Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## EmailForwardingsAliasesToggleActivation

> MessageResponse EmailForwardingsAliasesToggleActivation(ctx, domain, emailForwardingId, emailForwardingAliasId).EmailForwardingAliasesToggleActivation(emailForwardingAliasesToggleActivation).Execute()

Toggle alias activation

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingAliasId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding alias id
    emailForwardingAliasesToggleActivation := *openapiclient.NewEmailForwardingAliasesToggleActivation(false) // EmailForwardingAliasesToggleActivation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsAliasesToggleActivation(context.Background(), domain, emailForwardingId, emailForwardingAliasId).EmailForwardingAliasesToggleActivation(emailForwardingAliasesToggleActivation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsAliasesToggleActivation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesToggleActivation`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsAliasesToggleActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingAliasId** | **string** | Email forwarding alias id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesToggleActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **emailForwardingAliasesToggleActivation** | [**EmailForwardingAliasesToggleActivation**](EmailForwardingAliasesToggleActivation.md) |  | 

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


## EmailForwardingsAliasesUpdateRecipients

> MessageResponse EmailForwardingsAliasesUpdateRecipients(ctx, domain, emailForwardingId, emailForwardingAliasId).EmailForwardingAliasesRecipients(emailForwardingAliasesRecipients).Execute()

Update alias recipients

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingAliasId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding alias id
    emailForwardingAliasesRecipients := *openapiclient.NewEmailForwardingAliasesRecipients([]string{"Recipients_example"}) // EmailForwardingAliasesRecipients |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsAliasesUpdateRecipients(context.Background(), domain, emailForwardingId, emailForwardingAliasId).EmailForwardingAliasesRecipients(emailForwardingAliasesRecipients).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsAliasesUpdateRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesUpdateRecipients`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsAliasesUpdateRecipients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingAliasId** | **string** | Email forwarding alias id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesUpdateRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **emailForwardingAliasesRecipients** | [**EmailForwardingAliasesRecipients**](EmailForwardingAliasesRecipients.md) |  | 

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


## EmailForwardingsCatchAllActivate

> MessageResponse EmailForwardingsCatchAllActivate(ctx, domain).Execute()

Activate email forwarding catch all

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
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsCatchAllActivate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsCatchAllActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsCatchAllActivate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsCatchAllActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsCatchAllActivateRequest struct via the builder pattern


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


## EmailForwardingsCatchAllDeactivate

> MessageResponse EmailForwardingsCatchAllDeactivate(ctx, domain).Execute()

Deactivate email forwarding catch all

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
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsCatchAllDeactivate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsCatchAllDeactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsCatchAllDeactivate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsCatchAllDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsCatchAllDeactivateRequest struct via the builder pattern


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


## EmailForwardingsDeactivate

> MessageResponse EmailForwardingsDeactivate(ctx, domain).Execute()

Deactivate email forwarding

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
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsDeactivate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsDeactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsDeactivate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsDeactivateRequest struct via the builder pattern


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


## EmailForwardingsRecipientsDestroy

> MessageResponse EmailForwardingsRecipientsDestroy(ctx, domain, emailForwardingId, emailForwardingRecipientId).Execute()

Delete a recipient

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding recipient id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsRecipientsDestroy(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsRecipientsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsRecipientsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingRecipientId** | **string** | Email forwarding recipient id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsDestroyRequest struct via the builder pattern


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


## EmailForwardingsRecipientsIndex

> EmailForwardingRecipientsListData EmailForwardingsRecipientsIndex(ctx, domain, emailForwardingId).PerPage(perPage).Page(page).Execute()

List recipients of an email forwarding

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsRecipientsIndex(context.Background(), domain, emailForwardingId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsRecipientsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsIndex`: EmailForwardingRecipientsListData
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsRecipientsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**EmailForwardingRecipientsListData**](EmailForwardingRecipientsListData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## EmailForwardingsRecipientsResendVerification

> MessageResponse EmailForwardingsRecipientsResendVerification(ctx, domain, emailForwardingId, emailForwardingRecipientId).Execute()

Resend Verification

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding recipient id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsRecipientsResendVerification(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsRecipientsResendVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsResendVerification`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsRecipientsResendVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingRecipientId** | **string** | Email forwarding recipient id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsResendVerificationRequest struct via the builder pattern


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


## EmailForwardingsRecipientsSetDefault

> MessageResponse EmailForwardingsRecipientsSetDefault(ctx, domain, emailForwardingId, emailForwardingRecipientId).Execute()

Set default recipient

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding recipient id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsRecipientsSetDefault(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsRecipientsSetDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsSetDefault`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsRecipientsSetDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingRecipientId** | **string** | Email forwarding recipient id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsSetDefaultRequest struct via the builder pattern


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


## EmailForwardingsRecipientsStore

> EmailForwardingsRecipientsStore201Response EmailForwardingsRecipientsStore(ctx, domain, emailForwardingId).EmailForwardingRecipientsStore(emailForwardingRecipientsStore).Execute()

Create new recipient

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientsStore := *openapiclient.NewEmailForwardingRecipientsStore("Email_example") // EmailForwardingRecipientsStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsRecipientsStore(context.Background(), domain, emailForwardingId).EmailForwardingRecipientsStore(emailForwardingRecipientsStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsRecipientsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsStore`: EmailForwardingsRecipientsStore201Response
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsRecipientsStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailForwardingRecipientsStore** | [**EmailForwardingRecipientsStore**](EmailForwardingRecipientsStore.md) |  | 

### Return type

[**EmailForwardingsRecipientsStore201Response**](EmailForwardingsRecipientsStore201Response.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)


## EmailForwardingsRecipientsVerify

> MessageResponse EmailForwardingsRecipientsVerify(ctx, domain, emailForwardingId, emailForwardingRecipientId).EmailForwardingRecipientsVerify(emailForwardingRecipientsVerify).Execute()

Verify recipient

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
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding recipient id
    emailForwardingRecipientsVerify := *openapiclient.NewEmailForwardingRecipientsVerify("Code_example") // EmailForwardingRecipientsVerify |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsRecipientsVerify(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).EmailForwardingRecipientsVerify(emailForwardingRecipientsVerify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsRecipientsVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsVerify`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsRecipientsVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingRecipientId** | **string** | Email forwarding recipient id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **emailForwardingRecipientsVerify** | [**EmailForwardingRecipientsVerify**](EmailForwardingRecipientsVerify.md) |  | 

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


## EmailForwardingsStats

> EmailForwardingStatsData EmailForwardingsStats(ctx, domain).Execute()

Show stats of domain's email forwarding

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
    resp, r, err := apiClient.EmailForwardingApi.EmailForwardingsStats(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingApi.EmailForwardingsStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsStats`: EmailForwardingStatsData
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingApi.EmailForwardingsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailForwardingStatsData**](EmailForwardingStatsData.md)

### Authorization

[ApiKey](HOW-TO.md#ApiKey), [UserToken](HOW-TO.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints)
[[Back to Model list]](HOW-TO.md#documentation-for-models)
[[Back to README]](HOW-TO.md)

