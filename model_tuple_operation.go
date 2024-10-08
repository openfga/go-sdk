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
	"encoding/json"
	"fmt"
)

// TupleOperation the model 'TupleOperation'
type TupleOperation string

// List of TupleOperation
const (
	TUPLEOPERATION_WRITE  TupleOperation = "TUPLE_OPERATION_WRITE"
	TUPLEOPERATION_DELETE TupleOperation = "TUPLE_OPERATION_DELETE"
)

var allowedTupleOperationEnumValues = []TupleOperation{
	"TUPLE_OPERATION_WRITE",
	"TUPLE_OPERATION_DELETE",
}

func (v *TupleOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TupleOperation(value)
	for _, existing := range allowedTupleOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TupleOperation", value)
}

// NewTupleOperationFromValue returns a pointer to a valid TupleOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTupleOperationFromValue(v string) (*TupleOperation, error) {
	ev := TupleOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TupleOperation: valid values are %v", v, allowedTupleOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TupleOperation) IsValid() bool {
	for _, existing := range allowedTupleOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TupleOperation value
func (v TupleOperation) Ptr() *TupleOperation {
	return &v
}

type NullableTupleOperation struct {
	value *TupleOperation
	isSet bool
}

func (v NullableTupleOperation) Get() *TupleOperation {
	return v.value
}

func (v *NullableTupleOperation) Set(val *TupleOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleOperation(val *TupleOperation) *NullableTupleOperation {
	return &NullableTupleOperation{value: val, isSet: true}
}

func (v NullableTupleOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
