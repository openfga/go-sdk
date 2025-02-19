# BatchCheckItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | [**CheckRequestTupleKey**](CheckRequestTupleKey.md) |  | 
**ContextualTuples** | Pointer to [**ContextualTupleKeys**](ContextualTupleKeys.md) |  | [optional] 
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**CorrelationId** | **string** | correlation_id must be a string containing only letters, numbers, or hyphens, with length ≤ 36 characters. | 

## Methods

### NewBatchCheckItem

`func NewBatchCheckItem(tupleKey CheckRequestTupleKey, correlationId string, ) *BatchCheckItem`

NewBatchCheckItem instantiates a new BatchCheckItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCheckItemWithDefaults

`func NewBatchCheckItemWithDefaults() *BatchCheckItem`

NewBatchCheckItemWithDefaults instantiates a new BatchCheckItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *BatchCheckItem) GetTupleKey() CheckRequestTupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *BatchCheckItem) GetTupleKeyOk() (*CheckRequestTupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *BatchCheckItem) SetTupleKey(v CheckRequestTupleKey)`

SetTupleKey sets TupleKey field to given value.


### GetContextualTuples

`func (o *BatchCheckItem) GetContextualTuples() ContextualTupleKeys`

GetContextualTuples returns the ContextualTuples field if non-nil, zero value otherwise.

### GetContextualTuplesOk

`func (o *BatchCheckItem) GetContextualTuplesOk() (*ContextualTupleKeys, bool)`

GetContextualTuplesOk returns a tuple with the ContextualTuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualTuples

`func (o *BatchCheckItem) SetContextualTuples(v ContextualTupleKeys)`

SetContextualTuples sets ContextualTuples field to given value.

### HasContextualTuples

`func (o *BatchCheckItem) HasContextualTuples() bool`

HasContextualTuples returns a boolean if a field has been set.

### GetContext

`func (o *BatchCheckItem) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *BatchCheckItem) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *BatchCheckItem) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *BatchCheckItem) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCorrelationId

`func (o *BatchCheckItem) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *BatchCheckItem) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *BatchCheckItem) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


