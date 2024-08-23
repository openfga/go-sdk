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

// UnprocessableContentMessageResponse struct for UnprocessableContentMessageResponse
type UnprocessableContentMessageResponse struct {
	Code    *UnprocessableContentErrorCode `json:"code,omitempty"yaml:"code,omitempty"`
	Message *string                        `json:"message,omitempty"yaml:"message,omitempty"`
}

// NewUnprocessableContentMessageResponse instantiates a new UnprocessableContentMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnprocessableContentMessageResponse() *UnprocessableContentMessageResponse {
	this := UnprocessableContentMessageResponse{}
	var code UnprocessableContentErrorCode = UNPROCESSABLECONTENTERRORCODE_NO_THROTTLED_ERROR_CODE
	this.Code = &code
	return &this
}

// NewUnprocessableContentMessageResponseWithDefaults instantiates a new UnprocessableContentMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnprocessableContentMessageResponseWithDefaults() *UnprocessableContentMessageResponse {
	this := UnprocessableContentMessageResponse{}
	var code UnprocessableContentErrorCode = UNPROCESSABLECONTENTERRORCODE_NO_THROTTLED_ERROR_CODE
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UnprocessableContentMessageResponse) GetCode() UnprocessableContentErrorCode {
	if o == nil || o.Code == nil {
		var ret UnprocessableContentErrorCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnprocessableContentMessageResponse) GetCodeOk() (*UnprocessableContentErrorCode, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UnprocessableContentMessageResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given UnprocessableContentErrorCode and assigns it to the Code field.
func (o *UnprocessableContentMessageResponse) SetCode(v UnprocessableContentErrorCode) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UnprocessableContentMessageResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnprocessableContentMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UnprocessableContentMessageResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UnprocessableContentMessageResponse) SetMessage(v string) {
	o.Message = &v
}

func (o UnprocessableContentMessageResponse) MarshalJSON() ([]byte, error) {
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

type NullableUnprocessableContentMessageResponse struct {
	value *UnprocessableContentMessageResponse
	isSet bool
}

func (v NullableUnprocessableContentMessageResponse) Get() *UnprocessableContentMessageResponse {
	return v.value
}

func (v *NullableUnprocessableContentMessageResponse) Set(val *UnprocessableContentMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnprocessableContentMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnprocessableContentMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnprocessableContentMessageResponse(val *UnprocessableContentMessageResponse) *NullableUnprocessableContentMessageResponse {
	return &NullableUnprocessableContentMessageResponse{value: val, isSet: true}
}

func (v NullableUnprocessableContentMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnprocessableContentMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
