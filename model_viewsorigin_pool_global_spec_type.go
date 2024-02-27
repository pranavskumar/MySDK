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

// checks if the ViewsoriginPoolGlobalSpecType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewsoriginPoolGlobalSpecType{}

// ViewsoriginPoolGlobalSpecType Shape of the origin pool specification
type ViewsoriginPoolGlobalSpecType struct {
	AdvancedOptions *OriginPoolOriginPoolAdvancedOptions `json:"advanced_options,omitempty"`
	// This can be used for messages where no values are needed
	AutomaticPort map[string]interface{} `json:"automatic_port,omitempty"`
	EndpointSelection *ClusterEndpointSelectionPolicy `json:"endpoint_selection,omitempty"`
	// Exclusive with [same_as_endpoint_port]  Port used for performing health check  Validation Rules:   ves.io.schema.rules.uint32.lte: 65535 
	HealthCheckPort *int64 `json:"health_check_port,omitempty"`
	//  Reference to healthcheck configuration objects  Validation Rules:   ves.io.schema.rules.repeated.max_items: 4 
	Healthcheck []SchemaviewsObjectRefType `json:"healthcheck,omitempty"`
	// This can be used for messages where no values are needed
	LbPort map[string]interface{} `json:"lb_port,omitempty"`
	LoadbalancerAlgorithm *ClusterLoadbalancerAlgorithm `json:"loadbalancer_algorithm,omitempty"`
	// This can be used for messages where no values are needed
	NoTls map[string]interface{} `json:"no_tls,omitempty"`
	//  List of origin servers in this pool  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.repeated.max_items: 32   ves.io.schema.rules.repeated.min_items: 1   ves.io.schema.rules.repeated.unique: true 
	OriginServers []OriginPoolOriginServerType `json:"origin_servers,omitempty"`
	// Exclusive with [automatic_port lb_port]  Endpoint service is available on this port  Example: ` \"9080\"`  Validation Rules:   ves.io.schema.rules.uint32.gte: 1   ves.io.schema.rules.uint32.lte: 65535 
	Port *int64 `json:"port,omitempty"`
	// This can be used for messages where no values are needed
	SameAsEndpointPort map[string]interface{} `json:"same_as_endpoint_port,omitempty"`
	UseTls *OriginPoolUpstreamTlsParameters `json:"use_tls,omitempty"`
	ViewInternal *SchemaviewsObjectRefType `json:"view_internal,omitempty"`
}

// NewViewsoriginPoolGlobalSpecType instantiates a new ViewsoriginPoolGlobalSpecType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewsoriginPoolGlobalSpecType() *ViewsoriginPoolGlobalSpecType {
	this := ViewsoriginPoolGlobalSpecType{}
	var endpointSelection ClusterEndpointSelectionPolicy = DISTRIBUTED
	this.EndpointSelection = &endpointSelection
	var loadbalancerAlgorithm ClusterLoadbalancerAlgorithm = ROUND_ROBIN
	this.LoadbalancerAlgorithm = &loadbalancerAlgorithm
	return &this
}

// NewViewsoriginPoolGlobalSpecTypeWithDefaults instantiates a new ViewsoriginPoolGlobalSpecType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewsoriginPoolGlobalSpecTypeWithDefaults() *ViewsoriginPoolGlobalSpecType {
	this := ViewsoriginPoolGlobalSpecType{}
	var endpointSelection ClusterEndpointSelectionPolicy = DISTRIBUTED
	this.EndpointSelection = &endpointSelection
	var loadbalancerAlgorithm ClusterLoadbalancerAlgorithm = ROUND_ROBIN
	this.LoadbalancerAlgorithm = &loadbalancerAlgorithm
	return &this
}

// GetAdvancedOptions returns the AdvancedOptions field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetAdvancedOptions() OriginPoolOriginPoolAdvancedOptions {
	if o == nil || IsNil(o.AdvancedOptions) {
		var ret OriginPoolOriginPoolAdvancedOptions
		return ret
	}
	return *o.AdvancedOptions
}

// GetAdvancedOptionsOk returns a tuple with the AdvancedOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetAdvancedOptionsOk() (*OriginPoolOriginPoolAdvancedOptions, bool) {
	if o == nil || IsNil(o.AdvancedOptions) {
		return nil, false
	}
	return o.AdvancedOptions, true
}

