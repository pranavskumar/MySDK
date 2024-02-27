# ViewsoriginPoolCreateSpecType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedOptions** | Pointer to [**OriginPoolOriginPoolAdvancedOptions**](OriginPoolOriginPoolAdvancedOptions.md) |  | [optional] 
**AutomaticPort** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**EndpointSelection** | Pointer to [**ClusterEndpointSelectionPolicy**](ClusterEndpointSelectionPolicy.md) |  | [optional] [default to DISTRIBUTED]
**HealthCheckPort** | Pointer to **int64** | Exclusive with [same_as_endpoint_port]  Port used for performing health check  Validation Rules:   ves.io.schema.rules.uint32.lte: 65535  | [optional] 
**Healthcheck** | Pointer to [**[]SchemaviewsObjectRefType**](SchemaviewsObjectRefType.md) |  Reference to healthcheck configuration objects  Validation Rules:   ves.io.schema.rules.repeated.max_items: 4  | [optional] 
**LbPort** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**LoadbalancerAlgorithm** | Pointer to [**ClusterLoadbalancerAlgorithm**](ClusterLoadbalancerAlgorithm.md) |  | [optional] [default to ROUND_ROBIN]
**NoTls** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**OriginServers** | Pointer to [**[]OriginPoolOriginServerType**](OriginPoolOriginServerType.md) |  List of origin servers in this pool  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.repeated.max_items: 32   ves.io.schema.rules.repeated.min_items: 1   ves.io.schema.rules.repeated.unique: true  | [optional] 
**Port** | Pointer to **int64** | Exclusive with [automatic_port lb_port]  Endpoint service is available on this port  Example: &#x60; \&quot;9080\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.gte: 1   ves.io.schema.rules.uint32.lte: 65535  | [optional] 
**SameAsEndpointPort** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**UseTls** | Pointer to [**OriginPoolUpstreamTlsParameters**](OriginPoolUpstreamTlsParameters.md) |  | [optional] 

## Methods

### NewViewsoriginPoolCreateSpecType

`func NewViewsoriginPoolCreateSpecType() *ViewsoriginPoolCreateSpecType`

NewViewsoriginPoolCreateSpecType instantiates a new ViewsoriginPoolCreateSpecType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsoriginPoolCreateSpecTypeWithDefaults

`func NewViewsoriginPoolCreateSpecTypeWithDefaults() *ViewsoriginPoolCreateSpecType`

NewViewsoriginPoolCreateSpecTypeWithDefaults instantiates a new ViewsoriginPoolCreateSpecType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedOptions

`func (o *ViewsoriginPoolCreateSpecType) GetAdvancedOptions() OriginPoolOriginPoolAdvancedOptions`

GetAdvancedOptions returns the AdvancedOptions field if non-nil, zero value otherwise.

### GetAdvancedOptionsOk

`func (o *ViewsoriginPoolCreateSpecType) GetAdvancedOptionsOk() (*OriginPoolOriginPoolAdvancedOptions, bool)`

GetAdvancedOptionsOk returns a tuple with the AdvancedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedOptions

`func (o *ViewsoriginPoolCreateSpecType) SetAdvancedOptions(v OriginPoolOriginPoolAdvancedOptions)`

SetAdvancedOptions sets AdvancedOptions field to given value.

### HasAdvancedOptions

`func (o *ViewsoriginPoolCreateSpecType) HasAdvancedOptions() bool`

HasAdvancedOptions returns a boolean if a field has been set.

### GetAutomaticPort

`func (o *ViewsoriginPoolCreateSpecType) GetAutomaticPort() map[string]interface{}`

GetAutomaticPort returns the AutomaticPort field if non-nil, zero value otherwise.

### GetAutomaticPortOk

`func (o *ViewsoriginPoolCreateSpecType) GetAutomaticPortOk() (*map[string]interface{}, bool)`

GetAutomaticPortOk returns a tuple with the AutomaticPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticPort

`func (o *ViewsoriginPoolCreateSpecType) SetAutomaticPort(v map[string]interface{})`

SetAutomaticPort sets AutomaticPort field to given value.

### HasAutomaticPort

`func (o *ViewsoriginPoolCreateSpecType) HasAutomaticPort() bool`

HasAutomaticPort returns a boolean if a field has been set.

### GetEndpointSelection

`func (o *ViewsoriginPoolCreateSpecType) GetEndpointSelection() ClusterEndpointSelectionPolicy`

GetEndpointSelection returns the EndpointSelection field if non-nil, zero value otherwise.

### GetEndpointSelectionOk

`func (o *ViewsoriginPoolCreateSpecType) GetEndpointSelectionOk() (*ClusterEndpointSelectionPolicy, bool)`

GetEndpointSelectionOk returns a tuple with the EndpointSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSelection

`func (o *ViewsoriginPoolCreateSpecType) SetEndpointSelection(v ClusterEndpointSelectionPolicy)`

SetEndpointSelection sets EndpointSelection field to given value.

### HasEndpointSelection

`func (o *ViewsoriginPoolCreateSpecType) HasEndpointSelection() bool`

HasEndpointSelection returns a boolean if a field has been set.

### GetHealthCheckPort

`func (o *ViewsoriginPoolCreateSpecType) GetHealthCheckPort() int64`

GetHealthCheckPort returns the HealthCheckPort field if non-nil, zero value otherwise.

### GetHealthCheckPortOk

`func (o *ViewsoriginPoolCreateSpecType) GetHealthCheckPortOk() (*int64, bool)`

GetHealthCheckPortOk returns a tuple with the HealthCheckPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPort

