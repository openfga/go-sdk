/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// UsersetTree A UsersetTree contains the result of an Expansion.
type UsersetTree struct {
	Root *Node `json:"root,omitempty"`
}

// NewUsersetTree instantiates a new UsersetTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersetTree() *UsersetTree {
	this := UsersetTree{}
	return &this
}

// NewUsersetTreeWithDefaults instantiates a new UsersetTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersetTreeWithDefaults() *UsersetTree {
	this := UsersetTree{}
	return &this
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *UsersetTree) GetRoot() Node {
	if o == nil || o.Root == nil {
		var ret Node
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersetTree) GetRootOk() (*Node, bool) {
	if o == nil || o.Root == nil {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *UsersetTree) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given Node and assigns it to the Root field.
func (o *UsersetTree) SetRoot(v Node) {
	o.Root = &v
}

func (o UsersetTree) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Root != nil {
		toSerialize["root"] = o.Root
	}
	return json.Marshal(toSerialize)
}

type NullableUsersetTree struct {
	value *UsersetTree
	isSet bool
}

func (v NullableUsersetTree) Get() *UsersetTree {
	return v.value
}

func (v *NullableUsersetTree) Set(val *UsersetTree) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersetTree) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersetTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersetTree(val *UsersetTree) *NullableUsersetTree {
	return &NullableUsersetTree{value: val, isSet: true}
}

func (v NullableUsersetTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersetTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
