# ClusterEndpointSubsetSelectorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** |  List of keys that define a cluster subset class.  Example: &#x60; \&quot;production\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.repeated.items.string.not_empty: true   ves.io.schema.rules.repeated.max_items: 16  | [optional] 

## Methods

### NewClusterEndpointSubsetSelectorType

`func NewClusterEndpointSubsetSelectorType() *ClusterEndpointSubsetSelectorType`

NewClusterEndpointSubsetSelectorType instantiates a new ClusterEndpointSubsetSelectorType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEndpointSubsetSelectorTypeWithDefaults

`func NewClusterEndpointSubsetSelectorTypeWithDefaults() *ClusterEndpointSubsetSelectorType`

NewClusterEndpointSubsetSelectorTypeWithDefaults instantiates a new ClusterEndpointSubsetSelectorType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ClusterEndpointSubsetSelectorType) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ClusterEndpointSubsetSelectorType) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ClusterEndpointSubsetSelectorType) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ClusterEndpointSubsetSelectorType) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


