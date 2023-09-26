# RelationshipCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionName** | **string** | A reference (by name) of the relationship condition defined in the authorization model. | 
**Context** | Pointer to **map[string]interface{}** | Additional context/data to persist along with the condition. The keys must match the parameters defined by the condition, and the value types must match the parameter type definitions. | [optional] 

## Methods

### NewRelationshipCondition

`func NewRelationshipCondition(conditionName string, ) *RelationshipCondition`

NewRelationshipCondition instantiates a new RelationshipCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipConditionWithDefaults

`func NewRelationshipConditionWithDefaults() *RelationshipCondition`

NewRelationshipConditionWithDefaults instantiates a new RelationshipCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionName

`func (o *RelationshipCondition) GetConditionName() string`

GetConditionName returns the ConditionName field if non-nil, zero value otherwise.

### GetConditionNameOk

`func (o *RelationshipCondition) GetConditionNameOk() (*string, bool)`

GetConditionNameOk returns a tuple with the ConditionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionName

`func (o *RelationshipCondition) SetConditionName(v string)`

SetConditionName sets ConditionName field to given value.


### GetContext

`func (o *RelationshipCondition) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RelationshipCondition) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RelationshipCondition) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *RelationshipCondition) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


