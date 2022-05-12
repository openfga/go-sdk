# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Leaf** | Pointer to [**Leaf**](Leaf.md) |  | [optional] 
**Difference** | Pointer to [**UsersetTreeDifference**](UsersetTreeDifference.md) |  | [optional] 
**Union** | Pointer to [**Nodes**](Nodes.md) |  | [optional] 
**Intersection** | Pointer to [**Nodes**](Nodes.md) |  | [optional] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLeaf

`func (o *Node) GetLeaf() Leaf`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *Node) GetLeafOk() (*Leaf, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *Node) SetLeaf(v Leaf)`

SetLeaf sets Leaf field to given value.

### HasLeaf

`func (o *Node) HasLeaf() bool`

HasLeaf returns a boolean if a field has been set.

### GetDifference

`func (o *Node) GetDifference() UsersetTreeDifference`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *Node) GetDifferenceOk() (*UsersetTreeDifference, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *Node) SetDifference(v UsersetTreeDifference)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *Node) HasDifference() bool`

HasDifference returns a boolean if a field has been set.

### GetUnion

`func (o *Node) GetUnion() Nodes`

GetUnion returns the Union field if non-nil, zero value otherwise.

### GetUnionOk

`func (o *Node) GetUnionOk() (*Nodes, bool)`

GetUnionOk returns a tuple with the Union field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnion

`func (o *Node) SetUnion(v Nodes)`

SetUnion sets Union field to given value.

### HasUnion

`func (o *Node) HasUnion() bool`

HasUnion returns a boolean if a field has been set.

### GetIntersection

`func (o *Node) GetIntersection() Nodes`

GetIntersection returns the Intersection field if non-nil, zero value otherwise.

### GetIntersectionOk

`func (o *Node) GetIntersectionOk() (*Nodes, bool)`

GetIntersectionOk returns a tuple with the Intersection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntersection

`func (o *Node) SetIntersection(v Nodes)`

SetIntersection sets Intersection field to given value.

### HasIntersection

`func (o *Node) HasIntersection() bool`

HasIntersection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


