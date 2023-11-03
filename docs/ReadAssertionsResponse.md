# ReadAssertionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationModelId** | **string** |  | 
**Assertions** | Pointer to [**[]Assertion**](Assertion.md) |  | [optional] 

## Methods

### NewReadAssertionsResponse

`func NewReadAssertionsResponse(authorizationModelId string, ) *ReadAssertionsResponse`

NewReadAssertionsResponse instantiates a new ReadAssertionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadAssertionsResponseWithDefaults

`func NewReadAssertionsResponseWithDefaults() *ReadAssertionsResponse`

NewReadAssertionsResponseWithDefaults instantiates a new ReadAssertionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationModelId

`func (o *ReadAssertionsResponse) GetAuthorizationModelId() string`

GetAuthorizationModelId returns the AuthorizationModelId field if non-nil, zero value otherwise.

### GetAuthorizationModelIdOk

`func (o *ReadAssertionsResponse) GetAuthorizationModelIdOk() (*string, bool)`

GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationModelId

`func (o *ReadAssertionsResponse) SetAuthorizationModelId(v string)`

SetAuthorizationModelId sets AuthorizationModelId field to given value.


### GetAssertions

`func (o *ReadAssertionsResponse) GetAssertions() []Assertion`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *ReadAssertionsResponse) GetAssertionsOk() (*[]Assertion, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *ReadAssertionsResponse) SetAssertions(v []Assertion)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *ReadAssertionsResponse) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


