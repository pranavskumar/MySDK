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

// checks if the ClusterHttp2ProtocolOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterHttp2ProtocolOptions{}

// ClusterHttp2ProtocolOptions Http2 Protocol options for upstream connections
type ClusterHttp2ProtocolOptions struct {
	//  Enable/disable HTTP2 Protocol for upstream connections
	Enabled *bool `json:"enabled,omitempty"`
}

// NewClusterHttp2ProtocolOptions instantiates a new ClusterHttp2ProtocolOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterHttp2ProtocolOptions() *ClusterHttp2ProtocolOptions {
	this := ClusterHttp2ProtocolOptions{}
	return &this
}

// NewClusterHttp2ProtocolOptionsWithDefaults instantiates a new ClusterHttp2ProtocolOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterHttp2ProtocolOptionsWithDefaults() *ClusterHttp2ProtocolOptions {
	this := ClusterHttp2ProtocolOptions{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ClusterHttp2ProtocolOptions) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHttp2ProtocolOptions) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClusterHttp2ProtocolOptions) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClusterHttp2ProtocolOptions) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ClusterHttp2ProtocolOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterHttp2ProtocolOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableClusterHttp2ProtocolOptions struct {
	value *ClusterHttp2ProtocolOptions
	isSet bool
}

func (v NullableClusterHttp2ProtocolOptions) Get() *ClusterHttp2ProtocolOptions {
	return v.value
}

func (v *NullableClusterHttp2ProtocolOptions) Set(val *ClusterHttp2ProtocolOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterHttp2ProtocolOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterHttp2ProtocolOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterHttp2ProtocolOptions(val *ClusterHttp2ProtocolOptions) *NullableClusterHttp2ProtocolOptions {
	return &NullableClusterHttp2ProtocolOptions{value: val, isSet: true}
}

func (v NullableClusterHttp2ProtocolOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterHttp2ProtocolOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

