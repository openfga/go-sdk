# ConditionParamTypeRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeName** | [**TypeName**](TypeName.md) |  | [default to TYPENAME_UNSPECIFIED]
**GenericTypes** | Pointer to [**[]ConditionParamTypeRef**](ConditionParamTypeRef.md) |  | [optional] 

## Methods

### NewConditionParamTypeRef

`func NewConditionParamTypeRef(typeName TypeName, ) *ConditionParamTypeRef`

NewConditionParamTypeRef instantiates a new ConditionParamTypeRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionParamTypeRefWithDefaults

`func NewConditionParamTypeRefWithDefaults() *ConditionParamTypeRef`

NewConditionParamTypeRefWithDefaults instantiates a new ConditionParamTypeRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeName

`func (o *ConditionParamTypeRef) GetTypeName() TypeName`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *ConditionParamTypeRef) GetTypeNameOk() (*TypeName, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *ConditionParamTypeRef) SetTypeName(v TypeName)`

SetTypeName sets TypeName field to given value.


### GetGenericTypes

`func (o *ConditionParamTypeRef) GetGenericTypes() []ConditionParamTypeRef`

GetGenericTypes returns the GenericTypes field if non-nil, zero value otherwise.

### GetGenericTypesOk

`func (o *ConditionParamTypeRef) GetGenericTypesOk() (*[]ConditionParamTypeRef, bool)`

GetGenericTypesOk returns a tuple with the GenericTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericTypes

`func (o *ConditionParamTypeRef) SetGenericTypes(v []ConditionParamTypeRef)`

SetGenericTypes sets GenericTypes field to given value.

### HasGenericTypes

`func (o *ConditionParamTypeRef) HasGenericTypes() bool`

HasGenericTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


