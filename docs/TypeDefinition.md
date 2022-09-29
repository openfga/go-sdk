# TypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Relations** | Pointer to [**map[string]Userset**](Userset.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewTypeDefinition

`func NewTypeDefinition(type_ string, ) *TypeDefinition`

NewTypeDefinition instantiates a new TypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeDefinitionWithDefaults

`func NewTypeDefinitionWithDefaults() *TypeDefinition`

NewTypeDefinitionWithDefaults instantiates a new TypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TypeDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TypeDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TypeDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetRelations

`func (o *TypeDefinition) GetRelations() map[string]Userset`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *TypeDefinition) GetRelationsOk() (*map[string]Userset, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *TypeDefinition) SetRelations(v map[string]Userset)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *TypeDefinition) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetMetadata

`func (o *TypeDefinition) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TypeDefinition) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TypeDefinition) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TypeDefinition) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


