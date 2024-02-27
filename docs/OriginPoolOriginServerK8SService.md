# OriginPoolOriginServerK8SService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsideNetwork** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**OutsideNetwork** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**ServiceName** | Pointer to **string** | Exclusive with []  K8s service name of the origin server, including the namespace and cluster-id  The format is servicename.namespace:cluster-id. If the servicename is \&quot;frontend\&quot;,  namespace is \&quot;speedtest\&quot; and cluster-id is \&quot;prod\&quot;, then you will enter \&quot;frontend.speedtest:prod\&quot;.  Both namespace and cluster-id are optional.  Example: &#x60; \&quot;matching.default:production\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.string.ves_service_namespace_name: true  | [optional] 
**SiteLocator** | Pointer to [**ViewsSiteLocator**](ViewsSiteLocator.md) |  | [optional] 
**Vk8sNetworks** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 

## Methods

### NewOriginPoolOriginServerK8SService

`func NewOriginPoolOriginServerK8SService() *OriginPoolOriginServerK8SService`

NewOriginPoolOriginServerK8SService instantiates a new OriginPoolOriginServerK8SService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerK8SServiceWithDefaults

`func NewOriginPoolOriginServerK8SServiceWithDefaults() *OriginPoolOriginServerK8SService`

NewOriginPoolOriginServerK8SServiceWithDefaults instantiates a new OriginPoolOriginServerK8SService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsideNetwork

`func (o *OriginPoolOriginServerK8SService) GetInsideNetwork() map[string]interface{}`

GetInsideNetwork returns the InsideNetwork field if non-nil, zero value otherwise.

### GetInsideNetworkOk

`func (o *OriginPoolOriginServerK8SService) GetInsideNetworkOk() (*map[string]interface{}, bool)`

GetInsideNetworkOk returns a tuple with the InsideNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsideNetwork

`func (o *OriginPoolOriginServerK8SService) SetInsideNetwork(v map[string]interface{})`

SetInsideNetwork sets InsideNetwork field to given value.

### HasInsideNetwork

`func (o *OriginPoolOriginServerK8SService) HasInsideNetwork() bool`

HasInsideNetwork returns a boolean if a field has been set.

### GetOutsideNetwork

`func (o *OriginPoolOriginServerK8SService) GetOutsideNetwork() map[string]interface{}`

GetOutsideNetwork returns the OutsideNetwork field if non-nil, zero value otherwise.

### GetOutsideNetworkOk

`func (o *OriginPoolOriginServerK8SService) GetOutsideNetworkOk() (*map[string]interface{}, bool)`

GetOutsideNetworkOk returns a tuple with the OutsideNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideNetwork

`func (o *OriginPoolOriginServerK8SService) SetOutsideNetwork(v map[string]interface{})`

SetOutsideNetwork sets OutsideNetwork field to given value.

### HasOutsideNetwork

`func (o *OriginPoolOriginServerK8SService) HasOutsideNetwork() bool`

HasOutsideNetwork returns a boolean if a field has been set.

### GetServiceName

`func (o *OriginPoolOriginServerK8SService) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *OriginPoolOriginServerK8SService) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *OriginPoolOriginServerK8SService) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *OriginPoolOriginServerK8SService) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSiteLocator

`func (o *OriginPoolOriginServerK8SService) GetSiteLocator() ViewsSiteLocator`

GetSiteLocator returns the SiteLocator field if non-nil, zero value otherwise.

### GetSiteLocatorOk

`func (o *OriginPoolOriginServerK8SService) GetSiteLocatorOk() (*ViewsSiteLocator, bool)`

GetSiteLocatorOk returns a tuple with the SiteLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteLocator

`func (o *OriginPoolOriginServerK8SService) SetSiteLocator(v ViewsSiteLocator)`

SetSiteLocator sets SiteLocator field to given value.

### HasSiteLocator

`func (o *OriginPoolOriginServerK8SService) HasSiteLocator() bool`

HasSiteLocator returns a boolean if a field has been set.

### GetVk8sNetworks

`func (o *OriginPoolOriginServerK8SService) GetVk8sNetworks() map[string]interface{}`

GetVk8sNetworks returns the Vk8sNetworks field if non-nil, zero value otherwise.

### GetVk8sNetworksOk

`func (o *OriginPoolOriginServerK8SService) GetVk8sNetworksOk() (*map[string]interface{}, bool)`

GetVk8sNetworksOk returns a tuple with the Vk8sNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVk8sNetworks

`func (o *OriginPoolOriginServerK8SService) SetVk8sNetworks(v map[string]interface{})`

SetVk8sNetworks sets Vk8sNetworks field to given value.

### HasVk8sNetworks

`func (o *OriginPoolOriginServerK8SService) HasVk8sNetworks() bool`

HasVk8sNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


