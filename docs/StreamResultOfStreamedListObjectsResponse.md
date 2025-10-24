# StreamResultOfStreamedListObjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**StreamedListObjectsResponse**](StreamedListObjectsResponse.md) |  | [optional] 
**Error** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewStreamResultOfStreamedListObjectsResponse

`func NewStreamResultOfStreamedListObjectsResponse() *StreamResultOfStreamedListObjectsResponse`

NewStreamResultOfStreamedListObjectsResponse instantiates a new StreamResultOfStreamedListObjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfStreamedListObjectsResponseWithDefaults

`func NewStreamResultOfStreamedListObjectsResponseWithDefaults() *StreamResultOfStreamedListObjectsResponse`

NewStreamResultOfStreamedListObjectsResponseWithDefaults instantiates a new StreamResultOfStreamedListObjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *StreamResultOfStreamedListObjectsResponse) GetResult() StreamedListObjectsResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfStreamedListObjectsResponse) GetResultOk() (*StreamedListObjectsResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfStreamedListObjectsResponse) SetResult(v StreamedListObjectsResponse)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfStreamedListObjectsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *StreamResultOfStreamedListObjectsResponse) GetError() Status`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfStreamedListObjectsResponse) GetErrorOk() (*Status, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfStreamedListObjectsResponse) SetError(v Status)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfStreamedListObjectsResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


