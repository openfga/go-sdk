# ReadAuthorizationModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationModels** | Pointer to [**[]AuthorizationModel**](AuthorizationModel.md) |  | [optional] 
**ContinuationToken** | Pointer to **string** |  | [optional] 

## Methods

### NewReadAuthorizationModelsResponse

`func NewReadAuthorizationModelsResponse() *ReadAuthorizationModelsResponse`

NewReadAuthorizationModelsResponse instantiates a new ReadAuthorizationModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadAuthorizationModelsResponseWithDefaults

`func NewReadAuthorizationModelsResponseWithDefaults() *ReadAuthorizationModelsResponse`

NewReadAuthorizationModelsResponseWithDefaults instantiates a new ReadAuthorizationModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationModels

`func (o *ReadAuthorizationModelsResponse) GetAuthorizationModels() []AuthorizationModel`

GetAuthorizationModels returns the AuthorizationModels field if non-nil, zero value otherwise.

### GetAuthorizationModelsOk

`func (o *ReadAuthorizationModelsResponse) GetAuthorizationModelsOk() (*[]AuthorizationModel, bool)`

GetAuthorizationModelsOk returns a tuple with the AuthorizationModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModels

`func (o *ReadAuthorizationModelsResponse) SetAuthorizationModels(v []AuthorizationModel)`

SetAuthorizationModels sets AuthorizationModels field to given value.

### HasAuthorizationModels

`func (o *ReadAuthorizationModelsResponse) HasAuthorizationModels() bool`

HasAuthorizationModels returns a boolean if a field has been set.

### GetContinuationToken

`func (o *ReadAuthorizationModelsResponse) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ReadAuthorizationModelsResponse) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ReadAuthorizationModelsResponse) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ReadAuthorizationModelsResponse) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


