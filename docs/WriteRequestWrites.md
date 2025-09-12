# WriteRequestWrites

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKeys** | [**[]TupleKey**](TupleKey.md) |  | 
**OnDuplicate** | Pointer to **string** | On &#39;error&#39; ( or unspecified ), the API returns an error if an identical tuple already exists. On &#39;ignore&#39;, identical writes are treated as no-ops (matching on user, relation, object, and RelationshipCondition). | [optional] [default to "error"]

## Methods

### NewWriteRequestWrites

`func NewWriteRequestWrites(tupleKeys []TupleKey, ) *WriteRequestWrites`

NewWriteRequestWrites instantiates a new WriteRequestWrites object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteRequestWritesWithDefaults

`func NewWriteRequestWritesWithDefaults() *WriteRequestWrites`

NewWriteRequestWritesWithDefaults instantiates a new WriteRequestWrites object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKeys

`func (o *WriteRequestWrites) GetTupleKeys() []TupleKey`

GetTupleKeys returns the TupleKeys field if non-nil, zero value otherwise.

### GetTupleKeysOk

`func (o *WriteRequestWrites) GetTupleKeysOk() (*[]TupleKey, bool)`

GetTupleKeysOk returns a tuple with the TupleKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKeys

`func (o *WriteRequestWrites) SetTupleKeys(v []TupleKey)`

SetTupleKeys sets TupleKeys field to given value.


### GetOnDuplicate

`func (o *WriteRequestWrites) GetOnDuplicate() string`

GetOnDuplicate returns the OnDuplicate field if non-nil, zero value otherwise.

### GetOnDuplicateOk

`func (o *WriteRequestWrites) GetOnDuplicateOk() (*string, bool)`

GetOnDuplicateOk returns a tuple with the OnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDuplicate

`func (o *WriteRequestWrites) SetOnDuplicate(v string)`

SetOnDuplicate sets OnDuplicate field to given value.

### HasOnDuplicate

`func (o *WriteRequestWrites) HasOnDuplicate() bool`

HasOnDuplicate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


