# ExpandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | [**TupleKey**](TupleKey.md) |  | 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 

## Methods

### NewExpandRequest

`func NewExpandRequest(tupleKey TupleKey, ) *ExpandRequest`

NewExpandRequest instantiates a new ExpandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandRequestWithDefaults

`func NewExpandRequestWithDefaults() *ExpandRequest`

NewExpandRequestWithDefaults instantiates a new ExpandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *ExpandRequest) GetTupleKey() TupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *ExpandRequest) GetTupleKeyOk() (*TupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *ExpandRequest) SetTupleKey(v TupleKey)`

SetTupleKey sets TupleKey field to given value.


### GetAuthorizationModelId

`func (o *ExpandRequest) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *ExpandRequest) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *ExpandRequest) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *ExpandRequest) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


