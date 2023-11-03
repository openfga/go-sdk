# RelationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Relation** | Pointer to **string** |  | [optional] 
**Wildcard** | Pointer to **map[string]interface{}** |  | [optional] 
**Condition** | Pointer to **string** | The name of a condition that is enforced over the allowed relation. | [optional] 

## Methods

### NewRelationReference

`func NewRelationReference(type_ string, ) *RelationReference`

NewRelationReference instantiates a new RelationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationReferenceWithDefaults

`func NewRelationReferenceWithDefaults() *RelationReference`

NewRelationReferenceWithDefaults instantiates a new RelationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RelationReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationReference) SetType(v string)`

SetType sets Type field to given value.


### GetRelation

`func (o *RelationReference) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *RelationReference) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *RelationReference) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *RelationReference) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetWildcard

`func (o *RelationReference) GetWildcard() map[string]interface{}`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *RelationReference) GetWildcardOk() (*map[string]interface{}, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *RelationReference) SetWildcard(v map[string]interface{})`

SetWildcard sets Wildcard field to given value.

### HasWildcard

`func (o *RelationReference) HasWildcard() bool`

HasWildcard returns a boolean if a field has been set.

### GetCondition

`func (o *RelationReference) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RelationReference) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RelationReference) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RelationReference) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


