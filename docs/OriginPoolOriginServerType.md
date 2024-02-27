# OriginPoolOriginServerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsulService** | Pointer to [**OriginPoolOriginServerConsulService**](OriginPoolOriginServerConsulService.md) |  | [optional] 
**CustomEndpointObject** | Pointer to [**OriginPoolOriginServerCustomEndpoint**](OriginPoolOriginServerCustomEndpoint.md) |  | [optional] 
**K8sService** | Pointer to [**OriginPoolOriginServerK8SService**](OriginPoolOriginServerK8SService.md) |  | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  Add Labels for this origin server, these labels can be used to form subset.  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**PrivateIp** | Pointer to [**OriginPoolOriginServerPrivateIP**](OriginPoolOriginServerPrivateIP.md) |  | [optional] 
**PrivateName** | Pointer to [**OriginPoolOriginServerPrivateName**](OriginPoolOriginServerPrivateName.md) |  | [optional] 
**PublicIp** | Pointer to [**OriginPoolOriginServerPublicIP**](OriginPoolOriginServerPublicIP.md) |  | [optional] 
**PublicName** | Pointer to [**OriginPoolOriginServerPublicName**](OriginPoolOriginServerPublicName.md) |  | [optional] 
**VnPrivateIp** | Pointer to [**OriginPoolOriginServerVirtualNetworkIP**](OriginPoolOriginServerVirtualNetworkIP.md) |  | [optional] 
**VnPrivateName** | Pointer to [**OriginPoolOriginServerVirtualNetworkName**](OriginPoolOriginServerVirtualNetworkName.md) |  | [optional] 

## Methods

### NewOriginPoolOriginServerType

`func NewOriginPoolOriginServerType() *OriginPoolOriginServerType`

NewOriginPoolOriginServerType instantiates a new OriginPoolOriginServerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerTypeWithDefaults

`func NewOriginPoolOriginServerTypeWithDefaults() *OriginPoolOriginServerType`

NewOriginPoolOriginServerTypeWithDefaults instantiates a new OriginPoolOriginServerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsulService

`func (o *OriginPoolOriginServerType) GetConsulService() OriginPoolOriginServerConsulService`

GetConsulService returns the ConsulService field if non-nil, zero value otherwise.

### GetConsulServiceOk

`func (o *OriginPoolOriginServerType) GetConsulServiceOk() (*OriginPoolOriginServerConsulService, bool)`

GetConsulServiceOk returns a tuple with the ConsulService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsulService

`func (o *OriginPoolOriginServerType) SetConsulService(v OriginPoolOriginServerConsulService)`

SetConsulService sets ConsulService field to given value.

### HasConsulService

`func (o *OriginPoolOriginServerType) HasConsulService() bool`

HasConsulService returns a boolean if a field has been set.

### GetCustomEndpointObject

`func (o *OriginPoolOriginServerType) GetCustomEndpointObject() OriginPoolOriginServerCustomEndpoint`

GetCustomEndpointObject returns the CustomEndpointObject field if non-nil, zero value otherwise.

### GetCustomEndpointObjectOk

`func (o *OriginPoolOriginServerType) GetCustomEndpointObjectOk() (*OriginPoolOriginServerCustomEndpoint, bool)`

GetCustomEndpointObjectOk returns a tuple with the CustomEndpointObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEndpointObject

`func (o *OriginPoolOriginServerType) SetCustomEndpointObject(v OriginPoolOriginServerCustomEndpoint)`

SetCustomEndpointObject sets CustomEndpointObject field to given value.

### HasCustomEndpointObject

`func (o *OriginPoolOriginServerType) HasCustomEndpointObject() bool`

HasCustomEndpointObject returns a boolean if a field has been set.

### GetK8sService

`func (o *OriginPoolOriginServerType) GetK8sService() OriginPoolOriginServerK8SService`

GetK8sService returns the K8sService field if non-nil, zero value otherwise.

### GetK8sServiceOk

`func (o *OriginPoolOriginServerType) GetK8sServiceOk() (*OriginPoolOriginServerK8SService, bool)`

GetK8sServiceOk returns a tuple with the K8sService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sService

