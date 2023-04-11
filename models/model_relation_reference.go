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

// RelationReference RelationReference represents a relation of a particular object type (e.g. 'document#viewer').
type RelationReference struct {
	Type     string                  `json:"type"`
	Relation *string                 `json:"relation,omitempty"`
	Wildcard *map[string]interface{} `json:"wildcard,omitempty"`
}

// NewRelationReference instantiates a new RelationReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationReference(type_ string) *RelationReference {
	this := RelationReference{}
	this.Type = type_
	return &this
}

// NewRelationReferenceWithDefaults instantiates a new RelationReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationReferenceWithDefaults() *RelationReference {
	this := RelationReference{}
	return &this
}

// GetType returns the Type field value
func (o *RelationReference) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RelationReference) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RelationReference) SetType(v string) {
	o.Type = v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *RelationReference) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationReference) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *RelationReference) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *RelationReference) SetRelation(v string) {
	o.Relation = &v
}

// GetWildcard returns the Wildcard field value if set, zero value otherwise.
func (o *RelationReference) GetWildcard() map[string]interface{} {
	if o == nil || o.Wildcard == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Wildcard
}

// GetWildcardOk returns a tuple with the Wildcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationReference) GetWildcardOk() (*map[string]interface{}, bool) {
	if o == nil || o.Wildcard == nil {
		return nil, false
	}
	return o.Wildcard, true
}

// HasWildcard returns a boolean if a field has been set.
func (o *RelationReference) HasWildcard() bool {
	if o != nil && o.Wildcard != nil {
		return true
	}

	return false
}

// SetWildcard gets a reference to the given map[string]interface{} and assigns it to the Wildcard field.
func (o *RelationReference) SetWildcard(v map[string]interface{}) {
	o.Wildcard = &v
}

func (o RelationReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}
	if o.Wildcard != nil {
		toSerialize["wildcard"] = o.Wildcard
	}
	return json.Marshal(toSerialize)
}

type NullableRelationReference struct {
	value *RelationReference
	isSet bool
}

func (v NullableRelationReference) Get() *RelationReference {
	return v.value
}

func (v *NullableRelationReference) Set(val *RelationReference) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationReference) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationReference(val *RelationReference) *NullableRelationReference {
	return &NullableRelationReference{value: val, isSet: true}
}

func (v NullableRelationReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
