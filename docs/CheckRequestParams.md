# CheckRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | Pointer to [**TupleKey**](TupleKey.md) |  | [optional] 
**ContextualTuples** | Pointer to [**ContextualTupleKeys**](ContextualTupleKeys.md) |  | [optional] 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**Trace** | Pointer to **bool** | Defaults to false. Making it true has performance implications. | [optional] [readonly] 

## Methods

### NewCheckRequestParams

`func NewCheckRequestParams() *CheckRequestParams`

NewCheckRequestParams instantiates a new CheckRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckRequestParamsWithDefaults

`func NewCheckRequestParamsWithDefaults() *CheckRequestParams`

NewCheckRequestParamsWithDefaults instantiates a new CheckRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *CheckRequestParams) GetTupleKey() TupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *CheckRequestParams) GetTupleKeyOk() (*TupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *CheckRequestParams) SetTupleKey(v TupleKey)`

SetTupleKey sets TupleKey field to given value.

### HasTupleKey

`func (o *CheckRequestParams) HasTupleKey() bool`

HasTupleKey returns a boolean if a field has been set.

### GetContextualTuples

`func (o *CheckRequestParams) GetContextualTuples() ContextualTupleKeys`

GetContextualTuples returns the ContextualTuples field if non-nil, zero value otherwise.

### GetContextualTuplesOk

`func (o *CheckRequestParams) GetContextualTuplesOk() (*ContextualTupleKeys, bool)`

GetContextualTuplesOk returns a tuple with the ContextualTuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualTuples

`func (o *CheckRequestParams) SetContextualTuples(v ContextualTupleKeys)`

SetContextualTuples sets ContextualTuples field to given value.

### HasContextualTuples

`func (o *CheckRequestParams) HasContextualTuples() bool`

HasContextualTuples returns a boolean if a field has been set.

### GetAuthorizationModelId

`func (o *CheckRequestParams) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *CheckRequestParams) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *CheckRequestParams) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *CheckRequestParams) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetTrace

`func (o *CheckRequestParams) GetTrace() bool`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *CheckRequestParams) GetTraceOk() (*bool, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *CheckRequestParams) SetTrace(v bool)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *CheckRequestParams) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