// HasAdvancedOptions returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasAdvancedOptions() bool {
	if o != nil && !IsNil(o.AdvancedOptions) {
		return true
	}

	return false
}

// SetAdvancedOptions gets a reference to the given OriginPoolOriginPoolAdvancedOptions and assigns it to the AdvancedOptions field.
func (o *ViewsoriginPoolGlobalSpecType) SetAdvancedOptions(v OriginPoolOriginPoolAdvancedOptions) {
	o.AdvancedOptions = &v
}

// GetAutomaticPort returns the AutomaticPort field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetAutomaticPort() map[string]interface{} {
	if o == nil || IsNil(o.AutomaticPort) {
		var ret map[string]interface{}
		return ret
	}
	return o.AutomaticPort
}

// GetAutomaticPortOk returns a tuple with the AutomaticPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetAutomaticPortOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AutomaticPort) {
		return map[string]interface{}{}, false
	}
	return o.AutomaticPort, true
}

// HasAutomaticPort returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasAutomaticPort() bool {
	if o != nil && !IsNil(o.AutomaticPort) {
		return true
	}

	return false
}

// SetAutomaticPort gets a reference to the given map[string]interface{} and assigns it to the AutomaticPort field.
func (o *ViewsoriginPoolGlobalSpecType) SetAutomaticPort(v map[string]interface{}) {
	o.AutomaticPort = v
}

// GetEndpointSelection returns the EndpointSelection field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetEndpointSelection() ClusterEndpointSelectionPolicy {
	if o == nil || IsNil(o.EndpointSelection) {
		var ret ClusterEndpointSelectionPolicy
		return ret
	}
	return *o.EndpointSelection
}

// GetEndpointSelectionOk returns a tuple with the EndpointSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetEndpointSelectionOk() (*ClusterEndpointSelectionPolicy, bool) {
	if o == nil || IsNil(o.EndpointSelection) {
		return nil, false
	}
	return o.EndpointSelection, true
}

// HasEndpointSelection returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasEndpointSelection() bool {
	if o != nil && !IsNil(o.EndpointSelection) {
		return true
	}

	return false
}

// SetEndpointSelection gets a reference to the given ClusterEndpointSelectionPolicy and assigns it to the EndpointSelection field.
func (o *ViewsoriginPoolGlobalSpecType) SetEndpointSelection(v ClusterEndpointSelectionPolicy) {
	o.EndpointSelection = &v
}

// GetHealthCheckPort returns the HealthCheckPort field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetHealthCheckPort() int64 {
	if o == nil || IsNil(o.HealthCheckPort) {
		var ret int64
		return ret
	}
	return *o.HealthCheckPort
}

// GetHealthCheckPortOk returns a tuple with the HealthCheckPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetHealthCheckPortOk() (*int64, bool) {
	if o == nil || IsNil(o.HealthCheckPort) {
		return nil, false
	}
	return o.HealthCheckPort, true
}

// HasHealthCheckPort returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasHealthCheckPort() bool {
	if o != nil && !IsNil(o.HealthCheckPort) {
		return true
	}

	return false
}

// SetHealthCheckPort gets a reference to the given int64 and assigns it to the HealthCheckPort field.
func (o *ViewsoriginPoolGlobalSpecType) SetHealthCheckPort(v int64) {
	o.HealthCheckPort = &v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetHealthcheck() []SchemaviewsObjectRefType {
	if o == nil || IsNil(o.Healthcheck) {
		var ret []SchemaviewsObjectRefType
		return ret
	}
	return o.Healthcheck
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetHealthcheckOk() ([]SchemaviewsObjectRefType, bool) {
	if o == nil || IsNil(o.Healthcheck) {
		return nil, false
	}
	return o.Healthcheck, true
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasHealthcheck() bool {
	if o != nil && !IsNil(o.Healthcheck) {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given []SchemaviewsObjectRefType and assigns it to the Healthcheck field.
func (o *ViewsoriginPoolGlobalSpecType) SetHealthcheck(v []SchemaviewsObjectRefType) {
	o.Healthcheck = v
}

// GetLbPort returns the LbPort field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetLbPort() map[string]interface{} {
	if o == nil || IsNil(o.LbPort) {
		var ret map[string]interface{}
		return ret
	}
	return o.LbPort
}

// GetLbPortOk returns a tuple with the LbPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetLbPortOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LbPort) {
		return map[string]interface{}{}, false
	}
	return o.LbPort, true
}

// HasLbPort returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasLbPort() bool {
	if o != nil && !IsNil(o.LbPort) {
		return true
	}

	return false
}

// SetLbPort gets a reference to the given map[string]interface{} and assigns it to the LbPort field.
func (o *ViewsoriginPoolGlobalSpecType) SetLbPort(v map[string]interface{}) {
	o.LbPort = v
}

// GetLoadbalancerAlgorithm returns the LoadbalancerAlgorithm field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetLoadbalancerAlgorithm() ClusterLoadbalancerAlgorithm {
	if o == nil || IsNil(o.LoadbalancerAlgorithm) {
		var ret ClusterLoadbalancerAlgorithm
		return ret
	}
	return *o.LoadbalancerAlgorithm
}

