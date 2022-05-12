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

// UsersetTreeDifference struct for UsersetTreeDifference
type UsersetTreeDifference struct {
	Base     *Node `json:"base,omitempty"`
	Subtract *Node `json:"subtract,omitempty"`
}

// NewUsersetTreeDifference instantiates a new UsersetTreeDifference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersetTreeDifference() *UsersetTreeDifference {
	this := UsersetTreeDifference{}
	return &this
}

// NewUsersetTreeDifferenceWithDefaults instantiates a new UsersetTreeDifference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersetTreeDifferenceWithDefaults() *UsersetTreeDifference {
	this := UsersetTreeDifference{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *UsersetTreeDifference) GetBase() Node {
	if o == nil || o.Base == nil {
		var ret Node
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersetTreeDifference) GetBaseOk() (*Node, bool) {
	if o == nil || o.Base == nil {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *UsersetTreeDifference) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given Node and assigns it to the Base field.
func (o *UsersetTreeDifference) SetBase(v Node) {
	o.Base = &v
}

// GetSubtract returns the Subtract field value if set, zero value otherwise.
func (o *UsersetTreeDifference) GetSubtract() Node {
	if o == nil || o.Subtract == nil {
		var ret Node
		return ret
	}
	return *o.Subtract
}

// GetSubtractOk returns a tuple with the Subtract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersetTreeDifference) GetSubtractOk() (*Node, bool) {
	if o == nil || o.Subtract == nil {
		return nil, false
	}
	return o.Subtract, true
}

// HasSubtract returns a boolean if a field has been set.
func (o *UsersetTreeDifference) HasSubtract() bool {
	if o != nil && o.Subtract != nil {
		return true
	}

	return false
}

// SetSubtract gets a reference to the given Node and assigns it to the Subtract field.
func (o *UsersetTreeDifference) SetSubtract(v Node) {
	o.Subtract = &v
}

func (o UsersetTreeDifference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base != nil {
		toSerialize["base"] = o.Base
	}
	if o.Subtract != nil {
		toSerialize["subtract"] = o.Subtract
	}
	return json.Marshal(toSerialize)
}

type NullableUsersetTreeDifference struct {
	value *UsersetTreeDifference
	isSet bool
}

func (v NullableUsersetTreeDifference) Get() *UsersetTreeDifference {
	return v.value
}

func (v *NullableUsersetTreeDifference) Set(val *UsersetTreeDifference) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersetTreeDifference) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersetTreeDifference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersetTreeDifference(val *UsersetTreeDifference) *NullableUsersetTreeDifference {
	return &NullableUsersetTreeDifference{value: val, isSet: true}
}

func (v NullableUsersetTreeDifference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersetTreeDifference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
