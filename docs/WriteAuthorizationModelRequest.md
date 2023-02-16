# WriteAuthorizationModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeDefinitions** | [**[]TypeDefinition**](TypeDefinition.md) |  | 
**SchemaVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewWriteAuthorizationModelRequest

`func NewWriteAuthorizationModelRequest(typeDefinitions []TypeDefinition, ) *WriteAuthorizationModelRequest`

NewWriteAuthorizationModelRequest instantiates a new WriteAuthorizationModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteAuthorizationModelRequestWithDefaults

`func NewWriteAuthorizationModelRequestWithDefaults() *WriteAuthorizationModelRequest`

NewWriteAuthorizationModelRequestWithDefaults instantiates a new WriteAuthorizationModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeDefinitions

`func (o *WriteAuthorizationModelRequest) GetTypeDefinitions() []TypeDefinition`

GetTypeDefinitions returns the TypeDefinitions field if non-nil, zero value otherwise.

### GetTypeDefinitionsOk

`func (o *WriteAuthorizationModelRequest) GetTypeDefinitionsOk() (*[]TypeDefinition, bool)`

GetTypeDefinitionsOk returns a tuple with the TypeDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDefinitions

`func (o *WriteAuthorizationModelRequest) SetTypeDefinitions(v []TypeDefinition)`

SetTypeDefinitions sets TypeDefinitions field to given value.


### GetSchemaVersion

`func (o *WriteAuthorizationModelRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *WriteAuthorizationModelRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *WriteAuthorizationModelRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *WriteAuthorizationModelRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


