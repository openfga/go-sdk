# TupleChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | Pointer to [**TupleKey**](TupleKey.md) |  | [optional] 
**Operation** | Pointer to [**TupleOperation**](TupleOperation.md) |  | [optional] [default to WRITE]
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTupleChange

`func NewTupleChange() *TupleChange`

NewTupleChange instantiates a new TupleChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTupleChangeWithDefaults

`func NewTupleChangeWithDefaults() *TupleChange`

NewTupleChangeWithDefaults instantiates a new TupleChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *TupleChange) GetTupleKey() TupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *TupleChange) GetTupleKeyOk() (*TupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *TupleChange) SetTupleKey(v TupleKey)`

SetTupleKey sets TupleKey field to given value.

### HasTupleKey

`func (o *TupleChange) HasTupleKey() bool`

HasTupleKey returns a boolean if a field has been set.

### GetOperation

`func (o *TupleChange) GetOperation() TupleOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TupleChange) GetOperationOk() (*TupleOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TupleChange) SetOperation(v TupleOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *TupleChange) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetTimestamp

`func (o *TupleChange) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TupleChange) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TupleChange) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TupleChange) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


