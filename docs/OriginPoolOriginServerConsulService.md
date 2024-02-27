# OriginPoolOriginServerConsulService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsideNetwork** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**OutsideNetwork** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**ServiceName** | Pointer to **string** |  Consul service name of this origin server, including cluster-id.  The format is servicename:cluster-id. If the servicename is \&quot;frontend\&quot;,  and cluster-id is \&quot;prod\&quot;, then you will enter \&quot;frontend:prod\&quot;.  cluster-id is optional.  Example: &#x60; \&quot;matching:production\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true  | [optional] 
**SiteLocator** | Pointer to [**ViewsSiteLocator**](ViewsSiteLocator.md) |  | [optional] 

## Methods

### NewOriginPoolOriginServerConsulService

`func NewOriginPoolOriginServerConsulService() *OriginPoolOriginServerConsulService`

NewOriginPoolOriginServerConsulService instantiates a new OriginPoolOriginServerConsulService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerConsulServiceWithDefaults

`func NewOriginPoolOriginServerConsulServiceWithDefaults() *OriginPoolOriginServerConsulService`

NewOriginPoolOriginServerConsulServiceWithDefaults instantiates a new OriginPoolOriginServerConsulService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsideNetwork

`func (o *OriginPoolOriginServerConsulService) GetInsideNetwork() map[string]interface{}`

GetInsideNetwork returns the InsideNetwork field if non-nil, zero value otherwise.

### GetInsideNetworkOk

`func (o *OriginPoolOriginServerConsulService) GetInsideNetworkOk() (*map[string]interface{}, bool)`

GetInsideNetworkOk returns a tuple with the InsideNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsideNetwork

`func (o *OriginPoolOriginServerConsulService) SetInsideNetwork(v map[string]interface{})`

SetInsideNetwork sets InsideNetwork field to given value.

### HasInsideNetwork

`func (o *OriginPoolOriginServerConsulService) HasInsideNetwork() bool`

HasInsideNetwork returns a boolean if a field has been set.

### GetOutsideNetwork

`func (o *OriginPoolOriginServerConsulService) GetOutsideNetwork() map[string]interface{}`

GetOutsideNetwork returns the OutsideNetwork field if non-nil, zero value otherwise.

### GetOutsideNetworkOk

`func (o *OriginPoolOriginServerConsulService) GetOutsideNetworkOk() (*map[string]interface{}, bool)`

GetOutsideNetworkOk returns a tuple with the OutsideNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideNetwork

`func (o *OriginPoolOriginServerConsulService) SetOutsideNetwork(v map[string]interface{})`

SetOutsideNetwork sets OutsideNetwork field to given value.

### HasOutsideNetwork

`func (o *OriginPoolOriginServerConsulService) HasOutsideNetwork() bool`

HasOutsideNetwork returns a boolean if a field has been set.

### GetServiceName

`func (o *OriginPoolOriginServerConsulService) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *OriginPoolOriginServerConsulService) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *OriginPoolOriginServerConsulService) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *OriginPoolOriginServerConsulService) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSiteLocator

`func (o *OriginPoolOriginServerConsulService) GetSiteLocator() ViewsSiteLocator`

GetSiteLocator returns the SiteLocator field if non-nil, zero value otherwise.

### GetSiteLocatorOk

`func (o *OriginPoolOriginServerConsulService) GetSiteLocatorOk() (*ViewsSiteLocator, bool)`

GetSiteLocatorOk returns a tuple with the SiteLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteLocator

`func (o *OriginPoolOriginServerConsulService) SetSiteLocator(v ViewsSiteLocator)`

SetSiteLocator sets SiteLocator field to given value.

### HasSiteLocator

`func (o *OriginPoolOriginServerConsulService) HasSiteLocator() bool`

HasSiteLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


