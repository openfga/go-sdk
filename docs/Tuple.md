# Tuple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**TupleKey**](TupleKey.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTuple

`func NewTuple() *Tuple`

NewTuple instantiates a new Tuple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTupleWithDefaults

`func NewTupleWithDefaults() *Tuple`

NewTupleWithDefaults instantiates a new Tuple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Tuple) GetKey() TupleKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Tuple) GetKeyOk() (*TupleKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Tuple) SetKey(v TupleKey)`

SetKey sets Key field to given value.

### HasKey

`func (o *Tuple) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTimestamp

`func (o *Tuple) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Tuple) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Tuple) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Tuple) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