// GetLoadbalancerAlgorithmOk returns a tuple with the LoadbalancerAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetLoadbalancerAlgorithmOk() (*ClusterLoadbalancerAlgorithm, bool) {
	if o == nil || IsNil(o.LoadbalancerAlgorithm) {
		return nil, false
	}
	return o.LoadbalancerAlgorithm, true
}

// HasLoadbalancerAlgorithm returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasLoadbalancerAlgorithm() bool {
	if o != nil && !IsNil(o.LoadbalancerAlgorithm) {
		return true
	}

	return false
}

// SetLoadbalancerAlgorithm gets a reference to the given ClusterLoadbalancerAlgorithm and assigns it to the LoadbalancerAlgorithm field.
func (o *ViewsoriginPoolGlobalSpecType) SetLoadbalancerAlgorithm(v ClusterLoadbalancerAlgorithm) {
	o.LoadbalancerAlgorithm = &v
}

// GetNoTls returns the NoTls field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetNoTls() map[string]interface{} {
	if o == nil || IsNil(o.NoTls) {
		var ret map[string]interface{}
		return ret
	}
	return o.NoTls
}

// GetNoTlsOk returns a tuple with the NoTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetNoTlsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NoTls) {
		return map[string]interface{}{}, false
	}
	return o.NoTls, true
}

// HasNoTls returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasNoTls() bool {
	if o != nil && !IsNil(o.NoTls) {
		return true
	}

	return false
}

// SetNoTls gets a reference to the given map[string]interface{} and assigns it to the NoTls field.
func (o *ViewsoriginPoolGlobalSpecType) SetNoTls(v map[string]interface{}) {
	o.NoTls = v
}

// GetOriginServers returns the OriginServers field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetOriginServers() []OriginPoolOriginServerType {
	if o == nil || IsNil(o.OriginServers) {
		var ret []OriginPoolOriginServerType
		return ret
	}
	return o.OriginServers
}

// GetOriginServersOk returns a tuple with the OriginServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetOriginServersOk() ([]OriginPoolOriginServerType, bool) {
	if o == nil || IsNil(o.OriginServers) {
		return nil, false
	}
	return o.OriginServers, true
}

// HasOriginServers returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasOriginServers() bool {
	if o != nil && !IsNil(o.OriginServers) {
		return true
	}

	return false
}

// SetOriginServers gets a reference to the given []OriginPoolOriginServerType and assigns it to the OriginServers field.
func (o *ViewsoriginPoolGlobalSpecType) SetOriginServers(v []OriginPoolOriginServerType) {
	o.OriginServers = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ViewsoriginPoolGlobalSpecType) SetPort(v int64) {
	o.Port = &v
}

// GetSameAsEndpointPort returns the SameAsEndpointPort field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetSameAsEndpointPort() map[string]interface{} {
	if o == nil || IsNil(o.SameAsEndpointPort) {
		var ret map[string]interface{}
		return ret
	}
	return o.SameAsEndpointPort
}

// GetSameAsEndpointPortOk returns a tuple with the SameAsEndpointPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetSameAsEndpointPortOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SameAsEndpointPort) {
		return map[string]interface{}{}, false
	}
	return o.SameAsEndpointPort, true
}

// HasSameAsEndpointPort returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasSameAsEndpointPort() bool {
	if o != nil && !IsNil(o.SameAsEndpointPort) {
		return true
	}

	return false
}

