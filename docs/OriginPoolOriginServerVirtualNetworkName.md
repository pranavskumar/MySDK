# OriginPoolOriginServerVirtualNetworkName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | Pointer to **string** |  DNS Name  Example: &#x60; \&quot;value\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true  | [optional] 
**PrivateNetwork** | Pointer to [**SchemaviewsObjectRefType**](SchemaviewsObjectRefType.md) |  | [optional] 

## Methods

### NewOriginPoolOriginServerVirtualNetworkName

`func NewOriginPoolOriginServerVirtualNetworkName() *OriginPoolOriginServerVirtualNetworkName`

NewOriginPoolOriginServerVirtualNetworkName instantiates a new OriginPoolOriginServerVirtualNetworkName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerVirtualNetworkNameWithDefaults

`func NewOriginPoolOriginServerVirtualNetworkNameWithDefaults() *OriginPoolOriginServerVirtualNetworkName`

NewOriginPoolOriginServerVirtualNetworkNameWithDefaults instantiates a new OriginPoolOriginServerVirtualNetworkName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *OriginPoolOriginServerVirtualNetworkName) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *OriginPoolOriginServerVirtualNetworkName) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *OriginPoolOriginServerVirtualNetworkName) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *OriginPoolOriginServerVirtualNetworkName) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetPrivateNetwork

`func (o *OriginPoolOriginServerVirtualNetworkName) GetPrivateNetwork() SchemaviewsObjectRefType`

GetPrivateNetwork returns the PrivateNetwork field if non-nil, zero value otherwise.

### GetPrivateNetworkOk

`func (o *OriginPoolOriginServerVirtualNetworkName) GetPrivateNetworkOk() (*SchemaviewsObjectRefType, bool)`

GetPrivateNetworkOk returns a tuple with the PrivateNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetwork

`func (o *OriginPoolOriginServerVirtualNetworkName) SetPrivateNetwork(v SchemaviewsObjectRefType)`

SetPrivateNetwork sets PrivateNetwork field to given value.

### HasPrivateNetwork

`func (o *OriginPoolOriginServerVirtualNetworkName) HasPrivateNetwork() bool`

HasPrivateNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


