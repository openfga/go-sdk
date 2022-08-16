# ListObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Relation** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**ContextualTuples** | Pointer to [**ContextualTupleKeys**](ContextualTupleKeys.md) |  | [optional] 

## Methods

### NewListObjectsRequest

`func NewListObjectsRequest() *ListObjectsRequest`

NewListObjectsRequest instantiates a new ListObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListObjectsRequestWithDefaults

`func NewListObjectsRequestWithDefaults() *ListObjectsRequest`

NewListObjectsRequestWithDefaults instantiates a new ListObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationModelId

`func (o *ListObjectsRequest) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *ListObjectsRequest) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *ListObjectsRequest) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *ListObjectsRequest) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetType

`func (o *ListObjectsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListObjectsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListObjectsRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListObjectsRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelation

`func (o *ListObjectsRequest) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *ListObjectsRequest) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *ListObjectsRequest) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *ListObjectsRequest) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetUser

`func (o *ListObjectsRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListObjectsRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListObjectsRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ListObjectsRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetContextualTuples

`func (o *ListObjectsRequest) GetContextualTuples() ContextualTupleKeys`

GetContextualTuples returns the ContextualTuples field if non-nil, zero value otherwise.

### GetContextualTuplesOk

`func (o *ListObjectsRequest) GetContextualTuplesOk() (*ContextualTupleKeys, bool)`

GetContextualTuplesOk returns a tuple with the ContextualTuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualTuples

`func (o *ListObjectsRequest) SetContextualTuples(v ContextualTupleKeys)`

SetContextualTuples sets ContextualTuples field to given value.

### HasContextualTuples

`func (o *ListObjectsRequest) HasContextualTuples() bool`

HasContextualTuples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


