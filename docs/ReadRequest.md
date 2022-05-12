# ReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | Pointer to [**TupleKey**](TupleKey.md) |  | [optional] 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**ContinuationToken** | Pointer to **string** |  | [optional] 

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

`func (o *ReadRequest) GetTupleKey() TupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *ReadRequest) GetTupleKeyOk() (*TupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *ReadRequest) SetTupleKey(v TupleKey)`

SetTupleKey sets TupleKey field to given value.

### HasTupleKey

`func (o *ReadRequest) HasTupleKey() bool`

HasTupleKey returns a boolean if a field has been set.

### GetAuthorizationModelId

`func (o *ReadRequest) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *ReadRequest) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *ReadRequest) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *ReadRequest) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


