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

// CreateStoreRequest struct for CreateStoreRequest
type CreateStoreRequest struct {
	Name string `json:"name" yaml:"name"`
}

// NewCreateStoreRequest instantiates a new CreateStoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStoreRequest(name string) *CreateStoreRequest {
	this := CreateStoreRequest{}
	this.Name = name
	return &this
}

// NewCreateStoreRequestWithDefaults instantiates a new CreateStoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStoreRequestWithDefaults() *CreateStoreRequest {
	this := CreateStoreRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateStoreRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateStoreRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateStoreRequest) SetName(v string) {
	o.Name = v
}

func (o CreateStoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableCreateStoreRequest struct {
	value *CreateStoreRequest
	isSet bool
}

func (v NullableCreateStoreRequest) Get() *CreateStoreRequest {
	return v.value
}

func (v *NullableCreateStoreRequest) Set(val *CreateStoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStoreRequest(val *CreateStoreRequest) *NullableCreateStoreRequest {
	return &NullableCreateStoreRequest{value: val, isSet: true}
}

func (v NullableCreateStoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
