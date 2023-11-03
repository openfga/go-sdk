# Assertion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | [**CheckRequestTupleKey**](CheckRequestTupleKey.md) |  | 
**Expectation** | **bool** |  | 

## Methods

### NewAssertion

`func NewAssertion(tupleKey CheckRequestTupleKey, expectation bool, ) *Assertion`

NewAssertion instantiates a new Assertion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssertionWithDefaults

`func NewAssertionWithDefaults() *Assertion`

NewAssertionWithDefaults instantiates a new Assertion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTupleKey

`func (o *Assertion) GetTupleKey() CheckRequestTupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *Assertion) GetTupleKeyOk() (*CheckRequestTupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *Assertion) SetTupleKey(v CheckRequestTupleKey)`

SetTupleKey sets TupleKey field to given value.


### GetExpectation

`func (o *Assertion) GetExpectation() bool`

GetExpectation returns the Expectation field if non-nil, zero value otherwise.

### GetExpectationOk

`func (o *Assertion) GetExpectationOk() (*bool, bool)`

GetExpectationOk returns a tuple with the Expectation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectation

`func (o *Assertion) SetExpectation(v bool)`

SetExpectation sets Expectation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


