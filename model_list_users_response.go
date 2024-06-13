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

// ListUsersResponse struct for ListUsersResponse
type ListUsersResponse struct {
	Users []User `json:"users"yaml:"users"`
}

// NewListUsersResponse instantiates a new ListUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsersResponse(users []User) *ListUsersResponse {
	this := ListUsersResponse{}
	this.Users = users
	return &this
}

// NewListUsersResponseWithDefaults instantiates a new ListUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsersResponseWithDefaults() *ListUsersResponse {
	this := ListUsersResponse{}
	return &this
}

// GetUsers returns the Users field value
func (o *ListUsersResponse) GetUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ListUsersResponse) GetUsersOk() (*[]User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *ListUsersResponse) SetUsers(v []User) {
	o.Users = v
}

func (o ListUsersResponse) MarshalJSON() ([]byte, error) {
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

type NullableListUsersResponse struct {
	value *ListUsersResponse
	isSet bool
}

func (v NullableListUsersResponse) Get() *ListUsersResponse {
	return v.value
}

func (v *NullableListUsersResponse) Set(val *ListUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsersResponse(val *ListUsersResponse) *NullableListUsersResponse {
	return &NullableListUsersResponse{value: val, isSet: true}
}

func (v NullableListUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
