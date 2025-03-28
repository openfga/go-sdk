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

// Usersets struct for Usersets
type Usersets struct {
	Child []Userset `json:"child" yaml:"child"`
}

// NewUsersets instantiates a new Usersets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersets(child []Userset) *Usersets {
	this := Usersets{}
	this.Child = child
	return &this
}

// NewUsersetsWithDefaults instantiates a new Usersets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersetsWithDefaults() *Usersets {
	this := Usersets{}
	return &this
}

// GetChild returns the Child field value
func (o *Usersets) GetChild() []Userset {
	if o == nil {
		var ret []Userset
		return ret
	}

	return o.Child
}

// GetChildOk returns a tuple with the Child field value
// and a boolean to check if the value has been set.
func (o *Usersets) GetChildOk() (*[]Userset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Child, true
}

// SetChild sets field value
func (o *Usersets) SetChild(v []Userset) {
	o.Child = v
}

func (o Usersets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["child"] = o.Child
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableUsersets struct {
	value *Usersets
	isSet bool
}

func (v NullableUsersets) Get() *Usersets {
	return v.value
}

func (v *NullableUsersets) Set(val *Usersets) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersets) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersets(val *Usersets) *NullableUsersets {
	return &NullableUsersets{value: val, isSet: true}
}

func (v NullableUsersets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
