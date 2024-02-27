# OriginPoolOriginServerSegmentIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | x-displayName: \&quot;IP\&quot; x-example: \&quot;8.8.8.8\&quot; Private IPV4 address | [optional] 
**Ipv6** | Pointer to **string** | x-displayName: \&quot;IP6\&quot; x-example: \&quot;2001::10\&quot; Private IPV6 address | [optional] 
**Segment** | Pointer to [**SchemaviewsObjectRefType**](SchemaviewsObjectRefType.md) |  | [optional] 
**SiteLocator** | Pointer to [**ViewsSiteRegionLocator**](ViewsSiteRegionLocator.md) |  | [optional] 

## Methods

### NewOriginPoolOriginServerSegmentIP

`func NewOriginPoolOriginServerSegmentIP() *OriginPoolOriginServerSegmentIP`

NewOriginPoolOriginServerSegmentIP instantiates a new OriginPoolOriginServerSegmentIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerSegmentIPWithDefaults

`func NewOriginPoolOriginServerSegmentIPWithDefaults() *OriginPoolOriginServerSegmentIP`

NewOriginPoolOriginServerSegmentIPWithDefaults instantiates a new OriginPoolOriginServerSegmentIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *OriginPoolOriginServerSegmentIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OriginPoolOriginServerSegmentIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OriginPoolOriginServerSegmentIP) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OriginPoolOriginServerSegmentIP) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *OriginPoolOriginServerSegmentIP) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *OriginPoolOriginServerSegmentIP) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *OriginPoolOriginServerSegmentIP) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *OriginPoolOriginServerSegmentIP) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetSegment

`func (o *OriginPoolOriginServerSegmentIP) GetSegment() SchemaviewsObjectRefType`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *OriginPoolOriginServerSegmentIP) GetSegmentOk() (*SchemaviewsObjectRefType, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *OriginPoolOriginServerSegmentIP) SetSegment(v SchemaviewsObjectRefType)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *OriginPoolOriginServerSegmentIP) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetSiteLocator

`func (o *OriginPoolOriginServerSegmentIP) GetSiteLocator() ViewsSiteRegionLocator`

GetSiteLocator returns the SiteLocator field if non-nil, zero value otherwise.

### GetSiteLocatorOk

`func (o *OriginPoolOriginServerSegmentIP) GetSiteLocatorOk() (*ViewsSiteRegionLocator, bool)`

GetSiteLocatorOk returns a tuple with the SiteLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteLocator

`func (o *OriginPoolOriginServerSegmentIP) SetSiteLocator(v ViewsSiteRegionLocator)`

SetSiteLocator sets SiteLocator field to given value.

### HasSiteLocator

`func (o *OriginPoolOriginServerSegmentIP) HasSiteLocator() bool`

HasSiteLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


