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

// checks if the SchemaClearSecretInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaClearSecretInfoType{}

// SchemaClearSecretInfoType ClearSecretInfoType specifies information about the Secret that is not encrypted.
type SchemaClearSecretInfoType struct {
	//  Name of the Secret Management Access object that contains information about the store to get encrypted bytes  This field needs to be provided only if the url scheme is not string:///  Example: ` \"box-provider\"`
	Provider *string `json:"provider,omitempty"`
	//  URL of the secret. Currently supported URL schemes is string:///.  For string:/// scheme, Secret needs to be encoded Base64 format.  When asked for this secret, caller will get Secret bytes after Base64 decoding.  Example: ` \"string:///U2VjcmV0SW5mb3JtYXRpb24=\"`  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.string.max_bytes: 131072   ves.io.schema.rules.string.uri_ref: true 
	Url *string `json:"url,omitempty"`
}

// NewSchemaClearSecretInfoType instantiates a new SchemaClearSecretInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaClearSecretInfoType() *SchemaClearSecretInfoType {
	this := SchemaClearSecretInfoType{}
	return &this
}

// NewSchemaClearSecretInfoTypeWithDefaults instantiates a new SchemaClearSecretInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaClearSecretInfoTypeWithDefaults() *SchemaClearSecretInfoType {
	this := SchemaClearSecretInfoType{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SchemaClearSecretInfoType) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaClearSecretInfoType) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SchemaClearSecretInfoType) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *SchemaClearSecretInfoType) SetProvider(v string) {
	o.Provider = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SchemaClearSecretInfoType) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaClearSecretInfoType) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SchemaClearSecretInfoType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SchemaClearSecretInfoType) SetUrl(v string) {
	o.Url = &v
}

func (o SchemaClearSecretInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaClearSecretInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableSchemaClearSecretInfoType struct {
	value *SchemaClearSecretInfoType
	isSet bool
}

func (v NullableSchemaClearSecretInfoType) Get() *SchemaClearSecretInfoType {
	return v.value
}

func (v *NullableSchemaClearSecretInfoType) Set(val *SchemaClearSecretInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaClearSecretInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaClearSecretInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaClearSecretInfoType(val *SchemaClearSecretInfoType) *NullableSchemaClearSecretInfoType {
	return &NullableSchemaClearSecretInfoType{value: val, isSet: true}
}

func (v NullableSchemaClearSecretInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaClearSecretInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


