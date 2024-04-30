# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**FgaObject**](FgaObject.md) |  | [optional] 
**Userset** | Pointer to [**UsersetUser**](UsersetUser.md) |  | [optional] 
**Wildcard** | Pointer to [**TypedWildcard**](TypedWildcard.md) |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *User) GetObject() FgaObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *User) GetObjectOk() (*FgaObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *User) SetObject(v FgaObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *User) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUserset

`func (o *User) GetUserset() UsersetUser`

GetUserset returns the Userset field if non-nil, zero value otherwise.

### GetUsersetOk

`func (o *User) GetUsersetOk() (*UsersetUser, bool)`

GetUsersetOk returns a tuple with the Userset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserset

`func (o *User) SetUserset(v UsersetUser)`

SetUserset sets Userset field to given value.

### HasUserset

`func (o *User) HasUserset() bool`

HasUserset returns a boolean if a field has been set.

### GetWildcard

`func (o *User) GetWildcard() TypedWildcard`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *User) GetWildcardOk() (*TypedWildcard, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *User) SetWildcard(v TypedWildcard)`

SetWildcard sets Wildcard field to given value.

### HasWildcard

`func (o *User) HasWildcard() bool`

HasWildcard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


