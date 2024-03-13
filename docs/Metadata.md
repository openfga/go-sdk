# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relations** | Pointer to [**map[string]RelationMetadata**](RelationMetadata.md) |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 
**SourceInfo** | Pointer to [**SourceInfo**](SourceInfo.md) |  | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelations

`func (o *Metadata) GetRelations() map[string]RelationMetadata`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *Metadata) GetRelationsOk() (*map[string]RelationMetadata, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *Metadata) SetRelations(v map[string]RelationMetadata)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *Metadata) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetModule

`func (o *Metadata) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *Metadata) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *Metadata) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *Metadata) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetSourceInfo

`func (o *Metadata) GetSourceInfo() SourceInfo`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *Metadata) GetSourceInfoOk() (*SourceInfo, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *Metadata) SetSourceInfo(v SourceInfo)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *Metadata) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


