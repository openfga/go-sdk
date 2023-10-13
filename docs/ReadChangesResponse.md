# ReadChangesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | [**[]TupleChange**](TupleChange.md) |  | 
**ContinuationToken** | Pointer to **string** | The continuation token will be identical if there are no new changes. | [optional] 

## Methods

### NewReadChangesResponse

`func NewReadChangesResponse(changes []TupleChange, ) *ReadChangesResponse`

NewReadChangesResponse instantiates a new ReadChangesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadChangesResponseWithDefaults

`func NewReadChangesResponseWithDefaults() *ReadChangesResponse`

NewReadChangesResponseWithDefaults instantiates a new ReadChangesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *ReadChangesResponse) GetChanges() []TupleChange`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ReadChangesResponse) GetChangesOk() (*[]TupleChange, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ReadChangesResponse) SetChanges(v []TupleChange)`

SetChanges sets Changes field to given value.


### GetContinuationToken

`func (o *ReadChangesResponse) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ReadChangesResponse) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ReadChangesResponse) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ReadChangesResponse) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


