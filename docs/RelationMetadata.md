# RelationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectlyRelatedUserTypes** | Pointer to [**[]RelationReference**](RelationReference.md) |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 
**SourceInfo** | Pointer to [**SourceInfo**](SourceInfo.md) |  | [optional] 

## Methods

### NewRelationMetadata

`func NewRelationMetadata() *RelationMetadata`

NewRelationMetadata instantiates a new RelationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationMetadataWithDefaults

`func NewRelationMetadataWithDefaults() *RelationMetadata`

NewRelationMetadataWithDefaults instantiates a new RelationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectlyRelatedUserTypes

`func (o *RelationMetadata) GetDirectlyRelatedUserTypes() []RelationReference`

GetDirectlyRelatedUserTypes returns the DirectlyRelatedUserTypes field if non-nil, zero value otherwise.

### GetDirectlyRelatedUserTypesOk

`func (o *RelationMetadata) GetDirectlyRelatedUserTypesOk() (*[]RelationReference, bool)`

GetDirectlyRelatedUserTypesOk returns a tuple with the DirectlyRelatedUserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectlyRelatedUserTypes

`func (o *RelationMetadata) SetDirectlyRelatedUserTypes(v []RelationReference)`

SetDirectlyRelatedUserTypes sets DirectlyRelatedUserTypes field to given value.

### HasDirectlyRelatedUserTypes

`func (o *RelationMetadata) HasDirectlyRelatedUserTypes() bool`

HasDirectlyRelatedUserTypes returns a boolean if a field has been set.

### GetModule

`func (o *RelationMetadata) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *RelationMetadata) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *RelationMetadata) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *RelationMetadata) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetSourceInfo

`func (o *RelationMetadata) GetSourceInfo() SourceInfo`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *RelationMetadata) GetSourceInfoOk() (*SourceInfo, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *RelationMetadata) SetSourceInfo(v SourceInfo)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *RelationMetadata) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


