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

// ErrorCode the model 'ErrorCode'
type ErrorCode string

// List of ErrorCode
const (
	ERRORCODE_NO_ERROR                                         ErrorCode = "no_error"
	ERRORCODE_VALIDATION_ERROR                                 ErrorCode = "validation_error"
	ERRORCODE_AUTHORIZATION_MODEL_NOT_FOUND                    ErrorCode = "authorization_model_not_found"
	ERRORCODE_AUTHORIZATION_MODEL_RESOLUTION_TOO_COMPLEX       ErrorCode = "authorization_model_resolution_too_complex"
	ERRORCODE_INVALID_WRITE_INPUT                              ErrorCode = "invalid_write_input"
	ERRORCODE_CANNOT_ALLOW_DUPLICATE_TUPLES_IN_ONE_REQUEST     ErrorCode = "cannot_allow_duplicate_tuples_in_one_request"
	ERRORCODE_CANNOT_ALLOW_DUPLICATE_TYPES_IN_ONE_REQUEST      ErrorCode = "cannot_allow_duplicate_types_in_one_request"
	ERRORCODE_CANNOT_ALLOW_MULTIPLE_REFERENCES_TO_ONE_RELATION ErrorCode = "cannot_allow_multiple_references_to_one_relation"
	ERRORCODE_INVALID_CONTINUATION_TOKEN                       ErrorCode = "invalid_continuation_token"
	ERRORCODE_INVALID_TUPLE_SET                                ErrorCode = "invalid_tuple_set"
	ERRORCODE_INVALID_CHECK_INPUT                              ErrorCode = "invalid_check_input"
	ERRORCODE_INVALID_EXPAND_INPUT                             ErrorCode = "invalid_expand_input"
	ERRORCODE_UNSUPPORTED_USER_SET                             ErrorCode = "unsupported_user_set"
	ERRORCODE_INVALID_OBJECT_FORMAT                            ErrorCode = "invalid_object_format"
	ERRORCODE_WRITE_FAILED_DUE_TO_INVALID_INPUT                ErrorCode = "write_failed_due_to_invalid_input"
	ERRORCODE_AUTHORIZATION_MODEL_ASSERTIONS_NOT_FOUND         ErrorCode = "authorization_model_assertions_not_found"
	ERRORCODE_LATEST_AUTHORIZATION_MODEL_NOT_FOUND             ErrorCode = "latest_authorization_model_not_found"
	ERRORCODE_TYPE_NOT_FOUND                                   ErrorCode = "type_not_found"
	ERRORCODE_RELATION_NOT_FOUND                               ErrorCode = "relation_not_found"
	ERRORCODE_EMPTY_RELATION_DEFINITION                        ErrorCode = "empty_relation_definition"
	ERRORCODE_INVALID_USER                                     ErrorCode = "invalid_user"
	ERRORCODE_INVALID_TUPLE                                    ErrorCode = "invalid_tuple"
	ERRORCODE_UNKNOWN_RELATION                                 ErrorCode = "unknown_relation"
	ERRORCODE_STORE_ID_INVALID_LENGTH                          ErrorCode = "store_id_invalid_length"
	ERRORCODE_ASSERTIONS_TOO_MANY_ITEMS                        ErrorCode = "assertions_too_many_items"
	ERRORCODE_ID_TOO_LONG                                      ErrorCode = "id_too_long"
	ERRORCODE_AUTHORIZATION_MODEL_ID_TOO_LONG                  ErrorCode = "authorization_model_id_too_long"
	ERRORCODE_TUPLE_KEY_VALUE_NOT_SPECIFIED                    ErrorCode = "tuple_key_value_not_specified"
	ERRORCODE_TUPLE_KEYS_TOO_MANY_OR_TOO_FEW_ITEMS             ErrorCode = "tuple_keys_too_many_or_too_few_items"
	ERRORCODE_PAGE_SIZE_INVALID                                ErrorCode = "page_size_invalid"
	ERRORCODE_PARAM_MISSING_VALUE                              ErrorCode = "param_missing_value"
	ERRORCODE_DIFFERENCE_BASE_MISSING_VALUE                    ErrorCode = "difference_base_missing_value"
	ERRORCODE_SUBTRACT_BASE_MISSING_VALUE                      ErrorCode = "subtract_base_missing_value"
	ERRORCODE_OBJECT_TOO_LONG                                  ErrorCode = "object_too_long"
	ERRORCODE_RELATION_TOO_LONG                                ErrorCode = "relation_too_long"
	ERRORCODE_TYPE_DEFINITIONS_TOO_FEW_ITEMS                   ErrorCode = "type_definitions_too_few_items"
	ERRORCODE_TYPE_INVALID_LENGTH                              ErrorCode = "type_invalid_length"
	ERRORCODE_TYPE_INVALID_PATTERN                             ErrorCode = "type_invalid_pattern"
	ERRORCODE_RELATIONS_TOO_FEW_ITEMS                          ErrorCode = "relations_too_few_items"
	ERRORCODE_RELATIONS_TOO_LONG                               ErrorCode = "relations_too_long"
	ERRORCODE_RELATIONS_INVALID_PATTERN                        ErrorCode = "relations_invalid_pattern"
	ERRORCODE_OBJECT_INVALID_PATTERN                           ErrorCode = "object_invalid_pattern"
	ERRORCODE_QUERY_STRING_TYPE_CONTINUATION_TOKEN_MISMATCH    ErrorCode = "query_string_type_continuation_token_mismatch"
	ERRORCODE_EXCEEDED_ENTITY_LIMIT                            ErrorCode = "exceeded_entity_limit"
	ERRORCODE_INVALID_CONTEXTUAL_TUPLE                         ErrorCode = "invalid_contextual_tuple"
	ERRORCODE_DUPLICATE_CONTEXTUAL_TUPLE                       ErrorCode = "duplicate_contextual_tuple"
	ERRORCODE_INVALID_AUTHORIZATION_MODEL                      ErrorCode = "invalid_authorization_model"
	ERRORCODE_UNSUPPORTED_SCHEMA_VERSION                       ErrorCode = "unsupported_schema_version"
	ERRORCODE_CANCELLED                                        ErrorCode = "cancelled"
	ERRORCODE_INVALID_START_TIME                               ErrorCode = "invalid_start_time"
)