`func (o *ViewsoriginPoolCreateSpecType) SetHealthCheckPort(v int64)`

SetHealthCheckPort sets HealthCheckPort field to given value.

### HasHealthCheckPort

`func (o *ViewsoriginPoolCreateSpecType) HasHealthCheckPort() bool`

HasHealthCheckPort returns a boolean if a field has been set.

### GetHealthcheck

`func (o *ViewsoriginPoolCreateSpecType) GetHealthcheck() []SchemaviewsObjectRefType`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *ViewsoriginPoolCreateSpecType) GetHealthcheckOk() (*[]SchemaviewsObjectRefType, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *ViewsoriginPoolCreateSpecType) SetHealthcheck(v []SchemaviewsObjectRefType)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *ViewsoriginPoolCreateSpecType) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### GetLbPort

`func (o *ViewsoriginPoolCreateSpecType) GetLbPort() map[string]interface{}`

GetLbPort returns the LbPort field if non-nil, zero value otherwise.

### GetLbPortOk

`func (o *ViewsoriginPoolCreateSpecType) GetLbPortOk() (*map[string]interface{}, bool)`

GetLbPortOk returns a tuple with the LbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbPort

`func (o *ViewsoriginPoolCreateSpecType) SetLbPort(v map[string]interface{})`

SetLbPort sets LbPort field to given value.

### HasLbPort

`func (o *ViewsoriginPoolCreateSpecType) HasLbPort() bool`

HasLbPort returns a boolean if a field has been set.

### GetLoadbalancerAlgorithm

`func (o *ViewsoriginPoolCreateSpecType) GetLoadbalancerAlgorithm() ClusterLoadbalancerAlgorithm`

GetLoadbalancerAlgorithm returns the LoadbalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadbalancerAlgorithmOk

`func (o *ViewsoriginPoolCreateSpecType) GetLoadbalancerAlgorithmOk() (*ClusterLoadbalancerAlgorithm, bool)`

GetLoadbalancerAlgorithmOk returns a tuple with the LoadbalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancerAlgorithm

`func (o *ViewsoriginPoolCreateSpecType) SetLoadbalancerAlgorithm(v ClusterLoadbalancerAlgorithm)`

SetLoadbalancerAlgorithm sets LoadbalancerAlgorithm field to given value.

### HasLoadbalancerAlgorithm

`func (o *ViewsoriginPoolCreateSpecType) HasLoadbalancerAlgorithm() bool`

HasLoadbalancerAlgorithm returns a boolean if a field has been set.

### GetNoTls

`func (o *ViewsoriginPoolCreateSpecType) GetNoTls() map[string]interface{}`

GetNoTls returns the NoTls field if non-nil, zero value otherwise.

### GetNoTlsOk

`func (o *ViewsoriginPoolCreateSpecType) GetNoTlsOk() (*map[string]interface{}, bool)`

GetNoTlsOk returns a tuple with the NoTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTls

`func (o *ViewsoriginPoolCreateSpecType) SetNoTls(v map[string]interface{})`

SetNoTls sets NoTls field to given value.

### HasNoTls

`func (o *ViewsoriginPoolCreateSpecType) HasNoTls() bool`

HasNoTls returns a boolean if a field has been set.

### GetOriginServers

`func (o *ViewsoriginPoolCreateSpecType) GetOriginServers() []OriginPoolOriginServerType`

GetOriginServers returns the OriginServers field if non-nil, zero value otherwise.

### GetOriginServersOk

`func (o *ViewsoriginPoolCreateSpecType) GetOriginServersOk() (*[]OriginPoolOriginServerType, bool)`

GetOriginServersOk returns a tuple with the OriginServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginServers

`func (o *ViewsoriginPoolCreateSpecType) SetOriginServers(v []OriginPoolOriginServerType)`

SetOriginServers sets OriginServers field to given value.

### HasOriginServers

`func (o *ViewsoriginPoolCreateSpecType) HasOriginServers() bool`

HasOriginServers returns a boolean if a field has been set.

### GetPort

`func (o *ViewsoriginPoolCreateSpecType) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ViewsoriginPoolCreateSpecType) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ViewsoriginPoolCreateSpecType) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ViewsoriginPoolCreateSpecType) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSameAsEndpointPort

`func (o *ViewsoriginPoolCreateSpecType) GetSameAsEndpointPort() map[string]interface{}`

GetSameAsEndpointPort returns the SameAsEndpointPort field if non-nil, zero value otherwise.

### GetSameAsEndpointPortOk

`func (o *ViewsoriginPoolCreateSpecType) GetSameAsEndpointPortOk() (*map[string]interface{}, bool)`

GetSameAsEndpointPortOk returns a tuple with the SameAsEndpointPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameAsEndpointPort

`func (o *ViewsoriginPoolCreateSpecType) SetSameAsEndpointPort(v map[string]interface{})`

SetSameAsEndpointPort sets SameAsEndpointPort field to given value.

### HasSameAsEndpointPort

`func (o *ViewsoriginPoolCreateSpecType) HasSameAsEndpointPort() bool`

HasSameAsEndpointPort returns a boolean if a field has been set.

### GetUseTls

`func (o *ViewsoriginPoolCreateSpecType) GetUseTls() OriginPoolUpstreamTlsParameters`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *ViewsoriginPoolCreateSpecType) GetUseTlsOk() (*OriginPoolUpstreamTlsParameters, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *ViewsoriginPoolCreateSpecType) SetUseTls(v OriginPoolUpstreamTlsParameters)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *ViewsoriginPoolCreateSpecType) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


