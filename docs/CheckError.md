# CheckError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputError** | Pointer to [**ErrorCode**](ErrorCode.md) |  | [optional] [default to ERRORCODE_NO_ERROR]
**InternalError** | Pointer to [**InternalErrorCode**](InternalErrorCode.md) |  | [optional] [default to INTERNALERRORCODE_NO_INTERNAL_ERROR]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewCheckError

`func NewCheckError() *CheckError`

NewCheckError instantiates a new CheckError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckErrorWithDefaults

`func NewCheckErrorWithDefaults() *CheckError`

NewCheckErrorWithDefaults instantiates a new CheckError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputError

`func (o *CheckError) GetInputError() ErrorCode`

GetInputError returns the InputError field if non-nil, zero value otherwise.

### GetInputErrorOk

`func (o *CheckError) GetInputErrorOk() (*ErrorCode, bool)`

GetInputErrorOk returns a tuple with the InputError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputError

`func (o *CheckError) SetInputError(v ErrorCode)`

SetInputError sets InputError field to given value.

### HasInputError

`func (o *CheckError) HasInputError() bool`

HasInputError returns a boolean if a field has been set.

### GetInternalError

`func (o *CheckError) GetInternalError() InternalErrorCode`

GetInternalError returns the InternalError field if non-nil, zero value otherwise.

### GetInternalErrorOk

`func (o *CheckError) GetInternalErrorOk() (*InternalErrorCode, bool)`

GetInternalErrorOk returns a tuple with the InternalError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalError

`func (o *CheckError) SetInternalError(v InternalErrorCode)`

SetInternalError sets InternalError field to given value.

### HasInternalError

`func (o *CheckError) HasInternalError() bool`

HasInternalError returns a boolean if a field has been set.

### GetMessage

`func (o *CheckError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CheckError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CheckError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CheckError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


