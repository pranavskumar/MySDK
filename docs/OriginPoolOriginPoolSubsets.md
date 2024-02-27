# OriginPoolOriginPoolSubsets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnyEndpoint** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**DefaultSubset** | Pointer to [**OriginPoolOriginPoolDefaultSubset**](OriginPoolOriginPoolDefaultSubset.md) |  | [optional] 
**EndpointSubsets** | Pointer to [**[]ClusterEndpointSubsetSelectorType**](ClusterEndpointSubsetSelectorType.md) |  List of subset class. Subsets class is defined using list of keys. Every unique combination of values of these keys form a subset withing the class.  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.repeated.max_items: 32  | [optional] 
**FailRequest** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 

## Methods

### NewOriginPoolOriginPoolSubsets

`func NewOriginPoolOriginPoolSubsets() *OriginPoolOriginPoolSubsets`

NewOriginPoolOriginPoolSubsets instantiates a new OriginPoolOriginPoolSubsets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginPoolSubsetsWithDefaults

`func NewOriginPoolOriginPoolSubsetsWithDefaults() *OriginPoolOriginPoolSubsets`

NewOriginPoolOriginPoolSubsetsWithDefaults instantiates a new OriginPoolOriginPoolSubsets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnyEndpoint

`func (o *OriginPoolOriginPoolSubsets) GetAnyEndpoint() map[string]interface{}`

GetAnyEndpoint returns the AnyEndpoint field if non-nil, zero value otherwise.

### GetAnyEndpointOk

`func (o *OriginPoolOriginPoolSubsets) GetAnyEndpointOk() (*map[string]interface{}, bool)`

GetAnyEndpointOk returns a tuple with the AnyEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyEndpoint

`func (o *OriginPoolOriginPoolSubsets) SetAnyEndpoint(v map[string]interface{})`

SetAnyEndpoint sets AnyEndpoint field to given value.

### HasAnyEndpoint

`func (o *OriginPoolOriginPoolSubsets) HasAnyEndpoint() bool`

HasAnyEndpoint returns a boolean if a field has been set.

### GetDefaultSubset

`func (o *OriginPoolOriginPoolSubsets) GetDefaultSubset() OriginPoolOriginPoolDefaultSubset`

GetDefaultSubset returns the DefaultSubset field if non-nil, zero value otherwise.

### GetDefaultSubsetOk

`func (o *OriginPoolOriginPoolSubsets) GetDefaultSubsetOk() (*OriginPoolOriginPoolDefaultSubset, bool)`

GetDefaultSubsetOk returns a tuple with the DefaultSubset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubset

`func (o *OriginPoolOriginPoolSubsets) SetDefaultSubset(v OriginPoolOriginPoolDefaultSubset)`

SetDefaultSubset sets DefaultSubset field to given value.

### HasDefaultSubset

`func (o *OriginPoolOriginPoolSubsets) HasDefaultSubset() bool`

HasDefaultSubset returns a boolean if a field has been set.

### GetEndpointSubsets

`func (o *OriginPoolOriginPoolSubsets) GetEndpointSubsets() []ClusterEndpointSubsetSelectorType`

GetEndpointSubsets returns the EndpointSubsets field if non-nil, zero value otherwise.

### GetEndpointSubsetsOk

`func (o *OriginPoolOriginPoolSubsets) GetEndpointSubsetsOk() (*[]ClusterEndpointSubsetSelectorType, bool)`

GetEndpointSubsetsOk returns a tuple with the EndpointSubsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSubsets

`func (o *OriginPoolOriginPoolSubsets) SetEndpointSubsets(v []ClusterEndpointSubsetSelectorType)`

SetEndpointSubsets sets EndpointSubsets field to given value.

### HasEndpointSubsets

`func (o *OriginPoolOriginPoolSubsets) HasEndpointSubsets() bool`

HasEndpointSubsets returns a boolean if a field has been set.

### GetFailRequest

`func (o *OriginPoolOriginPoolSubsets) GetFailRequest() map[string]interface{}`

GetFailRequest returns the FailRequest field if non-nil, zero value otherwise.

### GetFailRequestOk

`func (o *OriginPoolOriginPoolSubsets) GetFailRequestOk() (*map[string]interface{}, bool)`

GetFailRequestOk returns a tuple with the FailRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailRequest

`func (o *OriginPoolOriginPoolSubsets) SetFailRequest(v map[string]interface{})`

SetFailRequest sets FailRequest field to given value.

### HasFailRequest

`func (o *OriginPoolOriginPoolSubsets) HasFailRequest() bool`

HasFailRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


