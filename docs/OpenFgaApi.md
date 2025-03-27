# \OpenFgaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCheck**](OpenFgaApi.md#BatchCheck) | **Post** /stores/{store_id}/batch-check | Send a list of &#x60;check&#x60; operations in a single request
[**Check**](OpenFgaApi.md#Check) | **Post** /stores/{store_id}/check | Check whether a user is authorized to access an object
[**CreateStore**](OpenFgaApi.md#CreateStore) | **Post** /stores | Create a store
[**DeleteStore**](OpenFgaApi.md#DeleteStore) | **Delete** /stores/{store_id} | Delete a store
[**Expand**](OpenFgaApi.md#Expand) | **Post** /stores/{store_id}/expand | Expand all relationships in userset tree format, and following userset rewrite rules.  Useful to reason about and debug a certain relationship
[**GetStore**](OpenFgaApi.md#GetStore) | **Get** /stores/{store_id} | Get a store
[**ListObjects**](OpenFgaApi.md#ListObjects) | **Post** /stores/{store_id}/list-objects | List all objects of the given type that the user has a relation with
[**ListStores**](OpenFgaApi.md#ListStores) | **Get** /stores | List all stores
[**ListUsers**](OpenFgaApi.md#ListUsers) | **Post** /stores/{store_id}/list-users | List the users matching the provided filter who have a certain relation to a particular type.
[**Read**](OpenFgaApi.md#Read) | **Post** /stores/{store_id}/read | Get tuples from the store that matches a query, without following userset rewrite rules
[**ReadAssertions**](OpenFgaApi.md#ReadAssertions) | **Get** /stores/{store_id}/assertions/{authorization_model_id} | Read assertions for an authorization model ID
[**ReadAuthorizationModel**](OpenFgaApi.md#ReadAuthorizationModel) | **Get** /stores/{store_id}/authorization-models/{id} | Return a particular version of an authorization model
[**ReadAuthorizationModels**](OpenFgaApi.md#ReadAuthorizationModels) | **Get** /stores/{store_id}/authorization-models | Return all the authorization models for a particular store
[**ReadChanges**](OpenFgaApi.md#ReadChanges) | **Get** /stores/{store_id}/changes | Return a list of all the tuple changes
[**Write**](OpenFgaApi.md#Write) | **Post** /stores/{store_id}/write | Add or delete tuples from the store
[**WriteAssertions**](OpenFgaApi.md#WriteAssertions) | **Put** /stores/{store_id}/assertions/{authorization_model_id} | Upsert assertions for an authorization model ID
[**WriteAuthorizationModel**](OpenFgaApi.md#WriteAuthorizationModel) | **Post** /stores/{store_id}/authorization-models | Create a new authorization model



## BatchCheck

> BatchCheckResponse BatchCheck(ctx).Body(body).Execute()

Send a list of `check` operations in a single request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    body := *openapiclient.NewBatchCheckRequest([]openapiclient.BatchCheckItem{*openapiclient.NewBatchCheckItem(*openapiclient.NewCheckRequestTupleKey("user:anne", "reader", "document:2021-budget"), "1cd93d8c-8e45-43c6-9a15-cbb3c7f394bc")}) // BatchCheckRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.BatchCheck(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.BatchCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `BatchCheck`: BatchCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.BatchCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**BatchCheckRequest**](BatchCheckRequest.md) |  | 

### Return type

[**BatchCheckResponse**](BatchCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Check

> CheckResponse Check(ctx).Body(body).Execute()

Check whether a user is authorized to access an object



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    body := *openapiclient.NewCheckRequest(*openapiclient.NewCheckRequestTupleKey("user:anne", "reader", "document:2021-budget")) // CheckRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.Check(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.Check``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `Check`: CheckResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.Check`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CheckRequest**](CheckRequest.md) |  | 

### Return type

[**CheckResponse**](CheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStore

> CreateStoreResponse CreateStore(ctx).Body(body).Execute()

Create a store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.CreateStore(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.CreateStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `CreateStore`: CreateStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.CreateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**CreateStoreRequest**](CreateStoreRequest.md) |  | 

### Return type

[**CreateStoreResponse**](CreateStoreResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStore

> DeleteStore(ctx).Execute()

Delete a store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.DeleteStore(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.DeleteStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Expand

> ExpandResponse Expand(ctx).Body(body).Execute()

Expand all relationships in userset tree format, and following userset rewrite rules.  Useful to reason about and debug a certain relationship



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    body := *openapiclient.NewExpandRequest(*openapiclient.NewExpandRequestTupleKey("reader", "document:2021-budget")) // ExpandRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.Expand(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.Expand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `Expand`: ExpandResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.Expand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiExpandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ExpandRequest**](ExpandRequest.md) |  | 

### Return type

[**ExpandResponse**](ExpandResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStore

> GetStoreResponse GetStore(ctx).Execute()

Get a store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.GetStore(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.GetStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `GetStore`: GetStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.GetStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GetStoreResponse**](GetStoreResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjects

> ListObjectsResponse ListObjects(ctx).Body(body).Execute()

List all objects of the given type that the user has a relation with



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    body := *openapiclient.NewListObjectsRequest("document", "reader", "user:anne") // ListObjectsRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ListObjects(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ListObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ListObjects`: ListObjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.ListObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ListObjectsRequest**](ListObjectsRequest.md) |  | 

### Return type

[**ListObjectsResponse**](ListObjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStores

> ListStoresResponse ListStores(ctx).PageSize(pageSize).ContinuationToken(continuationToken).Execute()

List all stores



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    continuationToken := "continuationToken_example" // string |  (optional)

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ListStores(context.Background()).PageSize(pageSize).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ListStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ListStores`: ListStoresResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.ListStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pageSize** | **int32** |  | 
**continuationToken** | **string** |  | 

### Return type

[**ListStoresResponse**](ListStoresResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsersResponse ListUsers(ctx).Body(body).Execute()

List the users matching the provided filter who have a certain relation to a particular type.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    body := *openapiclient.NewListUsersRequest(*openapiclient.NewFgaObject("document", "0bcdf6fa-a6aa-4730-a8eb-9cf172ff16d9"), "reader", []openapiclient.UserTypeFilter{*openapiclient.NewUserTypeFilter("group")}) // ListUsersRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ListUsers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ListUsers`: ListUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ListUsersRequest**](ListUsersRequest.md) |  | 

### Return type

[**ListUsersResponse**](ListUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> ReadResponse Read(ctx).Body(body).Execute()

Get tuples from the store that matches a query, without following userset rewrite rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    body := *openapiclient.NewReadRequest() // ReadRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.Read(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.Read``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `Read`: ReadResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ReadRequest**](ReadRequest.md) |  | 

### Return type

[**ReadResponse**](ReadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAssertions

> ReadAssertionsResponse ReadAssertions(ctx, authorizationModelId).Execute()

Read assertions for an authorization model ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    authorizationModelId := "authorizationModelId_example" // string | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ReadAssertions(context.Background(), authorizationModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ReadAssertions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadAssertions`: ReadAssertionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.ReadAssertions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationModelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ReadAssertionsResponse**](ReadAssertionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAuthorizationModel

> ReadAuthorizationModelResponse ReadAuthorizationModel(ctx, id).Execute()

Return a particular version of an authorization model



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    id := "id_example" // string | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ReadAuthorizationModel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ReadAuthorizationModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadAuthorizationModel`: ReadAuthorizationModelResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.ReadAuthorizationModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAuthorizationModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ReadAuthorizationModelResponse**](ReadAuthorizationModelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAuthorizationModels

> ReadAuthorizationModelsResponse ReadAuthorizationModels(ctx).PageSize(pageSize).ContinuationToken(continuationToken).Execute()

Return all the authorization models for a particular store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    pageSize := int32(56) // int32 |  (optional)
    continuationToken := "continuationToken_example" // string |  (optional)

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background()).PageSize(pageSize).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ReadAuthorizationModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadAuthorizationModels`: ReadAuthorizationModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.ReadAuthorizationModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiReadAuthorizationModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pageSize** | **int32** |  | 
**continuationToken** | **string** |  | 

### Return type

[**ReadAuthorizationModelsResponse**](ReadAuthorizationModelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadChanges

> ReadChangesResponse ReadChanges(ctx).Type_(type_).PageSize(pageSize).ContinuationToken(continuationToken).StartTime(startTime).Execute()

Return a list of all the tuple changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    type_ := "type__example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    continuationToken := "continuationToken_example" // string |  (optional)
    startTime := time.Now() // time.Time | Start date and time of changes to read. Format: ISO 8601 timestamp (e.g., 2022-01-01T00:00:00Z) If a continuation_token is provided along side start_time, the continuation_token will take precedence over start_time. (optional)

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ReadChanges(context.Background()).Type_(type_).PageSize(pageSize).ContinuationToken(continuationToken).StartTime(startTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ReadChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `ReadChanges`: ReadChangesResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.ReadChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiReadChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**type_** | **string** |  | 
**pageSize** | **int32** |  | 
**continuationToken** | **string** |  | 
**startTime** | **time.Time** | Start date and time of changes to read. Format: ISO 8601 timestamp (e.g., 2022-01-01T00:00:00Z) If a continuation_token is provided along side start_time, the continuation_token will take precedence over start_time. | 

### Return type

[**ReadChangesResponse**](ReadChangesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Write

> map[string]interface{} Write(ctx).Body(body).Execute()

Add or delete tuples from the store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    body := *openapiclient.NewWriteRequest() // WriteRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.Write(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.Write``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `Write`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.Write`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiWriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**WriteRequest**](WriteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteAssertions

> WriteAssertions(ctx, authorizationModelId).Body(body).Execute()

Upsert assertions for an authorization model ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    authorizationModelId := "authorizationModelId_example" // string | 
    body := *openapiclient.NewWriteAssertionsRequest([]openapiclient.Assertion{*openapiclient.NewAssertion(*openapiclient.NewAssertionTupleKey("document:2021-budget", "reader", "user:anne"), false)}) // WriteAssertionsRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.WriteAssertions(context.Background(), authorizationModelId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.WriteAssertions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorizationModelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWriteAssertionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**WriteAssertionsRequest**](WriteAssertionsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteAuthorizationModel

> WriteAuthorizationModelResponse WriteAuthorizationModel(ctx).Body(body).Execute()

Create a new authorization model



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openfga "github.com/openfga/go-sdk"
)

func main() {
    
    body := *openapiclient.NewWriteAuthorizationModelRequest([]openapiclient.TypeDefinition{*openapiclient.NewTypeDefinition("document")}, "SchemaVersion_example") // WriteAuthorizationModelRequest | 

    configuration, err := openfga.NewConfiguration(openfga.Configuration{
        ApiUrl:         os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
        StoreId:        os.Getenv("OPENFGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    })

    if err != nil {
    // .. Handle error
    }

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.WriteAuthorizationModel(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.WriteAuthorizationModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case FgaApiAuthenticationError:
            // Handle authentication error
        case FgaApiValidationError:
            // Handle parameter validation error
        case FgaApiNotFoundError:
            // Handle not found error
        case FgaApiInternalError:
            // Handle API internal error
        case FgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `WriteAuthorizationModel`: WriteAuthorizationModelResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.WriteAuthorizationModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Other parameters are passed through a pointer to a apiWriteAuthorizationModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**WriteAuthorizationModelRequest**](WriteAuthorizationModelRequest.md) |  | 

### Return type

[**WriteAuthorizationModelResponse**](WriteAuthorizationModelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

