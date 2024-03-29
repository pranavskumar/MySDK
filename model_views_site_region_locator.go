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

// checks if the ViewsSiteRegionLocator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewsSiteRegionLocator{}

// ViewsSiteRegionLocator x-displayName: \"Select Site or Virtual Site or Cloud Edge\" This message defines reference to site or virtual site or a cloud-re-region object
type ViewsSiteRegionLocator struct {
	CloudReRegion *SchemaviewsObjectRefType `json:"cloud_re_region,omitempty"`
	Site *SchemaviewsObjectRefType `json:"site,omitempty"`
	VirtualSite *SchemaviewsObjectRefType `json:"virtual_site,omitempty"`
}

// NewViewsSiteRegionLocator instantiates a new ViewsSiteRegionLocator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewsSiteRegionLocator() *ViewsSiteRegionLocator {
	this := ViewsSiteRegionLocator{}
	return &this
}

// NewViewsSiteRegionLocatorWithDefaults instantiates a new ViewsSiteRegionLocator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewsSiteRegionLocatorWithDefaults() *ViewsSiteRegionLocator {
	this := ViewsSiteRegionLocator{}
	return &this
}

// GetCloudReRegion returns the CloudReRegion field value if set, zero value otherwise.
func (o *ViewsSiteRegionLocator) GetCloudReRegion() SchemaviewsObjectRefType {
	if o == nil || IsNil(o.CloudReRegion) {
		var ret SchemaviewsObjectRefType
		return ret
	}
	return *o.CloudReRegion
}

// GetCloudReRegionOk returns a tuple with the CloudReRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsSiteRegionLocator) GetCloudReRegionOk() (*SchemaviewsObjectRefType, bool) {
	if o == nil || IsNil(o.CloudReRegion) {
		return nil, false
	}
	return o.CloudReRegion, true
}

// HasCloudReRegion returns a boolean if a field has been set.
func (o *ViewsSiteRegionLocator) HasCloudReRegion() bool {
	if o != nil && !IsNil(o.CloudReRegion) {
		return true
	}

	return false
}

// SetCloudReRegion gets a reference to the given SchemaviewsObjectRefType and assigns it to the CloudReRegion field.
func (o *ViewsSiteRegionLocator) SetCloudReRegion(v SchemaviewsObjectRefType) {
	o.CloudReRegion = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *ViewsSiteRegionLocator) GetSite() SchemaviewsObjectRefType {
	if o == nil || IsNil(o.Site) {
		var ret SchemaviewsObjectRefType
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsSiteRegionLocator) GetSiteOk() (*SchemaviewsObjectRefType, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ViewsSiteRegionLocator) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given SchemaviewsObjectRefType and assigns it to the Site field.
func (o *ViewsSiteRegionLocator) SetSite(v SchemaviewsObjectRefType) {
	o.Site = &v
}

// GetVirtualSite returns the VirtualSite field value if set, zero value otherwise.
func (o *ViewsSiteRegionLocator) GetVirtualSite() SchemaviewsObjectRefType {
	if o == nil || IsNil(o.VirtualSite) {
		var ret SchemaviewsObjectRefType
		return ret
	}
	return *o.VirtualSite
}

// GetVirtualSiteOk returns a tuple with the VirtualSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewsSiteRegionLocator) GetVirtualSiteOk() (*SchemaviewsObjectRefType, bool) {
	if o == nil || IsNil(o.VirtualSite) {
		return nil, false
	}
	return o.VirtualSite, true
}

// HasVirtualSite returns a boolean if a field has been set.
func (o *ViewsSiteRegionLocator) HasVirtualSite() bool {
	if o != nil && !IsNil(o.VirtualSite) {
		return true
	}

	return false
}

// SetVirtualSite gets a reference to the given SchemaviewsObjectRefType and assigns it to the VirtualSite field.
func (o *ViewsSiteRegionLocator) SetVirtualSite(v SchemaviewsObjectRefType) {
	o.VirtualSite = &v
}

func (o ViewsSiteRegionLocator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewsSiteRegionLocator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudReRegion) {
		toSerialize["cloud_re_region"] = o.CloudReRegion
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.VirtualSite) {
		toSerialize["virtual_site"] = o.VirtualSite
	}
	return toSerialize, nil
}

type NullableViewsSiteRegionLocator struct {
	value *ViewsSiteRegionLocator
	isSet bool
}

func (v NullableViewsSiteRegionLocator) Get() *ViewsSiteRegionLocator {
	return v.value
}

func (v *NullableViewsSiteRegionLocator) Set(val *ViewsSiteRegionLocator) {
	v.value = val
	v.isSet = true
}

func (v NullableViewsSiteRegionLocator) IsSet() bool {
	return v.isSet
}

func (v *NullableViewsSiteRegionLocator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewsSiteRegionLocator(val *ViewsSiteRegionLocator) *NullableViewsSiteRegionLocator {
	return &NullableViewsSiteRegionLocator{value: val, isSet: true}
}

func (v NullableViewsSiteRegionLocator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewsSiteRegionLocator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


