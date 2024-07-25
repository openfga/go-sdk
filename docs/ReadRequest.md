# ReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | Pointer to [**ReadRequestTupleKey**](ReadRequestTupleKey.md) |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**ContinuationToken** | Pointer to **string** |  | [optional] 
**Consistency** | Pointer to [**ConsistencyPreference**](ConsistencyPreference.md) |  | [optional] [default to UNSPECIFIED]

## Methods

### NewReadRequest

`func NewReadRequest() *ReadRequest`

NewReadRequest instantiates a new ReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadRequestWithDefaults

`func NewReadRequestWithDefaults() *ReadRequest`

NewReadRequestWithDefaults instantiates a new ReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *ReadRequest) GetTupleKey() ReadRequestTupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *ReadRequest) GetTupleKeyOk() (*ReadRequestTupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *ReadRequest) SetTupleKey(v ReadRequestTupleKey)`

SetTupleKey sets TupleKey field to given value.

### HasTupleKey

`func (o *ReadRequest) HasTupleKey() bool`

HasTupleKey returns a boolean if a field has been set.

### GetPageSize

`func (o *ReadRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ReadRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ReadRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ReadRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetContinuationToken

`func (o *ReadRequest) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ReadRequest) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ReadRequest) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ReadRequest) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetConsistency

`func (o *ReadRequest) GetConsistency() ConsistencyPreference`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *ReadRequest) GetConsistencyOk() (*ConsistencyPreference, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *ReadRequest) SetConsistency(v ConsistencyPreference)`

SetConsistency sets Consistency field to given value.

### HasConsistency

`func (o *ReadRequest) HasConsistency() bool`

HasConsistency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


