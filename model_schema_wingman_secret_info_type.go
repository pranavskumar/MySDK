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

// checks if the SchemaWingmanSecretInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaWingmanSecretInfoType{}

// SchemaWingmanSecretInfoType x-displayName: \"Wingman Secret\" WingmanSecretInfoType specifies the handle to the wingman secret
type SchemaWingmanSecretInfoType struct {
	// x-displayName: \"Name\" x-required x-example: \"ChargeBack-API-Key\" Name of the secret.
	Name *string `json:"name,omitempty"`
}

// NewSchemaWingmanSecretInfoType instantiates a new SchemaWingmanSecretInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaWingmanSecretInfoType() *SchemaWingmanSecretInfoType {
	this := SchemaWingmanSecretInfoType{}
	return &this
}

// NewSchemaWingmanSecretInfoTypeWithDefaults instantiates a new SchemaWingmanSecretInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaWingmanSecretInfoTypeWithDefaults() *SchemaWingmanSecretInfoType {
	this := SchemaWingmanSecretInfoType{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemaWingmanSecretInfoType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaWingmanSecretInfoType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemaWingmanSecretInfoType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemaWingmanSecretInfoType) SetName(v string) {
	o.Name = &v
}

func (o SchemaWingmanSecretInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaWingmanSecretInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSchemaWingmanSecretInfoType struct {
	value *SchemaWingmanSecretInfoType
	isSet bool
}

func (v NullableSchemaWingmanSecretInfoType) Get() *SchemaWingmanSecretInfoType {
	return v.value
}

func (v *NullableSchemaWingmanSecretInfoType) Set(val *SchemaWingmanSecretInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaWingmanSecretInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaWingmanSecretInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaWingmanSecretInfoType(val *SchemaWingmanSecretInfoType) *NullableSchemaWingmanSecretInfoType {
	return &NullableSchemaWingmanSecretInfoType{value: val, isSet: true}
}

func (v NullableSchemaWingmanSecretInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaWingmanSecretInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


