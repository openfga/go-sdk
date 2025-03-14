# BatchCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | [**[]BatchCheckItem**](BatchCheckItem.md) |  | 
**AuthorizationModelId** | Pointer to **string** |  | [optional] 
**Consistency** | Pointer to [**ConsistencyPreference**](ConsistencyPreference.md) |  | [optional] [default to CONSISTENCYPREFERENCE_UNSPECIFIED]

## Methods

### NewBatchCheckRequest

`func NewBatchCheckRequest(checks []BatchCheckItem, ) *BatchCheckRequest`

NewBatchCheckRequest instantiates a new BatchCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCheckRequestWithDefaults

`func NewBatchCheckRequestWithDefaults() *BatchCheckRequest`

NewBatchCheckRequestWithDefaults instantiates a new BatchCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *BatchCheckRequest) GetChecks() []BatchCheckItem`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *BatchCheckRequest) GetChecksOk() (*[]BatchCheckItem, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *BatchCheckRequest) SetChecks(v []BatchCheckItem)`

SetChecks sets Checks field to given value.


### GetAuthorizationModelId

`func (o *BatchCheckRequest) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *BatchCheckRequest) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *BatchCheckRequest) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.

### HasAuthorizationModelId

`func (o *BatchCheckRequest) HasAuthorizationModelId() bool`

HasAuthorizationModelId returns a boolean if a field has been set.

### GetConsistency

`func (o *BatchCheckRequest) GetConsistency() ConsistencyPreference`

GetConsistency returns the Consistency field if non-nil, zero value otherwise.

### GetConsistencyOk

`func (o *BatchCheckRequest) GetConsistencyOk() (*ConsistencyPreference, bool)`

GetConsistencyOk returns a tuple with the Consistency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsistency

`func (o *BatchCheckRequest) SetConsistency(v ConsistencyPreference)`

SetConsistency sets Consistency field to given value.

### HasConsistency

`func (o *BatchCheckRequest) HasConsistency() bool`

HasConsistency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


