# ListStoresResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stores** | [**[]Store**](Store.md) |  | 
**ContinuationToken** | **string** | The continuation token will be empty if there are no more stores. | 

## Methods

### NewListStoresResponse

`func NewListStoresResponse(stores []Store, continuationToken string, ) *ListStoresResponse`

NewListStoresResponse instantiates a new ListStoresResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStoresResponseWithDefaults

`func NewListStoresResponseWithDefaults() *ListStoresResponse`

NewListStoresResponseWithDefaults instantiates a new ListStoresResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStores

`func (o *ListStoresResponse) GetStores() []Store`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *ListStoresResponse) GetStoresOk() (*[]Store, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *ListStoresResponse) SetStores(v []Store)`

SetStores sets Stores field to given value.


### GetContinuationToken

`func (o *ListStoresResponse) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ListStoresResponse) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ListStoresResponse) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


