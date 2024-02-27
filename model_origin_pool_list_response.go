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

// checks if the OriginPoolListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginPoolListResponse{}

// OriginPoolListResponse This is the output message of 'List' RPC.
type OriginPoolListResponse struct {
	//  Errors(if any) while listing items from collection
	Errors []SchemaErrorType `json:"errors,omitempty"`
	//  items represents the collection in response
	Items []OriginPoolListResponseItem `json:"items,omitempty"`
}

// NewOriginPoolListResponse instantiates a new OriginPoolListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginPoolListResponse() *OriginPoolListResponse {
	this := OriginPoolListResponse{}
	return &this
}

// NewOriginPoolListResponseWithDefaults instantiates a new OriginPoolListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginPoolListResponseWithDefaults() *OriginPoolListResponse {
	this := OriginPoolListResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *OriginPoolListResponse) GetErrors() []SchemaErrorType {
	if o == nil || IsNil(o.Errors) {
		var ret []SchemaErrorType
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolListResponse) GetErrorsOk() ([]SchemaErrorType, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *OriginPoolListResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SchemaErrorType and assigns it to the Errors field.
func (o *OriginPoolListResponse) SetErrors(v []SchemaErrorType) {
	o.Errors = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OriginPoolListResponse) GetItems() []OriginPoolListResponseItem {
	if o == nil || IsNil(o.Items) {
		var ret []OriginPoolListResponseItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolListResponse) GetItemsOk() ([]OriginPoolListResponseItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OriginPoolListResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OriginPoolListResponseItem and assigns it to the Items field.
func (o *OriginPoolListResponse) SetItems(v []OriginPoolListResponseItem) {
	o.Items = v
}

func (o OriginPoolListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginPoolListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableOriginPoolListResponse struct {
	value *OriginPoolListResponse
	isSet bool
}

func (v NullableOriginPoolListResponse) Get() *OriginPoolListResponse {
	return v.value
}

func (v *NullableOriginPoolListResponse) Set(val *OriginPoolListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginPoolListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginPoolListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginPoolListResponse(val *OriginPoolListResponse) *NullableOriginPoolListResponse {
	return &NullableOriginPoolListResponse{value: val, isSet: true}
}

func (v NullableOriginPoolListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginPoolListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

