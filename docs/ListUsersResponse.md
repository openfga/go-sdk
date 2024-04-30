# ListUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]User**](User.md) |  | 
**ExcludedUsers** | [**[]ObjectOrUserset**](ObjectOrUserset.md) |  | 

## Methods

### NewListUsersResponse

`func NewListUsersResponse(users []User, excludedUsers []ObjectOrUserset, ) *ListUsersResponse`

NewListUsersResponse instantiates a new ListUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersResponseWithDefaults

`func NewListUsersResponseWithDefaults() *ListUsersResponse`

NewListUsersResponseWithDefaults instantiates a new ListUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ListUsersResponse) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListUsersResponse) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListUsersResponse) SetUsers(v []User)`

SetUsers sets Users field to given value.


### GetExcludedUsers

`func (o *ListUsersResponse) GetExcludedUsers() []ObjectOrUserset`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *ListUsersResponse) GetExcludedUsersOk() (*[]ObjectOrUserset, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *ListUsersResponse) SetExcludedUsers(v []ObjectOrUserset)`

SetExcludedUsers sets ExcludedUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


