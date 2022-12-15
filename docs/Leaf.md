# Leaf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**Users**](Users.md) |  | [optional] 
**Computed** | Pointer to [**Computed**](Computed.md) |  | [optional] 
**TupleToUserset** | Pointer to [**UsersetTreeTupleToUserset**](UsersetTreeTupleToUserset.md) |  | [optional] 

## Methods

### NewLeaf

`func NewLeaf() *Leaf`

NewLeaf instantiates a new Leaf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeafWithDefaults

`func NewLeafWithDefaults() *Leaf`

NewLeafWithDefaults instantiates a new Leaf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *Leaf) GetUsers() Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Leaf) GetUsersOk() (*Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Leaf) SetUsers(v Users)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Leaf) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetComputed

`func (o *Leaf) GetComputed() Computed`

GetComputed returns the Computed field if non-nil, zero value otherwise.

### GetComputedOk

`func (o *Leaf) GetComputedOk() (*Computed, bool)`

GetComputedOk returns a tuple with the Computed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputed

`func (o *Leaf) SetComputed(v Computed)`

SetComputed sets Computed field to given value.

### HasComputed

`func (o *Leaf) HasComputed() bool`

HasComputed returns a boolean if a field has been set.

### GetTupleToUserset

`func (o *Leaf) GetTupleToUserset() UsersetTreeTupleToUserset`

GetTupleToUserset returns the TupleToUserset field if non-nil, zero value otherwise.

### GetTupleToUsersetOk

`func (o *Leaf) GetTupleToUsersetOk() (*UsersetTreeTupleToUserset, bool)`

GetTupleToUsersetOk returns a tuple with the TupleToUserset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleToUserset

`func (o *Leaf) SetTupleToUserset(v UsersetTreeTupleToUserset)`

SetTupleToUserset sets TupleToUserset field to given value.

### HasTupleToUserset

`func (o *Leaf) HasTupleToUserset() bool`

HasTupleToUserset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


