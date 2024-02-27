# OriginPoolOriginServerSegmentName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | Pointer to **string** | x-displayName: \&quot;DNS Name\&quot; x-example: \&quot;value\&quot; x-required DNS Name | [optional] 
**RefreshInterval** | Pointer to **int64** | x-displayName: \&quot;DNS Refresh interval\&quot; x-example: \&quot;20\&quot; Interval for DNS refresh in seconds. Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767 | [optional] 
**Segment** | Pointer to [**SchemaviewsObjectRefType**](SchemaviewsObjectRefType.md) |  | [optional] 
**SiteLocator** | Pointer to [**ViewsSiteRegionLocator**](ViewsSiteRegionLocator.md) |  | [optional] 

## Methods

### NewOriginPoolOriginServerSegmentName

`func NewOriginPoolOriginServerSegmentName() *OriginPoolOriginServerSegmentName`

NewOriginPoolOriginServerSegmentName instantiates a new OriginPoolOriginServerSegmentName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerSegmentNameWithDefaults

`func NewOriginPoolOriginServerSegmentNameWithDefaults() *OriginPoolOriginServerSegmentName`

NewOriginPoolOriginServerSegmentNameWithDefaults instantiates a new OriginPoolOriginServerSegmentName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *OriginPoolOriginServerSegmentName) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *OriginPoolOriginServerSegmentName) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *OriginPoolOriginServerSegmentName) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *OriginPoolOriginServerSegmentName) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *OriginPoolOriginServerSegmentName) GetRefreshInterval() int64`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *OriginPoolOriginServerSegmentName) GetRefreshIntervalOk() (*int64, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *OriginPoolOriginServerSegmentName) SetRefreshInterval(v int64)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *OriginPoolOriginServerSegmentName) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetSegment

`func (o *OriginPoolOriginServerSegmentName) GetSegment() SchemaviewsObjectRefType`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *OriginPoolOriginServerSegmentName) GetSegmentOk() (*SchemaviewsObjectRefType, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *OriginPoolOriginServerSegmentName) SetSegment(v SchemaviewsObjectRefType)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *OriginPoolOriginServerSegmentName) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetSiteLocator

`func (o *OriginPoolOriginServerSegmentName) GetSiteLocator() ViewsSiteRegionLocator`

GetSiteLocator returns the SiteLocator field if non-nil, zero value otherwise.

### GetSiteLocatorOk

`func (o *OriginPoolOriginServerSegmentName) GetSiteLocatorOk() (*ViewsSiteRegionLocator, bool)`

GetSiteLocatorOk returns a tuple with the SiteLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteLocator

`func (o *OriginPoolOriginServerSegmentName) SetSiteLocator(v ViewsSiteRegionLocator)`

SetSiteLocator sets SiteLocator field to given value.

### HasSiteLocator

`func (o *OriginPoolOriginServerSegmentName) HasSiteLocator() bool`

HasSiteLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


