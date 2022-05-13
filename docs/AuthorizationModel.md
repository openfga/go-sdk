# AuthorizationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TypeDefinitions** | Pointer to [**[]TypeDefinition**](TypeDefinition.md) |  | [optional] 

## Methods

### NewAuthorizationModel

`func NewAuthorizationModel() *AuthorizationModel`

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

### HasId

`func (o *AuthorizationModel) HasId() bool`

HasId returns a boolean if a field has been set.

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

### HasTypeDefinitions

`func (o *AuthorizationModel) HasTypeDefinitions() bool`

HasTypeDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


