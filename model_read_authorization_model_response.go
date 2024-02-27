/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
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

// ReadAuthorizationModelResponse struct for ReadAuthorizationModelResponse
type ReadAuthorizationModelResponse struct {
	AuthorizationModel *AuthorizationModel `json:"authorization_model,omitempty"yaml:"authorization_model,omitempty"`
}

// NewReadAuthorizationModelResponse instantiates a new ReadAuthorizationModelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAuthorizationModelResponse() *ReadAuthorizationModelResponse {
	this := ReadAuthorizationModelResponse{}
	return &this
}

// NewReadAuthorizationModelResponseWithDefaults instantiates a new ReadAuthorizationModelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAuthorizationModelResponseWithDefaults() *ReadAuthorizationModelResponse {
	this := ReadAuthorizationModelResponse{}
	return &this
}

// GetAuthorizationModel returns the AuthorizationModel field value if set, zero value otherwise.
func (o *ReadAuthorizationModelResponse) GetAuthorizationModel() AuthorizationModel {
	if o == nil || o.AuthorizationModel == nil {
		var ret AuthorizationModel
		return ret
	}
	return *o.AuthorizationModel
}

// GetAuthorizationModelOk returns a tuple with the AuthorizationModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAuthorizationModelResponse) GetAuthorizationModelOk() (*AuthorizationModel, bool) {
	if o == nil || o.AuthorizationModel == nil {
		return nil, false
	}
	return o.AuthorizationModel, true
}

// HasAuthorizationModel returns a boolean if a field has been set.
func (o *ReadAuthorizationModelResponse) HasAuthorizationModel() bool {
	if o != nil && o.AuthorizationModel != nil {
		return true
	}

	return false
}

// SetAuthorizationModel gets a reference to the given AuthorizationModel and assigns it to the AuthorizationModel field.
func (o *ReadAuthorizationModelResponse) SetAuthorizationModel(v AuthorizationModel) {
	o.AuthorizationModel = &v
}

func (o ReadAuthorizationModelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationModel != nil {
		toSerialize["authorization_model"] = o.AuthorizationModel
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

type NullableReadAuthorizationModelResponse struct {
	value *ReadAuthorizationModelResponse
	isSet bool
}

func (v NullableReadAuthorizationModelResponse) Get() *ReadAuthorizationModelResponse {
	return v.value
}

func (v *NullableReadAuthorizationModelResponse) Set(val *ReadAuthorizationModelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAuthorizationModelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAuthorizationModelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAuthorizationModelResponse(val *ReadAuthorizationModelResponse) *NullableReadAuthorizationModelResponse {
	return &NullableReadAuthorizationModelResponse{value: val, isSet: true}
}

func (v NullableReadAuthorizationModelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAuthorizationModelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
