# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Expression** | **string** | A Google CEL expression, expressed as a string. | 
**Parameters** | Pointer to [**map[string]ConditionParamTypeRef**](ConditionParamTypeRef.md) | A map of parameter names to the parameter&#39;s defined type reference. | [optional] 

## Methods

### NewCondition

`func NewCondition(name string, expression string, ) *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Condition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Condition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Condition) SetName(v string)`

SetName sets Name field to given value.


### GetExpression

`func (o *Condition) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Condition) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Condition) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetParameters

`func (o *Condition) GetParameters() map[string]ConditionParamTypeRef`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Condition) GetParametersOk() (*map[string]ConditionParamTypeRef, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Condition) SetParameters(v map[string]ConditionParamTypeRef)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Condition) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


