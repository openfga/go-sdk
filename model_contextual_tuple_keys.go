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

// ContextualTupleKeys struct for ContextualTupleKeys
type ContextualTupleKeys struct {
	TupleKeys []TupleKey `json:"tuple_keys"yaml:"tuple_keys"`
}

// NewContextualTupleKeys instantiates a new ContextualTupleKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextualTupleKeys(tupleKeys []TupleKey) *ContextualTupleKeys {
	this := ContextualTupleKeys{}
	this.TupleKeys = tupleKeys
	return &this
}

// NewContextualTupleKeysWithDefaults instantiates a new ContextualTupleKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextualTupleKeysWithDefaults() *ContextualTupleKeys {
	this := ContextualTupleKeys{}
	return &this
}

// GetTupleKeys returns the TupleKeys field value
func (o *ContextualTupleKeys) GetTupleKeys() []TupleKey {
	if o == nil {
		var ret []TupleKey
		return ret
	}

	return o.TupleKeys
}

// GetTupleKeysOk returns a tuple with the TupleKeys field value
// and a boolean to check if the value has been set.
func (o *ContextualTupleKeys) GetTupleKeysOk() (*[]TupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TupleKeys, true
}

// SetTupleKeys sets field value
func (o *ContextualTupleKeys) SetTupleKeys(v []TupleKey) {
	o.TupleKeys = v
}

func (o ContextualTupleKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuple_keys"] = o.TupleKeys
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableContextualTupleKeys struct {
	value *ContextualTupleKeys
	isSet bool
}

func (v NullableContextualTupleKeys) Get() *ContextualTupleKeys {
	return v.value
}

func (v *NullableContextualTupleKeys) Set(val *ContextualTupleKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableContextualTupleKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableContextualTupleKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextualTupleKeys(val *ContextualTupleKeys) *NullableContextualTupleKeys {
	return &NullableContextualTupleKeys{value: val, isSet: true}
}

func (v NullableContextualTupleKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextualTupleKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
