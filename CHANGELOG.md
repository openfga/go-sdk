# Changelog

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
