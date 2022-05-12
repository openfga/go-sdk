# ObjectRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Relation** | Pointer to **string** |  | [optional] 

## Methods

### NewObjectRelation

`func NewObjectRelation() *ObjectRelation`

NewObjectRelation instantiates a new ObjectRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRelationWithDefaults

`func NewObjectRelationWithDefaults() *ObjectRelation`

NewObjectRelationWithDefaults instantiates a new ObjectRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ObjectRelation) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ObjectRelation) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ObjectRelation) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ObjectRelation) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelation

`func (o *ObjectRelation) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *ObjectRelation) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *ObjectRelation) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *ObjectRelation) HasRelation() bool`

HasRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


