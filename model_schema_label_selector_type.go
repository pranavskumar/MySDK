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

// checks if the SchemaLabelSelectorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaLabelSelectorType{}

// SchemaLabelSelectorType x-displayName: \"Label Selector\" This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expresssions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects. Label selector is immutable. expressions is a list of strings of label selection expression. Each string has \",\" separated values which are \"AND\" and all strings are logically \"OR\". BNF for expression string <selector-syntax>         ::= <requirement> | <requirement> \",\" <selector-syntax> <requirement>             ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ] <set-based-restriction>   ::= \"\" | <inclusion-exclusion> <value-set> <inclusion-exclusion>     ::= <inclusion> | <exclusion> <exclusion>               ::= \"notin\" <inclusion>               ::= \"in\" <value-set>               ::= \"(\" <values> \")\" <values>                  ::= VALUE | VALUE \",\" <values> <exact-match-restriction> ::= [\"=\"|\"==\"|\"!=\"] VALUE
type SchemaLabelSelectorType struct {
	// x-displayName: \"Selector Expression\" x-required x-example: \"region in (us-west1, us-west2),tier in (staging)\" expressions contains the kubernetes style label expression for selections.
	Expressions []string `json:"expressions,omitempty"`
}

// NewSchemaLabelSelectorType instantiates a new SchemaLabelSelectorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaLabelSelectorType() *SchemaLabelSelectorType {
	this := SchemaLabelSelectorType{}
	return &this
}

// NewSchemaLabelSelectorTypeWithDefaults instantiates a new SchemaLabelSelectorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaLabelSelectorTypeWithDefaults() *SchemaLabelSelectorType {
	this := SchemaLabelSelectorType{}
	return &this
}

// GetExpressions returns the Expressions field value if set, zero value otherwise.
func (o *SchemaLabelSelectorType) GetExpressions() []string {
	if o == nil || IsNil(o.Expressions) {
		var ret []string
		return ret
	}
	return o.Expressions
}

// GetExpressionsOk returns a tuple with the Expressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaLabelSelectorType) GetExpressionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Expressions) {
		return nil, false
	}
	return o.Expressions, true
}

// HasExpressions returns a boolean if a field has been set.
func (o *SchemaLabelSelectorType) HasExpressions() bool {
	if o != nil && !IsNil(o.Expressions) {
		return true
	}

	return false
}

// SetExpressions gets a reference to the given []string and assigns it to the Expressions field.
func (o *SchemaLabelSelectorType) SetExpressions(v []string) {
	o.Expressions = v
}

func (o SchemaLabelSelectorType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaLabelSelectorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expressions) {
		toSerialize["expressions"] = o.Expressions
	}
	return toSerialize, nil
}

type NullableSchemaLabelSelectorType struct {
	value *SchemaLabelSelectorType
	isSet bool
}

func (v NullableSchemaLabelSelectorType) Get() *SchemaLabelSelectorType {
	return v.value
}

func (v *NullableSchemaLabelSelectorType) Set(val *SchemaLabelSelectorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaLabelSelectorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaLabelSelectorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaLabelSelectorType(val *SchemaLabelSelectorType) *NullableSchemaLabelSelectorType {
	return &NullableSchemaLabelSelectorType{value: val, isSet: true}
}

func (v NullableSchemaLabelSelectorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaLabelSelectorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

