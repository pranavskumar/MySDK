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

// checks if the OriginPoolUpstreamTlsValidationContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginPoolUpstreamTlsValidationContext{}

// OriginPoolUpstreamTlsValidationContext Upstream TLS Validation Context
type OriginPoolUpstreamTlsValidationContext struct {
	// Exclusive with []  Inline Root CA certificate for verification of Server's certificate  Validation Rules:   ves.io.schema.rules.string.max_bytes: 131072   ves.io.schema.rules.string.min_bytes: 1   ves.io.schema.rules.string.truststore_url: true 
	TrustedCaUrl *string `json:"trusted_ca_url,omitempty"`
}

// NewOriginPoolUpstreamTlsValidationContext instantiates a new OriginPoolUpstreamTlsValidationContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginPoolUpstreamTlsValidationContext() *OriginPoolUpstreamTlsValidationContext {
	this := OriginPoolUpstreamTlsValidationContext{}
	return &this
}

// NewOriginPoolUpstreamTlsValidationContextWithDefaults instantiates a new OriginPoolUpstreamTlsValidationContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginPoolUpstreamTlsValidationContextWithDefaults() *OriginPoolUpstreamTlsValidationContext {
	this := OriginPoolUpstreamTlsValidationContext{}
	return &this
}

// GetTrustedCaUrl returns the TrustedCaUrl field value if set, zero value otherwise.
func (o *OriginPoolUpstreamTlsValidationContext) GetTrustedCaUrl() string {
	if o == nil || IsNil(o.TrustedCaUrl) {
		var ret string
		return ret
	}
	return *o.TrustedCaUrl
}

// GetTrustedCaUrlOk returns a tuple with the TrustedCaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolUpstreamTlsValidationContext) GetTrustedCaUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TrustedCaUrl) {
		return nil, false
	}
	return o.TrustedCaUrl, true
}

// HasTrustedCaUrl returns a boolean if a field has been set.
func (o *OriginPoolUpstreamTlsValidationContext) HasTrustedCaUrl() bool {
	if o != nil && !IsNil(o.TrustedCaUrl) {
		return true
	}

	return false
}

// SetTrustedCaUrl gets a reference to the given string and assigns it to the TrustedCaUrl field.
func (o *OriginPoolUpstreamTlsValidationContext) SetTrustedCaUrl(v string) {
	o.TrustedCaUrl = &v
}

func (o OriginPoolUpstreamTlsValidationContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginPoolUpstreamTlsValidationContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrustedCaUrl) {
		toSerialize["trusted_ca_url"] = o.TrustedCaUrl
	}
	return toSerialize, nil
}

type NullableOriginPoolUpstreamTlsValidationContext struct {
	value *OriginPoolUpstreamTlsValidationContext
	isSet bool
}

func (v NullableOriginPoolUpstreamTlsValidationContext) Get() *OriginPoolUpstreamTlsValidationContext {
	return v.value
}

func (v *NullableOriginPoolUpstreamTlsValidationContext) Set(val *OriginPoolUpstreamTlsValidationContext) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginPoolUpstreamTlsValidationContext) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginPoolUpstreamTlsValidationContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginPoolUpstreamTlsValidationContext(val *OriginPoolUpstreamTlsValidationContext) *NullableOriginPoolUpstreamTlsValidationContext {
	return &NullableOriginPoolUpstreamTlsValidationContext{value: val, isSet: true}
}

func (v NullableOriginPoolUpstreamTlsValidationContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginPoolUpstreamTlsValidationContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

