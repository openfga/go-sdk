# Changelog

## v0.3.1-go1.20

### [0.3.1-go1.20](https://github.com/openfga/go-sdk/compare/v0.3.1...v0.3.1-go1.20) (2023-12-19)

Same as v0.3.1, but with the target go version set to `1.20` (see https://github.com/lxc/incus/issues/315#issuecomment-1863382429)

## v0.3.1

### [0.3.1](https://github.com/openfga/go-sdk/compare/v0.3.0...v0.3.1) (2023-12-19)

- feat: oauth2 client credentials support (#62), thanks @le-yams
- fix: remove canonical import path from oauth2 packages (#64), thanks @bketelsen

## v0.3.0

### [0.3.0](https://github.com/openfga/go-sdk/compare/v0.2.3...v0.3.0) (2023-12-11)

- feat!: initial support for conditions
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
    AuthorizationModelId: os.Getenv("FGA_AUTHORIZATION_MODEL_ID"), // optional, recommended to be set for production
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
    e.g. response that was `{"object_ids":["roadmap"]}`, will now be `{"objects":["document:roadmap"]}`

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
