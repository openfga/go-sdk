/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// WriteRequestParams struct for WriteRequestParams
type WriteRequestParams struct {
	Writes               *TupleKeys `json:"writes,omitempty"`
	Deletes              *TupleKeys `json:"deletes,omitempty"`
	AuthorizationModelId *string    `json:"authorization_model_id,omitempty"`
}

// NewWriteRequestParams instantiates a new WriteRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteRequestParams() *WriteRequestParams {
	this := WriteRequestParams{}
	return &this
}

// NewWriteRequestParamsWithDefaults instantiates a new WriteRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteRequestParamsWithDefaults() *WriteRequestParams {
	this := WriteRequestParams{}
	return &this
}

// GetWrites returns the Writes field value if set, zero value otherwise.
func (o *WriteRequestParams) GetWrites() TupleKeys {
	if o == nil || o.Writes == nil {
		var ret TupleKeys
		return ret
	}
	return *o.Writes
}

// GetWritesOk returns a tuple with the Writes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteRequestParams) GetWritesOk() (*TupleKeys, bool) {
	if o == nil || o.Writes == nil {
		return nil, false
	}
	return o.Writes, true
}

// HasWrites returns a boolean if a field has been set.
func (o *WriteRequestParams) HasWrites() bool {
	if o != nil && o.Writes != nil {
		return true
	}

	return false
}

// SetWrites gets a reference to the given TupleKeys and assigns it to the Writes field.
func (o *WriteRequestParams) SetWrites(v TupleKeys) {
	o.Writes = &v
}

// GetDeletes returns the Deletes field value if set, zero value otherwise.
func (o *WriteRequestParams) GetDeletes() TupleKeys {
	if o == nil || o.Deletes == nil {
		var ret TupleKeys
		return ret
	}
	return *o.Deletes
}

// GetDeletesOk returns a tuple with the Deletes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteRequestParams) GetDeletesOk() (*TupleKeys, bool) {
	if o == nil || o.Deletes == nil {
		return nil, false
	}
	return o.Deletes, true
}

// HasDeletes returns a boolean if a field has been set.
func (o *WriteRequestParams) HasDeletes() bool {
	if o != nil && o.Deletes != nil {
		return true
	}

	return false
}

// SetDeletes gets a reference to the given TupleKeys and assigns it to the Deletes field.
func (o *WriteRequestParams) SetDeletes(v TupleKeys) {
	o.Deletes = &v
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *WriteRequestParams) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WriteRequestParams) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *WriteRequestParams) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *WriteRequestParams) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

func (o WriteRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Writes != nil {
		toSerialize["writes"] = o.Writes
	}
	if o.Deletes != nil {
		toSerialize["deletes"] = o.Deletes
	}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	return json.Marshal(toSerialize)
}

type NullableWriteRequestParams struct {
	value *WriteRequestParams
	isSet bool
}

func (v NullableWriteRequestParams) Get() *WriteRequestParams {
	return v.value
}

func (v *NullableWriteRequestParams) Set(val *WriteRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteRequestParams(val *WriteRequestParams) *NullableWriteRequestParams {
	return &NullableWriteRequestParams{value: val, isSet: true}
}

func (v NullableWriteRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
