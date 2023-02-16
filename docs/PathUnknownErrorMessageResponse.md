# PathUnknownErrorMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**NotFoundErrorCode**](NotFoundErrorCode.md) |  | [optional] [default to NO_NOT_FOUND_ERROR]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewPathUnknownErrorMessageResponse

`func NewPathUnknownErrorMessageResponse() *PathUnknownErrorMessageResponse`

NewPathUnknownErrorMessageResponse instantiates a new PathUnknownErrorMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathUnknownErrorMessageResponseWithDefaults

`func NewPathUnknownErrorMessageResponseWithDefaults() *PathUnknownErrorMessageResponse`

NewPathUnknownErrorMessageResponseWithDefaults instantiates a new PathUnknownErrorMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PathUnknownErrorMessageResponse) GetCode() NotFoundErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PathUnknownErrorMessageResponse) GetCodeOk() (*NotFoundErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PathUnknownErrorMessageResponse) SetCode(v NotFoundErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *PathUnknownErrorMessageResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *PathUnknownErrorMessageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PathUnknownErrorMessageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PathUnknownErrorMessageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PathUnknownErrorMessageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


