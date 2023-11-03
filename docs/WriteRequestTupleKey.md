# WriteRequestTupleKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | **string** |  | 
**Relation** | **string** |  | 
**Object** | **string** |  | 
**Condition** | Pointer to [**RelationshipCondition**](RelationshipCondition.md) |  | [optional] 

## Methods

### NewWriteRequestTupleKey

`func NewWriteRequestTupleKey(user string, relation string, object string, ) *WriteRequestTupleKey`

NewWriteRequestTupleKey instantiates a new WriteRequestTupleKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteRequestTupleKeyWithDefaults

`func NewWriteRequestTupleKeyWithDefaults() *WriteRequestTupleKey`

NewWriteRequestTupleKeyWithDefaults instantiates a new WriteRequestTupleKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *WriteRequestTupleKey) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WriteRequestTupleKey) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WriteRequestTupleKey) SetUser(v string)`

SetUser sets User field to given value.


### GetRelation

`func (o *WriteRequestTupleKey) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *WriteRequestTupleKey) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *WriteRequestTupleKey) SetRelation(v string)`

SetRelation sets Relation field to given value.


### GetObject

`func (o *WriteRequestTupleKey) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WriteRequestTupleKey) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WriteRequestTupleKey) SetObject(v string)`

SetObject sets Object field to given value.


### GetCondition

`func (o *WriteRequestTupleKey) GetCondition() RelationshipCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WriteRequestTupleKey) GetConditionOk() (*RelationshipCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WriteRequestTupleKey) SetCondition(v RelationshipCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WriteRequestTupleKey) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