var allowedErrorCodeEnumValues = []ErrorCode{
	"no_error",
	"validation_error",
	"authorization_model_not_found",
	"authorization_model_resolution_too_complex",
	"invalid_write_input",
	"cannot_allow_duplicate_tuples_in_one_request",
	"cannot_allow_duplicate_types_in_one_request",
	"cannot_allow_multiple_references_to_one_relation",
	"invalid_continuation_token",
	"invalid_tuple_set",
	"invalid_check_input",
	"invalid_expand_input",
	"unsupported_user_set",
	"invalid_object_format",
	"write_failed_due_to_invalid_input",
	"authorization_model_assertions_not_found",
	"latest_authorization_model_not_found",
	"type_not_found",
	"relation_not_found",
	"empty_relation_definition",
	"invalid_user",
	"invalid_tuple",
	"unknown_relation",
	"store_id_invalid_length",
	"assertions_too_many_items",
	"id_too_long",
	"authorization_model_id_too_long",
	"tuple_key_value_not_specified",
	"tuple_keys_too_many_or_too_few_items",
	"page_size_invalid",
	"param_missing_value",
	"difference_base_missing_value",
	"subtract_base_missing_value",
	"object_too_long",
	"relation_too_long",
	"type_definitions_too_few_items",
	"type_invalid_length",
	"type_invalid_pattern",
	"relations_too_few_items",
	"relations_too_long",
	"relations_invalid_pattern",
	"object_invalid_pattern",
	"query_string_type_continuation_token_mismatch",
	"exceeded_entity_limit",
	"invalid_contextual_tuple",
	"duplicate_contextual_tuple",
	"invalid_authorization_model",
	"unsupported_schema_version",
	"cancelled",
	"invalid_start_time",
}

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range allowedErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// NewErrorCodeFromValue returns a pointer to a valid ErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorCodeFromValue(v string) (*ErrorCode, error) {
	ev := ErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorCode: valid values are %v", v, allowedErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorCode) IsValid() bool {
	for _, existing := range allowedErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
