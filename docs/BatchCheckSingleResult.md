# BatchCheckSingleResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to [**CheckError**](CheckError.md) |  | [optional] 

## Methods

### NewBatchCheckSingleResult

`func NewBatchCheckSingleResult() *BatchCheckSingleResult`

NewBatchCheckSingleResult instantiates a new BatchCheckSingleResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCheckSingleResultWithDefaults

`func NewBatchCheckSingleResultWithDefaults() *BatchCheckSingleResult`

NewBatchCheckSingleResultWithDefaults instantiates a new BatchCheckSingleResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *BatchCheckSingleResult) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *BatchCheckSingleResult) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *BatchCheckSingleResult) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *BatchCheckSingleResult) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetError

`func (o *BatchCheckSingleResult) GetError() CheckError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchCheckSingleResult) GetErrorOk() (*CheckError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchCheckSingleResult) SetError(v CheckError)`

SetError sets Error field to given value.

### HasError

`func (o *BatchCheckSingleResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


