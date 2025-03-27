# BatchCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**map[string]BatchCheckSingleResult**](BatchCheckSingleResult.md) | map keys are the correlation_id values from the BatchCheckItems in the request | [optional] 

## Methods

### NewBatchCheckResponse

`func NewBatchCheckResponse() *BatchCheckResponse`

NewBatchCheckResponse instantiates a new BatchCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCheckResponseWithDefaults

`func NewBatchCheckResponseWithDefaults() *BatchCheckResponse`

NewBatchCheckResponseWithDefaults instantiates a new BatchCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *BatchCheckResponse) GetResult() map[string]BatchCheckSingleResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BatchCheckResponse) GetResultOk() (*map[string]BatchCheckSingleResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BatchCheckResponse) SetResult(v map[string]BatchCheckSingleResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *BatchCheckResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


