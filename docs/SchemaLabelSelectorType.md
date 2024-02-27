# SchemaLabelSelectorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expressions** | Pointer to **[]string** | x-displayName: \&quot;Selector Expression\&quot; x-required x-example: \&quot;region in (us-west1, us-west2),tier in (staging)\&quot; expressions contains the kubernetes style label expression for selections. | [optional] 

## Methods

### NewSchemaLabelSelectorType

`func NewSchemaLabelSelectorType() *SchemaLabelSelectorType`

NewSchemaLabelSelectorType instantiates a new SchemaLabelSelectorType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaLabelSelectorTypeWithDefaults

`func NewSchemaLabelSelectorTypeWithDefaults() *SchemaLabelSelectorType`

NewSchemaLabelSelectorTypeWithDefaults instantiates a new SchemaLabelSelectorType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpressions

`func (o *SchemaLabelSelectorType) GetExpressions() []string`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *SchemaLabelSelectorType) GetExpressionsOk() (*[]string, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *SchemaLabelSelectorType) SetExpressions(v []string)`

SetExpressions sets Expressions field to given value.

### HasExpressions

`func (o *SchemaLabelSelectorType) HasExpressions() bool`

HasExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


