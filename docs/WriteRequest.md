# WriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Writes** | Pointer to [**WriteRequestWrites**](WriteRequestWrites.md) |  | [optional] 
**Deletes** | Pointer to [**WriteRequestDeletes**](WriteRequestDeletes.md) |  | [optional] 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 

## Methods

### NewWriteRequest

`func NewWriteRequest() *WriteRequest`

NewWriteRequest instantiates a new WriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteRequestWithDefaults

`func NewWriteRequestWithDefaults() *WriteRequest`

NewWriteRequestWithDefaults instantiates a new WriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWrites

`func (o *WriteRequest) GetWrites() WriteRequestWrites`

GetWrites returns the Writes field if non-nil, zero value otherwise.

### GetWritesOk

`func (o *WriteRequest) GetWritesOk() (*WriteRequestWrites, bool)`

GetWritesOk returns a tuple with the Writes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrites

`func (o *WriteRequest) SetWrites(v WriteRequestWrites)`

SetWrites sets Writes field to given value.

### HasWrites

`func (o *WriteRequest) HasWrites() bool`

HasWrites returns a boolean if a field has been set.

### GetDeletes

`func (o *WriteRequest) GetDeletes() WriteRequestDeletes`

GetDeletes returns the Deletes field if non-nil, zero value otherwise.

### GetDeletesOk

`func (o *WriteRequest) GetDeletesOk() (*WriteRequestDeletes, bool)`

GetDeletesOk returns a tuple with the Deletes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletes

`func (o *WriteRequest) SetDeletes(v WriteRequestDeletes)`

SetDeletes sets Deletes field to given value.

### HasDeletes

`func (o *WriteRequest) HasDeletes() bool`

HasDeletes returns a boolean if a field has been set.

### GetAuthorizationModelId

`func (o *WriteRequest) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *WriteRequest) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *WriteRequest) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *WriteRequest) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


