# ReadRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | Pointer to [**TupleKey**](TupleKey.md) |  | [optional] 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**ContinuationToken** | Pointer to **string** |  | [optional] 

## Methods

### NewReadRequestParams

`func NewReadRequestParams() *ReadRequestParams`

NewReadRequestParams instantiates a new ReadRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadRequestParamsWithDefaults

`func NewReadRequestParamsWithDefaults() *ReadRequestParams`

NewReadRequestParamsWithDefaults instantiates a new ReadRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *ReadRequestParams) GetTupleKey() TupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *ReadRequestParams) GetTupleKeyOk() (*TupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *ReadRequestParams) SetTupleKey(v TupleKey)`

SetTupleKey sets TupleKey field to given value.

### HasTupleKey

`func (o *ReadRequestParams) HasTupleKey() bool`

HasTupleKey returns a boolean if a field has been set.

### GetAuthorizationModelId

`func (o *ReadRequestParams) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *ReadRequestParams) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *ReadRequestParams) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *ReadRequestParams) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetPageSize

`func (o *ReadRequestParams) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ReadRequestParams) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ReadRequestParams) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ReadRequestParams) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetContinuationToken

`func (o *ReadRequestParams) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ReadRequestParams) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ReadRequestParams) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ReadRequestParams) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


