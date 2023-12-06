# AuthorizationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SchemaVersion** | **string** |  | 
**TypeDefinitions** | [**[]TypeDefinition**](TypeDefinition.md) |  | 
**Conditions** | Pointer to [**map[string]Condition**](Condition.md) |  | [optional] 

## Methods

### NewAuthorizationModel

`func NewAuthorizationModel(id string, schemaVersion string, typeDefinitions []TypeDefinition, ) *AuthorizationModel`

NewAuthorizationModel instantiates a new AuthorizationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationModelWithDefaults

`func NewAuthorizationModelWithDefaults() *AuthorizationModel`

NewAuthorizationModelWithDefaults instantiates a new AuthorizationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorizationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationModel) SetId(v string)`

SetId sets Id field to given value.


### GetSchemaVersion

`func (o *AuthorizationModel) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *AuthorizationModel) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *AuthorizationModel) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetTypeDefinitions

`func (o *AuthorizationModel) GetTypeDefinitions() []TypeDefinition`

GetTypeDefinitions returns the TypeDefinitions field if non-nil, zero value otherwise.

### GetTypeDefinitionsOk

`func (o *AuthorizationModel) GetTypeDefinitionsOk() (*[]TypeDefinition, bool)`

GetTypeDefinitionsOk returns a tuple with the TypeDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDefinitions

`func (o *AuthorizationModel) SetTypeDefinitions(v []TypeDefinition)`

SetTypeDefinitions sets TypeDefinitions field to given value.


### GetConditions

`func (o *AuthorizationModel) GetConditions() map[string]Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizationModel) GetConditionsOk() (*map[string]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizationModel) SetConditions(v map[string]Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthorizationModel) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


