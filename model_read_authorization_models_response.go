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

// ReadAuthorizationModelsResponse struct for ReadAuthorizationModelsResponse
type ReadAuthorizationModelsResponse struct {
	AuthorizationModels []AuthorizationModel `json:"authorization_models" yaml:"authorization_models"`
	// The continuation token will be empty if there are no more models.
	ContinuationToken *string `json:"continuation_token,omitempty" yaml:"continuation_token,omitempty"`
}

// NewReadAuthorizationModelsResponse instantiates a new ReadAuthorizationModelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAuthorizationModelsResponse(authorizationModels []AuthorizationModel) *ReadAuthorizationModelsResponse {
	this := ReadAuthorizationModelsResponse{}
	this.AuthorizationModels = authorizationModels
	return &this
}

// NewReadAuthorizationModelsResponseWithDefaults instantiates a new ReadAuthorizationModelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAuthorizationModelsResponseWithDefaults() *ReadAuthorizationModelsResponse {
	this := ReadAuthorizationModelsResponse{}
	return &this
}

// GetAuthorizationModels returns the AuthorizationModels field value
func (o *ReadAuthorizationModelsResponse) GetAuthorizationModels() []AuthorizationModel {
	if o == nil {
		var ret []AuthorizationModel
		return ret
	}

	return o.AuthorizationModels
}

// GetAuthorizationModelsOk returns a tuple with the AuthorizationModels field value
// and a boolean to check if the value has been set.
func (o *ReadAuthorizationModelsResponse) GetAuthorizationModelsOk() (*[]AuthorizationModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationModels, true
}

// SetAuthorizationModels sets field value
func (o *ReadAuthorizationModelsResponse) SetAuthorizationModels(v []AuthorizationModel) {
	o.AuthorizationModels = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *ReadAuthorizationModelsResponse) GetContinuationToken() string {
	if o == nil || o.ContinuationToken == nil {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAuthorizationModelsResponse) GetContinuationTokenOk() (*string, bool) {
	if o == nil || o.ContinuationToken == nil {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *ReadAuthorizationModelsResponse) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken != nil {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *ReadAuthorizationModelsResponse) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

func (o ReadAuthorizationModelsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorization_models"] = o.AuthorizationModels
	if o.ContinuationToken != nil {
		toSerialize["continuation_token"] = o.ContinuationToken
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

type NullableReadAuthorizationModelsResponse struct {
	value *ReadAuthorizationModelsResponse
	isSet bool
}

func (v NullableReadAuthorizationModelsResponse) Get() *ReadAuthorizationModelsResponse {
	return v.value
}

func (v *NullableReadAuthorizationModelsResponse) Set(val *ReadAuthorizationModelsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAuthorizationModelsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAuthorizationModelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAuthorizationModelsResponse(val *ReadAuthorizationModelsResponse) *NullableReadAuthorizationModelsResponse {
	return &NullableReadAuthorizationModelsResponse{value: val, isSet: true}
}

func (v NullableReadAuthorizationModelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAuthorizationModelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
