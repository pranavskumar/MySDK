/*
F5 Distributed Cloud Services API for ves.io.schema.views.origin_pool

Origin pool is a view to create cluster and endpoints that can be used in HTTP loadbalancer or TCP loadbalancer  It will create following child objects  * cluster * endpoints * healthcheck

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OriginPoolOriginServerType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginPoolOriginServerType{}

// OriginPoolOriginServerType Various options to specify origin server
type OriginPoolOriginServerType struct {
	ConsulService *OriginPoolOriginServerConsulService `json:"consul_service,omitempty"`
	CustomEndpointObject *OriginPoolOriginServerCustomEndpoint `json:"custom_endpoint_object,omitempty"`
	K8sService *OriginPoolOriginServerK8SService `json:"k8s_service,omitempty"`
	//  Add Labels for this origin server, these labels can be used to form subset.  Example: ` \"value\"`
	Labels map[string]interface{} `json:"labels,omitempty"`
	PrivateIp *OriginPoolOriginServerPrivateIP `json:"private_ip,omitempty"`
	PrivateName *OriginPoolOriginServerPrivateName `json:"private_name,omitempty"`
	PublicIp *OriginPoolOriginServerPublicIP `json:"public_ip,omitempty"`
	PublicName *OriginPoolOriginServerPublicName `json:"public_name,omitempty"`
	VnPrivateIp *OriginPoolOriginServerVirtualNetworkIP `json:"vn_private_ip,omitempty"`
	VnPrivateName *OriginPoolOriginServerVirtualNetworkName `json:"vn_private_name,omitempty"`
}

// NewOriginPoolOriginServerType instantiates a new OriginPoolOriginServerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginPoolOriginServerType() *OriginPoolOriginServerType {
	this := OriginPoolOriginServerType{}
	return &this
}

// NewOriginPoolOriginServerTypeWithDefaults instantiates a new OriginPoolOriginServerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginPoolOriginServerTypeWithDefaults() *OriginPoolOriginServerType {
	this := OriginPoolOriginServerType{}
	return &this
}

// GetConsulService returns the ConsulService field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetConsulService() OriginPoolOriginServerConsulService {
	if o == nil || IsNil(o.ConsulService) {
		var ret OriginPoolOriginServerConsulService
		return ret
	}
	return *o.ConsulService
}

// GetConsulServiceOk returns a tuple with the ConsulService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetConsulServiceOk() (*OriginPoolOriginServerConsulService, bool) {
	if o == nil || IsNil(o.ConsulService) {
		return nil, false
	}
	return o.ConsulService, true
}

// HasConsulService returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasConsulService() bool {
	if o != nil && !IsNil(o.ConsulService) {
		return true
	}

	return false
}

// SetConsulService gets a reference to the given OriginPoolOriginServerConsulService and assigns it to the ConsulService field.
func (o *OriginPoolOriginServerType) SetConsulService(v OriginPoolOriginServerConsulService) {
	o.ConsulService = &v
}

// GetCustomEndpointObject returns the CustomEndpointObject field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetCustomEndpointObject() OriginPoolOriginServerCustomEndpoint {
	if o == nil || IsNil(o.CustomEndpointObject) {
		var ret OriginPoolOriginServerCustomEndpoint
		return ret
	}
	return *o.CustomEndpointObject
}

// GetCustomEndpointObjectOk returns a tuple with the CustomEndpointObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetCustomEndpointObjectOk() (*OriginPoolOriginServerCustomEndpoint, bool) {
	if o == nil || IsNil(o.CustomEndpointObject) {
		return nil, false
	}
	return o.CustomEndpointObject, true
}

// HasCustomEndpointObject returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasCustomEndpointObject() bool {
	if o != nil && !IsNil(o.CustomEndpointObject) {
		return true
	}

	return false
}

// SetCustomEndpointObject gets a reference to the given OriginPoolOriginServerCustomEndpoint and assigns it to the CustomEndpointObject field.
func (o *OriginPoolOriginServerType) SetCustomEndpointObject(v OriginPoolOriginServerCustomEndpoint) {
	o.CustomEndpointObject = &v
}

// GetK8sService returns the K8sService field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetK8sService() OriginPoolOriginServerK8SService {
	if o == nil || IsNil(o.K8sService) {
		var ret OriginPoolOriginServerK8SService
		return ret
	}
	return *o.K8sService
}

// GetK8sServiceOk returns a tuple with the K8sService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetK8sServiceOk() (*OriginPoolOriginServerK8SService, bool) {
	if o == nil || IsNil(o.K8sService) {
		return nil, false
	}
	return o.K8sService, true
}

// HasK8sService returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasK8sService() bool {
	if o != nil && !IsNil(o.K8sService) {
		return true
	}

	return false
}

// SetK8sService gets a reference to the given OriginPoolOriginServerK8SService and assigns it to the K8sService field.
func (o *OriginPoolOriginServerType) SetK8sService(v OriginPoolOriginServerK8SService) {
	o.K8sService = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetLabels() map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetLabelsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *OriginPoolOriginServerType) SetLabels(v map[string]interface{}) {
	o.Labels = v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetPrivateIp() OriginPoolOriginServerPrivateIP {
	if o == nil || IsNil(o.PrivateIp) {
		var ret OriginPoolOriginServerPrivateIP
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetPrivateIpOk() (*OriginPoolOriginServerPrivateIP, bool) {
	if o == nil || IsNil(o.PrivateIp) {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasPrivateIp() bool {
	if o != nil && !IsNil(o.PrivateIp) {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given OriginPoolOriginServerPrivateIP and assigns it to the PrivateIp field.
func (o *OriginPoolOriginServerType) SetPrivateIp(v OriginPoolOriginServerPrivateIP) {
	o.PrivateIp = &v
}

// GetPrivateName returns the PrivateName field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetPrivateName() OriginPoolOriginServerPrivateName {
	if o == nil || IsNil(o.PrivateName) {
		var ret OriginPoolOriginServerPrivateName
		return ret
	}
	return *o.PrivateName
}

// GetPrivateNameOk returns a tuple with the PrivateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetPrivateNameOk() (*OriginPoolOriginServerPrivateName, bool) {
	if o == nil || IsNil(o.PrivateName) {
		return nil, false
	}
	return o.PrivateName, true
}

// HasPrivateName returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasPrivateName() bool {
	if o != nil && !IsNil(o.PrivateName) {
		return true
	}

	return false
}

// SetPrivateName gets a reference to the given OriginPoolOriginServerPrivateName and assigns it to the PrivateName field.
func (o *OriginPoolOriginServerType) SetPrivateName(v OriginPoolOriginServerPrivateName) {
	o.PrivateName = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetPublicIp() OriginPoolOriginServerPublicIP {
	if o == nil || IsNil(o.PublicIp) {
		var ret OriginPoolOriginServerPublicIP
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetPublicIpOk() (*OriginPoolOriginServerPublicIP, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given OriginPoolOriginServerPublicIP and assigns it to the PublicIp field.
func (o *OriginPoolOriginServerType) SetPublicIp(v OriginPoolOriginServerPublicIP) {
	o.PublicIp = &v
}

// GetPublicName returns the PublicName field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetPublicName() OriginPoolOriginServerPublicName {
	if o == nil || IsNil(o.PublicName) {
		var ret OriginPoolOriginServerPublicName
		return ret
	}
	return *o.PublicName
}

// GetPublicNameOk returns a tuple with the PublicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetPublicNameOk() (*OriginPoolOriginServerPublicName, bool) {
	if o == nil || IsNil(o.PublicName) {
		return nil, false
	}
	return o.PublicName, true
}

// HasPublicName returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasPublicName() bool {
	if o != nil && !IsNil(o.PublicName) {
		return true
	}

	return false
}

// SetPublicName gets a reference to the given OriginPoolOriginServerPublicName and assigns it to the PublicName field.
func (o *OriginPoolOriginServerType) SetPublicName(v OriginPoolOriginServerPublicName) {
	o.PublicName = &v
}

// GetVnPrivateIp returns the VnPrivateIp field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetVnPrivateIp() OriginPoolOriginServerVirtualNetworkIP {
	if o == nil || IsNil(o.VnPrivateIp) {
		var ret OriginPoolOriginServerVirtualNetworkIP
		return ret
	}
	return *o.VnPrivateIp
}

// GetVnPrivateIpOk returns a tuple with the VnPrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetVnPrivateIpOk() (*OriginPoolOriginServerVirtualNetworkIP, bool) {
	if o == nil || IsNil(o.VnPrivateIp) {
		return nil, false
	}
	return o.VnPrivateIp, true
}

// HasVnPrivateIp returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasVnPrivateIp() bool {
	if o != nil && !IsNil(o.VnPrivateIp) {
		return true
	}

	return false
}

// SetVnPrivateIp gets a reference to the given OriginPoolOriginServerVirtualNetworkIP and assigns it to the VnPrivateIp field.
func (o *OriginPoolOriginServerType) SetVnPrivateIp(v OriginPoolOriginServerVirtualNetworkIP) {
	o.VnPrivateIp = &v
}

// GetVnPrivateName returns the VnPrivateName field value if set, zero value otherwise.
func (o *OriginPoolOriginServerType) GetVnPrivateName() OriginPoolOriginServerVirtualNetworkName {
	if o == nil || IsNil(o.VnPrivateName) {
		var ret OriginPoolOriginServerVirtualNetworkName
		return ret
	}
	return *o.VnPrivateName
}

// GetVnPrivateNameOk returns a tuple with the VnPrivateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerType) GetVnPrivateNameOk() (*OriginPoolOriginServerVirtualNetworkName, bool) {
	if o == nil || IsNil(o.VnPrivateName) {
		return nil, false
	}
	return o.VnPrivateName, true
}

// HasVnPrivateName returns a boolean if a field has been set.
func (o *OriginPoolOriginServerType) HasVnPrivateName() bool {
	if o != nil && !IsNil(o.VnPrivateName) {
		return true
	}

	return false
}

// SetVnPrivateName gets a reference to the given OriginPoolOriginServerVirtualNetworkName and assigns it to the VnPrivateName field.
func (o *OriginPoolOriginServerType) SetVnPrivateName(v OriginPoolOriginServerVirtualNetworkName) {
	o.VnPrivateName = &v
}

func (o OriginPoolOriginServerType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginPoolOriginServerType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsulService) {
		toSerialize["consul_service"] = o.ConsulService
	}
	if !IsNil(o.CustomEndpointObject) {
		toSerialize["custom_endpoint_object"] = o.CustomEndpointObject
	}
	if !IsNil(o.K8sService) {
		toSerialize["k8s_service"] = o.K8sService
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.PrivateIp) {
		toSerialize["private_ip"] = o.PrivateIp
	}
	if !IsNil(o.PrivateName) {
		toSerialize["private_name"] = o.PrivateName
	}
	if !IsNil(o.PublicIp) {
		toSerialize["public_ip"] = o.PublicIp
	}
	if !IsNil(o.PublicName) {
		toSerialize["public_name"] = o.PublicName
	}
	if !IsNil(o.VnPrivateIp) {
		toSerialize["vn_private_ip"] = o.VnPrivateIp
	}
	if !IsNil(o.VnPrivateName) {
		toSerialize["vn_private_name"] = o.VnPrivateName
	}
	return toSerialize, nil
}

type NullableOriginPoolOriginServerType struct {
	value *OriginPoolOriginServerType
	isSet bool
}

func (v NullableOriginPoolOriginServerType) Get() *OriginPoolOriginServerType {
	return v.value
}

func (v *NullableOriginPoolOriginServerType) Set(val *OriginPoolOriginServerType) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginPoolOriginServerType) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginPoolOriginServerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginPoolOriginServerType(val *OriginPoolOriginServerType) *NullableOriginPoolOriginServerType {
	return &NullableOriginPoolOriginServerType{value: val, isSet: true}
}

func (v NullableOriginPoolOriginServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginPoolOriginServerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

