# \OpenFgaApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Check**](OpenFgaApi.md#Check) | **Post** /stores/{store_id}/check | Check whether a user is authorized to access an object
[**Expand**](OpenFgaApi.md#Expand) | **Post** /stores/{store_id}/expand | Expand all relationships in userset tree format, and following userset rewrite rules.  Useful to reason about and debug a certain relationship
[**Read**](OpenFgaApi.md#Read) | **Post** /stores/{store_id}/read | Get tuples from the store that matches a query, without following userset rewrite rules
[**ReadAssertions**](OpenFgaApi.md#ReadAssertions) | **Get** /stores/{store_id}/assertions/{authorization_model_id} | Read assertions for an authorization model ID
[**ReadAuthorizationModel**](OpenFgaApi.md#ReadAuthorizationModel) | **Get** /stores/{store_id}/authorization-models/{id} | Return a particular version of an authorization model
[**ReadAuthorizationModels**](OpenFgaApi.md#ReadAuthorizationModels) | **Get** /stores/{store_id}/authorization-models | Return all the authorization model IDs for a particular store
[**ReadChanges**](OpenFgaApi.md#ReadChanges) | **Get** /stores/{store_id}/changes | Return a list of all the tuple changes
[**Write**](OpenFgaApi.md#Write) | **Post** /stores/{store_id}/write | Add or delete tuples from the store
[**WriteAssertions**](OpenFgaApi.md#WriteAssertions) | **Put** /stores/{store_id}/assertions/{authorization_model_id} | Upsert assertions for an authorization model ID
[**WriteAuthorizationModel**](OpenFgaApi.md#WriteAuthorizationModel) | **Post** /stores/{store_id}/authorization-models | Create a new authorization model



## Check

> CheckResponse Check(ctx).Params(params).Execute()

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
    
    params := *openapiclient.NewCheckRequestParams() // CheckRequestParams | 

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.Check(context.Background()).Params(params).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.Check``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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
**params** | [**CheckRequestParams**](CheckRequestParams.md) |  | 

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


## Expand

> ExpandResponse Expand(ctx).Params(params).Execute()

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
    
    params := *openapiclient.NewExpandRequestParams() // ExpandRequestParams | 

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.Expand(context.Background()).Params(params).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.Expand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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
**params** | [**ExpandRequestParams**](ExpandRequestParams.md) |  | 

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


## Read

> ReadResponse Read(ctx).Params(params).Execute()

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
    
    params := *openapiclient.NewReadRequestParams() // ReadRequestParams | 

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.Read(context.Background()).Params(params).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.Read``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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
**params** | [**ReadRequestParams**](ReadRequestParams.md) |  | 

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

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ReadAssertions(context.Background(), authorizationModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ReadAssertions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ReadAuthorizationModel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ReadAuthorizationModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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

Return all the authorization model IDs for a particular store



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

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ReadAuthorizationModels(context.Background()).PageSize(pageSize).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ReadAuthorizationModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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

> ReadChangesResponse ReadChanges(ctx).Type_(type_).PageSize(pageSize).ContinuationToken(continuationToken).Execute()

Return a list of all the tuple changes



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
    
    type_ := "type__example" // string |  (optional)
    pageSize := int32(56) // int32 |  (optional)
    continuationToken := "continuationToken_example" // string |  (optional)

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.ReadChanges(context.Background()).Type_(type_).PageSize(pageSize).ContinuationToken(continuationToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.ReadChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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

> map[string]interface{} Write(ctx).Params(params).Execute()

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
    
    params := *openapiclient.NewWriteRequestParams() // WriteRequestParams | 

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.Write(context.Background()).Params(params).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.Write``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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
**params** | [**WriteRequestParams**](WriteRequestParams.md) |  | 

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

> map[string]interface{} WriteAssertions(ctx, authorizationModelId).Params(params).Execute()

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
    params := *openapiclient.NewWriteAssertionsRequestParams([]openapiclient.Assertion{*openapiclient.NewAssertion(false)}) // WriteAssertionsRequestParams | 

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.WriteAssertions(context.Background(), authorizationModelId).Params(params).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.WriteAssertions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
            // Exponential backoff in handling rate limit error
        default:
            // Handle unknown/undefined error
        }
    }
    // response from `WriteAssertions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpenFgaApi.WriteAssertions`: %v\n", resp)
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
**params** | [**WriteAssertionsRequestParams**](WriteAssertionsRequestParams.md) |  | 

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


## WriteAuthorizationModel

> WriteAuthorizationModelResponse WriteAuthorizationModel(ctx).TypeDefinitions(typeDefinitions).Execute()

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
    
    typeDefinitions := *openapiclient.NewTypeDefinitions() // TypeDefinitions | 

    configuration := openfga.NewConfiguration(UserConfiguration{
        StoreId:      os.Getenv("OPENFGA_STORE_ID"),
        ClientId:     os.Getenv("OPENFGA_CLIENT_ID"),
        ClientSecret: os.Getenv("OPENFGA_CLIENT_SECRET"),
        Environment:  os.Getenv("OPENFGA_ENVIRONMENT"),
    })

    apiClient := openfga.NewAPIClient(configuration)

    resp, r, err := apiClient.OpenFgaApi.WriteAuthorizationModel(context.Background()).TypeDefinitions(typeDefinitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenFgaApi.WriteAuthorizationModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        switch v := err.(type) {
        case OpenFgaApiAuthenticationError:
            // Handle authentication error
        case OpenFgaApiValidationError:
            // Handle parameter validation error
        case OpenFgaApiNotFoundError:
            // Handle not found error
        case OpenFgaApiInternalError:
            // Handle API internal error
        case OpenFgaApiRateLimitError:
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
**typeDefinitions** | [**TypeDefinitions**](TypeDefinitions.md) |  | 

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

