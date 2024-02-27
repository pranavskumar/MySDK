# OriginPoolOriginPoolDefaultSubset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSubset** | Pointer to **map[string]interface{}** |  List of key-value pairs that define default subset.  which gets used when route specifies no metadata or no subset matching the metadata exists.  Example: &#x60; \&quot;key:value\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.map.max_pairs: 32  | [optional] 

## Methods

### NewOriginPoolOriginPoolDefaultSubset

`func NewOriginPoolOriginPoolDefaultSubset() *OriginPoolOriginPoolDefaultSubset`

NewOriginPoolOriginPoolDefaultSubset instantiates a new OriginPoolOriginPoolDefaultSubset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginPoolDefaultSubsetWithDefaults

`func NewOriginPoolOriginPoolDefaultSubsetWithDefaults() *OriginPoolOriginPoolDefaultSubset`

NewOriginPoolOriginPoolDefaultSubsetWithDefaults instantiates a new OriginPoolOriginPoolDefaultSubset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSubset

`func (o *OriginPoolOriginPoolDefaultSubset) GetDefaultSubset() map[string]interface{}`

GetDefaultSubset returns the DefaultSubset field if non-nil, zero value otherwise.

### GetDefaultSubsetOk

`func (o *OriginPoolOriginPoolDefaultSubset) GetDefaultSubsetOk() (*map[string]interface{}, bool)`

GetDefaultSubsetOk returns a tuple with the DefaultSubset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubset

`func (o *OriginPoolOriginPoolDefaultSubset) SetDefaultSubset(v map[string]interface{})`

SetDefaultSubset sets DefaultSubset field to given value.

### HasDefaultSubset

`func (o *OriginPoolOriginPoolDefaultSubset) HasDefaultSubset() bool`

HasDefaultSubset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


