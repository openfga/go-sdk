# UnprocessableContentMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**UnprocessableContentErrorCode**](UnprocessableContentErrorCode.md) |  | [optional] [default to UNPROCESSABLECONTENTERRORCODE_NO_THROTTLED_ERROR_CODE]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewUnprocessableContentMessageResponse

`func NewUnprocessableContentMessageResponse() *UnprocessableContentMessageResponse`

NewUnprocessableContentMessageResponse instantiates a new UnprocessableContentMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnprocessableContentMessageResponseWithDefaults

`func NewUnprocessableContentMessageResponseWithDefaults() *UnprocessableContentMessageResponse`

NewUnprocessableContentMessageResponseWithDefaults instantiates a new UnprocessableContentMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnprocessableContentMessageResponse) GetCode() UnprocessableContentErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnprocessableContentMessageResponse) GetCodeOk() (*UnprocessableContentErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnprocessableContentMessageResponse) SetCode(v UnprocessableContentErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnprocessableContentMessageResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *UnprocessableContentMessageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnprocessableContentMessageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnprocessableContentMessageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnprocessableContentMessageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


