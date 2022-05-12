# Userset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**This** | Pointer to **map[string]interface{}** | A DirectUserset is a sentinel message for referencing the direct members specified by an object/relation mapping. | [optional] 
**ComputedUserset** | Pointer to [**ObjectRelation**](ObjectRelation.md) |  | [optional] 
**TupleToUserset** | Pointer to [**AuthorizationmodelTupleToUserset**](AuthorizationmodelTupleToUserset.md) |  | [optional] 
**Union** | Pointer to [**Usersets**](Usersets.md) |  | [optional] 
**Intersection** | Pointer to [**Usersets**](Usersets.md) |  | [optional] 
**Difference** | Pointer to [**AuthorizationmodelDifference**](AuthorizationmodelDifference.md) |  | [optional] 

## Methods

### NewUserset

`func NewUserset() *Userset`

NewUserset instantiates a new Userset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersetWithDefaults

`func NewUsersetWithDefaults() *Userset`

NewUsersetWithDefaults instantiates a new Userset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThis

`func (o *Userset) GetThis() map[string]interface{}`

GetThis returns the This field if non-nil, zero value otherwise.

### GetThisOk

`func (o *Userset) GetThisOk() (*map[string]interface{}, bool)`

GetThisOk returns a tuple with the This field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThis

`func (o *Userset) SetThis(v map[string]interface{})`

SetThis sets This field to given value.

### HasThis

`func (o *Userset) HasThis() bool`

HasThis returns a boolean if a field has been set.

### GetComputedUserset

`func (o *Userset) GetComputedUserset() ObjectRelation`

GetComputedUserset returns the ComputedUserset field if non-nil, zero value otherwise.

### GetComputedUsersetOk

`func (o *Userset) GetComputedUsersetOk() (*ObjectRelation, bool)`

GetComputedUsersetOk returns a tuple with the ComputedUserset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedUserset

`func (o *Userset) SetComputedUserset(v ObjectRelation)`

SetComputedUserset sets ComputedUserset field to given value.

### HasComputedUserset

`func (o *Userset) HasComputedUserset() bool`

HasComputedUserset returns a boolean if a field has been set.

### GetTupleToUserset

`func (o *Userset) GetTupleToUserset() AuthorizationmodelTupleToUserset`

GetTupleToUserset returns the TupleToUserset field if non-nil, zero value otherwise.

### GetTupleToUsersetOk

`func (o *Userset) GetTupleToUsersetOk() (*AuthorizationmodelTupleToUserset, bool)`

GetTupleToUsersetOk returns a tuple with the TupleToUserset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleToUserset

`func (o *Userset) SetTupleToUserset(v AuthorizationmodelTupleToUserset)`

SetTupleToUserset sets TupleToUserset field to given value.

### HasTupleToUserset

`func (o *Userset) HasTupleToUserset() bool`

HasTupleToUserset returns a boolean if a field has been set.

### GetUnion

`func (o *Userset) GetUnion() Usersets`

GetUnion returns the Union field if non-nil, zero value otherwise.

### GetUnionOk

`func (o *Userset) GetUnionOk() (*Usersets, bool)`

GetUnionOk returns a tuple with the Union field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnion

`func (o *Userset) SetUnion(v Usersets)`

SetUnion sets Union field to given value.

### HasUnion

`func (o *Userset) HasUnion() bool`

HasUnion returns a boolean if a field has been set.

### GetIntersection

`func (o *Userset) GetIntersection() Usersets`

GetIntersection returns the Intersection field if non-nil, zero value otherwise.

### GetIntersectionOk

`func (o *Userset) GetIntersectionOk() (*Usersets, bool)`

GetIntersectionOk returns a tuple with the Intersection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntersection

`func (o *Userset) SetIntersection(v Usersets)`

SetIntersection sets Intersection field to given value.

### HasIntersection

`func (o *Userset) HasIntersection() bool`

HasIntersection returns a boolean if a field has been set.

### GetDifference

`func (o *Userset) GetDifference() AuthorizationmodelDifference`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *Userset) GetDifferenceOk() (*AuthorizationmodelDifference, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *Userset) SetDifference(v AuthorizationmodelDifference)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *Userset) HasDifference() bool`

HasDifference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


