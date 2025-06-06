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

// Users struct for Users
type Users struct {
	Users []string `json:"users" yaml:"users"`
}

// NewUsers instantiates a new Users object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsers(users []string) *Users {
	this := Users{}
	this.Users = users
	return &this
}

// NewUsersWithDefaults instantiates a new Users object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersWithDefaults() *Users {
	this := Users{}
	return &this
}

// GetUsers returns the Users field value
func (o *Users) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *Users) GetUsersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *Users) SetUsers(v []string) {
	o.Users = v
}

func (o Users) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableUsers struct {
	value *Users
	isSet bool
}

func (v NullableUsers) Get() *Users {
	return v.value
}

func (v *NullableUsers) Set(val *Users) {
	v.value = val
	v.isSet = true
}

func (v NullableUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsers(val *Users) *NullableUsers {
	return &NullableUsers{value: val, isSet: true}
}

func (v NullableUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
