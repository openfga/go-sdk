# UnauthenticatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**ErrorCode**](ErrorCode.md) |  | [optional] [default to ERRORCODE_NO_ERROR]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewUnauthenticatedResponse

`func NewUnauthenticatedResponse() *UnauthenticatedResponse`

NewUnauthenticatedResponse instantiates a new UnauthenticatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthenticatedResponseWithDefaults

`func NewUnauthenticatedResponseWithDefaults() *UnauthenticatedResponse`

NewUnauthenticatedResponseWithDefaults instantiates a new UnauthenticatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnauthenticatedResponse) GetCode() ErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnauthenticatedResponse) GetCodeOk() (*ErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnauthenticatedResponse) SetCode(v ErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnauthenticatedResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *UnauthenticatedResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthenticatedResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthenticatedResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnauthenticatedResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


