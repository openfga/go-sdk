/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// AuthorizationModel struct for AuthorizationModel
type AuthorizationModel struct {
	Id              string                `json:"id"yaml:"id"`
	SchemaVersion   string                `json:"schema_version"yaml:"schema_version"`
	TypeDefinitions []TypeDefinition      `json:"type_definitions"yaml:"type_definitions"`
	Conditions      *map[string]Condition `json:"conditions,omitempty"yaml:"conditions,omitempty"`
}

// NewAuthorizationModel instantiates a new AuthorizationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationModel(id string, schemaVersion string, typeDefinitions []TypeDefinition) *AuthorizationModel {
	this := AuthorizationModel{}
	this.Id = id
	this.SchemaVersion = schemaVersion
	this.TypeDefinitions = typeDefinitions
	return &this
}

// NewAuthorizationModelWithDefaults instantiates a new AuthorizationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationModelWithDefaults() *AuthorizationModel {
	this := AuthorizationModel{}
	return &this
}

// GetId returns the Id field value
func (o *AuthorizationModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthorizationModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthorizationModel) SetId(v string) {
	o.Id = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *AuthorizationModel) GetSchemaVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *AuthorizationModel) GetSchemaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *AuthorizationModel) SetSchemaVersion(v string) {
	o.SchemaVersion = v
}

// GetTypeDefinitions returns the TypeDefinitions field value
func (o *AuthorizationModel) GetTypeDefinitions() []TypeDefinition {
	if o == nil {
		var ret []TypeDefinition
		return ret
	}

	return o.TypeDefinitions
}

// GetTypeDefinitionsOk returns a tuple with the TypeDefinitions field value
// and a boolean to check if the value has been set.
func (o *AuthorizationModel) GetTypeDefinitionsOk() (*[]TypeDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeDefinitions, true
}

// SetTypeDefinitions sets field value
func (o *AuthorizationModel) SetTypeDefinitions(v []TypeDefinition) {
	o.TypeDefinitions = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AuthorizationModel) GetConditions() map[string]Condition {
	if o == nil || o.Conditions == nil {
		var ret map[string]Condition
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationModel) GetConditionsOk() (*map[string]Condition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AuthorizationModel) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given map[string]Condition and assigns it to the Conditions field.
func (o *AuthorizationModel) SetConditions(v map[string]Condition) {
	o.Conditions = &v
}

func (o AuthorizationModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schema_version"] = o.SchemaVersion
	toSerialize["type_definitions"] = o.TypeDefinitions
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationModel struct {
	value *AuthorizationModel
	isSet bool
}

func (v NullableAuthorizationModel) Get() *AuthorizationModel {
	return v.value
}

func (v *NullableAuthorizationModel) Set(val *AuthorizationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationModel(val *AuthorizationModel) *NullableAuthorizationModel {
	return &NullableAuthorizationModel{value: val, isSet: true}
}

func (v NullableAuthorizationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
