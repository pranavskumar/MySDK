# OriginPoolOriginServerPrivateIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsideNetwork** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**Ip** | Pointer to **string** | Exclusive with []  Private IPV4 address  Example: &#x60; \&quot;8.8.8.8\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.string.ipv4: true  | [optional] 
**OutsideNetwork** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**SiteLocator** | Pointer to [**ViewsSiteLocator**](ViewsSiteLocator.md) |  | [optional] 

## Methods

### NewOriginPoolOriginServerPrivateIP

`func NewOriginPoolOriginServerPrivateIP() *OriginPoolOriginServerPrivateIP`

NewOriginPoolOriginServerPrivateIP instantiates a new OriginPoolOriginServerPrivateIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerPrivateIPWithDefaults

`func NewOriginPoolOriginServerPrivateIPWithDefaults() *OriginPoolOriginServerPrivateIP`

NewOriginPoolOriginServerPrivateIPWithDefaults instantiates a new OriginPoolOriginServerPrivateIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsideNetwork

`func (o *OriginPoolOriginServerPrivateIP) GetInsideNetwork() map[string]interface{}`

GetInsideNetwork returns the InsideNetwork field if non-nil, zero value otherwise.

### GetInsideNetworkOk

`func (o *OriginPoolOriginServerPrivateIP) GetInsideNetworkOk() (*map[string]interface{}, bool)`

GetInsideNetworkOk returns a tuple with the InsideNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsideNetwork

`func (o *OriginPoolOriginServerPrivateIP) SetInsideNetwork(v map[string]interface{})`

SetInsideNetwork sets InsideNetwork field to given value.

### HasInsideNetwork

`func (o *OriginPoolOriginServerPrivateIP) HasInsideNetwork() bool`

HasInsideNetwork returns a boolean if a field has been set.

### GetIp

`func (o *OriginPoolOriginServerPrivateIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OriginPoolOriginServerPrivateIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OriginPoolOriginServerPrivateIP) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OriginPoolOriginServerPrivateIP) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetOutsideNetwork

`func (o *OriginPoolOriginServerPrivateIP) GetOutsideNetwork() map[string]interface{}`

GetOutsideNetwork returns the OutsideNetwork field if non-nil, zero value otherwise.

### GetOutsideNetworkOk

`func (o *OriginPoolOriginServerPrivateIP) GetOutsideNetworkOk() (*map[string]interface{}, bool)`

GetOutsideNetworkOk returns a tuple with the OutsideNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideNetwork

`func (o *OriginPoolOriginServerPrivateIP) SetOutsideNetwork(v map[string]interface{})`

SetOutsideNetwork sets OutsideNetwork field to given value.

### HasOutsideNetwork

`func (o *OriginPoolOriginServerPrivateIP) HasOutsideNetwork() bool`

HasOutsideNetwork returns a boolean if a field has been set.

### GetSiteLocator

`func (o *OriginPoolOriginServerPrivateIP) GetSiteLocator() ViewsSiteLocator`

GetSiteLocator returns the SiteLocator field if non-nil, zero value otherwise.

### GetSiteLocatorOk

`func (o *OriginPoolOriginServerPrivateIP) GetSiteLocatorOk() (*ViewsSiteLocator, bool)`

GetSiteLocatorOk returns a tuple with the SiteLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteLocator

`func (o *OriginPoolOriginServerPrivateIP) SetSiteLocator(v ViewsSiteLocator)`

SetSiteLocator sets SiteLocator field to given value.

### HasSiteLocator

`func (o *OriginPoolOriginServerPrivateIP) HasSiteLocator() bool`

HasSiteLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


