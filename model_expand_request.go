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

// ExpandRequest struct for ExpandRequest
type ExpandRequest struct {
	TupleKey             ExpandRequestTupleKey  `json:"tuple_key" yaml:"tuple_key"`
	AuthorizationModelId *string                `json:"authorization_model_id,omitempty" yaml:"authorization_model_id,omitempty"`
	Consistency          *ConsistencyPreference `json:"consistency,omitempty" yaml:"consistency,omitempty"`
	ContextualTuples     *ContextualTupleKeys   `json:"contextual_tuples,omitempty" yaml:"contextual_tuples,omitempty"`
}

// NewExpandRequest instantiates a new ExpandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandRequest(tupleKey ExpandRequestTupleKey) *ExpandRequest {
	this := ExpandRequest{}
	this.TupleKey = tupleKey
	var consistency = CONSISTENCYPREFERENCE_UNSPECIFIED
	this.Consistency = &consistency
	return &this
}

// NewExpandRequestWithDefaults instantiates a new ExpandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandRequestWithDefaults() *ExpandRequest {
	this := ExpandRequest{}
	var consistency = CONSISTENCYPREFERENCE_UNSPECIFIED
	this.Consistency = &consistency
	return &this
}

// GetTupleKey returns the TupleKey field value
func (o *ExpandRequest) GetTupleKey() ExpandRequestTupleKey {
	if o == nil {
		var ret ExpandRequestTupleKey
		return ret
	}

	return o.TupleKey
}

// GetTupleKeyOk returns a tuple with the TupleKey field value
// and a boolean to check if the value has been set.
func (o *ExpandRequest) GetTupleKeyOk() (*ExpandRequestTupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TupleKey, true
}

// SetTupleKey sets field value
func (o *ExpandRequest) SetTupleKey(v ExpandRequestTupleKey) {
	o.TupleKey = v
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *ExpandRequest) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandRequest) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *ExpandRequest) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *ExpandRequest) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

// GetConsistency returns the Consistency field value if set, zero value otherwise.
func (o *ExpandRequest) GetConsistency() ConsistencyPreference {
	if o == nil || o.Consistency == nil {
		var ret ConsistencyPreference
		return ret
	}
	return *o.Consistency
}

// GetConsistencyOk returns a tuple with the Consistency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandRequest) GetConsistencyOk() (*ConsistencyPreference, bool) {
	if o == nil || o.Consistency == nil {
		return nil, false
	}
	return o.Consistency, true
}

// HasConsistency returns a boolean if a field has been set.
func (o *ExpandRequest) HasConsistency() bool {
	if o != nil && o.Consistency != nil {
		return true
	}

	return false
}

// SetConsistency gets a reference to the given ConsistencyPreference and assigns it to the Consistency field.
func (o *ExpandRequest) SetConsistency(v ConsistencyPreference) {
	o.Consistency = &v
}

// GetContextualTuples returns the ContextualTuples field value if set, zero value otherwise.
func (o *ExpandRequest) GetContextualTuples() ContextualTupleKeys {
	if o == nil || o.ContextualTuples == nil {
		var ret ContextualTupleKeys
		return ret
	}
	return *o.ContextualTuples
}

// GetContextualTuplesOk returns a tuple with the ContextualTuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandRequest) GetContextualTuplesOk() (*ContextualTupleKeys, bool) {
	if o == nil || o.ContextualTuples == nil {
		return nil, false
	}
	return o.ContextualTuples, true
}

// HasContextualTuples returns a boolean if a field has been set.
func (o *ExpandRequest) HasContextualTuples() bool {
	if o != nil && o.ContextualTuples != nil {
		return true
	}

	return false
}

// SetContextualTuples gets a reference to the given ContextualTupleKeys and assigns it to the ContextualTuples field.
func (o *ExpandRequest) SetContextualTuples(v ContextualTupleKeys) {
	o.ContextualTuples = &v
}

func (o ExpandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuple_key"] = o.TupleKey
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	if o.Consistency != nil {
		toSerialize["consistency"] = o.Consistency
	}
	if o.ContextualTuples != nil {
		toSerialize["contextual_tuples"] = o.ContextualTuples
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

type NullableExpandRequest struct {
	value *ExpandRequest
	isSet bool
}

func (v NullableExpandRequest) Get() *ExpandRequest {
	return v.value
}

func (v *NullableExpandRequest) Set(val *ExpandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandRequest(val *ExpandRequest) *NullableExpandRequest {
	return &NullableExpandRequest{value: val, isSet: true}
}

func (v NullableExpandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
