# Changelog

## [Unreleased](https://github.com/openfga/go-sdk/compare/v0.7.3...HEAD)

- feat: add a generic API Executor `fgaClient.GetAPIExecutor()` to allow calling any OpenFGA API method. See [Calling Other Endpoints](./README.md#calling-other-endpoints) for more.
- feat: add generic `ToPtr[T any](v T) *T` function for creating pointers to any type
- deprecation: `PtrBool`, `PtrInt`, `PtrInt32`, `PtrInt64`, `PtrFloat32`, `PtrFloat64`, `PtrString`, and `PtrTime` are now deprecated in favor of the generic `ToPtr` function
- feat: add a top-level makefile in go-sdk to simplify running tests and linters: (#250)
- feat: add support for StreamedListObjects endpoint (#252)

## v0.7.3

### [0.7.3](https://github.com/openfga/go-sdk/compare/v0.7.2...v0.7.3) (2025-10-08)

- feat: add support for custom headers per request. See [documentation](https://github.com/openfga/go-sdk#custom-headers).
- feat: add support for conflict options for Write operations**: (#229)
  The client now supports setting `Conflict` on `ClientWriteOptions` to control behavior when writing duplicate tuples or deleting non-existent tuples. This feature requires OpenFGA server [v1.10.0](https://github.com/openfga/openfga/releases/tag/v1.10.0) or later.
  See [Conflict Options for Write Operations](./README.md#conflict-options-for-write-operations) for more.

## v0.7.2

### [0.7.2](https://github.com/openfga/go-sdk/compare/v0.7.1...v0.7.2) (2025-09-15)

- feat: add contextual tuples support in Expand requests (https://github.com/openfga/sdk-generator/pull/547) - thanks @SoulPancake
- feat: Support passing name filter to ListStores (#186, #213) - thanks @Oscmage!
- fix: 5xx errors were not being properly retried (#204) - thanks @maxlegault for reporting

## v0.7.1

### [0.7.1](https://github.com/openfga/go-sdk/compare/v0.7.0...v0.7.1) (2025-04-07)

- fix: resolves issue where SDK could stall when MaxParallelRequests * MaxBatchSize < body.Checks length
- fix: resolves issue where private singleBatchCheck method in SdkClient prevented mocking

## v0.7.0

### [0.7.0](https://github.com/openfga/go-sdk/compare/v0.6.5...v0.7.0) (2025-04-02)

- feat: fix and improve retries and rate limit handling. (#176)
  The SDK now retries on network errors and the default retry handling has been fixed
  for both the calls to the OpenFGA API and the API Token Issuer for those using ClientCredentials
  The SDK now also respects the rate limit headers (`Retry-After`) returned by the server and will retry the request after the specified time.
  If the header is not sent or on network errors, it will fall back to exponential backoff.
- feat: retry on network errors when calling the token issuer (#182)
- feat: add support for server-side BatchCheck (#187)
- fix: use defaults when transaction options were only partially set (#183)
- chore: log retry attempts when debug mode is enabled (#182)

[!WARNING]
BREAKING CHANGES:
This release contains a breaking change around it's handling of `BatchCheck`
- The new `BatchCheck` requires OpenFGA [v1.8.0+](https://github.com/openfga/openfga/releases/tag/v1.8.0) server.
- The existing `BatchCheck` method has been renamed to `ClientBatchCheck`. The existing `BatchCheckResponse` has been renamed to `ClientBatchCheckResponse`.

NOTE:
This release was previously released as `v0.6.6`, but has been re-released as `v0.7.0` due to the breaking changes.

## v0.6.6

### [0.6.6](https://github.com/openfga/go-sdk/compare/v0.6.5...v0.6.6) (2025-04-02)
[!WARNING]
BREAKING CHANGES: This has been re-released as [v0.7.0](https://github.com/openfga/go-sdk/releases/tag/v0.7.0) due to breaking changes around `BatchCheck`.

## v0.6.5

### [0.6.5](https://github.com/openfga/go-sdk/compare/v0.6.4...v0.6.5) (2025-02-06)

- feat: support assertions with context and contextual tuples (#169)
- fix: populate start_time in request (#166)

## v0.6.4

### [0.6.4](https://github.com/openfga/go-sdk/compare/v0.6.3...v0.6.4) (2025-01-29)

- feat: add support for `start_time` parameter in `ReadChanges` endpoint (#158)
- fix: correctly set request level storeId in non-transactional write (#162)
- fix: api client should set default telemetry if not specified (#160)
- docs: replace readable names with uuid (#146) - thanks @sccalabr 
- fix: support marshaling client.ClientWriteResponse (#145) - thanks @Fedot-Compot
- fix: update client interface with store and model getter/setter (#155)
- fix: api client should set default telemetry if not specified (#160) 
- fix: Do not ignore request level storeId in non-transactional write (#162)

## v0.6.3

### [0.6.3](https://github.com/openfga/go-sdk/compare/v0.6.2...v0.6.3) (2024-10-22)

- fix: fix metrics data race issues (#139)

## v0.6.2

### [0.6.2](https://github.com/openfga/go-sdkk/compare/v0.6.1...v0.6.2) (2024-10-21)

- fix: fix batch check consistency (#131)
- fix: fix data race on TelemetryInstances (#136) - thanks @Kryvchun!

NOTE: `TelemetryInstances` in `telemetry.go` has been deprecated, as its usage is intended to be internal. It will be removed in a future release.

## v0.6.1

### [0.6.1](https://github.com/openfga/go-sdk/compare/v0.6.0...v0.6.1) (2024-09-23)

- refactor(OpenTelemetry): move configuration API into public package namespace (#122)
- docs(OpenTelemetry): initial documentation and example (#123)

## v0.6.0

### [0.6.0](https://github.com/openfga/go-sdk/compare/v0.5.0...v0.6.0) (2024-08-29)

- feat: support OpenTelemetry metrics reporting (#115)
- feat!: support for sending the consistency parameter to the read, check, list users, list objects, and expand endpoints (#117)
- chore(docs): update stale README (#113) - thanks @Code2Life

BREAKING CHANGE:

When the generator converts enums in the open API definition, by default it removes the type prefix. For example, `TYPE_NAME_UNSPECIFIED` is converted to a const named `UNSPECIFIED`. This leads to potential collisions with other enums, and as the consistency type is a new enum, we finally got a collision (was just a matter of time).

The fix for this is to specify `"enumClassPrefix": true` in the generation config. This will then include the class name on the const name, which resoles collision issues. This means any enum value, such as `INT` now becomes `TYPENAME_INT`. The main impact of this is the `TypeName` consts and error codes. The fix is to add the class name prefix as discussed above.

## v0.5.0

### [0.5.0](https://github.com/openfga/go-sdk/compare/v0.4.0...v0.5.0) (2024-06-14)
- fix: correctly set HTTPClient - thanks @wonyx
- chore!: remove excluded users from ListUsers response

BREAKING CHANGE:

This version removes the `ExcludedUsers` field from the `ListUsersResponse` and `ClientListUsersResponse` structs,
for more details see the [associated API change](https://github.com/openfga/api/pull/171).

## v0.4.0

### [0.4.0](https://github.com/openfga/go-sdk/compare/v0.3.7...v0.4.0) (2024-05-30)
- feat!: remove store ID from API config, allow store ID override per-request (see README for additional documentation and examples)
- fix: only retry on client credential requests that are 429 or 5x

BREAKING CHANGE:

This version removes the `StoreId` from the API client configuration. Instead, the `StoreId` parameter 
must now be passed to each of the API methods that require a store ID. 

**If you are using `api_open_fga.go` directly, you will now need to pass the `StoreId` parameter.**

## v0.3.7

### [0.3.7](https://github.com/openfga/go-sdk/compare/v0.3.6...v0.3.7) (2024-05-08)
- feat: Add MaxParallelRequests option in ListRelations (#93) - thanks @gurleensethi
- chore: lower required go version to 1.21 (fixes #94)

## v0.3.6

### [0.3.6](https://github.com/openfga/go-sdk/compare/v0.3.5...v0.3.6) (2024-04-30)

- feat: support the [ListUsers](https://github.com/openfga/rfcs/blob/main/20231214-listUsers-api.md) endpoint (#81)
- fix: do not call ReadAuthorizationModel on BatchCheck or non-Transactional Write (#78)
- chore: fix typos in the readme (#91) - thanks @balaji-dongare
- chore!: raise required go version to 1.21.9

## v0.3.5

### [0.3.5](https://github.com/openfga/go-sdk/compare/v0.3.4...v0.3.5) (2024-02-13)

- fix: don't escape HTML characters in conditions when marshalling a model

## v0.3.4

### [0.3.4](https://github.com/openfga/go-sdk/compare/v0.3.3...v0.3.4) (2024-01-22)

- feat: configurable client credentials token url - thanks @le-yams
- fix: WriteAuthorizationModel was not passing conditions to API

## v0.3.3

### [0.3.3](https://github.com/openfga/go-sdk/compare/v0.3.2...v0.3.3) (2023-12-21)

- fix: WriteAuthorizationModel was not passing conditions to API
- chore: add [example project](./example)

## v0.3.2

### [0.3.2](https://github.com/openfga/go-sdk/compare/v0.3.1...v0.3.2) (2023-12-20)

- fix: ListObjects was not passing context to API
- chore: downgrade target go version to 1.20

## v0.3.1

### [0.3.1](https://github.com/openfga/go-sdk/compare/v0.3.0...v0.3.1) (2023-12-19)

- feat: oauth2 client credentials support (#62), thanks @le-yams
- fix: remove canonical import path from oauth2 packages (#64), thanks @bketelsen

## v0.3.0

### [0.3.0](https://github.com/openfga/go-sdk/compare/v0.2.3...v0.3.0) (2023-12-11)

- feat!: initial support for [conditions](https://openfga.dev/blog/conditional-tuples-announcement)
- feat: support specifying a port and path for the API (You can now set the `ApiUrl` to something like: `https://api.fga.exampleL8080/some_path`)
- fix: resolve a bug in `NewCredentials` (#60) - thanks @harper
- chore!: use latest API interfaces
- chore: dependency updates


BREAKING CHANGES:
Note: This release comes with substantial breaking changes, especially to the interfaces due to the protobuf changes in the last release.

While the http interfaces did not break (you can still use `v0.2.3` SDK with a `v1.3.8+` server),
the grpc interface did and this caused a few changes in the interfaces of the SDK.

You will have to modify some parts of your code, but we hope this will be to the better as a lot of the parameters are now correctly marked as required,
and so the Pointer-to-String conversion is no longer needed.

Some of the changes to expect:

* When initializing a client, please use `ApiUrl`. The separate `ApiScheme` and `ApiHost` fields have been deprecated
```go
fgaClient, err := NewSdkClient(&ClientConfiguration{
    ApiUrl:  os.Getenv("FGA_API_URL"), // required, e.g. https://api.fga.example
    StoreId: os.Getenv("FGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
    AuthorizationModelId: os.Getenv("FGA_MODEL_ID"), // optional, recommended to be set for production
})
```
- When initializing a client, `AuthorizationModelId` is no longer a pointer, and you can just pass the string directly
- The `OpenFgaClient` now has methods to get and set the model ID `GetAuthorizationModelId` and `SetAuthorizationModelId`
- The following request interfaces changed:
    - `CheckRequest`: the `TupleKey` field is now of interface `CheckRequestTupleKey`, you can also now pass in `Context`
    - `ExpandRequest`: the `TupleKey` field is now of interface `ExpandRequestTupleKey`
    - `ReadRequest`: the `TupleKey` field is now of interface `ReadRequestTupleKey`
    - `WriteRequest`: now takes `WriteRequestWrites` and `WriteRequestDeletes`
    - And more
- The following interfaces had fields that were pointers are are now the direct value:
    - `CreateStoreResponse`
    - `GetStoreResponse`
    - `ListStoresResponse`
    - `ListObjectsResponse`
    - `ReadChangesResponse`
    - `ReadResponse`
    - `AuthorizationModel` and several interfaces under it
    - And more

## v0.2.3

### [0.2.3](https://github.com/openfga/go-sdk/compare/v0.2.2...v0.2.3) (2023-10-13)

- fix: allow setting user agent
- fix(client): resolve null pointer exceptions when getting auth model id
- fix(client): allow read to contain empty fields
- fix(client): require auth model id and store id to be ulids
- fix(client): resolve cases where req options was not respected
- fix: add retry logic to oauth
- chore: target go1.21.3 and upgrade dependencies

## v0.2.2

### [0.2.2](https://github.com/openfga/go-sdk/compare/v0.2.1...v0.2.2) (2023-04-21)

- feat(client): add OpenFgaClient wrapper see [docs](https://github.com/openfga/go-sdk/tree/main#readme), see the `v0.2.1` docs for [the OpenFgaApi docs](https://github.com/openfga/go-sdk/tree/v0.2.1#readme)
- feat(client): implement `BatchCheck` to check multiple tuples in parallel
- feat(client): implement `ListRelations` to check in one call whether a user has multiple relations to an objects
- feat(client): add support for a non-transactional `Write`
- chore(config): bump default max retries to `15`
- fix(config)!: make the capitalization of the json equivalent of the configuration consistent
- fix: retry on 5xx errors

## v0.2.1

### [0.2.1](https://github.com/openfga/go-sdk/compare/v0.2.0...v0.2.1) (2023-01-17)

- chore(deps): upgrade `golang.org/x/net` dependency

## v0.2.0

### [0.2.0](https://github.com/openfga/go-sdk/compare/v0.1.1...v0.2.0) (2022-12-14)

Updated to include support for [OpenFGA 0.3.0](https://github.com/openfga/openfga/releases/tag/v0.3.0)

Changes:
- [BREAKING] feat(list-objects)!: response has been changed to include the object type
    e.g. response that was `{"object_ids":["roadmap"]}`, will now be `{"objects":["document:0192ab2a-d83f-756d-9397-c5ed9f3cb69a"]}`

Fixes:
- [BREAKING] fix(models): update interfaces that had incorrectly optional fields to make them required

Chore:
- chore(deps): update dev dependencies

## v0.1.1

### [0.1.1](https://github.com/openfga/go-sdk/compare/v0.1.0...v0.1.1) (2022-09-30)

- chore(deps): upgrade dependencies - dependency updates were accidentally reverted in v0.1.0 release

## v0.1.0

### [0.1.0](https://github.com/openfga/go-sdk/compare/v0.0.3...v0.1.0) (2022-09-29)

- BREAKING: exported interface `TypeDefinitions` is now `WriteAuthorizationModelRequest`
    This is only a breaking change on the SDK, not the API. It was changed to conform to the proto changes in [openfga/api](https://github.com/openfga/api/pull/27).
- chore(deps): upgrade dependencies

## v0.0.3

### [0.0.3](https://github.com/openfga/go-sdk/compare/v0.0.2...v0.0.3) (2022-09-07)

- Fix incorrectly applying client_credentials validation to api_token cred method [openfga/sdk-generator#21](https://github.com/openfga/sdk-generator/pull/21)
- Target go 1.19
- Bump golang.org/x/net
- Use [govulncheck](https://go.dev/blog/vuln) in CI to check for issues

## v0.0.2

### [0.0.2](https://github.com/openfga/go-sdk/compare/v0.0.1...v0.0.2) (2022-08-15)

Support for [ListObjects API]](https://openfga.dev/api/service#/Relationship%20Queries/ListObjects)

You call the API and receive the list of object ids from a particular type that the user has a certain relation with.

For example, to find the list of documents that Anne can read:

```golang
body := openfga.ListObjectsRequest{
    AuthorizationModelId: PtrString(""),
    User:                 PtrString("anne"),
    Relation:             PtrString("can_view"),
    Type:                 PtrString("document"),
}
data, response, err := apiClient.OpenFgaApi.ListObjects(context.Background()).Body(body).Execute()

// response.object_ids = ["roadmap"]
```

## v0.0.1

### [0.0.1](https://github.com/openfga/go-sdk/releases/tag/v0.0.1) (2022-06-16)

Initial OpenFGA Go SDK release
- Support for [OpenFGA](https://github.com/openfga/openfga) API
  - CRUD stores
  - Create, read & list authorization models
  - Writing and Reading Tuples
  - Checking authorization
  - Using Expand to understand why access was granted
