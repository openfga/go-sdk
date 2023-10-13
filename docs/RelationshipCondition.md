# RelationshipCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A reference (by name) of the relationship condition defined in the authorization model. | 
**Context** | **map[string]interface{}** | Additional context/data to persist along with the condition. The keys must match the parameters defined by the condition, and the value types must match the parameter type definitions. | 

## Methods

### NewRelationshipCondition

`func NewRelationshipCondition(name string, context map[string]interface{}, ) *RelationshipCondition`

NewRelationshipCondition instantiates a new RelationshipCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipConditionWithDefaults

`func NewRelationshipConditionWithDefaults() *RelationshipCondition`

NewRelationshipConditionWithDefaults instantiates a new RelationshipCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RelationshipCondition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelationshipCondition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelationshipCondition) SetName(v string)`

SetName sets Name field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


