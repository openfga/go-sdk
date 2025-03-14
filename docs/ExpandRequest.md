# ExpandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | [**ExpandRequestTupleKey**](ExpandRequestTupleKey.md) |  | 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**Consistency** | Pointer to [**ConsistencyPreference**](ConsistencyPreference.md) |  | [optional] [default to CONSISTENCYPREFERENCE_UNSPECIFIED]
**ContextualTuples** | Pointer to [**ContextualTupleKeys**](ContextualTupleKeys.md) |  | [optional] 

## Methods

### NewExpandRequest

`func NewExpandRequest(tupleKey ExpandRequestTupleKey, ) *ExpandRequest`

NewExpandRequest instantiates a new ExpandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandRequestWithDefaults

`func NewExpandRequestWithDefaults() *ExpandRequest`

NewExpandRequestWithDefaults instantiates a new ExpandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *ExpandRequest) GetTupleKey() ExpandRequestTupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *ExpandRequest) GetTupleKeyOk() (*ExpandRequestTupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *ExpandRequest) SetTupleKey(v ExpandRequestTupleKey)`

SetTupleKey sets TupleKey field to given value.


### GetAuthorizationModelId

`func (o *ExpandRequest) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *ExpandRequest) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *ExpandRequest) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *ExpandRequest) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetConsistency

`func (o *ExpandRequest) GetConsistency() ConsistencyPreference`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *ExpandRequest) GetConsistencyOk() (*ConsistencyPreference, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *ExpandRequest) SetConsistency(v ConsistencyPreference)`

SetConsistency sets Consistency field to given value.

### HasConsistency

`func (o *ExpandRequest) HasConsistency() bool`

HasConsistency returns a boolean if a field has been set.

### GetContextualTuples

`func (o *ExpandRequest) GetContextualTuples() ContextualTupleKeys`

GetContextualTuples returns the ContextualTuples field if non-nil, zero value otherwise.

### GetContextualTuplesOk

`func (o *ExpandRequest) GetContextualTuplesOk() (*ContextualTupleKeys, bool)`

GetContextualTuplesOk returns a tuple with the ContextualTuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualTuples

`func (o *ExpandRequest) SetContextualTuples(v ContextualTupleKeys)`

SetContextualTuples sets ContextualTuples field to given value.

### HasContextualTuples

`func (o *ExpandRequest) HasContextualTuples() bool`

HasContextualTuples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