`func (o *OriginPoolOriginServerType) SetK8sService(v OriginPoolOriginServerK8SService)`

SetK8sService sets K8sService field to given value.

### HasK8sService

`func (o *OriginPoolOriginServerType) HasK8sService() bool`

HasK8sService returns a boolean if a field has been set.

### GetLabels

`func (o *OriginPoolOriginServerType) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OriginPoolOriginServerType) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OriginPoolOriginServerType) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OriginPoolOriginServerType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPrivateIp

`func (o *OriginPoolOriginServerType) GetPrivateIp() OriginPoolOriginServerPrivateIP`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *OriginPoolOriginServerType) GetPrivateIpOk() (*OriginPoolOriginServerPrivateIP, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *OriginPoolOriginServerType) SetPrivateIp(v OriginPoolOriginServerPrivateIP)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *OriginPoolOriginServerType) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPrivateName

`func (o *OriginPoolOriginServerType) GetPrivateName() OriginPoolOriginServerPrivateName`

GetPrivateName returns the PrivateName field if non-nil, zero value otherwise.

### GetPrivateNameOk

`func (o *OriginPoolOriginServerType) GetPrivateNameOk() (*OriginPoolOriginServerPrivateName, bool)`

GetPrivateNameOk returns a tuple with the PrivateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateName

`func (o *OriginPoolOriginServerType) SetPrivateName(v OriginPoolOriginServerPrivateName)`

SetPrivateName sets PrivateName field to given value.

### HasPrivateName

`func (o *OriginPoolOriginServerType) HasPrivateName() bool`

HasPrivateName returns a boolean if a field has been set.

### GetPublicIp

`func (o *OriginPoolOriginServerType) GetPublicIp() OriginPoolOriginServerPublicIP`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *OriginPoolOriginServerType) GetPublicIpOk() (*OriginPoolOriginServerPublicIP, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *OriginPoolOriginServerType) SetPublicIp(v OriginPoolOriginServerPublicIP)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *OriginPoolOriginServerType) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetPublicName

`func (o *OriginPoolOriginServerType) GetPublicName() OriginPoolOriginServerPublicName`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *OriginPoolOriginServerType) GetPublicNameOk() (*OriginPoolOriginServerPublicName, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *OriginPoolOriginServerType) SetPublicName(v OriginPoolOriginServerPublicName)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *OriginPoolOriginServerType) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### GetVnPrivateIp

`func (o *OriginPoolOriginServerType) GetVnPrivateIp() OriginPoolOriginServerVirtualNetworkIP`

GetVnPrivateIp returns the VnPrivateIp field if non-nil, zero value otherwise.

### GetVnPrivateIpOk

`func (o *OriginPoolOriginServerType) GetVnPrivateIpOk() (*OriginPoolOriginServerVirtualNetworkIP, bool)`

GetVnPrivateIpOk returns a tuple with the VnPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnPrivateIp

`func (o *OriginPoolOriginServerType) SetVnPrivateIp(v OriginPoolOriginServerVirtualNetworkIP)`

SetVnPrivateIp sets VnPrivateIp field to given value.

### HasVnPrivateIp

`func (o *OriginPoolOriginServerType) HasVnPrivateIp() bool`

HasVnPrivateIp returns a boolean if a field has been set.

### GetVnPrivateName

`func (o *OriginPoolOriginServerType) GetVnPrivateName() OriginPoolOriginServerVirtualNetworkName`

GetVnPrivateName returns the VnPrivateName field if non-nil, zero value otherwise.

### GetVnPrivateNameOk

`func (o *OriginPoolOriginServerType) GetVnPrivateNameOk() (*OriginPoolOriginServerVirtualNetworkName, bool)`

GetVnPrivateNameOk returns a tuple with the VnPrivateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnPrivateName

`func (o *OriginPoolOriginServerType) SetVnPrivateName(v OriginPoolOriginServerVirtualNetworkName)`

SetVnPrivateName sets VnPrivateName field to given value.

### HasVnPrivateName

`func (o *OriginPoolOriginServerType) HasVnPrivateName() bool`

HasVnPrivateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


