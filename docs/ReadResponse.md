# ReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tuples** | Pointer to [**[]Tuple**](Tuple.md) |  | [optional] 
**ContinuationToken** | Pointer to **string** |  | [optional] 

## Methods

### NewReadResponse

`func NewReadResponse() *ReadResponse`

NewReadResponse instantiates a new ReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadResponseWithDefaults

`func NewReadResponseWithDefaults() *ReadResponse`

NewReadResponseWithDefaults instantiates a new ReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTuples

`func (o *ReadResponse) GetTuples() []Tuple`

GetTuples returns the Tuples field if non-nil, zero value otherwise.

### GetTuplesOk

`func (o *ReadResponse) GetTuplesOk() (*[]Tuple, bool)`

GetTuplesOk returns a tuple with the Tuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuples

`func (o *ReadResponse) SetTuples(v []Tuple)`

SetTuples sets Tuples field to given value.

### HasTuples

`func (o *ReadResponse) HasTuples() bool`

HasTuples returns a boolean if a field has been set.

### GetContinuationToken

`func (o *ReadResponse) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ReadResponse) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ReadResponse) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ReadResponse) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


