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

// PathUnknownErrorMessageResponse struct for PathUnknownErrorMessageResponse
type PathUnknownErrorMessageResponse struct {
	Code    *NotFoundErrorCode `json:"code,omitempty" yaml:"code,omitempty"`
	Message *string            `json:"message,omitempty" yaml:"message,omitempty"`
}

// NewPathUnknownErrorMessageResponse instantiates a new PathUnknownErrorMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathUnknownErrorMessageResponse() *PathUnknownErrorMessageResponse {
	this := PathUnknownErrorMessageResponse{}
	var code = NOTFOUNDERRORCODE_NO_NOT_FOUND_ERROR
	this.Code = &code
	return &this
}

// NewPathUnknownErrorMessageResponseWithDefaults instantiates a new PathUnknownErrorMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathUnknownErrorMessageResponseWithDefaults() *PathUnknownErrorMessageResponse {
	this := PathUnknownErrorMessageResponse{}
	var code = NOTFOUNDERRORCODE_NO_NOT_FOUND_ERROR
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PathUnknownErrorMessageResponse) GetCode() NotFoundErrorCode {
	if o == nil || o.Code == nil {
		var ret NotFoundErrorCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathUnknownErrorMessageResponse) GetCodeOk() (*NotFoundErrorCode, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PathUnknownErrorMessageResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given NotFoundErrorCode and assigns it to the Code field.
func (o *PathUnknownErrorMessageResponse) SetCode(v NotFoundErrorCode) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PathUnknownErrorMessageResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathUnknownErrorMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PathUnknownErrorMessageResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PathUnknownErrorMessageResponse) SetMessage(v string) {
	o.Message = &v
}

func (o PathUnknownErrorMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
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

type NullablePathUnknownErrorMessageResponse struct {
	value *PathUnknownErrorMessageResponse
	isSet bool
}

func (v NullablePathUnknownErrorMessageResponse) Get() *PathUnknownErrorMessageResponse {
	return v.value
}

func (v *NullablePathUnknownErrorMessageResponse) Set(val *PathUnknownErrorMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePathUnknownErrorMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePathUnknownErrorMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathUnknownErrorMessageResponse(val *PathUnknownErrorMessageResponse) *NullablePathUnknownErrorMessageResponse {
	return &NullablePathUnknownErrorMessageResponse{value: val, isSet: true}
}

func (v NullablePathUnknownErrorMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathUnknownErrorMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
