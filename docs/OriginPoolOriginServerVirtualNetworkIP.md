# OriginPoolOriginServerVirtualNetworkIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Exclusive with []  IPV4 address  Example: &#x60; \&quot;1.1.1.1\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.string.ipv4: true  | [optional] 
**VirtualNetwork** | Pointer to [**SchemaviewsObjectRefType**](SchemaviewsObjectRefType.md) |  | [optional] 

## Methods

### NewOriginPoolOriginServerVirtualNetworkIP

`func NewOriginPoolOriginServerVirtualNetworkIP() *OriginPoolOriginServerVirtualNetworkIP`

NewOriginPoolOriginServerVirtualNetworkIP instantiates a new OriginPoolOriginServerVirtualNetworkIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerVirtualNetworkIPWithDefaults

`func NewOriginPoolOriginServerVirtualNetworkIPWithDefaults() *OriginPoolOriginServerVirtualNetworkIP`

NewOriginPoolOriginServerVirtualNetworkIPWithDefaults instantiates a new OriginPoolOriginServerVirtualNetworkIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *OriginPoolOriginServerVirtualNetworkIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OriginPoolOriginServerVirtualNetworkIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OriginPoolOriginServerVirtualNetworkIP) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OriginPoolOriginServerVirtualNetworkIP) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *OriginPoolOriginServerVirtualNetworkIP) GetVirtualNetwork() SchemaviewsObjectRefType`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *OriginPoolOriginServerVirtualNetworkIP) GetVirtualNetworkOk() (*SchemaviewsObjectRefType, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *OriginPoolOriginServerVirtualNetworkIP) SetVirtualNetwork(v SchemaviewsObjectRefType)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *OriginPoolOriginServerVirtualNetworkIP) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


