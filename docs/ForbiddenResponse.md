# ForbiddenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**AuthErrorCode**](AuthErrorCode.md) |  | [optional] [default to AUTHERRORCODE_NO_AUTH_ERROR]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewForbiddenResponse

`func NewForbiddenResponse() *ForbiddenResponse`

NewForbiddenResponse instantiates a new ForbiddenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenResponseWithDefaults

`func NewForbiddenResponseWithDefaults() *ForbiddenResponse`

NewForbiddenResponseWithDefaults instantiates a new ForbiddenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ForbiddenResponse) GetCode() AuthErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ForbiddenResponse) GetCodeOk() (*AuthErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ForbiddenResponse) SetCode(v AuthErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *ForbiddenResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ForbiddenResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ForbiddenResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ForbiddenResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ForbiddenResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


