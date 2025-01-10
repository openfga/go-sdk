# Assertion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TupleKey** | [**AssertionTupleKey**](AssertionTupleKey.md) |  | 
**Expectation** | **bool** |  | 
**ContextualTuples** | Pointer to [**[]TupleKey**](TupleKey.md) |  | [optional] 
**Context** | Pointer to **map[string]interface{}** | Additional request context that will be used to evaluate any ABAC conditions encountered in the query evaluation. | [optional] 

## Methods

### NewAssertion

`func NewAssertion(tupleKey AssertionTupleKey, expectation bool, ) *Assertion`

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

`func (o *Assertion) GetTupleKey() AssertionTupleKey`

GetTupleKey returns the TupleKey field if non-nil, zero value otherwise.

### GetTupleKeyOk

`func (o *Assertion) GetTupleKeyOk() (*AssertionTupleKey, bool)`

GetTupleKeyOk returns a tuple with the TupleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTupleKey

`func (o *Assertion) SetTupleKey(v AssertionTupleKey)`

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


### GetContextualTuples

`func (o *Assertion) GetContextualTuples() []TupleKey`

GetContextualTuples returns the ContextualTuples field if non-nil, zero value otherwise.

### GetContextualTuplesOk

`func (o *Assertion) GetContextualTuplesOk() (*[]TupleKey, bool)`

GetContextualTuplesOk returns a tuple with the ContextualTuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualTuples

`func (o *Assertion) SetContextualTuples(v []TupleKey)`

SetContextualTuples sets ContextualTuples field to given value.

### HasContextualTuples

`func (o *Assertion) HasContextualTuples() bool`

HasContextualTuples returns a boolean if a field has been set.

### GetContext

`func (o *Assertion) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Assertion) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Assertion) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *Assertion) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


