# ConditionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **string** |  | [optional] 
**SourceInfo** | Pointer to [**SourceInfo**](SourceInfo.md) |  | [optional] 

## Methods

### NewConditionMetadata

`func NewConditionMetadata() *ConditionMetadata`

NewConditionMetadata instantiates a new ConditionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionMetadataWithDefaults

`func NewConditionMetadataWithDefaults() *ConditionMetadata`

NewConditionMetadataWithDefaults instantiates a new ConditionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *ConditionMetadata) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ConditionMetadata) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ConditionMetadata) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *ConditionMetadata) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetSourceInfo

`func (o *ConditionMetadata) GetSourceInfo() SourceInfo`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *ConditionMetadata) GetSourceInfoOk() (*SourceInfo, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *ConditionMetadata) SetSourceInfo(v SourceInfo)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *ConditionMetadata) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


