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
	"bytes"

	"encoding/json"
)

// RelationshipCondition struct for RelationshipCondition
type RelationshipCondition struct {
	// A reference (by name) of the relationship condition defined in the authorization model.
	Name string `json:"name"yaml:"name"`
	// Additional context/data to persist along with the condition. The keys must match the parameters defined by the condition, and the value types must match the parameter type definitions.
	Context *map[string]interface{} `json:"context,omitempty"yaml:"context,omitempty"`
}

// NewRelationshipCondition instantiates a new RelationshipCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipCondition(name string) *RelationshipCondition {
	this := RelationshipCondition{}
	this.Name = name
	return &this
}

// NewRelationshipConditionWithDefaults instantiates a new RelationshipCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipConditionWithDefaults() *RelationshipCondition {
	this := RelationshipCondition{}
	return &this
}

// GetName returns the Name field value
func (o *RelationshipCondition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RelationshipCondition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RelationshipCondition) SetName(v string) {
	o.Name = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RelationshipCondition) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipCondition) GetContextOk() (*map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RelationshipCondition) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *RelationshipCondition) SetContext(v map[string]interface{}) {
	o.Context = &v
}

func (o RelationshipCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableRelationshipCondition struct {
	value *RelationshipCondition
	isSet bool
}

func (v NullableRelationshipCondition) Get() *RelationshipCondition {
	return v.value
}

func (v *NullableRelationshipCondition) Set(val *RelationshipCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationshipCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationshipCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationshipCondition(val *RelationshipCondition) *NullableRelationshipCondition {
	return &NullableRelationshipCondition{value: val, isSet: true}
}

func (v NullableRelationshipCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationshipCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
