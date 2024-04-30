# ObjectOrUserset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**FgaObject**](FgaObject.md) |  | [optional] 
**Userset** | Pointer to [**UsersetUser**](UsersetUser.md) |  | [optional] 

## Methods

### NewObjectOrUserset

`func NewObjectOrUserset() *ObjectOrUserset`

NewObjectOrUserset instantiates a new ObjectOrUserset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectOrUsersetWithDefaults

`func NewObjectOrUsersetWithDefaults() *ObjectOrUserset`

NewObjectOrUsersetWithDefaults instantiates a new ObjectOrUserset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ObjectOrUserset) GetObject() FgaObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ObjectOrUserset) GetObjectOk() (*FgaObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ObjectOrUserset) SetObject(v FgaObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *ObjectOrUserset) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUserset

`func (o *ObjectOrUserset) GetUserset() UsersetUser`

GetUserset returns the Userset field if non-nil, zero value otherwise.

### GetUsersetOk

`func (o *ObjectOrUserset) GetUsersetOk() (*UsersetUser, bool)`

GetUsersetOk returns a tuple with the Userset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserset

`func (o *ObjectOrUserset) SetUserset(v UsersetUser)`

SetUserset sets Userset field to given value.

### HasUserset

`func (o *ObjectOrUserset) HasUserset() bool`

HasUserset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


