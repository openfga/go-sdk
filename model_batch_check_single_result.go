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

// BatchCheckSingleResult struct for BatchCheckSingleResult
type BatchCheckSingleResult struct {
	Allowed *bool       `json:"allowed,omitempty" yaml:"allowed,omitempty"`
	Error   *CheckError `json:"error,omitempty" yaml:"error,omitempty"`
}

// NewBatchCheckSingleResult instantiates a new BatchCheckSingleResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchCheckSingleResult() *BatchCheckSingleResult {
	this := BatchCheckSingleResult{}
	return &this
}

// NewBatchCheckSingleResultWithDefaults instantiates a new BatchCheckSingleResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchCheckSingleResultWithDefaults() *BatchCheckSingleResult {
	this := BatchCheckSingleResult{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *BatchCheckSingleResult) GetAllowed() bool {
	if o == nil || o.Allowed == nil {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchCheckSingleResult) GetAllowedOk() (*bool, bool) {
	if o == nil || o.Allowed == nil {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *BatchCheckSingleResult) HasAllowed() bool {
	if o != nil && o.Allowed != nil {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *BatchCheckSingleResult) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BatchCheckSingleResult) GetError() CheckError {
	if o == nil || o.Error == nil {
		var ret CheckError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchCheckSingleResult) GetErrorOk() (*CheckError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BatchCheckSingleResult) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given CheckError and assigns it to the Error field.
func (o *BatchCheckSingleResult) SetError(v CheckError) {
	o.Error = &v
}

func (o BatchCheckSingleResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allowed != nil {
		toSerialize["allowed"] = o.Allowed
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
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

type NullableBatchCheckSingleResult struct {
	value *BatchCheckSingleResult
	isSet bool
}

func (v NullableBatchCheckSingleResult) Get() *BatchCheckSingleResult {
	return v.value
}

func (v *NullableBatchCheckSingleResult) Set(val *BatchCheckSingleResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchCheckSingleResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchCheckSingleResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchCheckSingleResult(val *BatchCheckSingleResult) *NullableBatchCheckSingleResult {
	return &NullableBatchCheckSingleResult{value: val, isSet: true}
}

func (v NullableBatchCheckSingleResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchCheckSingleResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
