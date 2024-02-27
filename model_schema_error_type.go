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

// checks if the SchemaErrorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaErrorType{}

// SchemaErrorType Information about a error in API operation
type SchemaErrorType struct {
	Code *SchemaErrorCode `json:"code,omitempty"`
	ErrorObj *ProtobufAny `json:"error_obj,omitempty"`
	//  A human readable string of the error  Example: ` \"value\"`
	Message *string `json:"message,omitempty"`
}

// NewSchemaErrorType instantiates a new SchemaErrorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaErrorType() *SchemaErrorType {
	this := SchemaErrorType{}
	var code SchemaErrorCode = EOK
	this.Code = &code
	return &this
}

// NewSchemaErrorTypeWithDefaults instantiates a new SchemaErrorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaErrorTypeWithDefaults() *SchemaErrorType {
	this := SchemaErrorType{}
	var code SchemaErrorCode = EOK
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SchemaErrorType) GetCode() SchemaErrorCode {
	if o == nil || IsNil(o.Code) {
		var ret SchemaErrorCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaErrorType) GetCodeOk() (*SchemaErrorCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SchemaErrorType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given SchemaErrorCode and assigns it to the Code field.
func (o *SchemaErrorType) SetCode(v SchemaErrorCode) {
	o.Code = &v
}

// GetErrorObj returns the ErrorObj field value if set, zero value otherwise.
func (o *SchemaErrorType) GetErrorObj() ProtobufAny {
	if o == nil || IsNil(o.ErrorObj) {
		var ret ProtobufAny
		return ret
	}
	return *o.ErrorObj
}

// GetErrorObjOk returns a tuple with the ErrorObj field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaErrorType) GetErrorObjOk() (*ProtobufAny, bool) {
	if o == nil || IsNil(o.ErrorObj) {
		return nil, false
	}
	return o.ErrorObj, true
}

// HasErrorObj returns a boolean if a field has been set.
func (o *SchemaErrorType) HasErrorObj() bool {
	if o != nil && !IsNil(o.ErrorObj) {
		return true
	}

	return false
}

// SetErrorObj gets a reference to the given ProtobufAny and assigns it to the ErrorObj field.
func (o *SchemaErrorType) SetErrorObj(v ProtobufAny) {
	o.ErrorObj = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SchemaErrorType) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaErrorType) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SchemaErrorType) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SchemaErrorType) SetMessage(v string) {
	o.Message = &v
}

func (o SchemaErrorType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaErrorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.ErrorObj) {
		toSerialize["error_obj"] = o.ErrorObj
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableSchemaErrorType struct {
	value *SchemaErrorType
	isSet bool
}

func (v NullableSchemaErrorType) Get() *SchemaErrorType {
	return v.value
}

func (v *NullableSchemaErrorType) Set(val *SchemaErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaErrorType(val *SchemaErrorType) *NullableSchemaErrorType {
	return &NullableSchemaErrorType{value: val, isSet: true}
}

func (v NullableSchemaErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


