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

// checks if the OriginPoolOriginServerSegmentName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginPoolOriginServerSegmentName{}

// OriginPoolOriginServerSegmentName x-displayName: \"DNS Name of Origin Server in Segment on given Sites\" Specify origin server with DNS name in Segment on given Site
type OriginPoolOriginServerSegmentName struct {
	// x-displayName: \"DNS Name\" x-example: \"value\" x-required DNS Name
	DnsName *string `json:"dns_name,omitempty"`
	// x-displayName: \"DNS Refresh interval\" x-example: \"20\" Interval for DNS refresh in seconds. Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767
	RefreshInterval *int64 `json:"refresh_interval,omitempty"`
	Segment *SchemaviewsObjectRefType `json:"segment,omitempty"`
	SiteLocator *ViewsSiteRegionLocator `json:"site_locator,omitempty"`
}

// NewOriginPoolOriginServerSegmentName instantiates a new OriginPoolOriginServerSegmentName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginPoolOriginServerSegmentName() *OriginPoolOriginServerSegmentName {
	this := OriginPoolOriginServerSegmentName{}
	return &this
}

// NewOriginPoolOriginServerSegmentNameWithDefaults instantiates a new OriginPoolOriginServerSegmentName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginPoolOriginServerSegmentNameWithDefaults() *OriginPoolOriginServerSegmentName {
	this := OriginPoolOriginServerSegmentName{}
	return &this
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *OriginPoolOriginServerSegmentName) GetDnsName() string {
	if o == nil || IsNil(o.DnsName) {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerSegmentName) GetDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsName) {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *OriginPoolOriginServerSegmentName) HasDnsName() bool {
	if o != nil && !IsNil(o.DnsName) {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *OriginPoolOriginServerSegmentName) SetDnsName(v string) {
	o.DnsName = &v
}

// GetRefreshInterval returns the RefreshInterval field value if set, zero value otherwise.
func (o *OriginPoolOriginServerSegmentName) GetRefreshInterval() int64 {
	if o == nil || IsNil(o.RefreshInterval) {
		var ret int64
		return ret
	}
	return *o.RefreshInterval
}

// GetRefreshIntervalOk returns a tuple with the RefreshInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerSegmentName) GetRefreshIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshInterval) {
		return nil, false
	}
	return o.RefreshInterval, true
}

// HasRefreshInterval returns a boolean if a field has been set.
func (o *OriginPoolOriginServerSegmentName) HasRefreshInterval() bool {
	if o != nil && !IsNil(o.RefreshInterval) {
		return true
	}

	return false
}

// SetRefreshInterval gets a reference to the given int64 and assigns it to the RefreshInterval field.
func (o *OriginPoolOriginServerSegmentName) SetRefreshInterval(v int64) {
	o.RefreshInterval = &v
}

// GetSegment returns the Segment field value if set, zero value otherwise.
func (o *OriginPoolOriginServerSegmentName) GetSegment() SchemaviewsObjectRefType {
	if o == nil || IsNil(o.Segment) {
		var ret SchemaviewsObjectRefType
		return ret
	}
	return *o.Segment
}

// GetSegmentOk returns a tuple with the Segment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerSegmentName) GetSegmentOk() (*SchemaviewsObjectRefType, bool) {
	if o == nil || IsNil(o.Segment) {
		return nil, false
	}
	return o.Segment, true
}

// HasSegment returns a boolean if a field has been set.
func (o *OriginPoolOriginServerSegmentName) HasSegment() bool {
	if o != nil && !IsNil(o.Segment) {
		return true
	}

	return false
}

// SetSegment gets a reference to the given SchemaviewsObjectRefType and assigns it to the Segment field.
func (o *OriginPoolOriginServerSegmentName) SetSegment(v SchemaviewsObjectRefType) {
	o.Segment = &v
}

// GetSiteLocator returns the SiteLocator field value if set, zero value otherwise.
func (o *OriginPoolOriginServerSegmentName) GetSiteLocator() ViewsSiteRegionLocator {
	if o == nil || IsNil(o.SiteLocator) {
		var ret ViewsSiteRegionLocator
		return ret
	}
	return *o.SiteLocator
}

// GetSiteLocatorOk returns a tuple with the SiteLocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginPoolOriginServerSegmentName) GetSiteLocatorOk() (*ViewsSiteRegionLocator, bool) {
	if o == nil || IsNil(o.SiteLocator) {
		return nil, false
	}
	return o.SiteLocator, true
}

// HasSiteLocator returns a boolean if a field has been set.
func (o *OriginPoolOriginServerSegmentName) HasSiteLocator() bool {
	if o != nil && !IsNil(o.SiteLocator) {
		return true
	}

	return false
}

// SetSiteLocator gets a reference to the given ViewsSiteRegionLocator and assigns it to the SiteLocator field.
func (o *OriginPoolOriginServerSegmentName) SetSiteLocator(v ViewsSiteRegionLocator) {
	o.SiteLocator = &v
}

func (o OriginPoolOriginServerSegmentName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginPoolOriginServerSegmentName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnsName) {
		toSerialize["dns_name"] = o.DnsName
	}
	if !IsNil(o.RefreshInterval) {
		toSerialize["refresh_interval"] = o.RefreshInterval
	}
	if !IsNil(o.Segment) {
		toSerialize["segment"] = o.Segment
	}
	if !IsNil(o.SiteLocator) {
		toSerialize["site_locator"] = o.SiteLocator
	}
	return toSerialize, nil
}

type NullableOriginPoolOriginServerSegmentName struct {
	value *OriginPoolOriginServerSegmentName
	isSet bool
}

func (v NullableOriginPoolOriginServerSegmentName) Get() *OriginPoolOriginServerSegmentName {
	return v.value
}

func (v *NullableOriginPoolOriginServerSegmentName) Set(val *OriginPoolOriginServerSegmentName) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginPoolOriginServerSegmentName) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginPoolOriginServerSegmentName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginPoolOriginServerSegmentName(val *OriginPoolOriginServerSegmentName) *NullableOriginPoolOriginServerSegmentName {
	return &NullableOriginPoolOriginServerSegmentName{value: val, isSet: true}
}

func (v NullableOriginPoolOriginServerSegmentName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginPoolOriginServerSegmentName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


