# ListUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**Object** | [**FgaObject**](FgaObject.md) |  | 
**Relation** | **string** |  | 
**UserFilters** | [**[]UserTypeFilter**](UserTypeFilter.md) | The type of results returned. Only accepts exactly one value. | 
**ContextualTuples** | Pointer to [**[]TupleKey**](TupleKey.md) |  | [optional] 
**Context** | Pointer to **map[string]interface{}** | Additional request context that will be used to evaluate any ABAC conditions encountered in the query evaluation. | [optional] 

## Methods

### NewListUsersRequest

`func NewListUsersRequest(object FgaObject, relation string, userFilters []UserTypeFilter, ) *ListUsersRequest`

NewListUsersRequest instantiates a new ListUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersRequestWithDefaults

`func NewListUsersRequestWithDefaults() *ListUsersRequest`

NewListUsersRequestWithDefaults instantiates a new ListUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationModelId

`func (o *ListUsersRequest) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *ListUsersRequest) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *ListUsersRequest) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *ListUsersRequest) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetObject

`func (o *ListUsersRequest) GetObject() FgaObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ListUsersRequest) GetObjectOk() (*FgaObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ListUsersRequest) SetObject(v FgaObject)`

SetObject sets Object field to given value.


### GetRelation

`func (o *ListUsersRequest) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *ListUsersRequest) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *ListUsersRequest) SetRelation(v string)`

SetRelation sets Relation field to given value.


### GetUserFilters

`func (o *ListUsersRequest) GetUserFilters() []UserTypeFilter`

GetUserFilters returns the UserFilters field if non-nil, zero value otherwise.

### GetUserFiltersOk

`func (o *ListUsersRequest) GetUserFiltersOk() (*[]UserTypeFilter, bool)`

GetUserFiltersOk returns a tuple with the UserFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFilters

`func (o *ListUsersRequest) SetUserFilters(v []UserTypeFilter)`

SetUserFilters sets UserFilters field to given value.


### GetContextualTuples

`func (o *ListUsersRequest) GetContextualTuples() []TupleKey`

GetContextualTuples returns the ContextualTuples field if non-nil, zero value otherwise.

### GetContextualTuplesOk

`func (o *ListUsersRequest) GetContextualTuplesOk() (*[]TupleKey, bool)`

GetContextualTuplesOk returns a tuple with the ContextualTuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualTuples

`func (o *ListUsersRequest) SetContextualTuples(v []TupleKey)`

SetContextualTuples sets ContextualTuples field to given value.

### HasContextualTuples

`func (o *ListUsersRequest) HasContextualTuples() bool`

HasContextualTuples returns a boolean if a field has been set.

### GetContext

`func (o *ListUsersRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ListUsersRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ListUsersRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ListUsersRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