// SetSameAsEndpointPort gets a reference to the given map[string]interface{} and assigns it to the SameAsEndpointPort field.
func (o *ViewsoriginPoolGlobalSpecType) SetSameAsEndpointPort(v map[string]interface{}) {
	o.SameAsEndpointPort = v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetUseTls() OriginPoolUpstreamTlsParameters {
	if o == nil || IsNil(o.UseTls) {
		var ret OriginPoolUpstreamTlsParameters
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetUseTlsOk() (*OriginPoolUpstreamTlsParameters, bool) {
	if o == nil || IsNil(o.UseTls) {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasUseTls() bool {
	if o != nil && !IsNil(o.UseTls) {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given OriginPoolUpstreamTlsParameters and assigns it to the UseTls field.
func (o *ViewsoriginPoolGlobalSpecType) SetUseTls(v OriginPoolUpstreamTlsParameters) {
	o.UseTls = &v
}

// GetViewInternal returns the ViewInternal field value if set, zero value otherwise.
func (o *ViewsoriginPoolGlobalSpecType) GetViewInternal() SchemaviewsObjectRefType {
	if o == nil || IsNil(o.ViewInternal) {
		var ret SchemaviewsObjectRefType
		return ret
	}
	return *o.ViewInternal
}

// GetViewInternalOk returns a tuple with the ViewInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsoriginPoolGlobalSpecType) GetViewInternalOk() (*SchemaviewsObjectRefType, bool) {
	if o == nil || IsNil(o.ViewInternal) {
		return nil, false
	}
	return o.ViewInternal, true
}

// HasViewInternal returns a boolean if a field has been set.
func (o *ViewsoriginPoolGlobalSpecType) HasViewInternal() bool {
	if o != nil && !IsNil(o.ViewInternal) {
		return true
	}

	return false
}

// SetViewInternal gets a reference to the given SchemaviewsObjectRefType and assigns it to the ViewInternal field.
func (o *ViewsoriginPoolGlobalSpecType) SetViewInternal(v SchemaviewsObjectRefType) {
	o.ViewInternal = &v
}

func (o ViewsoriginPoolGlobalSpecType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewsoriginPoolGlobalSpecType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvancedOptions) {
		toSerialize["advanced_options"] = o.AdvancedOptions
	}
	if !IsNil(o.AutomaticPort) {
		toSerialize["automatic_port"] = o.AutomaticPort
	}
	if !IsNil(o.EndpointSelection) {
		toSerialize["endpoint_selection"] = o.EndpointSelection
	}
	if !IsNil(o.HealthCheckPort) {
		toSerialize["health_check_port"] = o.HealthCheckPort
	}
	if !IsNil(o.Healthcheck) {
		toSerialize["healthcheck"] = o.Healthcheck
	}
	if !IsNil(o.LbPort) {
		toSerialize["lb_port"] = o.LbPort
	}
	if !IsNil(o.LoadbalancerAlgorithm) {
		toSerialize["loadbalancer_algorithm"] = o.LoadbalancerAlgorithm
	}
	if !IsNil(o.NoTls) {
		toSerialize["no_tls"] = o.NoTls
	}
	if !IsNil(o.OriginServers) {
		toSerialize["origin_servers"] = o.OriginServers
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.SameAsEndpointPort) {
		toSerialize["same_as_endpoint_port"] = o.SameAsEndpointPort
	}
	if !IsNil(o.UseTls) {
		toSerialize["use_tls"] = o.UseTls
	}
	if !IsNil(o.ViewInternal) {
		toSerialize["view_internal"] = o.ViewInternal
	}
	return toSerialize, nil
}

type NullableViewsoriginPoolGlobalSpecType struct {
	value *ViewsoriginPoolGlobalSpecType
	isSet bool
}

func (v NullableViewsoriginPoolGlobalSpecType) Get() *ViewsoriginPoolGlobalSpecType {
	return v.value
}

func (v *NullableViewsoriginPoolGlobalSpecType) Set(val *ViewsoriginPoolGlobalSpecType) {
	v.value = val
	v.isSet = true
}

func (v NullableViewsoriginPoolGlobalSpecType) IsSet() bool {
	return v.isSet
}

func (v *NullableViewsoriginPoolGlobalSpecType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewsoriginPoolGlobalSpecType(val *ViewsoriginPoolGlobalSpecType) *NullableViewsoriginPoolGlobalSpecType {
	return &NullableViewsoriginPoolGlobalSpecType{value: val, isSet: true}
}

func (v NullableViewsoriginPoolGlobalSpecType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewsoriginPoolGlobalSpecType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

