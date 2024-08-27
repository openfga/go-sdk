# ValidationErrorMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**ErrorCode**](ErrorCode.md) |  | [optional] [default to ERRORCODE_NO_ERROR]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewValidationErrorMessageResponse

`func NewValidationErrorMessageResponse() *ValidationErrorMessageResponse`

NewValidationErrorMessageResponse instantiates a new ValidationErrorMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorMessageResponseWithDefaults

`func NewValidationErrorMessageResponseWithDefaults() *ValidationErrorMessageResponse`

NewValidationErrorMessageResponseWithDefaults instantiates a new ValidationErrorMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ValidationErrorMessageResponse) GetCode() ErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationErrorMessageResponse) GetCodeOk() (*ErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationErrorMessageResponse) SetCode(v ErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *ValidationErrorMessageResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ValidationErrorMessageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationErrorMessageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationErrorMessageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationErrorMessageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


