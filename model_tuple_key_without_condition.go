/**
 * Go SDK for OpenFGA
 *
 * API version: 1.x
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://openfga.dev/community
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"bytes"

	"encoding/json"
)

// TupleKeyWithoutCondition struct for TupleKeyWithoutCondition
type TupleKeyWithoutCondition struct {
	User     string `json:"user" yaml:"user"`
	Relation string `json:"relation" yaml:"relation"`
	Object   string `json:"object" yaml:"object"`
}

// NewTupleKeyWithoutCondition instantiates a new TupleKeyWithoutCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTupleKeyWithoutCondition(user string, relation string, object string) *TupleKeyWithoutCondition {
	this := TupleKeyWithoutCondition{}
	this.User = user
	this.Relation = relation
	this.Object = object
	return &this
}

// NewTupleKeyWithoutConditionWithDefaults instantiates a new TupleKeyWithoutCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleKeyWithoutConditionWithDefaults() *TupleKeyWithoutCondition {
	this := TupleKeyWithoutCondition{}
	return &this
}

// GetUser returns the User field value
func (o *TupleKeyWithoutCondition) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TupleKeyWithoutCondition) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TupleKeyWithoutCondition) SetUser(v string) {
	o.User = v
}

// GetRelation returns the Relation field value
func (o *TupleKeyWithoutCondition) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *TupleKeyWithoutCondition) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value
func (o *TupleKeyWithoutCondition) SetRelation(v string) {
	o.Relation = v
}

// GetObject returns the Object field value
func (o *TupleKeyWithoutCondition) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *TupleKeyWithoutCondition) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *TupleKeyWithoutCondition) SetObject(v string) {
	o.Object = v
}

func (o TupleKeyWithoutCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	toSerialize["relation"] = o.Relation
	toSerialize["object"] = o.Object
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableTupleKeyWithoutCondition struct {
	value *TupleKeyWithoutCondition
	isSet bool
}

func (v NullableTupleKeyWithoutCondition) Get() *TupleKeyWithoutCondition {
	return v.value
}

func (v *NullableTupleKeyWithoutCondition) Set(val *TupleKeyWithoutCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleKeyWithoutCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleKeyWithoutCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleKeyWithoutCondition(val *TupleKeyWithoutCondition) *NullableTupleKeyWithoutCondition {
	return &NullableTupleKeyWithoutCondition{value: val, isSet: true}
}

func (v NullableTupleKeyWithoutCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleKeyWithoutCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
