# WriteRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Writes** | Pointer to [**TupleKeys**](TupleKeys.md) |  | [optional] 
**Deletes** | Pointer to [**TupleKeys**](TupleKeys.md) |  | [optional] 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 

## Methods

### NewWriteRequestParams

`func NewWriteRequestParams() *WriteRequestParams`

NewWriteRequestParams instantiates a new WriteRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteRequestParamsWithDefaults

`func NewWriteRequestParamsWithDefaults() *WriteRequestParams`

NewWriteRequestParamsWithDefaults instantiates a new WriteRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWrites

`func (o *WriteRequestParams) GetWrites() TupleKeys`

GetWrites returns the Writes field if non-nil, zero value otherwise.

### GetWritesOk

`func (o *WriteRequestParams) GetWritesOk() (*TupleKeys, bool)`

GetWritesOk returns a tuple with the Writes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrites

`func (o *WriteRequestParams) SetWrites(v TupleKeys)`

SetWrites sets Writes field to given value.

### HasWrites

`func (o *WriteRequestParams) HasWrites() bool`

HasWrites returns a boolean if a field has been set.

### GetDeletes

`func (o *WriteRequestParams) GetDeletes() TupleKeys`

GetDeletes returns the Deletes field if non-nil, zero value otherwise.

### GetDeletesOk

`func (o *WriteRequestParams) GetDeletesOk() (*TupleKeys, bool)`

GetDeletesOk returns a tuple with the Deletes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletes

`func (o *WriteRequestParams) SetDeletes(v TupleKeys)`

SetDeletes sets Deletes field to given value.

### HasDeletes

`func (o *WriteRequestParams) HasDeletes() bool`

HasDeletes returns a boolean if a field has been set.

### GetAuthorizationModelId

`func (o *WriteRequestParams) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *WriteRequestParams) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *WriteRequestParams) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *WriteRequestParams) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


