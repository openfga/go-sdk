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

// ReadChangesResponse struct for ReadChangesResponse
type ReadChangesResponse struct {
	Changes []TupleChange `json:"changes"yaml:"changes"`
	// The continuation token will be identical if there are no new changes.
	ContinuationToken *string `json:"continuation_token,omitempty"yaml:"continuation_token,omitempty"`
}

// NewReadChangesResponse instantiates a new ReadChangesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadChangesResponse(changes []TupleChange) *ReadChangesResponse {
	this := ReadChangesResponse{}
	this.Changes = changes
	return &this
}

// NewReadChangesResponseWithDefaults instantiates a new ReadChangesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadChangesResponseWithDefaults() *ReadChangesResponse {
	this := ReadChangesResponse{}
	return &this
}

// GetChanges returns the Changes field value
func (o *ReadChangesResponse) GetChanges() []TupleChange {
	if o == nil {
		var ret []TupleChange
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *ReadChangesResponse) GetChangesOk() (*[]TupleChange, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Changes, true
}

// SetChanges sets field value
func (o *ReadChangesResponse) SetChanges(v []TupleChange) {
	o.Changes = v
}

// GetContinuationToken returns the ContinuationToken field value if set, zero value otherwise.
func (o *ReadChangesResponse) GetContinuationToken() string {
	if o == nil || o.ContinuationToken == nil {
		var ret string
		return ret
	}
	return *o.ContinuationToken
}

// GetContinuationTokenOk returns a tuple with the ContinuationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadChangesResponse) GetContinuationTokenOk() (*string, bool) {
	if o == nil || o.ContinuationToken == nil {
		return nil, false
	}
	return o.ContinuationToken, true
}

// HasContinuationToken returns a boolean if a field has been set.
func (o *ReadChangesResponse) HasContinuationToken() bool {
	if o != nil && o.ContinuationToken != nil {
		return true
	}

	return false
}

// SetContinuationToken gets a reference to the given string and assigns it to the ContinuationToken field.
func (o *ReadChangesResponse) SetContinuationToken(v string) {
	o.ContinuationToken = &v
}

func (o ReadChangesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["changes"] = o.Changes
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

type NullableReadChangesResponse struct {
	value *ReadChangesResponse
	isSet bool
}

func (v NullableReadChangesResponse) Get() *ReadChangesResponse {
	return v.value
}

func (v *NullableReadChangesResponse) Set(val *ReadChangesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadChangesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadChangesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadChangesResponse(val *ReadChangesResponse) *NullableReadChangesResponse {
	return &NullableReadChangesResponse{value: val, isSet: true}
}

func (v NullableReadChangesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadChangesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
