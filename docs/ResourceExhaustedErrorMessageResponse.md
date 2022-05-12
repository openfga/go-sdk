# ResourceExhaustedErrorMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**ResourceExhaustedErrorCode**](ResourceExhaustedErrorCode.md) |  | [optional] [default to NO_RESOURCE_EXHAUSTED_ERROR]
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceExhaustedErrorMessageResponse

`func NewResourceExhaustedErrorMessageResponse() *ResourceExhaustedErrorMessageResponse`

NewResourceExhaustedErrorMessageResponse instantiates a new ResourceExhaustedErrorMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceExhaustedErrorMessageResponseWithDefaults

`func NewResourceExhaustedErrorMessageResponseWithDefaults() *ResourceExhaustedErrorMessageResponse`

NewResourceExhaustedErrorMessageResponseWithDefaults instantiates a new ResourceExhaustedErrorMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ResourceExhaustedErrorMessageResponse) GetCode() ResourceExhaustedErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResourceExhaustedErrorMessageResponse) GetCodeOk() (*ResourceExhaustedErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResourceExhaustedErrorMessageResponse) SetCode(v ResourceExhaustedErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResourceExhaustedErrorMessageResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ResourceExhaustedErrorMessageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceExhaustedErrorMessageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceExhaustedErrorMessageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceExhaustedErrorMessageResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


