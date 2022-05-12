# InternalErrorMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**InternalErrorCode**](InternalErrorCode.md) |  | [optional] [default to NO_INTERNAL_ERROR]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalErrorMessageResponse

`func NewInternalErrorMessageResponse() *InternalErrorMessageResponse`

NewInternalErrorMessageResponse instantiates a new InternalErrorMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalErrorMessageResponseWithDefaults

`func NewInternalErrorMessageResponseWithDefaults() *InternalErrorMessageResponse`

NewInternalErrorMessageResponseWithDefaults instantiates a new InternalErrorMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InternalErrorMessageResponse) GetCode() InternalErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InternalErrorMessageResponse) GetCodeOk() (*InternalErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InternalErrorMessageResponse) SetCode(v InternalErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *InternalErrorMessageResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *InternalErrorMessageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InternalErrorMessageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InternalErrorMessageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InternalErrorMessageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


