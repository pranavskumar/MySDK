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

// checks if the SchemaStatusType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaStatusType{}

// SchemaStatusType Status is a return value for calls that don't return other objects.
type SchemaStatusType struct {
	//  Suggested HTTP return code for this status, 0 if not set.  Example: ` \"0\"`
	Code *int32 `json:"code,omitempty"`
	//  A human-readable description of why this operation is in the  \"Failure\" status. If this value is empty there  is no information available.  Example: ` \"value\"`
	Reason *string `json:"reason,omitempty"`
	//  Status of the operation.  One of: \"Success\" or \"Failure\".  Example: ` \"value\"`
	Status *string `json:"status,omitempty"`
}

// NewSchemaStatusType instantiates a new SchemaStatusType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaStatusType() *SchemaStatusType {
	this := SchemaStatusType{}
	return &this
}

// NewSchemaStatusTypeWithDefaults instantiates a new SchemaStatusType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaStatusTypeWithDefaults() *SchemaStatusType {
	this := SchemaStatusType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SchemaStatusType) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaStatusType) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SchemaStatusType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *SchemaStatusType) SetCode(v int32) {
	o.Code = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SchemaStatusType) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaStatusType) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SchemaStatusType) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SchemaStatusType) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SchemaStatusType) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaStatusType) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SchemaStatusType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SchemaStatusType) SetStatus(v string) {
	o.Status = &v
}

func (o SchemaStatusType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaStatusType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSchemaStatusType struct {
	value *SchemaStatusType
	isSet bool
}

func (v NullableSchemaStatusType) Get() *SchemaStatusType {
	return v.value
}

func (v *NullableSchemaStatusType) Set(val *SchemaStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaStatusType(val *SchemaStatusType) *NullableSchemaStatusType {
	return &NullableSchemaStatusType{value: val, isSet: true}
}

func (v NullableSchemaStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


