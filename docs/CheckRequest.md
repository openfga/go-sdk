# CheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | [**CheckRequestTupleKey**](CheckRequestTupleKey.md) |  | 
**ContextualTuples** | Pointer to [**ContextualTupleKeys**](ContextualTupleKeys.md) |  | [optional] 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**Trace** | Pointer to **bool** | Defaults to false. Making it true has performance implications. | [optional] [readonly] 
**Context** | Pointer to **map[string]interface{}** | Additional request context that will be used to evaluate any ABAC conditions encountered in the query evaluation. | [optional] 
**Consistency** | Pointer to [**ConsistencyPreference**](ConsistencyPreference.md) |  | [optional] [default to CONSISTENCYPREFERENCE_UNSPECIFIED]

## Methods

### NewCheckRequest

`func NewCheckRequest(tupleKey CheckRequestTupleKey, ) *CheckRequest`

NewCheckRequest instantiates a new CheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckRequestWithDefaults

`func NewCheckRequestWithDefaults() *CheckRequest`

NewCheckRequestWithDefaults instantiates a new CheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *CheckRequest) GetTupleKey() CheckRequestTupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *CheckRequest) GetTupleKeyOk() (*CheckRequestTupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *CheckRequest) SetTupleKey(v CheckRequestTupleKey)`

SetTupleKey sets TupleKey field to given value.


### GetContextualTuples

`func (o *CheckRequest) GetContextualTuples() ContextualTupleKeys`

GetContextualTuples returns the ContextualTuples field if non-nil, zero value otherwise.

### GetContextualTuplesOk

`func (o *CheckRequest) GetContextualTuplesOk() (*ContextualTupleKeys, bool)`

GetContextualTuplesOk returns a tuple with the ContextualTuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualTuples

`func (o *CheckRequest) SetContextualTuples(v ContextualTupleKeys)`

SetContextualTuples sets ContextualTuples field to given value.

### HasContextualTuples

`func (o *CheckRequest) HasContextualTuples() bool`

HasContextualTuples returns a boolean if a field has been set.

### GetAuthorizationModelId

`func (o *CheckRequest) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *CheckRequest) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *CheckRequest) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *CheckRequest) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetTrace

`func (o *CheckRequest) GetTrace() bool`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *CheckRequest) GetTraceOk() (*bool, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *CheckRequest) SetTrace(v bool)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *CheckRequest) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetContext

`func (o *CheckRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CheckRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CheckRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *CheckRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetConsistency

`func (o *CheckRequest) GetConsistency() ConsistencyPreference`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *CheckRequest) GetConsistencyOk() (*ConsistencyPreference, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *CheckRequest) SetConsistency(v ConsistencyPreference)`

SetConsistency sets Consistency field to given value.

### HasConsistency

`func (o *CheckRequest) HasConsistency() bool`

HasConsistency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


