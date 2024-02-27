# SchemaHashAlgorithms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HashAlgorithms** | Pointer to [**[]SchemaHashAlgorithm**](SchemaHashAlgorithm.md) |  Ordered list of hash algorithms to be used.  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.repeated.max_items: 4   ves.io.schema.rules.repeated.min_items: 1   ves.io.schema.rules.repeated.unique: true  | [optional] 

## Methods

### NewSchemaHashAlgorithms

`func NewSchemaHashAlgorithms() *SchemaHashAlgorithms`

NewSchemaHashAlgorithms instantiates a new SchemaHashAlgorithms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaHashAlgorithmsWithDefaults

`func NewSchemaHashAlgorithmsWithDefaults() *SchemaHashAlgorithms`

NewSchemaHashAlgorithmsWithDefaults instantiates a new SchemaHashAlgorithms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashAlgorithms

`func (o *SchemaHashAlgorithms) GetHashAlgorithms() []SchemaHashAlgorithm`

GetHashAlgorithms returns the HashAlgorithms field if non-nil, zero value otherwise.

### GetHashAlgorithmsOk

`func (o *SchemaHashAlgorithms) GetHashAlgorithmsOk() (*[]SchemaHashAlgorithm, bool)`

GetHashAlgorithmsOk returns a tuple with the HashAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithms

`func (o *SchemaHashAlgorithms) SetHashAlgorithms(v []SchemaHashAlgorithm)`

SetHashAlgorithms sets HashAlgorithms field to given value.

### HasHashAlgorithms

`func (o *SchemaHashAlgorithms) HasHashAlgorithms() bool`

HasHashAlgorithms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


