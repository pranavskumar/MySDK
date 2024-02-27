/*
F5 Distributed Cloud Services API for ves.io.schema.views.origin_pool

Origin pool is a view to create cluster and endpoints that can be used in HTTP loadbalancer or TCP loadbalancer  It will create following child objects  * cluster * endpoints * healthcheck

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SchemaHashAlgorithm Specifies the Hash Algorithm to be used  Invalid hash algorithm sha256 hash algorithm sha1 hash algorithm
type SchemaHashAlgorithm string

// List of schemaHashAlgorithm
const (
	INVALID_HASH_ALGORITHM SchemaHashAlgorithm = "INVALID_HASH_ALGORITHM"
	SHA256 SchemaHashAlgorithm = "SHA256"
	SHA1 SchemaHashAlgorithm = "SHA1"
)

// All allowed values of SchemaHashAlgorithm enum
var AllowedSchemaHashAlgorithmEnumValues = []SchemaHashAlgorithm{
	"INVALID_HASH_ALGORITHM",
	"SHA256",
	"SHA1",
}

func (v *SchemaHashAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SchemaHashAlgorithm(value)
	for _, existing := range AllowedSchemaHashAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SchemaHashAlgorithm", value)
}

// NewSchemaHashAlgorithmFromValue returns a pointer to a valid SchemaHashAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSchemaHashAlgorithmFromValue(v string) (*SchemaHashAlgorithm, error) {
	ev := SchemaHashAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SchemaHashAlgorithm: valid values are %v", v, AllowedSchemaHashAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SchemaHashAlgorithm) IsValid() bool {
	for _, existing := range AllowedSchemaHashAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to schemaHashAlgorithm value
func (v SchemaHashAlgorithm) Ptr() *SchemaHashAlgorithm {
	return &v
}

type NullableSchemaHashAlgorithm struct {
	value *SchemaHashAlgorithm
	isSet bool
}

func (v NullableSchemaHashAlgorithm) Get() *SchemaHashAlgorithm {
	return v.value
}

func (v *NullableSchemaHashAlgorithm) Set(val *SchemaHashAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaHashAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaHashAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaHashAlgorithm(val *SchemaHashAlgorithm) *NullableSchemaHashAlgorithm {
	return &NullableSchemaHashAlgorithm{value: val, isSet: true}
}

func (v NullableSchemaHashAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaHashAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

