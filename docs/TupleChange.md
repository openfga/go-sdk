# TupleChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | [**TupleKey**](TupleKey.md) |  | 
**Operation** | [**TupleOperation**](TupleOperation.md) |  | [default to TUPLEOPERATION_WRITE]
**Timestamp** | **time.Time** |  | 

## Methods

### NewTupleChange

`func NewTupleChange(tupleKey TupleKey, operation TupleOperation, timestamp time.Time, ) *TupleChange`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


