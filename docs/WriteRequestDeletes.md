# WriteRequestDeletes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKeys** | [**[]TupleKeyWithoutCondition**](TupleKeyWithoutCondition.md) |  | 
**OnMissing** | Pointer to **string** | On &#39;error&#39;, the API returns an error when deleting a tuple that does not exist. On &#39;ignore&#39;, deletes of non-existent tuples are treated as no-ops. | [optional] [default to "error"]

## Methods

### NewWriteRequestDeletes

`func NewWriteRequestDeletes(tupleKeys []TupleKeyWithoutCondition, ) *WriteRequestDeletes`

NewWriteRequestDeletes instantiates a new WriteRequestDeletes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteRequestDeletesWithDefaults

`func NewWriteRequestDeletesWithDefaults() *WriteRequestDeletes`

NewWriteRequestDeletesWithDefaults instantiates a new WriteRequestDeletes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKeys

`func (o *WriteRequestDeletes) GetTupleKeys() []TupleKeyWithoutCondition`

GetTupleKeys returns the TupleKeys field if non-nil, zero value otherwise.

### GetTupleKeysOk

`func (o *WriteRequestDeletes) GetTupleKeysOk() (*[]TupleKeyWithoutCondition, bool)`

GetTupleKeysOk returns a tuple with the TupleKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKeys

`func (o *WriteRequestDeletes) SetTupleKeys(v []TupleKeyWithoutCondition)`

SetTupleKeys sets TupleKeys field to given value.


### GetOnMissing

`func (o *WriteRequestDeletes) GetOnMissing() string`

GetOnMissing returns the OnMissing field if non-nil, zero value otherwise.

### GetOnMissingOk

`func (o *WriteRequestDeletes) GetOnMissingOk() (*string, bool)`

GetOnMissingOk returns a tuple with the OnMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnMissing

`func (o *WriteRequestDeletes) SetOnMissing(v string)`

SetOnMissing sets OnMissing field to given value.

### HasOnMissing

`func (o *WriteRequestDeletes) HasOnMissing() bool`

HasOnMissing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


